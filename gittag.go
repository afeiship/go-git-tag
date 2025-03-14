package gittag

import (
	"fmt"
	"os/exec"
	"strings"
)

// CreateLocal 创建一个本地 Git 标签
// @param tagName - 标签名称，例如："v1.0.0"
// @param message - 标签信息（可选），如果不提供则使用默认格式："Release <tagName>"
// @return error - 如果创建过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Create a tag with default message
//	err := gittag.CreateLocal("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create a tag with custom message
//	err = gittag.CreateLocal("v2.0.0", "Major release with new features")
//	if err != nil {
//		log.Fatal(err)
//	}
func CreateLocal(tagName string, message ...string) error {
	tagMessage := "chore(release): " + tagName
	if len(message) > 0 && message[0] != "" {
		tagMessage = message[0]
	}
	cmd := exec.Command("git", "tag", "-a", tagName, "-m", tagMessage)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("创建本地标签失败: %v", err)
	}
	return nil
}

// CreateRemote 将本地标签推送到远程仓库
// @param tagName - 标签名称，例如："v1.0.0"
// @return error - 如果推送过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Push a local tag to remote repository
//	err := gittag.CreateRemote("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
func CreateRemote(tagName string) error {
	cmd := exec.Command("git", "push", "origin", tagName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("推送标签到远程仓库失败: %v", err)
	}
	return nil
}

// CreateTag creates a tag both locally and remotely in one operation
// @param tagName - 标签名称，例如："v1.0.0"
// @param message - 标签信息（可选），如果不提供则使用默认格式："chore(release): <tagName>"
// @return error - 如果创建过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Create a tag with default message and push to remote
//	err := gittag.CreateTag("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create a tag with custom message and push to remote
//	err = gittag.CreateTag("v2.0.0", "Major release with new features")
//	if err != nil {
//		log.Fatal(err)
//	}
func CreateTag(tagName string, message ...string) error {
	if err := CreateLocal(tagName, message...); err != nil {
		return err
	}
	if err := CreateRemote(tagName); err != nil {
		return err
	}
	return nil
}

// DeleteLocal 删除本地标签
// @param tagName - 要删除的标签名称
// @return error - 如果删除过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Delete a specific local tag
//	err := gittag.DeleteLocal("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
func DeleteLocal(tagName string) error {
	cmd := exec.Command("git", "tag", "-d", tagName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("删除本地标签失败: %v", err)
	}
	return nil
}

// DeleteRemote 删除远程仓库中的标签
// @param tagName - 要删除的标签名称
// @return error - 如果删除过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Delete a tag from remote repository
//	err := gittag.DeleteRemote("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
func DeleteRemote(tagName string) error {
	cmd := exec.Command("git", "push", "origin", "--delete", tagName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("删除远程标签失败: %v", err)
	}
	return nil
}

// DeleteTag deletes a tag both locally and remotely in one operation
// @param tagName - 要删除的标签名称
// @return error - 如果删除过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Delete a tag both locally and from remote repository
//	err := gittag.DeleteTag("v1.0.0")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Delete multiple tags in a loop
//	tags := []string{"v1.0.0", "v1.1.0", "v1.2.0"}
//	for _, tag := range tags {
//		if err := gittag.DeleteTag(tag); err != nil {
//			log.Printf("Failed to delete tag %s: %v", tag, err)
//			continue
//		}
//	}
func DeleteTag(tagName string) error {
	if err := DeleteLocal(tagName); err != nil {
		return err
	}
	if err := DeleteRemote(tagName); err != nil {
		return err
	}
	return nil
}

// DeleteLocalAll 删除所有匹配指定模式的本地标签
// @param pattern - 标签匹配模式，例如："v1.*" 将匹配所有以 v1. 开头的标签
// @return error - 如果删除过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Delete all local tags starting with "v1."
//	err := gittag.DeleteLocalAll("v1.*")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Delete all beta tags
//	err = gittag.DeleteLocalAll("*-beta.*")
//	if err != nil {
//		log.Fatal(err)
//	}
func DeleteLocalAll(pattern string) error {
	tags, err := FindMany(pattern)
	if err != nil {
		return nil // 如果没有找到标签，直接返回
	}

	for _, tag := range tags {
		if err := DeleteLocal(tag); err != nil {
			return fmt.Errorf("删除标签 %s 失败: %v", tag, err)
		}
	}
	return nil
}

// DeleteRemoteAll 删除所有匹配指定模式的远程标签
// @param pattern - 标签匹配模式，例如："v1.*" 将匹配所有以 v1. 开头的标签
// @return error - 如果删除过程中出现错误，返回相应的错误信息
//
// Example:
//
//	// Delete all remote tags starting with "v1."
//	err := gittag.DeleteRemoteAll("v1.*")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Delete all remote release candidate tags
//	err = gittag.DeleteRemoteAll("*-rc.*")
//	if err != nil {
//		log.Fatal(err)
//	}
func DeleteRemoteAll(pattern string) error {
	tags, err := FindMany(pattern)
	if err != nil {
		return nil // 如果没有找到标签，直接返回
	}

	for _, tag := range tags {
		if err := DeleteRemote(tag); err != nil {
			return fmt.Errorf("删除远程标签 %s 失败: %v", tag, err)
		}
	}
	return nil
}

// DeleteAllTags deletes all tags both locally and remotely in one operation.
// This is a convenience function that combines DeleteLocalAll and DeleteRemoteAll
// to remove all Git tags from both the local repository and the remote repository.
// Use this function with caution as it will remove ALL tags.
//
// @return error - If any error occurs during the deletion process
//
// Example:
//
//	// Delete all tags (both local and remote)
//	 err := gittag.DeleteAllTags()
//	 if err != nil {
//		 log.Printf("Failed to delete all tags: %v", err)
//		 return err
//	 }
//
//	 // Using in a cleanup function
//	 func cleanupTags() error {
//		 fmt.Println("Removing all Git tags...")
//		 if err := gittag.DeleteAllTags(); err != nil {
//			 return fmt.Errorf("failed to cleanup tags: %v", err)
//		 }
//		 fmt.Println("All tags have been removed successfully")
//		 return nil
//	 }
func DeleteAllTags() error {
	if err := DeleteLocalAll("*"); err != nil {
		return err
	}
	if err := DeleteRemoteAll("*"); err != nil {
		return err
	}
	return nil
}

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
