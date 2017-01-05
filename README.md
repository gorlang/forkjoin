# forkjoin

A simple example of the ForkJoin pattern using goroutines and channels for collection of the result.
ForkJoin is a algorithm that splits the processing of data into subsequential smaller pieces recursively.
If size of the indata array is larger than fork_size, computing of data will be split (forked) into two parts
Thus there could be 1, 2, 4, 8, 16, 32 or more goroutines processing the data depending on initial size of data and
selected fork_size.

The End.