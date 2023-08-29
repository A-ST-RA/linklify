package encoder

type EncoderService interface {
	Encode(n string) (string, error)
}
