package usecase

import "jargonjester/domain"

type permissionUsecase struct {
	permissionRepository domain.PermissionRepository
}

func NewPermissionUsecase(
	conversationRepository domain.ConversationRepository,
	openaiRepository domain.OpenaiRepository,
	configRepository domain.ConfigRepository) domain.ChatUsecase {
	return &chatUsecase{
		conversationRepository: conversationRepository,
		openaiRepository:       openaiRepository,
		configRepository:       configRepository,
	}
}
