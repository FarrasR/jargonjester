package usecase

import (
	"fmt"
	"jargonjester/entity"
	"time"

	"github.com/pkoukk/tiktoken-go"
)

func (u *chatUsecase) PrivateChat(channelID string, username string, message string) (string, error) {
	//for now lets use the same key for all
	err := u.configRepository.IsLimited("chat")

	if err != nil {
		return "", err
	}

	previousConversations, err := u.conversationRepository.GetMessagesInAChannel(channelID)
	if err != nil {
		return "", err
	}

	var messages []entity.Message

	userPrompt := fmt.Sprintf(promptPrivateChatContent, username)

	promptMessage := entity.Message{
		Role:    promptRole,
		Content: userPrompt,
	}

	tokenCount, err := tiktoken.EncodingForModel(gptModel)
	if err != nil {
		return "", err
	}

	promptcontenttokencount := tokenCount.Encode(promptGroupChatContent, nil, nil)
	promptroletokencount := tokenCount.Encode(promptRole, nil, nil)
	promptToken := 4 + len(promptroletokencount) + len(promptcontenttokencount)
	totalToken := 3 + promptToken

	userMessage := entity.Message{
		Role:    userRole,
		Content: message,
	}

	totalToken = totalToken + len(tokenCount.Encode(userMessage.Role, nil, nil)) + len(tokenCount.Encode(userMessage.Content, nil, nil)) + 4

	for i := len(previousConversations) - 1; i >= 0; i-- {
		role := userRole

		if previousConversations[i].Username == botNameString {
			role = assistantRole
		}

		totalToken = totalToken + len(tokenCount.Encode(previousConversations[i].Content, nil, nil)) + len(tokenCount.Encode(role, nil, nil)) + 4

		if totalToken >= maximumToken {
			break
		}

		messages = append(messages, entity.Message{
			Role:    role,
			Content: previousConversations[i].Content,
		})
	}

	messages = append(messages, promptMessage)

	//swap the slices
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	messages = append(messages, userMessage)

	userConversation := entity.Conversation{
		ChannelID:   channelID,
		Username:    username,
		Content:     message,
		MessageTime: time.Now(),
	}

	response, err := u.openaiRepository.CompleteChat(gptModel, messages)
	if err != nil {
		return "", err
	}

	responseConversation := entity.Conversation{
		ChannelID:   channelID,
		Username:    botNameString,
		Content:     response.Content,
		MessageTime: time.Now(),
	}

	if err := u.conversationRepository.InsertMessage(userConversation); err != nil {
		return "", err
	}

	if err := u.conversationRepository.InsertMessage(responseConversation); err != nil {
		return "", err
	}

	return response.Content, nil
}
