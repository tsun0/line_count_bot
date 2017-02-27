package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fujiwara/ridge"
	"github.com/line/line-bot-sdk-go/linebot"
)

var mux = http.NewServeMux()

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log.SetFlags(log.Lshortfile)
	// 新しいbotクライアントインスタンスを作成します
	// botにはClient型が入ります
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Println(err)
		return
	}

	// LINEプラットフォームのリクエストを受け取るためのHTTPサーバをセットアップ
	mux.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		// ParseRequestはリクエストを受け取って[]*Eventを返します
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		// count := countNumberOfPeople(events)
		// log.Println(count)

		wg := &sync.WaitGroup{}
		for _, event := range events {
			log.Printf("%+v", event)
			if event.ReplyToken == "00000000000000000000000000000000" {
				log.Println("verify message received")
				return
			}
			wg.Add(1)
			go func() {
				handleMessage(event, bot)
				wg.Done()
			}()

		}
		wg.Wait()

	})
	ridge.Run(":8080", "/api", mux)
}

func handleMessage(event *linebot.Event, bot *linebot.Client) {

	log.Println(event.Type)
	if event.Type == linebot.EventTypeFollow {

		m := generateJoinMessage()
		_, err := bot.ReplyMessage(event.ReplyToken, m...).Do()
		if err != nil {
			log.Println(err)
		}
	}

	if event.Type == linebot.EventTypeMessage {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:

			m, err := generateMessage(message)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.ReplyMessage(event.ReplyToken, m...).Do()
			if err != nil {
				log.Println(err)
			}
		}
	}

}
