package codec

import (
	"bytes"
	"encoding/gob"
	"sync"
)

type GobCodec struct {
	codec
}

func NewGobCodec() *GobCodec {
	encBuffer := bytes.NewBuffer(nil)
	decBuffer := bytes.NewBuffer(nil)

	return &GobCodec{
		codec{
			encoder:   gob.NewEncoder(encBuffer),
			decoder:   gob.NewDecoder(decBuffer),
			encBuffer: encBuffer,
			decBuffer: decBuffer,
			encLock:   &sync.Mutex{},
			decLock:   &sync.Mutex{},
		},
	}
}

func (c *codec) Encode(msg any) ([]byte, error) {
	c.encLock.Lock()
	defer c.encLock.Unlock()

	err := c.encoder.Encode(msg)

	if err != nil {
		return nil, err
	}

	encoded := c.encBuffer.Bytes()
	c.encBuffer.Reset()

	return encoded, nil
}

func (c *codec) Decode(msg []byte, target any) error {
	c.decLock.Lock()
	defer c.decLock.Unlock()

	_, err := c.decBuffer.Write(msg)

	if err != nil {
		return err
	}

	err = c.decoder.Decode(target)
	c.decBuffer.Reset()

	return err
}
