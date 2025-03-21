package gittag

import (
	"fmt"
	"os/exec"
)

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