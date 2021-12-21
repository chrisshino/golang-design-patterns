package main
import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {

}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle")
}

type RasterRenderer struct {
	Dpi int
}

func (v *RasterRenderer) RenderCircle(dpi int) {
	fmt.Println("Drawing pixels for circle of radius")
}

type Circle struct {
	renderer Renderer 
	radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}



func main() {

}