package gcss

import (
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"io"
	"reflect"
)

type (
	// Props represents the standard CSS properties that can be applied to a CSS rule.
	Props struct {
		AlignItems              props.AlignItems         `css:"align-items"`
		Appearance              props.Appearance         `css:"appearance"`
		BackgroundColor         props.Color              `css:"background-color"`
		BackgroundImage         props.BackgroundImage    `css:"background-image"`
		BackgroundPosition      props.BackgroundPosition `css:"background-position"`
		BackgroundRepeat        props.BackgroundRepeat   `css:"background-repeat"`
		BackgroundSize          props.BackgroundSize     `css:"background-size"`
		BorderColor             props.Color              `css:"border-color"`
		Border                  props.Border             `css:"border"`
		BorderBottom            props.Border             `css:"border-bottom"`
		BorderBottomLeftRadius  props.Unit               `css:"border-bottom-left-radius"`
		BorderBottomRightRadius props.Unit               `css:"border-bottom-right-radius"`
		BorderLeft              props.Border             `css:"border-left"`
		BorderRadius            props.Unit               `css:"border-radius"`
		BorderRight             props.Border             `css:"border-right"`
		BorderStyle             props.BorderStyle        `css:"border-style"`
		BorderTop               props.Border             `css:"border-top"`
		BorderTopLeftRadius     props.Unit               `css:"border-top-left-radius"`
		BorderTopRightRadius    props.Unit               `css:"border-top-right-radius"`
		BorderWidth             props.Unit               `css:"border-width"`
		Bottom                  props.Unit               `css:"bottom"`
		BoxSizing               props.BoxSizing          `css:"box-sizing"`
		CaptionSide             props.CaptionSide        `css:"caption-side"`
		Color                   props.Color              `css:"color"`
		ColumnGap               props.Unit               `css:"column-gap"`
		Cursor                  props.Cursor             `css:"cursor"`
		Display                 props.Display            `css:"display"`
		FlexBasis               props.Unit               `css:"flex-basis"`
		FlexDirection           props.FlexDirection      `css:"flex-direction"`
		FlexGrow                props.Unit               `css:"flex-grow"`
		FlexShrink              props.Unit               `css:"flex-shrink"`
		FlexWrap                props.FlexWrap           `css:"flex-wrap"`
		Float                   props.Float              `css:"float"`
		FontSize                props.Unit               `css:"font-size"`
		FontStyle               props.FontStyle          `css:"font-style"`
		FontWeight              props.FontWeight         `css:"font-weight"`
		Gap                     props.Unit               `css:"gap"`
		Height                  props.Unit               `css:"height"`
		JustifyContent          props.JustifyContent     `css:"justify-content"`
		JustifyItems            props.JustifyItems       `css:"justify-items"`
		JustifySelf             props.JustifySelf        `css:"justify-self"`
		Left                    props.Unit               `css:"left"`
		LineHeight              props.Unit               `css:"line-height"`
		ListStylePosition       props.ListStylePosition  `css:"list-style-position"`
		ListStyleType           props.ListStyleType      `css:"list-style-type"`
		Margin                  props.Unit               `css:"margin"`
		MarginBottom            props.Unit               `css:"margin-bottom"`
		MarginLeft              props.Unit               `css:"margin-left"`
		MarginRight             props.Unit               `css:"margin-right"`
		MarginTop               props.Unit               `css:"margin-top"`
		MaxWidth                props.Unit               `css:"max-width"`
		MinWidth                props.Unit               `css:"min-width"`
		Overflow                props.Overflow           `css:"overflow"`
		OverflowX               props.Overflow           `css:"overflow-x"`
		OverflowY               props.Overflow           `css:"overflow-y"`
		Padding                 props.Unit               `css:"padding"`
		PaddingBottom           props.Unit               `css:"padding-bottom"`
		PaddingLeft             props.Unit               `css:"padding-left"`
		PaddingRight            props.Unit               `css:"padding-right"`
		PaddingTop              props.Unit               `css:"padding-top"`
		Position                props.Position           `css:"position"`
		PrintColorAdjust        props.PrintColorAdjust   `css:"print-color-adjust"`
		Right                   props.Unit               `css:"right"`
		RowGap                  props.Unit               `css:"row-gap"`
		TextAlign               props.TextAlign          `css:"text-align"`
		TextOverflow            props.TextOverflow       `css:"text-overflow"`
		Top                     props.Unit               `css:"top"`
		VerticalAlign           props.VerticalAlign      `css:"vertical-align"`
		WhiteSpace              props.WhiteSpace         `css:"white-space"`
		Width                   props.Unit               `css:"width"`
	}
	// CustomProp represents an additional CSS property that is not covered by the Props struct.
	CustomProp struct {
		Attr, Value string
	}
	// Style represents a CSS style rule.
	Style struct {
		// Selector is the CSS selector to which the properties will be applied.
		// It can be any valid CSS selector like class, id, element type, etc.
		Selector string

		// Props contains the standard CSS properties that will be applied to the selector.
		// These properties are represented by the Props struct and are strongly typed.
		Props Props

		// CustomProps contains any additional CSS properties that are not covered by the Props struct.
		// These properties are directly added to the CSS rule as is.
		CustomProps []CustomProp
	}
)

// CSS writes the CSS representation of the style to the writer.
func (s *Style) CSS(w io.Writer) error {
	propsValue := reflect.ValueOf(s.Props)
	propsType := reflect.TypeOf(s.Props)

	if _, err := fmt.Fprintf(w, "%s{", s.Selector); err != nil {
		return err
	}

	// Iterate over the fields of the Props struct and write the CSS properties to the writer.
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

	// Write the custom properties to the writer.
	for _, prop := range s.CustomProps {
		if _, err := fmt.Fprintf(w, "%s:%s;", prop.Attr, prop.Value); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}

	return nil
}
