<!DOCTYPE html>
<html>
<head>
	<title>文件列表</title>
	<meta charset="utf-8">
	<script type="text/javascript" src="/static/jquery.min.js"></script>
	<script type="text/javascript" src="/static/bootstrap/js/bootstrap.min.js"></script>
	<link href="/static/bootstrap/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
	<link rel="stylesheet" type="text/css" href="/static/css/common.css">
</head>					
<body>
<div class="container col-xs-8 col-xs-offset-2">
	<div class="row header">
		<div class="col-xs-8 col-xs-offset-2">
			<h2>文件列表</h2>
		</div>
	</div>
	<div class="row margin-top">
		<div class="col-xs-12">
		    <div class="panel panel-default">
		        <div class="panel-heading">
		            Table
		        </div>
		        <!-- /.panel-heading -->
		        <div class="panel-body">
		            <div class="table-responsive">
		                <table class="table">
		                    <thead>
		                        <tr>
		                            <th>文件名</th>
		                            <th>日期</th>
		                            <th>删除</th>
		                        </tr>
		                    </thead>
		                    <tbody>
		                    {{range .}}
		                        <tr class="success">
		                            <td><a href="{{.Url}}">{{.Name}}</a></td>
		                            <td>{{.Created_at}}</td>
		                            <td><button fileId="{{.Id}}" fileName="{{.Name}}" onclick="DeleteFile(this)" type="button" class="btn btn-primary">Delete</button></td>
		                        </tr>
		                   {{end}}
		                    </tbody>
		                </table>
		            </div>
		            <!-- /.table-responsive -->
		        </div>
		        <!-- /.panel-body -->
		    </div>
		    <!-- /.panel -->
		</div>
	</div>
	<div class="row movedown" style="margin-left: 15%;">
		<div class="col-lg-12">
			<a href="/create_file"><button type="button" class="btn btn-primary col-lg-8">继续上传</button></a>
		</div>
	</div>
</div>
	<script type="text/javascript">
		function DeleteFile(dd){
			file = dd.getAttribute("fileName");
			id = dd.getAttribute("fileId")
			$.ajax({
				type: "post",
				url: "/file/delete?file="+file+"&fileId="+id,
				success:function(data){
					if(data.Code == "200"){
						window.location.reload()
					}
				}
			});
		}
	</script>
</body>
</html>