package Structs

// TODO: Make more params (like: age of audience, url of picture and etc)
// TODO: Make a genres
type Movie struct {
	ID          int
	MovieTitle  string
	Director    string
	Producer    string
	Description string
	Realesed    int
	Category    []Category
	Cards       []Cards // TODO: Linking movies with cards
	URLPoster   string
}
