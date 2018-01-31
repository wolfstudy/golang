## pprof 分析

###pprof是一个优秀的分析gorutine的工具，本篇将简单介绍pprof的一些使用技巧：

`go tool pprof ip:port/debug/pprof/goroutine`

### pprof 有四个功能

* cpuprofie是分析cpu的
* heap是分析内存的
* goroutinue是分析携程的
* thread是分析线程的，应用的相对少一点

关于pprof官方文档描述的比较详细，所以不做过多的叙述：

[pprof官方文档](https://blog.golang.org/profiling-go-programs)