package xval

import "net/http"

// PageHandle 主页
func PageHandle(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(html))
}

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>XVAL.CN</title>
    <link rel="stylesheet" href="https://unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css">
</head>
<body style="text-align: center;">
	<h1>个人网站</h1>
	<span>冀ICP备13021099号</span>
</body>
</html>
`