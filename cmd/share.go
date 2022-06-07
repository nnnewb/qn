/*
Copyright © 2022 weak_ptr <weak_ptr@outlook.com>

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
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "获取文件下载链接",
	Long:  `获取文件下载链接`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(checkConfig())

		credential := auth.New(args[0], args[1])
		mgr := storage.NewBucketManager(credential, &storage.Config{UseCdnDomains: false})

		domain, err := cmd.Flags().GetString("domain")
		cobra.CheckErr(err)

		if domain == "" {
			domains, err := mgr.ListBucketDomains(viper.GetViper().GetString("bucket"))
			cobra.CheckErr(err)
			if len(domains) > 0 {
				domain = domains[0].Domain
			}
		}

		expire, err := cmd.Flags().GetDuration("duration")
		cobra.CheckErr(err)

		url := storage.MakePrivateURLv2(credential, domain, args[0], int64(expire.Seconds()))
		cmd.Println(url)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	shareCmd.Flags().StringP("domain", "d", "", "链接使用的域名")
	shareCmd.Flags().Duration("duration", time.Hour/2, "下载链接的签名有效期")
}
