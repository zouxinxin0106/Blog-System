package main

import (
	"blog-system/database"
	"blog-system/handlers"
	"blog-system/models"
	"fmt"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== 题目1 & 3: 建表与钩子测试 ==========")

	var user models.User
	if err := db.First(&user,2).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查到用户: %s (ID=%d)\n", user.Name, user.ID)

	post1 := models.Post{Title: "Go语言入门3", Content: "GORM使用教程3", UserID: user.ID}
	db.Create(&post1)
	fmt.Printf("创建文章: %s (ID=%d) - 自动触发钩子更新用户文章数\n", post1.Title, post1.ID)

	db.First(&user, user.ID)
	fmt.Printf("用户 %s 的文章数: %d\n", user.Name, user.PostCount)

	comment1 := models.Comment{Content: "第一条评论3", UserID: user.ID, PostID: post1.ID}
	db.Create(&comment1)
	fmt.Printf("创建评论: %s (ID=%d)\n", comment1.Content, comment1.ID)

	comment2 := models.Comment{Content: "第二条评论3", UserID: user.ID, PostID: post1.ID}
	db.Create(&comment2)
	fmt.Printf("创建评论: %s (ID=%d)\n", comment2.Content, comment2.ID)

	comment3 := models.Comment{Content: "第三条评论3", UserID: user.ID, PostID: post1.ID}
	db.Create(&comment3)
	fmt.Printf("创建评论: %s (ID=%d)\n", comment3.Content, comment3.ID)

	fmt.Println("\n========== 题目2: 关联查询测试 ==========")

	// 调用 handlers.QueryUserPosts 查询用户所有文章及评论
	posts, err := handlers.QueryUserPosts(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n用户 %s 的所有文章及评论:\n", user.Name)
	for _, p := range posts {
		fmt.Printf("  文章: %s (评论数: %d)\n", p.Title, len(p.Comments))
		for _, c := range p.Comments {
			fmt.Printf("    - 评论: %s (by 用户%d)\n", c.Content, c.UserID)
		}
	}

	// 调用 handlers.QueryTopPostsByCommentCount 查询评论最多的文章
	topPosts, err := handlers.QueryTopPostsByCommentCount(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n评论最多的文章:\n")
	for i, p := range topPosts {
		fmt.Printf("  %d. %s (评论数: %d)\n", i+1, p.Title, len(p.Comments))
	}

	fmt.Println("\n========== 题目3: 删除评论钩子测试 ==========")
	fmt.Printf("删除评论前 - 文章 '%s' 状态: %s\n", post1.Title, post1.CommentStatus)

	db.Delete(&comment1)

	db.First(&post1, post1.ID)
	fmt.Printf("删除评论后 - 文章 '%s' 状态: %s\n", post1.Title, post1.CommentStatus)

	fmt.Println("\n========== 全部测试完成 ==========")
}