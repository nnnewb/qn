# qn

七牛云命令行工具。提供一些常用命令，方便和其他命令行工具结合实现更多功能。

## 例子

```
$ qn config --set ak=<access key>
$ qn config --set sk=<secret key>
$ qn config --set bucket=<bucket>

$ qn put Meslo.zip
$ qn ls Meslo
2021-08-25T15:35:32+08:00    63.8 MiB   Meslo.zip

$ qn stat Meslo.zip
key:      Meslo.zip
hash:     lndAhW4ceYL4BnfUnCT2dLYi5sce
mimetype: application/zip
size:     63.8 MiB
putTime:  2021-08-25T15:35:32+08:00   
type:     0
status:   0

$ qn mv Meslo.zip Meslo-font.zip
$ qn cp Meslo-font.zip Meslo.zip
$ qn ls Meslo
2021-08-25T15:39:13+08:00    63.8 MiB   Meslo-font.zip
2021-08-25T15:39:28+08:00    63.8 MiB   Meslo.zip

$ qn rm Meslo-font.zip
$ qn get welcome.jpg
4.21 KiB / 4.21 KiB [---------------------------------------] 100.00% ? p/s 0s
```

## 用法

```
七牛云命令行实用工具

Usage:
  qn [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  config      查看和修改设置
  cp          复制文件
  get         下载指定的文件到本地
  help        Help about any command
  ls          列出文件
  mv          移动文件
  put         上传文件到七牛云
  rm          删除文件
  stat        获取存储文件状态

Flags:
      --config string   config file (default is $HOME/.qnrc)
  -h, --help            help for qn
  -t, --toggle          Help message for toggle

Use "qn [command] --help" for more information about a command.
```
