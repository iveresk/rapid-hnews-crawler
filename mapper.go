package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func mapMerge(ms, ams map[string]string) map[string]string {
	for k, v := range ms {
		ams[k] = v
	}
	return ams
}

func blogEqual(ms, ams map[string]string) (map[string]string, bool) {
	res := make(map[string]string)
	isequal := true

	for k, v := range ms {
		if ams[k] != v {
			isequal = false
			res[k] = v
		}
	}
	return res, isequal
}

func mapReader(fn string) (map[string]string, error) {
	file, ferr := ioutil.ReadFile(fn)
	if ferr != nil {
		return nil, ferr
	}
	file_res := make(map[string]string)
	jerr := json.Unmarshal(file, &file_res)
	if jerr != nil {
		return nil, jerr
	}
	return file_res, nil
}

func mapWriter(ms map[string]string, fn string) error {
	jsonout, jerr := json.Marshal(ms)
	if jerr != nil {
		return jerr
	}
	err := ioutil.WriteFile(fn, jsonout, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func mapToTelegram(ms map[string]string) (bool, error) {
	for k, v := range ms {
		reformed := reformat(k) + reformat(" ") + reformat(v)
		post_url := "https://api.telegram.org/bot" + os.Getenv("botToken") + "/sendMessage?chat_id=" + os.Getenv("chatID") + "&text=" + reformed
		resp, err := http.Post(post_url, "", nil)
		if err != nil {
			log.Println("Something with the Telegram or link, POST request didn't pass through")
			return false, err
		}
		if !strings.Contains(resp.Status, "200") {
			return false, nil
		}
	}
	return true, nil
}

func reformat(s string) string {
	res := ""
	for c := range s {
		switch s[c] {
		case '/':
			res = res + "%2F"
		case ':':
			res = res + "%3A"
		case ' ':
			res = res + "%20"
		case '!':
			res = res + "%21"
		case '"':
			res = res + "%22"
		case '#':
			res = res + "%23"
		case '$':
			res = res + "%24"
		case '%':
			res = res + "%25"
		case '&':
			res = res + "%26"
		case '\'':
			res = res + "%27"
		case '(':
			res = res + "%28"
		case ')':
			res = res + "%29"
		case '*':
			res = res + "%2A"
		case '+':
			res = res + "%2B"
		case ',':
			res = res + "%2C"
		case '-':
			res = res + "%2D"
		case '.':
			res = res + "%2E"
		case ';':
			res = res + "%3B"
		case '<':
			res = res + "%3C"
		case '=':
			res = res + "%3D"
		case '>':
			res = res + "%3E"
		case '?':
			res = res + "%3F"
		case '@':
			res = res + "%40"
		case '[':
			res = res + "%5B"
		case '\\':
			res = res + "%5C"
		case ']':
			res = res + "%5D"
		case '^':
			res = res + "%5E"
		case '_':
			res = res + "%5F"
		case '`':
			res = res + "%60"
		default:
			res = res + string(s[c])
		}
	}
	return res
}

// TODO mapreader and mapwriter for the MongoDB
