package greeting

import (
	"testing"
	"strings"
)
//TDD: Test Driven Development
func TestGreetingYourName(t *testing.T){

	given := "Hello, Bob."
	want := "Hello, Bob."

	get := Greet(given)

	if want != get {
		t.Errorf("given a name %s want greeting %q,but got %q",given,want,get)
	}
}

func TestGreetingYourFriend(t *testing.T){

	given := "Hello"
	want := "Hello, My friend."
	given += ", My friend."


	get := Greet(given)

	if want != get {
		t.Errorf("given a name %s want greeting %q,but got %q",given,want,get)
	}
}


func TestGreetingCaptital(t *testing.T){

	given := "bob."
	want := "HELLO, BOB."
	given = "HELLO, " +strings.ToUpper(given)
	get := Greet(given)

	// if given == strings.ToUpper(name) {
	// 	// ...
	// }

	if want != get {
		t.Errorf("given a name %s want greeting %q,but got %q",given,want,get)
	}
}