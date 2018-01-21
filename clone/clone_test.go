package clone

import "testing"

func TestCloneRepository(t *testing.T) {
	if err := CloneRepository("myRepo", "", false); err != nil {
		t.Fail()
	}
}
