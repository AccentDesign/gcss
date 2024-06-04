package gcss

import (
	"bytes"
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"image/color"
	"testing"
)

func runTest(t *testing.T, st *Style, expected string) {
	var buf bytes.Buffer
	err := st.CSS(&buf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestStyle_Empty(t *testing.T) {
	st := &Style{Selector: ".test", Props: Props{}}
	css := ".test{}"
	runTest(t, st, css)
}

func TestStyle_MultipleProps(t *testing.T) {
	st := &Style{Selector: ".test", Props: Props{
		BackgroundColor: props.ColorRGBA(0, 0, 0, 255),
		Height:          props.UnitPx(100),
		Width:           props.UnitPx(100),
	}}
	css := ".test{background-color:rgba(0,0,0,1.00);height:100px;width:100px;}"
	runTest(t, st, css)
}

func TestStyle_CustomProps(t *testing.T) {
	st := &Style{
		Selector: ".test",
		CustomProps: []CustomProp{
			{Attr: "--color", Value: "red"},
			{Attr: "background-color", Value: "var(--color)"},
		},
	}
	css := ".test{--color:red;background-color:var(--color);}"
	runTest(t, st, css)
}

func TestStyle_AlignContent(t *testing.T) {
	testCases := map[props.AlignContent]string{
		props.AlignContentNormal:        "normal",
		props.AlignContentStart:         "start",
		props.AlignContentCenter:        "center",
		props.AlignContentEnd:           "end",
		props.AlignContentFlexStart:     "flex-start",
		props.AlignContentFlexEnd:       "flex-end",
		props.AlignContentBaseline:      "baseline",
		props.AlignContentFirstBaseline: "first baseline",
		props.AlignContentLastBaseline:  "last baseline",
		props.AlignContentSpaceBetween:  "space-between",
		props.AlignContentSpaceAround:   "space-around",
		props.AlignContentSpaceEvenly:   "space-evenly",
		props.AlignContentStretch:       "stretch",
		props.AlignContent("inherit"):   "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{AlignContent: prop}}
		css := fmt.Sprintf(".test{align-content:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_AlignItems(t *testing.T) {
	testCases := map[props.AlignItems]string{
		props.AlignItemsNormal:      "normal",
		props.AlignItemsStart:       "start",
		props.AlignItemsCenter:      "center",
		props.AlignItemsEnd:         "end",
		props.AlignItemsFlexStart:   "flex-start",
		props.AlignItemsFlexEnd:     "flex-end",
		props.AlignItemsSelfStart:   "self-start",
		props.AlignItemsSelfEnd:     "self-end",
		props.AlignItemsBaseline:    "baseline",
		props.AlignItemsStretch:     "stretch",
		props.AlignItems("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{AlignItems: prop}}
		css := fmt.Sprintf(".test{align-items:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_AlignSelf(t *testing.T) {
	testCases := map[props.AlignSelf]string{
		props.AlignSelfAuto:        "auto",
		props.AlignSelfNormal:      "normal",
		props.AlignSelfStretch:     "stretch",
		props.AlignSelfCenter:      "center",
		props.AlignSelfStart:       "start",
		props.AlignSelfEnd:         "end",
		props.AlignSelfFlexStart:   "flex-start",
		props.AlignSelfFlexEnd:     "flex-end",
		props.AlignSelfSelfStart:   "self-start",
		props.AlignSelfSelfEnd:     "self-end",
		props.AlignSelfBaseline:    "baseline",
		props.AlignSelf("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{AlignSelf: prop}}
		css := fmt.Sprintf(".test{align-self:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Appearance: prop}}
		css := fmt.Sprintf(".test{appearance:%s;}", expected)
		runTest(t, st, css)
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
		{Keyword: "#efefef"}:                           "#efefef",
		{Keyword: "red"}:                               "red",
		{Color: color.Gray16{0}}:                       "rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BackgroundColor: prop}}
		css := fmt.Sprintf(".test{background-color:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BackgroundImage(t *testing.T) {
	testCases := map[props.BackgroundImage]string{
		props.BackgroundImageURL("image.jpg"):              `url("image.jpg")`,
		props.BackgroundImageLinearGradient("red", "blue"): "linear-gradient(red,blue)",
		props.BackgroundImage("inherit"):                   "inherit",
		props.BackgroundImages(
			props.BackgroundImageLinearGradient("red", "blue"),
			props.BackgroundImageURL("image.jpg"),
		): `linear-gradient(red,blue),url("image.jpg")`,
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BackgroundImage: prop}}
		css := fmt.Sprintf(".test{background-image:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BackgroundPosition(t *testing.T) {
	testCases := map[props.BackgroundPosition]string{
		props.BackgroundPositionXY(props.UnitPx(10), props.UnitPx(20)): "10px 20px",
		props.BackgroundPositionEdges(
			props.BackgroundPositionEdge{Position: props.BackgroundPositionTop, Unit: props.UnitPx(10)},
			props.BackgroundPositionEdge{Position: props.BackgroundPositionRight, Unit: props.UnitPx(20)},
		): "top 10px right 20px",
		props.BackgroundPositionTop:         "top",
		props.BackgroundPositionBottom:      "bottom",
		props.BackgroundPositionLeft:        "left",
		props.BackgroundPositionRight:       "right",
		props.BackgroundPositionCenter:      "center",
		props.BackgroundPositionTopLeft:     "top left",
		props.BackgroundPositionTopRight:    "top right",
		props.BackgroundPositionBottomLeft:  "bottom left",
		props.BackgroundPositionBottomRight: "bottom right",
		props.BackgroundPositions(
			props.BackgroundPositionTop,
			props.BackgroundPositionLeft,
		): "top,left",
		props.BackgroundPosition("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BackgroundPosition: prop}}
		css := fmt.Sprintf(".test{background-position:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{BackgroundRepeat: prop}}
		css := fmt.Sprintf(".test{background-repeat:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{BackgroundSize: prop}}
		css := fmt.Sprintf(".test{background-size:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Border(t *testing.T) {
	testCases := map[props.Border]string{
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleSolid,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleDouble,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
		{
			Style: props.BorderStyleNone,
		}: "none",
		{
			Style: props.BorderStyle("initial"),
		}: "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Border: prop}}
		css := fmt.Sprintf(".test{border:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderBottom(t *testing.T) {
	testCases := map[props.Border]string{
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleSolid,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleDouble,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
		{
			Style: props.BorderStyleNone,
		}: "none",
		{
			Style: props.BorderStyle("initial"),
		}: "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderBottom: prop}}
		css := fmt.Sprintf(".test{border-bottom:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderBottomLeftRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):         "10px",
		props.UnitPercent(50):    "50.00%",
		props.UnitRem(10):        "10.000rem",
		props.UnitRaw("20% 10%"): "20% 10%",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderBottomLeftRadius: prop}}
		css := fmt.Sprintf(".test{border-bottom-left-radius:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderBottomRightRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):         "10px",
		props.UnitPercent(50):    "50.00%",
		props.UnitRem(10):        "10.000rem",
		props.UnitRaw("20% 10%"): "20% 10%",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderBottomRightRadius: prop}}
		css := fmt.Sprintf(".test{border-bottom-right-radius:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderCollapse(t *testing.T) {
	testCases := map[props.BorderCollapse]string{
		props.BorderCollapseSeparate:    "separate",
		props.BorderCollapseCollapse:    "collapse",
		props.BorderCollapse("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderCollapse: prop}}
		css := fmt.Sprintf(".test{border-collapse:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderColor(t *testing.T) {
	testCases := map[props.Color]string{
		props.ColorRGBA(0, 0, 0, 255):       "rgba(0,0,0,1.00)",
		props.ColorRGBA(255, 255, 255, 230): "rgba(255,255,255,0.90)",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderColor: prop}}
		css := fmt.Sprintf(".test{border-color:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderLeft(t *testing.T) {
	testCases := map[props.Border]string{
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleSolid,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleDouble,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px double rgba(0,0,0,1.00)",
		{
			Style: props.BorderStyleNone,
		}: "none",
		{
			Style: props.BorderStyle("initial"),
		}: "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderLeft: prop}}
		css := fmt.Sprintf(".test{border-left:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderRadius: prop}}
		css := fmt.Sprintf(".test{border-radius:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderRight(t *testing.T) {
	testCases := map[props.Border]string{
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleSolid,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		{
			Width: props.UnitPx(5),
			Style: props.BorderStyleDouble,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "5px double rgba(0,0,0,1.00)",
		{
			Style: props.BorderStyleNone,
		}: "none",
		{
			Style: props.BorderStyle("initial"),
		}: "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderRight: prop}}
		css := fmt.Sprintf(".test{border-right:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{BorderStyle: prop}}
		css := fmt.Sprintf(".test{border-style:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderTop(t *testing.T) {
	testCases := map[props.Border]string{
		{
			Width: props.UnitPx(10),
			Style: props.BorderStyleSolid,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "10px solid rgba(0,0,0,1.00)",
		{
			Width: props.UnitPx(5),
			Style: props.BorderStyleDouble,
			Color: props.ColorRGBA(0, 0, 0, 255),
		}: "5px double rgba(0,0,0,1.00)",
		{
			Style: props.BorderStyleNone,
		}: "none",
		{
			Style: props.BorderStyle("initial"),
		}: "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderTop: prop}}
		css := fmt.Sprintf(".test{border-top:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderTopLeftRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):         "10px",
		props.UnitPercent(50):    "50.00%",
		props.UnitRem(10):        "10.000rem",
		props.UnitRaw("20% 10%"): "20% 10%",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderTopLeftRadius: prop}}
		css := fmt.Sprintf(".test{border-top-left-radius:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderTopRightRadius(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):         "10px",
		props.UnitPercent(50):    "50.00%",
		props.UnitRem(10):        "10.000rem",
		props.UnitRaw("20% 10%"): "20% 10%",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderTopRightRadius: prop}}
		css := fmt.Sprintf(".test{border-top-right-radius:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BorderWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BorderWidth: prop}}
		css := fmt.Sprintf(".test{border-width:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Bottom(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Bottom: prop}}
		css := fmt.Sprintf(".test{bottom:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_BoxSizing(t *testing.T) {
	testCases := map[props.BoxSizing]string{
		props.BoxSizingBorderBox:   "border-box",
		props.BoxSizingContentBox:  "content-box",
		props.BoxSizing("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{BoxSizing: prop}}
		css := fmt.Sprintf(".test{box-sizing:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_CaptionSide(t *testing.T) {
	testCases := map[props.CaptionSide]string{
		props.CaptionSideTop:         "top",
		props.CaptionSideBottom:      "bottom",
		props.CaptionSide("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{CaptionSide: prop}}
		css := fmt.Sprintf(".test{caption-side:%s;}", expected)
		runTest(t, st, css)
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
		{Keyword: "#efefef"}:                           "#efefef",
		{Keyword: "red"}:                               "red",
		{Color: color.Gray16{0}}:                       "rgba(0,0,0,1.00)",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Color: prop}}
		css := fmt.Sprintf(".test{color:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_ColumnGap(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{ColumnGap: prop}}
		css := fmt.Sprintf(".test{column-gap:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Cursor(t *testing.T) {
	testCases := map[props.Cursor]string{
		props.CursorAuto:        "auto",
		props.CursorDefault:     "default",
		props.CursorNone:        "none",
		props.CursorHelp:        "help",
		props.CursorPointer:     "pointer",
		props.CursorWait:        "wait",
		props.CursorText:        "text",
		props.CursorMove:        "move",
		props.CursorNotAllowed:  "not-allowed",
		props.Cursor("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Cursor: prop}}
		css := fmt.Sprintf(".test{cursor:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Display: prop}}
		css := fmt.Sprintf(".test{display:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_FlexBasis(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitAuto(): "auto",
		props.UnitRaw(0): "0",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FlexBasis: prop}}
		css := fmt.Sprintf(".test{flex-basis:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{FlexDirection: prop}}
		css := fmt.Sprintf(".test{flex-direction:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_FlexGrow(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1):         "1",
		props.UnitRaw(0):         "0",
		props.UnitRaw("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FlexGrow: prop}}
		css := fmt.Sprintf(".test{flex-grow:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_FlexShrink(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1):         "1",
		props.UnitRaw(0):         "0",
		props.UnitRaw("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FlexShrink: prop}}
		css := fmt.Sprintf(".test{flex-shrink:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{FlexWrap: prop}}
		css := fmt.Sprintf(".test{flex-wrap:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Float: prop}}
		css := fmt.Sprintf(".test{float:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_FontFamily(t *testing.T) {
	testCases := map[props.FontFamily]string{
		props.FontFamilySans:  "ui-sans-serif, system-ui, sans-serif, \"Apple Color Emoji\", \"Segoe UI Emoji\", \"Segoe UI Symbol\", \"Noto Color Emoji\"",
		props.FontFamilySerif: "ui-serif, Georgia, Cambria, \"Times New Roman\", Times, serif",
		props.FontFamilyMono:  "ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace",
		props.FontFamily("\"Gill Sans Extrabold\", sans-serif"): "\"Gill Sans Extrabold\", sans-serif",
		props.FontFamily("sans-serif"):                          "sans-serif",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FontFamily: prop}}
		css := fmt.Sprintf(".test{font-family:%s;}", expected)
		runTest(t, st, css)
	}

}
func TestStyle_FontSize(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(20): "20px",
		props.UnitRem(2): "2.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FontSize: prop}}
		css := fmt.Sprintf(".test{font-size:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_FontStyle(t *testing.T) {
	testCases := map[props.FontStyle]string{
		props.FontStyleNormal:      "normal",
		props.FontStyleItalic:      "italic",
		props.FontStyle("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{FontStyle: prop}}
		css := fmt.Sprintf(".test{font-style:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{FontWeight: prop}}
		css := fmt.Sprintf(".test{font-weight:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Gap(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Gap: prop}}
		css := fmt.Sprintf(".test{gap:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Height: prop}}
		css := fmt.Sprintf(".test{height:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_JustifyContent(t *testing.T) {
	testCases := map[props.JustifyContent]string{
		props.JustifyContentNormal:       "normal",
		props.JustifyContentCenter:       "center",
		props.JustifyContentStart:        "start",
		props.JustifyContentEnd:          "end",
		props.JustifyContentFlexStart:    "flex-start",
		props.JustifyContentFlexEnd:      "flex-end",
		props.JustifyContentLeft:         "left",
		props.JustifyContentRight:        "right",
		props.JustifyContentSpaceBetween: "space-between",
		props.JustifyContentSpaceAround:  "space-around",
		props.JustifyContentSpaceEvenly:  "space-evenly",
		props.JustifyContentStretch:      "stretch",
		props.JustifyContent("initial"):  "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{JustifyContent: prop}}
		css := fmt.Sprintf(".test{justify-content:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_JustifyItems(t *testing.T) {
	testCases := map[props.JustifyItems]string{
		props.JustifyItemsNormal:      "normal",
		props.JustifyItemsStretch:     "stretch",
		props.JustifyItemsCenter:      "center",
		props.JustifyItemsStart:       "start",
		props.JustifyItemsEnd:         "end",
		props.JustifyItemsFlexStart:   "flex-start",
		props.JustifyItemsFlexEnd:     "flex-end",
		props.JustifyItemsSelfStart:   "self-start",
		props.JustifyItemsSelfEnd:     "self-end",
		props.JustifyItemsLeft:        "left",
		props.JustifyItemsRight:       "right",
		props.JustifyItemsBaseline:    "baseline",
		props.JustifyItems("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{JustifyItems: prop}}
		css := fmt.Sprintf(".test{justify-items:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_JustifySelf(t *testing.T) {
	testCases := map[props.JustifySelf]string{
		props.JustifySelfAuto:        "auto",
		props.JustifySelfNormal:      "normal",
		props.JustifySelfStretch:     "stretch",
		props.JustifySelfCenter:      "center",
		props.JustifySelfStart:       "start",
		props.JustifySelfEnd:         "end",
		props.JustifySelfFlexStart:   "flex-start",
		props.JustifySelfFlexEnd:     "flex-end",
		props.JustifySelfSelfStart:   "self-start",
		props.JustifySelfSelfEnd:     "self-end",
		props.JustifySelfLeft:        "left",
		props.JustifySelfRight:       "right",
		props.JustifySelfBaseline:    "baseline",
		props.JustifySelf("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{JustifySelf: prop}}
		css := fmt.Sprintf(".test{justify-self:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Left(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Left: prop}}
		css := fmt.Sprintf(".test{left:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_ListStylePosition(t *testing.T) {
	testCases := map[props.ListStylePosition]string{
		props.ListStylePositionInside:  "inside",
		props.ListStylePositionOutside: "outside",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{ListStylePosition: prop}}
		css := fmt.Sprintf(".test{list-style-position:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_ListStyleType(t *testing.T) {
	testCases := map[props.ListStyleType]string{
		props.ListStyleTypeNone:    "none",
		props.ListStyleTypeDisc:    "disc",
		props.ListStyleTypeDecimal: "decimal",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{ListStyleType: prop}}
		css := fmt.Sprintf(".test{list-style-type:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_LineHeight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1.5): "1.5",
		props.UnitPx(20):   "20px",
		props.UnitEm(3):    "3.000em",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{LineHeight: prop}}
		css := fmt.Sprintf(".test{line-height:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Margin(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Margin: prop}}
		css := fmt.Sprintf(".test{margin:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MarginBottom(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MarginBottom: prop}}
		css := fmt.Sprintf(".test{margin-bottom:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MarginLeft(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MarginLeft: prop}}
		css := fmt.Sprintf(".test{margin-left:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MarginRight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MarginRight: prop}}
		css := fmt.Sprintf(".test{margin-right:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MarginTop(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10): "10px",
		props.UnitRem(2): "2.000rem",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MarginTop: prop}}
		css := fmt.Sprintf(".test{margin-top:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MaxHeight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):  "10px",
		props.UnitRem(2):  "2.000rem",
		props.UnitVh(100): "100vh",
		props.UnitAuto():  "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MaxHeight: prop}}
		css := fmt.Sprintf(".test{max-height:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MaxWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):  "10px",
		props.UnitRem(2):  "2.000rem",
		props.UnitVw(100): "100vw",
		props.UnitAuto():  "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MaxWidth: prop}}
		css := fmt.Sprintf(".test{max-width:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MinHeight(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):  "10px",
		props.UnitRem(2):  "2.000rem",
		props.UnitVh(100): "100vh",
		props.UnitAuto():  "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MinHeight: prop}}
		css := fmt.Sprintf(".test{min-height:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_MinWidth(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):  "10px",
		props.UnitRem(2):  "2.000rem",
		props.UnitVw(100): "100vw",
		props.UnitAuto():  "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{MinWidth: prop}}
		css := fmt.Sprintf(".test{min-width:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Opacity(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(0):         "0",
		props.UnitRaw(0.33):      "0.33",
		props.UnitPercent(90):    "90.00%",
		props.UnitRaw("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Opacity: prop}}
		css := fmt.Sprintf(".test{opacity:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Overflow: prop}}
		css := fmt.Sprintf(".test{overflow:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{OverflowX: prop}}
		css := fmt.Sprintf(".test{overflow-x:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{OverflowY: prop}}
		css := fmt.Sprintf(".test{overflow-y:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Padding: prop}}
		css := fmt.Sprintf(".test{padding:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{PaddingBottom: prop}}
		css := fmt.Sprintf(".test{padding-bottom:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{PaddingLeft: prop}}
		css := fmt.Sprintf(".test{padding-left:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{PaddingRight: prop}}
		css := fmt.Sprintf(".test{padding-right:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{PaddingTop: prop}}
		css := fmt.Sprintf(".test{padding-top:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Position: prop}}
		css := fmt.Sprintf(".test{position:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_PrintColorAdjust(t *testing.T) {
	testCases := map[props.PrintColorAdjust]string{
		props.PrintColorAdjustEconomy:  "economy",
		props.PrintColorAdjustExact:    "exact",
		props.PrintColorAdjust("auto"): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{PrintColorAdjust: prop}}
		css := fmt.Sprintf(".test{print-color-adjust:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Right(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Right: prop}}
		css := fmt.Sprintf(".test{right:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_RowGap(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{RowGap: prop}}
		css := fmt.Sprintf(".test{row-gap:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{TextAlign: prop}}
		css := fmt.Sprintf(".test{text-align:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextDecorationColor(t *testing.T) {
	testCases := map[props.Color]string{
		props.ColorRGBA(0, 0, 0, 255):                  "rgba(0,0,0,1.00)",
		props.ColorRGBA(255, 255, 255, 230):            "rgba(255,255,255,0.90)",
		props.ColorRGBA(255, 255, 255, 255).Alpha(230): "rgba(255,255,255,0.90)",
		props.ColorCurrentColor():                      "currentColor",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextDecorationColor: prop}}
		css := fmt.Sprintf(".test{text-decoration-color:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextDecorationLine(t *testing.T) {
	testCases := map[props.TextDecorationLine]string{
		props.TextDecorationLineNone:        "none",
		props.TextDecorationLineUnderline:   "underline",
		props.TextDecorationLineOverline:    "overline",
		props.TextDecorationLineLineThrough: "line-through",
		props.TextDecorationLine("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextDecorationLine: prop}}
		css := fmt.Sprintf(".test{text-decoration-line:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextDecorationStyle(t *testing.T) {
	testCases := map[props.TextDecorationStyle]string{
		props.TextDecorationStyleSolid:       "solid",
		props.TextDecorationStyleDouble:      "double",
		props.TextDecorationStyleDotted:      "dotted",
		props.TextDecorationStyleDashed:      "dashed",
		props.TextDecorationStyleWavy:        "wavy",
		props.TextDecorationStyle("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextDecorationStyle: prop}}
		css := fmt.Sprintf(".test{text-decoration-style:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextDecorationThickness(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(1):            "1px",
		props.UnitAuto():           "auto",
		props.UnitRaw("from-font"): "from-font",
		props.UnitInherit():        "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextDecorationThickness: prop}}
		css := fmt.Sprintf(".test{text-decoration-thickness:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextIndent(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(0):      "0px",
		props.UnitRem(0.875): "0.875rem",
		props.UnitInherit():  "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextIndent: prop}}
		css := fmt.Sprintf(".test{text-indent:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextUnderlineOffset(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(1):  "1px",
		props.UnitAuto(): "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextUnderlineOffset: prop}}
		css := fmt.Sprintf(".test{text-underline-offset:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextOverflow(t *testing.T) {
	testCases := map[props.TextOverflow]string{
		props.TextOverflowClip:        "clip",
		props.TextOverflowEllipsis:    "ellipsis",
		props.TextOverflow("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextOverflow: prop}}
		css := fmt.Sprintf(".test{text-overflow:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextTransform(t *testing.T) {
	testCases := map[props.TextTransform]string{
		props.TextTransformNone:        "none",
		props.TextTransformCapitalize:  "capitalize",
		props.TextTransformUppercase:   "uppercase",
		props.TextTransformLowercase:   "lowercase",
		props.TextTransform("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextTransform: prop}}
		css := fmt.Sprintf(".test{text-transform:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_TextWrap(t *testing.T) {
	testCases := map[props.TextWrap]string{
		props.TextWrapWrap:        "wrap",
		props.TextWrapNoWrap:      "nowrap",
		props.TextWrapBalance:     "balance",
		props.TextWrap("inherit"): "inherit",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{TextWrap: prop}}
		css := fmt.Sprintf(".test{text-wrap:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Top(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitPx(10):      "10px",
		props.UnitPercent(50): "50.00%",
		props.UnitRem(10):     "10.000rem",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Top: prop}}
		css := fmt.Sprintf(".test{top:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_VerticalAlign(t *testing.T) {
	testCases := map[props.VerticalAlign]string{
		props.VerticalAlignBaseline:    "baseline",
		props.VerticalAlignSub:         "sub",
		props.VerticalAlignSuper:       "super",
		props.VerticalAlignTextTop:     "text-top",
		props.VerticalAlignTextBottom:  "text-bottom",
		props.VerticalAlignMiddle:      "middle",
		props.VerticalAlignTop:         "top",
		props.VerticalAlignBottom:      "bottom",
		props.VerticalAlign("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{VerticalAlign: prop}}
		css := fmt.Sprintf(".test{vertical-align:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_WhiteSpace(t *testing.T) {
	testCases := map[props.WhiteSpace]string{
		props.WhiteSpaceNormal:      "normal",
		props.WhiteSpaceNowrap:      "nowrap",
		props.WhiteSpacePre:         "pre",
		props.WhiteSpacePreLine:     "pre-line",
		props.WhiteSpacePreWrap:     "pre-wrap",
		props.WhiteSpace("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{WhiteSpace: prop}}
		css := fmt.Sprintf(".test{white-space:%s;}", expected)
		runTest(t, st, css)
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
		st := &Style{Selector: ".test", Props: Props{Width: prop}}
		css := fmt.Sprintf(".test{width:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_Visibility(t *testing.T) {
	testCases := map[props.Visibility]string{
		props.VisibilityVisible:     "visible",
		props.VisibilityHidden:      "hidden",
		props.VisibilityCollapse:    "collapse",
		props.Visibility("initial"): "initial",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{Visibility: prop}}
		css := fmt.Sprintf(".test{visibility:%s;}", expected)
		runTest(t, st, css)
	}
}

func TestStyle_ZIndex(t *testing.T) {
	testCases := map[props.Unit]string{
		props.UnitRaw(1):  "1",
		props.UnitRaw(10): "10",
		props.UnitAuto():  "auto",
	}

	for prop, expected := range testCases {
		st := &Style{Selector: ".test", Props: Props{ZIndex: prop}}
		css := fmt.Sprintf(".test{z-index:%s;}", expected)
		runTest(t, st, css)
	}
}
