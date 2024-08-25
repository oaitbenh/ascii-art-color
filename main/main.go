package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"Ascii"
	Cl "Ascii/Colors"
)

var (
	OutputFile = ""
	Text       = ""
	Banner     = "standard.txt"
	Align      = "left"
	Size       = Ascii.SizeOfTerminal()
	Justify    = false
	Color      = ""
	ColorFlag  = false
	SubString  = ""
)

func main() {
	ArgsHandler()
	_, Folderename, _, _ := runtime.Caller(0)
	SplitBanner := strings.Split("/"+Banner, "/")
	Banner = SplitBanner[len(SplitBanner)-1]
	AsciiFile, err := os.ReadFile(path.Join(path.Dir(Folderename), "/../Banners", Banner))
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		os.Exit(0)
	}
	AsciiFile = []byte(strings.ReplaceAll(string(AsciiFile), "\r", ""))
	AsciiFile = AsciiFile[1 : len(AsciiFile)-1]
	SplitedFileContent := strings.Split(string(AsciiFile), "\n\n")
	for _, char := range Text {
		if int(char) > 126 || int(char) < 32 {
			log.Fatal("This String Contains A Character That Is Imprintable")
		}
	}
	SliceText := strings.Split(Text, "\\n")
	if strings.ReplaceAll(Text, "\\n", "") == "" {
		SliceText = SliceText[1:]
	}
	AsciiArt := ""
	for _, Txt := range SliceText {
		if Txt == "" {
			AsciiArt += "\n"
			continue
		}
		AsciiArt += Ascii.PrintAscii(Align, SplitedFileContent, Size, Txt, ColorFlag, Color, SubString)
	}
	if OutputFile != "" {
		if !strings.HasSuffix(OutputFile, ".txt") {
			OutputFile += ".txt"
		}
		SplitOutput := strings.Split("/"+OutputFile, "/")
		OutputFile = SplitOutput[len(SplitOutput)-1]
		os.WriteFile("../Output/"+OutputFile, []byte(AsciiArt), 0o666)
	} else {
		fmt.Print(AsciiArt)
	}
}

func ArgsHandler() {
	switch len(os.Args) {
	case 5:
		if strings.HasPrefix(os.Args[1], "--color=") {
			ColorFlag = true
			Color = strings.TrimPrefix(os.Args[1], "--color=")
			if strings.HasPrefix(Color, "#") {
				Color, _ = Cl.HexToRgb(Color)
			} else {
				Color = Cl.RgbColor(Cl.Colors[Color].R, Cl.Colors[Color].G, Cl.Colors[Color].B)
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			os.Exit(0)
		}
		SubString = os.Args[2]
		Text = os.Args[3]
		Banner = os.Args[4] + ".txt"
	case 4:
		if strings.HasPrefix(os.Args[1], "--color=") {
			ColorFlag = true
			Color = strings.TrimPrefix(os.Args[1], "--color=")
			if strings.HasPrefix(Color, "#") {
				Color, _ = Cl.HexToRgb(Color)
			} else {
				Color = Cl.RgbColor(Cl.Colors[Color].R, Cl.Colors[Color].G, Cl.Colors[Color].B)
			}
			if os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy" {
				SubString = os.Args[2]
				Text = os.Args[3]
			} else {
				Text = os.Args[2]
				Banner = os.Args[3] + ".txt"
			}
		} else if strings.HasPrefix(os.Args[1], "--align=") {
			Align = strings.TrimPrefix(os.Args[1], "--align=")
			Text = os.Args[2]
			Banner = os.Args[3] + ".txt"
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			OutputFile = strings.TrimPrefix(os.Args[1], "--output=")
			if OutputFile == "" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				os.Exit(0)
			}
			Text = os.Args[2]
			Banner = os.Args[3] + ".txt"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			os.Exit(0)
		}
	case 3:
		if strings.HasPrefix(os.Args[1], "--color=") {
			ColorFlag = true
			Color = strings.TrimPrefix(os.Args[1], "--color=")
			if strings.HasPrefix(Color, "#") {
				Color, _ = Cl.HexToRgb(Color)
			} else {
				Color = Cl.RgbColor(Cl.Colors[Color].R, Cl.Colors[Color].G, Cl.Colors[Color].B)
			}
			Text = os.Args[2]
		} else if strings.HasPrefix(os.Args[1], "--align=") {
			Align = strings.TrimPrefix(os.Args[1], "--align=")
			Text = os.Args[2]
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			OutputFile = strings.TrimPrefix(os.Args[1], "--output=")
			if OutputFile == "" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				os.Exit(0)
			}
			Text = os.Args[2]
		} else {
			Banner = os.Args[2] + ".txt"
			Text = os.Args[1]
		}
	case 2:
		Text = os.Args[1]
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		os.Exit(0)
	}
}
