# xd

[![build-img]][build-url]
[![version-img]][version-url]

xd is just hexdump with a few parameters.

## Features

* Simple
* Works with `stdin`
* Parameters:
  * `-n` how many bytes to output 
  * `-s` how many bytes to skip before processing
  * `-f` file to read input
  * `-o` file to write output

## Install

Go version 1.17+

```
go install github.com/cristaloleg/xd
```

## Example

```sh
$ xd -n 40 < xd
00000000  cf fa ed fe 0c 00 00 01  00 00 00 00 02 00 00 00  |................|
00000010  0f 00 00 00 a8 09 00 00  04 00 20 00 00 00 00 00  |.......... .....|
00000020  19 00 00 00 48 00 00 00  % 

$ xd -n=128 -s=256 -f=xd
00000000  5f 5f 73 79 6d 62 6f 6c  5f 73 74 75 62 31 00 00  |__symbol_stub1..|
00000010  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000020  a0 af 09 00 01 00 00 00  88 02 00 00 00 00 00 00  |................|
00000030  a0 af 09 00 05 00 00 00  00 00 00 00 00 00 00 00  |................|
00000040  08 04 00 80 00 00 00 00  06 00 00 00 00 00 00 00  |................|
00000050  5f 5f 72 6f 64 61 74 61  00 00 00 00 00 00 00 00  |__rodata........|
00000060  5f 5f 54 45 58 54 00 00  00 00 00 00 00 00 00 00  |__TEXT..........|
00000070  40 b2 09 00 01 00 00 00  a8 4c 02 00 00 00 00 00  |@........L......|
```

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristaloleg/xd/workflows/build/badge.svg
[build-url]: https://github.com/cristaloleg/xd/actions
[version-img]: https://img.shields.io/github/v/release/cristaloleg/xd
[version-url]: https://github.com/cristaloleg/xd/releases
