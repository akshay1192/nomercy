Complexity

    Worst Case : O(n^2) time and O(1) space
    Avg Case   : O(n^2) time and O(1) space
    Best Case  : O(n) time and O(1) space


    Note : Best case is O(n) because of the swappedOnce flag. In case array is already sorted, we are break the loop there itself
           If we dont add swappedOnce logic then our best case will also become O(n^2)


Run test cases

    run : cd insertionSort/helperFunc dir
    run : go test


Run benchmarking

    run : cd insertionSort dir
    run : go test -bench=.
    
    
    // This -run flag takes in a regex pattern that can filter out the benchmarks we want to run.
    // And this will trigger both our TestTargetSumCase1,TestTargetSumCase2 and BenchmarkTargetSum functions:
    run : go test -run=TargetSum -bench=.
    
    
    // If we wanted to only run our BenchmarkTargetSum function then we could change our regex pattern to be:
    run : go test -run=Bench -bench=.


Run main program

    // execution time of the program is also being shown in the output
    
    run : cd to root dir
    run : go run main.go

    
    
    