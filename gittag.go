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
func CreateRemote(tagName string) error {
	cmd := exec.Command("git", "push", "origin", tagName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("推送标签到远程仓库失败: %v", err)
	}
	return nil
}

// DeleteLocal 删除本地标签
// @param tagName - 要删除的标签名称
// @return error - 如果删除过程中出现错误，返回相应的错误信息
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
func DeleteRemote(tagName string) error {
	cmd := exec.Command("git", "push", "origin", "--delete", tagName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("删除远程标签失败: %v", err)
	}
	return nil
}

// DeleteLocalAll 删除所有匹配指定模式的本地标签
// @param pattern - 标签匹配模式，例如："v1.*" 将匹配所有以 v1. 开头的标签
// @return error - 如果删除过程中出现错误，返回相应的错误信息
func DeleteLocalAll(pattern string) error {
	cmd := exec.Command("git", "tag", "-l", pattern)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取标签列表失败: %v", err)
	}

	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return nil
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
func DeleteRemoteAll(pattern string) error {
	cmd := exec.Command("git", "tag", "-l", pattern)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取标签列表失败: %v", err)
	}

	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return nil
	}

	for _, tag := range tags {
		if err := DeleteRemote(tag); err != nil {
			return fmt.Errorf("删除远程标签 %s 失败: %v", tag, err)
		}
	}
	return nil
}
