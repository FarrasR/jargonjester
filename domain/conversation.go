package domain

import "jargonjester/entity"

type ConversationRepository interface {
	GetMessagesInAChannel(string) ([]entity.Conversation, error)
	InsertMessage(entity.Conversation) error
}
