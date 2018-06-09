package idealBoardGame

type boardGame struct {
	people []int
}

func newBoardGame(numberOfPeople int) boardGame {
	return boardGame{
		people: makePeople(numberOfPeople),
	}
}

func makePeople(numberOfPeople int) []int {
	people := []int{}
	for i := 1; i <= numberOfPeople; i++ {
		people = append(people, i)
	}
	return people
}

func (b *boardGame) startRound() int {
	peopleInNextRound := []int{}
	lastPeople := len(b.people) - 1
	for index, people := range b.people {
		if index%2 == 0 {
			if index == lastPeople {
				peopleInNextRound = peopleInNextRound[1:]
			}
			peopleInNextRound = append(peopleInNextRound, people)
		}
	}
	b.people = peopleInNextRound
	return len(b.people)
}
