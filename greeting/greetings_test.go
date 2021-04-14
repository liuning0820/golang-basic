package greetings_test
 
import "testing"
import "math"
import "fmt"
 
func TestMathBasics(t *testing.T) {
    v := math.Min(10, 0)
    if v != 0 {
        t.Error("Failed the test!")
    }
}


func WrongMin(x, y float64) float64 {
    if x > y {
        return y
    } else {
        return x
    }
}

func TestMathBasicsFail(t *testing.T) {
    v := WrongMin(10, 0)
	
	fmt.Printf("%T,", v)
	fmt.Println(v)
    if v != 0 {
        t.Error("Failed the test!")
    }
}