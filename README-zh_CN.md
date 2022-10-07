# hash-rename

[![release](https://img.shields.io/github/v/release/ciphersaw/hash-rename)](https://github.com/ciphersaw/hash-rename) [![go](https://img.shields.io/badge/go-1.19.1-blue)](https://golang.org/)

[English](README.md) | 简体中文

hash-rename 作为一个常用工具，能将文件用其哈希值重命名，并保留原有后缀名不变。

## 使用说明

以下用法在 Linux 系统下演示，在 Windows 等其他系统下用法相同，也可作参考。

### 获取帮助

输入 `./hash-rename --help` 命令获取帮助：

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename --help
Usages:
hash-rename <-f /path/to/file> [-h hash_func]
hash-rename <-d /path/to/dir -s suffix1,suffix2,...> [-h hash_func]
  -d, --dir string      Set the directory path including files to be renamed.
  -f, --file string     Set the file path to be renamed.
  -h, --hash string     Set the hash function for renaming.
                        Currently available for md5, sha1, sha256. (default "md5")
      --help            Print the usage of hash-rename.
  -s, --suffix string   Set the suffixes of files.
  -v, --version         Print the version of hash-rename.
```

### 获取版本号

输入 `./hash-rename --version` 命令获取版本号：

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename --version
v1.0.0
```

### 重命名单个文件

使用 `-f, --file` 将一个系统文件重命名为其 md5 哈希值：

注意到此系统文件没有后缀名。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./passwd 
Result of renameOneFile:
[*] passwd --> bf52fc29f3fd754693ce4a6ff11575e7
```

使用 `-f, --file` 将一个 jpg 图片文件重命名为其 sha1 哈希值：

注意到可使用 `-h, --hash` 指定一种哈希算法用于重命名。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./test01.jpg -h sha1
Result of renameOneFile:
[*] test01.jpg --> 440a91d6fabaf5f8865faf97dfe574345b37a5a7.jpg
```

### 重命名目录下多个文件

使用 `-d, --dir` 将目录下的所有 jpg 与 png 图片文件，重命名为各自 sha256 哈希值：

注意到必须使用 `-s, --suffix`  指定需要重命名的文件的后缀名。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s jpg,png -h sha256
Result of renameBulkFiles:
[1] 440a91d6fabaf5f8865faf97dfe574345b37a5a7.jpg --> 71bfe668469aa882c3422100b1cbd89c4b83dbce9ea279854966e8ef084ffe0e.jpg
[2] test02.jpg --> 0014e9a4bb731b6060e9476dd6ad25f8423fd27451fa9d5c1ef1a9cec1bd45e8.jpg
[3] test03.png --> 428e6d35fe78cab5c792657088a124d91076e97b9cad5036b46698ea7341985e.png
[4] test04.png --> d8429ab7f39582146710a8afbc7d5bbe8adc0f9c7ee16b6e50c8738d0caafcf9.png
```

其中一种特殊的后缀名 `null/none`，用于重命名没有后缀名的文件：

注意到此处  hash-rename 工具本身也没有后缀名，所以它也被重命名为其 sha1 哈希值。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s null -h sha1     
Result of renameBulkFiles:
[1] bf52fc29f3fd754693ce4a6ff11575e7 --> 9ee17a7aa5a9cfb91dfc27a13a3f29732dd1f051
[2] hash-rename --> c2eccfd7430d8e8b37272b7cfb5f75ccbff41056
[3] zsh --> 5a2e990f3ae4ca940f9078826708ec9bdd273baf
```

另一种特殊的后缀名 `all`，用于重命名目录下的所有文件，无论是否有后缀名：

注意到 /tmp/test 目录下的所有文件，都被重命名为各自的 md5 哈希值。

```bash
┌──(root💀kali)-[/tmp/test]
└─# mv ./c2eccfd7430d8e8b37272b7cfb5f75ccbff41056 /tmp/hash-rename

┌──(root💀kali)-[/tmp/test]
└─# /tmp/hash-rename -d /tmp/test -s all -h md5
Result of renameBulkFiles:
[1] 0014e9a4bb731b6060e9476dd6ad25f8423fd27451fa9d5c1ef1a9cec1bd45e8.jpg --> 200852747245ddc1a9282a8006c72068.jpg
[2] 428e6d35fe78cab5c792657088a124d91076e97b9cad5036b46698ea7341985e.png --> 50197874009730f5a5d366baf52ed102.png
[3] 5a2e990f3ae4ca940f9078826708ec9bdd273baf --> f7889fc1a97bb6786b79ceb63d9c6ca4
[4] 71bfe668469aa882c3422100b1cbd89c4b83dbce9ea279854966e8ef084ffe0e.jpg --> bcc60e314d22ac5048299327c54d5e83.jpg
[5] 9ee17a7aa5a9cfb91dfc27a13a3f29732dd1f051 --> bf52fc29f3fd754693ce4a6ff11575e7
[6] d8429ab7f39582146710a8afbc7d5bbe8adc0f9c7ee16b6e50c8738d0caafcf9.png --> 80dabfe444567e35ee03d8c053b54d71.png
```