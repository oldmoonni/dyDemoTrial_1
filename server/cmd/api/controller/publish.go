package controller

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/dao"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/ffmpeg"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/handlers"
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/minio"
	"log"
	path2 "path"
	"strconv"
	"strings"
	"time"
)

type VideoListResponse struct {
	handlers.Response
	VideoList []handlers.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	duserlock, flag := dao.UserLockInfoByToken(token)
	if flag == false {
		c.JSON(consts.StatusOK, handlers.Response{
			StatusCode: 1,
			StatusMsg:  "wrong user information",
		})
	}

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(consts.StatusOK, handlers.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}

	//生成的id
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	id := int64(node.Generate())

	//保存到minio上
	title := c.PostForm("title")
	minio.UploadMinio(file, id, file.Filename, file.Size)
	path := fmt.Sprintf("http://172.22.105.178:9000/videos/%s/%s", strconv.FormatInt(id, 10), file.Filename)
	//封面上传到minio, 这个filename是包含.mp4后缀的
	img_name := strings.TrimSuffix(file.Filename, path2.Ext(file.Filename)) + ".jpeg"
	ffmpeg.GetFrame(path, img_name, id)

	author := duserlock.Id
	img_path := fmt.Sprintf("http://172.22.105.178:9000/videos/%s/%s", strconv.FormatInt(id, 10), img_name)
	timeUnix := time.Now().Unix()
	dao.VideoInsert(id, author, path, img_path, title, timeUnix)

	c.JSON(consts.StatusOK, handlers.Response{
		StatusCode: 0,
		StatusMsg:  file.Filename + " uploaded successfully",
	})
}