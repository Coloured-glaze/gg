package main

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"

	"github.com/Coloured-glaze/gg"
)

// go test -benchmem -bench .
// go test -benchmem -run=^$ -bench ^BenchmarkByte$

func BenchmarkByte(b *testing.B) {
	//by, _ := ioutil.ReadFile("../examples/james-webb.png")

	b.ResetTimer() //重置时间
	for i := 0; i < b.N; i++ {
		gg.GetWH("../examples/gopher.png")
	}
	/*
		goos: windows
		goarch: amd64
		cpu: Intel(R) Core(TM) i3-10100F CPU @ 3.60GHz
		BenchmarkByte-8
		39444             31425 ns/op            5800 B/op          7 allocs/op
		PASS
		ok      github.com/Coloured-glaze/gg/test       1.576s
	*/
}

func BenchmarkByte2(b *testing.B) {
	//by, _ := ioutil.ReadFile("../examples/james-webb.png")
	by, _ := ioutil.ReadFile("../examples/gopher.png")

	b.ResetTimer() //重置时间
	for i := 0; i < b.N; i++ {
		png.DecodeConfig(bytes.NewReader(by))
	}
	/*
	   goos: windows
	   goarch: amd64
	   cpu: Intel(R) Core(TM) i3-10100F CPU @ 3.60GHz
	   BenchmarkByte2-8   	 3977091	       298.1 ns/op	    1088 B/op	       3 allocs/op
	   PASS
	   ok  	_/e_/1/github/gg/test	1.514s
	*/
}
