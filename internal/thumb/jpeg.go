package thumb

import (
	"image"
	"path/filepath"
	"os/exec"
	"strconv"
	"fmt"

	"github.com/photoprism/photoprism/pkg/txt"
)

func Jpeg(srcFilename, jpgFilename string, size int) (img image.Image, err error) {
	resize := fmt.Sprintf("%dx%d>", size, size)
	cmd := exec.Command("convert", srcFilename + "[0]", "-resize", resize, "-quality", strconv.Itoa(JpegQuality), jpgFilename)
	if err = cmd.Run(); err != nil {
		log.Errorf("resample: failed to save %s", txt.Quote(filepath.Base(jpgFilename)))
		return nil, err
	}

	return nil, nil
}
