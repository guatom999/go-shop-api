package exception

import "fmt"

type (
	PlayerCreating struct {
		PlayerID string
	}
)

func (c *PlayerCreating) Error() string {
	return fmt.Sprintf("creating playerID %s failed", c.PlayerID)
}
