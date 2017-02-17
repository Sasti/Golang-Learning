package main

import (
	"fmt"
	"time"

	"github.com/sasti9/imagewebservice/image/imagestore"
)

func main() {
	// Starting the image-store application.

	now := time.Now()
	var m imagestore.Metadata

	m.Uploaddate = &now
	m.Uploadername = "don't know"

	fmt.Printf("Aktuelle Zeit %s", m.Uploaddate)
}
