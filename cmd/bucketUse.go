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
	"fmt"

	"github.com/nnnewb/qn/internal/utils"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "使用指定的bucket",
	Long:  `使用指定的bucket`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		credential := auth.New(viper.GetString("ak"), viper.GetString("sk"))

		rid, err := cmd.Flags().GetString("region")
		cobra.CheckErr(err)

		regionID, err := utils.CheckRegion(rid)
		cobra.CheckErr(err)

		region, _ := storage.GetRegionByID(regionID)
		mgr := storage.NewBucketManager(credential, &storage.Config{UseCdnDomains: false, Region: &region})

		buckets, err := mgr.Buckets(true)
		bucket := args[0]
		for _, b := range buckets {
			if b == bucket {
				viper.Set("bucket", bucket)
				err = viper.GetViper().WriteConfig()
				cobra.CheckErr(err)

				fmt.Println("bucket =>", bucket)
				return
			}
		}
		cobra.CheckErr(fmt.Errorf("bucket %s not exists", bucket))
	},
}

func init() {
	bucketCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
