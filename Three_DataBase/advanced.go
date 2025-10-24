package main

import (
	"encoding/json"
	"fmt"
	"go-base-practice/utils"
	"reflect"

	"gorm.io/gorm"
)

// User 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
//
//	使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//	编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	ID        uint64 `gorm:"index;"`
	Name      string
	PostCount int
	PostArr   []Post `gorm:"foreigner:UserId;references:ID"`
}

type Post struct {
	ID         uint64 `db:"id"`
	Title      string `db:"title"`
	Status     string
	UserId     uint64
	CommentArr []Comment `gorm:"foreigner:PostId;references:ID"`
}

type Comment struct {
	ID      uint64 `db:"id"`
	Content string `db:"content"`
	PostId  uint64
}

func main() {
	// 题目1：模型定义
	//createInit()

	// 题目2：关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	//  编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var arrUsers []User
	utils.DB().Debug().Preload("PostArr.CommentArr").Where("Name = ?", "小王").Find(&arrUsers)
	for i := range arrUsers {
		js, _ := json.MarshalIndent(arrUsers[i], "", " ")
		fmt.Println(string(js))
	}

	//  编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var posts Post
	var postId uint64
	err := utils.DB().Table("comments").Select("post_id").Group("post_id").Order("count(1) Desc").Limit(1).Scan(&postId).Error
	if err != nil {
		fmt.Println("查询错误", err)
		return
	}
	err = utils.DB().Debug().Model(&posts).Where("id = ?", postId).Scan(&posts).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< 题目3 ")
	// 题目3：继续使用博客系统的模型。
	// 要求 ：
	//  为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	var postOne = Post{
		Title:  "文章P",
		Status: "0",
		UserId: 5,
	}
	utils.DB().Create(&postOne)
	//  为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var c = Comment{
		ID: 1,
	}
	utils.DB().Debug().Delete(&c)
}

// BeforeCreate Post 模型添加一个钩子函数，在文章创建时自动更新用户文章数量统计字段。
func (post *Post) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("创建更新文章用户数量- 钩子函数 -begin")
	fmt.Println("创建更新文章用户ID:", post.UserId)
	var num int
	// 判断文章类型 是数组还是单体
	if reflect.ValueOf(post).Kind() == reflect.Slice {
		fmt.Println("是 数组~~~~~~~~~~~~~~~~")
		// 数组
		postArr := reflect.ValueOf(post)
		num = postArr.Len()
	} else {
		fmt.Println("是 单体~~~~~~~~~~~~~~~~")
		num = 1
	}
	fmt.Println("需要更新用户文章的数量为:", num)
	// 查用户信息并更新
	err = db.Debug().Model(&User{}).Where("id = ?", post.UserId).UpdateColumn("PostCount", gorm.Expr("post_count + ?", num)).Error
	if err != nil {
		fmt.Println("更新用户文章数量失败:", err)
	}
	return err
}

// BeforeDelete Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) BeforeDelete(db *gorm.DB) (err error) {
	fmt.Println("删除评论主键id为：", comment.ID)

	var postId uint64
	var count int64
	var status string
	// 根据主键id查该文章的评论有几条 和 文章中评论状态是什么? 校验
	err = db.Debug().Model(&Comment{}).Select("PostId").Where("ID = ?", comment.ID).Find(&postId).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除评论的文章id为：", postId)

	// 查该文章评论数有几条
	err = db.Debug().Model(&Comment{}).Where("post_id = ?", postId).Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("该文章的评论数量为:", count)

	// 查对应文章状态是什么
	err = db.Debug().Model(&Post{}).Select("Status").Where("ID = ?", postId).Find(&status).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("该文章状态为:", status)

	// 更新文章状态
	if (count == 0 || count == 1) || status == "0" {
		fmt.Println("无评论,需要更新文章状态为无,文章id:", postId)
		db.Debug().Model(&Post{}).Where("id = ?", postId).Update("status", "无评论")
	}
	return
}

var usArr = []User{
	{
		Name:      "小张",
		PostCount: 1,
		PostArr: []Post{
			{
				Title:  "文章A",
				Status: "1", // 0 无评论, 1 有评论
				CommentArr: []Comment{
					{
						Content: "文章A真好",
					},
				},
			},
		},
	},
	{
		Name:      "小王",
		PostCount: 2,
		PostArr: []Post{
			{
				Title:      "文章B",
				Status:     "0", // 0 无评论, 1 有评论
				CommentArr: []Comment{},
			},
			{
				Title:  "文章C",
				Status: "1", // 0 无评论, 1 有评论
				CommentArr: []Comment{
					{
						Content: "文章C真好",
					},
					{
						Content: "文章C真帅",
					},
					{
						Content: "文章C真美",
					},
				},
			},
		}},
}

func createInit() {
	utils.DB().AutoMigrate(&User{}, &Post{}, &Comment{})
	utils.DB().Create(&usArr)
}
