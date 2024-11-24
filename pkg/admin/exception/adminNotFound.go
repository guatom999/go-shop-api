package exception

import "fmt"

type (
	AdminNotFound struct {
		AdminID string
	}
)

func (e *AdminNotFound) Error() string {
	return fmt.Sprintf("AdminID %s not found ", e.AdminID)
}
