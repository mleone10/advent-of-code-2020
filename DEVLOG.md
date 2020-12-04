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

## 12/3
Nothing really new today!  A fun modulo operator to progressively loop through the static grid, an in-line boolean expression in an `append` statement, and a rarely-encountered `*=` operator were about all it took.  This is the first grid of the year though, so I'll probably go back and start extracting a Grid type into the library.

## 12/4
Today's puzzle was just tedious, but there were a few important lessons learned.  Reading in the input was interesting.  I read line-by-line until an empty line was encountered, but I've seen a few other solutions that used the double newline ("\n\n") as a split delimiter.  One coworker's Awk solution actually started by condensing the whole input into one semicolon-delimited line.  So for future puzzles, I'll be thinking about transformations I can do to aid reading the input.

This problem also gave me a good opportunity to refactor the original solution to make it cleaner.  The first pass (which actually got me the stars) stored passports as `map[string]string`s.  The problem with that was just the repetitive, messy-looking `p["cid"]` access pattern.  On the second pass, I did two main things:
1. Converted passports to a struct with dedicated field strings
2. Broke each field's validation into their own functions

If I were to continue cleaning it up, I might try to refactor further so that my `validateData` method was essentially:
```go
func validateData(p passport) bool {
    for f := range p.fields {
        if !f.isValid() {
            return false
        }
    }

    return true
}
```

Such a loop would look cleaner, but would mean implementing a `type field interface` for each field.  Not the worst thing, but it would be more code for not much gain.  If the passport concept comes around again (as many of my friends believe), I may end up doing just that.

Other than that, there were only three other observations.  First, since Go doesn't have a `set` type, I just used a `map[string]bool` global variable to store the valid eye colors.  This made validation easy - just try to access `ecls[p.ecl]`.  If `p.ecl` is in the map, it returns `true`.  I'm sure the introduction of generics in 2022 will lead to a standard way of doing sets, though.

Second, an error in my input loop delayed my first star by a few minutes.  I *think* the `bufio.Scanner.Scan()` method was returning `false` after reading the last `\n`, which meant the last line never triggered the creation of the final passport.

Finally, I validated hair color by trying to convert the given hex value into an integer.  This can be done with `strconv.ParseUint`, which apparently will work for any base from 2 to 36.  If conversion fails, `err != nil`.  I've also seen solutions that used regex matches for this (and all) validation, but I like to think this is more "correct".
