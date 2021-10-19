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
	"fmt"
	"time"

	"github.com/nnnewb/qn/internal/utils"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "列出文件",
	Long:  `列出文件`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		credential := auth.New(viper.GetString("ak"), viper.GetString("sk"))
		mgr := storage.NewBucketManager(credential, &storage.Config{UseCdnDomains: false})
		bucket := viper.GetString("bucket")

		delimiter, err := cmd.Flags().GetString("delimiter")
		cobra.CheckErr(err)

		noDelimiter, err := cmd.Flags().GetBool("no-delimiter")
		cobra.CheckErr(err)

		if noDelimiter {
			delimiter = ""
		}

		ch, err := mgr.ListBucket(bucket, args[0], delimiter, "")
		cobra.CheckErr(err)

		for v := range ch {
			fmt.Printf("%s  %10s   %s\n", time.UnixMicro(v.Item.PutTime/10).Format(time.RFC3339), utils.ByteCountIEC(v.Item.Fsize), v.Item.Key)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	lsCmd.Flags().BoolP("no-delimiter", "", false, "不使用任何路径分隔符，ls命令会递归显示所有文件。")
	lsCmd.Flags().StringP("delimiter", "d", "/", "指定路径分隔符，ls命令会以unix-like的方式工作。")
}
