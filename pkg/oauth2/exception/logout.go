package exceoption

type (
	Logout struct {
	}
)

func (e *Logout) Error() string {
	return "logout failed"
}
