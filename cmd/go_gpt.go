package main

import (
	"fmt"
	"os"

	commands "github.com/siddhantprateek/go-gpt/core/commands"

	"github.com/spf13/cobra"
)

const go_gpt_greeting = `
	 ____               ____ ____ _____ 
	/ ___|  ___        / ___|  _ \_   _|
	| |  _ / _ \ _____| |  _| |_) || |  
	| |_| | (_) |_____| |_| |  __/ | |  
	\____| \___/       \____|_|    |_|  

  Go-GPT is a simple way to communicate with OpenAI  
    using your CLI created by @siddhantprateek`

// Global variables
var Source string
var Verbose bool
var author string

func main() {

	// fmt.Println(chatgpt.ChatGPT("What is Iot"))
	// api.LatestCoinData()

	rootCmd := &cobra.Command{Use: "go-gpt"}
	rootCmd.AddCommand(InitCmd())
	rootCmd.AddCommand(
		commands.Echos(),
		commands.PrintName(),
		commands.Times(),
		commands.ChatCompletion3(),
		commands.GPTImage(),
		commands.GPTModels(),
		commands.Moderation(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func InitCmd() *cobra.Command {

	fmt.Printf("%s\n\n", go_gpt_greeting)

	gpt_cmd := &cobra.Command{
		Use:   "about",
		Short: "Go-GPT is a simple openAI CLI built using Go.",
		Long:  `All software has versions. This is Go-GPTS's version 0.1.0`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, " + args[0] + "!")
		},
	}

	gpt_cmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	gpt_cmd.PersistentFlags().StringVar(&author, "author", "siddhantprateek", "Copyrights belongs to @siddhantprateek")
	return gpt_cmd
}
