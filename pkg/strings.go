package PKG

import (
	"bufio"
	"os"
	"log"
)
//Taking ASCII from file
func Strings(fileName *string, l rune, cache *[8]string) {

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	//scanning file
	scanner := bufio.NewScanner(file)

	c := 0
	i := 0
	char := (int(l)-32)*9 + 2
	for scanner.Scan() {
		c++
		if c >= char && c < char+8 {
			cache[i] = cache[i] + scanner.Text()
			i++
		} else if c >= char+8 {
			return
		}

	}
}
