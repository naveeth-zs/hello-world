package employee

import (
	"testing"
)

type output struct {
	check  bool
	person Employee
}

func TestCheckEmployee(t *testing.T) {
	tables := []struct {
		input Employee
		out   output
	}{
		{input: Employee{name: "abc", age: 21}, out: output{}},
		{input: Employee{name: "xyz", age: 24}, out: output{true, Employee{"xyz", 24}}},
	}
	for _, v := range tables {
		details, result := checkEmployee(v.input)

		if details != v.out.check {
			t.Errorf("accepted %v but got %v", v.out, details)
		}
		if result != v.out.person {
			t.Errorf("accepted %v but got %v", v.out, details)
		}
	}
}
