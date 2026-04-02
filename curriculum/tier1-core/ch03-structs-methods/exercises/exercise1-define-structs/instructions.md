# Exercise 1: Define and Use Structs

## Task

Create a program that defines a `Movie` struct with fields for title, director, year, and rating. Create several movie instances, print their details, and demonstrate that structs are copied on assignment.

## Requirements

- Define a `Movie` struct with fields: `Title` (string), `Director` (string), `Year` (int), `Rating` (float64)
- Create at least two movie instances using named field initialization
- Print each movie's details in a formatted way
- Demonstrate copy behavior: copy a movie, change the copy's rating, and show the original is unchanged

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Movie Collection
================
"The Matrix" (1999) directed by The Wachowskis — Rating: 8.7
"Inception" (2010) directed by Christopher Nolan — Rating: 8.8

Copy behavior:
  Original rating: 8.7
  Copy rating:     9.0
  (Original unchanged — structs are value types)
```

## Hints

<details>
<summary>Hint 1</summary>

Define the struct with `type Movie struct { ... }` and create instances with `Movie{Title: "...", Director: "...", Year: 1999, Rating: 8.7}`.

</details>

<details>
<summary>Hint 2</summary>

To demonstrate copy behavior, assign one movie to a new variable (`copy := original`), change a field on the copy, then print both to show they're independent.

</details>
