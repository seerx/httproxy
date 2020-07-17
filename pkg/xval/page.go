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
</head>
<body style="text-align: center;">
	<h1>XVAL.CN</h1>
	<span>冀ICP备13021099号</span>
</body>
</html>
`

// AppHandle 主页
func AppHandle(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(htmlApp))
}

const htmlApp = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>XVAL.CN</title>
</head>
<body style="text-align: center;">
	<a href="syslog://go.app">App</a>
</body>
</html>
`
