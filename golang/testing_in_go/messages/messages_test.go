package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestGreetTableDriver(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result. Wanted %q, got: %q\n", s.expect, got)
		}
	}

}
func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

// func TestFailureTypes(t *testing.T) {
// 	t.Error("Error signals a failed test, but doesn't stop the rest of the test from executing")
// 	t.Fatal("Fatal will fail the test and stop its execution")
// 	t.Error("This will never print")
// }
