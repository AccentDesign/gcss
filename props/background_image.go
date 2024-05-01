package props

import (
	"fmt"
	"strings"
)

type (
	BackgroundImage struct {
		Value string
	}
	BackgroundImageLayer string
)

func (b BackgroundImage) String() string {
	return b.Value
}

func BackgroundImageLayers(layers ...BackgroundImageLayer) BackgroundImage {
	ls := make([]string, len(layers))
	for i, layer := range layers {
		ls[i] = string(layer)
	}
	return BackgroundImage{Value: strings.Join(ls, ",")}
}

func BackgroundImageLinearGradient(segments ...string) BackgroundImageLayer {
	return BackgroundImageLayer(fmt.Sprintf("linear-gradient(%s)", strings.Join(segments, ",")))
}

func BackgroundImageURL(url string) BackgroundImageLayer {
	return BackgroundImageLayer(fmt.Sprintf(`url("%s")`, url))
}

func BackgroundImageInherit() BackgroundImage {
	return BackgroundImage{Value: "inherit"}
}

func BackgroundImageInitial() BackgroundImage {
	return BackgroundImage{Value: "initial"}
}

func BackgroundImageRevert() BackgroundImage {
	return BackgroundImage{Value: "revert"}
}

func BackgroundImageRevertLayer() BackgroundImage {
	return BackgroundImage{Value: "revert-layer"}
}

func BackgroundImageUnset() BackgroundImage {
	return BackgroundImage{Value: "unset"}
}
