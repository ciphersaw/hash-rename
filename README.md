# hash-rename

[![release](https://img.shields.io/github/v/release/ciphersaw/hash-rename)](https://github.com/ciphersaw/hash-rename) [![go](https://img.shields.io/badge/go-1.19.1-blue)](https://golang.org/)

English | [简体中文](README-zh_CN.md)

The hash-rename is a common tool to rename file with its hash value, whose suffix would be kept unchanged.

## Usage

The usages below are demonstrated in Linux, as the same as in Windows and macOS.

### Help

Input `./hash-rename --help` to get usages:

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

### Version

Input `./hash-rename --version` to get the current version:

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename --version
v1.1.0
```

### Rename One File

Use `-f, --file` to rename a system file with its MD5 lowercase value:

Note that the system file has no suffix.

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./passwd 
Result of renameOneFile:
[*] passwd --> bf52fc29f3fd754693ce4a6ff11575e7
```

Use `-f, --file` to rename a jpg file with its SHA1 uppercase value:

Note that `-h, --hash` can be used to specify a hash function (default "md5"), and `-u, --uppercase` to set the uppercase of hash value for renaming.

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f ./test01.jpg -h sha1 -u
Result of renameOneFile:
[*] test01.jpg --> 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg
```

In order to avoid repetitive work, file would be not renamed if its current name matches the corresponding hash value.

Try to rename 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg with its SHA1 uppercase value again, and get the prompt of no need to rename again:

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg -h sha1 -u
Result of renameOneFile:
[-] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg has already been renamed with sha1 value, no need to rename again.
```

Nevertheless, `-F, --force` can be used to ignore file name check and rename forcibly:

```bash
──(root💀kali)-[/tmp/test]
└─# ./hash-rename -f 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg -h sha1 -u -F
Result of renameOneFile:
[*] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg --> 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg
```

### Rename Bulk Files

Use `-d, --dir` to rename a bulk of jpg and png files with their respective SHA1 lowercase value, and set 10 goroutines concurrency:

Note that `-s, --suffix` must be used to specify the suffixes that the files have, and `-c, --concurrency` can be used to set the concurrency for renaming files (default 4).

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s jpg,png -h sha1 -c 10
Result of renameBulkFiles:
[1] 440A91D6FABAF5F8865FAF97DFE574345B37A5A7.jpg --> 440a91d6fabaf5f8865faf97dfe574345b37a5a7.jpg
[2] test02.jpg --> 6705507d67c4d56eb3d273e03e0952f5daa2aea9.jpg
[3] test03.png --> e2ef055966ef72dfeb6bb3a8d6dd0b6746166055.png
[4] test04.png --> 5aa00caa44b1dfc9f0d341825b04bb2a006d8976.png
```

One special type of suffix is `null/none` that only renames the files without any suffix:

Note that hash-rename itself is also renamed with its SHA256 uppercase value, because of no suffix.

```bash
┌──(root💀kali)-[/tmp/test]
└─# ./hash-rename -d ./ -s null -h sha256 -u
Result of renameBulkFiles:
[1] bf52fc29f3fd754693ce4a6ff11575e7 --> E56B457E3F3B8104DDEAB52028E934863C2A28E49EEAA557EA68F274E2893BC2
[2] zsh --> C3F5891EC3CAB3D0534BFCB3CFB44B224236C8100459704CB8AE0388229DFBE5
[3] hash-rename --> 15144B2E8ED998AB4E1813925AFD56CF114D2828FC34D7519BC6DFF23256AE15
```

The other special type of suffix is `all` that renames all files ignoring suffix:

Note that all files in /tmp/test are renamed with their respective MD5 lowercase value.

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

Try to rename all files in /tmp/test with their respective MD5 lowercase value again, and get the prompt of possible reasons including no need to rename again:

Note that the other possible reasons also include suffixes mismatching, and errors in getting file hash or renaming file.

```bash
┌──(root💀kali)-[/tmp/test]
└─# /tmp/hash-rename -d /tmp/test -s all -h md5
Result of renameBulkFiles:
[-] No files have been renamed, and the possible reasons are as follows:
 1. The suffixes you specify do not match any files.
 2. The files in /tmp/test have already been renamed with md5 value, no need to rename again.
 3. Errors happen in getting file hash or renaming file with its hash value.
```

As the same above, `-F, --force` can also be used to rename all files in /tmp/test forcibly:

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

