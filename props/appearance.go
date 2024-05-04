package props

type Appearance string

const (
	AppearanceNone           Appearance = "none"
	AppearanceAuto           Appearance = "auto"
	AppearanceMenuListButton Appearance = "menulist-button"
	AppearanceTextField      Appearance = "textfield"
	AppearanceInherit        Appearance = "inherit"
	AppearanceInitial        Appearance = "initial"
	AppearanceRevert         Appearance = "revert"
	AppearanceRevertLater    Appearance = "revert-later"
	AppearanceUnset          Appearance = "unset"
	AppearanceButton         Appearance = "button"
	AppearanceCheckbox       Appearance = "checkbox"
)

func (a Appearance) String() string {
	return string(a)
}
