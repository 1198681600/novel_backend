package storage

var (
	NovelStorage INovelStorage
)

func init() {
	NovelStorage = newNovelStorage()
}
