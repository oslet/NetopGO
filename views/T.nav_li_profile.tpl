{{define "nav_li_profile" }}
<li class="nav-header">
    <div class="dropdown profile-element">
        <span>
            <img alt="image" class="img-circle" width="48" height="48" src="/static/img/{{if eq .Role 1}}root.png{{ else if eq .Role 2}}a4.jpg{{else}}a6.jpg{{end}}" />
        </span>
        <a data-toggle="dropdown" class="dropdown-toggle" href="#">
            <span class="clear">
                <span class="block m-t-xs">
                    <strong class="font-bold">{{if eq .Role 1}}超级管理员{{else if eq .Role 2}}数据库管理员{{else}}来宾用户{{end}}<span style="color: #8095a8"></span></strong>
                </span>
                <span class="text-muted text-xs block">
                    {{.Uname}} <b class="caret"></b>
                </span>
            </span>
        </a>
        <ul class="dropdown-menu animated fadeInRight m-t-xs">
            <li><a href="/user/detail?id={{.Id}}">个人信息</a></li>
            <li><a href="/user/reset_password?id={{.Id}}&action=view">修改密码</a></li>
            <li><a href="/logout">注销</a></li>
        </ul>
    </div>

    <div class="logo-element">
        JS+
    </div>
</li>
{{end}}