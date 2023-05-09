package downloader

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
)

type Downloader struct {
	VideoId string
}

func (d *Downloader) Download(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	client := youtube.Client{}
	fmt.Println(d)

	video, err := client.GetVideo(d.VideoId)
	if err != nil {
		return err
	}

	formats := video.Formats.WithAudioChannels()
	stream, size, err := client.GetStream(video, &formats[0])
	if err != nil {
		return err
	}
	fmt.Println("Size:", size/1024/1024, "MB")

	newUuid := uuid.New()
	file, err := os.Create("./downloads/" + newUuid.String() + ".mp4")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}
