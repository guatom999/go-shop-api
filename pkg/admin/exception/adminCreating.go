package exception

import "fmt"

type (
	AdminCreating struct {
		AdminID string
	}
)

func (c *AdminCreating) Error() string {
	return fmt.Sprintf("creating AdminID %s failed", c.AdminID)
}
