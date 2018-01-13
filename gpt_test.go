package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// BenchmarkGenFiles 基准测试
func BenchmarkGenFiles(b *testing.B) {
	Convey("BenchmarkGenFiles", b, func() {
		for i := 0; i < b.N; i++ {
			GenFiles(AllFiles(), time.Now())
		}
	})
}
