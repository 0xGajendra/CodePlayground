package main
import "fmt"

//a stringer is a type that can describe itself as string. The fmt package and many more look at this interface to print values. If a type implements the String() method, it is considered a stringer.

type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
}