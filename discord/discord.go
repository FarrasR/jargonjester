package discord

import (
	"jargonjester/domain"

	"github.com/bwmarrin/discordgo"
)

func BuildHandlers(session *discordgo.Session, handlers ...domain.DiscordHandler) {
	for _, handler := range handlers {
		handler.Register(session)
	}
}
