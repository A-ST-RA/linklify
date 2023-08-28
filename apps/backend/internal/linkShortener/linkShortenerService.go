package linkShortener

type ShortenedLinkModel struct {
	FullLink string `json:"fullLink"`

	ShortedSlug string `json:"shortedSlug"`
}

type LinkShortenerService struct {
}

func NewLinkShortenerService() *LinkShortenerService {
	return &LinkShortenerService{}
}

func (s *LinkShortenerService) createShortedLink(name string) {

}
