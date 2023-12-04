## Advent of Code 2023 - Day 1

This is [Day 1](https://adventofcode.com/2023/day/1) of Advent of Code 2023, written in Go.

It is able to find the correct
answers.

This was a bit challenging, as
Go's built-in [regexp](https://pkg.go.dev/regexp) package does not have support for regex expressions with lookback and lookforward. This missing functionality meant that it had trouble initially with overlapping matches (e.g., "oneight" for "one" and "eight").

While I did spend some time looking into alternative expression patterns that would create a similar effect, I eventually decided to use two regex expressions, with the second being the numbers spelled in reverse. I would then reverse the entries and then use the regex matcher.

