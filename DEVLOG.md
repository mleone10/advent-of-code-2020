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

## 12/5
This was a nice, easy puzzle after yesterday's slog.  While reading the puzzle, I immediately recognized that the boarding pass sequence was describing a binary search.  However, it's kind of coming at it from the opposite direction as a usual binary search - rather than ask for an algorithm to determine whether to use the top or bottom half of an array, today's problem *gives* us those movements.  I have a hunch that Go's built in `sort.Search()` method can be used to simplify my approach, so I might give that a shot at some point.

Instead, I sort of brute forced it by initializing a slice of `N` integers and looping through the given instructions.  Depending on the instruction, I used slice expressions to chop off the top or bottom half of the array.  The puzzle guaranteed that we would have exactly enough instructions, so I knew that after this chopping loop I would have a single element left - the row/column ID.

For part two, I got to use `sort.Ints()`!  First I created a slice of seat IDs, then sorted it.  Looping through to find the missing number just meant finding the first `id` that was not `previousID + 1`.

One recurring lesson that these problems teach is the decomposition of a large problem into smaller, easier problems.  Today's is a great example of that.  For part one, I have to find the largest seat ID, so I made a `calcMaxSeatID()` function.  How do I get the max seat ID?  I need to calculate all of the individual seat IDs and find the max.  How do I calculate an individual seat ID?  Easy, `row * 8 + col`.  How do I find `row` and `col`?  Functions for `calcRow()` and `calcCol()`.  Here we finally get to the meat of the puzzle - given a string, how do I binary search through a slice of integers?  That's the real challenge, and everything before that is just getting there.

This decomposition has the added benefit of making part two easier, usually.  Finding my seat ID first required that I have a list of *all* seat IDs, which I already had a function for!

## 12/6
Today's puzzle was surprisingly straightforward.  Again we encounter the idea of splitting groups by a double `\n` character.  I think I'll try to find a more idiomatic way to do that than just detecting empty lines.  The `bufio` package has the concept of a `SplitFunc`, which I might be able to use to scan entire groups in at a time.  Then I can split those multi-line strings by `\n` to process the given group.  Plus, that custom `SplitFunc` can be extracted and reused.

The solution for today's puzzle was really just about the clever usage of `map`s.  In my initial part one solution, I stored "questions" in a `map[string]bool`, which is effectively a set.  Part two required that I change the pseudo-set to also store a count of each "question", so I ended up modifying the part one solution to match.

My solution for part two was delayed by not fully considering my algorithm before writing it.  The idea that I would have to reuse my part one solution, but with a count, was fairly immediate.  Subsequently *using* that count wasn't entirely clear, so my first attempt at part two involved something like `if len(rs) == len(g)`.  Completely off the mark, but I was in guess-and-check mode.  If I would have taken the time to think it through, I might have finished a minute before I did.

In any case, part two's solution started with constructing the `map[string]int` of responses within a group.  With that, I just had to count the number of "questions" whose count matched the number of people in the group.
