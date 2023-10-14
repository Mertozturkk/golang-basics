package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

/*
$ go test -bench=.
benchmark test yazmaya cok benzer ve dilin birinci sinif ozelliklerinden birisidir
https://pkg.go.dev/testing#hdr-Benchmarks
*/
