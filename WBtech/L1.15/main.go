package main
import (
	"fmt"
)

func createHugeString(size int) string {
	new_string := make([]rune, size)
	for i := 0; i < size; i++ {
		new_string[i] = 'a'
	}
	return string(new_string)
}

func someFunc() string{
  v := createHugeString(1 << 10)
  return string(v[:100])

}
func main() {
	new_string2 := someFunc()
  	fmt.Println(new_string2)

}