package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"
)


/*

https://contest.yandex.ru/contest/19036/problems/H/
 */

func serve(ctx context.Context, cancel context.CancelFunc) (err error) {

	mux := http.NewServeMux()

	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		cancel()
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/validatePhoneNumber", func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("phone_number")
		type PhoneValidateResponse struct {
			Status     bool   `json:"status"`
			Normalized string `json:"normalized,omitempty"`
		}
		if number == "" {
			response := PhoneValidateResponse{Status: false}
			reponseJson, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(reponseJson))
			return
		}
		if checkIsPatternCorrect(number) == false {
			response := PhoneValidateResponse{Status: false}
			reponseJson, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(reponseJson))
			return

		}

		normalized := normalizePhone(number)

		response := PhoneValidateResponse{Status: true, Normalized: normalized}
		reponseJson, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reponseJson))

	})

	srv := &http.Server{
		Addr:    ":7777",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
func checkIsPatternCorrect(phone string) bool {

	patterns := []string{
		`^\+7\s(982|986|912|934)\s\d\d\d\s\d\d\s\d\d$`,
		`^\+7\s(982|986|912|934)\s\d\d\d\s\d\d\d\d$`,
		`^\+7\s\((982|986|912|934)\)\s\d\d\d-\d\d\d\d$`,
		`^\+7(982|986|912|934)\d\d\d\d\d\d\d$`,
		`^8\s(982|986|912|934)\s\d\d\d\s\d\d\d\d$`,
		`^8\s(982|986|912|934)\s\d\d\d\s\d\d\s\d\d$`,
		`^8\s\((982|986|912|934)\)\s\d\d\d-\d\d\d\d$`,
		`^8(982|986|912|934)\d\d\d\d\d\d\d$`,
	}
	for _, p := range patterns {
		matched, _ := regexp.MatchString(p, phone)
		if matched == true {
			return true
		}
	}
	return false
}

//+7-###-###-####
func normalizePhone(number string) string {
	number = strings.Replace(number, " ", "", -1)
	number = strings.Replace(number, "(", "", -1)
	number = strings.Replace(number, ")", "", -1)
	number = strings.Replace(number, "-", "", -1)
	number = strings.Replace(number, "+", "", -1)

	numrunes := []rune(number)
	formattedNumber := ""
	for idx, char := range numrunes {
		if idx == 0 {
			formattedNumber += `+7`
			continue
		}
		if idx == 1 || idx == 4 || idx == 7 {
			formattedNumber += "-" + string(char)
			continue
		}
		formattedNumber += string(char)
	}
	return formattedNumber
}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	if err := serve(ctx, cancel); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
