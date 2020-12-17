package advent20201208

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

// Op is an instruction in this advent challenge
type Op struct {
	instruction string
	offset      int64
}

// Flip turns a nop into a jmp or vice versa
func (o *Op) Flip() {
	if o.instruction == "nop" {
		o.instruction = "jmp"
	} else if o.instruction == "jmp" {
		o.instruction = "nop"
	}
}

// RecordsFromFile returns a list of Ops for a file
func RecordsFromFile(filename string) []Op {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	ops := make([]Op, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Fields(line)
		offset, err := strconv.ParseInt(sides[1], 10, 64)
		if err != nil {
			panic(err)
		}
		op := Op{sides[0], offset}
		ops = append(ops, op)
	}
	return ops
}

// SumAcc sums up the accumulator for a file
func SumAcc(opList []Op) (finishes bool, acc int64) {
	seen := mapset.NewSet()
	idx := int64(0)
	for int(idx) < len(opList) {
		if seen.Contains(idx) {
			return false, acc
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

// Part1 answers part 1
func Part1(filename string) (acc int64) {
	opList := RecordsFromFile(filename)
	_, acc = SumAcc(opList)
	return acc
}

// ReportTerminates tells us what the accumulator was if this run succeeded
func ReportTerminates(ops []Op, send chan<- int64) {
	finishes, acc := SumAcc(ops)
	if finishes {
		send <- acc
	}
}

// SuccessfulAccumulator tries to run the opList once for each possible single instruction flip
// and returns the accumulator of the successful run
func SuccessfulAccumulator(opList []Op) int64 {
	send := make(chan int64)
	for idx := range opList {
		tmp := make([]Op, len(opList))
		copy(tmp, opList)
		tmp[idx].Flip()
		go func(ops []Op) {
			go ReportTerminates(tmp, send)
		}(tmp)
	}
	for value := range send {
		return value
	}
	return 0
}

// Part2 answers part 2
func Part2(filename string) int64 {
	opList := RecordsFromFile(filename)
	return SuccessfulAccumulator(opList)
}
