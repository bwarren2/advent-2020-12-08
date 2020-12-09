package advent20201208_test

import (
	"testing"

	advent "github.com/bwarren2/advent20201208"
)

func TestRecordsFromFile(t *testing.T) {
	// for value := range advent.RecordsFromFile("sample.txt") {
	// fmt.Println(value)
	// }
}
func TestOpList(t *testing.T) {
	// fmt.Println(advent.OpList("sample.txt"))
}

func TestPart1(t *testing.T) {
	value := advent.Part1("input.txt")
	if value != 2014 {
		t.Errorf("Got the wrong value! %v", value)
	}
}
func TestPart2(t *testing.T) {
	value := advent.Part2("input.txt")
	if value != 2251 {
		t.Errorf("Got the wrong value! %v", value)
	}
}
