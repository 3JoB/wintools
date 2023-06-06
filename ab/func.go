package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/3JoB/unsafeConvert"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func AB(c *cli.Context) error {
	uri := ""
	if c.Args().Len() != 0 {
		l := c.Args().Len() - 1
		if l < 0 {
			l = 0
		}
		uri = c.Args().Slice()[l]
	}
	if uri == "" {
		return PrintUsage()
	}
	n := CliFlagRequestNumber.Get(c)
	concurrency := CliFlagConcurrentRequestNumber.Get(c)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)

	content_type := CliFlagContentType.Get(c)
	cookie := CliFlagCookieNameValue.Get(c)

	headers := CliFlagCustomHeader.Get(c)
	if len(headers) != 0 {
		for _, h := range headers {
			split := strings.SplitN(h, ":", 2)
			if len(split) != 2 {
				break
			}
			req.Header.Add(split[0], split[1])
		}
	}

	basic_auth := CliFlagAuthUsernamePassword.Get(c)

	if basic_auth != "" {
		auth := strings.SplitN(basic_auth, ":", 2)
		req.Header.Set("Authorization", "Basic "+basicAuth(auth[0], auth[1]))
	}

	if cookie != "" {
		req.Header.Add("cookie", cookie)
	}

	if content_type != "" {
		req.Header.Set("Content-Type", content_type)
	}

	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}
	defer fasthttp.ReleaseResponse(resp)

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
	results.requests = n
	fstart := time.Now()
	for i := 0; i < n; i++ {
		go func() {
			start := time.Now()
			respx := fasthttp.AcquireResponse()
			if err := fasthttp.Do(req, respx); err != nil {
				results.failures++
			} else {
				results.durations = append(results.durations, time.Since(start))
			}
			fasthttp.ReleaseResponse(respx)
		}()
	}

	time.Sleep(time.Duration(n/concurrency) * time.Second)

	results.mean = mean(results.durations)
	results.median = median(results.durations)
	results.stddev = stddev(results.durations)
	end := time.Now()
	results.throughput = float64(n) / end.Sub(fstart).Seconds()

	// 打印结果
	fmt.Printf("Server Software: %s\n", unsafeConvert.StringReflect(resp.Header.Server()))
	fmt.Printf("Concurrency Level: %d\n", concurrency)
	fmt.Printf("Time taken for tests: %d secondsv\n", int(time.Now().Add(time.Duration(n/concurrency)*time.Second).Sub(fstart)/time.Second))
	fmt.Printf("Complete requests: %d\n", n)
	fmt.Printf("Failed requests: %d\n", results.failures)
	fmt.Printf("Total transferred: %d bytes\n", n*1024)
	fmt.Printf("Requests per second: %.2f [#/sec]\n", results.throughput)
	fmt.Printf("Time per request: %.3f [ms]\n", float64(results.mean)/float64(time.Millisecond))
	fmt.Printf("Mean: %v [ms]\n", results.mean.Milliseconds())
	fmt.Printf("Median: %v [ms]\n", results.median.Milliseconds())
	fmt.Printf("Stddev: %v [ms]\n", results.stddev.Milliseconds())
	return nil
}
