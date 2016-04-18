{{define "nav"}}
<nav class="navbar-default navbar-static-side" role="navigation">
    <div class="sidebar-collapse">
        <ul class="nav" id="side-menu">
            {{template "nav_li_profile" .}}
            <li id="index">
               <a href="/"><i class="fa fa-dashboard"></i> <span class="nav-label">仪表盘</span><span class="label label-info pull-right"></span></a>
            </li>
            <li id="user">
                <a href="#"><i class="fa fa-group"></i> <span class="nav-label">用户管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="group"><a href="/user/list?page=">用户列表</a></li>
                </ul>
            </li>
            <li id="host">
                <a><i class="fa fa-inbox"></i> <span class="nav-label">主机管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="idc"> <a href="#">主机列表</a></li>
                </ul>
            </li>
            <li id="db">
                <a href="#"><i class="fa fa-edit"></i> <span class="nav-label">DB管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="rule">
                        <a href="#">实例列表</a>
                    </li>
                    <li class="rule">
                        <a href="#">查询窗口</a>
                    </li>
                    <li class="rule">
                        <a href="#">分区监控</a>
                    </li>
                </ul>
            </li>
            <li id="record">
                <a href="#"><i class="fa fa-edit"></i> <span class="nav-label">记录管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="upload"><a href="#">系统升级</a></li>
                    <li class="download"><a href="#">DB升级</a></li>
                    <li class="download"><a href="#">故障记录</a></li>
                    <li class="download"><a href="#">知识库</a></li>
                </ul>
            </li>
            <li id="log">
               <a href="#"><i class="fa fa-files-o"></i> <span class="nav-label">日志审计</span><span class="label label-info pull-right"></span></a>
            </li>
            <li class="special_link">
                <a href="http://www.tingyun.com" target="_blank"><i class="fa fa-database"></i> <span class="nav-label">访问官网</span></a>
            </li>
        </ul>

    </div>
</nav>
{{end}}