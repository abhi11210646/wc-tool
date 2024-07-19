## Build Your Own wc Tool - Coding Challenge


Build the GO program first.

```console
go build -o ./wc ./cmd/main.go
```

Execute:

```console
abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc test.txt test2.txt
2  5  12   test.txt
6  8  18   test2222.txt
8  13 30   total

```
```console
abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc test.txt
2  5  12   test.txt

```

```console
abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc -h

        USAGE:  wc [OPTIONS] FILE(S)
                -c      print the byte counts
                -l      print the newline counts
                -w      print the word counts  
        EXAMPLE: 
                wc -l logs.txt
                wc -w logs.txt
                wc -l logs.txt logs2.txt
          

abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc test.txt test2.txt test.txt -l -w
 2  5 test.txt
 6  8 test2.txt
 2  5 test.txt
```
![wc-tool](https://github.com/abhi11210646/wc-tool/assets/16542492/a39a1ff9-40b5-4a06-a9be-32ee4a49d5be)

```console
abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc -h

        USAGE:  wc [OPTIONS] FILE(S)
                -c      print the byte counts
                -l      print the newline counts
                -w      print the word counts  
        EXAMPLE: 
                wc -l logs.txt
                wc -w logs.txt
                wc -l logs.txt logs2.txt
          

abhisekhs@abhisekhs-ThinkPad-L490:~/personal/Projects/wc-tool$ ./wc test.txt test2.txt test.txt -l -w
 2  5 test.txt
 6  8 test2.txt
 2  5 test.txt
```
