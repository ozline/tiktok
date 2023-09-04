package sensitive_words

import (
	"bufio"
	"io"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/ozline/tiktok/pkg/utils"
)

var St *utils.SensitiveTrie

func Init(path string) {
	St = utils.NewSensitiveTrie()
	fileHandle, err := os.OpenFile(path+"/words.txt", os.O_RDONLY, 0666)
	if err != nil {
		klog.Warn(err)
		return
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	var words []string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		words = append(words, string(line))
	}

	St.AddWords(words)
}
