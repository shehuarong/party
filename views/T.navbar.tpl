{{define "navbar"}}
<a class="navbar-brand" href="/">社区党员信息管理系统</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
		<li {{if .IsInformation}}class="active"{{end}}><a href="/information">信息管理</a></li>
		<li {{if .IsHonor}}class="active"{{end}}><a href="/honor">党员荣誉</a></li>
		<li {{if .IsLife}}class="active"{{end}}><a href="/life">工作生活</a></li>
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
		<li><a href="/login?exit=true">退出登录</a></li>
		{{else}}
		<li><a href="/login">登录</a></li>
		<li {{if .IsRegister}}class="active"{{end}}><a href="/register"> 注册</a></li>
		{{end}}

	</ul>
</div>
{{end}}