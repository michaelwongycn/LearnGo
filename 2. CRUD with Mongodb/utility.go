package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func check(err error) {

	if err != nil {
		panic(err)
	}
}

func ReadConfig() (Config map[string]string) {

	config := make(map[string]string)

	file, err := os.Open("config.txt")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		check(err)
	}
	return config
}

func connect(uri string) (*mongo.Client, context.Context) {

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	check(err)

	return client, ctx
}

func ping(client *mongo.Client, ctx context.Context) {

	err := client.Ping(ctx, readpref.Primary())

	check(err)
}

func disconnect(client *mongo.Client, ctx context.Context) {

	defer func() {
		err := client.Disconnect(ctx)
		check(err)
	}()
}
