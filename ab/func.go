package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/3JoB/unsafeConvert"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func AB(c *cli.Context) error {
	// 从 flag 获取选项
	uri := c.String("url")
	if uri == "" {
		uri = c.Args().Get(c.Args().Len() - 1)
	}
	if _, err := url.Parse(uri); err != nil {
		return err
	}
	n := c.Int("n")
	concurrency := c.Int("c")
	// 构建请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)
	// 添加 header
	headers := c.StringSlice("H")
	if len(headers) != 0 {
		for _, h := range headers {
			split := strings.SplitN(h, ":", 2)
			if len(split) != 2 {
				break
			}
			req.Header.Add(split[0], split[1])
		}
	}
	// Basic 认证
	if c.String("A") != "" {
		auth := strings.SplitN(c.String("A"), ":", 2)
		req.Header.Set("Authorization", "Basic "+basicAuth(auth[0], auth[1]))
	}
	// Cookie
	if c.String("C") != "" {
		req.Header.Add("cookie", c.String("C"))
	}
	// 发出请求
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}
	defer fasthttp.ReleaseResponse(resp)
	// 统计结果
	type stats struct {
		requests   int
		failures   int
		throughput float64
		durations  []time.Duration
		mean       time.Duration
		median     time.Duration
		stddev     time.Duration
	}
	results := stats{}
	results.requests = n // 总请求数
	fstart := time.Now()
	for i := 0; i < n; i++ {
		// 并发请求
		go func() {
			start := time.Now()
			respx := fasthttp.AcquireResponse()
			if err := fasthttp.Do(req, respx); err != nil {
				results.failures++ // 失败的请求数
			} else {
				results.durations = append(results.durations, time.Since(start))
			}
			fasthttp.ReleaseResponse(respx)
		}()
	}

	// 等待所有请求完成
	time.Sleep(time.Duration(n/concurrency) * time.Second)

	// 计算统计信息
	results.mean = mean(results.durations)     // 平均响应时间
	results.median = median(results.durations) // 中位数响应时间
	results.stddev = stddev(results.durations) // 标准偏差
	end := time.Now()
	results.throughput = float64(n) / end.Sub(fstart).Seconds() // 吞吐量

	// 打印结果
	fmt.Printf("Server Software: %s\n", unsafeConvert.StringReflect(resp.Header.Server()))
	fmt.Printf("Concurrency Level: %d\n", concurrency)
	fmt.Printf("Time taken for tests: %d secondsv\n", int(time.Now().Add(time.Duration(n/concurrency)*time.Second).Sub(fstart)/time.Second))
	fmt.Printf("Complete requests: %d\n", n)
	fmt.Printf("Failed requests: %d\n", results.failures)
	fmt.Printf("Total transferred: %d bytes\n", n*1024) // 估计值
	fmt.Printf("Requests per second: %.2f [#/sec]\n", results.throughput)
	fmt.Printf("Time per request: %.3f [ms]\n", float64(results.mean)/float64(time.Millisecond))
	fmt.Printf("Mean: %v [ms]\n", results.mean.Milliseconds())
	fmt.Printf("Median: %v [ms]\n", results.median.Milliseconds())
	fmt.Printf("Stddev: %v [ms]\n", results.stddev.Milliseconds())
	return nil
}
