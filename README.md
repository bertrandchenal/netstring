# Netstring

Yet another netstring implementation! Simple usage:


```
	payload := []byte("12:Hello world!,0:,13:Goodbye world,")
	res, err := Decode(payload)
	// res is not the list of bytes containing:
	//  "Hello world!", "" and "Goodbye world"

    items := [][]byte{
		[]byte("Hello world!"),
		[]byte(""),
		[]byte("Goodbye world"),
	}
	encoded, err := Encode(byte_items...)
	// encoded is now equal to payload
```

