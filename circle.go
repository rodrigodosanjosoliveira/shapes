package shapes

import "fmt"

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%0.2f)", c.Radius)
}

// Area function
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
