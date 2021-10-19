package comment

import "fmt"

func (c *DummyCommentService) Remove(comment_id uint64) (bool, error) {
	id, ok := c.findElementIDByCommentID(comment_id)
	if !ok {
		return false, fmt.Errorf("id does not exist")
	}

	if c.comments[id].ID == c.comments[len(c.comments)-1].ID {
		c.comments = c.comments[:id]
		return true, nil
	}

	c.comments = append(c.comments[:id], c.comments[id+1:]...)

	return true, nil
}
