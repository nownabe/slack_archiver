package slack

type slackError struct {
	Message string
}

// RateLimitError is an error for rate limit.
type RateLimitError struct {
	slackError
	RetryAfter int
}

func (e slackError) Error() string {
	return e.Message
}
