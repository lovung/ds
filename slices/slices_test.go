package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertInnerTypeTo(t *testing.T) {
	r := ConvertInnerTypeTo[int]([]any{1, 2, 3, 4})
	if fmt.Sprintf("%T", r) != "[]int" {
		t.Errorf("ConvertInnerTypeTo[int] failed")
	}
	r2 := ConvertInnerTypeTo[string]([]any{1, 2, 3, 4})
	if fmt.Sprintf("%T", r2) != "[]string" {
		t.Errorf("ConvertInnerTypeTo[string] failed")
	}
	fmt.Println(reflect.ValueOf(r2))
	t.Error("")
	r3 := ConvertInnerTypeTo[string]([]any{"a", "b", "c", "d"})
	if fmt.Sprintf("%T", r3) != "[]string" {
		t.Errorf("ConvertInnerTypeTo[string] failed")
	}
}
