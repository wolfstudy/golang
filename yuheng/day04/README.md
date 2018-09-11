# 变量和常量

## 虚拟内存

### 实模式（8086、boot）与保护模式
实模式：直接操作物理内存，之前很早的时候，现在的话OS启动的时候，cpu会让os直接操作内存
* 缺点：越界、安全问题
* 好处：修改游戏数据，只要找到内存地址就可以直接修改。

保护模式：80286开始，win95把保护模式发扬广大，

* os让每一个进程认为自己拥有一个完整的内存空间。总共分为两段，一部分给os使用，一部分给用户使用。
* 我们目前看到的地址都是虚拟地址，看不到真实的物理地址，虚拟内存是os给每一个进程构建出来的，但是必须和真实的物理地址一一映射。
* 每一个进程都可以有相同的地址空间，注意是虚拟出来的。

### 虚拟内存器 、VA/PA、 MMU/TLB 、PT/PTE

* cpu中有一个MMU会把虚拟地址翻译为物理地址，所以就有：VA--->MMU---->PA
* 好处：这样的话你肯定不会越界，每一个进程都有一个独立的虚拟地址空间，它会映射到某一块真实的物理内存上，这也是保护模式的一种优势。
* os如何与MMU打交道？每个进程都有一个页表，是按照页为单位管理内存的。flag为0，没有在物理内存上映射，这时候，os会在物理内存上分配一个对应的页，然后虚拟内存的
页指向真实物理内存的页，分配之后，flag置为1。页在每一个进程的头部。
* MMU是cpu上的，对cpu来说，主存是很慢的，MMU内部有一个小的缓存叫TLB，把每一个查过的这种映射关系存储到TLB中。
* 怎么去分配？程序去申请内存的时候，给你返回的是一个VA地址和一个根本不存在的内存，这个时候根本没有分配内存，os只是标记，这段地址被人用过了。

### 缺页异常、 swap in/out、 thrashing

* 当进行文件读写的时候，cpu会调用MMU去VA的页表中去查，如果没有找到的话，会触发一个缺页异常，造成系统中断，它会按照你的需求，分配物理内存，写回给VA，并将页表的flag置为1，这个时候
再去读写的时候，就可以找到真实的物理内存地址。缺页异常是因为你的虚拟地址和物理地址并没有建立好映射关系。这时候cpu会去分配真实的物理内存，然后重新执行
读写操作。

* 4GB的物理内存是否能够分配8GB的对象？
    * 首先去分配空间：64为的地址空间，最大到达上万PB，肯定可以给你分配出来，然后初始化这8GB的空间，往里面写东西，问题了来了，PM只有4GB，而VM有8GB，真实的4GB
    映射完了，它会触发一定的换出算法，把这个页的数据写到disk上，这个页变为空，就可以重新建立映射关系。所以只要主存和disk的size加起来才是真实的虚拟地址空间的映射。
    * 主存中一般保存的是热数据，一些不常用的数据换出到disk上。很像我们拿redis做cache，mysql做持久化存储。当我们需要读取disk上的文件到主存上的时候，我们称为换入。
    * 如果我们的内存足够大，关掉swap，禁止换出。swap一般是我们内存小、不够用的时候才会打开。
    * 主存其实就是在虚拟内存和真实的disk上的存储的中间缓存。
* mongodb早期的引擎，超出主存大小的时候，性能会出现很大的抖动，主要消耗在数据的换入换出。

### 存储器层次

disk-->main mem-->L3-->L2-->L1-->寄存器 （多级缓存体系）

访问权限分为：
* Ring0 os内核使用（可以访问所有的物理地址）
* Ring1、Ring2 驱动程序使用
* Ring3 用户程序使用

### 机会主义内存分配

内存分配，os一般会采取机会主义的手段来做。

还是从分配8GB内存开始说起（4GB主存）：

1. os会给你分配这么大的一个地址空间，注意是地址空间，只是给你保留这么一个空间，确保别人不会使用。
2. 当你进行读写操作的时候，因为没有物理内存映射，会触发缺页异常，给你分配物理内存，缺页异常不回一次性给你分配这么大的内存，而是你写一个页，os给你补一个
页，所以PA可能不是连续的，虽然你看到的VA是连续的。

mem.go

```
package main

import (
	"fmt"
	"time"
)

func main() {

	var x [1 << 20 * 100]byte //100MB

	for i := 0; i < len(x); i++ {
		x[i] = byte(i % 255)

		if i%(1<<20) == 0 {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
```

Makefile

```
run: mem.go
    go build -o test $^ && ./test
stat:
    pidstat -r -C test 1
clean:
    -rm test

```


可以看到VSZ虚拟内存地址是一次性分配好的，但是RSS是不断往上分配的。监控进程的内存相关状态。
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go# pidstat -r -C test 1
Linux 4.4.0-62-generic (iZm5eakzlq9sm4ycm9lw0wZ) 	08/30/2018 	_x86_64_	(4 CPU)

08:56:59 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:00 PM     0     25417    108.91      0.00  110596    8276   0.03  test

08:57:00 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:01 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:02 PM     0     25417      1.00      0.00  110596   10324   0.03  test

08:57:02 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:03 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:04 PM     0     25417      1.00      0.00  110596   12372   0.04  test

08:57:04 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:05 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:06 PM     0     25417      1.00      0.00  110596   14420   0.04  test

08:57:06 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:07 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:08 PM     0     25417      1.00      0.00  110596   16468   0.05  test

08:57:08 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:09 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:10 PM     0     25417      1.00      0.00  110596   18516   0.06  test

08:57:10 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command

08:57:11 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
08:57:12 PM     0     25417      1.00      0.00  110596   20564   0.06  test
```

### dstat、pidstat 使用示例


dstat

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go# dstat
You did not select any stats, using -cdngy by default.
----total-cpu-usage---- -dsk/total- -net/total- ---paging-- ---system--
usr sys idl wai hiq siq| read  writ| recv  send|  in   out | int   csw
  8   4  87   1   0   1|  52k 1966k|   0     0 |   0     0 |1253    12k
 12   0  88   0   0   0|   0     0 |  34k 7943k|   0     0 | 904   695
 14   1  86   0   0   0|   0     0 |  19k 8274k|   0     0 | 817   679
 12   1  86   0   0   0|   0    32k|  11k 6551k|   0     0 | 776   782
 20   1  79   0   0   0|   0     0 |  52k 1475k|   0     0 |1200   911
 15   0  85   0   0   0|   0     0 |  13k 6629k|   0     0 | 853   809
 13   0  86   0   0   0|   0  4096B|  12k 8072k|   0     0 | 740   674
 17   3  80   0   0   0|   0     0 |  20k 8638k|   0     0 |1169   921
 23  23  54   0   0   1|   0     0 |  17k 4400k|   0     0 |1961  1388
 18   2  80   0   0   0|   0    52k|  49k 6635k|   0     0 |1175   861
 15   1  84   0   0   0|   0     0 |5936B 8735k|   0     0 | 722   777
 14   1  85   0   0   0|   0  8192B|  13k 8569k|   0     0 | 676   654
 18   1  80   0   0   0|   0     0 |  28k 3849k|   0     0 |1552  1031
 17   1  82   0   0   0|   0     0 |  34k 6438k|   0     0 | 812   685
 12   1  87   0   0   0|   0    28k|5874B 9028k|   0     0 | 768   877
 23  17  60   0   0   1|   0     0 |  23k 9735k|   0     0 |1627  1220
 22  11  67   0   0   0|   0   828k|  36k 2523k|   0     0 |1349  1173
 19   1  80   0   0   0|   0     0 |  16k 9366k|   0     0 | 906   751
 14   0  85   0   0   0|   0     0 |6152B 9148k|   0     0 | 747   760
 16   1  83   0   0   0|   0    56k|5678B 8999k|   0     0 | 736   837
 20   1  79   0   0   0|   0     0 |  45k 2028k|   0     0 |1118   803
 20   0  79   0   0   0|   0     0 |  31k   10M|   0     0 | 970   648
```

## 处理器


### 对称多处理器（SMP）NUMA的问题

* SMP：目前使用的，多核，共享主线、主存等资源，地位都是平等的。
* NUMA：超级计算机，计算能力强，有多个核，有自己的主线、主存。

### 物理核心、逻辑核心

`nproc` 查看有多少个cpu核

### 缓存：L1（d、i）L2、L3

L1：数据和指令，离cpu最近

### 超线程（HT），可能因共享cache造成性能问题

计算核心多了。


### lscpu使用示例

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go# lscpu
Architecture:          x86_64
CPU op-mode(s):        32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                4
On-line CPU(s) list:   0-3
Thread(s) per core:    2
Core(s) per socket:    2
Socket(s):             1
NUMA node(s):          1
Vendor ID:             GenuineIntel
CPU family:            6
Model:                 63
Model name:            Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz
Stepping:              2
CPU MHz:               2494.224
BogoMIPS:              4988.44
Hypervisor vendor:     KVM
Virtualization type:   full
L1d cache:             32K
L1i cache:             32K
L2 cache:              256K
L3 cache:              30720K
NUMA node0 CPU(s):     0-3
```

## 进程内存模型

### 可执行文件和进程的差异

* 可执行文件，编译的时候就将ELF文件映射到虚拟地址空间，叫虚拟地址将这段空间地址给我预留着。所有的代码合并到`.text`段，起始地址加偏移量就可以确定应该写到哪一个真实
的物理地址上。但是物理地址就不可以了，因为你压根不知道。也算是虚拟地址的好处。

* 如何映射到内存中的？

`readelf -l test`这个才是程序加载到内存时你需要使用的。
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go# readelf -l test

Elf file type is EXEC (Executable file)
Entry point 0x4509d0
There are 7 program headers, starting at offset 64

Program Headers:
  Type           Offset             VirtAddr           PhysAddr
                 FileSiz            MemSiz              Flags  Align
  PHDR           0x0000000000000040 0x0000000000400040 0x0000000000400040
                 0x0000000000000188 0x0000000000000188  R      1000
  NOTE           0x0000000000000f9c 0x0000000000400f9c 0x0000000000400f9c
                 0x0000000000000064 0x0000000000000064  R      4
  LOAD           0x0000000000000000 0x0000000000400000 0x0000000000400000
                 0x0000000000083810 0x0000000000083810  R E    1000
  LOAD           0x0000000000084000 0x0000000000484000 0x0000000000484000
                 0x00000000000915e8 0x00000000000915e8  R      1000
  LOAD           0x0000000000116000 0x0000000000516000 0x0000000000516000
                 0x0000000000013700 0x00000000000324d8  RW     1000
  GNU_STACK      0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000  RW     8
  LOOS+5041580   0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000         8

 Section to Segment mapping:
  Segment Sections...
   00
   01     .note.go.buildid
   02     .text .note.go.buildid
   03     .rodata .typelink .itablink .gosymtab .gopclntab
   04     .noptrdata .data .bss .noptrbss
   05
   06
```

这二者之间有映射关系，但不是一一对应的。
### 进程内存模型


### 使用readelf -l查看内存段映射



### 使用gdb查看运行期内存模型


### heap、stack、mmap

## 常量和变量

* 常量：`const x = 0x1000` 没有运行期常量这种概念，常量会在执行过程中被展开到使用的地方。所以是没有地址的。常量是一个数据，数据本身没有地址，是只读的。
* 变量：变量代表的是某一段内存空间，是有地址的，变量名就是这段虚拟地址空间的一个别名。

### 常量陷阱

* 两个程序，B引用A的全局常量X，第一次的时候X=100，编译.so文件，连接到B使用，后来把X修改为200，但是B引用的X还是100。
    * 将常量改为变量（变量不会展开）有一个VA地址空间
    * 将常量改为函数（每次都是动态输出）
这里要注意。