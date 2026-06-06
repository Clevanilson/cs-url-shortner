package assert

import "testing"

func Equal[TValue comparable](t *testing.T, value1, value2 TValue) {
	if value1 != value2 {
		t.Fatalf("🔴 Expected: %v to equals %v", value1, value2)
	}
}
