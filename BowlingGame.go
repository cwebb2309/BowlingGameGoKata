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
		if g.isStrike(rollPointer) {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1] + g.rolls[rollPointer+2])
			rollPointer++
		} else if g.isSpare(rollPointer) {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1] + g.rolls[rollPointer+2])
			rollPointer += 2
		} else {
			retVal += (g.rolls[rollPointer] + g.rolls[rollPointer+1])
			rollPointer += 2
		}

	}

	return retVal
}

func (g game) isStrike(rollPointer int) bool {
	var retVal bool

	if g.rolls[rollPointer] == 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}

func (g game) isSpare(rollPointer int) bool {
	var retVal bool

	if g.rolls[rollPointer]+g.rolls[rollPointer+1] == 10 && g.rolls[rollPointer] != 10 {
		retVal = true
	} else {
		retVal = false
	}
	return retVal
}
