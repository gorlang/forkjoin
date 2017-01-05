package main

import (
	"fmt"
	"testing"
	"time"
)

/*
	Creates a test data array and checks the compute result briefly.
*/

func TestCompute(t *testing.T) {

	data := []string{"namenamename0namenamename1"}
	var test_data []string
	test_size := 1000000
	for i := 1; i <= test_size; i++ {
		test_data = append(test_data, data...)
	}

	collector := make(chan []string, 1)
	fork_size := 250000
	start := time.Now().UnixNano()

	fmt.Println("Computing...")

	Compute(test_data, collector, fork_size)
	collected := <-collector

	end := time.Now().UnixNano()
	fmt.Println("Started", GoCount, "go routines from", ForkCount, "forks.")
	fmt.Println("Processed", len(collected), "words.")
	fmt.Println("Took", (end-start)/1000000, "ms.")

	expected_len := len(test_data)
	got_len := len(collected)
	if expected_len != got_len {
		t.Error("Expected", expected_len, "but got", got_len)
	}

	expected_result := "1emanemaneman0emanemaneman"
	got_result := collected[0]
	if expected_result != collected[0] {
		t.Error("Expected", expected_result, "but got", got_result)
	}

}
