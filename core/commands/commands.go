package commands

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	svc "github.com/siddhantprateek/go-gpt/core"

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
	cmdres := &cobra.Command{
		Use:   "chat [string to query.]",
		Short: "Write anything to GPT-3 chat completion.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// gptClient := svc.GPTInstance{}
			responseMsg := svc.RapidChatGPT(strings.Join(args, " "))
			fmt.Println(responseMsg)
			// cmd.Print(responseMsg)
		},
	}
	return cmdres
}

func GPTImage() *cobra.Command {
	cmdImage := &cobra.Command{
		Use:   "image [image description.]",
		Short: "Generate image using OpenAI",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			img1, img2 := svc.ImageGeneration(strings.Join(args, " "))
			fmt.Println(img1)
			fmt.Println(img2)
		},
	}

	return cmdImage
}

func GPTModels() *cobra.Command {

	cmdModel := &cobra.Command{
		Use:   "models",
		Short: "Get all available OpenAI models",
		Run: func(_ *cobra.Command, _ []string) {
			models := svc.GetAllModels()
			fmt.Println("All Available Models.")
			fmt.Printf("\n%4s |  %s\n", "Model Idx", "Models")
			for idx, model := range models {
				fmt.Printf("%8d  :  %s\n", idx, model)
			}
		},
	}
	return cmdModel
}

func Moderation() *cobra.Command {
	cmdModeration := &cobra.Command{
		Use:   "moderation [string *.]",
		Short: "Text Moderation Analysis.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			svc.TxtModeration(strings.Join(args, " "))
		},
	}
	return cmdModeration
}

func SentimentCmd() *cobra.Command {
	cmdSentiment := &cobra.Command{
		Use:   "sentiment [string *.]",
		Short: "Sentiment Analysis.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			response := svc.SentimentAnalysis(strings.Join(args, " "))

			fmt.Println("Positive:", response.Postive)
			fmt.Println("Negative:", response.Negative)
			fmt.Println("Neutral:", response.Neutral)

		},
	}
	return cmdSentiment
}
