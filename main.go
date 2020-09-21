package main

/*

$> afk --message "Lunch!" --emoji ":salad:" --duration "1h"
Slack status until 3:23pm: :salad: Lunch!
<press any key to stop>

*/

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const defaultDuration = "1h"

var emoji string
var duration time.Duration
var message string
var doNotDisturb bool

type slackAPIResponse struct {
	Ok    bool   `json:"ok,omitempty"`
	Error string `json:"error,omitempty"`
}

type slackProfile struct {
	slackAPIResponse
	Profile struct {
		slackStatus
	} `json:"profile"`
}
type slackStatus struct {
	StatusText  string `json:"status_text"` // max 100 chars
	StatusEmoji string `json:"status_emoji"`
	// epoch time, setting to 0 or ommitting this in API call results in a status that never expires
	StatusExpiration int64 `json:"status_expiration"`
}

func main() {
	flag.StringVar(&message, "message", "Away from keyboard", "The message to display while AFK")
	flag.StringVar(&emoji, "emoji", ":speech-bubble:", "Emoji to display while AFK")
	flag.DurationVar(&duration, "duration", getDefaultDuration(), "How long the AFK status should last")
	flag.BoolVar(&doNotDisturb, "dnd", false, "Enable Do Not Disturb")
	flag.Parse()

	currentStatus := getCurrentStatus()
	fmt.Printf("=== Current Status ===\n%v\n", currentStatus.String())

	newStatus := slackStatus{
		StatusEmoji:      emoji,
		StatusText:       message,
		StatusExpiration: time.Now().Add(duration).Unix(),
	}
	fmt.Printf("=== New Status ===\n%v\n", newStatus.String())
	// set new status
	if doNotDisturb {
		// TODO: set DND
	}
	// wait for expiration or input
	// replace previous status
}

func (status *slackStatus) String() string {
	emoji := status.StatusEmoji
	if emoji == "" {
		emoji = "<none>"
	}
	text := status.StatusText
	if text == "" {
		text = "<none>"
	}
	expirationText := ""
	if status.StatusExpiration == 0 {
		expirationText = "<none>"
	} else {
		expirationText = time.Unix(status.StatusExpiration, 0).String()
		timeRemaining := timeDurationInWords(time.Until(time.Unix(status.StatusExpiration, 0)))
		expirationText = fmt.Sprintf("%v (%v from now)", expirationText, timeRemaining)
	}
	return fmt.Sprintf("Emoji: %v\nText: %v\nExpires: %v", emoji, text, expirationText)
}
func getCurrentStatus() slackStatus {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://slack.com/api/users.profile.get", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+getSlackToken())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var slackProfile slackProfile
	err = json.Unmarshal(body, &slackProfile)
	if err != nil {
		log.Fatal(err)
	}
	if slackProfile.Ok != true {
		log.Fatal("Slack API error: " + slackProfile.Error)
	}
	return slackProfile.Profile.slackStatus
}

func getSlackToken() string {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		configFilePath := path.Join(os.Getenv("HOME"), ".afk-slack.yml")
		token = getSlackTokenFromFile(configFilePath)
		if token == "" {
			log.Fatal("Could not find a Slack API token. Checked ENV var: $SLACK_API_TOKEN and file: " + configFilePath)
		}
	}

	return token
}

func getSlackTokenFromFile(configFilePath string) string {
	token := ""
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return token
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, ":")
		if splitText[0] == "token" {
			token = strings.Trim(splitText[1], " ")
			break
		}
	}
	return token
}

func getDefaultDuration() time.Duration {
	duration, err := time.ParseDuration(defaultDuration)
	if err != nil {
		log.Fatal(err)
	}
	return duration
}
