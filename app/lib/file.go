package lib

func GetContentType()  {
	gContentType := make(map[string][]string)
	// 图片
	gContentType["image/gif"] = []string{"gif", "img"}
	gContentType["image/jpeg"] = []string{"jpg", "img"}
	gContentType["image/jpg"] = []string{"jpg", "img"}
	gContentType["image/png"] = []string{"png", "img"}
	gContentType["image/x-png"] = []string{"png", "img"}
	gContentType["image/x-png"] = []string{"png", "img"}
	gContentType["image/bmp"] = []string{"bmp", "img"}

	// 视频
	gContentType["video/mp4"] = []string{"mp4", "video"}
	gContentType["video/x-matroska"] = []string{"mkv", "video"}
	gContentType["video/x-msvideo"] = []string{"avi", "video"}
	gContentType["application/vnd.rn-realmedia-vbr"] = []string{"rmvb", "video"}
	gContentType["video/3gpp"] = []string{"3gp", "video"}
	gContentType["video/x-flv"] = []string{"flv", "video"}
	gContentType["video/mpeg"] = []string{"mpg", "video"}
	gContentType["video/quicktime"] = []string{"mov", "video"}
	gContentType["video/x-ms-wmv"] = []string{"wmv", "video"}
}

