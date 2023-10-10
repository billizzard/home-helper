package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ByteToHuman(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func FileToIcon(file os.FileInfo) string {
	if file.IsDir() {
		return "folder"
	}

	switch strings.ToLower(filepath.Ext(file.Name())) {
	case ".pdf":
		return "file-earmark-pdf"
	case ".jpg", ".jpeg", ".png", ".bmp":
		return "file-image"
	case ".zip":
		return "file-earmark-zip"
	case ".mp3", ".wav", ".wma":
		return "file-earmark-music"
	case ".avi", ".mp4", ".wmv", ".mkv", ".mov":
		return "file-earmark-play"
	case ".txt":
		return "file-earmark-text"
	case ".doc", ".docx":
		return "file-earmark-word"
	case ".exe":
		return "filetype-exe"
	case ".html", ".htm":
		return "filetype-html"
	case ".yml", ".yaml":
		return "filetype-yml"
	default:
		return "filetype-earmark"
	}
}
