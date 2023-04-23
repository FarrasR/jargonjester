package discord

import (
	"fmt"
	"jargonjester/domain"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type ChatHandler struct {
	ChatUsecase domain.ChatUsecase
}

func (h *ChatHandler) Register(session *discordgo.Session) {
	session.AddHandler(h.handleGroupChat)
	session.AddHandler(h.handlePrivateChat)
}

func NewChatHandler(chatUsecase domain.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		ChatUsecase: chatUsecase,
	}
}

func (h *ChatHandler) handleGroupChat(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	userChannel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		fmt.Println("Error detecting channel please debug")
		return
	}

	messageContent := strings.Split(m.Message.Content, " ")

	if m.ChannelID == userChannel.ID {
		return
	}

	if messageContent[0] != "!chat" {
		return
	}

	response, err := h.ChatUsecase.GroupChat(m.ChannelID, m.Author.Username, strings.Join(messageContent[1:], " "))

	if err != nil {
		s.ChannelMessageSend(
			m.ChannelID,
			err.Error(),
		)
	}

	s.ChannelMessageSend(m.ChannelID, response)

}

func (h *ChatHandler) handlePrivateChat(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	userChannel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		fmt.Println("Error detecting channel please debug")
		return
	}

	//only works in DM
	if m.ChannelID != userChannel.ID {
		return
	}

	response, err := h.ChatUsecase.PrivateChat(m.ChannelID, m.Author.Username, m.Message.Content)

	if err != nil {
		s.ChannelMessageSend(
			m.ChannelID,
			err.Error(),
		)
	}

	s.ChannelMessageSend(m.ChannelID, response)
}
