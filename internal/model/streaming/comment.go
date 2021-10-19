package streaming

import "fmt"

type Comment struct {
	ID   uint64
	Text string
}

func (c *Comment) String() string {
	return fmt.Sprintf(`ID: %d, Text: %s`, c.ID, c.Text)
}
