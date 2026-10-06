// Command election 是选举场景入口。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/election-terms/election"
)

func emit(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	sample := flag.String("sample", "", "dupvote|majority|stale|reject|work")
	flag.Parse()

	for _, s := range election.Samples() {
		if s.Name == *sample {
			emit(s.Run())
			return
		}
	}
	if *sample == "work" {
		emit(election.Work(2000))
		return
	}
	fmt.Fprintln(os.Stderr, "未知场景："+*sample)
	os.Exit(2)
}
