package maps

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"
)

//go:embed data/*
var mapData embed.FS

type MapMask [MapY][MapX]int

var MapMasks = map[string]MapMask{}

// init populates MapMasks with data from the data/ directory.
func init() {
	entries, err := fs.Glob(mapData, "data/*")
	if err != nil {
		log.Fatalln("bad map data:", err)
	}
	for _, entry := range entries {
		name, mask, err := fileToMask(mapData, entry)
		if err != nil {
			log.Fatalln("bad map data:", err)
		}
		MapMasks[name] = mask
	}
}

// fileToMask converts a file in the data/ directory into a MapMask.
func fileToMask(fsys fs.FS, filename string) (name string, mask MapMask, err error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return "", mask, fmt.Errorf("open failed: %w", err)
	}
	defer f.Close()
	hasName := false
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	y := 0
	for scanner.Scan() {
		if !hasName {
			name = scanner.Text()
			hasName = true
			continue
		}
		if scanner.Text() == "" {
			continue
		}
		for x, char := range strings.Split(scanner.Text(), "") {
			result := 0
			if char == "1" {
				result = 1
			}
			mask[y][x] = result
		}
		y += 1
	}
	if err := scanner.Err(); err != nil {
		return "", mask, fmt.Errorf("scan failed: %w", err)
	}
	return name, mask, nil
}
