package modrinth

type SearchResponse struct {
	Hits []SearchHit `json:"hits"`
}
type SearchHit struct {
	ProjectID   string   `json:"project_id"`
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Versions    []string `json:"versions"`
	Categories  []string `json:"categories"`
}
