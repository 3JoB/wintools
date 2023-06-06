package main

import (
	"encoding/base64"
	"math"
	"sort"
	"time"

	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

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
