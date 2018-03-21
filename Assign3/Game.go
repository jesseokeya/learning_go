package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game struct holds game private variables
type Game struct {
	potAmount int
	betAmount int
	betType   string
}

func init() {
	fmt.Println("Welcome to Double or Nothing Dice Game..bet an amount and type ")
	fmt.Println("     -if you are correct, you win twice your bet,")
	fmt.Println("     -otherwise you lose your bet")
	fmt.Println("A bet of 0 ends the game")
}

func (g *Game) getBetFromUser() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("\nYour current pot is " + strconv.Itoa(g.potAmount))
	if g.potAmount == 0 {
		return
	}
	fmt.Print("Enter your bet amount: ")
	betAmount, _ := scanner.ReadString('\n')
	convBetAmount, err := strconv.Atoi(strings.TrimSuffix(betAmount, "\n"))
	if err != nil {
		panic(err)
	}
	g.betAmount = convBetAmount
	for g.betAmount < 0 || g.betAmount > g.potAmount {
		fmt.Printf("Error - cannot bet less than 0 or more than %s...Enter your bet amount: ", strconv.Itoa(g.potAmount))
		betAmount, _ := scanner.ReadString('\n')
		convBetAmount, err := strconv.Atoi(strings.TrimSuffix(betAmount, "\n"))
		if err != nil {
			panic(err)
		}
		g.betAmount = convBetAmount
	}
	if g.betAmount == 0 {
		return
	}
	fmt.Print("Enter your bet type: ")
	betType, _ := scanner.ReadString('\n')
	convBetType := strings.TrimSuffix(betType, "\n")
	convBetType = strings.ToLower(convBetType)
	g.betType = convBetType
	for !(g.betType == "low" || g.betType == "high" || g.betType == "odd" || g.betType == "even") {
		fmt.Print("Please enter odd, even, high, or low ....Enter your bet type: ")
		betType, _ := scanner.ReadString('\n')
		convBetType := strings.TrimSuffix(betType, "\n")
		convBetType = strings.ToLower(convBetType)
		g.betType = convBetType
	}
}

func (g *Game) playGame() {
	for g.betAmount > 0 || g.potAmount > 0 {
		g.getBetFromUser()
		if g.betAmount <= 0 || g.potAmount <= 0 {
			fmt.Println("You end the game with pot of " + strconv.Itoa(g.potAmount))
			return
		}
		const defaultValue int = 0
		sum := 0
		dieArray := []Die{
			Die{defaultValue},
			Die{defaultValue},
			Die{defaultValue},
		}
		fmt.Print("Your dies are: ")
		for i, die := range dieArray {
			die.rollDie()
			die.display()
			if i != len(dieArray)-1 {
				fmt.Print(" and ")
			} else {
				fmt.Println()
			}
			sum += die.getValue()
		}
		switch g.betType {
		case "odd":
			if sum%2 != 0 {
				fmt.Println("You WIN....double your bet ")
				g.potAmount += g.betAmount
			} else {
				fmt.Println("You LOSE....your bet ")
				g.potAmount -= g.betAmount
			}
			break
		case "even":
			if sum%2 == 0 {
				fmt.Println("You WIN....double your bet ")
				g.potAmount += g.betAmount
			} else {
				fmt.Println("You LOSE....your bet ")
				g.potAmount -= g.betAmount
			}
			break
		case "low":
			if sum < 9 {
				fmt.Println("You WIN....double your bet ")
				g.potAmount += g.betAmount
			} else {
				fmt.Println("You LOSE....your bet ")
				g.potAmount -= g.betAmount
			}
			break
		case "high":
			if sum >= 9 {
				fmt.Println("You WIN....double your bet ")
				g.potAmount += g.betAmount
			} else {
				fmt.Println("You LOSE....your bet ")
				g.potAmount -= g.betAmount
			}
		default:
			break
		}
	}
}
