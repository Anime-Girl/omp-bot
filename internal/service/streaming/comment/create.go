package comment

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Create(comment streaming.Comment) (uint64, error) {
	comment.Text = strings.Trim(comment.Text, " ")

	if comment.Text == "" {
		return 0, fmt.Errorf("comment's text is empty")
	}

	var newID uint64 = 1

	if len(c.comments) != 0 {
		newID = c.comments[len(c.comments)-1].ID + 1
	}

	c.comments = append(c.comments, streaming.Comment{
		ID:   newID,
		Text: comment.Text,
	})

	return newID, nil
}
