package shop

import "image"
import "math/rand"
import "fmt"

// Article is the main model for the articles in the shop.
type Article struct {
	ID          uint64
	Description string
	Netto       float32
	VAT         float32
	Images      map[string]image.Image
}

// CalculateBrutto calculates the brutto price based on netto and VAT
func (a *Article) CalculateBrutto() float32 {
	return a.Netto * 100 * a.VAT
}

// GenerateIDForArticle generates a random ID for an Article.
func GenerateIDForArticle() uint64 {
	articleIDPart1 := uint64(rand.Uint32())
	articleIDPart2 := uint64(rand.Uint32()) << 32

	var articleID uint64
	articleID = articleIDPart1 | articleIDPart2

	return articleID
}

// Implement stringer interface
func (a *Article) String() string {
	brutto := a.CalculateBrutto()
	articlestring := fmt.Sprintf("ArticleID: %d, Price: %f", a.ID, brutto)
	return articlestring
}
