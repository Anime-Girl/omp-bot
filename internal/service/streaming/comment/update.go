package comment

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Update(comment_id uint64, comment streaming.Comment) error {
	comment.Text = strings.Trim(comment.Text, " ")

	if comment.Text == "" {
		return fmt.Errorf("comment's text is empty")
	}

	id, ok := c.findElementIDByCommentID(comment_id)
	if !ok {
		return fmt.Errorf("id does not exist")
	}

	c.comments[id] = comment

	return nil
}
