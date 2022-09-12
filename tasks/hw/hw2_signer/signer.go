package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var (
	md5M sync.Mutex
)

func dataSignerMd5(data string) string {
	md5M.Lock()
	defer md5M.Unlock()
	return DataSignerMd5(data)
}

func dataSignerCrc32multi(data ...string) (rv []string) {
	rv = make([]string, len(data))
	wg := sync.WaitGroup{}
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(i int) {
			rv[i] = DataSignerCrc32(data[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	return rv
}

func ExecutePipeline(jobs ...job) {
	chs := make([]chan any, len(jobs)+1)
	for i := 0; i < len(jobs)+1; i++ {
		chs[i] = make(chan any)
	}

	wg := sync.WaitGroup{}
	for i := range jobs {
		wg.Add(1)
		go func(j job, in, out chan any) {
			j(in, out)
			wg.Done()
			close(out)
		}(jobs[i], chs[i], chs[i+1])
	}
	wg.Wait()
}

func singleHash(data string) string {
	return strings.Join(dataSignerCrc32multi(data, dataSignerMd5(data)), "~")
}

func SingleHash(in, out chan any) {
	wg := sync.WaitGroup{}
	for d := range in {
		str := ""
		switch data := d.(type) {
		case int:
			str = fmt.Sprintf("%d", data)
		}

		wg.Add(1)
		go func(str string) {
			out <- singleHash(str)
			wg.Done()
		}(str)
	}
	wg.Wait()
}

func multiHash(data string) string {
	rvs := make([]string, 6)
	for i := 0; i < 6; i++ {
		rvs[i] = fmt.Sprintf("%d%s", i, data)
	}

	return strings.Join(dataSignerCrc32multi(rvs...), "")
}

func MultiHash(in, out chan any) {
	wg := sync.WaitGroup{}
	for d := range in {
		str := ""
		switch data := d.(type) {
		case string:
			str = data
		}

		wg.Add(1)
		go func(str string) {
			out <- multiHash(str)
			wg.Done()
		}(str)
	}
	wg.Wait()
}

func CombineResults(in, out chan any) {
	rv := []string{}
	for d := range in {
		switch data := d.(type) {
		case string:
			rv = append(rv, data)
		}
	}

	sort.Strings(rv)
	out <- strings.Join(rv, "_")
}
