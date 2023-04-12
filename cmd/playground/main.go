package main

import (
	"fmt"
	chatUsecase "jargonjester/chat/usecase"
	conversationRepository "jargonjester/conversation/repository"
	"jargonjester/database"
	openaiRepository "jargonjester/openai/repository"
	"jargonjester/utils"
	"os"
)

func main() {
	utils.LoadEnvironment()

	db := database.InitDB()

	conversationRepository := conversationRepository.NewConversationRepository(db)
	openaiRepository := openaiRepository.NewOpenaiRepository(os.Getenv("OPENAI_HOST"), os.Getenv("OPENAI_KEY"))

	chatUsecase := chatUsecase.NewChatUsercase(conversationRepository, openaiRepository)

	response, err := chatUsecase.ChannelChat("123456", "Farras", "What is my name?")

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
