// An interface defines a set of method signatures that a type must implement.
// If a type has all the methods required by an interface, it implicitly implements that interface.

package main
import ("fmt" ; "math")

// define interface
type shape interface {
	area() float64
}

// define rectangle
type rectangle struct {
	l, w float64
}
// implement area() of shape interface for rectangle
func (r *rectangle) area() float64 {
	return r.l * r.w
}

//define circle
type circle struct {
	r float64
}

// implement area() of shape interface for circle
func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

// variadic function for adding areas
func totalArea (shapes ...shape) float64{
	var total float64
	for _, sh := range shapes {
		total += sh.area()
	}
	return total
}

// define multishape (a shape made of multishape)
type multishape struct {
	shapes []shape
}

// multishape implementing shape
func (m *multishape) area() float64{
	var total float64
	for _, value := range m.shapes {
		total += value.area()
	}
	return total
}

//main function
func main(){
	r := &rectangle{l : 5, w : 4}
	c := &circle{r : 3}
	fmt.Println(r.area(), c.area())
	fmt.Println("total area of rectangle and circle using totalArea() is", totalArea(r, c))

	multi := &multishape{
		shapes : []shape{r, c},
	}
	fmt.Println("Area of multishape is", multi.area())
}