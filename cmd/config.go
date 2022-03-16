/*
Copyright © 2021 weak_ptr <weak_ptr@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "查看和修改设置",
	Long:  `查看和修改设置`,
	Run: func(cmd *cobra.Command, args []string) {
		setConfig, err := cmd.Flags().GetString("set")
		if err != nil {
			panic(err)
		}

		if setConfig != "" {
			pairs := strings.SplitN(setConfig, "=", 2)
			if len(pairs) < 2 {
				println("--set 仅接受 entry=value 形式的参数")
				os.Exit(1)
			}

			viper.Set(pairs[0], pairs[1])
			err = viper.GetViper().WriteConfig()
			cobra.CheckErr(err)
			return
		}

		getConfig, err := cmd.Flags().GetString("get")
		if err != nil {
			panic(err)
		}

		if getConfig != "" {
			fmt.Printf("%s => %v\n", getConfig, viper.Get(getConfig))
			return
		}

		showConfig, err := cmd.Flags().GetBool("show")
		if err != nil {
			panic(err)
		}

		if showConfig {
			for _, k := range viper.AllKeys() {
				println(k, "=>", fmt.Sprintf("%v", viper.Get(k)))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")
	configCmd.Flags().StringP("set", "", "", "接受 entry=value 形式的设置")
	configCmd.Flags().StringP("get", "", "", "获取指定配置项的值")
	configCmd.Flags().BoolP("show", "", false, "展示配置")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkConfig() error {
	if viper.GetString("ak") == "" {
		return errors.New("配置项 `ak` 必须有值")
	}

	if viper.GetString("sk") == "" {
		return errors.New("配置项 `sk` 必须有值")
	}

	if viper.GetString("bucket") == "" {
		return errors.New("配置项 `bucket` 必须有值")
	}

	return nil
}
