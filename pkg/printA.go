package PKG

// adjusting the text
func PrintA(arr *[8]string, textArt *string) {
	for i, line := range arr {
		*textArt += line + "<br>"
		arr[i] = ""
	}
}
