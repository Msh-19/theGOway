package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    var guess int
    const maxAttempts = 3 
    attempts := 0

    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I have selected a random number between 1 and 100.")
    fmt.Printf("You have %d attempts to guess it.\n", maxAttempts)

    for attempts < maxAttempts {
        attempts++
        fmt.Printf("Attempt %d\nEnter your guess: ", attempts)
        fmt.Scanln(&guess)

        if guess < target {
            fmt.Println("Too low!")
        } else if guess > target {
            fmt.Println("Too high!")
        } else {
            fmt.Printf("Correct! You've guessed the number in %d attempts.\n", attempts)
            return 
        }
    }

    fmt.Printf("You used up all of your %d attempts. The correct answer was %d.\n", maxAttempts, target)
}
