package ffmpeg

import (
	"github.com/trial_1/dyDemoTrial_1/server/cmd/api/minio"
	"log"
	"os/exec"
)

// 获取视频中一帧的图片
func GetFrame(url string, name string, id int64) {
	cmd := exec.Command("ffmpeg",
		"-i", url,
		"-ss", "00:00:00",
		"-frames:v", "1",
		"./images/"+name)
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to extract frame:", err)
	}

	minio.UploadMinioJpg(id, name, "./images/"+name)

}
