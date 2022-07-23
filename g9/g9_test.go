package g9

import (
	"testing"
)

func TestSetDuration(t *testing.T) {
	NoteDuration.SetDuration(Quarter, 140)
	for k, v := range NoteDuration {
		t.Logf("Note: %v, \t\tDuration: %v", k, v.Duration)
	}
}
