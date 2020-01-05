package main

import (
	"testing"
)

func TestBonusAmount(t *testing.T) {
	tests := []struct {
		name    string
		score   []int
		want    int
	}{
		{ name: "More than 10_000", score: []int{15_000, 12_000, 11_000}, want: 400},
		{ name: "Less than 10_000", score: []int{8000, 9000, 5, 4, 400}, want:0},
		{ name: "Equally 10_000", score: []int{10_000, 10_000, 10_000}, want: 0},
		{ name: "Mixed", score: []int{10_000, 15_000, 12_000, 8000}, want:350},
	}

	for _, tests := range tests{
		got := BonusAmount(5, tests.score)
		if got != tests.want{
			t.Error("BonusAmount test", tests.name, "got", got, "want", tests.want)
		}
	}
}