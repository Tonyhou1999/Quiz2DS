# Quiz2DS
This is the Test 2
- There are two files, suggested as the name, for the two seperate parts
- Running is very simple, just by run by go run Quiz2QA.go number, where A = 1 or 2, number is how many numbers will be randomly generated
- The outcome will be the following, for Question 1


The sum of the array is 113089414
It takes 517400 nanoseconds to compute the summation directly<br />
The sum of the array is 113085944                            \\
It takes 0 nanoseconds to compute the summation directly  
(Note, that the sum is different I think mostly due to stack overflow, which is unfortunate, but you should try a number big enough, I tried with 15000, but as you can see, the second one takes about 0 nanoseconds, but it was a nonzero when I input like 15000000, so keep that in mind if you want to play with the input parameter)

- The outcome for Q2 will be the following

The Sorting used to sort the slice takes 18144400 nanoseconds to sort
The SliceStable used to sort the slice takes 57100100 nanoseconds to sort

(Note: I tried with the input 15000, but you are certainly welcome to try larger ones by modifying the parseInt method)

Analysis on Notation:
Yes I do agree with the Go official documental time complexity in sorting, as you can see they are differed by a few times, but that is consistant with the notation as  I blieve that O(anlog(n)) = O(nlog(n)), where a is a constant number.
