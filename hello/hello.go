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

func nononmepty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
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

	integerSlice := make([]int, 3)

	integerSlice[0] = 1
	integerSlice[1] = 2
	integerSlice[2] = 3

	integerSlicesliced := integerSlice[0:2]

	fmt.Printf("slice cap: %d slice len: %d \n", cap(integerSlicesliced), len(integerSlicesliced))
	fmt.Printf("slice position zero: %d \n", integerSlicesliced[2])

}
