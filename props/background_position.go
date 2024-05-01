package props

import (
	"strings"
)

type (
	BackgroundPosition struct {
		Value string
	}
	BackgroundPositionEdge     string
	BackgroundPositionEdgeItem struct {
		Edge BackgroundPositionEdge
		Unit Unit
	}
)

const (
	BackgroundPositionEdgeTop         BackgroundPositionEdge = "top"
	BackgroundPositionEdgeBottom      BackgroundPositionEdge = "bottom"
	BackgroundPositionEdgeLeft        BackgroundPositionEdge = "left"
	BackgroundPositionEdgeRight       BackgroundPositionEdge = "right"
	BackgroundPositionEdgeCenter      BackgroundPositionEdge = "center"
	BackgroundPositionEdgeTopLeft     BackgroundPositionEdge = "top left"
	BackgroundPositionEdgeTopRight    BackgroundPositionEdge = "top right"
	BackgroundPositionEdgeBottomLeft  BackgroundPositionEdge = "bottom left"
	BackgroundPositionEdgeBottomRight BackgroundPositionEdge = "bottom right"
)

func (b BackgroundPosition) String() string {
	return b.Value
}

func BackgroundPositionXY(x, y Unit) BackgroundPosition {
	return BackgroundPosition{x.String() + " " + y.String()}
}

func BackgroundPositionEdgeOffset(edges ...BackgroundPositionEdgeItem) BackgroundPosition {
	value := ""
	for _, edge := range edges {
		value += " " + string(edge.Edge)
		if edge.Unit.String() != "" {
			value += " " + edge.Unit.String()
		}
	}
	return BackgroundPosition{strings.TrimSpace(value)}
}

func BackgroundPositionMultiple(positions ...BackgroundPosition) BackgroundPosition {
	value := ""
	for i, position := range positions {
		if i > 0 {
			value += ", "
		}
		value += position.String()
	}
	return BackgroundPosition{value}
}

func BackgroundPositionTop() BackgroundPosition {
	return BackgroundPosition{"top"}
}

func BackgroundPositionBottom() BackgroundPosition {
	return BackgroundPosition{"bottom"}
}

func BackgroundPositionLeft() BackgroundPosition {
	return BackgroundPosition{"left"}
}

func BackgroundPositionRight() BackgroundPosition {
	return BackgroundPosition{"right"}
}

func BackgroundPositionCenter() BackgroundPosition {
	return BackgroundPosition{"center"}
}
