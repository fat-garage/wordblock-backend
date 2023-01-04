package request

type UploadBlock struct {
	Content string `json:"content" binding:"required"`
	DID     string `json:"did" binding:"required"`
}
