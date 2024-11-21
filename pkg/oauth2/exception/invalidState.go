package exceoption

type (
	InvalidState struct {
	}
)

func (e *InvalidState) Error() string {
	return "invalid state"
}
