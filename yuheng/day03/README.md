## hello.c

```
#include <stdio.h>
#include <stdlib.h>

int x = 0x1234;
int y;
char *s = "x = %x, y = %x\n";

int main()
{
    printf(s,x,y);
    return 0;
}
```

`gcc -g -O0 -o hello_c hello.c` -g 表示生成调试字符信息 -O表示优化的第一个字母 -o输出的文件名字

## 可执行文件结构

`readelf -S hello_c` 对数据表的具体描述

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# readelf -S hello_c
There are 36 section headers, starting at offset 0x1d68:

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .interp           PROGBITS         0000000000400238  00000238
       000000000000001c  0000000000000000   A       0     0     1
  [ 2] .note.ABI-tag     NOTE             0000000000400254  00000254
       0000000000000020  0000000000000000   A       0     0     4
  [ 3] .note.gnu.build-i NOTE             0000000000400274  00000274
       0000000000000024  0000000000000000   A       0     0     4
  [ 4] .gnu.hash         GNU_HASH         0000000000400298  00000298
       000000000000001c  0000000000000000   A       5     0     8
  [ 5] .dynsym           DYNSYM           00000000004002b8  000002b8
       0000000000000060  0000000000000018   A       6     1     8
  [ 6] .dynstr           STRTAB           0000000000400318  00000318
       000000000000003f  0000000000000000   A       0     0     1
  [ 7] .gnu.version      VERSYM           0000000000400358  00000358
       0000000000000008  0000000000000002   A       5     0     2
  [ 8] .gnu.version_r    VERNEED          0000000000400360  00000360
       0000000000000020  0000000000000000   A       6     1     8
  [ 9] .rela.dyn         RELA             0000000000400380  00000380
       0000000000000018  0000000000000018   A       5     0     8
  [10] .rela.plt         RELA             0000000000400398  00000398
       0000000000000030  0000000000000018  AI       5    24     8
  [11] .init             PROGBITS         00000000004003c8  000003c8
       000000000000001a  0000000000000000  AX       0     0     4
  [12] .plt              PROGBITS         00000000004003f0  000003f0
       0000000000000030  0000000000000010  AX       0     0     16
  [13] .plt.got          PROGBITS         0000000000400420  00000420
       0000000000000008  0000000000000000  AX       0     0     8
  [14] .text             PROGBITS         0000000000400430  00000430
       00000000000001a2  0000000000000000  AX       0     0     16
  [15] .fini             PROGBITS         00000000004005d4  000005d4
       0000000000000009  0000000000000000  AX       0     0     4
  [16] .rodata           PROGBITS         00000000004005e0  000005e0
       0000000000000014  0000000000000000   A       0     0     4
  [17] .eh_frame_hdr     PROGBITS         00000000004005f4  000005f4
       0000000000000034  0000000000000000   A       0     0     4
  [18] .eh_frame         PROGBITS         0000000000400628  00000628
       00000000000000f4  0000000000000000   A       0     0     8
  [19] .init_array       INIT_ARRAY       0000000000600e10  00000e10
       0000000000000008  0000000000000000  WA       0     0     8
  [20] .fini_array       FINI_ARRAY       0000000000600e18  00000e18
       0000000000000008  0000000000000000  WA       0     0     8
  [21] .jcr              PROGBITS         0000000000600e20  00000e20
       0000000000000008  0000000000000000  WA       0     0     8
  [22] .dynamic          DYNAMIC          0000000000600e28  00000e28
       00000000000001d0  0000000000000010  WA       6     0     8
  [23] .got              PROGBITS         0000000000600ff8  00000ff8
       0000000000000008  0000000000000008  WA       0     0     8
  [24] .got.plt          PROGBITS         0000000000601000  00001000
       0000000000000028  0000000000000008  WA       0     0     8
  [25] .data             PROGBITS         0000000000601028  00001028
       0000000000000020  0000000000000000  WA       0     0     8
  [26] .bss              NOBITS           0000000000601048  00001048
       0000000000000008  0000000000000000  WA       0     0     4
  [27] .comment          PROGBITS         0000000000000000  00001048
       0000000000000034  0000000000000001  MS       0     0     1
  [28] .debug_aranges    PROGBITS         0000000000000000  0000107c
       0000000000000030  0000000000000000           0     0     1
  [29] .debug_info       PROGBITS         0000000000000000  000010ac
       00000000000000de  0000000000000000           0     0     1
  [30] .debug_abbrev     PROGBITS         0000000000000000  0000118a
       000000000000005c  0000000000000000           0     0     1
  [31] .debug_line       PROGBITS         0000000000000000  000011e6
       000000000000003e  0000000000000000           0     0     1
  [32] .debug_str        PROGBITS         0000000000000000  00001224
       00000000000000cd  0000000000000001  MS       0     0     1
  [33] .shstrtab         STRTAB           0000000000000000  00001c16
       000000000000014c  0000000000000000           0     0     1
  [34] .symtab           SYMTAB           0000000000000000  000012f8
       0000000000000708  0000000000000018          35    52     8
  [35] .strtab           STRTAB           0000000000000000  00001a00
       0000000000000216  0000000000000000           0     0     1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings), l (large)
  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)
  O (extra OS processing required) o (OS specific), p (processor specific)
```


数量名字什么的都不是固定的，对于各种编译器都是不一样的。对cpu没用，cpu只关心地址，从哪里开始执行，到哪里结束。

* .text 存储的全部是机器码，指令
* .data  存储全局变量，静态局部变量，值已经初始化好的
* .bss 没有初始化值的全局变量，映射地址信息，但不存储任何数据，否则他们之间的offset是不可能相等的。
* .rodata 存储一些只读数据

注意：

* 机器码和汇编的关系：机器码和汇编是一一对应的，从汇编到机器码是编译，从机器码到汇编属于反编译
* .bss文件段是存储的没有初始化的一些变量，它分为两个阶段，一个是ELF，一个是进程执行的时候，在ELF阶段是不需要填充其数据段的（也就是说，我初始化的时候
没有数据），但是我在运行时的时候，可能会被别人调用到，所以就需要为其分配内存空间。

在这里要区分指令和数据。指令和数据不是一回事。

`readelf -x .text hello_c`

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# readelf -x .text hello_c

Hex dump of section '.text':
  0x00400430 31ed4989 d15e4889 e24883e4 f0505449 1.I..^H..H...PTI
  0x00400440 c7c0d005 400048c7 c1600540 0048c7c7 ....@.H..`.@.H..
  0x00400450 26054000 e8b7ffff fff4660f 1f440000 &.@.......f..D..
  0x00400460 b84f1060 0055482d 48106000 4883f80e .O.`.UH-H.`.H...
  0x00400470 4889e576 1bb80000 00004885 c074115d H..v......H..t.]
  0x00400480 bf481060 00ffe066 0f1f8400 00000000 .H.`...f........
  0x00400490 5dc30f1f 4000662e 0f1f8400 00000000 ]...@.f.........
  0x004004a0 be481060 00554881 ee481060 0048c1fe .H.`.UH..H.`.H..
  0x004004b0 034889e5 4889f048 c1e83f48 01c648d1 .H..H..H..?H..H.
  0x004004c0 fe7415b8 00000000 4885c074 0b5dbf48 .t......H..t.].H
  0x004004d0 106000ff e00f1f00 5dc3660f 1f440000 .`......].f..D..
  0x004004e0 803d610b 20000075 11554889 e5e86eff .=a. ..u.UH...n.
  0x004004f0 ffff5dc6 054e0b20 0001f3c3 0f1f4000 ..]..N. ......@.
  0x00400500 bf200e60 0048833f 007505eb 930f1f00 . .`.H.?.u......
  0x00400510 b8000000 004885c0 74f15548 89e5ffd0 .....H..t.UH....
  0x00400520 5de97aff ffff5548 89e58b15 1c0b2000 ].z...UH...... .
  0x00400530 8b0d020b 2000488b 05030b20 0089ce48 .... .H.... ...H
  0x00400540 89c7b800 000000e8 b4feffff b8000000 ................
  0x00400550 005dc366 2e0f1f84 00000000 000f1f00 .].f............
  0x00400560 41574156 4189ff41 5541544c 8d259e08 AWAVA..AUATL.%..
  0x00400570 20005548 8d2d9e08 20005349 89f64989  .UH.-.. .SI..I.
  0x00400580 d54c29e5 4883ec08 48c1fd03 e837feff .L).H...H....7..
  0x00400590 ff4885ed 742031db 0f1f8400 00000000 .H..t 1.........
  0x004005a0 4c89ea4c 89f64489 ff41ff14 dc4883c3 L..L..D..A...H..
  0x004005b0 014839eb 75ea4883 c4085b5d 415c415d .H9.u.H...[]A\A]
  0x004005c0 415e415f c390662e 0f1f8400 00000000 A^A_..f.........
  0x004005d0 f3c3
```


`gdb hello_c` 启动gdb调试

```
Reading symbols from hello_c...done.
(gdb) l //查看代码信息
1	#include <stdio.h>
2	#include <stdlib.h>
3
4	int x = 0x1234;
5	int y;
6	char *s = "x = %x, y = %x\n";
7
8	int main()
9	{
10	    printf(s,x,y);
(gdb)
11	    return 0;
12	}
(gdb)
Line number 13 out of range; hello.c has 12 lines.
(gdb) b 12 //断点
Breakpoint 1 at 0x400551: file hello.c, line 12.
(gdb) r //执行
Starting program: /root/rxl/go/day03/hello_c
x = 1234, y = 0

Breakpoint 1, main () at hello.c:12
12	}
(gdb) p/x &x //取地址
$1 = 0x601038
(gdb) p/x &y
$2 = 0x60104c
(gdb) p/x &s
$3 = 0x601040
(gdb) x/xg &s //去获取char s的string地址信息
0x601040 <s>:	0x00000000004005e4
(gdb) info files //当前打开的可执行文件的一些信息（包括地址信息）
Symbols from "/root/rxl/go/day03/hello_c".
Native process:
	Using the running image of child process 9419.
	While running this, GDB does not access memory from...
Local exec file:
	`/root/rxl/go/day03/hello_c', file type elf64-x86-64.
	Entry point: 0x400430
	0x0000000000400238 - 0x0000000000400254 is .interp
	0x0000000000400254 - 0x0000000000400274 is .note.ABI-tag
	0x0000000000400274 - 0x0000000000400298 is .note.gnu.build-id
	0x0000000000400298 - 0x00000000004002b4 is .gnu.hash
	0x00000000004002b8 - 0x0000000000400318 is .dynsym
	0x0000000000400318 - 0x0000000000400357 is .dynstr
	0x0000000000400358 - 0x0000000000400360 is .gnu.version
	0x0000000000400360 - 0x0000000000400380 is .gnu.version_r
	0x0000000000400380 - 0x0000000000400398 is .rela.dyn
	0x0000000000400398 - 0x00000000004003c8 is .rela.plt
	0x00000000004003c8 - 0x00000000004003e2 is .init
	0x00000000004003f0 - 0x0000000000400420 is .plt
	0x0000000000400420 - 0x0000000000400428 is .plt.got
	0x0000000000400430 - 0x00000000004005d2 is .text
	0x00000000004005d4 - 0x00000000004005dd is .fini
	0x00000000004005e0 - 0x00000000004005f4 is .rodata //char 定义的s的string信息存储在这里。
	0x00000000004005f4 - 0x0000000000400628 is .eh_frame_hdr
	0x0000000000400628 - 0x000000000040071c is .eh_frame
	0x0000000000600e10 - 0x0000000000600e18 is .init_array
	0x0000000000600e18 - 0x0000000000600e20 is .fini_array
	0x0000000000600e20 - 0x0000000000600e28 is .jcr
	0x0000000000600e28 - 0x0000000000600ff8 is .dynamic
	0x0000000000600ff8 - 0x0000000000601000 is .got
	0x0000000000601000 - 0x0000000000601028 is .got.plt
	0x0000000000601028 - 0x0000000000601048 is .data //可以看到x的地址是在.data段中的。
	0x0000000000601048 - 0x0000000000601050 is .bss  // y 是没有初始化的变量，可以看到在.bss段中 ,这里存储的是char 那个指针s，后面的string信息是单独存储的。
```


`objcopy --add-section .xxx=hello.c hello_c` 自己添加数据段

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# objcopy --add-section .xxx=hello.c hello_c
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# readelf -S hello_c
There are 37 section headers, starting at offset 0x1e08:

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .interp           PROGBITS         0000000000400238  00000238
       000000000000001c  0000000000000000   A       0     0     1
  [ 2] .note.ABI-tag     NOTE             0000000000400254  00000254
       0000000000000020  0000000000000000   A       0     0     4
  [ 3] .note.gnu.build-i NOTE             0000000000400274  00000274
       0000000000000024  0000000000000000   A       0     0     4
  [ 4] .gnu.hash         GNU_HASH         0000000000400298  00000298
       000000000000001c  0000000000000000   A       5     0     8
  [ 5] .dynsym           DYNSYM           00000000004002b8  000002b8
       0000000000000060  0000000000000018   A       6     1     8
  [ 6] .dynstr           STRTAB           0000000000400318  00000318
       000000000000003f  0000000000000000   A       0     0     1
  [ 7] .gnu.version      VERSYM           0000000000400358  00000358
       0000000000000008  0000000000000002   A       5     0     2
  [ 8] .gnu.version_r    VERNEED          0000000000400360  00000360
       0000000000000020  0000000000000000   A       6     1     8
  [ 9] .rela.dyn         RELA             0000000000400380  00000380
       0000000000000018  0000000000000018   A       5     0     8
  [10] .rela.plt         RELA             0000000000400398  00000398
       0000000000000030  0000000000000018  AI       5    24     8
  [11] .init             PROGBITS         00000000004003c8  000003c8
       000000000000001a  0000000000000000  AX       0     0     4
  [12] .plt              PROGBITS         00000000004003f0  000003f0
       0000000000000030  0000000000000010  AX       0     0     16
  [13] .plt.got          PROGBITS         0000000000400420  00000420
       0000000000000008  0000000000000000  AX       0     0     8
  [14] .text             PROGBITS         0000000000400430  00000430
       00000000000001a2  0000000000000000  AX       0     0     16
  [15] .fini             PROGBITS         00000000004005d4  000005d4
       0000000000000009  0000000000000000  AX       0     0     4
  [16] .rodata           PROGBITS         00000000004005e0  000005e0
       0000000000000014  0000000000000000   A       0     0     4
  [17] .eh_frame_hdr     PROGBITS         00000000004005f4  000005f4
       0000000000000034  0000000000000000   A       0     0     4
  [18] .eh_frame         PROGBITS         0000000000400628  00000628
       00000000000000f4  0000000000000000   A       0     0     8
  [19] .init_array       INIT_ARRAY       0000000000600e10  00000e10
       0000000000000008  0000000000000000  WA       0     0     8
  [20] .fini_array       FINI_ARRAY       0000000000600e18  00000e18
       0000000000000008  0000000000000000  WA       0     0     8
  [21] .jcr              PROGBITS         0000000000600e20  00000e20
       0000000000000008  0000000000000000  WA       0     0     8
  [22] .dynamic          DYNAMIC          0000000000600e28  00000e28
       00000000000001d0  0000000000000010  WA       6     0     8
  [23] .got              PROGBITS         0000000000600ff8  00000ff8
       0000000000000008  0000000000000008  WA       0     0     8
  [24] .got.plt          PROGBITS         0000000000601000  00001000
       0000000000000028  0000000000000008  WA       0     0     8
  [25] .data             PROGBITS         0000000000601028  00001028
       0000000000000020  0000000000000000  WA       0     0     8
  [26] .bss              NOBITS           0000000000601048  00001048
       0000000000000008  0000000000000000  WA       0     0     4
  [27] .comment          PROGBITS         0000000000000000  00001048
       0000000000000034  0000000000000001  MS       0     0     1
  [28] .debug_aranges    PROGBITS         0000000000000000  0000107c
       0000000000000030  0000000000000000           0     0     1
  [29] .debug_info       PROGBITS         0000000000000000  000010ac
       00000000000000de  0000000000000000           0     0     1
  [30] .debug_abbrev     PROGBITS         0000000000000000  0000118a
       000000000000005c  0000000000000000           0     0     1
  [31] .debug_line       PROGBITS         0000000000000000  000011e6
       000000000000003e  0000000000000000           0     0     1
  [32] .debug_str        PROGBITS         0000000000000000  00001224
       00000000000000cd  0000000000000001  MS       0     0     1
  [33] .xxx              PROGBITS         0000000000000000  000012f1 //刚才加进去的数据段
       000000000000008e  0000000000000000           0     0     1
  [34] .shstrtab         STRTAB           0000000000000000  00001cb6
       0000000000000151  0000000000000000           0     0     1
  [35] .symtab           SYMTAB           0000000000000000  00001380
       0000000000000720  0000000000000018          36    53     8
  [36] .strtab           STRTAB           0000000000000000  00001aa0
       0000000000000216  0000000000000000           0     0     1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings), l (large)
  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)
  O (extra OS processing required) o (OS specific), p (processor specific)
```

`strip --remove-section=.xxx hello_c`


`readelf -x .xxx hello_c`
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# readelf -x .xxx hello_c

Hex dump of section '.xxx':
  0x00000000 23696e63 6c756465 203c7374 64696f2e #include <stdio.
  0x00000010 683e0a23 696e636c 75646520 3c737464 h>.#include <std
  0x00000020 6c69622e 683e0a0a 696e7420 78203d20 lib.h>..int x =
  0x00000030 30783132 33343b0a 696e7420 793b0a63 0x1234;.int y;.c
  0x00000040 68617220 2a73203d 20227820 3d202578 har *s = "x = %x
  0x00000050 2c207920 3d202578 5c6e223b 0a0a696e , y = %x\n";..in
  0x00000060 74206d61 696e2829 0a7b0a20 20202070 t main().{.    p
  0x00000070 72696e74 6628732c 782c7929 3b0a2020 rintf(s,x,y);.
  0x00000080 20207265 7475726e 20303b0a 7d0a       return 0;.}.
```

## hello.go


go 的data分为两个，有指针和没指针的。（.data 和 .noptrdata）

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# readelf -S hello_go
There are 23 section headers, starting at offset 0x1c8:

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .text             PROGBITS         0000000000401000  00001000
       000000000008224b  0000000000000000  AX       0     0     16
  [ 2] .rodata           PROGBITS         0000000000484000  00084000
       000000000004200f  0000000000000000   A       0     0     32
  [ 3] .shstrtab         STRTAB           0000000000000000  000c6020
       000000000000010b  0000000000000000           0     0     1
  [ 4] .typelink         PROGBITS         00000000004c6140  000c6140
       0000000000000b64  0000000000000000   A       0     0     32
  [ 5] .itablink         PROGBITS         00000000004c6ca8  000c6ca8
       0000000000000040  0000000000000000   A       0     0     8
  [ 6] .gosymtab         PROGBITS         00000000004c6ce8  000c6ce8
       0000000000000000  0000000000000000   A       0     0     1
  [ 7] .gopclntab        PROGBITS         00000000004c6d00  000c6d00
       000000000004e498  0000000000000000   A       0     0     32
  [ 8] .noptrdata        PROGBITS         0000000000516000  00116000
       000000000000cbdc  0000000000000000  WA       0     0     32
  [ 9] .data             PROGBITS         0000000000522be0  00122be0
       0000000000006b10  0000000000000000  WA       0     0     32
  [10] .bss              NOBITS           0000000000529700  00129700
       000000000001c688  0000000000000000  WA       0     0     32
  [11] .noptrbss         NOBITS           0000000000545da0  00145da0
       0000000000002738  0000000000000000  WA       0     0     32
  [12] .debug_abbrev     PROGBITS         0000000000549000  0012a000
       00000000000001b5  0000000000000000           0     0     1
  [13] .debug_line       PROGBITS         00000000005491b5  0012a1b5
       0000000000010658  0000000000000000           0     0     1
  [14] .debug_frame      PROGBITS         000000000055980d  0013a80d
       000000000001219c  0000000000000000           0     0     1
  [15] .debug_pubnames   PROGBITS         000000000056b9a9  0014c9a9
       0000000000007f8a  0000000000000000           0     0     1
  [16] .debug_pubtypes   PROGBITS         0000000000573933  00154933
       000000000000a509  0000000000000000           0     0     1
  [17] .debug_gdb_script PROGBITS         000000000057de3c  0015ee3c
       000000000000002a  0000000000000000           0     0     1
  [18] .debug_info       PROGBITS         000000000057de66  0015ee66
       0000000000064461  0000000000000000           0     0     1
  [19] .debug_ranges     PROGBITS         00000000005e22c7  001c32c7
       00000000000060b0  0000000000000000           0     0     1
  [20] .note.go.buildid  NOTE             0000000000400f9c  00000f9c
       0000000000000064  0000000000000000   A       0     0     4
  [21] .symtab           SYMTAB           0000000000000000  001ca000
       00000000000122a0  0000000000000018          22   102     8
  [22] .strtab           STRTAB           0000000000000000  001dc2a0
       0000000000012416  0000000000000000           0     0     1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings), l (large)
  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)
  O (extra OS processing required) o (OS specific), p (processor specific)
```


`go build -gcflags "-N -l" -o test test.go`
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# go build -gcflags "-N -l" -o test test.go
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# ./test
hello world 4660 0
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# cat test.go
package main

import "fmt"

var s = "hello world"
var x = 0x1234
var y int

func main(){
	fmt.Println(s,x,y)
}
```

启动gdb调试，go有package的概念，gdb调试的时候需要带上包名。

```
(gdb) b 10
Breakpoint 1 at 0x48322f: file /root/rxl/go/day03/test.go, line 10.
(gdb) r
Starting program: /root/rxl/go/day03/test
[New LWP 10136]
[New LWP 10137]
[New LWP 10138]
[New LWP 10139]
[New LWP 10140]

Thread 1 "test" hit Breakpoint 1, main.main () at /root/rxl/go/day03/test.go:10
10		fmt.Println(s,x,y)
(gdb) p/x main.s
$1 = 0x4b41ca "hello world"
(gdb) p/x main.x
$2 = 0x1234
(gdb) p/x main.y
$3 = 0x0
```

## 符号

`nm test`

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# nm test
0000000000601048 B __bss_start
0000000000601048 b completed.7585
0000000000601028 D __data_start
0000000000601028 W data_start
0000000000400460 t deregister_tm_clones
00000000004004e0 t __do_global_dtors_aux
0000000000600e18 t __do_global_dtors_aux_fini_array_entry
0000000000601030 D __dso_handle
0000000000600e28 d _DYNAMIC
0000000000601048 D _edata
0000000000601050 B _end
00000000004005d4 T _fini
0000000000400500 t frame_dummy
0000000000600e10 t __frame_dummy_init_array_entry
0000000000400718 r __FRAME_END__
0000000000601000 d _GLOBAL_OFFSET_TABLE_
                 w __gmon_start__
00000000004005f4 r __GNU_EH_FRAME_HDR
00000000004003c8 T _init
0000000000600e18 t __init_array_end
0000000000600e10 t __init_array_start
00000000004005e0 R _IO_stdin_used
                 w _ITM_deregisterTMCloneTable
                 w _ITM_registerTMCloneTable
0000000000600e20 d __JCR_END__
0000000000600e20 d __JCR_LIST__
                 w _Jv_RegisterClasses
00000000004005d0 T __libc_csu_fini
0000000000400560 T __libc_csu_init
                 U __libc_start_main@@GLIBC_2.2.5
0000000000400526 T main
                 U printf@@GLIBC_2.2.5//没有地址，因为是动态链接的。
00000000004004a0 t register_tm_clones
0000000000601040 D s
0000000000400430 T _start
0000000000601048 D __TMC_END__
0000000000601038 D x
000000000060104c B y
```

## 反汇编


`objdump -d -M intel test`

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day03# objdump -d -M intel test

test:     file format elf64-x86-64


Disassembly of section .init:

00000000004003c8 <_init>:
  4003c8:	48 83 ec 08          	sub    rsp,0x8
  4003cc:	48 8b 05 25 0c 20 00 	mov    rax,QWORD PTR [rip+0x200c25]        # 600ff8 <_DYNAMIC+0x1d0>
  4003d3:	48 85 c0             	test   rax,rax
  4003d6:	74 05                	je     4003dd <_init+0x15>
  4003d8:	e8 43 00 00 00       	call   400420 <__libc_start_main@plt+0x10>
  4003dd:	48 83 c4 08          	add    rsp,0x8
  4003e1:	c3                   	ret

Disassembly of section .plt:

00000000004003f0 <printf@plt-0x10>:
  4003f0:	ff 35 12 0c 20 00    	push   QWORD PTR [rip+0x200c12]        # 601008 <_GLOBAL_OFFSET_TABLE_+0x8>
  4003f6:	ff 25 14 0c 20 00    	jmp    QWORD PTR [rip+0x200c14]        # 601010 <_GLOBAL_OFFSET_TABLE_+0x10>
  4003fc:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]

0000000000400400 <printf@plt>:
  400400:	ff 25 12 0c 20 00    	jmp    QWORD PTR [rip+0x200c12]        # 601018 <_GLOBAL_OFFSET_TABLE_+0x18>
  400406:	68 00 00 00 00       	push   0x0
  40040b:	e9 e0 ff ff ff       	jmp    4003f0 <_init+0x28>

0000000000400410 <__libc_start_main@plt>:
  400410:	ff 25 0a 0c 20 00    	jmp    QWORD PTR [rip+0x200c0a]        # 601020 <_GLOBAL_OFFSET_TABLE_+0x20>
  400416:	68 01 00 00 00       	push   0x1
  40041b:	e9 d0 ff ff ff       	jmp    4003f0 <_init+0x28>

Disassembly of section .plt.got:

0000000000400420 <.plt.got>:
  400420:	ff 25 d2 0b 20 00    	jmp    QWORD PTR [rip+0x200bd2]        # 600ff8 <_DYNAMIC+0x1d0>
  400426:	66 90                	xchg   ax,ax

Disassembly of section .text:

0000000000400430 <_start>:
  400430:	31 ed                	xor    ebp,ebp
  400432:	49 89 d1             	mov    r9,rdx
  400435:	5e                   	pop    rsi
  400436:	48 89 e2             	mov    rdx,rsp
  400439:	48 83 e4 f0          	and    rsp,0xfffffffffffffff0
  40043d:	50                   	push   rax
  40043e:	54                   	push   rsp
  40043f:	49 c7 c0 d0 05 40 00 	mov    r8,0x4005d0
  400446:	48 c7 c1 60 05 40 00 	mov    rcx,0x400560
  40044d:	48 c7 c7 26 05 40 00 	mov    rdi,0x400526
  400454:	e8 b7 ff ff ff       	call   400410 <__libc_start_main@plt>
  400459:	f4                   	hlt
  40045a:	66 0f 1f 44 00 00    	nop    WORD PTR [rax+rax*1+0x0]

0000000000400460 <deregister_tm_clones>:
  400460:	b8 4f 10 60 00       	mov    eax,0x60104f
  400465:	55                   	push   rbp
  400466:	48 2d 48 10 60 00    	sub    rax,0x601048
  40046c:	48 83 f8 0e          	cmp    rax,0xe
  400470:	48 89 e5             	mov    rbp,rsp
  400473:	76 1b                	jbe    400490 <deregister_tm_clones+0x30>
  400475:	b8 00 00 00 00       	mov    eax,0x0
  40047a:	48 85 c0             	test   rax,rax
  40047d:	74 11                	je     400490 <deregister_tm_clones+0x30>
  40047f:	5d                   	pop    rbp
  400480:	bf 48 10 60 00       	mov    edi,0x601048
  400485:	ff e0                	jmp    rax
  400487:	66 0f 1f 84 00 00 00 	nop    WORD PTR [rax+rax*1+0x0]
  40048e:	00 00
  400490:	5d                   	pop    rbp
  400491:	c3                   	ret
  400492:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]
  400496:	66 2e 0f 1f 84 00 00 	nop    WORD PTR cs:[rax+rax*1+0x0]
  40049d:	00 00 00

00000000004004a0 <register_tm_clones>:
  4004a0:	be 48 10 60 00       	mov    esi,0x601048
  4004a5:	55                   	push   rbp
  4004a6:	48 81 ee 48 10 60 00 	sub    rsi,0x601048
  4004ad:	48 c1 fe 03          	sar    rsi,0x3
  4004b1:	48 89 e5             	mov    rbp,rsp
  4004b4:	48 89 f0             	mov    rax,rsi
  4004b7:	48 c1 e8 3f          	shr    rax,0x3f
  4004bb:	48 01 c6             	add    rsi,rax
  4004be:	48 d1 fe             	sar    rsi,1
  4004c1:	74 15                	je     4004d8 <register_tm_clones+0x38>
  4004c3:	b8 00 00 00 00       	mov    eax,0x0
  4004c8:	48 85 c0             	test   rax,rax
  4004cb:	74 0b                	je     4004d8 <register_tm_clones+0x38>
  4004cd:	5d                   	pop    rbp
  4004ce:	bf 48 10 60 00       	mov    edi,0x601048
  4004d3:	ff e0                	jmp    rax
  4004d5:	0f 1f 00             	nop    DWORD PTR [rax]
  4004d8:	5d                   	pop    rbp
  4004d9:	c3                   	ret
  4004da:	66 0f 1f 44 00 00    	nop    WORD PTR [rax+rax*1+0x0]

00000000004004e0 <__do_global_dtors_aux>:
  4004e0:	80 3d 61 0b 20 00 00 	cmp    BYTE PTR [rip+0x200b61],0x0        # 601048 <__TMC_END__>
  4004e7:	75 11                	jne    4004fa <__do_global_dtors_aux+0x1a>
  4004e9:	55                   	push   rbp
  4004ea:	48 89 e5             	mov    rbp,rsp
  4004ed:	e8 6e ff ff ff       	call   400460 <deregister_tm_clones>
  4004f2:	5d                   	pop    rbp
  4004f3:	c6 05 4e 0b 20 00 01 	mov    BYTE PTR [rip+0x200b4e],0x1        # 601048 <__TMC_END__>
  4004fa:	f3 c3                	repz ret
  4004fc:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]

0000000000400500 <frame_dummy>:
  400500:	bf 20 0e 60 00       	mov    edi,0x600e20
  400505:	48 83 3f 00          	cmp    QWORD PTR [rdi],0x0
  400509:	75 05                	jne    400510 <frame_dummy+0x10>
  40050b:	eb 93                	jmp    4004a0 <register_tm_clones>
  40050d:	0f 1f 00             	nop    DWORD PTR [rax]
  400510:	b8 00 00 00 00       	mov    eax,0x0
  400515:	48 85 c0             	test   rax,rax
  400518:	74 f1                	je     40050b <frame_dummy+0xb>
  40051a:	55                   	push   rbp
  40051b:	48 89 e5             	mov    rbp,rsp
  40051e:	ff d0                	call   rax
  400520:	5d                   	pop    rbp
  400521:	e9 7a ff ff ff       	jmp    4004a0 <register_tm_clones>

0000000000400526 <main>:
  400526:	55                   	push   rbp
  400527:	48 89 e5             	mov    rbp,rsp
  40052a:	8b 15 1c 0b 20 00    	mov    edx,DWORD PTR [rip+0x200b1c]        # 60104c <y>
  400530:	8b 0d 02 0b 20 00    	mov    ecx,DWORD PTR [rip+0x200b02]        # 601038 <x>
  400536:	48 8b 05 03 0b 20 00 	mov    rax,QWORD PTR [rip+0x200b03]        # 601040 <s>
  40053d:	89 ce                	mov    esi,ecx
  40053f:	48 89 c7             	mov    rdi,rax
  400542:	b8 00 00 00 00       	mov    eax,0x0
  400547:	e8 b4 fe ff ff       	call   400400 <printf@plt>
  40054c:	b8 00 00 00 00       	mov    eax,0x0
  400551:	5d                   	pop    rbp
  400552:	c3                   	ret
  400553:	66 2e 0f 1f 84 00 00 	nop    WORD PTR cs:[rax+rax*1+0x0]
  40055a:	00 00 00
  40055d:	0f 1f 00             	nop    DWORD PTR [rax]

0000000000400560 <__libc_csu_init>:
  400560:	41 57                	push   r15
  400562:	41 56                	push   r14
  400564:	41 89 ff             	mov    r15d,edi
  400567:	41 55                	push   r13
  400569:	41 54                	push   r12
  40056b:	4c 8d 25 9e 08 20 00 	lea    r12,[rip+0x20089e]        # 600e10 <__frame_dummy_init_array_entry>
  400572:	55                   	push   rbp
  400573:	48 8d 2d 9e 08 20 00 	lea    rbp,[rip+0x20089e]        # 600e18 <__init_array_end>
  40057a:	53                   	push   rbx
  40057b:	49 89 f6             	mov    r14,rsi
  40057e:	49 89 d5             	mov    r13,rdx
  400581:	4c 29 e5             	sub    rbp,r12
  400584:	48 83 ec 08          	sub    rsp,0x8
  400588:	48 c1 fd 03          	sar    rbp,0x3
  40058c:	e8 37 fe ff ff       	call   4003c8 <_init>
  400591:	48 85 ed             	test   rbp,rbp
  400594:	74 20                	je     4005b6 <__libc_csu_init+0x56>
  400596:	31 db                	xor    ebx,ebx
  400598:	0f 1f 84 00 00 00 00 	nop    DWORD PTR [rax+rax*1+0x0]
  40059f:	00
  4005a0:	4c 89 ea             	mov    rdx,r13
  4005a3:	4c 89 f6             	mov    rsi,r14
  4005a6:	44 89 ff             	mov    edi,r15d
  4005a9:	41 ff 14 dc          	call   QWORD PTR [r12+rbx*8]
  4005ad:	48 83 c3 01          	add    rbx,0x1
  4005b1:	48 39 eb             	cmp    rbx,rbp
  4005b4:	75 ea                	jne    4005a0 <__libc_csu_init+0x40>
  4005b6:	48 83 c4 08          	add    rsp,0x8
  4005ba:	5b                   	pop    rbx
  4005bb:	5d                   	pop    rbp
  4005bc:	41 5c                	pop    r12
  4005be:	41 5d                	pop    r13
  4005c0:	41 5e                	pop    r14
  4005c2:	41 5f                	pop    r15
  4005c4:	c3                   	ret
  4005c5:	90                   	nop
  4005c6:	66 2e 0f 1f 84 00 00 	nop    WORD PTR cs:[rax+rax*1+0x0]
  4005cd:	00 00 00

00000000004005d0 <__libc_csu_fini>:
  4005d0:	f3 c3                	repz ret

Disassembly of section .fini:

00000000004005d4 <_fini>:
  4005d4:	48 83 ec 08          	sub    rsp,0x8
  4005d8:	48 83 c4 08          	add    rsp,0x8
  4005dc:	c3                   	ret
```

> 删除符号表信息会对反汇编有一定的影响

在mac下面编译出来的不是elf文件，mac会把一些符号信息删除掉，这可能会造成一定影响，
这个时候需要使用go自带的编译工具：  `go build -ldflags "-w -s" -o test test.go`

#### upx

可执行文件的压缩工具