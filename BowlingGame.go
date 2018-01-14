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
	var frameStartIndex int

	frameStartIndex = 0
	for frameCounter := 0; frameCounter < 10; frameCounter++ {
		if g.isStrike(frameStartIndex) {
			retVal += (g.rolls[frameStartIndex] + g.rolls[frameStartIndex+1] + g.rolls[frameStartIndex+2])
			frameStartIndex++
		} else if g.isSpare(frameStartIndex) {
			retVal += (g.rolls[frameStartIndex] + g.rolls[frameStartIndex+1] + g.rolls[frameStartIndex+2])
			frameStartIndex += 2
		} else {
			retVal += (g.rolls[frameStartIndex] + g.rolls[frameStartIndex+1])
			frameStartIndex += 2
		}

	}

	return retVal
}

func (g game) isStrike(frameStartIndex int) bool {
	var retVal bool

	if g.rolls[frameStartIndex] == 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}

func (g game) isSpare(frameStartIndex int) bool {
	var retVal bool

	if g.rolls[frameStartIndex]+g.rolls[frameStartIndex+1] == 10 && g.rolls[frameStartIndex] != 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}
