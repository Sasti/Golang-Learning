package main

import shop "github.com/sasti9/golang-learning/maps/shop"
import "fmt"
import "encoding/json"
import "os"
import "io/ioutil"

type MinimalArticle struct {
	Description string
	Netto       float32
	Variants    []string
}

var x int

func squares() func(a int) int {
	f := func(a int) int {

		x += a
		x++
		return x * x
	}
	return f
}

func sum(vals ...int) int {
	sum := 0

	for val := range vals {
		sum += val
	}

	return sum
}

func main() {

	nerfgun := shop.Article{
		Description: "Super duper Nerf-Gun.",
		Netto:       299.99,
		VAT:         0.19,
		ID:          shop.GenerateIDForArticle(),
		Variants:    []string{"Rot", "Grün"},
		Images: map[string]string{
			"redgun":   "c/redgun.jpg",
			"greengun": "c/greengun.jrp",
		},
	}

	articles := map[string]*shop.Article{
		"Nerf3000": &nerfgun,
		"Nerf1337": &shop.Article{
			Description: "An even better gun compared to the Nert3000. With this baby John will feel the revenge.",
			Netto:       399.99,
			VAT:         0.19,
			ID:          shop.GenerateIDForArticle(),
		},
	}

	fmt.Println(articles)

	for _, artciclepointer := range articles {
		fmt.Println(artciclepointer)
	}

	data, err := json.MarshalIndent(nerfgun, "", "     ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	f, err := os.Create("nerfgun.json")

	if err != nil {
		fmt.Println(err)
	}
	f.Write(data)
	f.Close()

	f, err = os.Open("nerfgun.json")
	dataFromFile, _ := ioutil.ReadAll(f)

	var minimalArticle MinimalArticle
	json.Unmarshal(dataFromFile, &minimalArticle)

	fmt.Println(minimalArticle)

	squares1 := squares()
	squares2 := squares()

	fmt.Println(squares1(1))
	fmt.Println(squares2(-1))

	sum(1, 2)

	sum(1, 2, 3)

	v := 1

	// Use a slice to implement a stack
	stack := make([]int, 0, 20)
	//stack = stack[0:1]

	fmt.Println(stack)
	stack = append(stack, v)
	fmt.Println(stack)
	top := stack[len(stack)-1]
	fmt.Println(stack)

	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	fmt.Println(top)
}
