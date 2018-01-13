package main

type game struct {
	rolls      [21]int
	rollNumber int
}

func main() {

}

func (g *game) roll(pinsDown int) {
	g.rolls[g.rollNumber] = pinsDown
	g.rollNumber++
}

func (g game) score() int {
	var retVal int
	var rollPointer int

	rollPointer = 0
	for frameCounter := 0; frameCounter < 10; frameCounter++ {
		if isStrike(g.rolls[rollPointer]) {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1] + g.rolls[rollPointer+2])
			rollPointer++
		} else if isSpare(g.rolls[rollPointer], g.rolls[rollPointer+1]) {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1] + g.rolls[rollPointer+2])
			rollPointer += 2
		} else {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1])
			rollPointer += 2
		}

	}

	return retVal
}

func isStrike(ballOne int) bool {
	var retVal bool

	if ballOne == 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}

func isSpare(ballOne int, ballTwo int) bool {
	var retVal bool

	if ballOne+ballTwo == 10 && ballOne != 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}
