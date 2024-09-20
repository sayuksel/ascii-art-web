package PKG

import (
	"fmt"
	"os"
	"strings"
)
//Handling any errors if file is not present and sending the correct status
func Text(text string, art string) (string, error) {
	fmt.Println(text)
	fmt.Println(art)

	var fileName string
	switch art {
	case "standard":
		fileName = ("../banners/standard.txt")
	case "shadow":
		fileName = ("../banners/shadow.txt")
	case "thinkertoy":
		fileName = ("../banners/thinkertoy.txt")
	default:
		return "500", nil		
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("1")
		return "500", fmt.Errorf("Error: cannot read banner file ")
	}
	defer file.Close()
	for i, v := range text {
		if v < 32 || v > 126 {
			if v == 10 {
				fmt.Println(v, i)
				continue
			}
			if v == 13 {
				fmt.Println(text[:i], v, i)
				text = text[:i] + "\\n" + text[i+2:]
				i += 2
				continue
			}
			return "400", nil
		}
	}
	if len(strings.Trim(text, " ")) == 0 {
		return "400", nil
	}

	sections := strings.Split(text, "\\n")
	if len(sections) >= 2 && sections[0] == "" && sections[1] == "" {
		fmt.Println()
		sections = sections[2:]

	}
	textArt := "<pre>"
	for _, s := range sections {
		if s == "" {

			textArt += "<br>"
			continue
		}
		cache := [8]string{}
		for _, r := range s {
			Strings(&fileName, r, &cache)
		}
		PrintA(&cache, &textArt)
	}

	return textArt + "</pre>", nil
}
