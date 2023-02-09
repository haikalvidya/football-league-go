package payload

type ContainLetterRequest struct {
	FirstWord  string `json:"first_word"`
	SecondWord string `json:"second_word"`
}
