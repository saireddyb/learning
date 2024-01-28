Concurrency:
Imagine you're cooking in the kitchen. Concurrency is like chopping vegetables while your soup is simmering. You switch between tasks (chopping and stirring) to make progress on both, even though you're not doing them at the exact same moment.

Parallelism:
Now, picture having two chefs working side by side in the kitchen. Both are chopping, stirring, and doing different tasks simultaneously. That's parallelism—doing multiple things at the same time.


go routines
- a lightweight thread that has a separate independent execution.
- can execute concurrently with other go-routines
- Managed entirely by Go runtime

To execute a method or function in go routine 
call it with go keyword
go calculate()


Anonymous go routime

go func() {

}

context switching
functions calls
garbage collections
network calls
channel operations
on using go keyword





Go-routine vs Threads
go routine are cheaper
go routine are multiplexed to a fewer number of OS threads
context-switching time of go-routines in much faster
go routines communicates using channels

