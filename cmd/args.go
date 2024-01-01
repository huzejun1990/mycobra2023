// @Author huzejun 2024/1/1 19:05:00
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var level2Cmd = &cobra.Command{
	Use: "l2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("l2 cmd run begin")
		fmt.Println(cmd.Flags().Lookup("viper").Value)
		fmt.Println(cmd.Flags().Lookup("author").Value)
		fmt.Println(cmd.Flags().Lookup("config").Value)
		fmt.Println(cmd.Flags().Lookup("license").Value)
		fmt.Println(viper.GetString("author"))
		fmt.Println(viper.GetString("license"))
		fmt.Println("l2 cmd run end")
	},
}

var curArgsCheckCmd = &cobra.Command{
	Use: "cuscheck",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cuscheck cmd run end")
		fmt.Println(args)
		fmt.Println("cuscheck cmd run end")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("至多输入两个参数")
		}
		return nil
	},
}

var onlyArgsCmd = &cobra.Command{
	Use:       "only",
	ValidArgs: []string{"123", "abc", "nick"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("only cmd run end")
		fmt.Println(args)
		fmt.Println("only cmd run end")
	},
}

var rangeArgsCmd = &cobra.Command{
	Use: "range",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("range cmd run end")
		fmt.Println(args)
		fmt.Println("range cmd run end")
	},
	Args: cobra.RangeArgs(2, 4),
}

func init() {
	initCmd.AddCommand(level2Cmd)
	rootCmd.AddCommand(curArgsCheckCmd)
	rootCmd.AddCommand(onlyArgsCmd)
	rootCmd.AddCommand(rangeArgsCmd)
}
