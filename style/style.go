package style

import (
	"fmt"
	"io"
	"reflect"

	"github.com/AccentDesign/gostyle/props/align"
	"github.com/AccentDesign/gostyle/props/appearance"
	"github.com/AccentDesign/gostyle/props/background"
	"github.com/AccentDesign/gostyle/props/border"
	"github.com/AccentDesign/gostyle/props/colors"
	"github.com/AccentDesign/gostyle/props/display"
	"github.com/AccentDesign/gostyle/props/flex"
	"github.com/AccentDesign/gostyle/props/font"
	"github.com/AccentDesign/gostyle/props/justify"
	"github.com/AccentDesign/gostyle/props/margin"
	"github.com/AccentDesign/gostyle/props/print"
	"github.com/AccentDesign/gostyle/props/table"
	"github.com/AccentDesign/gostyle/props/text"
	"github.com/AccentDesign/gostyle/props/unit"
	"github.com/AccentDesign/gostyle/props/width"
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
		BorderRadius       unit.Unit             `css:"border-radius"`
		Border             border.Border         `css:"border"`
		BorderTop          border.Border         `css:"border-top"`
		BorderRight        border.Border         `css:"border-right"`
		BorderBottom       border.Border         `css:"border-bottom"`
		BorderLeft         border.Border         `css:"border-left"`
		BorderColor        colors.Color          `css:"border-color"`
		BorderStyle        border.Style          `css:"border-style"`
		BorderWidth        unit.Unit             `css:"border-width"`
		CaptionSide        table.CaptionSide     `css:"caption-side"`
		Color              colors.Color          `css:"color"`
		Display            display.Display       `css:"display"`
		FlexDirection      flex.Direction        `css:"flex-direction"`
		FontSize           unit.Unit             `css:"font-size"`
		FontWeight         font.Weight           `css:"font-weight"`
		Height             unit.Unit             `css:"height"`
		JustifyContent     justify.Content       `css:"justify-content"`
		LineHeight         unit.Unit             `css:"line-height"`
		Margin             margin.Margin         `css:"margin"`
		MarginTop          margin.Margin         `css:"margin-top"`
		MarginRight        margin.Margin         `css:"margin-right"`
		MarginBottom       margin.Margin         `css:"margin-bottom"`
		MarginLeft         margin.Margin         `css:"margin-left"`
		MaxWidth           unit.Unit             `css:"max-width"`
		MinWidth           unit.Unit             `css:"min-width"`
		Padding            unit.Unit             `css:"padding"`
		PaddingTop         unit.Unit             `css:"padding-top"`
		PaddingRight       unit.Unit             `css:"padding-right"`
		PaddingBottom      unit.Unit             `css:"padding-bottom"`
		PaddingLeft        unit.Unit             `css:"padding-left"`
		PrintColorAdjust   print.ColorAdjust     `css:"print-color-adjust"`
		TextAlign          text.Align            `css:"text-align"`
		Width              width.Width           `css:"width"`
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
