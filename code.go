package advent20201208

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

type Op struct {
	instruction string
	offset      int64
}

// RecordsFromFile returns a channel that gives records from a file
func RecordsFromFile(filename string) <-chan Op {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	iterator := make(chan Op)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	go func(scanner *bufio.Scanner) {

		for scanner.Scan() {
			line := scanner.Text()
			sides := strings.Fields(line)
			offset, err := strconv.ParseInt(sides[1], 10, 64)
			if err != nil {
				panic(err)
			}
			op := Op{sides[0], offset}
			iterator <- op
		}
		close(iterator)
	}(scanner)
	return iterator
}

func OpList(filename string) []Op {
	ops := make([]Op, 0)
	iterator := RecordsFromFile(filename)
	for op := range iterator {
		ops = append(ops, op)
	}
	return ops
}

func TerminatesAndAcc(filename string) (bool, int64) {
	opList := OpList(filename)
	seen := mapset.NewSet()
	var idx, acc int64

	for idx < int64(len(opList)) {
		if seen.Contains(idx) {
			return false, 0
		}
		seen.Add(idx)
		op := opList[idx]
		if op.instruction == "nop" {
			idx++
		} else if op.instruction == "acc" {
			acc += op.offset
			idx++
		} else if op.instruction == "jmp" {
			idx += op.offset
		}
	}
	return true, acc
}

func OpsTerminates(opList []Op) bool {
	seen := mapset.NewSet()
	idx := 0
	for idx < len(opList) {
		if seen.Contains(idx) {
			return false
		}
		seen.Add(idx)
		op := opList[idx]
		if op.instruction == "nop" {
			idx++
		} else if op.instruction == "acc" {
			idx++
		} else if op.instruction == "jmp" {
			idx += int(op.offset)
		}
	}
	return true
}

func Part1(filename string) (acc int64) {
	opList := OpList(filename)
	seen := mapset.NewSet()
	idx := int64(0)
	for !seen.Contains(idx) && int(idx) != len(opList) {
		seen.Add(idx)
		op := opList[idx]
		if op.instruction == "nop" {
			idx++
		} else if op.instruction == "acc" {
			acc += op.offset
			idx++
		} else if op.instruction == "jmp" {
			idx += op.offset
		}
	}
	return
}

func ReportTerminates(ops []Op, idx int, send chan<- int) {
	if OpsTerminates(ops) {
		send <- idx
	}
}

func Part2(filename string) int {
	opList := OpList(filename)
	send := make(chan int)
	for idx, op := range opList {
		tmp := make([]Op, len(opList))
		copy(tmp, opList)
		if op.instruction == "jmp" {
			tmp[idx].instruction = "nop"
			go func(i int) {
				go ReportTerminates(tmp, i, send)
			}(idx)
		} else if op.instruction == "nop" {
			tmp[idx].instruction = "jmp"
			go func(i int) {
				ReportTerminates(tmp, i, send)
			}(idx)
		}
	}
	for value := range send {
		return value
	}
	return 0
}
