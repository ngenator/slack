package slack

type Message struct {
}

type Attachment struct {
	ID       int     `json:"id,omitempty"`
	Text     string  `json:"text"`
	Fallback string  `json:"fallback,omitempty"`
	Color    string  `json:"color,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}
