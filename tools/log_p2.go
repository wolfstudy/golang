package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type reader interface {
	read(chan []byte)
}

type writer interface {
	write(chan string)
}

type ReadFromFile struct {
	path string
}

func (r *ReadFromFile) read(rc chan []byte) {
	// open file
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file failure：%s", err.Error()))
	}

	br := bufio.NewReader(f)
	for {
		line, err := br.ReadSlice('\n')
		if err == io.EOF {
			close(rc)
			break
		} else if err != nil {
			panic(fmt.Sprintf("Read Byts err:%s", err.Error()))
		} else {
			rc <- line
		}
	}
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  reader
	write writer
}

func (l *LogProcess) Process() {
	// 解析模块
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}

	isEnd := false
	for {
		select {
		case tmp, ok := <-l.rc:
			if !ok {
				isEnd = true
				close(l.wc)
				break
			} else {
				l.wc <- strings.ToUpper(string(tmp))
			}
		}
		if isEnd {
			break
		}
	}
}

type WriteToInfluxDB struct {
	influxDBDsn string //influx data source
}

func (w *WriteToInfluxDB) write(wc chan string) {
	// 写入模块
	for {
		select {
		case tmp, ok := <-wc:
			if !ok {
				os.Exit(0)
			} else {
				fmt.Println(string(tmp)[:len(tmp)-2])
			}
		}
	}
}

func main() {
	r := &ReadFromFile{
		path: "access.log",
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

	go lp.read.read(lp.rc)
	go lp.Process()
	go lp.write.write(lp.wc)

	time.Sleep(5 * time.Second)
}