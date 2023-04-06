package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	services "github.com/siddhantprateek/go-gpt/core"

	"github.com/spf13/cobra"
)

func PrintName() *cobra.Command {

	cmdPrint := &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
			 For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	return cmdPrint
}

func Echos() *cobra.Command {
	cmdEcho := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Print: " + strings.Join(args, " "))
		},
	}
	return cmdEcho
}

func Times() *cobra.Command {
	var echoTimes int
	cmdTimes := &cobra.Command{
		Use:   "times [# times] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
	a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				pterm.DefaultParagraph.Println("Echo: " + pterm.LightMagenta(strings.Join(args, " ")))
			}
		},
	}
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	return cmdTimes
}

func ChatCompletion3() *cobra.Command {

	cmdGPT3 := &cobra.Command{
		Use:   "chat [string to query.]",
		Short: "Write anything to GPT-3 chat completion.",
		Long:  "Write anything to GPT-3 chat completion.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			gptClient := services.GPTInstance{}
			responseMsg := gptClient.RapidChatGPT(strings.Join(args, " "))
			fmt.Println(responseMsg)
		},
	}

	return cmdGPT3
}

func GPTImage() *cobra.Command {

	cmdGPTImage := &cobra.Command{
		Use:   "image [image description.]",
		Short: "Generate image using OpenAI",
		Long:  "Generate image using OpenAI.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			img1, img2 := services.GPTImageGen(strings.Join(args, " "))
			fmt.Println(img1)
			fmt.Println(img2)
		},
	}

	return cmdGPTImage
}

func GPTModels() *cobra.Command {

	cmdModel := &cobra.Command{
		Use:   "models",
		Short: "Get all available OpenAI models",
		Long:  "Get all available OpenAI models",
		Run: func(cmd *cobra.Command, args []string) {
			models := services.GPTModels()
			for idx, model := range models {
				fmt.Println("[" + strconv.Itoa(idx) + "]" + ":" + model)
			}

		},
	}

	return cmdModel
}

func Moderation() *cobra.Command {

	cmdModeration := &cobra.Command{
		Use:   "moderation [string *.]",
		Short: "Text Moderation Analysis.",
		Long:  "Text Moderation Analysis.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			services.GPTModeration(strings.Join(args, " "))

		},
	}

	return cmdModeration
}
