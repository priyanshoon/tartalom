package utils_test

import (
	"testing"

	"tartalom/utils"
)

func TestApiKeyGenerator(t *testing.T) {
	got := utils.GenerateApi("8f22448d-12d7-475f-9c45-295ea256c6be")
	want := "OGYyMjQ0OGQtMTJkNy00NzVmLTljNDUtMjk1ZWEyNTZjNmJl"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
