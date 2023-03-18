package listeners

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/bot"
	"context"
	"fmt"
	"log"
)

func ConsoleListener(ctx context.Context, server *conf.Server, console chan string, commandSent chan bool) {
	var (
		consoleOutput []string
	)
	for {
		select {
		case output := <-console:
			consoleOutput = append(consoleOutput, output)
			if len(consoleOutput) > server.Console.MessageLines {
				sendConsoleOutput(server, consoleOutput)
			}
		case <-commandSent:
			sendConsoleOutput(server, consoleOutput)
		case <-ctx.Done():
			return
		}
	}
}

func sendConsoleOutput(server *conf.Server, consoleOutput []string) {
	var message string
	for _, line := range consoleOutput {
		message += fmt.Sprintf("\n%v", line)
	}
	for _, channelID := range server.Console.ChannelID {
		_, err := bot.Session.ChannelMessageSend(channelID, fmt.Sprintf("```%v```", message))
		if err != nil {
			log.Printf("Failed to send console to discord: %v (ChannelID: %v)", err, channelID)
		}
	}
}
