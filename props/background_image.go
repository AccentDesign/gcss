package props

import (
	"fmt"
	"strings"
)

type BackgroundImage string

func (b BackgroundImage) String() string {
	return string(b)
}

func BackgroundImages(images ...BackgroundImage) BackgroundImage {
	im := make([]string, len(images))
	for i, image := range images {
		im[i] = string(image)
	}
	return BackgroundImage(strings.Join(im, ","))
}

func BackgroundImageLinearGradient(segments ...string) BackgroundImage {
	return BackgroundImage(fmt.Sprintf("linear-gradient(%s)", strings.Join(segments, ",")))
}

func BackgroundImageURL(url string) BackgroundImage {
	return BackgroundImage(fmt.Sprintf(`url("%s")`, url))
}
