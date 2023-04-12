package usecase

import "jargonjester/domain"

type chatUsecase struct {
	conversationRepository domain.ConversationRepository
	openaiRepository       domain.OpenaiRepository
}

func NewChatUsercase(conversationRepository domain.ConversationRepository, openaiRepository domain.OpenaiRepository) domain.ChatUsecase {
	return &chatUsecase{
		conversationRepository: conversationRepository,
		openaiRepository:       openaiRepository,
	}
}
