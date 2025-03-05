# binspect

a tool for quickly viewing what encapsulates every occurrence of a specific string in a file

this is basically just [bgrep](https://github.com/tmbinc/bgrep/tree/master) with an output that is simpler to read 

## usage
```
binspect -file <infile> -target <target-string> [options]
```

the options are :
```
-b int
    number of bytes to be read before the target string (default 8)
-a int
    number of bytes to be read starting from the target string (default 8)
-B string
    output format for the bytes that come before the target string: hex or char or mixed (default hex)
-A string
    output format for the bytes that start from the target string: hex or char or mixed (default hex)
-O string
    output format for the offset of each occurrence: hex or decimal (default hex)
-order string
    the ordering of each row: any permutation of "boa" (default boa)
```
results are given in TSV format to stdout, by default in the order `before`-`offset`-`after`, e.g:
```
$ ./binspect -file ./binspect -target "hi"
00 80 39 09 FD 63 D3 C6 0001B735        68 69 38 09 89 60 D3 EA
03 00 F9 E9 1B 40 F9 CB 00051ECD        68 69 F8 2B 07 00 F9 C8
01 07 8B 69 01 7D D3 EA 0005AC35        68 69 F8 4A 05 00 91 EA
```


## features
what makes this different from bgrep ?? 

### output modes
in bgrep, if a byte falls within the printable ascii range, it will be displayed as such. otherwise, it will be displayed as hex. i find this hard to read:
```
$ ./bgrep -B 8 -A 8 6869 ./binspect
byteinspect/binspect: 0001b735
\x00\x809\x09\xfdc\xd3\xc6hi8\x09\x89`\xd3\xea
byteinspect/binspect: 00051ecd
\x03\x00\xf9\xe9\x1b@\xf9\xcbhi\xf8+\x07\x00\xf9\xc8
byteinspect/binspect: 0005ac35
\x01\x07\x8bi\x01}\xd3\xeahi\xf8J\x05\x00\x91\xea
```
so, binspect displays all output as hex by default,
```
$ ./binspect -file ./binspect -target "hi" -b 8 -a 8
00 80 39 09 FD 63 D3 C6 0001B735        68 69 38 09 89 60 D3 EA
03 00 F9 E9 1B 40 F9 CB 00051ECD        68 69 F8 2B 07 00 F9 C8
01 07 8B 69 01 7D D3 EA 0005AC35        68 69 F8 4A 05 00 91 EA
```
and lets u choose how u want the bytes to be displayed:
```
$ ./binspect -file ./binspect -target "hi" -b 8 -a 8 -B hex -A char
00 80 39 09 FD 63 D3 C6 0001B735        hi8..`..
03 00 F9 E9 1B 40 F9 CB 00051ECD        hi.+....
01 07 8B 69 01 7D D3 EA 0005AC35        hi.J....
```
output mode `mixed` will give a similar result to bgrep:
```
$ ./binspect -file ./binspect -target "hi" -B mixed -A mixed
\x00\x809\x09\xFDc\xD3\xC6      0001B735        hi8\x09\x89`\xD3\xEA
\x03\x00\xF9\xE9\x1B@\xF9\xCB   00051ECD        hi\xF8+\x07\x00\xF9\xC8
\x01\x07\x8Bi\x01}\xD3\xEA      0005AC35        hi\xF8J\x05\x00\x91\xEA
```
### offset lists
if u only want a list of all the offsets, bgrep lets u set the before/after bytes to 0:
```
$ bgrep -B 0 -A 0 6869 ./binspect
./binspect: 0001b735
./binspect: 00051ecd
./binspect: 0005ac35
```
binspect does the same thing but it doesnt include the filename at the start of each line , which makes the list easier to feed into other programs :
```
$ ./binspect -file ./binspect -target "hi" -b 0 -a 0
0001B735
00051ECD
0005AC35
```

### other options

change the output mode for the offsets with `-O`
```
$ ./binspect -file ./binspect -target "hi" -O decimal
00 80 39 09 FD 63 D3 C6 112437  68 69 38 09 89 60 D3 EA
03 00 F9 E9 1B 40 F9 CB 335565  68 69 F8 2B 07 00 F9 C8
01 07 8B 69 01 7D D3 EA 371765  68 69 F8 4A 05 00 91 EA
```

control the order of each row by using `-order`
```
$ ./binspect -file ./binspect -target "hi" -order bao
00 80 39 09 FD 63 D3 C6 68 69 38 09 89 60 D3 EA0001B735
03 00 F9 E9 1B 40 F9 CB 68 69 F8 2B 07 00 F9 C800051ECD
01 07 8B 69 01 7D D3 EA 68 69 F8 4A 05 00 91 EA0005AC35
```
combining all the options can give a completely different output
```
$ ./binspect -file ./binspect -target "hi" -B mixed -A mixed -O decimal -order oba 
112437  \x00\x809\x09\xFDc\xD3\xC6      hi8\x09\x89`\xD3\xEA
335565  \x03\x00\xF9\xE9\x1B@\xF9\xCB   hi\xF8+\x07\x00\xF9\xC8
371765  \x01\x07\x8Bi\x01}\xD3\xEA      hi\xF8J\x05\x00\x91\xEA
```
## install
needs go to build

clone to ur machine then build with `go build -o binspect main.go`

