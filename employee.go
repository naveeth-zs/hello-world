package employee

import "fmt"

type Employee struct {
	name string
	age  int
}

func checkEmployee(details Employee) (check bool, Details Employee) {
	if details.age < 22 {
		return
	} else {
		return true, details
	}
}

func main() {
	a := Employee{name: "xyz", age: 23}
	b := Employee{name: "abc", age: 20}
	fmt.Println(checkEmployee(a))
	fmt.Println(checkEmployee(b))
}
