package linkShortener

type ShortenedLinkModel struct {
	FullLink string `json:"fullLink"`

	ShortedSlug string `json:"shortedSlug"`
}

type LinkShortenerService struct {
}

func (s *LinkShortenerService) createShortedLink(name string) {
	
}
