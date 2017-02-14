package main

import "fmt"

type Vector2D struct {
	x, y int
}

type Vector3D struct {
	Vector2D
	z int
}

func (v Vector2D) String() string {
	return fmt.Sprintf("X: %d, Y: %d", v.x, v.y)
}

func (v Vector3D) String() string {
	return fmt.Sprintf("%s, Z: %d", v.Vector2D, v.z)
}

type mygreatint int

func (i mygreatint) String() string {
	return fmt.Sprintf("%d", i)
}

func main() {
	var vec2 Vector2D
	vec2.x = 1
	vec2.y = 2

	var vec3 Vector3D

	fmt.Printf("Mein Vector2D sieht so aus %s. \n", vec2)
	fmt.Printf("Mein Vector3D sieht so aus %s. \n", vec3)

	var myint mygreatint = 666
	fmt.Printf("%s \n", string(myint))
}
