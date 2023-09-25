package upload

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
文件上传
*/

type SysFile struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT"`
	CreateAt     int64  `gorm:"column:create_at"`
	CreateBy     string `gorm:"column:create_by"`
	DeletedAt    int64  `gorm:"column:deleted_at"` //删除时间
	OriginName   string `gorm:"column:origin_name"`
	RelativePath string `gorm:"column:relative_path"`
	Size         int64  `gorm:"column:size"`
	FileType     string `gorm:"column:file_type"`
	SeeType      string `gorm:"column:see_type"`
}

// 第一步上传文件到cos

// 第二步生成相对路径

func UploadFile(c *gin.Context) {

	// 获取到上传的文件
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// // 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
	}
	// 上传到cos

	// 入数据库

	// 给出response
}
