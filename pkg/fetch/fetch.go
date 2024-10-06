package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type Fetch struct {
	token  *string
	logger *zap.SugaredLogger
}

func NewFetch(ctx context.Context, token *string) *Fetch {
	return &Fetch{
		token:  token,
		logger: logging.FromContext(ctx),
	}
}

func (f Fetch) Get(url string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating GET request: %w", err)
	}

	f.addAuthorizationHeader(req)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			f.logger.Error(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	return json.Unmarshal(body, response)
}

// PostForm sends a POST request to the given url with the given data and stores the response in the given response.
// The Content-Type of the request is set to application/x-www-form-urlencoded.
func (f Fetch) PostForm(url string, data url.Values, response interface{}) error {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return f.post(req, response)
}

// PostJson sends a POST request to the given url with the given data and stores the response in the given response.
// The Content-Type of the request is set to application/json.
func (f Fetch) PostJson(url string, data, response interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error creating POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return f.post(req, response)
}

func (f Fetch) post(req *http.Request, response interface{}) error {

	f.addAuthorizationHeader(req)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	return json.Unmarshal(body, response)
}

func (f Fetch) addAuthorizationHeader(req *http.Request) {
	if f.token != nil {
		req.Header.Set("Authorization", "Bearer "+*f.token)
	}
}
