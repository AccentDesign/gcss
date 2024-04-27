package display

type Display string

const (
	Block            Display = "block"
	InlineBlock      Display = "inline-block"
	Inline           Display = "inline"
	Flex             Display = "flex"
	InlineFlex       Display = "inline-flex"
	Table            Display = "table"
	InlineTable      Display = "inline-table"
	TableCaption     Display = "table-caption"
	TableCell        Display = "table-cell"
	TableColumn      Display = "table-column"
	TableColumnGroup Display = "table-column-group"
	TableFooterGroup Display = "table-footer-group"
	TableHeaderGroup Display = "table-header-group"
	TableRowGroup    Display = "table-row-group"
	TableRow         Display = "table-row"
	FlowRoot         Display = "flow-root"
	Grid             Display = "grid"
	InlineGrid       Display = "inline-grid"
	Contents         Display = "contents"
	ListItem         Display = "list-item"
	None             Display = "none"
)

func (d Display) String() string {
	return string(d)
}
