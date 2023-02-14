package minio

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadMinio(fileheader *multipart.FileHeader, id int64, name string, size int64) {
	endpoint := "192.168.64.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false //注意没有安装证书的填false

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("初始化完成！")
	// 链接一个叫videos的存储桶。
	bucketName := "videos"
	//下面注释部分是创建一个叫videos的存储桶。
	//bucketName := "videos"
	//location := "cn-north-1"
	//err = minioClient.MakeBucket(bucketName, location)
	//if err != nil {
	//	log.Println("创建bucket失败！")
	//	// 检查存储桶是否已经存在。
	//	exists, err := minioClient.BucketExists(bucketName)
	//	if err == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Println("打印失败！")
	//		log.Fatalln(err)
	//	}
	//}
	//log.Printf("Successfully created %s\n", bucketName)

	// 上传一个mp4文件。
	objectName := strconv.FormatInt(id, 10) + "/" + name
	//filePath := "C:\\Users\\胡旭旭\\Pictures\\Camera Roll\\dy2.mp4"
	contentType := "video/mp4"

	// 使用FPutObject上传一个mp4文件。
	file, err := fileheader.Open()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := minioClient.PutObject(bucketName, objectName, file, size, minio.PutObjectOptions{ContentType: contentType})
	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func UploadMinioJpg(id int64, name string, filePath string) {
	endpoint := "192.168.64.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false //注意没有安装证书的填false

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("初始化完成！")
	// 链接一个叫videos的存储桶。
	bucketName := "videos"

	objectName := strconv.FormatInt(id, 10) + "/" + name
	//filePath := "C:\\Users\\胡旭旭\\Pictures\\Camera Roll\\dy2.mp4"
	contentType := "image/jpeg"

	//n, err := minioClient.PutObject(bucketName, objectName, file, size, minio.PutObjectOptions{ContentType: contentType})
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	//在本地的封面文件上传到minio后，将其删除
	os.Remove(filePath)
	if err != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}

}
