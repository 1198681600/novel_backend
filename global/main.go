package global

func init() {
	initConfig()
	initLogger()
	initDB()
	initRedis()
}
