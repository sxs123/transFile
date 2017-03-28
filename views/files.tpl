<!DOCTYPE html>
<html>
<head>
	<title>文件列表</title>
	<meta charset="utf-8">
</head>
<body>
	{{range .}}
		<div class="row">
			<div class="item">
				<a href="{{.Url}}">{{.Name}}</a>
				<div class="date">{{.Created_at}}</div>
			</div>
		</div>
	{{end}}
</body>
</html>