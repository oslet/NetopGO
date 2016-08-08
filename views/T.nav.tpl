{{define "nav"}}
<nav class="navbar-default navbar-static-side" role="navigation">
    <div class="sidebar-collapse">
        <ul class="nav" id="side-menu">
            {{template "nav_li_profile" .}}
            <li id="index">
               <a href="/netopgo"><i class="fa fa-dashboard"></i> <span class="nav-label">仪表盘</span><span class="label label-info pull-right"></span></a>
            </li>
            <li id="user">
                <a href="javascript:void(0)"><i class="fa fa-group"></i> <span class="nav-label">用户管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="group"><a href="/user/list">用户列表</a></li>
                </ul>
            </li>
            <li id="host">
                <a><i class="fa fa-inbox"></i> <span class="nav-label">主机管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="idc"> <a href="/host/list">主机列表</a></li>
                    <li class="idc"> <a href="/group/list">业务组列表</a></li>
                    <li class="idc"> <a href="/line/list">线路列表</a></li>
                    <li class="idc"> <a href="/system/list">系统列表</a></li>
                </ul>
            </li>
            <li id="db">
                <a><i class="fa fa-gears"></i> <span class="nav-label">DB管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">                
                    <li class="rule"><a href="/db/list">实例列表</a></li>
                    <li class="rule"><a href="/db/query">查询窗口</a></li>
                    <li class="rule"><a href="/schema/list">数据源列表</a></li>                    
                </ul>
            </li>
            <li id="record">
                <a href="javascript:void(0)"><i class="fa fa-edit"></i> <span class="nav-label">系统发布</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="download"><a href="/workorder/my/list">应用工单</a></li>
                    <li class="download"><a href="/workorder/mydb/list">数据库工单</a></li>
                    <li class="download"><a href="/workorder/app">提交应用工单</a></li>
                    <li class="download"><a href="/workorder/db">提交数据库工单</a></li>
                    <li class="download"><a href="/record/fault/list">故障记录</a></li>
                    <li class="download"><a href="/record/quest/list">问题记录</a></li>
                </ul>
            </li>      
            <!--      
            <li id="record">
                <a href="#"><i class="fa fa-edit"></i> <span class="nav-label">记录管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="upload"><a href="/record/app/list">系统升级</a></li>
                    <li class="download"><a href="/record/db/list">DB升级</a></li>
                    <li class="download"><a href="#">知识库</a></li>
                </ul>
            </li>
            -->
            <li id="log">
               <a href="/audit/list"><i class="fa fa-files-o"></i> <span class="nav-label">日志审计</span><span class="label label-info pull-right"></span></a>
            </li>
            <li id="report">
                <a href="javascript:void(0)"><i class="fa fa-edit"></i> <span class="nav-label">报表管理</span><span class="fa arrow"></span></a>
                <ul class="nav nav-second-level">
                    <li class="report_host"><a href="/report/host/list/">新增主机</a></li>
                    <li class="report_host"><a href="/report/recycle/list/">回收主机</a></li>
                </ul>
            </li>  
            <li class="special_link">
                <a href="http://www.platenogroup.com" target="_blank"><i class="fa fa-database"></i> <span class="nav-label">访问官网</span></a>
            </li>
        </ul>
        </div>
</nav>
{{end}}