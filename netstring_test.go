package netstring

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestNS(t *testing.T) {
	var encoded = []string{
		"Hello world!",
		"",
		"Goodbye world",
	}
	ns := NewNetString()
	ns.EncodeString(encoded[0], encoded[1], encoded[2])
	checkError(t, ns.err)
	out := ns.buffer.String()
	assertEqual(t, "12:Hello world!,0:,13:Goodbye world,", string(out))

	ns = NewNetString(out)
	decoded := ns.DecodeString()
	checkError(t, ns.err)
	for pos, part := range decoded {
		assertEqual(t, part, encoded[pos])
	}
}
