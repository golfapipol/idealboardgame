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
	if reflect.DeepEqual(expectedPeople, actualPeople) != true {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}

func Test_IdealBoardGame_Round_2(t *testing.T) {
	expectedPeopleLeft := 25
	expectedPeople := []int{1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97}
	game := newBoardGame(100)
	game.startRound()
	actualPeopleLeft := game.startRound()
	actualPeople := game.people

	if expectedPeopleLeft != actualPeopleLeft {
		t.Errorf("Expected : %d but it is %d", expectedPeopleLeft, actualPeopleLeft)
	}
	if reflect.DeepEqual(expectedPeople, actualPeople) != true {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}

func Test_IdealBoardGame_Round_3(t *testing.T) {
	expectedPeopleLeft := 12
	expectedPeople := []int{9, 17, 25, 33, 41, 49, 57, 65, 73, 81, 89, 97}
	game := newBoardGame(100)
	game.startRound()
	game.startRound()
	actualPeopleLeft := game.startRound()
	actualPeople := game.people

	if expectedPeopleLeft != actualPeopleLeft {
		t.Errorf("Expected : %d but it is %d", expectedPeopleLeft, actualPeopleLeft)
	}
	if reflect.DeepEqual(expectedPeople, actualPeople) != true {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}

func Test_IdealBoardGame_Round_4(t *testing.T) {
	expectedPeopleLeft := 6
	expectedPeople := []int{9, 25, 41, 57, 73, 89}
	game := newBoardGame(100)
	game.startRound()
	game.startRound()
	game.startRound()
	actualPeopleLeft := game.startRound()
	actualPeople := game.people

	if expectedPeopleLeft != actualPeopleLeft {
		t.Errorf("Expected : %d but it is %d", expectedPeopleLeft, actualPeopleLeft)
	}
	if reflect.DeepEqual(expectedPeople, actualPeople) != true {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}

func Test_IdealBoardGame_Round_6(t *testing.T) {
	expectedPeopleLeft := 1
	expectedPeople := []int{73}
	game := newBoardGame(100)
	game.startRound()
	game.startRound()
	game.startRound()
	game.startRound()
	game.startRound()
	actualPeopleLeft := game.startRound()
	actualPeople := game.people

	if expectedPeopleLeft != actualPeopleLeft {
		t.Errorf("Expected : %d but it is %d", expectedPeopleLeft, actualPeopleLeft)
	}
	if reflect.DeepEqual(expectedPeople, actualPeople) != true {
		t.Errorf("Expected : %v but it is %v", expectedPeople, actualPeople)
	}
}
