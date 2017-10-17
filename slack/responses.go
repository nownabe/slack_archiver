package slack

type response interface {
	ok() bool
	error() string
}

type responseCommon struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func (r responseCommon) ok() bool {
	return r.OK
}

func (r responseCommon) error() string {
	return r.Error
}

type channelsListResp struct {
	responseCommon
	Channels []Channel `json:"channels"`
}

type channelsHistoryResp struct {
	responseCommon
	History
}
