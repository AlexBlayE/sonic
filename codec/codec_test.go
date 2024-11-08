package codec

import "testing"

func TestGobEncodeDecode(t *testing.T) {
	c := NewGobCodec()

	msg := "Test"

	enc, err := c.Encode(msg)

	if err != nil {
		t.Fatal(err)
		return
	}

	var dec string
	err = c.Decode(enc, &dec)

	if err != nil {
		t.Fatal(err)
		return
	}

	if msg != dec {
		t.Fatal("Not the same")
	}
}
