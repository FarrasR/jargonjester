package usecase

import (
	"fmt"
	"jargonjester/entity"
	"strings"
	"time"

	"github.com/pkoukk/tiktoken-go"
)

func (u *chatUsecase) GroupChat(channelID string, username string, message string) (string, error) {

	previousConversations, err := u.conversationRepository.GetMessagesInAChannel(channelID)
	if err != nil {
		return "", err
	}

	var messages []entity.Message

	promptMessage := entity.Message{
		Role:    promptRole,
		Content: promptGroupChatContent,
	}

	tokenCount, err := tiktoken.EncodingForModel(gptModel)
	if err != nil {
		panic(err)
	}

	//counting token
	promptcontenttokencount := tokenCount.Encode(promptGroupChatContent, nil, nil)
	promptroletokencount := tokenCount.Encode(promptRole, nil, nil)
	promptToken := 4 + len(promptroletokencount) + len(promptcontenttokencount)
	totalToken := 3 + promptToken

	userMessage := entity.Message{
		Role: userRole,
		Content: fmt.Sprintf("%s: %s",
			username,
			message,
		),
	}

	totalToken = totalToken + len(tokenCount.Encode(userMessage.Role, nil, nil)) + len(tokenCount.Encode(userMessage.Content, nil, nil)) + 4

	//iterate through previous messages to count tokens
	for i := 0; i < len(previousConversations); i++ {
		role := userRole
		content := fmt.Sprintf("%s: %s",
			previousConversations[i].Username,
			previousConversations[i].Content,
		)

		if previousConversations[i].Username == botNameString {
			role = assistantRole
		}

		totalToken = totalToken + len(tokenCount.Encode(content, nil, nil)) + len(tokenCount.Encode(role, nil, nil)) + 4

		// will break if more than allowed tokens
		if totalToken >= maximumToken {
			break
		}

		messages = append(messages, entity.Message{
			Role:    role,
			Content: content,
		})
	}

	//append prompt message
	messages = append(messages, promptMessage)

	//swap the slices
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	userConversation := entity.Conversation{
		ChannelID:   channelID,
		Username:    username,
		Content:     message,
		MessageTime: time.Now(),
	}

	//append user message
	messages = append(messages, userMessage)

	//get response
	response, err := u.openaiRepository.CompleteChat(gptModel, messages)

	if err != nil {
		return "", err
	}
	//append message and payload

	conversation := strings.SplitN(response.Content, ": ", 2)

	if len(conversation) != 2 {
		fmt.Println("length of response is not two, debug this")
		fmt.Println(response.Content)
		return response.Content, nil
	}

	responseConversation := entity.Conversation{
		ChannelID:   channelID,
		Username:    botNameString,
		Content:     conversation[1],
		MessageTime: time.Now(),
	}

	u.conversationRepository.InsertMessage(userConversation)
	u.conversationRepository.InsertMessage(responseConversation)

	return conversation[1], nil
}
