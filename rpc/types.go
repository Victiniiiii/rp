package rpc

import (
	"time"

	"github.com/Victiniiiii/rp/ipc"
)

type Handshake struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

type Args struct {
	Pid      int              `json:"pid"`
	Activity *PayloadActivity `json:"activity"`
}

type PayloadActivity struct {
	Details    string             `json:"details,omitempty"`
	State      string             `json:"state,omitempty"`
	Assets     PayloadAssets      `json:"assets,omitempty"`
	Party      *PayloadParty      `json:"party,omitempty"`
	Timestamps *PayloadTimestamps `json:"timestamps,omitempty"`
	Secrets    *PayloadSecrets    `json:"secrets,omitempty"`
	Buttons    []*PayloadButton   `json:"buttons,omitempty"`
	Type       int                `json:"type,omitempty"`
}

type PayloadAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type PayloadParty struct {
	ID   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type PayloadTimestamps struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type PayloadSecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

type PayloadButton struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}

type Client struct {
	ClientID string
	IPC      *ipc.IPC
	Logged   bool
}

// Activity holds the data for discord rich presence
type Activity struct {
	// What the player is currently doing
	Details string
	// The user's current party status
	State string
	// The id for a large asset of the activity, usually a snowflake
	LargeImage string
	// Text displayed when hovering over the large image of the activity
	LargeText string
	// The id for a small asset of the activity, usually a snowflake
	SmallImage string
	// Text displayed when hovering over the small image of the activity
	SmallText string
	// Information for the current party of the player
	Party *Party
	// Unix timestamps for start and/or end of the game
	Timestamps *Timestamps
	// Secrets for Rich Presence joining and spectating
	Secrets *Secrets
	// Clickable buttons that open a URL in the browser
	Buttons []*Button
}

// Button holds a label and the corresponding URL that is opened on press
type Button struct {
	// The label of the button
	Label string
	// The URL of the button
	Url string
}

// Party holds information for the current party of the player
type Party struct {
	// The ID of the party
	ID string
	// Used to show the party's current size
	Players int
	// Used to show the party's maximum size
	MaxPlayers int
}

// Timestamps holds unix timestamps for start and/or end of the game
type Timestamps struct {
	// unix time (in milliseconds) of when the activity started
	Start *time.Time
	// unix time (in milliseconds) of when the activity ends
	End *time.Time
}

// Secrets holds secrets for Rich Presence joining and spectating
type Secrets struct {
	// The secret for a specific instanced match
	Match string
	// The secret for joining a party
	Join string
	// The secret for spectating a game
	Spectate string
}
