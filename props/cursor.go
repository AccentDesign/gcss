package props

type Cursor string

const (
	CursorAuto       Cursor = "auto"
	CursorDefault    Cursor = "default"
	CursorPointer    Cursor = "pointer"
	CursorWait       Cursor = "wait"
	CursorText       Cursor = "text"
	CursorMove       Cursor = "move"
	CursorHelp       Cursor = "help"
	CursorNotAllowed Cursor = "not-allowed"
	CursorNone       Cursor = "none"
)

func (c Cursor) String() string {
	return string(c)
}
