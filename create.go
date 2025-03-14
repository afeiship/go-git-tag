package gittag

import (
	"fmt"
	"os/exec"
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