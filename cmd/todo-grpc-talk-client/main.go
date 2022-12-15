package main

import (
	"fmt"
	"os"

	"github.com/CDimonaco/todo-grpc-talk/pkg/client"
	"github.com/spf13/cobra"
)

var serviceUrl string

var rootCmd = &cobra.Command{
	Use:   "todo-grpc",
	Short: "Todo grpc client",
}

var addTodo = &cobra.Command{
	Use:   "add-todo [name] [description]",
	Short: "Create a new todo, pass name and description as arguments",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := client.NewTodoGrpcClient(serviceUrl)
		if err != nil {
			return err
		}

		t, err := client.AddTodo(cmd.Context(), args[0], args[1])
		if err != nil {
			return err
		}

		os.Stdout.WriteString(
			fmt.Sprintf("New todo created \nName: %s\nDescription: %s\nID: %s", t.Title, t.Description, t.ID),
		)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addTodo)

	rootCmd.Flags().StringVarP(&serviceUrl, "url", "u", "localhost:9090", "grpc service root url")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
