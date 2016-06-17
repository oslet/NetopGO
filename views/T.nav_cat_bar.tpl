{{define "nav_cat_bar"}}
<div class="row wrapper border-bottom white-bg page-heading">
    <div class="col-sm-10">
        <h2></h2>
        <ol class="breadcrumb">
            <li>
                <a href="/netopgo">仪表盘</a>
            </li>
            <li>
                <a href="{{.Href}}">{{.Path1}}</a>
            </li>
            <li class="active">
                <strong>{{.Path2}}</strong>
            </li>
        </ol>
    </div>
    <div class="col-sm-2">
    </div>
</div>
{{end}}