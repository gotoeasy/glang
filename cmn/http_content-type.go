package cmn

import "mime"

// 按扩展名如 .html 取ContentType，优先取自定义再取mime，都取不到返回"application/octet-stream"
func ContentType(dotExt string, contentTypeMap map[string]string) string {
	contentType := contentTypeMap[dotExt]
	if contentType == "" {
		contentType = mime.TypeByExtension(dotExt)
	}
	if contentType == "" {
		contentType = "application/octet-stream" // 无法检测到 ContentType 时的默认值
	}
	return contentType
}
