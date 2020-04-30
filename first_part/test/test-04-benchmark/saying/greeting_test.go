package saying

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	result := Greeting("bob")
	if result != "welcome bob !!!" {
		t.Error("expected welcome bob !!! but got ", result)
	}
}

func ExampleGreeting() {
	fmt.Println(Greeting("bob"))
	//Output:
	//welcome bob !!!
}
func BenchmarkGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("bob")
	}
}
