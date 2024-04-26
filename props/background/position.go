package background

import (
	"github.com/AccentDesign/gostyle/props"
	"strings"
)

type (
	Position struct {
		Value string
	}
	PositionEdge     string
	PositionEdgeItem struct {
		Edge PositionEdge
		Unit props.Unit
	}
)

const (
	PositionEdgeTop         PositionEdge = "top"
	PositionEdgeBottom      PositionEdge = "bottom"
	PositionEdgeLeft        PositionEdge = "left"
	PositionEdgeRight       PositionEdge = "right"
	PositionEdgeCenter      PositionEdge = "center"
	PositionEdgeTopLeft     PositionEdge = "top left"
	PositionEdgeTopRight    PositionEdge = "top right"
	PositionEdgeBottomLeft  PositionEdge = "bottom left"
	PositionEdgeBottomRight PositionEdge = "bottom right"
)

func (b Position) String() string {
	return b.Value
}

func PositionXY(x, y props.Unit) Position {
	return Position{x.String() + " " + y.String()}
}

func PositionEdgeOffset(edges ...PositionEdgeItem) Position {
	value := ""
	for _, edge := range edges {
		value += " " + string(edge.Edge)
		if edge.Unit.String() != "" {
			value += " " + edge.Unit.String()
		}
	}
	return Position{strings.TrimSpace(value)}
}

func PositionMultiple(positions ...Position) Position {
	value := ""
	for i, position := range positions {
		if i > 0 {
			value += ", "
		}
		value += position.String()
	}
	return Position{value}
}

func PositionTop() Position {
	return Position{"top"}
}

func PositionBottom() Position {
	return Position{"bottom"}
}

func PositionLeft() Position {
	return Position{"left"}
}

func PositionRight() Position {
	return Position{"right"}
}

func PositionCenter() Position {
	return Position{"center"}
}
