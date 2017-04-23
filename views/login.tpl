<!DOCTYPE html>
<html>
<head>
	<title>文件传输系统</title>
	<meta charset="utf-8"><script type="text/javascript" src="/static/jquery.min.js"></script>
	<script type="text/javascript" src="/static/bootstrap/js/bootstrap.min.js"></script>
	<link href="/static/bootstrap/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
	<link rel="stylesheet" type="text/css" href="/static/css/common.css"></head>
</head>
<body>
	<div class="container">
		<div class="row movedown-lg moveup-lg">
			<div class="col-md-4 col-md-offset-4 title">
				登录文件传输系统
			</div>
		</div>
        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <div class="login-panel panel panel-default">
                    <div class="panel-heading">
                        <h3 class="panel-title">登录</h3>
                    </div>
                    <div class="panel-body">
                        <form role="form" action="http://127.0.0.1:3000/login" method="post">
                            <fieldset>
	                           	<form >
	                                <div class="form-group">
	                                    <input class="form-control" placeholder="E-mail" name="email" id="email" type="email" autofocus>
	                                </div>
	                                <div class="form-group">
	                                    <input class="form-control" placeholder="Password" name="password" id="password" type="password" value="">
	                                </div>
	                                <div class="checkbox">
	                                    <label>
	                                        <input name="remember" type="checkbox" value="Remember Me">记住我
	                                    </label>
	                                </div>
	                                <!-- Change this to a button or input when using this as a form -->
	                                <input type="submit" class="btn btn-lg btn-success btn-block" name="submit" value="登录">
	                            </form>
                            </fieldset>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <script type="text/javascript">
    	function login(){
    		var email = $("#email").val();
    		var pwd = $("#password").val();
    		var data = new Object();
    		data.email = email;
    		data.pwd = pwd;
    		var url = "/login";
    		$.ajax({
    			type: "post",
    			url: url,
    			data: data,
    			success:function(data){

    			}
    		});
    	}
    </script>
</body>
</html>