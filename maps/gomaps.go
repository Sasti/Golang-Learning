package main

import shop "github.com/sasti9/maps/shop"
import "fmt"

func main() {

	nerfgun := shop.Article{
		Description: "Super duper Nerf-Gun.",
		Netto:       299.99,
		VAT:         0.19,
		ID:          shop.GenerateIDForArticle(),
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
}
