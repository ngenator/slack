package slack

type Message struct {
}

type Attachment struct {
	ID int `json:"id,omitempty"`

	Text     string `json:"text,omitempty"`
	PreText  string `json:"pretext,omitempty"`
	Fallback string `json:"fallback,omitempty"`

	Color string `json:"color,omitempty"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`

	ImageUrl string `json:"image_url,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`

	Fields []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
