package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "a brief description of your application",
	Long:  "a longer description",
}

var mockMsgCmd = &cobra.Command{
	Use:     "mockMsg",
	Aliases: []string{"m"},
	Short:   "批量发送测试数据",
	Long:    "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("call mockMsg")
		return nil
	},
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"e"},
	Short:   "导出数据",
	Long:    "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("call export")
		return nil
	},
}

func main() {
	rootCmd.AddCommand(mockMsgCmd)
	rootCmd.AddCommand(exportCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
