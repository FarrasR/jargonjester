package domain

import "github.com/bwmarrin/discordgo"

type DiscordHandler interface {
	Register(session *discordgo.Session)
}
