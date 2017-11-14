package slack

import (
	"fmt"
	"net/url"

	"appengine"
)

// ChannelsArchive archives a channel.
func (c *Client) ChannelsArchive(ctx appengine.Context, channel string) error {
	values := url.Values{}
	values.Add("channel", channel)

	resp := responseCommon{}

	if err := c.call(ctx, "channels.archive", values, &resp); err != nil {
		return err
	}

	return nil
}

// ChannelsHistory fetches history of messages and events from a channel.
func (c *Client) ChannelsHistory(ctx appengine.Context, channel string, count int) (*History, error) {
	values := url.Values{}
	values.Add("channel", channel)
	values.Add("count", fmt.Sprintf("%d", count))

	resp := channelsHistoryResp{}

	if err := c.call(ctx, "channels.history", values, &resp); err != nil {
		return nil, err
	}

	return &resp.History, nil
}

// ChannelsList lists all channels in a Slack team.
func (c *Client) ChannelsList(ctx appengine.Context, excludeArchived bool, excludeMembers bool) ([]Channel, error) {
	values := url.Values{}
	if excludeArchived {
		values.Add("exclude_archived", "1")
	}
	if excludeMembers {
		values.Add("exclude_members", "1")
	}

	resp := channelsListResp{}

	err := c.call(ctx, "channels.list", values, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Channels, nil
}
