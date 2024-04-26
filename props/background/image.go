package background

import (
	"fmt"
	"strings"
)

type (
	Image struct {
		Value string
	}
	ImageLayer string
)

func (b Image) String() string {
	return b.Value
}

func ImageLayers(layers ...ImageLayer) Image {
	ls := make([]string, len(layers))
	for i, layer := range layers {
		ls[i] = string(layer)
	}
	return Image{Value: strings.Join(ls, ",")}
}

func ImageLinearGradient(segments ...string) ImageLayer {
	return ImageLayer(fmt.Sprintf("linear-gradient(%s)", strings.Join(segments, ",")))
}

func ImageURL(url string) ImageLayer {
	return ImageLayer(fmt.Sprintf(`url("%s")`, url))
}

func ImageInherit() Image {
	return Image{Value: "inherit"}
}

func ImageInitial() Image {
	return Image{Value: "initial"}
}

func ImageRevert() Image {
	return Image{Value: "revert"}
}

func ImageRevertLayer() Image {
	return Image{Value: "revert-layer"}
}

func ImageUnset() Image {
	return Image{Value: "unset"}
}
