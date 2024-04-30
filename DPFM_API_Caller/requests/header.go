package requests

type Header struct {
	Site				int     `json:"Site"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
