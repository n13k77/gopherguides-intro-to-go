# Assignment Week 02

This document describes my solution for the assignment of week 02. The code can be found in the file `assignment02_test.go` where this document contains a description of:

- the design choices that I made
- difficulties that I ran in to
- surprises that I found

## Design Choices

First of all, I started out by creating 3 empty test functions, split according to the different collection types that I had to write tests for. After this, I copied and pasted the requirements from the assignments into each of the functions where they had to be implemented.

The requirements are still present in the code. I left them there because it is easily visible which part of the requirements is implemented in which part of the code. If this is not desirable then do let me know in this PR and I will take them out.

After this part was done, it was fairly easy to implement the required lines of code per requirement. In two cases there would have been a solution possible that in my opinion would have been more straightforward from a coding perspective:

- lines 12-14 could have been replaced with:

```go
act = exp
```

- lines 29-31 could have been replaced with:

```go
act = append(act, exp...)
```

In both cases, I implemented the solution with a loop construct. This choice was made in order to follow the requirements from the assignments as closely as possible. However, from a clean code perspective, I would actually prefer the constructs mentioned above.

The final design choice that I made is that I put the written assignment in a `README.md` (this file) instead of in the description of the PR. This has, in my opinion, the benefit that it is also under source countrol.

## Difficulties

I did not run into actual difficulties while coding the solution. I did find it a bit hard to make the choices as mentioned in the section "Design Choices". Follow the requirements or go for a cleaner solution from a coding perspective. I decided in the end to consistently follow the requirements from the assignment.

To make sure that the tests actually fail when the conditions are not met, I added some one-liners that would make them fail, ran the tests and took the oneliners out again. That worked properly.

However, I prefer the construct where a particular test is used to test matching and non-matching situations in one go. This avoids the need to smuggle one-liners in and out. I believe this is called Table Driven Testing in Go, but I could be wrong.

## Surprises

There were two surprises for me while doing this assignment:

1) I had my first experience with workflows in Github. I was impressed by it, particularly the matrix functionality.

2) I knew Go testing, but then the specific construct where a module has a corresponding file with tests. Now only the file with tests is present, not the module file that contains the module subjected to tests. I was surprised that this also works.

## Closing Remark

I am fully aware that I have not reached the required amount of words. I have nothing more to add, so I decided to leave it at this. Adding more words just to reach the count feels counterintuitive to me.
