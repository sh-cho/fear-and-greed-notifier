package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	FgiApi = "https://fear-and-greed-index.p.rapidapi.com/v1/fgi"
)

type VVT struct {
	Value     int
	ValueText string
}

type FgiResult struct {
	LastUpdated struct {
		EpochUnixSeconds int
		HumanDate        string
	}
	Fgi struct {
		Now           VVT
		PreviousClose VVT
		OneWeekAgo    VVT
		OneMonthAgo   VVT
		OneYearAgo    VVT
	}
}

func (fr FgiResult) toString() string {
	return fmt.Sprintf(`[lastUpdate: %s]
- now: %d (%s)
- prev: %d (%s)
- 1w ago: %d (%s)
- 1m ago: %d (%s)
- 1y ago: %d (%s)`, fr.LastUpdated.HumanDate,
		fr.Fgi.Now.Value, fr.Fgi.Now.ValueText,
		fr.Fgi.PreviousClose.Value, fr.Fgi.PreviousClose.ValueText,
		fr.Fgi.OneWeekAgo.Value, fr.Fgi.OneWeekAgo.ValueText,
		fr.Fgi.OneMonthAgo.Value, fr.Fgi.OneMonthAgo.ValueText,
		fr.Fgi.OneYearAgo.Value, fr.Fgi.OneYearAgo.ValueText)
}

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
