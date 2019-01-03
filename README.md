# Netstring

Yet another netstring implementation! Simple usage:


``` golang
	payload := []byte("12:Hello world!,0:,13:Goodbye world,")
	res, err := netstring.Decode(payload)
	// res is not the list of bytes containing:
	//  "Hello world!", "" and "Goodbye world"

    items := [][]byte{
		[]byte("Hello world!"),
		[]byte(""),
		[]byte("Goodbye world"),
	}
	encoded, err := netstring.Encode(byte_items...)
	// encoded is now equal to payload
```

