package props

type Display string

const (
	DisplayBlock            Display = "block"
	DisplayInlineBlock      Display = "inline-block"
	DisplayInline           Display = "inline"
	DisplayFlex             Display = "flex"
	DisplayInlineFlex       Display = "inline-flex"
	DisplayTable            Display = "table"
	DisplayInlineTable      Display = "inline-table"
	DisplayTableCaption     Display = "table-caption"
	DisplayTableCell        Display = "table-cell"
	DisplayTableColumn      Display = "table-column"
	DisplayTableColumnGroup Display = "table-column-group"
	DisplayTableFooterGroup Display = "table-footer-group"
	DisplayTableHeaderGroup Display = "table-header-group"
	DisplayTableRowGroup    Display = "table-row-group"
	DisplayTableRow         Display = "table-row"
	DisplayFlowRoot         Display = "flow-root"
	DisplayGrid             Display = "grid"
	DisplayInlineGrid       Display = "inline-grid"
	DisplayContents         Display = "contents"
	DisplayListItem         Display = "list-item"
	DisplayHidden           Display = "none"
)

func (d Display) String() string {
	return string(d)
}
