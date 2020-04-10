package dataStructs

// Identification object that represents an identification in the World State
type Identification struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Ip        string `json:"ip"`
	PublicKey string `json:"publicKey"`
}