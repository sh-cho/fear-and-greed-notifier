package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	FgiApi = "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"
)

func main() {
	res := getFearAndGreedIndex()
	sendMesage(res)
}

func getFearAndGreedIndex() FgiResult {
	req, _ := http.NewRequest("GET", FgiApi, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", "fear-and-greed-index.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	r := FgiResult{}
	err := json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	return r
}

func sendMesage(fr FgiResult) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	//bot.Debug = true

	ch, err := strconv.ParseInt(os.Getenv("CHATID"), 10, 64)
	m := tgbotapi.NewMessage(ch, fr.toString())
	_, err = bot.Send(m)
	if err != nil {
		panic(err)
	}
}
