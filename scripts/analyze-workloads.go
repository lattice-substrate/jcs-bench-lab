//go:build ignore

// analyze-workloads reads each JSON file in workloads/valid/ and emits a CSV
// characterizing the structure: total values, number values, number fraction,
// max nesting depth, and unique keys.
//
// Definition: Array and object containers are counted as values; their elements
// and members are counted separately.
//
// Usage: go run scripts/analyze-workloads.go [-workloads workloads/valid]
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	dir := flag.String("workloads", "workloads/valid", "path to valid workloads directory")
	flag.Parse()

	entries, err := os.ReadDir(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("name,bytes,total_values,num_values,num_fraction_pct,max_depth,keys")

	type result struct {
		name     string
		bytes    int
		total    int
		nums     int
		frac     float64
		depth    int
		keys     int
	}

	var results []result
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		path := filepath.Join(*dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: %v\n", e.Name(), err)
			continue
		}

		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: %v\n", e.Name(), err)
			continue
		}

		total, nums, depth, keys := walk(raw, 0)
		frac := 0.0
		if total > 0 {
			frac = float64(nums) / float64(total) * 100
		}
		name := strings.TrimSuffix(e.Name(), ".json")
		results = append(results, result{name, len(data), total, nums, frac, depth, keys})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].name < results[j].name
	})

	for _, r := range results {
		fmt.Printf("%s,%d,%d,%d,%.1f,%d,%d\n", r.name, r.bytes, r.total, r.nums, r.frac, r.depth, r.keys)
	}
}

// walk recursively counts values, numbers, max depth, and unique keys.
// The container itself (array/object) counts as one value.
func walk(v interface{}, depth int) (totalValues, numValues, maxDepth, uniqueKeys int) {
	switch val := v.(type) {
	case map[string]interface{}:
		totalValues = 1 // the object itself
		maxDepth = depth + 1
		keys := map[string]bool{}
		for k, child := range val {
			keys[k] = true
			ct, cn, cd, ck := walk(child, depth+1)
			totalValues += ct
			numValues += cn
			if cd > maxDepth {
				maxDepth = cd
			}
			uniqueKeys += ck
		}
		uniqueKeys += len(keys)
	case []interface{}:
		totalValues = 1 // the array itself
		maxDepth = depth + 1
		for _, child := range val {
			ct, cn, cd, ck := walk(child, depth+1)
			totalValues += ct
			numValues += cn
			if cd > maxDepth {
				maxDepth = cd
			}
			uniqueKeys += ck
		}
	case float64:
		totalValues = 1
		numValues = 1
		maxDepth = depth
	case string:
		totalValues = 1
		maxDepth = depth
	case bool:
		totalValues = 1
		maxDepth = depth
	case nil:
		totalValues = 1
		maxDepth = depth
	}
	return
}
