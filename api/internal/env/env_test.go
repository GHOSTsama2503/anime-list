package env

import (
	"testing"
)

func TestLoad(t *testing.T) {
	if IsProduction {
		t.Fatal("unexpected environment")
	}
}
