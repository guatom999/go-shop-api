package exception

type (
	PlayerCoinShowing struct {
	}
)

func (e *PlayerCoinShowing) Error() string {
	return "showing coin failed"
}
