package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name       string
		arr1       []string
		arr2       []string
		wantResArr []string
	}{
		{
			"Empty arrays",
			[]string{},
			[]string{},
			[]string{},
		},
		{
			"No intersection",
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
			[]string{},
		},
		{
			"Full intersection",
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			"Partial intersection",
			[]string{"a", "b", "c"},
			[]string{"a", "b", "d"},
			[]string{"a", "b"},
		},
		{
			"Array of size 1",
			[]string{"a"},
			[]string{"a"},
			[]string{"a"},
		},
		{
			"Array of size 1 (no intersection)",
			[]string{"a"},
			[]string{"b"},
			[]string{},
		},
		{
			"Multiple intersected values",
			[]string{"a", "b", "c"},
			[]string{"a", "c", "d", "a", "c", "a"},
			[]string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResArr := intersection(tt.arr1, tt.arr2); !reflect.DeepEqual(gotResArr, tt.wantResArr) {
				t.Errorf("intersection() = %v, want %v", gotResArr, tt.wantResArr)
			}
		})
	}
}
