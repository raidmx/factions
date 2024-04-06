package items

import (
	"image"
	"os"
)

// texture returns the texture of a custom item
func texture(id string) image.Image {
	imgFile, err := os.Open("./textures/" + id)

	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imgFile)

	if err != nil {
		panic(err)
	}

	return img.(image.Image)
}
