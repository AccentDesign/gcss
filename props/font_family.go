package props

type FontFamily string

const (
	FontFamilySans  FontFamily = "ui-sans-serif, system-ui, sans-serif, \"Apple Color Emoji\", \"Segoe UI Emoji\", \"Segoe UI Symbol\", \"Noto Color Emoji\""
	FontFamilySerif FontFamily = "ui-serif, Georgia, Cambria, \"Times New Roman\", Times, serif"
	FontFamilyMono  FontFamily = "ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace"
)

func (f FontFamily) String() string {
	return string(f)
}
