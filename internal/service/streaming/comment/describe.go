package comment

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *DummyCommentService) Describe(comment_id uint64) (*streaming.Comment, error) {
	id, ok := c.findElementIDByCommentID(comment_id)
	if !ok {
		return nil, fmt.Errorf("id does not exist")
	}

	comment := c.comments[id]

	return &comment, nil
}
