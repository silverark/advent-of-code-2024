package part1

import (
	"log"
	"silverark/aoc-2024/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {

	value := process(file.GetData("../example.txt"))
	expect := 123
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetData("../input1.txt"))
	log.Println("The answer is", value)

}
