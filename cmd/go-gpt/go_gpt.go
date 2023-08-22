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

var (
	creator = "siddhantprateek"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-gpt",
		Short: "Go-GPT is a simple openAI CLI built using Go.",
		Long:  `All software has versions. This is Go-GPTS's version 0.1.0`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n\n", go_gpt_greeting)
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	rootCmd.PersistentFlags().StringVar(&author, "author", creator, "Copyrights belongs to @siddhantprateek")

	rootCmd.AddCommand(
		commands.PrintName(),
		commands.ChatCompletion3(),
		commands.GPTImage(),
		commands.GPTModels(),
		commands.Moderation(),
		commands.SentimentCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
