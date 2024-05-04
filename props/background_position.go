package props

import (
	"fmt"
	"strings"
)

type (
	BackgroundPosition         string
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
	return string(b)
}

func (b BackgroundPositionEdgeItem) String() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", b.Edge, b.Unit))
}

func BackgroundPositionXY(x, y Unit) BackgroundPosition {
	return BackgroundPosition(fmt.Sprintf("%s %s", x.String(), y.String()))
}

func BackgroundPositionEdgeOffset(edges ...BackgroundPositionEdgeItem) BackgroundPosition {
	ed := make([]string, len(edges))
	for i, edge := range edges {
		ed[i] = edge.String()
	}
	return BackgroundPosition(strings.Join(ed, " "))
}

func BackgroundPositions(positions ...BackgroundPosition) BackgroundPosition {
	po := make([]string, len(positions))
	for i, position := range positions {
		po[i] = string(position)
	}
	return BackgroundPosition(strings.Join(po, ","))
}

func BackgroundPositionTop() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeTop)
}

func BackgroundPositionBottom() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeBottom)
}

func BackgroundPositionLeft() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeLeft)
}

func BackgroundPositionRight() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeRight)
}

func BackgroundPositionCenter() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeCenter)
}

func BackgroundPositionTopLeft() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeTopLeft)
}

func BackgroundPositionTopRight() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeTopRight)
}

func BackgroundPositionBottomLeft() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeBottomLeft)
}

func BackgroundPositionBottomRight() BackgroundPosition {
	return BackgroundPosition(BackgroundPositionEdgeBottomRight)
}
