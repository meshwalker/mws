package types

type ErrMsg struct {
	Status	string	`json:"status,omitempty"`
	Message	string	`json:"message,omitempty"`
}

type RespMsg struct {
	Status	string			`json:"status,omitempty"`
	Message	map[string]string	`json:"message,omitempty"`
}

type Token struct {

}