package netstring

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)",
		a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestString(t *testing.T) {
	var items = []string{
		"Hello world!",
		"",
		"Goodbye world",
	}
	ns := NewNetString()
	ns.EncodeString(items[0], items[1], items[2])
	checkError(t, ns.err)
	out := ns.buffer.String()
	assertEqual(t, "12:Hello world!,0:,13:Goodbye world,", string(out))

	ns = NewNetString(out)
	decoded := ns.DecodeString()
	checkError(t, ns.err)
	for pos, part := range decoded {
		assertEqual(t, part, items[pos])
	}
}

func TestBytes(t *testing.T) {
	items := []string{
		"Hello world!",
		"",
		"Goodbye world",
	}
	payload := []byte("12:Hello world!,0:,13:Goodbye world,")

	// Test decode
	res, err := Decode(payload)
	checkError(t, err)
	for pos, part := range res {
		assertEqual(t, string(part), items[pos])
	}

	// Test encode
	byte_items := [][]byte{
		[]byte("Hello world!"),
		[]byte(""),
		[]byte("Goodbye world"),
	}
	encoded, err := Encode(byte_items...)
	checkError(t, err)
	assertEqual(t, string(encoded), string(payload))
}
