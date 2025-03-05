# byteinspect

tool for quickly viewing what encapsulates every occurrence of a specific string in a file

## usage
```
byteinspect -infile <infile> -string <target-string> [options]
```

the options are :
```
- before int
    number of bytes to be read before each occurrence of target string (default 8)
- after int
    number of bytes to be read after each occurrence of target string (default 16)
- beforeFormat string
    output format for the bytes that precede the target string: hex or char (default hex)
- afterFormat string
    output format for the bytes that follow the target string: hex or char (default hex)
```

results are given in TSV format to stdout, e.g:
```
$ ./byteinspect -infile ./byteinspect -string "hi"
00 80 39 09 FD 63 D3 C6 0001B735        68 69 38 09 89 60 D3 EA 03 40 B2 49 21 C9 9A 26
03 00 F9 E9 1B 40 F9 CB 00051ECD        68 69 F8 2B 07 00 F9 C8 68 29 F8 61 08 40 F9 3F
01 07 8B 69 01 7D D3 EA 0005AC35        68 69 F8 4A 05 00 91 EA 68 29 F8 FF 7F 05 A9 FF
[...]
```

if you only want to get the offset of every occurrence and dont care about the surrounding data , u can just set `before` and `after` to 0, e.g.:
```
./byteinspect -infile ./byteinspect -string "hi" -before 0 -after 0
0001B735
00051ECD
0005AC35
```

## install
needs go to build

clone to ur machine then build with `go build -o byteinspect main.go`

