package console

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"redis_client/pkg/cache"
	"strconv"
	"strings"
)

var url = os.Getenv("redis_url")
var db = 0
var promptTemplate = "%s[%s]"
var historyLimit = 10
var redisClient *redis.Client
var ipDb = make(map[string]int, 0)

func Execute(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(w, "cannot read body", err)
			return
		}
		var request Request
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(w, "cannot unmarshal body", err)
			return
		}
		hist := History{Input: request.Command}
		params := strings.Split(request.Command, " ")
		if len(params) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(w, "unknown command", err)
			return
		}
		theCommand := strings.ToLower(params[0])

		if currDb, ok := ipDb[r.RemoteAddr]; !ok {
			ipDb[r.RemoteAddr] = 0
			db = currDb
		}
		redisClient = cache.NewClient(db, url)
		if theCommand == "select" {
			newDb, err := strconv.Atoi(params[1])
			if err != nil {
				hist.Output = "(error) ERR invalid DB index"
				res, err := updateAndGetHistory(r.RemoteAddr, hist)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				respondWithHistory(w, res)
				return

			} else {
				redisClient = cache.NewClient(newDb, url)
				hist.Output = "OK"
				ipDb[r.RemoteAddr] = newDb
				res, err := updateAndGetHistory(r.RemoteAddr, hist)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				respondWithHistory(w, res)
				return

			}
		} else {
			if comm, ok := cache.Commands[theCommand]; !ok {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("command not supported: " + theCommand))
				return
			} else {
				result := comm(redisClient, context.Background(), params[1:]...)
				hist.Output = result
				res, err := updateAndGetHistory(r.RemoteAddr, hist)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				respondWithHistory(w, res)
				return
			}

		}

	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, err := template.ParseGlob(consoleTemplate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(w, "Unable to load template")
		}
        ipDb[r.RemoteAddr] = 0
		res, err := updateAndGetHistory(r.RemoteAddr, History{})
		resp := Response{History: res, Prompt: fmt.Sprintf(promptTemplate, url, strconv.Itoa(db))}

		t.Execute(w, resp)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed: " + r.Method ))
}

func respondWithHistory(w http.ResponseWriter, res []History) {
	resp := Response{History: res, Prompt: fmt.Sprintf(promptTemplate, url, strconv.Itoa(db))}
	marsh, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marsh)
	return
}

func updateAndGetHistory(ip string, item History) ([]History, error) {
	fileName := fmt.Sprintf("./%s_history", ip)
	res := make([]History, 0)
	if _, err := os.Stat(fileName); err == nil {
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Println("error reading history file", err)
			return nil, err
		}
		history := make([]History, 0)
		err = json.Unmarshal(file, &history)
		if err != nil {
			log.Println("can't unmarshal history", err)
			return nil, err
		}
		res = history[:historyLimit-1]
		if item.Input != "" {
			res = append(res, item)
		}
		os.Remove(fileName)
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist

	} else {
		log.Println("unexpected error", err)
		return nil, err
	}

	marshalled, err := json.Marshal(res)
	if err != nil {
		log.Println("error marshaling", err)
		return nil, err
	}
	err = ioutil.WriteFile(fileName, marshalled, 0644)
	return res, nil
}

type History struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Request struct {
	Command string `json:"command"`
}

type Response struct {
	History []History `json:"history"`
	Prompt  string    `json:"prompt"`
}
