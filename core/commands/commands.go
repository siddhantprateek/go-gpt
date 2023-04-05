package commands

import (
	"fmt"
	"strings"

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
			fmt.Println("Print: " + strings.Join(args, " "))
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
				fmt.Println("Echo: " + strings.Join(args, " "))
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
