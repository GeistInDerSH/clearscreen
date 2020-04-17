package clearscreen

import (
	"runtime"
	"testing"
)

func TestSupportedOS(t *testing.T) {
	result := IsSupportedOS()
	if !result {
		t.Errorf("%s is not in the list of supported operating systems!", runtime.GOOS)
	}
}
