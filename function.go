package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func generateJoinMessage() []linebot.Message {

	var messages []linebot.Message
	text := linebot.NewTextMessage("友だち追加ありがとうございます!")
	log.Println(text.Text)
	sticker := linebot.NewStickerMessage("2", "144")
	log.Println(sticker.StickerID)
	messages = append(messages, text, sticker)
	log.Println(messages)
	return messages

}

func generateMessage(message *linebot.TextMessage) []linebot.Message {

	var messages []linebot.Message

	if message.Text == "おしえて" {
		messages = createReplyMessage("いいよー")
		return messages
	}
	messages = createReplyMessage("...")
	return messages

}

func randomSticker() string {

	max, min := 179, 140
	return strconv.Itoa(rand.Intn(max-min) + min)

}

func createReplyMessage(text string) []linebot.Message {

	message := linebot.NewTextMessage(text)
	sticker := linebot.NewStickerMessage("2", randomSticker())
	var messages []linebot.Message
	messages = append(messages, message, sticker)
	log.Println(messages)
	return messages

}
