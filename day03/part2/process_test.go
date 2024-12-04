package part1

import (
	"log"
	"silverark/aoc-2024/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {

	value := process(file.GetFile("../example2.txt"))
	expect := 48
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetFile("../input1.txt"))
	log.Println("The answer is", value)

	if value != 89349241 {
		t.Fatalf("Received %v, but expected %v", value, 89349241)
	}
}
