package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func getResponse(input string, responder Responder) (string, error) {
	response, _ := responder.Respond(input)
	fmt.Println("Recieved Response ID: ", response.ID, " at time: ", response.Time, " numChoices: ", response.Choices)

	return response.Message, nil

	// return testOpenAI(input)
}

func cleanInput(input string) string {
	input = input[:len(input)-1]
	if len(input) < 1 {
		input = "Empty"
	}
	return input
}

// Takes in a line and returns corresponding response!
// Uses chat GPT underneath for the reponse
func main() {
	// Get input from user
	var input string
	inpBuf := bufio.NewReader(os.Stdin)

	input, _ = inpBuf.ReadString('\n')
	input = cleanInput(input)

	// Create a responder client
	token := os.Getenv("OPENAPI_TOKEN")
	responder := openAIConn{
		client: openai.NewClient(token),
	}

	// Get response from responder client
	response, err := getResponse(input, responder)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}
