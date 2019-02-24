Complexity

    O(n) time and O(1) space with bubble sort (actually it will be 3 * O(n)) ----- After 1 bubble sort iteration, you get the largest pushed to last
    O(n) time and O(1) space with maintaining array of 3 and swapping accordingly


Run test cases

    run : cd threeLargestNumber/helperFunc dir
    run : go test


Run benchmarking

    run : cd threeLargestNumber dir
    run : go test -bench=.
    
    
    // This -run flag takes in a regex pattern that can filter out the benchmarks we want to run.
    // And this will trigger both our TestTargetSumCase1,TestTargetSumCase2 and BenchmarkTargetSum functions:
    run : go test -run=threeLargestNumber -bench=.
    
    
    // If we wanted to only run our BenchmarkTargetSum function then we could change our regex pattern to be:
    run : go test -run=Bench -bench=.


Run main program

    // execution time of the program is also being shown in the output
    
    run : cd to root dir
    run : go run main.go

    
    
    