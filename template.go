// generated by $%U%$ at $%D%$-$%M%$-$%Y%$
// template version: 2022-01-11

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// 0000.A
func solve() {
	// here goes solution
}

func main() {
	in := bufio.NewReader(os.Stdin)

	t := ReadInt(in)
	for i := 0; i < t; i++ {
		// here goes input reading
		// input := ReadIntSlice(in)
		solve()
	}
}

// helpers for convenient input reading

func ReadInt(in *bufio.Reader) int {
	line, _ := in.ReadString('\n')
	line = strings.Trim(line, "\n\r")
	n, _ := strconv.Atoi(line)
	return n
}

func ReadLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	line = strings.Trim(line, "\n\r")
	return line
}

func ReadLineParts(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.Trim(line, "\n\r")
	parts := strings.Split(line, " ")
	return parts
}

func ReadIntSlice(in *bufio.Reader) []int {
	parts := ReadLineParts(in)
	slice := make([]int, len(parts))
	for i, n := range parts {
		slice[i] = MustAtoi(n)
	}
	return slice
}

func ReadInt64Slice(in *bufio.Reader) []int64 {
	parts := ReadLineParts(in)
	slice := make([]int64, len(parts))
	for i, n := range parts {
		slice[i] = MustParseInt(n, 10, 64)
	}
	return slice
}

func ReadUint64Slice(in *bufio.Reader) []uint64 {
	parts := ReadLineParts(in)
	slice := make([]uint64, len(parts))
	for i, n := range parts {
		slice[i] = MustParseUint(n, 10, 64)
	}
	return slice
}

func MustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func MustParseInt(s string, base, bitSize int) int64 {
	v, _ := strconv.ParseInt(s, base, bitSize)
	return v
}

func MustParseUint(s string, base, bitSize int) uint64 {
	v, _ := strconv.ParseUint(s, base, bitSize)
	return v
}

// helper for convenient answer output

func Output(v interface{}) {
	switch v.(type) {
	case bool:
		if v == true {
			Println("YES")
		} else {
			Println("NO")
		}
	default:
		Println(v)
	}
}
