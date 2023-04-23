package main

import (
	"fmt"
	chatDiscordHandler "jargonjester/chat/delivery/discord"
	chatUsecase "jargonjester/chat/usecase"
	conversationRepository "jargonjester/conversation/repository"
	"jargonjester/database"
	"jargonjester/discord"
	openaiRepository "jargonjester/openai/repository"
	"jargonjester/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	utils.LoadEnvironment()

	db := database.InitDB()

	conversationRepository := conversationRepository.NewConversationRepository(db)
	openaiRepository := openaiRepository.NewOpenaiRepository(os.Getenv("OPENAI_HOST"), os.Getenv("OPENAI_KEY"))

	chatUsecase := chatUsecase.NewChatUsercase(conversationRepository, openaiRepository)

	chatHandler := chatDiscordHandler.NewChatHandler(chatUsecase)

	discordSession, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.BuildHandlers(discordSession, chatHandler)

	discordSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsDirectMessages | discordgo.IntentsGuildMessages)

	err = discordSession.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discordSession.Close()
}
