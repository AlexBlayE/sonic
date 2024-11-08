package interfaces

type Decoder interface {
	Decode(any) error
}
