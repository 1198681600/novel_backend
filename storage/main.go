package storage

var (
	NovelStorage INovelStorage
	BookStorage  IBookStorage
)

func init() {
	NovelStorage = newNovelStorage()
	BookStorage = newBookStorage()
}
