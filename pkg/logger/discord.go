package logger

import "github.com/bwmarrin/discordgo"

type DiscordErrorLog struct {
	Level     string `json:"level"`
	Caller    string `json:"caller"`
	Message   string `json:"msg"`
	Timestamp string `json:"timestamp"`
}

func WebhookSend(webhookID, webhookToken, text string) {
	dc, _ := discordgo.New("Bot")
	defer dc.Close()

	hookMsg := &discordgo.WebhookParams{
		Content: text,
	}

	_, _ = dc.WebhookExecute(webhookID, webhookToken, false, hookMsg)
}
