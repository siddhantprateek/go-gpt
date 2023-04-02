package main

import (
	"fmt"

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

func main() {
	fmt.Printf("%s\n\n", go_gpt_greeting)
}

func InitCmd() *cobra.Command {

	var gpt_cmd = &cobra.Command{
		Use:   "Go-GPT",
		Short: "Go-GPT is a simple openAI CLI built using Go.",
	}

	return gpt_cmd
}
