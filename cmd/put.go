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
	"path"

	"github.com/nnnewb/qn/internal/utils"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "上传文件到七牛云",
	Long:  "上传文件到七牛云",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(checkConfig())

		credential := auth.New(viper.GetString("ak"), viper.GetString("sk"))
		dest, err := cmd.Flags().GetString("dest")
		cobra.CheckErr(err)

		for _, filepath := range args {
			key := path.Join(dest, path.Base(filepath))
			err = utils.Upload(viper.GetString("bucket"), key, filepath, credential)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	putCmd.Flags().StringP("dest", "d", "", "上传位置，最终key会拼接成 <dest>/filename 的形式")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
