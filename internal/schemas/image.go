package schemas

type ImageAdded struct {
	URL string
}

func NewImageAdded(URL string) *ImageAdded {
	return &ImageAdded{
		URL: URL,
	}
}
