package diff

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
)

var diffeableInterface = reflect.TypeOf((*Diffeable)(nil)).Elem()

func Diff(left interface{}, right interface{}) (removed, added interface{}) {
	if traceEnabled {
		zlog.Debug("checking diff between elements", zap.Stringer("left", reflectType{left}), zap.Stringer("right", reflectType{right}))
	}

	if left == nil && right == nil {
		if traceEnabled {
			zlog.Debug("both end are == nil, returning them as-is")
		}

		// Hopefully types will be all right using straight received values
		return left, right
	}

	if left == nil {
		if traceEnabled {
			zlog.Debug("nil -> right, returning no removed and right added")
		}

		return reflect.Zero(reflect.TypeOf(right)).Interface(), right
	}

	if right == nil {
		if traceEnabled {
			zlog.Debug("left -> nil, returning left removed and no added")
		}

		return left, reflect.Zero(reflect.TypeOf(left)).Interface()
	}

	leftValue := reflect.ValueOf(left)
	leftType := leftValue.Type()

	rightValue := reflect.ValueOf(right)
	rightType := rightValue.Type()
	if leftType != rightType {
		panic(fmt.Errorf("type mistmatch, left != right (type %s != type %s)", leftType, rightType))
	}

	// This is costly because it means we deeply compare at each level of check. There is probably a much better
	// way to do this. We should implement a full walking reflection based diff instead that walks the whole thing
	// and allocate a removed/added struct and set the field as we go. Of course, will need to be public otherwise
	// the caller implements Diffeable and peforms the job himself. Will see, probably a good start anyway.
	if reflect.DeepEqual(left, right) {
		if traceEnabled {
			zlog.Debug("left == right, returning no removed and no added")
		}

		return reflect.Zero(leftType).Interface(), reflect.Zero(rightType).Interface()
	}

	// They are the same type, so we can check either left or right to ensure that both implements diff.Diffeable
	if leftType.Implements(diffeableInterface) {
		if traceEnabled {
			zlog.Debug("delegating to Diffeable to perform its job on struct")
		}

		return left.(Diffeable).Diff(right)
	}

	if leftValue.Kind() == reflect.Slice {
		if traceEnabled {
			zlog.Debug("performing slice compare")
		}

		return diffSlice(left, leftValue, right, rightValue)
	}

	// We know at this point that left & right are not deeply equal, not a slice and does not implement Diffeable, simply return them
	if leftType.Comparable() {
		return left, right
	}

	panic(fmt.Errorf("type incomparable, type %s is not a slice, nor a comparable and does not implement diff.Diffeable", leftType))
}

func diffSlice(left interface{}, leftValue reflect.Value, right interface{}, rightValue reflect.Value) (interface{}, interface{}) {
	leftLen := leftValue.Len()
	rightLen := rightValue.Len()

	if leftLen == 0 && rightLen == 0 {
		return nil, nil
	}

	if leftLen == 0 {
		return reflect.Zero(leftValue.Type()).Interface(), right
	}

	if rightLen == 0 {
		return left, reflect.Zero(rightValue.Type()).Interface()
	}

	removed := reflect.Zero(rightValue.Type())
	added := reflect.Zero(rightValue.Type())

	for i := 0; i < leftLen; i++ {
		if i < rightLen {
			// Both set has the same value
			leftAt := leftValue.Index(i).Interface()
			rightAt := rightValue.Index(i).Interface()

			// FIXME: Re-use Diff(...) logic that same element gives already nothing so we avoid this...
			if !reflect.DeepEqual(leftAt, rightAt) {
				removedAt, addedAt := Diff(leftAt, rightAt)

				if traceEnabled {
					zlog.Debug("slice elements at index different", zap.Int("index", i), zap.Stringer("removed", reflectType{removedAt}), zap.Stringer("added", reflectType{addedAt}))
				}

				removed = reflect.Append(removed, reflect.ValueOf(removedAt))
				added = reflect.Append(added, reflect.ValueOf(addedAt))
			}
		} else {
			// Left is bigger than right, every element here has been removed from left
			removed = reflect.Append(removed, leftValue.Index(i))
		}
	}

	// Right is bigger than left, every element after (left len - 1) has been added from right
	if rightLen > leftLen {
		for i := leftLen; i < rightLen; i++ {
			added = reflect.Append(added, rightValue.Index(i))
		}
	}

	return removed.Interface(), added.Interface()
}

// FIXME: We could most probably get rid of the Diffeable interface and diff everything ourself, should not be hard follwing
//        reflect.DeepEqual rules, probably not worth it just yet.
type Diffeable interface {
	// Diff performs the structural difference between itself (i.e. the receiver implementing the interface) which
	// we call the "left" and a "right" element returning two new structure of the same type that contains only the
	// difference between left and right. The first is the "removed" set (i.e. ) The left (receiver), the right (parameter) and the out (result) will be all of
	// the same type.
	//
	// The implementer is responsible of validating the "right" elment's type and returning the appropiate "out".
	//
	// For a given struct,
	Diff(right interface{}) (removed, added interface{})
}
