package main

/*
	A simple example of the ForkJoin pattern using goroutines and channels for collection of the result.
	ForkJoin is a algorithm that splits the processing of data into subsequential smaller pieces recursively.
	If size of the indata array is larger than fork_size, computing of data will be split (forked) into two parts
	Thus there could be 1, 2, 4, 8, 16, 32 or more goroutines processing the data depending on initial size of data and
	selected fork_size.
*/

/*
	Compute method handles the forking of data and call to the actual compute go routines.
*/

func Compute(indata []string, result chan<- []string, fork_size int) {

	size_indata := len(indata)

	if size_indata <= fork_size {
		go ComputeDirect(indata, result)
		return
	}

	result_fork_a := make(chan []string)
	result_fork_b := make(chan []string)

	split_at := size_indata / 2

	go Compute(indata[:split_at], result_fork_a, fork_size)
	go Compute(indata[split_at:size_indata], result_fork_b, fork_size)

	var subres []string
	result <- append(append(subres, <-result_fork_a...), <-result_fork_b...)
}

/*
	Implementation of the actual computing logic.
	Could be reduce, sum or whatever you like.
	In this example we are just reversing the text in an array of strings. ("Golang" will be "gnaloG")
*/

func ComputeDirect(indata []string, result chan<- []string) {

	var res []string
	for i := 0; i < len(indata); i++ {
		tmp := ""
		for j := (len(indata[i]) - 1); j > -1; j-- {
			tmp += string(indata[i][j])
		}
		res = append(res, tmp)
	}
	result <- res
}
