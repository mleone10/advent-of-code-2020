# AoC 2020 Developer Log

This doc is a raw journal of my progress toward AoC 2020. As I refine those thoughts, I'm posting them to [my blog](https://marioleone.me/tags/advent-of-code/) and deleting them from here.

## 12/8

- Classic! Implement a virtual machine!
- Part one, very easy if you know how computers work
- Part two, changing the program?? A twist!

## 12/9

TwoSum is back! Thanks to the small function I wrote as an example in my blog post on [Day 1](https://marioleone.me/posts/2020/aoc-2020-p1/#day-1), part one was just iterating through the list of ints until I found a target that didn't have a valid two-sum. I was initially bitten by a lack of understanding about how slices get passed into functions. As I iterated through the main slice and sorted each sub-slice, the main slice was being modified. This is because of the way slices are passed by value in Go - changes to the data in a slice persist outside of function calls, but you can't add or remove elements.

There's likely an elegant way to approach part 2, but I brute-forced it by using nested for-loops to check all sub-slices until a match was found. One improvement I saw was to save the result of sum(ints[i:j+1]) and use it when calculating `sum(ints[i:j+1])`. Another would be to exit the inner for-loop when the `sum()` exceeded the `target`.
