# 一些常用命令

* gcc --version
* gdb --version
* gdb
* dstat
* pidstat -w
* htop
* objdump
* readelf

[源码地址下载](https://golang.org/dl/) 注意：需要翻墙

推荐自己去下载源码进行编译，go支持自举，可以使用低版本编译高版本。

## add.c

```
//gcc -O0 -o test1 add.c //gcc 的开关 把优化都关掉
//gcc -O2 -o test2 add.c // O1 O2 O3 各种优化级别
//objdump -d -M intel test1
//objdump -d -M intel test2
```


`gcc -c teest.o`继续生成 .o 文件


`readelf -r test.o`
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# readelf -r test.o

Relocation section '.rela.text' at offset 0x268 contains 3 entries:
  Offset          Info           Type           Sym. Value    Sym. Name + Addend
00000000003c  000900000002 R_X86_64_PC32     0000000000000000 add - 4
000000000049  00050000000a R_X86_64_32       0000000000000000 .rodata + 0
000000000053  000b00000002 R_X86_64_PC32     0000000000000000 printf - 4

Relocation section '.rela.eh_frame' at offset 0x2b0 contains 2 entries:
  Offset          Info           Type           Sym. Value    Sym. Name + Addend
000000000020  000200000002 R_X86_64_PC32     0000000000000000 .text + 0
000000000040  000200000002 R_X86_64_PC32     0000000000000000 .text + 14
```


`objdump -S test.o`
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# objdump -S test.o

test.o:     file format elf64-x86-64


Disassembly of section .text:

0000000000000000 <add>:
   0:	55                   	push   %rbp
   1:	48 89 e5             	mov    %rsp,%rbp
   4:	89 7d fc             	mov    %edi,-0x4(%rbp)
   7:	89 75 f8             	mov    %esi,-0x8(%rbp)
   a:	8b 55 fc             	mov    -0x4(%rbp),%edx
   d:	8b 45 f8             	mov    -0x8(%rbp),%eax
  10:	01 d0                	add    %edx,%eax
  12:	5d                   	pop    %rbp
  13:	c3                   	retq

0000000000000014 <main>:
  14:	55                   	push   %rbp
  15:	48 89 e5             	mov    %rsp,%rbp
  18:	48 83 ec 20          	sub    $0x20,%rsp
  1c:	89 7d ec             	mov    %edi,-0x14(%rbp)
  1f:	48 89 75 e0          	mov    %rsi,-0x20(%rbp)
  23:	c7 45 f4 00 01 00 00 	movl   $0x100,-0xc(%rbp)
  2a:	c7 45 f8 00 02 00 00 	movl   $0x200,-0x8(%rbp)
  31:	8b 55 f8             	mov    -0x8(%rbp),%edx
  34:	8b 45 f4             	mov    -0xc(%rbp),%eax
  37:	89 d6                	mov    %edx,%esi
  39:	89 c7                	mov    %eax,%edi
  3b:	e8 00 00 00 00       	callq  40 <main+0x2c>
  40:	89 45 fc             	mov    %eax,-0x4(%rbp)
  43:	8b 45 fc             	mov    -0x4(%rbp),%eax
  46:	89 c6                	mov    %eax,%esi
  48:	bf 00 00 00 00       	mov    $0x0,%edi
  4d:	b8 00 00 00 00       	mov    $0x0,%eax
  52:	e8 00 00 00 00       	callq  57 <main+0x43>
  57:	b8 00 00 00 00       	mov    $0x0,%eax
  5c:	c9                   	leaveq
  5d:	c3                   	retq
```

## add.go

`go build -gcflags "-m" -o test add.go ` 输出优化器可以做的优化的策略

```
 wolf4j@xiaolong  ~/go/src/golang/yuheng/day01   master ●✚  go build -gcflags "-m" -o test add.go
# command-line-arguments
./add.go:3:6: can inline add
./add.go:7:6: can inline main
./add.go:11:5: inlining call to add

```

`go tool objdump -s "main\.main" test`

```
 wolf4j@xiaolong  ~/go/src/golang/yuheng/day01   master ●✚ go tool objdump -s "main\.main" test
TEXT main.main(SB) /Users/wolf4j/go/src/golang/yuheng/day01/add.go
  add.go:7              0x104f360               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  add.go:7              0x104f369               483b6110                CMPQ 0x10(CX), SP
  add.go:7              0x104f36d               7634                    JBE 0x104f3a3
  add.go:7              0x104f36f               4883ec10                SUBQ $0x10, SP
  add.go:7              0x104f373               48896c2408              MOVQ BP, 0x8(SP)
  add.go:7              0x104f378               488d6c2408              LEAQ 0x8(SP), BP
  add.go:12             0x104f37d               e82e3ffdff              CALL runtime.printlock(SB)
  add.go:12             0x104f382               48c7042400030000        MOVQ $0x300, 0(SP)
  add.go:12             0x104f38a               e80147fdff              CALL runtime.printint(SB)
  add.go:12             0x104f38f               e8cc41fdff              CALL runtime.printnl(SB)
  add.go:12             0x104f394               e8a73ffdff              CALL runtime.printunlock(SB)
  add.go:13             0x104f399               488b6c2408              MOVQ 0x8(SP), BP
  add.go:13             0x104f39e               4883c410                ADDQ $0x10, SP
  add.go:13             0x104f3a2               c3                      RET
  add.go:7              0x104f3a3               e8c883ffff              CALL runtime.morestack_noctxt(SB)
  add.go:7              0x104f3a8               ebb6                    JMP main.main(SB)
  :-1                   0x104f3aa               cc                      INT $0x3
  :-1                   0x104f3ab               cc                      INT $0x3
  :-1                   0x104f3ac               cc                      INT $0x3
  :-1                   0x104f3ad               cc                      INT $0x3
  :-1                   0x104f3ae               cc                      INT $0x3
  :-1                   0x104f3af               cc                      INT $0x3
```


`go build -x -o add add.go`
```
 wolf4j@xiaolong  ~/go/src/golang/yuheng/day01   master ●✚  go build -x -o add add.go
WORK=/var/folders/wk/y6tt5z4d0236rbrh14vpmqz40000gp/T/go-build310512725
mkdir -p $WORK/command-line-arguments/_obj/
mkdir -p $WORK/command-line-arguments/_obj/exe/
cd /Users/wolf4j/go/src/golang/yuheng/day01
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/command-line-arguments.a -trimpath $WORK -goversion go1.9.2 -p main -complete -buildid 490ab27c53105f459b886069b1de57e37fadf810 -D _/Users/wolf4j/go/src/golang/yuheng/day01 -I $WORK -pack ./add.go
cd .
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/command-line-arguments/_obj/exe/a.out -L $WORK -extld=clang -buildmode=exe -buildid=490ab27c53105f459b886069b1de57e37fadf810 $WORK/command-line-arguments.a
mv $WORK/command-line-arguments/_obj/exe/a.out add
```