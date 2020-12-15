package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	type pair struct {
		left  interface{}
		right interface{}
	}

	type expected struct {
		removed interface{}
		added   interface{}
	}

	tests := []struct {
		name     string
		in       pair
		expected expected
	}{
		// Plain

		{"plain - left nil, right nil",
			pair{nil, nil},
			expected{removed: nil, added: nil},
		},

		{"plain - left nil, right set",
			pair{nil, &plainStruct{}},
			expected{removed: (*plainStruct)(nil), added: &plainStruct{}},
		},

		{"plain - left set, right nil",
			pair{&plainStruct{}, nil},
			expected{removed: &plainStruct{}, added: (*plainStruct)(nil)},
		},

		{"plain - equal",
			pair{&plainStruct{}, &plainStruct{}},
			expected{removed: (*plainStruct)(nil), added: (*plainStruct)(nil)},
		},

		{"plain - diff",
			pair{&plainStruct{field: 1}, &plainStruct{field: 2}},
			expected{removed: &plainStruct{field: 1}, added: &plainStruct{field: 2}},
		},

		// Slice

		{"slice - equal both nil",
			pair{[]string(nil), []string(nil)},
			expected{
				removed: []string(nil),
				added:   []string(nil),
			},
		},

		{"slice - equal both length 0",
			pair{[]string{}, []string{}},
			expected{
				removed: []string(nil),
				added:   []string(nil),
			},
		},

		{"slice - diff both length 1",
			pair{[]string{"a"}, []string{"b"}},
			expected{
				removed: []string{"a"},
				added:   []string{"b"},
			},
		},

		{"slice - diff both length 2 re-ordered",
			pair{[]string{"a", "b"}, []string{"b", "a"}},
			expected{
				removed: []string{"a", "b"},
				added:   []string{"b", "a"},
			},
		},

		{"slice - diff left is longer than right, all different",
			pair{[]string{"a", "b"}, []string{"c"}},
			expected{
				removed: []string{"a", "b"},
				added:   []string{"c"},
			},
		},

		{"slice - diff left is longer than right, some equals",
			pair{[]string{"a", "b"}, []string{"a"}},
			expected{
				removed: []string{"b"},
				added:   []string(nil),
			},
		},

		{"slice - diff left is smaller than right, all different",
			pair{[]string{"a"}, []string{"b", "c"}},
			expected{
				removed: []string{"a"},
				added:   []string{"b", "c"},
			},
		},

		{"slice - diff left is smaller than right, some equals",
			pair{[]string{"a"}, []string{"a", "b"}},
			expected{
				removed: []string(nil),
				added:   []string{"b"},
			},
		},

		// Full

		{"full - everything diff",
			pair{
				&topStruct{
					literal: "a",
					pointer: &plainStruct{field: 1},
					array:   []string{"a", "b"},
					child:   &childStruct{literal: "1", pointer: &plainStruct{field: 10}, array: []string{"1", "2"}},
				},
				&topStruct{
					literal: "b",
					pointer: &plainStruct{field: 2},
					array:   []string{"b", "c"},
					child:   &childStruct{literal: "2", pointer: &plainStruct{field: 20}, array: []string{"2"}},
				},
			},
			expected{
				removed: &topStruct{
					literal: "a",
					pointer: &plainStruct{field: 1},
					array:   []string{"a", "b"},
					child:   &childStruct{literal: "1", pointer: &plainStruct{field: 10}, array: []string{"1", "2"}},
				},
				added: &topStruct{
					literal: "b",
					pointer: &plainStruct{field: 2},
					array:   []string{"b", "c"},
					child:   &childStruct{literal: "2", pointer: &plainStruct{field: 20}, array: []string{"2"}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			removed, added := Diff(test.in.left, test.in.right)
			assert.Equal(t, test.expected.removed, removed, "removed set is different")
			assert.Equal(t, test.expected.added, added, "added set is different")
		})
	}
}

type topStruct struct {
	literal string
	pointer *plainStruct
	array   []string
	child   *childStruct
}

func (s *topStruct) Diff(rightRaw interface{}) (interface{}, interface{}) {
	left := s
	right := rightRaw.(*topStruct)

	removed := &topStruct{}
	added := &topStruct{}
	if left.literal != right.literal {
		removed.literal = left.literal
		added.literal = right.literal
	}

	removedPointer, addedPointer := Diff(left.pointer, right.pointer)
	removed.pointer = removedPointer.(*plainStruct)
	added.pointer = addedPointer.(*plainStruct)

	removedArray, addedArray := Diff(left.array, right.array)
	removed.array = removedArray.([]string)
	added.array = addedArray.([]string)

	removedChild, addedChild := Diff(left.child, right.child)
	removed.child = removedChild.(*childStruct)
	added.child = addedChild.(*childStruct)

	return removed, added
}

type childStruct struct {
	literal string
	pointer *plainStruct
	array   []string
}

func (s *childStruct) Diff(rightRaw interface{}) (interface{}, interface{}) {
	left := s
	right := rightRaw.(*childStruct)

	removed := &childStruct{}
	added := &childStruct{}
	if left.literal != right.literal {
		removed.literal = left.literal
		added.literal = right.literal
	}

	removedPointer, addedPointer := Diff(left.pointer, right.pointer)
	removed.pointer = removedPointer.(*plainStruct)
	added.pointer = addedPointer.(*plainStruct)

	removedArray, addedArray := Diff(left.array, right.array)
	removed.array = removedArray.([]string)
	added.array = addedArray.([]string)

	return removed, added
}

type plainStruct struct {
	field int
}

func (s *plainStruct) Diff(rightRaw interface{}) (interface{}, interface{}) {
	left := s
	right := rightRaw.(*plainStruct)

	removed := &plainStruct{}
	added := &plainStruct{}
	if left.field != right.field {
		removed.field = left.field
		added.field = right.field
	}

	return removed, added
}
