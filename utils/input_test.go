// Only making this to iteravely test things unfamalliar to me like requesting content of a web page
package utils

import (
	"fmt"
	"net/http"
	"testing"
)

func TestInputFromRequest(t *testing.T) {
	resp, err := http.Get("https://adventofcode.com/2022/day/1/input")
	if err != nil {
		t.Fatal(err)
	}

	lines := ReadLinesFromReader(resp.Body)
	fmt.Println(lines)
}
