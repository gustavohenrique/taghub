package stringutils

import "testing"

func TestSliceContains(t *testing.T) {
	arr := []string{"gustavo", "henrique"}
	isExists := SliceContains(arr, "henrique")
	if isExists == false {
		t.Errorf("Failed. Expected true but got false")
	}
}
