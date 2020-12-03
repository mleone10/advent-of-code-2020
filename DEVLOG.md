# AoC 2020 Developer Log

## 12/1
According to my coworker, Day 1's puzzle is actually a common-ish interview question.  The brute force solution is to just double or triple-nest for-loops.  An optimized solution to the 2-sum problem involves starting with indices 0 and `len(nums)`.  If the sum is greater than the target, decrement the upper index.  If the sum is less than the target, increment the lower index.

The 3-sum problem is a bit trickier, but ultimately just means subtracting `nums[len(nums)]` from the target, then running the 2-sum problem on `nums[:-1]`.  If the result isn't found, start over with `nums[:-1]`.

## 12/2
Most of today's was just parsing the input into a struct.  That said, I'm pretty happy with my `isValidIndex()` function implementation, as well as the signature of the `countValidPasswords()` function, which takes in a validation function.  That way I don't have to duplicate the loop.

Today I also learned how [heredocs](https://en.wikipedia.org/wiki/Here_document) work!  Specifically, I'm using one in the [newday](newday) script for my `main.go` template.

I also learned about the `-e` and `-u` Bash flags.  There's a brief writeup about them [here](https://web.archive.org/web/20110314180918/http://www.davidpashley.com/articles/writing-robust-shell-scripts.html), but basically:

* `-e` sets the **errexit** flag.  If any statement returns a non-zero return code, the script exists.
* `-u` sets the **nounset** flag.  If an uninitialized variable is encountered, the script exists.
    
Together, they ensure that a) a suffix-less directory is never created, and b) that days are never overwritten, since `mkdir` fails if invoked for a directory that already exists.