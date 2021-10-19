package comment

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *StreamingCommentCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	texts := strings.Split(args, `\\\`)
	var createdIDs []uint64

	for i := range texts {
		comment := streaming.Comment{
			Text: texts[i],
		}

		id, err := c.commentService.Create(comment)
		if err != nil {
			log.Printf("fail to create comment: %v", err)
			continue
		}

		createdIDs = append(createdIDs, id)
	}

	response := fmt.Sprintf("Comment with %v id's created", createdIDs)

	if len(createdIDs) == 0 {
		response = "Comment should not be empty"
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		response,
	)

	c.bot.Send(msg)
}
