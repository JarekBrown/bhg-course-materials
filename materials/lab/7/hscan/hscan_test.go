// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f") // Currently function returns only number of open ports
	want := "foo"
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func TestGuessSingleAgain(t *testing.T) {
	got := GuessSingle("b3995d3c36af97aacd90b10cfaa7b4ff07cb44b89a594954f5228cf271211b81") // Currently function returns only number of open ports
	want := "gopherman"
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
