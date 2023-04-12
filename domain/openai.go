package domain

import (
	"jargonjester/entity"
)

type OpenaiRepository interface {
	CompleteChat(model string, messages []entity.Message) (entity.Message, error)
}
