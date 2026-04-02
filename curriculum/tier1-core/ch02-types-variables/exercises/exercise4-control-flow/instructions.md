# Exercise 4: Implement Control Flow Patterns

## Task

Write a FizzBuzz program using Go's control flow constructs. For numbers 1 through 20, print "Fizz" for multiples of 3, "Buzz" for multiples of 5, "FizzBuzz" for multiples of both, and the number itself otherwise. Use a conditionless `switch` statement for the logic.

## Requirements

- Loop from 1 to 20 using a `for` loop
- Use a conditionless `switch` (no expression after `switch`) to determine the output
- Print "FizzBuzz" for multiples of both 3 and 5
- Print "Fizz" for multiples of 3 only
- Print "Buzz" for multiples of 5 only
- Print the number for all other values
- After the loop, print a summary of how many Fizz, Buzz, and FizzBuzz results there were

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
FizzBuzz (1–20)
===============
  1
  2
  Fizz
  4
  Buzz
  Fizz
  7
  8
  Fizz
  Buzz
  11
  Fizz
  13
  14
  FizzBuzz
  16
  17
  Fizz
  19
  Buzz

Summary: 5 Fizz, 3 Buzz, 1 FizzBuzz
```

## Hints

<details>
<summary>Hint 1</summary>

A conditionless switch looks like: `switch { case i%15 == 0: ... case i%3 == 0: ... }`. Check the most specific case (divisible by both) first.

</details>

<details>
<summary>Hint 2</summary>

Keep three counters (`fizz`, `buzz`, `fizzBuzz`) and increment the appropriate one in each case.

</details>
