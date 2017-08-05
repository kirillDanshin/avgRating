package avgRating

import "testing"

var samples = []struct {
	rating      [5]int
	expectAvg   float64
	expectScore float64
}{
	{
		rating:      [5]int{0, 0, 0, 0, 0},
		expectAvg:   0,
		expectScore: 0,
	},
	{
		rating:      [5]int{1, 1, 1, 1, 1},
		expectAvg:   3.0,
		expectScore: 0.17,
	}, {
		rating:      [5]int{2, 2, 2, 2, 2},
		expectAvg:   3.0,
		expectScore: 0.24,
	}, {
		rating:      [5]int{3, 3, 3, 3, 3},
		expectAvg:   3.0,
		expectScore: 0.27,
	}, {
		rating:      [5]int{4, 4, 4, 4, 4},
		expectAvg:   3.0,
		expectScore: 0.3,
	}, {
		rating:      [5]int{5, 5, 5, 5, 5},
		expectAvg:   3.0,
		expectScore: 0.32,
	}, {
		rating:      [5]int{5, 4, 3, 2, 1},
		expectAvg:   2.3,
		expectScore: 0.15,
	}, {
		rating:      [5]int{5, 0, 0, 0, 5},
		expectAvg:   3.0,
		expectScore: 0.24,
	}, {
		rating:      [5]int{5, 0, 0, 4, 5},
		expectAvg:   3.3,
		expectScore: 0.33,
	}, {
		rating:      [5]int{5, 4, 0, 0, 5},
		expectAvg:   2.7,
		expectScore: 0.21,
	}, {
		rating:      [5]int{0, 0, 0, 0, 5},
		expectAvg:   5,
		expectScore: 0.57,
	}, {
		rating:      [5]int{0, 0, 0, 4, 5},
		expectAvg:   4.6,
		expectScore: 0.56,
	}, {
		rating:      [5]int{0, 0, 3, 4, 5},
		expectAvg:   4.2,
		expectScore: 0.51,
	}, {
		rating:      [5]int{0, 2, 3, 4, 5},
		expectAvg:   3.9,
		expectScore: 0.45,
	}, {
		rating:      [5]int{1, 2, 3, 4, 5},
		expectAvg:   3.7,
		expectScore: 0.42,
	}, {
		rating:      [5]int{9524, 4158, 10177, 25971, 68669},
		expectAvg:   4.2,
		expectScore: 0.79,
	}, {
		rating:      [5]int{134055, 57472, 143135, 365957, 1448459},
		expectAvg:   4.4,
		expectScore: 0.84,
	},
}

func TestAverage(t *testing.T) {
	for _, v := range samples {
		avg := Average(v.rating)
		if avg != v.expectAvg {
			t.Fatalf("given %v, got %v but expected %v", v.rating, avg, v.expectAvg)
		}
		t.Logf("given %v, got the expected %v", v.rating, avg)
	}
}

func TestScore(t *testing.T) {
	for _, v := range samples {
		score := Rate(v.rating)
		if score != v.expectScore {
			t.Fatalf("given %v, got %v but expected %v", v.rating, score, v.expectScore)
		}
		t.Logf("given %v, got the expected %v", v.rating, score)
	}
}

func BenchmarkAverage(b *testing.B) {
	avg := 0.0
	rating := [5]int{134055, 57472, 143135, 365957, 1448459}
	for i := 0; i < b.N; i++ {
		avg = Average(rating)
	}

	_ = avg
}

func BenchmarkRate(b *testing.B) {
	rate := 0.0
	rating := [5]int{134055, 57472, 143135, 365957, 1448459}
	for i := 0; i < b.N; i++ {
		rate = Rate(rating)
	}

	_ = rate
}

func BenchmarkScore(b *testing.B) {
	score := 0.0
	for i := 0; i < b.N; i++ {
		score = GetWilsonScore(345, 345345)
	}

	_ = score
}
