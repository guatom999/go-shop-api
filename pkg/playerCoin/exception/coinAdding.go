package exception

type (
	CoinAdding struct {
	}
)

func (e *CoinAdding) Error() string {
	return "coin adding failed"
}
