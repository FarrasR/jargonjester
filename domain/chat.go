package domain

type ChatUsecase interface {
	ChannelChat(channelID string, username string, message string) (string, error)
}
