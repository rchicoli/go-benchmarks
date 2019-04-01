package main

import (
	"io/ioutil"
	"log"
	"testing"
)

var veryLongText = `
munmap(0x7f5734638000, 16400)           = 0
munmap(0x7f57342f5000, 280672)          = 0
munmap(0x7f5734632000, 20496)           = 0
munmap(0x7f573462c000, 20496)           = 0
munmap(0x7f5734627000, 16400)           = 0
munmap(0x7f5734622000, 16400)           = 0
munmap(0x7f573427f000, 480352)          = 0
munmap(0x7f573461d000, 16400)           = 0
munmap(0x7f5734618000, 16400)           = 0
munmap(0x7f5734613000, 16400)           = 0
munmap(0x7f573460e000, 16400)           = 0
munmap(0x7f5734590000, 16400)           = 0
munmap(0x7f5734586000, 36880)           = 0
munmap(0x7f5734581000, 16400)           = 0
munmap(0x7f5734209000, 480352)          = 0
munmap(0x7f573457c000, 16400)           = 0
munmap(0x7f5734203000, 20504)           = 0
munmap(0x7f57341fe000, 16400)           = 0
munmap(0x7f57341f2000, 45072)           = 0
munmap(0x7f57341e4000, 54592)           = 0
munmap(0x7f57341df000, 16400)           = 0
munmap(0x7f57341d8000, 24592)           = 0
munmap(0x7f57341d3000, 16400)           = 0
munmap(0x7f573418d000, 282688)          = 0
munmap(0x7f5734187000, 20496)           = 0
munmap(0x7f5734182000, 16400)           = 0
munmap(0x7f573417c000, 20496)           = 0
getpid()                                = 20037
`

var discard = log.New(ioutil.Discard, "", 0)

func Println(args ...interface{}) { discard.Println(args...) }

func BenchmarkLogToDevNull(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Println(veryLongText)
	}
	b.StopTimer()
}

func BenchmarkLogIfVerboseTrue(b *testing.B) {
	verbose := true
	for n := 0; n < b.N; n++ {
		if verbose {
			Println(veryLongText)
		}
	}
	b.StopTimer()
}
func BenchmarkLogIfVerboseFalse(b *testing.B) {
	verbose := false
	for n := 0; n < b.N; n++ {
		if verbose {
			Println(veryLongText)
		}
	}
	b.StopTimer()
}

func BenchmarkLogIfVerboseGreaterThan4(b *testing.B) {
	verbose := 5
	for n := 0; n < b.N; n++ {
		if verbose > 4 {
			Println(veryLongText)
		}
	}
	b.StopTimer()
}
func BenchmarkLogIfVerboseGreaterThan5(b *testing.B) {
	verbose := 5
	for n := 0; n < b.N; n++ {
		if verbose > 5 {
			Println(veryLongText)
		}
	}
	b.StopTimer()
}
