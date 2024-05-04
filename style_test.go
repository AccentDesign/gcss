package gcss

import (
	"bytes"
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"image/color"
	"testing"
)

func TestStyle_Empty(t *testing.T) {
	st := &Style{Selector: ".test"}
	var buf bytes.Buffer
	err := st.CSS(&buf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	css := ".test{}"
	if buf.String() != css {
		t.Errorf("expected %q, got %q", css, buf.String())
	}
}

func TestStyle_MultipleProps(t *testing.T) {
	st := &Style{Selector: ".test", Props: Props{
		BackgroundColor: props.ColorRGBA(0, 0, 0, 255),
		Height:          props.UnitPx(100),
		Width:           props.UnitPx(100),
	}}
	var buf bytes.Buffer
	err := st.CSS(&buf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	css := ".test{background-color:rgba(0,0,0,1.00);height:100px;width:100px;}"
	if buf.String() != css {
		t.Errorf("expected %q, got %q", css, buf.String())
	}
}

func TestStyle_AlignItems(t *testing.T) {
	testCases := map[props.AlignItems]string{
		props.AlignItemsStart:       "flex-start",
		props.AlignItemsEnd:         "flex-end",
		props.AlignItemsCenter:      "center",
		props.AlignItemsBaseline:    "baseline",
		props.AlignItemsStretch:     "stretch",
		props.AlignItems("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{AlignItems: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{align-items:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_Appearance(t *testing.T) {
	testCases := map[props.Appearance]string{
		props.AppearanceNone:           "none",
		props.AppearanceAuto:           "auto",
		props.AppearanceMenuListButton: "menulist-button",
		props.AppearanceTextField:      "textfield",
		props.AppearanceInherit:        "inherit",
		props.AppearanceInitial:        "initial",
		props.AppearanceRevert:         "revert",
		props.AppearanceRevertLater:    "revert-later",
		props.AppearanceUnset:          "unset",
		props.AppearanceButton:         "button",
		props.AppearanceCheckbox:       "checkbox",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Appearance: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{appearance:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BackgroundColor(t *testing.T) {
	testCases := map[props.Color]string{
		props.ColorRGBA(0, 0, 0, 255):                  "rgba(0,0,0,1.00)",
		props.ColorRGBA(255, 255, 255, 230):            "rgba(255,255,255,0.90)",
		props.ColorRGBA(255, 255, 255, 255).Alpha(230): "rgba(255,255,255,0.90)",
		props.ColorCurrentColor():                      "currentColor",
		props.ColorInherit():                           "inherit",
		props.ColorTransparent():                       "transparent",
		props.Color{Keyword: "#efefef"}:                "#efefef",
		props.Color{Keyword: "red"}:                    "red",
		props.Color{RGBA: color.RGBA{0, 0, 0, 255}}:    "rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BackgroundColor: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{background-color:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BackgroundImage(t *testing.T) {
	testCases := map[props.BackgroundImage]string{
		props.BackgroundImageURL("image.jpg"):              `url("image.jpg")`,
		props.BackgroundImageLinearGradient("red", "blue"): "linear-gradient(red,blue)",
		props.BackgroundImageInherit():                     "inherit",
		props.BackgroundImageInitial():                     "initial",
		props.BackgroundImageRevert():                      "revert",
		props.BackgroundImageRevertLayer():                 "revert-layer",
		props.BackgroundImageUnset():                       "unset",
		props.BackgroundImages(
			props.BackgroundImageLinearGradient("red", "blue"),
			props.BackgroundImageURL("image.jpg"),
		): `linear-gradient(red,blue),url("image.jpg")`,
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BackgroundImage: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{background-image:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BackgroundPosition(t *testing.T) {
	testCases := map[props.BackgroundPosition]string{
		props.BackgroundPositionXY(props.UnitPx(10), props.UnitPx(20)): "10px 20px",
		props.BackgroundPositionEdgeOffset(
			props.BackgroundPositionEdgeItem{Edge: props.BackgroundPositionEdgeTop, Unit: props.UnitPx(10)},
			props.BackgroundPositionEdgeItem{Edge: props.BackgroundPositionEdgeRight, Unit: props.UnitPx(20)},
		): "top 10px right 20px",
		props.BackgroundPositionTop():         "top",
		props.BackgroundPositionBottom():      "bottom",
		props.BackgroundPositionLeft():        "left",
		props.BackgroundPositionRight():       "right",
		props.BackgroundPositionCenter():      "center",
		props.BackgroundPositionTopLeft():     "top left",
		props.BackgroundPositionTopRight():    "top right",
		props.BackgroundPositionBottomLeft():  "bottom left",
		props.BackgroundPositionBottomRight(): "bottom right",
		props.BackgroundPositions(
			props.BackgroundPositionTop(),
			props.BackgroundPositionLeft(),
		): "top,left",
		props.BackgroundPosition("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BackgroundPosition: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{background-position:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BackgroundRepeat(t *testing.T) {
	testCases := map[props.BackgroundRepeat]string{
		props.BackgroundRepeatRepeat:   "repeat",
		props.BackgroundRepeatRepeatX:  "repeat-x",
		props.BackgroundRepeatRepeatY:  "repeat-y",
		props.BackgroundRepeatNoRepeat: "no-repeat",
		props.BackgroundRepeatSpace:    "space",
		props.BackgroundRepeatRound:    "round",
		props.BackgroundRepeats(
			props.BackgroundRepeatRepeat,
			props.BackgroundRepeatNoRepeat,
		): "repeat,no-repeat",
		props.BackgroundRepeat("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BackgroundRepeat: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{background-repeat:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BackgroundSize(t *testing.T) {
	testCases := map[props.BackgroundSize]string{
		props.BackgroundSizeWidth(props.UnitPx(10)): "10px",
		props.BackgroundSizeDimension(
			props.UnitPx(10),
			props.UnitPx(20),
		): "10px 20px",
		props.BackgroundSizes(
			props.BackgroundSizeWidth(props.UnitPx(10)),
			props.BackgroundSizeDimension(props.UnitPx(10), props.UnitPx(20)),
		): "10px,10px 20px",
		props.BackgroundSizeCover():     "cover",
		props.BackgroundSizeContain():   "contain",
		props.BackgroundSize("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BackgroundSize: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{background-size:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderColor(t *testing.T) {
	testCases := map[props.Color]string{
		props.ColorRGBA(0, 0, 0, 255):       "rgba(0,0,0,1.00)",
		props.ColorRGBA(255, 255, 255, 230): "rgba(255,255,255,0.90)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderColor: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-color:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_Border(t *testing.T) {
	testCases := map[props.Border]string{
		props.Border{
			props.UnitPx(10),
			props.BorderStyleSolid,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		props.Border{
			props.UnitPx(10),
			props.BorderStyleDouble,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Border: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderBottom(t *testing.T) {
	testCases := map[props.Border]string{
		props.Border{
			props.UnitPx(10),
			props.BorderStyleSolid,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		props.Border{
			props.UnitPx(10),
			props.BorderStyleDouble,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderBottom: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-bottom:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderLeft(t *testing.T) {
	testCases := map[props.Border]string{
		props.Border{
			props.UnitPx(10),
			props.BorderStyleSolid,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		props.Border{
			props.UnitPx(10),
			props.BorderStyleDouble,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderLeft: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-left:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderRadius: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-radius:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderRight(t *testing.T) {
	testCases := map[props.Border]string{
		props.Border{
			props.UnitPx(10),
			props.BorderStyleSolid,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		props.Border{
			props.UnitPx(5),
			props.BorderStyleDouble,
			props.ColorRGBA(0, 0, 0, 255),
		}: "5px double rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderRight: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-right:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderStyle(t *testing.T) {
	testCases := map[props.BorderStyle]string{
		props.BorderStyleNone:        "none",
		props.BorderStyleHidden:      "hidden",
		props.BorderStyleDotted:      "dotted",
		props.BorderStyleDashed:      "dashed",
		props.BorderStyleSolid:       "solid",
		props.BorderStyleDouble:      "double",
		props.BorderStyle("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderStyle: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-style:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderTop(t *testing.T) {
	testCases := map[props.Border]string{
		props.Border{
			props.UnitPx(10),
			props.BorderStyleSolid,
			props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		props.Border{
			props.UnitPx(5),
			props.BorderStyleDouble,
			props.ColorRGBA(0, 0, 0, 255),
		}: "5px double rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderTop: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-top:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_BorderWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{BorderWidth: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{border-width:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_CaptionSide(t *testing.T) {
	testCases := map[props.CaptionSide]string{
		props.CaptionSideTop:         "top",
		props.CaptionSideBottom:      "bottom",
		props.CaptionSide("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{CaptionSide: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			css := fmt.Sprintf(".test{caption-side:%s;}", expected)
			if buf.String() != css {
				t.Errorf("expected %q, got %q", css, buf.String())
			}
		})
	}
}

func TestStyle_Color(t *testing.T) {
	testCases := map[props.Color]string{
		props.ColorRGBA(0, 0, 0, 255):                  "rgba(0,0,0,1.00)",
		props.ColorRGBA(255, 255, 255, 230):            "rgba(255,255,255,0.90)",
		props.ColorRGBA(255, 255, 255, 255).Alpha(230): "rgba(255,255,255,0.90)",
		props.ColorCurrentColor():                      "currentColor",
		props.ColorInherit():                           "inherit",
		props.ColorTransparent():                       "transparent",
		props.Color{Keyword: "#efefef"}:                "#efefef",
		props.Color{Keyword: "red"}:                    "red",
		props.Color{RGBA: color.RGBA{0, 0, 0, 255}}:    "rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Color: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if buf.String() != fmt.Sprintf(".test{color:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Display(t *testing.T) {
	testCases := map[props.Display]string{
		props.DisplayBlock:            "block",
		props.DisplayInline:           "inline",
		props.DisplayInlineBlock:      "inline-block",
		props.DisplayInlineFlex:       "inline-flex",
		props.DisplayInlineGrid:       "inline-grid",
		props.DisplayFlex:             "flex",
		props.DisplayGrid:             "grid",
		props.DisplayNone:             "none",
		props.DisplayTable:            "table",
		props.DisplayInlineTable:      "inline-table",
		props.DisplayTableCaption:     "table-caption",
		props.DisplayTableCell:        "table-cell",
		props.DisplayTableColumn:      "table-column",
		props.DisplayTableColumnGroup: "table-column-group",
		props.DisplayTableFooterGroup: "table-footer-group",
		props.DisplayTableHeaderGroup: "table-header-group",
		props.DisplayTableRowGroup:    "table-row-group",
		props.DisplayTableRow:         "table-row",
		props.DisplayFlowRoot:         "flow-root",
		props.DisplayContents:         "contents",
		props.DisplayListItem:         "list-item",
		props.Display("inherit"):      "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Display: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if buf.String() != fmt.Sprintf(".test{display:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FlexBasis(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FlexBasis: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{flex-basis:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FlexDirection(t *testing.T) {
	testCases := map[props.FlexDirection]string{
		props.FlexDirectionRow:           "row",
		props.FlexDirectionRowReverse:    "row-reverse",
		props.FlexDirectionColumn:        "column",
		props.FlexDirectionColumnReverse: "column-reverse",
		props.FlexDirection("initial"):   "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FlexDirection: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{flex-direction:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FlexGrow(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1):         "1",
		props.UnitRaw(0):         "0",
		props.UnitRaw("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FlexGrow: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{flex-grow:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FlexShrink(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1):         "1",
		props.UnitRaw(0):         "0",
		props.UnitRaw("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FlexShrink: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{flex-shrink:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FlexWrap(t *testing.T) {
	testCases := map[props.FlexWrap]string{
		props.FlexWrapNoWrap:      "nowrap",
		props.FlexWrapWrap:        "wrap",
		props.FlexWrapWrapReverse: "wrap-reverse",
		props.FlexWrap("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FlexWrap: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{flex-wrap:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Float(t *testing.T) {
	testCases := map[props.Float]string{
		props.FloatLeft:        "left",
		props.FloatRight:       "right",
		props.FloatNone:        "none",
		props.Float("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Float: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{float:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FontSize(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(20): "20px",
		props.UnitRem(2): "2.000rem",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FontSize: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{font-size:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FontStyle(t *testing.T) {
	testCases := map[props.FontStyle]string{
		props.FontStyleNormal:      "normal",
		props.FontStyleItalic:      "italic",
		props.FontStyle("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FontStyle: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{font-style:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_FontWeight(t *testing.T) {
	testCases := map[props.FontWeight]string{
		props.FontWeightThin:        "100",
		props.FontWeightExtraLight:  "200",
		props.FontWeightLight:       "300",
		props.FontWeightNormal:      "400",
		props.FontWeightMedium:      "500",
		props.FontWeightSemiBold:    "600",
		props.FontWeightBold:        "700",
		props.FontWeightExtraBold:   "800",
		props.FontWeightBlack:       "900",
		props.FontWeight("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{FontWeight: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{font-weight:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Height(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(20):      "20px",
		props.UnitRem(2):      "2.000rem",
		props.UnitPercent(50): "50.00%",
		props.UnitAuto():      "auto",
		props.UnitRaw(0):      "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Height: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{height:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_JustifyContent(t *testing.T) {
	testCases := map[props.JustifyContent]string{
		props.JustifyContentNormal:       "normal",
		props.JustifyContentFlexStart:    "flex-start",
		props.JustifyContentFlexEnd:      "flex-end",
		props.JustifyContentCenter:       "center",
		props.JustifyContentSpaceBetween: "space-between",
		props.JustifyContentSpaceAround:  "space-around",
		props.JustifyContentSpaceEvenly:  "space-evenly",
		props.JustifyContentStretch:      "stretch",
		props.JustifyContent("initial"):  "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{JustifyContent: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{justify-content:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_JustifyItems(t *testing.T) {
	testCases := map[props.JustifyItems]string{
		props.JustifyItemsStretch:     "stretch",
		props.JustifyItemsCenter:      "center",
		props.JustifyItemsStart:       "start",
		props.JustifyItemsEnd:         "end",
		props.JustifyItems("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{JustifyItems: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{justify-items:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_JustifySelf(t *testing.T) {
	testCases := map[props.JustifySelf]string{
		props.JustifySelfStretch:     "stretch",
		props.JustifySelfCenter:      "center",
		props.JustifySelfStart:       "start",
		props.JustifySelfEnd:         "end",
		props.JustifySelf("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{JustifySelf: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{justify-self:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_LineHeight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1.5): "1.5",
		props.UnitPx(20):   "20px",
		props.UnitEm(3):    "3.000em",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{LineHeight: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{line-height:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Margin(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Margin: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{margin:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MarginBottom(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MarginBottom: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{margin-bottom:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MarginLeft(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MarginLeft: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{margin-left:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MarginRight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MarginRight: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{margin-right:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MarginTop(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MarginTop: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{margin-top:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MaxWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MaxWidth: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{max-width:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_MinWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{MinWidth: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{min-width:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Overflow(t *testing.T) {
	testCases := map[props.Overflow]string{
		props.OverflowVisible:     "visible",
		props.OverflowHidden:      "hidden",
		props.OverflowScroll:      "scroll",
		props.OverflowAuto:        "auto",
		props.OverflowClip:        "clip",
		props.Overflow("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Overflow: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{overflow:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_OverflowX(t *testing.T) {
	testCases := map[props.Overflow]string{
		props.OverflowVisible:     "visible",
		props.OverflowHidden:      "hidden",
		props.OverflowScroll:      "scroll",
		props.OverflowAuto:        "auto",
		props.OverflowClip:        "clip",
		props.Overflow("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{OverflowX: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{overflow-x:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_OverflowY(t *testing.T) {
	testCases := map[props.Overflow]string{
		props.OverflowVisible:     "visible",
		props.OverflowHidden:      "hidden",
		props.OverflowScroll:      "scroll",
		props.OverflowAuto:        "auto",
		props.OverflowClip:        "clip",
		props.Overflow("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{OverflowY: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{overflow-y:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Padding(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Padding: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{padding:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_PaddingBottom(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{PaddingBottom: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{padding-bottom:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_PaddingLeft(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{PaddingLeft: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{padding-left:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_PaddingRight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{PaddingRight: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{padding-right:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_PaddingTop(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{PaddingTop: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{padding-top:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Position(t *testing.T) {
	testCases := map[props.Position]string{
		props.PositionStatic:      "static",
		props.PositionRelative:    "relative",
		props.PositionAbsolute:    "absolute",
		props.PositionFixed:       "fixed",
		props.PositionSticky:      "sticky",
		props.Position("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Position: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{position:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_PrintColorAdjust(t *testing.T) {
	testCases := map[props.PrintColorAdjust]string{
		props.PrintColorAdjustEconomy:  "economy",
		props.PrintColorAdjustExact:    "exact",
		props.PrintColorAdjust("auto"): "auto",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{PrintColorAdjust: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if buf.String() != fmt.Sprintf(".test{print-color-adjust:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_TextAlign(t *testing.T) {
	testCases := map[props.TextAlign]string{
		props.TextAlignLeft:        "left",
		props.TextAlignRight:       "right",
		props.TextAlignCenter:      "center",
		props.TextAlignJustify:     "justify",
		props.TextAlignStart:       "start",
		props.TextAlignEnd:         "end",
		props.TextAlign("initial"): "initial",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{TextAlign: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if buf.String() != fmt.Sprintf(".test{text-align:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}

func TestStyle_Width(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(20):      "20px",
		props.UnitRem(2):      "2.000rem",
		props.UnitPercent(50): "50.00%",
		props.UnitAuto():      "auto",
		props.UnitRaw(0):      "0",
	}

	for prop, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			st := &Style{Selector: ".test", Props: Props{Width: prop}}
			var buf bytes.Buffer
			err := st.CSS(&buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if buf.String() != fmt.Sprintf(".test{width:%s;}", expected) {
				t.Errorf("expected %q, got %q", expected, buf.String())
			}
		})
	}
}
