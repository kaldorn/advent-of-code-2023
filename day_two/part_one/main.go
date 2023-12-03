package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Roll struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID    int
	Rolls []Roll
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	// Define max allowed cubes per roll
	maxAllowed := Roll{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	// Loop through the games parsing them and checking if they're compatible
	compatibleGames := []Game{}
	for fileScanner.Scan() {
		game := formatPlayedGames(fileScanner.Text())
		if checkCompatibleGame(maxAllowed, game) {
			compatibleGames = append(compatibleGames, game)
			fmt.Println(game.ID)
		}
	}
	fmt.Println(addAll(compatibleGames))
}

func formatPlayedGames(game string) Game {
	total := Game{}

	// Assign game ID from the string rather than use some iterative counter
	gameId := regexp.MustCompile(`\d+`)
	total.ID, _ = strconv.Atoi(string(gameId.Find([]byte(game))))

	// Do not need the `Game 1: `
	prefix := regexp.MustCompile(`Game \d+: `)
	games := prefix.ReplaceAllString(game, "")
	//fmt.Println(games)

	// From: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// To: ["3 blue, 4 red", " 1 red, 2 green, 6 blue", " 2 green"]
	gamesArr := strings.Split(games, ";")
	//fmt.Println(gamesArr)

	// Take each of the rolls and parse them
	for _, v := range gamesArr {
		total.Rolls = append(total.Rolls, parseRoll(v))
	}

	return total
}

// Takes a string of a roll in format " 2 red, 4 green, 0 blue"
// and uses regular expression to parse out the number of each color
func parseRoll(roll string) Roll {
	tmp := Roll{}

	rollRegex := regexp.MustCompile(`\d+ (red|green|blue)`)
	findNum := regexp.MustCompile(`\d+`)
	draws := rollRegex.FindAll([]byte(roll), -1)

	for _, draw := range draws {
		if strings.Contains(string(draw), "red") {
			// Parse first value for int
			num, err := strconv.Atoi(string(findNum.Find(draw)))
			if err != nil {
				fmt.Println(err)
			}
			tmp.Red = num
		}
		if strings.Contains(string(draw), "blue") {
			// Parse first value for int
			num, err := strconv.Atoi(string(findNum.Find(draw)))
			if err != nil {
				fmt.Println(err)
			}
			tmp.Blue = num
		}
		if strings.Contains(string(draw), "green") {
			// Parse first value for int
			num, err := strconv.Atoi(string(findNum.Find(draw)))
			if err != nil {
				fmt.Println(err)
			}
			tmp.Green = num
		}
	}
	return tmp
}

func checkCompatibleGame(roll Roll, game Game) bool {
	compatible := true
	for _, v := range game.Rolls {
		//fmt.Println(v)
		if v.Blue > roll.Blue || v.Green > roll.Green || v.Red > roll.Red {
			compatible = false
		}
	}
	return compatible
}

func addAll(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.ID
	}
	return sum
}
