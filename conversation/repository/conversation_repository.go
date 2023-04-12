package repository

import (
	"jargonjester/domain"
	"jargonjester/entity"

	"gorm.io/gorm"
)

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) domain.ConversationRepository {
	return &conversationRepository{
		db: db,
	}
}

func (r *conversationRepository) GetLast10MessageInAChannel(channelID string) ([]entity.Conversation, error) {
	var conversations []entity.Conversation

	if result := r.db.Limit(10).Where(&entity.Conversation{ChannelID: channelID}).Order("message_time desc").Find(&conversations); result.Error != nil {
		return []entity.Conversation{}, result.Error
	}

	return conversations, nil
}

func (r *conversationRepository) InsertMessage(conversation entity.Conversation) error {
	if result := r.db.Create(&conversation); result.Error != nil {
		return result.Error
	}

	return nil
}
