package interfaces

type Codec interface {
	Encode(msg any) ([]byte, error)
	Decode(msg []byte, target any) error
}
