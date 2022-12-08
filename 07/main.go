package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type File struct {
	parent   *File
	name     string
	size     int
	children []*File
}

//go:embed input.txt
var input []byte

const smallDirSize = 100000
const diskSpace = 70000000
const unusedSpaceNeeded = 30000000

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	root := File{
		parent:   nil,
		name:     "---",
		size:     0,
		children: []*File{makeDir(nil, "/")},
	}

	file := &root

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")

		if line[:4] == "$ cd" {
			file = cd(file, lineParts[2])
		} else if line[:4] == "$ ls" {
			continue
		} else if lineParts[0] == "dir" {
			file.children = append(file.children, makeDir(file, lineParts[1]))
		} else {
			size, _ := strconv.Atoi(lineParts[0])
			file.children = append(file.children, makeFile(lineParts[1], size))
		}
	}
	calculateFileSizes(&root)

	fmt.Printf("part one: %v\n", getTotalSizeOfSmallDirs(&root))
	fmt.Printf("part two: %v\n", partTwo(&root))
}

func getTotalSizeOfSmallDirs(f *File) int {
	if f.children == nil {
		return 0
	}

	size := 0

	if f.size <= smallDirSize {
		size += f.size
	}

	for _, f := range f.children {
		size += getTotalSizeOfSmallDirs(f)
	}
	return size
}

func partTwo(root *File) int {
	sizeToDelete := root.size - diskSpace + unusedSpaceNeeded
	return getSmallestDirectoryToDeleteSize(root, sizeToDelete)
}

func getSmallestDirectoryToDeleteSize(dir *File, sizeToDelete int) int {
	if dir.children == nil {
		return math.MaxInt32
	}

	size := math.MaxInt32

	if dir.size >= sizeToDelete {
		size = dir.size
	}

	for _, f := range dir.children {
		fsize := getSmallestDirectoryToDeleteSize(f, sizeToDelete)
		if fsize <= size {
			size = fsize
		}
	}
	return size
}

func calculateFileSizes(f *File) int {
	if f.children == nil {
		return f.size
	}

	size := 0
	for _, f := range f.children {
		size += calculateFileSizes(f)
	}

	f.size = size
	return size
}

func cd(dir *File, name string) *File {
	if name == ".." {
		return dir.parent
	}

	for _, f := range dir.children {
		if f.name == name {
			return f
		}
	}

	return nil // No such file or directory - should not happen
}

func makeDir(parent *File, name string) *File {
	return &File{
		parent:   parent,
		name:     name,
		children: []*File{},
	}
}

func makeFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}
