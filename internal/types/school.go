package types

type School struct {
	UUID    string `json:"uuid,omitempty"`
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
