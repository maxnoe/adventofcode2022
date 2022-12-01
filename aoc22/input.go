package aoc22

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var NoCookie = errors.New("You need to set the AOC_SESSION env variable")

func GetInput(year int, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", NoCookie
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	input := string(body)
	return input, nil
}
