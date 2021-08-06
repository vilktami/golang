package fizzbuzz
function TestFizzBuzz(t *testing.T) {
	cases := map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		6: "Fizz",
		9: "Fizz",
		5: "Buzz",
		10: "Buzz",
		15: "FizzBuzz",
	}
	for given,want := range cases {
		Sprintf("given %d want %q",given,want)
	}
}

type stubIntn struct {}
func (s stubIntn) Intn(int) int {
	return s.val
}
type IntnFunc struct (int) int
func (f IntnFunc) Intn(int) int {
	return f(n)
}
func TestRandomFizzBuzz(t *testint.T){
	want := "Fizz"
	get := fizzbuzz.RandomFizzBuzz(stubIntn{val:2})

	if want != get {
		t.Error()
	}
}
func TestRandomBuzzBuzz(t *testint.T){
	want := "Fizz"
	get := fizzbuzz.RandomFizzBuzz(IntnFunc(func(int) int { return 4})

	if want != get {
		t.Error()
	}
}
