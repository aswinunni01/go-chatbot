package main

type DumbResponder struct {
}

func (conn DumbResponder) Respond(input string) (Response, error) {
	var response string
	switch input {
	case "What is my name":
		response = "Aswin"
	case "What is my age":
		response = "22"
	default:
		response = "Invalid input"
	}

	return Response{
		ID:      "0",
		Message: response,
		Time:    0,
		Choices: 1,
	}, nil
}
