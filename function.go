package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
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

type Count struct {
	Total int `json:"total"`
}

func generateMessage(message *linebot.TextMessage) ([]linebot.Message, error) {

	var messages []linebot.Message

	if message.Text == "おしえて" {

		log.Println("おしえて")
		countJSON, err := requestCountsJSON()
		if err != nil {
			log.Println(err)
			text := linebot.NewTextMessage("エラーでーす")
			messages = append(messages, text)
			return messages, nil
		}
		log.Printf("%T\n", countJSON)

		count := &Count{}
		log.Println(count)

		if err := json.Unmarshal(countJSON, count); err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(count)

		messages = createReplyMessage(strconv.Itoa(count.Total))
		return messages, nil
	}
	messages = createReplyMessage("...")
	return messages, nil

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

	return messages

}

func requestCountsJSON() ([]byte, error) {

	res, err := http.Get("https://s3-ap-northeast-1.amazonaws.com/tsuchi-line-bot-count/count.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	count, err := readResponseBody(res)
	return count, err
}

func readResponseBody(res *http.Response) ([]byte, error) {

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(body))
	return body, nil
}
