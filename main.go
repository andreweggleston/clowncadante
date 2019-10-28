package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var api = slack.New(
	os.Getenv("OAUTH_TOKEN"),
	slack.OptionDebug(true),
	slack.OptionLog(log.New(os.Stdout, "clown-bot: ", log.Lshortfile|log.LstdFlags)),
)

func main() {
	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: os.Getenv("VERIFICATION_TOKEN")}))
		if e != nil {
			fmt.Println("Error: ", e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text")
			_, _ = w.Write([]byte(r.Challenge))
		}
		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				fmt.Println(ev.Username, ": ", ev.Text)
				if ev.User == os.Getenv("MERKY_UID"){
					fmt.Println("A clown has been detected")
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					num := r.Intn(100)
					fmt.Println("Rand genned: ", num)
					if num > 80 {
						fmt.Println("ITS CLOWN TIME!!!!!")
						if api.AddReaction("clown_face", slack.NewRefToMessage(ev.Channel, ev.TimeStamp)) != nil {
							fmt.Println("Couldn't add reaction")
						}
					}
				}
			//case *slackevents.:
			//
			}
		}
	})
	fmt.Println("[INFO] Server listening")
	_ = http.ListenAndServe(":8080", nil)
}
