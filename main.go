package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	godotenv.Load()
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".

		Token:  os.Getenv("TELEGRAM_TOKEN_BOT"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	fmt.Println(os.Getenv("TELEGRAM_TOKEN_BOT"))
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Chat, "Hello World")
	})

	// Private to bot
	b.Handle("/hello", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		b.Send(m.Chat, "Hello!")
	})

	b.Start()
}
