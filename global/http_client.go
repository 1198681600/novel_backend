package global

import "github.com/imroc/req/v3"

var (
	ReqClient *req.Client
)

func init() {
	// TODO enable dev mod only on dev
	ReqClient = req.C() //.DevMode()
}
