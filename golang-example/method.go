package main

import "fmt"

type rect struct {
	widht, height int
}

func (r *rect) area() int {
	return r.widht * r.height

}

func (r rect) perim() int {
	return 2*r.widht + 2*r.height
}

func main() {

	r := rect{widht: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())
}