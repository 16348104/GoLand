package _0Switch

func readByExt(ext string) {
	if ext == "json" {
		println("read json file")
	} else if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif" {
		println("read image file")
	} else if ext == "txt" || ext == "md" {
		println("read text file")
	} else if ext == "yml" || ext == "yaml" {
		println("read yaml file")
	} else if ext == "ini" {
		println("read ini file")
	} else {
		println("unsupported file extension:", ext)
	}
}
