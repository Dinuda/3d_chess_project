package messages

const (
	InMessageNick      = "nick"
	InMessageStartGame = "start-game"
	InMessageJoinGame  = "join-game"
	InMessageMove      = "move"
	InMessagePlace     = "place"
)

type BaseMessage struct {
	Msg string `json:"msg"`
}

type MessageNick struct {
	Nick string `json:"nick"`
}

type MessageJoinGame struct {
	ID string `json:"id"`
}

type MessageMove struct {
	Secret string `json:"secret"`
	FromX  int    `json:"fx"`
	FromY  int    `json:"fy"`
	ToX    int    `json:"tx"`
	ToY    int    `json:"ty"`
}

type MessagePlace struct {
	Secret string `json:"secret"`
	ID     string `json:"id"`
	ToX    int    `json:"tx"`
	ToY    int    `json:"ty"`
}
