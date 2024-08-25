package Ascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func PrintAscii(Align string, data []string, Size int, Txt string, ColorFlag bool, Color string, SubString string) string {
	ColorsIndex := CheckSubString(Txt, SubString)

	if len(data) != 95 {
		log.Fatal("The file does not contain all the Ascii art symbols!")
	}
	for _, char := range data {
		if len(strings.Split(char, "\n")) != 8 {
			log.Fatal("the size of one of the characters does not match required size")
		}
	}
	Art := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(Txt); j++ {
			if ColorFlag && (ColorsIndex[j] || SubString == "") {
				Art += Color + strings.Split(data[int(Txt[j]-32)], "\n")[i] + "\033[0m"
			} else if strings.ToLower(Align) == "justify" && Txt[j] == ' ' {
				SZ := (Size - len(ArtWithoutSpace(Txt, data))) / SizeOfChar(Txt, ' ')
				Art += Pr(" ", SZ)
			} else {
				Art += strings.Split(data[int(Txt[j]-32)], "\n")[i]
			}
		}
		Art += "\n"
	}
	switch strings.ToLower(Align) {
	case "left":
		Size = 0
	case "right":
		Size = Size - len(strings.Split(Art, "\n")[0])
	case "center":
		Size = Size/2 - len(strings.Split(Art, "\n")[0])/2
	case "justify":
		Size = 0
	default:
		log.Fatal(Align, "is not a Position")
	}
	ResArt := ""
	for i, line := range strings.Split(Art, "\n") {
		if i != len(strings.Split(Art, "\n"))-1 {
			ResArt += Pr(" ", Size) + line + "\n"
		}
	}
	return ResArt
}

func Pr(add string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += add
	}
	return res
}

func ArtWithoutSpace(txt string, data []string) string {
	Art := ""
	for j := 0; j < len(txt); j++ {
		if txt[j] != ' ' {
			Art += strings.Split(data[int(txt[j]-32)], "\n")[0]
		}
	}
	return Art
}

func SizeOfChar(txt string, char rune) int {
	count := 0
	for _, c := range txt {
		if char == c {
			count++
		}
	}
	return count
}

func SizeOfTerminal() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	Out, err := cmd.Output()
	if err != nil {
		log.Fatal(err, "Error in get Terminal size")
	}
	Out = Out[:len(Out)-1]
	Number, err := strconv.Atoi(strings.Split(string(Out), " ")[1])
	if err != nil {
		log.Fatal(err, "Error in convert to Number")
	}
	return Number
}

func CheckSubString(Text, SubString string) map[int]bool {
	res := map[int]bool{}
	for i := 0; i < len(Text)-len(SubString)+1; i++ {
		if SubString == Text[i:i+len(SubString)] {
			for j := i; j < i+len(SubString); j++ {
				res[j] = true
			}
		}
	}
	return res
}
