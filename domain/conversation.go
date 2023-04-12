package domain

import "jargonjester/entity"

type ConversationRepository interface {
	GetLast10MessageInAChannel(string) ([]entity.Conversation, error)
	InsertMessage(entity.Conversation) error
}
