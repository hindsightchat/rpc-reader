package richpresence

import "time"

// opcodes
const (
	OpHandshake = iota
	OpFrame
	OpClose
	OpPing
	OpPong
)

// known commands
const (
	CmdDispatch           = "DISPATCH"
	CmdSetActivity        = "SET_ACTIVITY"
	CmdSubscribe          = "SUBSCRIBE"
	CmdUnsubscribe        = "UNSUBSCRIBE"
	CmdGetGuild           = "GET_GUILD"
	CmdGetGuilds          = "GET_GUILDS"
	CmdGetChannel         = "GET_CHANNEL"
	CmdGetChannels        = "GET_CHANNELS"
	CmdGetSelectedVoice   = "GET_SELECTED_VOICE_CHANNEL"
	CmdSelectVoiceChannel = "SELECT_VOICE_CHANNEL"
	CmdSelectTextChannel  = "SELECT_TEXT_CHANNEL"
	CmdVoiceSettings      = "GET_VOICE_SETTINGS"
	CmdSetVoiceSettings   = "SET_VOICE_SETTINGS"
)

// events
const (
	EvtReady              = "READY"
	EvtError              = "ERROR"
	EvtGuildStatus        = "GUILD_STATUS"
	EvtGuildCreate        = "GUILD_CREATE"
	EvtChannelCreate      = "CHANNEL_CREATE"
	EvtVoiceStateCreate   = "VOICE_STATE_CREATE"
	EvtVoiceStateUpdate   = "VOICE_STATE_UPDATE"
	EvtVoiceStateDelete   = "VOICE_STATE_DELETE"
	EvtVoiceSettingsUpdate = "VOICE_SETTINGS_UPDATE"
	EvtVoiceConnectionStat = "VOICE_CONNECTION_STATUS"
	EvtSpeakingStart      = "SPEAKING_START"
	EvtSpeakingStop       = "SPEAKING_STOP"
	EvtActivityJoin       = "ACTIVITY_JOIN"
	EvtActivitySpectate   = "ACTIVITY_SPECTATE"
	EvtActivityJoinReq    = "ACTIVITY_JOIN_REQUEST"
)

type Handshake struct {
	V        int    `json:"v"`
	ClientID string `json:"client_id"`
}

type Frame struct {
	Cmd   string      `json:"cmd"`
	Nonce string      `json:"nonce,omitempty"`
	Evt   string      `json:"evt,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Args  interface{} `json:"args,omitempty"`
}

type Activity struct {
	State      string      `json:"state,omitempty"`
	Details    string      `json:"details,omitempty"`
	Timestamps *Timestamps `json:"timestamps,omitempty"`
	Assets     *Assets     `json:"assets,omitempty"`
	Party      *Party      `json:"party,omitempty"`
	Secrets    *Secrets    `json:"secrets,omitempty"`
	Buttons    []Button    `json:"buttons,omitempty"`
	Instance   bool        `json:"instance,omitempty"`
}

type Timestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type Assets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type Party struct {
	ID   string `json:"id,omitempty"`
	Size []int  `json:"size,omitempty"` // [current, max]
}

type Secrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type Button struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type SetActivityArgs struct {
	PID      int       `json:"pid"`
	Activity *Activity `json:"activity"`
}

type ReadyData struct {
	V    int       `json:"v"`
	User *UserData `json:"user,omitempty"`
}

type UserData struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar,omitempty"`
}

type PresenceUpdate struct {
	ClientID  string    `json:"client_id"`
	Activity  *Activity `json:"activity"`
	PID       int       `json:"pid"`
	Timestamp time.Time `json:"timestamp"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
