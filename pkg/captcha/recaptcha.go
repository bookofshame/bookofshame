package captcha

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/fetch"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type ReCaptcha struct {
	host    string
	secret  string
	skip    bool
	fetcher *fetch.Fetch
	logger  *zap.SugaredLogger
}

const (
	ReCaptchaPayloadSecret   = "secret"
	ReCaptchaPayloadResponse = "response"
	ReCaptchaScoreThreshold  = 0.5
)

type ReCaptchaResponse struct {
	Success bool    `json:"success"`
	Score   float32 `json:"score"`
	Action  string  `json:"action"`
}

func NewReCaptcha(ctx context.Context, cfg config.Config) *ReCaptcha {
	return &ReCaptcha{
		skip:    cfg.Env == "development",
		secret:  cfg.ReCaptchaSecretKey,
		host:    cfg.ReCaptchaHost,
		fetcher: fetch.NewFetch(ctx, nil),
		logger:  logging.FromContext(ctx),
	}
}

func (c *ReCaptcha) Verify(token string) error {
	// skip if we are in localhost
	if c.skip {
		return nil
	}

	payload := url.Values{
		ReCaptchaPayloadSecret:   {c.secret},
		ReCaptchaPayloadResponse: {token},
	}

	response := &ReCaptchaResponse{}

	err := c.fetcher.PostForm(c.host, payload, response)

	if !response.Success || err != nil {
		return fmt.Errorf("failed to verify captcha. refresh the page and try again")
	}

	if response.Score < ReCaptchaScoreThreshold {
		return fmt.Errorf("possible suspicious activity. please try again later")
	}

	return nil
}
