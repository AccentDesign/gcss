package style

import (
	"fmt"
	"github.com/AccentDesign/gcss/props/overflow"
	"io"
	"reflect"

	"github.com/AccentDesign/gcss/props/align"
	"github.com/AccentDesign/gcss/props/appearance"
	"github.com/AccentDesign/gcss/props/background"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/display"
	"github.com/AccentDesign/gcss/props/flex"
	"github.com/AccentDesign/gcss/props/font"
	"github.com/AccentDesign/gcss/props/justify"
	"github.com/AccentDesign/gcss/props/print"
	"github.com/AccentDesign/gcss/props/table"
	"github.com/AccentDesign/gcss/props/text"
	"github.com/AccentDesign/gcss/props/unit"
)

type (
	Props struct {
		AlignItems         align.Items           `css:"align-items"`
		Appearance         appearance.Appearance `css:"appearance"`
		BackgroundColor    colors.Color          `css:"background-color"`
		BackgroundImage    background.Image      `css:"background-image"`
		BackgroundPosition background.Position   `css:"background-position"`
		BackgroundRepeat   background.Repeat     `css:"background-repeat"`
		BackgroundSize     background.Size       `css:"background-size"`
		BorderColor        colors.Color          `css:"border-color"`
		Border             border.Border         `css:"border"`
		BorderBottom       border.Border         `css:"border-bottom"`
		BorderLeft         border.Border         `css:"border-left"`
		BorderRadius       unit.Unit             `css:"border-radius"`
		BorderRight        border.Border         `css:"border-right"`
		BorderStyle        border.Style          `css:"border-style"`
		BorderTop          border.Border         `css:"border-top"`
		BorderWidth        unit.Unit             `css:"border-width"`
		CaptionSide        table.CaptionSide     `css:"caption-side"`
		Color              colors.Color          `css:"color"`
		Display            display.Display       `css:"display"`
		FlexDirection      flex.Direction        `css:"flex-direction"`
		FlexGrow           unit.Unit             `css:"flex-grow"`
		FlexShrink         unit.Unit             `css:"flex-shrink"`
		FlexWrap           flex.Wrap             `css:"flex-wrap"`
		FontSize           unit.Unit             `css:"font-size"`
		FontStyle          font.Style            `css:"font-style"`
		FontWeight         font.Weight           `css:"font-weight"`
		Height             unit.Unit             `css:"height"`
		JustifyContent     justify.Content       `css:"justify-content"`
		justifyItems       justify.Items         `css:"justify-items"`
		justifySelf        justify.Self          `css:"justify-self"`
		LineHeight         unit.Unit             `css:"line-height"`
		Margin             unit.Unit             `css:"margin"`
		MarginBottom       unit.Unit             `css:"margin-bottom"`
		MarginLeft         unit.Unit             `css:"margin-left"`
		MarginRight        unit.Unit             `css:"margin-right"`
		MarginTop          unit.Unit             `css:"margin-top"`
		MaxWidth           unit.Unit             `css:"max-width"`
		MinWidth           unit.Unit             `css:"min-width"`
		Overflow           overflow.Overflow     `css:"overflow"`
		OverflowX          overflow.Overflow     `css:"overflow-x"`
		OverflowY          overflow.Overflow     `css:"overflow-y"`
		Padding            unit.Unit             `css:"padding"`
		PaddingBottom      unit.Unit             `css:"padding-bottom"`
		PaddingLeft        unit.Unit             `css:"padding-left"`
		PaddingRight       unit.Unit             `css:"padding-right"`
		PaddingTop         unit.Unit             `css:"padding-top"`
		PrintColorAdjust   print.ColorAdjust     `css:"print-color-adjust"`
		TextAlign          text.Align            `css:"text-align"`
		Width              unit.Unit             `css:"width"`
	}
	Style struct {
		Selector string
		Props    Props
	}
)

// CSS returns the CSS string representation of the style.
func (s *Style) CSS(w io.Writer) error {
	propsValue := reflect.ValueOf(s.Props)
	propsType := reflect.TypeOf(s.Props)

	if _, err := fmt.Fprintf(w, "%s{", s.Selector); err != nil {
		return err
	}

	for i := 0; i < propsValue.NumField(); i++ {
		fieldValue := propsValue.Field(i)
		fieldType := propsType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		fieldName := fieldType.Tag.Get("css")

		if v, ok := fieldValue.Interface().(fmt.Stringer); ok {
			if _, err := fmt.Fprintf(w, "%s:%s;", fieldName, v.String()); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}

	return nil
}
