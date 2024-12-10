package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sea_battle/game"
	"strconv"
	"strings"
)

func GetUserInput(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	result, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occurred while reading the entered text! Please try again")
	}

	return strings.TrimRight(result, "\r\n")
}

func ParseUserInputToCords(input string) (*game.Cords, error) {
	if strings.Contains(input, " ") {
		return nil, errors.New("input should not contain spaces")
	}

	if len(input) != 2 {
		return nil, errors.New("input should be like 'A6' or 'J5'")
	}

	var x int
	var isFoundRune bool = false

	firstRune, secondRune := input[0], input[1]

	for i, value := range game.RowsAlphabet {
		if strings.ToUpper(string(firstRune)) == value {
			x = i + 1
			isFoundRune = true
		}
	}

	if !isFoundRune {
		return nil, errors.New("first symbol is incorrect")
	}

	value, err := strconv.ParseInt(string(secondRune), 10, 8)

	if err != nil {
		return nil, errors.New("second symbol is incorrect")
	}

	if value > 10 {
		return nil, errors.New("second symbol too big. Should be less equal then 10")
	}

	coords := game.Cords{
		X: x,
		Y: int(value),
	}

	return &coords, nil
}
