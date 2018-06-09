package idealBoardGame

import (
	"reflect"
	"testing"
)

func Test_IdealBoardGame_Round_1(t *testing.T) {
	expectedPeopleLeft := 50
	expectedPeople := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99}
	game := newBoardGame(100)
	actualPeopleLeft := game.startRound()
	actualPeople := game.people

	if expectedPeopleLeft != actualPeopleLeft {
		t.Errorf("Expected : %d but it is %d", expectedPeopleLeft, actualPeopleLeft)
	}
	if reflect.DeepEqual(expectedPeople, actualPeople) {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}
