package gittag

import (
	"fmt"
	"os/exec"
	"strings"
)

// FindOne searches for and returns a single Git tag matching the given pattern.
// @param pattern - The pattern to match tags against, e.g., "v1.*" matches all tags starting with "v1."
// @return (string, error) - Returns the first matching tag and any error that occurred
//
// Example:
//
//	// Find the first tag matching a pattern
//	tag, err := gittag.FindOne("v1.*")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Found tag: %s\n", tag)
func FindOne(pattern string) (string, error) {
	cmd := exec.Command("git", "tag", "-l", pattern)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("查找标签失败: %v", err)
	}

	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return "", fmt.Errorf("未找到匹配的标签")
	}

	return tags[0], nil
}

// FindMany searches for and returns all Git tags matching the given pattern.
// @param pattern - The pattern to match tags against, e.g., "v1.*" matches all tags starting with "v1."
// @return ([]string, error) - Returns all matching tags and any error that occurred
//
// Example:
//
//	// Find all tags matching a pattern
//	tags, err := gittag.FindMany("v1.*")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, tag := range tags {
//		fmt.Printf("Found tag: %s\n", tag)
//	}
func FindMany(pattern string) ([]string, error) {
	cmd := exec.Command("git", "tag", "-l", pattern)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("查找标签失败: %v", err)
	}

	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return nil, fmt.Errorf("未找到匹配的标签")
	}

	return tags, nil
}