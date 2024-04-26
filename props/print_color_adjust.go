package props

type PrintColorAdjust string

const (
	PrintColorAdjustEconomy     PrintColorAdjust = "economy"
	PrintColorAdjustExact       PrintColorAdjust = "exact"
	PrintColorAdjustInherit     PrintColorAdjust = "inherit"
	PrintColorAdjustInitial     PrintColorAdjust = "initial"
	PrintColorAdjustRevert      PrintColorAdjust = "revert"
	PrintColorAdjustRevertLayer PrintColorAdjust = "revert-layer"
	PrintColorAdjustUnset       PrintColorAdjust = "unset"
)

func (p PrintColorAdjust) String() string {
	return string(p)
}
