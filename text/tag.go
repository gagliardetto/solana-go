package text

import (
	"reflect"
	"strings"
)

type fieldTag struct {
	Linear     bool
	Skip       bool
	Label      string
	NoTypeName bool
}

func parseFieldTag(tag reflect.StructTag) *fieldTag {
	t := &fieldTag{}
	tagStr := tag.Get("text")
	if tagStr == "" {
		return t
	}
	for _, s := range strings.Split(tagStr, ",") {
		if strings.HasPrefix(s, "linear") {
			t.Linear = true
		} else if strings.HasPrefix(s, "notype") {
			t.NoTypeName = true
		} else if s == "-" {
			t.Skip = true
		} else {
			t.Label = s
		}
	}
	return t
}
