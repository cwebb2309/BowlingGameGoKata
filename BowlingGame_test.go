package main

import "testing"

func TestGutterGame(t *testing.T) {

	var testGame game
	scores := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 0 {
		t.Errorf("Score expected 0. Actual %d", testGame.score())
	}
}

func TestGameEveryRollIsOne(t *testing.T) {

	var testGame game
	scores := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 20 {
		t.Errorf("Score expected 20. Actual %d", testGame.score())
	}
}

func TestGameRollASpareFollowedByThree(t *testing.T) {

	var testGame game
	scores := []int{3, 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 16 {
		t.Errorf("Score expected 16. Actual %d", testGame.score())
	}
}

func TestGameEveryRollIsFive(t *testing.T) {

	var testGame game
	scores := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 150 {
		t.Errorf("Score expected 150. Actual %d", testGame.score())
	}
}

func TestGameRollAStrikeFollowedByThreeThenFour(t *testing.T) {

	var testGame game
	scores := []int{10, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 24 {
		t.Errorf("Score expected 24. Actual %d", testGame.score())
	}
}

func TestGameRollAllStrikes(t *testing.T) {

	var testGame game
	scores := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

	for _, score := range scores {
		testGame.roll(score)
	}

	if testGame.score() != 300 {
		t.Errorf("Score expected 300. Actual %d", testGame.score())
	}
}
