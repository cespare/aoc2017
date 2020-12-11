package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func init() {
	addSolutions(1, problem1)
}

func problem1(ctx *problemContext) {
	b, err := ioutil.ReadAll(ctx.f)
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.TrimSpace(b)
	ctx.reportLoad()

	ctx.reportPart1(captcha(b, 1))
	ctx.reportPart2(captcha(b, len(b)/2))
}

func captcha(digits []byte, delta int) int64 {
	var n int64
	for i, c := range digits {
		j := (i + delta) % len(digits)
		d0 := c - '0'
		d1 := digits[j] - '0'
		if d0 == d1 {
			n += int64(d0)
		}
	}
	return n
}
