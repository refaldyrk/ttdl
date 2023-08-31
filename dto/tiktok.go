package dto

type TikTokData struct {
	AuthorAvatar string `json:"author_avatar"`
	AuthorID     string `json:"author_id"`
	AuthorName   string `json:"author_name"`
	CommentCount int    `json:"comment_count"`
	CreateTime   string `json:"create_time"`
	ID           string `json:"id"`
	LikeCount    int    `json:"like_count"`
	ShareCount   int    `json:"share_count"`
	Success      bool   `json:"success"`
	Token        string `json:"token"`
}
