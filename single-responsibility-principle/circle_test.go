package single

import "testing"

func TestNewTask(t *testing.T) {
	var testEqualFloat = func(got, want float64, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got '%f' want '%f' \n", got, want)
		}
	}

	t.Run("Create a new Circle", func(t *testing.T) {
		c := Circle{10.10}
		got := c.radius
		want := 10.10
		testEqualFloat(got, want, t)
	})
}
