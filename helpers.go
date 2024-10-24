package swissknife

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

const introMessage = "If you have a bit of extra coin, you can drop the developer a coffee.\n" +
	"So that he gets drunk on coffee and happily runs to develop this and other projects"
const previewColor = "green"

// PrintIntroMessage - publishes a funny message in the log asking to donate to the developer
func PrintIntroMessage(appName, cryptoAddress string, cryptoCoinTag ...string) {
	coinTag := ""
	if len(cryptoCoinTag) > 0 {
		coinTag = cryptoCoinTag[0]
	}

	figure.NewColorFigure(appName, "", previewColor, true).Print()
	fmt.Println()
	fmt.Println(introMessage)
	color.Green(cryptoAddress + " " + coinTag)
	fmt.Println()
}

func PrintObject(o any) {
	data, err := json.MarshalIndent(o, "", "	")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
