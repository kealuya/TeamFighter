package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

/*
利用协程，高效读取超大文件
 sync.Pool 可以重复利用对象，但是不会保存对象，也就是 长链接之类的内容不能存储，只能存临时存储数据的对象
*/
func main() {

	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 4*1024*1024)
		return lines
	}}
	f := "/Volumes/生态硬盘/software/myeclipse-2017-ci-3-offline-installer-macosx.dmg"
	ff, _ := os.OpenFile(f, os.O_RDONLY, 0666)
	b := bufio.NewReader(ff)

	wg := sync.WaitGroup{}

	// 判断需要多少个协程的案例
	chunkSize := 100 //一个协程打算处理多少条数据
	n := len([]string{"预计有多少条需要处理的数据"})
	noOfThread := n / chunkSize
	if n%chunkSize != 0 { //总数 除以 每协程处理数 如果有余数，就再追加一个协程
		noOfThread++
	}

	noOfThread = 20
	for {
		bufByte := linesPool.Get().([]byte)
		n, err := b.Read(bufByte)
		bufByte = bufByte[:n] //因为bufByte是全尺寸buf，需要获取真实读取了内容的部分
		//fmt.Println("read n ", n)

		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				fmt.Println(err)
				break
			}
		}
		for i := 0; i < noOfThread; i++ {
			wg.Add(1)
			go func() {
				// do something
				// todo
				fmt.Println(runtime.NumGoroutine(), "::", len(bufByte))
				wg.Done()
			}()
		}
	}
	wg.Wait()
	fmt.Println("over")
}
