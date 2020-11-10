package model

type Tag struct {
	Id   int
	Name string
}

type Category struct {
	Id   int
	Name string
}

type Pet struct {
	Id        int    `json:"id"`
	Category  Category `json:"category,omitempty"`
	Name      string   `json:"name"`
	PhotoURLs []string `json:"photoUrls,omitempty"`
	Tags      []Tag    `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}
