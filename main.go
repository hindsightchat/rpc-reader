package main

import (
	"fmt"

	richpresence "github.com/rmcord/discord-rpc-reader/src"
)

func main() {
	reader := richpresence.New()

	reader.OnPresence(func(update richpresence.PresenceUpdate) {
		fmt.Printf("App: %s\n", update.ClientID)
		if update.Activity != nil {
			fmt.Printf("Playing: %s - %s\n", update.Activity.Details, update.Activity.State)
		}
	})

	reader.Start()

	// keep the program running
	select {}
	
}
