package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", all)
	fout, err := os.Create("result.html")
	defer fout.Close()

	if err != nil {
		panic(err)
		return
	}

	for i := 0; i < 10; i++ {
		fout.Write(all)
	}
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
