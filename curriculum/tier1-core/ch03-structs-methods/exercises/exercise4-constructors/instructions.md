# Exercise 4: Implement Constructor Functions

## Task

Create a `Player` struct for a game with a `NewPlayer` constructor that validates inputs and sets defaults. The constructor should reject empty names and negative health values, and default the health to 100 if zero is provided.

## Requirements

- Define a `Player` struct with `Name` (string), `Health` (int), and `Level` (int) fields
- Implement `NewPlayer(name string, health int)` that returns `(*Player, error)`
- Return an error if `name` is empty
- Return an error if `health` is negative
- Default `health` to 100 if the caller passes 0
- Set `Level` to 1 for all new players
- Add a `Status` method that returns a formatted string

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Constructor Functions Demo
==========================
Player: Hero (Health: 100, Level: 1)
Player: Tank (Health: 200, Level: 1)
NewPlayer("", 100): name cannot be empty
NewPlayer("Bad", -5): health cannot be negative
```

## Hints

<details>
<summary>Hint 1</summary>

Check for empty name with `if name == ""` and negative health with `if health < 0`. Return `nil, errors.New("...")` for each error case.

</details>

<details>
<summary>Hint 2</summary>

For the default health, check `if health == 0 { health = 100 }` before creating the struct. Always set `Level: 1` in the struct literal.

</details>
