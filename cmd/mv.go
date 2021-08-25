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
	"strings"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "移动文件",
	Long:  `移动文件`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		credential := auth.New(viper.GetString("ak"), viper.GetString("sk"))
		mgr := storage.NewBucketManager(credential, &storage.Config{UseCdnDomains: false})
		bucket := viper.GetString("bucket")
		overwrite, err := cmd.Flags().GetBool("force")
		cobra.CheckErr(err)

		if len(args) == 2 {
			// 模拟复制到文件夹里这种情况
			src := args[0]
			dest := args[1]
			if strings.HasSuffix(args[1], "/") {
				dest = path.Join(dest, path.Base(src))
			}
			err := mgr.Move(bucket, src, bucket, dest, overwrite)
			cobra.CheckErr(err)
		} else {
			batchOps := make([]string, 0, len(args)-1)
			prefix := args[len(args)-1]
			for i := 0; i < len(args)-1; i++ {
				batchOps = append(batchOps, storage.URIMove(bucket, args[i], bucket, path.Join(prefix, args[i]), overwrite))
			}
			ret, err := mgr.Batch(batchOps)
			cobra.CheckErr(err)
			for idx, item := range ret {
				if item.Code != 200 {
					println(args[idx], "复制失败")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")
	mvCmd.Flags().BoolP("force", "f", false, "覆盖已存在的文件")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
