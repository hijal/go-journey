package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersions(a, b string) (int, error) {

	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	for i := range max(len(as), len(bs)) {
		av, bv := 0, 0

		if i < len(as) {
			n, err := strconv.Atoi(as[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", a, as[i])
			}
			av = n
		}

		if i < len(bs) {
			n, err := strconv.Atoi(bs[i])
			if err != nil {
				return 0, fmt.Errorf("bad version %q: part %q is not a number", b, bs[i])
			}
			bv = n
		}

		switch {
		case av < bv:
			return -1, nil
		case av > bv:
			return 1, nil
		}
	}

	return 0, nil
}

func main() {
	pairs := [][2]string{
		{"1.10.0", "1.9.5"},
		{"2.0", "2.0.0"},
		{"0.9", "1.0"},
		{"1.x", "1.0"},
	}

	for _, p := range pairs {
		result, err := compareVersions(p[0], p[1])
		if err != nil {
			fmt.Printf("%s vs %s -> error: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%s vs %s -> %d\n", p[0], p[1], result)
	}
}
