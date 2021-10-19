package comment

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type CommentService interface {
	Describe(comment_id uint64) (*streaming.Comment, error)
	List(cursor uint64, limit uint64) ([]streaming.Comment, error)
	Create(streaming.Comment) (uint64, error)
	Update(comment_id uint64, comment streaming.Comment) error
	Remove(comment_id uint64) (bool, error)

	CommentsCount() int
}

type DummyCommentService struct {
	comments []streaming.Comment
}

func NewDummyCommentService() *DummyCommentService {
	return &DummyCommentService{}
}

func (c *DummyCommentService) CommentsCount() int {
	return len(c.comments)
}

func (c *DummyCommentService) checkIndexOutOfRange(id int) bool {
	return id < 0 || id >= len(c.comments)
}

func (c *DummyCommentService) findElementIDByCommentID(target uint64) (int, bool) {
	l := 0
	r := len(c.comments) - 1

	for l <= r {
		m := (l + r) / 2
		if c.comments[m].ID == target {
			return m, true
		}

		if c.comments[m].ID < target {
			l = m + 1
			continue
		}

		r = m - 1
	}

	return 0, false
}
