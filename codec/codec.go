package codec

import (
	"bytes"
	"sync"

	"github.com/AlexBlayE/sonic/interfaces"
)

type codec struct {
	encoder interfaces.Encoder
	decoder interfaces.Decoder

	encBuffer *bytes.Buffer
	decBuffer *bytes.Buffer

	decLock sync.Locker
	encLock sync.Locker
}
