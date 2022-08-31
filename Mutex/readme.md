**Mutex**<br><br>
When a program runs concurrently,
the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time.
This section of code that modifies shared resources is called critical section
race condition:<br>
where the output of the program depends on the sequence of execution of Goroutines is called race condition.

- Mutex is available in the sync package
- There are two methods defined on Mutex namely Lock and Unlock.
- If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.
- we can solve race condition with channel or mutex(what to use when? it depends on problem but In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code.

)


