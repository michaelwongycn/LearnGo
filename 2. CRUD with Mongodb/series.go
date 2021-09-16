package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// type series struct {
// 	id       primitive.ObjectID
// 	title    string
// 	author   string
// 	language string
// }

var (
	config         (map[string]string) = ReadConfig()
	database       string              = config["database"]
	collectionName string              = config["seriesCollection"]
)

func printSeriesMenu() {
	fmt.Println("Welcome to BookInc series management system")
	fmt.Println("Command list")
	fmt.Println("1. View All Series")
	fmt.Println("2. View a Series")
	fmt.Println("3. Insert New Series")
	fmt.Println("4. Update a Series")
	fmt.Println("5. Delete a Series")
	fmt.Println("0. Back")
	fmt.Print("Please Input your command : ")
}

func printSeriesSelector() {
	fmt.Println("Choose field to search")
	fmt.Println("Command list")
	fmt.Println("1. Title")
	fmt.Println("2. Author")
	fmt.Println("3. Language")
	fmt.Println("0. Back")
	fmt.Print("Please Input your command : ")
}

func printSeries(result []bson.M) {
	fmt.Println("=========================")
	for index, doc := range result {
		fmt.Printf("Book No.	: %d\n", index+1)
		fmt.Printf("Series Title	: %s\n", doc["title"])
		fmt.Printf("Series Author	: %s\n", doc["author"])
		fmt.Printf("Series Language	: %s\n", doc["language"])
	}
	fmt.Println("=========================")
}

func findAllSeries(client *mongo.Client, ctx context.Context) []bson.M {

	filter := bson.M{}
	cursor := findDocument(client, ctx, database, collectionName, filter)

	var result []bson.M
	cursor.All(ctx, &result)

	printSeries(result)

	return result
}

func findSeries(client *mongo.Client, ctx context.Context) {
	key := ""
	query := ""
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printSeriesSelector()
		scanner.Scan()
		command, _ := strconv.Atoi(scanner.Text())

		if command == 0 {
			break
		}

		switch command {
		case 1:
			fmt.Print("Input Title to Search : ")
			scanner.Scan()
			query = scanner.Text()
			key = "title"
		case 2:
			fmt.Print("Input Author to Search : ")
			scanner.Scan()
			query = scanner.Text()
			key = "author"
		case 3:
			fmt.Print("Input Language to Search : ")
			scanner.Scan()
			query = scanner.Text()
			key = "language"
		}

		filter := bson.M{key: query}
		cursor := findDocument(client, ctx, database, collectionName, filter)

		var result []bson.M
		cursor.All(ctx, &result)

		printSeries(result)
	}
}

func insertSeries(client *mongo.Client, ctx context.Context) {
	title := ""
	author := ""
	language := ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input Series's Title : ")
	scanner.Scan()
	title = scanner.Text()
	fmt.Print("Input Series's Author : ")
	scanner.Scan()
	author = scanner.Text()
	fmt.Print("Input Series's Language : ")
	scanner.Scan()
	language = scanner.Text()

	filter := bson.M{"title": title, "author": author, "language": language}
	cursor := findDocument(client, ctx, database, collectionName, filter)

	var result []bson.M
	cursor.All(ctx, &result)

	if len(result) > 0 {
		fmt.Println("Exact Series Already Exist")
	} else {
		insertDocument(client, ctx, database, collectionName, filter)
	}
}

func updateSeries(client *mongo.Client, ctx context.Context) {
	title := ""
	author := ""
	language := ""
	result := findAllSeries(client, ctx)
	scanner := bufio.NewScanner(os.Stdin)

	if len(result) == 0 {
		fmt.Println("There are No Series In the System")
	} else {
		fmt.Printf("Input The Series No You Want to Edit (1 - %d) : \n", len(result))
		scanner.Scan()
		command, _ := strconv.Atoi(scanner.Text())

		if command < 1 || command > len(result) {
			fmt.Printf("Please Input Between (1 - %d) : \n", len(result))
		} else {
			fmt.Print("Input Series's New Title : ")
			scanner.Scan()
			title = scanner.Text()
			fmt.Print("Input Series's New Author : ")
			scanner.Scan()
			author = scanner.Text()
			fmt.Print("Input Series's New Language : ")
			scanner.Scan()
			language = scanner.Text()

			_id := bson.M{"_id": result[command-1]["_id"]}
			document := bson.M{"$set": bson.M{"title": title, "author": author, "language": language}}

			updateDocument(client, ctx, database, collectionName, _id, document)
		}
	}

}

func deleteSeries(client *mongo.Client, ctx context.Context) {
	stringCommand := ""
	result := findAllSeries(client, ctx)
	scanner := bufio.NewScanner(os.Stdin)

	if len(result) == 0 {
		fmt.Println("There are No Series In the System")
	} else {
		fmt.Printf("Input The Series No You Want to Delete (1 - %d) : \n", len(result))
		scanner.Scan()
		command, _ := strconv.Atoi(scanner.Text())

		if command < 1 || command > len(result) {
			fmt.Printf("Please Input Between (1 - %d) : \n", len(result))
		} else {
			_id := bson.M{"_id": result[command-1]["_id"]}

			fmt.Printf("Are You Sure Want to Delete Series : %s ?", result[command-1]["title"])
			scanner.Scan()
			stringCommand = scanner.Text()

			switch stringCommand {
			case "yes", "Yes", "yEs", "yeS", "YEs", "yES", "YeS", "YES":
				deleteDocument(client, ctx, database, collectionName, _id)
			default:
				break
			}
		}
	}
}
