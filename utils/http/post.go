package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type client struct {
	client *http.Client
}

type TYPE int

const (
	POST TYPE = iota
	GET
)

func NewClient() Client {
	client := new(client)
	client.client = &http.Client{}
	return client
}

func (c *client) Post(url string, body ...string) (int, []byte, error) {

	request, err := c.request(POST, url, body[0])
	if err != nil {
		return 0, nil, err
	}
	if len(body) > 1 {
		fmt.Printf("timestamp is %v", body[1])
		request.Header.Set("Timestamp", body[1])
	}

	request.Header.Set("Content-Type", "application/json")

	return c.do(request)
}

func (c *client) Get(url string) (int, []byte, error) {
	request, err := c.request(GET, url, "")
	if err != nil {
		return 0, nil, err
	}

	return c.do(request)
}

func (c *client) request(typ TYPE, url, body string) (*http.Request, error) {
	switch typ {
	case GET:
		return http.NewRequest("GET", url, nil)
	case POST:
		return http.NewRequest("POST", url, strings.NewReader(body))
	}
	return nil, errors.New("error TYPE!")
}

func (c *client) do(req *http.Request) (int, []byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, b, nil
}
