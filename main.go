package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) < 3 {
		panic("arg len < 3")
	}

	imgFile := os.Args[1]

	f, err := os.Open(imgFile)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed: %v", imgFile, err))
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(fmt.Sprintf("stat file %s failed: %v", imgFile, err))
	}

	mem, err := syscall.Mmap(int(f.Fd()), 0, int(fi.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(fmt.Sprintf("mmap file %s failed: %v", imgFile, err))
	}

	img := NewImage(mem)

	command := os.Args[2]

	switch command {
	case "ls":
		if len(os.Args) < 4 {
			panic("arg len < 4")
		}

		path := os.Args[3]

		err := ls(img, path)
		if err != nil {
			panic(fmt.Sprintf("ls %s failed: %v", path, err))
		}
	case "get":
	case "put":
	case "rm":
	}
}

func ls(img *Image, path string) error {
	return nil
}
