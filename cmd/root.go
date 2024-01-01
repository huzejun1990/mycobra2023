// @Author huzejun 2023/12/30 9:02:00
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	// 命令行使用
	Use: "root",
	//简短的提示
	Short: "short desc",
	//详细的提示
	Long: "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(cmd.Flags().Lookup("viper").Value)
		fmt.Println(cmd.Flags().Lookup("author").Value)
		fmt.Println(cmd.Flags().Lookup("config").Value)
		fmt.Println(cmd.Flags().Lookup("license").Value)
		fmt.Println(cmd.Flags().Lookup("source").Value)
		fmt.Println(viper.GetString("author"))
		fmt.Println(viper.GetString("license"))

		fmt.Println("root cmd run end")
	},
}

func Execute() {
	rootCmd.Execute()
}

var cfgFile string
var userLicense string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().Bool("viper", true, "")
	rootCmd.PersistentFlags().StringP("author", "a", "Your Name", "")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")
	rootCmd.Flags().StringP("source", "s", "", "")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	//viper.SetDefault()
}

func initConfig() {
	if cfgFile == "" {
		return
	}
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		cobra.CheckErr(err)
	}
	fmt.Println("using config file:", viper.ConfigFileUsed())
}
