package exception

type (
	NoPerMission struct {
	}
)

func (e *NoPerMission) Error() string {
	return "no permission"
}
