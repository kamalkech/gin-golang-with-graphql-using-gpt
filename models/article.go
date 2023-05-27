package models

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []Article{
	{
		ID:      "1",
		Title:   "First Article",
		Content: "This is the content of the first article.",
	},
	{
		ID:      "2",
		Title:   "Second Article",
		Content: "This is the content of the second article.",
	},
}

func GetArticles() []Article {
	return articles
}

func GetArticleByID(id string) *Article {
	for _, article := range articles {
		if article.ID == id {
			return &article
		}
	}
	return nil
}

