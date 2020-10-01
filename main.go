package main

import (
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

		Token:       os.Getenv("TELEGRAM_TOKEN_BOT"),
		Synchronous: true,
		Poller:      &tb.LongPoller{Timeout: 10 * time.Second},
	})

	var (
		// Universal markup builders.
		selector = &tb.ReplyMarkup{}
		// Inline buttons.
		btnMoon = selector.Data("Moon 🌚", "moon")
		btnSun  = selector.Data("Sun 🌞", "sun")
	)

	selector.Inline(
		selector.Row(btnMoon, btnSun),
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(&btnMoon, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Message.Chat, "Imam says 'Bacot'!")
	})

	b.Handle(&btnSun, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Message.Chat, "Imam says 'BGST'!")
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(
			m.Chat,
			"Word, you choose",
			selector,
		)
	})

	b.Start()
}
