package main

import (
	"encoding/json"
	"os/exec"
	"net/http"

	"fmt"
	"time"
)

const (
	githubApiURL = "https://api.github.com/notifications"
	Token = "enter your token here!!!"
)

type Subject struct{
	Title string `json:"title"`
	URL string `json:"url"`
}

type Notification struct{
	Subject
	UpdatedAt time.Time `json:uploaded_at`
}

func sendDesktopNotification(title, url string) {
    cmd := exec.Command("notify-send", title, url)
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error sending desktop notification: %v\n", err)
    }
}

func main()  {
	// token := os.Getenv("GITHUB_ACCESS_TOKEN")
	// if token == "" {
    //     fmt.Println("Error: GitHub access token not provided")
    //     os.Exit(1)
    // }

	for{
		notifications,err := FetchNotifications()
		if err!= nil{
			fmt.Printf("Error fetching notifications: %v\n", err)
		}else {
			for _, notification := range notifications{
				title := notification.Title
                if title == "" {
                    title = "No title"
                }
				// fmt.Printf("Title: %s, URL: %s\n", notification.Subject.Title, notification.Subject.URL)
				sendDesktopNotification(notification.Title, notification.URL)
			}
		}

		time.Sleep(60 * time.Second) // Check for new notifications every minute
	}
}

func FetchNotifications() ([]Notification, error) {
    req, err := http.NewRequest("GET", githubApiURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "token "+Token)
    req.Header.Set("Accept", "application/vnd.github.v3+json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var notifications []Notification
    var data []map[string]json.RawMessage

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    for _, item := range data {
        var notification Notification
        if err := json.Unmarshal(item["subject"], &notification); err != nil {
            return nil, err
        }
        notification.URL = string(item["url"])
        notifications = append(notifications, notification)
    }

    return notifications, nil
}