package text

type TextEncodable interface {
	TextEncode(encoder *Encoder, option *Option) error
}
