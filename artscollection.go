package artscollection

type Loader interface {
	Load() *Collection
}

type PropLoader interface {
	LoadProperties(string) *[]Property
}

type ArtwLoader interface {
	LoadArtworks(string) *[]Artwork
}

type ArtwAdder interface {
	AddArtwork(*Artwork) error
}

type ArtwUpdater interface {
	UpdateArtwork(*Artwork) error
}

type Collection struct {
	Artworks       []Artwork
	Title          string
	Description    string
	DataProperties []Property
}

type Artwork struct {
	Path       string
	Filename   string
	Properties []Property
	Data       Data
}

type Property struct {
	ID     string
	Name   string
	Type   string
	Group  string
	Order  int
	Select []string
}

type Data struct {
	Strings  map[string]string
	Integers map[string]int
}
