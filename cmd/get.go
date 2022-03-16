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
	"time"

	"github.com/nnnewb/qn/internal/utils"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "下载指定的文件到本地",
	Long:  "下载指定的文件到本地",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(checkConfig())

		credential := auth.New(viper.GetString("ak"), viper.GetString("sk"))
		mgr := storage.NewBucketManager(credential, &storage.Config{UseCdnDomains: false})

		info, err := mgr.GetBucketInfo(viper.GetString("bucket"))
		cobra.CheckErr(err)

		domain, err := utils.GetBucketDownloadDomain(viper.GetString("bucket"), mgr)
		cobra.CheckErr(err)

		links := make([]string, 0, len(args))

		for _, key := range args {
			if info.IsPrivate() {
				links = append(links, storage.MakePrivateURLv2(credential, domain, key, int64(time.Hour)))
			} else {
				links = append(links, storage.MakePublicURLv2(domain, key))
			}
		}

		for _, link := range links {
			basename := path.Base(link)
			err = utils.DownloadWithProgress(link, basename)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
