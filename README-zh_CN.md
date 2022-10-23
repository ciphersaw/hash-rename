# hash-rename

[![release](https://img.shields.io/github/v/release/ciphersaw/hash-rename)](https://github.com/ciphersaw/hash-rename) [![go](https://img.shields.io/badge/go-1.19.1-blue)](https://golang.org/)

[English](README.md) | 简体中文

hash-rename 作为一个常用工具，能将文件用其哈希值重命名，并保留原有后缀名不变。

## 使用说明

以下用法在 Linux 系统下演示，在 Windows 与 macOS 等其他系统下用法相同，也可作参考。

### 获取帮助

输入 `./hash-rename --help` 命令获取帮助：

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename --help
Usages:
hash-rename <-f /path/to/file> [-h hash_func] [-u] [-F]
hash-rename <-d /path/to/dir -s suffix1,suffix2,...> [-h hash_func] [-c num] [-u] [-F]
  -c, --concurrency uint8   Set the goroutine concurrency for renaming files. (default 4)
  -d, --dir string          Set the directory path including files to be renamed.
  -f, --file string         Set the file path to be renamed.
  -F, --force               Force to rename ignoring file name check.
  -h, --hash string         Set the hash function for renaming.
                            Currently available for md5, sha1, sha256. (default "md5")
      --help                Print the usage of hash-rename.
  -s, --suffix string       Set the suffixes of files.
  -u, --uppercase           Set the uppercase of hash value for renaming.
  -v, --version             Print the version of hash-rename.
```

### 获取版本号

输入 `./hash-rename --version` 命令获取版本号：

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename --version
v1.1.0
```

### 重命名单个文件

使用 `-f, --file` 将一个系统文件重命名为其 MD5 小写哈希值：

注意到此系统文件没有后缀名。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./passwd 
Result of renameOneFile:
[*] passwd --> bf52fc29f3fd754693ce4a6ff11575e7
```

使用 `-f, --file` 将一个 jpg 图片文件重命名为其 SHA1 大写哈希值：

注意到 `-h, --hash` 可指定一种哈希算法用于重命名（默认值为 md5），而 `-u, --uppercase` 可指定用大写哈希值重命名。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./test01.jpg -h sha1 -u
Result of renameOneFile:
[*] test01.jpg --> 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg
```

为了避免重复工作，若当前文件名与指定对应的哈希值一致，则不会进行重命名。

尝试将 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg 文件再次重命名为其 SHA1 大写哈希值，会提示文件无需再次重命名：

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg -h sha1 -u
Result of renameOneFile:
[-] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg has already been renamed with sha1 value, no need to rename again.
```

尽管如此，可使用 `-F, --force` 来忽略文件名检查，并进行强制重命名：

```bash
──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg -h sha1 -u -F
Result of renameOneFile:
[*] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg --> 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg
```

### 重命名目录下多个文件

使用 `-d, --dir` 将目录下的所有 jpg 与 png 图片文件，重命名为各自的 SHA1 小写哈希值，并将 Go 协程并发数设置为 10：

注意到必须使用 `-s, --suffix`  指定需要重命名的文件的后缀名，而可使用 `-c, --concurrency` 设置文件重命名并发数（默认值为 4）。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s jpg,png -h sha1 -c 10
Result of renameBulkFiles:
[1] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg --> 440a91d6fabaf5f8865faf97dfe574345b37a5a7.jpg
[2] test02.jpg --> 6705507d67c4d56eb3d273e03e0952f5daa2aea9.jpg
[3] test03.png --> e2ef055966ef72dfeb6bb3a8d6dd0b6746166055.png
[4] test04.png --> 5aa00caa44b1dfc9f0d341825b04bb2a006d8976.png
```

其中一种特殊的后缀名 `null/none`，用于重命名没有后缀名的文件：

注意到此处  hash-rename 工具本身也没有后缀名，所以它也被重命名为其 SHA256 大写哈希值。

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s null -h sha256 -u
Result of renameBulkFiles:
[1] bf52fc29f3fd754693ce4a6ff11575e7 --> E56B457E3F3B8104DDEAB52028E934863C2A28E49EEAA557EA68F274E2893BC2
[2] zsh --> C3F5891EC3CAB3D0534BFCB3CFB44B224236C8100459704CB8AE0388229DFBE5
[3] hash-rename --> 15144B2E8ED998AB4E1813925AFD56CF114D2828FC34D7519BC6DFF23256AE15
```

另一种特殊的后缀名 `all`，用于重命名目录下的所有文件，无论是否有后缀名：

注意到 /tmp/test 目录下的所有文件，都被重命名为各自的 MD5 小写哈希值。

```bash
┌──(root💀kali)-[/tmp/test]
└─# mv ./15144B2E8ED998AB4E1813925AFD56CF114D2828FC34D7519BC6DFF23256AE15 /tmp/hash-rename

┌──(root💀kali)-[/tmp/test]
└─# /tmp/hash-rename -d /tmp/test -s all -h md5
Result of renameBulkFiles:
[1] 440a91d6fabaf5f8865faf97dfe574345b37a5a7.jpg --> bcc60e314d22ac5048299327c54d5e83.jpg
[2] 5aa00caa44b1dfc9f0d341825b04bb2a006d8976.png --> 80dabfe444567e35ee03d8c053b54d71.png
[3] 6705507d67c4d56eb3d273e03e0952f5daa2aea9.jpg --> 200852747245ddc1a9282a8006c72068.jpg
[4] e2ef055966ef72dfeb6bb3a8d6dd0b6746166055.png --> 50197874009730f5a5d366baf52ed102.png
[5] E56B457E3F3B8104DDEAB52028E934863C2A28E49EEAA557EA68F274E2893BC2 --> bf52fc29f3fd754693ce4a6ff11575e7
[6] C3F5891EC3CAB3D0534BFCB3CFB44B224236C8100459704CB8AE0388229DFBE5 --> f7889fc1a97bb6786b79ceb63d9c6ca4
```

尝试将 /tmp/test 目录下的所有文件再次重命名为其 MD5 小写哈希值，会提示未进行重命名可能有哪些原因，其中包括文件无需再次重命名：

注意到其他可能的原因，还包括文件后缀未匹配，以及计算哈希值或文件重命名过程中发生的错误等。

```bash
┌──(root💀kali)-[/tmp/test]
└─# /tmp/hash-rename -d /tmp/test -s all -h md5
Result of renameBulkFiles:
[-] No files have been renamed, and the possible reasons are as follows:
 1. The suffixes you specify do not match any files.
 2. The files in /tmp/test have already been renamed with md5 value, no need to rename again.
 3. Errors happen in getting file hash or renaming file with its hash value.
```

同上，使用 `-F, --force` 依旧能够对 /tmp/test 目录下的所有文件，进行强制重命名:

```bash
┌──(root💀kali)-[/tmp/test]
└─# /tmp/hash-rename -d /tmp/test -s all -h md5 -F
Result of renameBulkFiles:
[1] 200852747245ddc1a9282a8006c72068.jpg --> 200852747245ddc1a9282a8006c72068.jpg
[2] 50197874009730f5a5d366baf52ed102.png --> 50197874009730f5a5d366baf52ed102.png
[3] bcc60e314d22ac5048299327c54d5e83.jpg --> bcc60e314d22ac5048299327c54d5e83.jpg
[4] 80dabfe444567e35ee03d8c053b54d71.png --> 80dabfe444567e35ee03d8c053b54d71.png
[5] bf52fc29f3fd754693ce4a6ff11575e7 --> bf52fc29f3fd754693ce4a6ff11575e7
[6] f7889fc1a97bb6786b79ceb63d9c6ca4 --> f7889fc1a97bb6786b79ceb63d9c6ca4
```