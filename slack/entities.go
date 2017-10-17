package slack

// Channel is a channel.
type Channel struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Created    int    `json:"created"`
	Creator    string `json:"creator"`
	IsArchived bool   `json:"is_archived"`
	IsMember   bool   `json:"is_member"`
	NumMembers int    `json:"num_members"`
	Topic      struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"topic"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"purpose"`
}

// History is a history of a channel.
type History struct {
	Latest   string    `json:"latest"`
	Messages []Message `json:"messages"`
	HasMore  bool      `json:"has_more"`
}

// Message is a message.
type Message struct {
	Type      string     `json:"type"`
	TS        string     `json:"ts"`
	User      string     `json:"user"`
	Text      string     `json:"text"`
	IsStarred bool       `json:"is_starred"`
	Reactions []Reaction `json:"reactions"`
}

// Reaction is a reaction.
type Reaction struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Users []string `json:"users"`
}
