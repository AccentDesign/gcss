package props

import (
	"fmt"
	"strings"
)

type (
	BackgroundPosition     string
	BackgroundPositionEdge struct {
		Position BackgroundPosition
		Unit     Unit
	}
)

const (
	BackgroundPositionTop         BackgroundPosition = "top"
	BackgroundPositionBottom      BackgroundPosition = "bottom"
	BackgroundPositionLeft        BackgroundPosition = "left"
	BackgroundPositionRight       BackgroundPosition = "right"
	BackgroundPositionCenter      BackgroundPosition = "center"
	BackgroundPositionTopLeft     BackgroundPosition = "top left"
	BackgroundPositionTopRight    BackgroundPosition = "top right"
	BackgroundPositionBottomLeft  BackgroundPosition = "bottom left"
	BackgroundPositionBottomRight BackgroundPosition = "bottom right"
)

func (b BackgroundPosition) String() string {
	return string(b)
}

func (b BackgroundPositionEdge) String() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", b.Position, b.Unit))
}

func BackgroundPositionXY(x, y Unit) BackgroundPosition {
	return BackgroundPosition(fmt.Sprintf("%s %s", x.String(), y.String()))
}

func BackgroundPositionEdges(edges ...BackgroundPositionEdge) BackgroundPosition {
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
