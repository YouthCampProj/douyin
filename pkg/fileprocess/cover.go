package fileprocess

import (
	"github.com/YouthCampProj/douyin/pkg/config"
	"os/exec"
)

func GetCoverFromLocal(videoPath string, coverPath string) (string, error) {
	// capture first video frame as jpg
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", coverPath)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	coverURL := config.Conf.Site.Domain + "/" + coverPath
	return coverURL, nil
}
