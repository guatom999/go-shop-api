package exception

type (
	CoinNotEnough struct {
	}
)

func (e *CoinNotEnough) Error() string {
	return "player coin is not enough"
}
