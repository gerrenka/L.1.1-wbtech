package main
import "fmt"

func main() {
	stroka := "главрыба"
	runes := []rune(stroka)
	for i := len(runes)-1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
}