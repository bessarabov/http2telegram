package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"time"
)

func telegramSendMessage(chat_id, text string) {

	data := map[string]string{"chat_id": chat_id, "text": text}
	json_data, err := json.Marshal(data)

	_, err = http.Post("https://api.telegram.org/"+os.Getenv("TELEGRAM_TOKEN")+"/sendMessage", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	now := time.Now().Format(time.RFC3339)

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	normalizedHeadersText := ""

	headers := make([]string, 0, len(r.Header))

	for header, _ := range r.Header {
		headers = append(headers, header)
	}

	sort.Strings(headers)

	for _, header := range headers {
		for _, v := range r.Header[header] {
			normalizedHeadersText += header + ": " + v + "\n"
		}
	}

	fmt.Fprintf(os.Stderr, now+" "+ip+" "+r.Method+" "+r.URL.String()+"\n")

	text := ""
	text += "---" + "\n"
	text += time.Now().Format(time.RFC3339) + "\n"
	text += "Request from " + ip + "\n"
	text += r.Method + " " + r.URL.String() + "\n"
	text += "\n"
	text += "Normalized headers:" + "\n"
	text += normalizedHeadersText
	text += "\n"
	text += "Body:\n"
	text += string(bodyBytes) + "\n"
	text += "---" + "\n"

	telegramSendMessage(os.Getenv("CHAT_ID"), text)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
