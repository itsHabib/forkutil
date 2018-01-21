package fork

import "testing"

func TestForkRepository(t *testing.T) {
	if err := ForkRepository("myRepo"); err != nil {
		t.Fail()
	}
}
