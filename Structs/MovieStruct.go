package Structs

type Movie struct {
	ID          int
	MovieTitle  string
	Director    string
	Producer    string
	Description string
	Realesed    int
	Category    []Category
	Cards       []string // TODO: make cards
}
