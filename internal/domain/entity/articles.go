package entity

type Articles struct {
	ArticlesID int32   `json:"articles_id"`
	Title      string  `json:"title"`
	Body       string  `json:"body"`
	AuthorID   int32   `json:"author_id"`
	Image      []byte  `json:"image"`
	Created_at []uint8 `json:"created_at"`
	Updated_at []uint8 `json:"updated_At"`
}
