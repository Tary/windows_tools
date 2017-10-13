package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const (
	K  = 1024
	Kf = float64(K)
	M  = K * K
	Mf = float64(M)
	G  = M * K
	Gf = float64(G)
	T  = G * K
	Tf = float64(T)
)

func parseSize(size, precision int) string {
	if size < K {
		return fmt.Sprintf("%dB", size)
	}

	var base float64
	var measure string
	if size < M {
		base = Kf
		measure = "K"
	} else if size < G {
		base = Mf
		measure = "M"
	} else if size < T {
		base = Gf
		measure = "G"
	} else {
		base = Tf
		measure = "T"
	}

	var d int = int(math.Pow10(precision))
	val := int((float64(size) / base) * float64(d))

	integral := val / d
	decimal := val % d

	for ; decimal > 0 && decimal%10 == 0; decimal /= 10 {
	}

	if decimal != 0 {
		return fmt.Sprintf("%d.%d%s", integral, decimal, measure)
	}
	return fmt.Sprintf("%d%s", integral, measure)

}

func main() {
	var help = flag.Bool("h", false, "使用说明")
	var dirName *string = flag.String("c", ".", "目录")
	var precision *int = flag.Int("n", 2, "小数部分精确度")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	fileInfo, err := ioutil.ReadDir(*dirName)
	if err != nil {
		os.Stderr.WriteString("输入路径错误\n")
		return
	}

	format := fmt.Sprintf("%%s %%s %%%ds %%s\n", *precision+5) //三位整数,小数点及单位

	const timeFMTLayout string = "2006-01-02 03:04:05"
	for _, v := range fileInfo {
		ln := fmt.Sprintf(format,
			v.Mode().String(),
			v.ModTime().Format(timeFMTLayout),
			parseSize(int(v.Size()), *precision),
			v.Name())
		os.Stdout.WriteString(ln)
	}
}
