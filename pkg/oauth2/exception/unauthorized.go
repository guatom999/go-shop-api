package exception

type (
	UnAuthorized struct {
	}
)

func (e *UnAuthorized) Error() string {
	return "status unAuthorized"
}
