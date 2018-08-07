package main

import (
	"strings"
	"fmt"
	"time"
	"os"
	"bufio"
	"io"
)

type Read interface {
	Reader(rc chan []byte)
}

type Write interface {
	Writer(wc chan string)
}

type ReadFromFile struct {
	path string //读取文件的路径
}

type WriteToInfluxDB struct {
	influxDBDsn string //influx data source
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Read
	write Write
}

// 读取模块
func (r *ReadFromFile) Reader(rc chan []byte) {
	// open file
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file failure：%s", err.Error()))
	}

	// 从文件末尾开始逐行读取
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("Read Byts err:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

func (l *LogProcess) Process() {
	// 解析模块
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

func (w *WriteToInfluxDB) Writer(wc chan string) {
	// 写入模块
	for v := range wc {
		fmt.Println(v)
	}
}

func main() {
	r := &ReadFromFile{
		path: "/Users/wolf4j/go/src/golang/imoock/access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "",
	}

	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lp.read.Reader(lp.rc)
	go lp.Process()
	go lp.write.Writer(lp.wc)

	time.Sleep(30 * time.Second)
}
