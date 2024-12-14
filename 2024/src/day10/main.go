package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Read the input file as bytes
func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return fileBytes
}

// Read the input file as lines
func ReadFileLines(path string) []string {
	fileBytes := ReadFileBytes(path)
	lines := []string{}
	separator := []byte("\n")
	for _, line := range bytes.Split(fileBytes, separator) {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}
	return lines
}

// Parse the disk map into files and free space representation
func GetFiles(code string) []int {
	var fileList []int
	for i, char := range code {
		n, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			// If even index, it's a file
			for j := 0; j < n; j++ {
				fileList = append(fileList, i/2)
			}
		} else {
			// If odd index, it's free space
			for j := 0; j < n; j++ {
				fileList = append(fileList, -1)
			}
		}
	}
	return fileList
}

// Get blocks and free space as separate mappings
func GetMappedBlocks(files []int) (map[int]int, map[int]int) {
	blocks := make(map[int]int)
	frees := make(map[int]int)

	start := 0
	for i, file := range files {
		if file != -1 {
			blocks[start] = i
		} else {
			frees[start] = i
		}
		start = i
	}
	return blocks, frees
}

// Sort the keys of a map (used for files and free space)
func SortMap(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func MoveFiles(files []int, blocks map[int]int, frees map[int]int) []int {
    // Identify unique files in descending order
    uniqueFiles := make([]int, 0)
    fileSet := make(map[int]bool)
    for _, f := range files {
        if f != -1 && !fileSet[f] {
            uniqueFiles = append(uniqueFiles, f)
            fileSet[f] = true
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(uniqueFiles)))

    // Move each file
    for _, fileID := range uniqueFiles {
        // Find current file block
        fileStart, fileEnd := -1, -1
        for i := 0; i < len(files); i++ {
            if files[i] == fileID {
                if fileStart == -1 {
                    fileStart = i
                }
                fileEnd = i
            }
        }
        
        fileSize := fileEnd - fileStart + 1

        // Find leftmost free space
        for _, freeStart := range SortMap(frees) {
            freeEnd := frees[freeStart]
            freeSize := freeEnd - freeStart + 1

            if freeSize >= fileSize {
                // Move whole file
                copy(files[freeStart:freeStart+fileSize], files[fileStart:fileStart+fileSize])
                
                // Clear original file location
                for j := fileStart; j <= fileEnd; j++ {
                    files[j] = -1
                }
                break
            }
        }
    }

    return files
}

// Calculate the checksum based on file positions and IDs
func CalculateChecksum(files []int) int {
	checksum := 0
	for i, file := range files {
		if file != -1 {
			checksum += i * file
		}
	}
	return checksum
}

func main() {
	lines := ReadFileLines("sample.txt")
	fileList := GetFiles(lines[0])
	blocks, frees := GetMappedBlocks(fileList)

	// Move files and calculate checksum
	fmt.Println(fileList)
	movedFiles := MoveFiles(fileList, blocks, frees)
	checksum := CalculateChecksum(movedFiles)

	fmt.Println("Moved files:", movedFiles)
	fmt.Println("Checksum:", checksum)
}
