package main

import "localstack-examples/config"

func main() {
	client := config.NewAWSClient()
	err := client.ListSQSQueues()
	if err != nil {
		panic(err)
	}
	err = client.ListSNSTopics()
	if err != nil {
		panic(err)
	}
}
