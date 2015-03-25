package expr

import (
	"bosun.org/opentsdb"
	"fmt"
	"math/rand"
	"testing"
)

func TestSlowUnion(t *testing.T) {
	ra, rb := buildFakeResults()
	e := State{}
	e.unjoinedOk = true
	x := e.union(ra, rb, "")
	if len(x) != 1000 {
		t.Errorf("Bad length %d != 1000", len(x))
	}
}

func buildFakeResults() (ra, rb *Results) {
	ra = &Results{}
	rb = &Results{}
	rand.Seed(42)
	for i := 0; i < 50000; i++ {
		tags := opentsdb.TagSet{}
		tags["disk"] = fmt.Sprint("a%d", i)
		tags["host"] = fmt.Sprint("b%d", i)
		if i < 1000 {
			ra.Results = append(ra.Results, &Result{Value: Number(rand.Int63()), Group: tags})
		}
		rb.Results = append(ra.Results, &Result{Value: Number(rand.Int63()), Group: tags})
	}
	rb.Results = rb.Results.DescByValue()
	return ra, rb
}

func BenchmarkSlowUnion(b *testing.B) {
	ra, rb := buildFakeResults()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := State{}
		e.unjoinedOk = true
		e.union(ra, rb, "")
	}
}
