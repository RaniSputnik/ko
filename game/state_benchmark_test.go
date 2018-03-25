package game_test

import (
	"testing"

	"github.com/RaniSputnik/ko/game"
)

// BenchmarkPlayEverywhere5-8   	   50000	     38695 ns/op	    3408 B/op	      81 allocs/op
// BenchmarkPlayEverywhere5-8   	   10000	    118797 ns/op	    9168 B/op	     801 allocs/op

func BenchmarkPlayEverywhere5(b *testing.B) {
	benchmarkPlayEverywhere(b, 5)
}

// BenchmarkPlayEverywhere9-8   	    3000	    397907 ns/op	   17040 B/op	     251 allocs/op
func BenchmarkPlayEverywhere9(b *testing.B) {
	benchmarkPlayEverywhere(b, 9)
}

// BenchmarkPlayEverywhere13-8   	    1000	   1754967 ns/op	   48736 B/op	     516 allocs/op
func BenchmarkPlayEverywhere13(b *testing.B) {
	benchmarkPlayEverywhere(b, 13)
}

// BenchmarkPlayEverywhere19-8   	     200	   8446118 ns/op	  178096 B/op	    1093 allocs/op
// BenchmarkPlayEverywhere19-8   	      50	  26381842 ns/op	 1655572 B/op	  185773 allocs/op
func BenchmarkPlayEverywhere19(b *testing.B) {
	benchmarkPlayEverywhere(b, 19)
}

func benchmarkPlayEverywhere(b *testing.B, boardSize int) {
	for n := 0; n < b.N; n++ {
		m := game.Match{
			Owner:    &Alice,
			Opponent: &Bob,
			Board: game.Board{
				Size: boardSize,
			},
		}

		var err error
		for y := 0; y < boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				m, err = m.Play(m.Next(), x, y)
				if err != nil {
					b.Fatalf("Failed to play: %s", err)
				}
			}
		}
	}
}
