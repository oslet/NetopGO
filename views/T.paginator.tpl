{{define "paginator"}}
<div class="col-sm-6">
	{{$IsSearch := .IsSearch}}
	{{$Keyword := .Keyword}}
	{{$Category := .Category}}
	{{$IsSlowLog := .IsSlowLog}}
	{{$Idc := .Idc}}
	{{$Group := .Group}}
	{{$Schema := .Schema}}
	<div class="dataTables_paginate paging_simple_numbers" id="editable_paginate">
	   <ul class="pagination" style="margin-top: 0; float: right">
	    <li>
	      <a href="/{{$Category}}/{{if $IsSearch}}search?page={{.paginator.firstpage}}&keyword={{$Keyword}}&idc={{$Idc}}&group={{$Group}}{{else if $IsSlowLog}}slowlog?name={{$Schema}}&page={{.paginator.firstpage}}{{else}}list?page={{.paginator.firstpage}}{{end}}" aria-label="Previous">
	        <span aria-hidden="true">&laquo;</span>
	      </a>
	    </li>
	    {{range $index,$page := .paginator.pages}}
	    <li {{if eq $.paginator.currpage $page }} class="active" {{end}}><a href="/{{$Category}}/{{if $IsSearch}}search?page={{$page}}&keyword={{$Keyword}}&idc={{$Idc}}&group={{$Group}}{{else if $IsSlowLog}}slowlog?name={{$Schema}}&page={{$page}}{{else}}list?page={{$page}}{{end}}">{{$page}}</a></li>
	    {{end}}
	    <li>
	      <a href="/{{$Category}}/{{if $IsSearch}}search?page={{.paginator.lastpage}}&keyword={{$Keyword}}&idc={{$Idc}}&group={{$Group}}{{else if $IsSlowLog}}slowlog?name={{$Schema}}&page={{.paginator.firstpage}}{{else}}list?page={{.paginator.lastpage}}{{end}}" aria-label="Next">
	        <span aria-hidden="true">&raquo;</span>
	      </a>
	    </li>
	  </ul>
	</div>
</div>
{{end}}