package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

type block struct {
	pos  int
	size int
}

func main() {
	part, file, start := utils.Setup()
	strNums := strings.Split(string(file), "")
	input := parse(strNums)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(strNums)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func parse(strNums []string) []string {
	var res []string
	id := 0

	for i := 0; i < len(strNums); i += 2 {
		emptySpace := 0
		if i+1 < len(strNums) {
			emptySpace, _ = strconv.Atoi(strNums[i+1])
		}
		fileSize, _ := strconv.Atoi(strNums[i])
		for range fileSize {
			strId := strconv.Itoa(id)
			res = append(res, strId)
		}
		for range emptySpace {
			res = append(res, ".")
		}
		id++
	}

	return res
}

func checkSum(diskMap []string) int {
	sum := 0
	for pos, val := range diskMap {
		if val == "." {
			break
		}
		fileId, _ := strconv.Atoi(val)
		sum += pos * fileId
	}
	return sum
}

// [0 0 . . . 1 1 1 . . . 2 . . . 3 3 3 . 4 4 . 5 5 5 5 . 6 6 6 6 . 7 7 7 . 8 8 8 8 9 9]
func part1(diskMap []string) int {
	l, r := 0, len(diskMap)-1
	freeSpace := "."
	for l < r {
		if diskMap[l] != freeSpace {
			l++
		} else {
			diskMap[l] = diskMap[r]
			diskMap[r] = freeSpace
			r--
		}
	}

	return checkSum(diskMap)
}

func getBlocks(diskMap []string) (map[int]block, []block, int) {
	files := make(map[int]block)
	var blanks []block
	fid := 0
	pos := 0
	for i, char := range diskMap {
		size, _ := strconv.Atoi(char)
		if size == 0 {
			continue
		}
		if i%2 == 0 {
			// file
			files[fid] = block{size: size, pos: pos}
			fid++
		} else {
			// blank
			blanks = append(blanks, block{size: size, pos: pos})
		}
		pos += size
	}
	return files, blanks, fid - 1
}

func checkSum2(files map[int]block) int {
	sum := 0
	for k, v := range files {
		for i := 0; i < v.size; i++ {
			sum += k * (v.pos + i)
		}

	}
	return sum
}

func part2(fileSystem []string) int {
	files, blanks, fid := getBlocks(fileSystem)

	for fid > 0 {
		fsize := files[fid].size
		fpos := files[fid].pos
		for i, v := range blanks {
			blankSize := v.size
			blankPos := v.pos
			if fid == 2 {
			}
			if fpos <= blankPos {
				// too far
				blanks = blanks[:i]
				break
			}
			if fsize <= blankSize {
				// it fits so move it into the blank spot
				files[fid] = block{size: fsize, pos: blankPos}

				if blankSize == fsize {
					blanks = append(blanks[:i], blanks[i+1:]...)
				} else {
					blanks[i] = block{pos: blanks[i].pos + fsize, size: blanks[i].size - fsize}
				}
				break
			}
		}
		fid--
	}

	return checkSum2(files)
}
