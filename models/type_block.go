package models

// Block .
type Block struct {
	GormModel
	Did      string `json:"did"`
	Cid      string `json:"cid"`
	FileName string `json:"file_name"`
}
