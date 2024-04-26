package appearance

type Appearance string

const (
	None        Appearance = "none"
	Auto        Appearance = "auto"
	Default     Appearance = "menulist-button"
	TextField   Appearance = "textfield"
	Inherit     Appearance = "inherit"
	Initial     Appearance = "initial"
	Revert      Appearance = "revert"
	RevertLater Appearance = "revert-later"
	Unset       Appearance = "unset"
	Button      Appearance = "button"
	Checkbox    Appearance = "checkbox"
)

func (a Appearance) String() string {
	return string(a)
}
