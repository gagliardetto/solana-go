package solana

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/lunixbochs/struc"
	"go.uber.org/zap"
)

type VariantDefinition struct {
	typeIDToType map[Varuint16]reflect.Type
	typeIDToName map[Varuint16]string
	typeNameToID map[string]Varuint16
}

type VariantType struct {
	Name string
	Type interface{}
}

func NewVariantDefinition(types []VariantType) (out *VariantDefinition) {
	if len(types) < 0 {
		panic("it's not valid to create a variant definition without any types")
	}

	typeCount := len(types)
	out = &VariantDefinition{
		typeIDToType: make(map[Varuint16]reflect.Type, typeCount),
		typeIDToName: make(map[Varuint16]string, typeCount),
		typeNameToID: make(map[string]Varuint16, typeCount),
	}

	for i, typeDef := range types {
		typeID := Varuint16(i)

		// FIXME: Check how the reflect.Type is used and cache all its usage in the definition.
		//        Right now, on each Unmarshal, we re-compute some expensive stuff that can be
		//        re-used like the `typeGo.Elem()` which is always the same. It would be preferable
		//        to have those already pre-defined here so we can actually speed up the
		//        Unmarshal code.
		out.typeIDToType[typeID] = reflect.TypeOf(typeDef.Type)
		out.typeIDToName[typeID] = typeDef.Name
		out.typeNameToID[typeDef.Name] = typeID
	}

	return out
}

func (d *VariantDefinition) IDForType(impl interface{}) Varuint16 {
	for typeNum, reflectType := range d.typeIDToType {
		if reflect.TypeOf(impl) == reflectType {
			return typeNum
		}
	}
	panic(fmt.Sprintf("type %T undefined on variant definition %#v", impl, d))
}

type BaseVariant struct {
	Type Varuint16
	Impl interface{}
}

func (bv BaseVariant) Pack(p []byte, opt *struc.Options) (written int, err error) {
	if traceEnabled {
		zlog.Debug("packing variant to binary", zap.Uint16("type", uint16(bv.Type)))
	}

	w := &byteCounterWritter{Writer: bytes.NewBuffer(p)}

	err = struc.Pack(w, bv.Type)
	if err != nil {
		return 0, fmt.Errorf("pack type: %w", err)
	}

	err = struc.Pack(w, bv.Impl)
	if err != nil {
		return 0, fmt.Errorf("pack impl: %w", err)
	}

	return w.byteCount, nil
}

func (bv BaseVariant) Size(opt *struc.Options) int {
	return 0
}

func (bv BaseVariant) String() string {
	return fmt.Sprintf("%T (%d)", bv.Type, bv.Impl)
}

func (bv *BaseVariant) Unpack(def *VariantDefinition, r io.Reader, length int, opt *struc.Options) (err error) {
	if err = struc.Unpack(r, &bv.Type); err != nil {
		return
	}

	typeGo := def.typeIDToType[bv.Type]
	if typeGo == nil {
		return fmt.Errorf("no known type for type %d on %#v", bv.Type, def)
	}

	if typeGo.Kind() == reflect.Ptr {
		bv.Impl = reflect.New(typeGo.Elem()).Interface()
		if err = struc.Unpack(r, bv.Impl); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %s", bv.Type, err)
		}
	} else {
		// This is not the most optimal way of doing things for "value"
		// types (over "pointer" types) as we always allocate a new pointer
		// element, unmarshal it and then either keep the pointer type or turn
		// it into a value type.
		//
		// However, in non-reflection based code, one would do like this and
		// avoid an `new` memory allocation:
		//
		// ```
		// name := string("")
		// json.Unmarshal(data, &name)
		// ```
		//
		// This would work without a problem. In reflection code however, I
		// did not find how one can go from `reflect.Zero(typeGo)` (which is
		// the equivalence of doing `name := string("")`) and take the
		// pointer to it so it can be unmarshalled correctly.
		//
		// A played with various iteration, and nothing got it working. Maybe
		// the next step would be to explore the `unsafe` package and obtain
		// an unsafe pointer and play with it.
		value := reflect.New(typeGo)
		if err = struc.Unpack(r, value.Interface()); err != nil {
			return fmt.Errorf("unable to decode variant type %d: %s", bv.Type, err)
		}

		bv.Impl = value.Elem().Interface()
	}
	return nil
}

type byteCounterWritter struct {
	io.Writer
	byteCount int
}

func (w *byteCounterWritter) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	w.byteCount += n
	return
}
