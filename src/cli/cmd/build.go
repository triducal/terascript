/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"terascript/src/ast"
	"terascript/src/compiler"
	"terascript/src/lexer"
	"terascript/src/parser"

	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "build",
	Short: ".sb3 builder and compiler",
	Long:  `This command is used to run the compiler and outputs a .sb3 file in the main directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %s", filePath)
		}

		bytes, _ := os.ReadFile(filePath)
		source := string(bytes)
		lex := lexer.Tokenize(source)

		var sprites = make(map[string]ast.Sprite, 0)

		sprite, _ := parser.Parse("Stage", lex)

		sprites[sprite.Name] = sprite

		project := ast.Project{
			Sprites: sprites,
		}

		sb3 := compiler.Compile(project)
		litter.Dump(sb3)

		// jsonObj, err := json.Marshal(sb3)

		// if err != nil {
		// 	panic(err)
		// }

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
