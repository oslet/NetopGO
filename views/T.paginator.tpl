{{define "paginator"}}
<div class="col-sm-6">
	{{$IsSearch:= .IsSearch}}
	{{$Keyword := .Keyword}}
	<div class="dataTables_paginate paging_simple_numbers" id="editable_paginate">
	   <ul class="pagination" style="margin-top: 0; float: right">
	    <li>
	      <a href="/user/{{if $IsSearch}}search?page={{.paginator.firstpage}}&keyword={{$Keyword}}{{else}}list?page={{.paginator.firstpage}}{{end}}" aria-label="Previous">
	        <span aria-hidden="true">&laquo;</span>
	      </a>
	    </li>
	    {{range $index,$page := .paginator.pages}}
	    <li {{if eq $.paginator.currpage $page }} class="active" {{end}}><a href="/user/{{if $IsSearch}}search?page={{$page}}&keyword={{$Keyword}}{{else}}list?page={{$page}}{{end}}">{{$page}}</a></li>
	    {{end}}
	    <li>
	      <a href="/user/{{if $IsSearch}}search?page={{.paginator.lastpage}}&keyword={{$Keyword}}{{else}}list?page={{.paginator.lastpage}}{{end}}" aria-label="Next">
	        <span aria-hidden="true">&raquo;</span>
	      </a>
	    </li>
	  </ul>
	</div>
</div>
{{end}}