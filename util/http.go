package util

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type AuthCredentials struct {
	IsBasicAuth bool
	APIKey      string
	User        string
	Password    string
}

func HystrixPost(
	method string,
	url string,
	data io.Reader,
	auth AuthCredentials,
	commandName string,
	fallback func(err error) error) error {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Panic("create new request object with error: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	if auth.IsBasicAuth {
		req.SetBasicAuth(auth.User, auth.Password)
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth.APIKey))
	}

	resultChannel := make(chan string, 1)
	errChannel := hystrix.Go(commandName, func() error {
		response, e := client.Do(req)
		defer response.Body.Close()

		log.Debugf("Sent %s request to: %s with body: %s", method, url, data)

		if e != nil {
			return e
		}

		body, _ := ioutil.ReadAll(response.Body)

		if !(response.StatusCode >= 200 && response.StatusCode < 300) {
			log.Warnf("Request did not return success with status: %s, %d", body, response.StatusCode)
		}

		log.Debugf("Request returned with response: %s", body)

		resultChannel <- string(body)

		return nil
	}, fallback)

	select {
	case result := <- resultChannel:
		log.Info(result)
	case err := <- errChannel:
		log.Warnf("Hystrix failure with error: %v", err)
		return err
	}

	return nil
}
