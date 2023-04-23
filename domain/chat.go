package domain

type ChatUsecase interface {
	GroupChat(channelID string, username string, message string) (string, error)
	PrivateChat(channelID string, username string, message string) (string, error)
}
