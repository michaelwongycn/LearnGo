package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printPrompt() {
	fmt.Println("Welcome to BookInc management system")
	fmt.Println("Command list")
	fmt.Println("1. Series Menu")
	fmt.Println("2. Book Menu")
	fmt.Println("0. Exit")
	fmt.Print("Please Input your command : ")
}

func main() {
	config := ReadConfig()
	uri := config["uri"]

	scanner := bufio.NewScanner(os.Stdin)

	client, ctx := connect(uri)
	ping(client, ctx)
	defer disconnect(client, ctx)

	for {
		printPrompt()
		scanner.Scan()
		command, _ := strconv.Atoi(scanner.Text())

		if command == 0 {
			break
		}

		switch command {
		case 1:
			for {
				printSeriesMenu()
				scanner.Scan()
				command, _ := strconv.Atoi(scanner.Text())

				if command == 0 {
					break
				}

				switch command {
				case 1:
					findAllSeries(client, ctx)
				case 2:
					findSeries(client, ctx)
				case 3:
					insertSeries(client, ctx)
				case 4:
					updateSeries(client, ctx)
				case 5:
					deleteSeries(client, ctx)
				}
			}
		case 2:
			for {
				printPrompt()
				break
			}
		}
	}

}
