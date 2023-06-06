package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

const Usage string = `Usage:
win bench [OPTIONS] <url>

Benchmark and analyze the performance of an HTTP web server.

Options:
-n, --requests INTEGER     Number of requests to perform. Default is 1000.
-c, --concurrency INTEGER  Number of concurrent connections. Default is 100.
-m, --method TEXT          HTTP method to use. Possible values are 'GET', 'POST', 'PUT' or 'DELETE'. Default is 'GET'.
-H, --header TEXT          Additional header to include. Example: 'Connection: close'. You can add multiple headers by mentioning this argument multiple times.
-A, --auth TEXT            Basic authentication credentials in the form of 'username:password' to use for the requests. 
-C, --cookie TEXT          Cookie in the form of 'name=value' to add to requests.   
-T, --content-type TEXT    Content type header to use for the requests.     
-b, --body TEXT            File containing data to POST.     
-p, --post-file TEXT       File containing HTML form data to POST.       
-u, --upload-file TEXT     File to upload (via PUT or POST). 
-x, --table-attributes TEXT   
						   <table> attributes. Example: 'border=1 cellspacing=0'. 
-y, --tr-attributes TEXT   <tr> attributes. 
-z, --td-attributes TEXT   <td> attributes.       
-X, --proxy TEXT           Proxy URL and port. Example: 'http://myproxy:8080'. 
-Z, --ciphersuite TEXT     Cipher suite to use for SSL/TLS connections. Example: 'AES128-GCM-SHA256'. 
-k, --keepalive            Use HTTP KeepAlive feature. Default is disabled.   
-w, --html                 Print out results in an HTML table. Default is disabled. 
-S, --skip-stderr          Do not display errors and warnings. Default is disabled. 
-q, --quiet                Do not display progress messages. Default is disabled. 
-g, --gnuplot              Generate a gnuplot graph file. Extension is .gnuplot. 
-e, --csv                  Generate a CSV output file. Extension is .csv.   `

func PrintUsage() error {
	fmt.Println(Usage)
	return nil
}

func basicAuth(username, password string) string {
	auth := litefmt.Sprint(username, ":", password)
	return base64.StdEncoding.EncodeToString(unsafeConvert.BytesReflect(auth))
}

func mean(nums []time.Duration) time.Duration {
	total := time.Duration(0)
	for _, n := range nums {
		total += n
	}
	return total / time.Duration(len(nums))
}

func median(nums []time.Duration) time.Duration {
	sort.Sort(timeDurations(nums))
	middle := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[middle]
	}
	return (nums[middle-1] + nums[middle]) / 2
}

func stddev(nums []time.Duration) time.Duration {
	m := mean(nums)
	var sum time.Duration
	for _, n := range nums {
		d := n - m
		sum += d * d
	}
	return time.Duration(math.Sqrt(float64(sum) / float64(len(nums))))
}

type timeDurations []time.Duration

func (t timeDurations) Len() int { return len(t) }

func (t timeDurations) Less(i, j int) bool { return t[i] < t[j] }

func (t timeDurations) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
