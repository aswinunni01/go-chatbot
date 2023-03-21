package main

import "testing"

func TestGetResponse(t *testing.T) {

	t.Run("Testing name response", func(t *testing.T) {
		want := "Aswin"
		get, _ := getResponse("What is my name", DumbResponder{})

		if want != get {
			t.Errorf("got %q want %q", get, want)
		}
	})

	t.Run("Testing age response", func(t *testing.T) {
		want := "22"
		get, _ := getResponse("What is my age", DumbResponder{})

		if want != get {
			t.Errorf("got %q want %q", get, want)
		}
	})
}

func TestCleanInput(t *testing.T) {

	want := "Aswin"
	get := cleanInput("Aswin\n")

	if want != get {
		t.Errorf("got %q want %q", get, want)
	}

}
