package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/appengine/urlfetch"
)

// Client is a slack api client.
type Client struct {
	token string
}

// New returns a new Client instance.
func New(token string) *Client {
	return &Client{token: token}
}

func (c *Client) call(ctx context.Context, method string, values url.Values, rv interface{}) error {
	endpoint := "https://slack.com/api/" + method

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := urlfetch.Client(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
	case 429:
		retryAfter, err := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			return err
		}
		return RateLimitError{
			slackError: slackError{Message: fmt.Sprintf("HTTP Error: %s", resp.Status)},
			RetryAfter: retryAfter,
		}
	default:
		return fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &rv)
	if err != nil {
		return err
	}

	if r, ok := rv.(response); ok {
		if !r.ok() {
			return fmt.Errorf("Slack API Errlr: %s", r.error())
		}
	}

	return nil
}
