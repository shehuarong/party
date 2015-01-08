{{template "header"}}
    	<title>登录 - 社区党员信息管理系统</title>
  	</head>
	
	<body>
		<div class="container" style="width: 500px;">
			<form class="form-horizontal" method="post" action="/login">
			  <div class="form-group">
			    <label class="col-lg-4 control-label">请输入昵称：</label>
			    <div class="col-lg-6">
			      <input id="uname" class="form-control" name="uname" placeholder="Account">
			    </div>
			  </div>

			  <div class="form-group">
			    <label class="col-lg-4 control-label">请输入密码：</label>
			    <div class="col-lg-6">
			      <input id="pwd" type="password" class="form-control" name="pwd" placeholder="Password">
			    </div>
			  </div>



			  <div class="form-group">
			    <div class="col-lg-offset-2 col-lg-10">
			    	<a href="/login"><button type="submit" class="btn btn-default" onclick="return checkInput();">登录</button></a>
			      	<button class="btn btn-default" onclick="return backToHome();">返回</button>

			  
					
			      	<script type="text/javascript">
				      	function backToHome() {
				      		window.location.href = "/";
				      		return false;
				      	}

				      	function checkInput() {
				      		var uname = document.getElementById("uname");
				      		if (uname.value.length == 0) {
				      			alert("请输入用户名");
				      			return false;
				      		}

				      		var pwd = document.getElementById("pwd");
				      		if (pwd.value.length == 0) {
				      			alert("请输入密码");
				      			return false;
				      		}
				      	}
			      	</script>
			    </div>
			  </div>
			</form>
		</div>
	</body>
</html>