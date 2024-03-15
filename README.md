# xd

[![build-img]][build-url]
[![reportcard-img]][reportcard-url]
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
$ xd -n=42 -s=12345 -f=xd
00000000  1c 2c 6e 8c 01 1f 4e 88  85 a8 4e 1e 1d 24 6e c4  |.,n...N...N..$n.|
00000010  57 27 4f c4 47 27 6f a1  84 a1 4e 2d 1c 2d 6e ad  |W'O.G'o...N-.-n.|
00000020  09 60 6e a9 85 a9 4e 3e  1d 25 % 

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
[reportcard-img]: https://goreportcard.com/badge/cristaloleg/xd
[reportcard-url]: https://goreportcard.com/report/cristaloleg/xd
[version-img]: https://img.shields.io/github/v/release/cristaloleg/xd
[version-url]: https://github.com/cristaloleg/xd/releases
