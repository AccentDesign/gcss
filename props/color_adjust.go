package props

type ColorAdjust string

const (
	ColorAdjustEconomy     ColorAdjust = "economy"
	ColorAdjustExact       ColorAdjust = "exact"
	ColorAdjustInherit     ColorAdjust = "inherit"
	ColorAdjustInitial     ColorAdjust = "initial"
	ColorAdjustRevert      ColorAdjust = "revert"
	ColorAdjustRevertLayer ColorAdjust = "revert-layer"
	ColorAdjustUnset       ColorAdjust = "unset"
)

func (c ColorAdjust) String() string {
	return string(c)
}
