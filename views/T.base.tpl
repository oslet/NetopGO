{{define "page_top"}}
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>NetopGO | 听云NetopGO系统</title>

    <link rel="shortcut icon" href="/static/img/facio.ico" type="image/x-icon">
    {{template "link_css"}}
    {{template "head_script"}}
</head>

<body>

    <div id="wrapper">
    	{{template "nav"}}
        <div id="page-wrapper" class="gray-bg">
            <div class="row border-bottom">
            	{{template "nav_bar_header"}}
            </div>

{{end}}