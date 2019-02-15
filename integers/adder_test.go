package integers

import "testing"

func TestAdder(t *testing.T) {
	assertMessage := func(t *testing.T, expected, sum int) {
		t.Helper()
		if expected != sum {
			t.Errorf("expected '%d' but got '%d'", expected, sum)

		}
	}

	t.Run("Add two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertMessage(t, expected, sum)
	})
}
