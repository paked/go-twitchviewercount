package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var stream = flag.String("stream", "tehhcwool", "Your (or someone elses) streamname")
var clientID = flag.String("clientID", "9po4ts2jz2niigqq3o9gtt2ntw69njf", "Your client ID (set in settings/connections)")

type Repsonse struct {
	Stream StreamInfo `json: stream`
}
type StreamInfo struct {
	Viewers int `json: "viewers"`
}

func main() {
	flag.Parse()
	ticker := time.NewTicker(5 * time.Second)

	url := "https://api.twitch.tv/kraken/streams/" + *stream + "?client_id=" + *clientID
	log.Printf("Final URL is %v", url)

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(url)
			if err != nil {
				log.Println("Something went wrong with your internet/connection/twitches servers")
				break
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			response := Repsonse{StreamInfo{}}
			json.Unmarshal(body, &response)
			log.Printf("%v viewers are watching "+*stream+"'s stream", response.Stream.Viewers)

		}
	}
}
