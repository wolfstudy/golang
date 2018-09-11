run: mem.go
    go build -o test $^ && ./test
stat:
    pidstat -r -C test 1
clean:
    -rm test
