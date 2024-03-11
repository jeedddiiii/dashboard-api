package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"

	"database/sql"

	_ "github.com/lib/pq"
)

type EmotionCount struct {
	Name    string
	Emotion string
	Count   int
}

var bot *linebot.Client

func buildMessage(rows *sql.Rows) (string, error) {
	emotionCounts := make(map[string]map[string]int)

	for rows.Next() {
		var ec EmotionCount
		err := rows.Scan(&ec.Name, &ec.Emotion, &ec.Count)
		if err != nil {
			return "", err
		}

		if _, ok := emotionCounts[ec.Name]; !ok {
			emotionCounts[ec.Name] = make(map[string]int)
		}

		emotionCounts[ec.Name][ec.Emotion] = ec.Count
	}

	var message strings.Builder
	for name, emotions := range emotionCounts {
		message.WriteString("คุณ " + name + "\n")
		for emotion, count := range emotions {
			message.WriteString(emotion + " " + strconv.Itoa(count) + " ครั้ง\n")
		}
		message.WriteString("\n")
	}

	return message.String(), nil
}

func BotInit() {
	var err error
	bot, err = linebot.New("3fd320c4fd8338776f0a82010cfdd0d2", "+1asOfxTo17N4YeN7/fFPRaoyygseuHsxzgdloG7wzmtxUXPS+BW27A4E/psoW/eZf88tNQxHp/tWCaAoiUiedaJi7LLHZisBH8CzViMyAlkalOvmy6mokMIGq1UWX1EqibP9JoPd17OyvOWDhWstwdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}
}

func SentMessage() {
	rows, err := db.Query(`
        SELECT name, emotion, COUNT(*) as count
        FROM transactions
        WHERE name != 'Unknown'
        GROUP BY name, emotion
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	message, err := buildMessage(rows)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.PushMessage("U499b899d22f729bb8056763b6e11f49a", linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Fatal(err)
	}
}
