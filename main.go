package main

import (
	"encoding/base32"
	"flag"
	"io"
	"os"
)

func main() {
	decode := flag.Bool("d", false, "Decode data")
	// padding := flag.Bool("p", false, "Use padding, default is no padding")
	hex := flag.Bool("hex", false, "Use hex encoding, default is standard encoding in RFC4648")
	flag.Parse()
	var r io.Reader
	switch flag.Arg(0) {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		r = f
	}
	var enc base32.Encoding
	if *hex {
		enc = *base32.HexEncoding
	} else {
		enc = *base32.StdEncoding
	}
	var w io.WriteCloser
	w = os.Stdout
	if *decode {
		decoder := base32.NewDecoder(&enc, r)
		r = decoder
	} else {
		encoder := base32.NewEncoder(&enc, w)
		w = encoder
	}
	defer w.Close()
	_, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
}
