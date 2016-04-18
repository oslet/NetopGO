{{define "paginator"}}
<div class="col-sm-6">
	<div class="dataTables_paginate paging_simple_numbers" id="editable_paginate">
	   <ul class="pagination" style="margin-top: 0; float: right">
	    <li>
	      <a href="#" aria-label="Previous">
	        <span aria-hidden="true">&laquo;</span>
	      </a>
	    </li>
	    {{range $index,$page := .paginator.pages}}
	    <li {{if eq $.paginator.currpage $page }} class="active" {{end}}><a href="/user/list?page={{$page}}">{{$page}}</a></li>
	    {{end}}
	    <li>
	      <a href="#" aria-label="Next">
	        <span aria-hidden="true">&raquo;</span>
	      </a>
	    </li>
	  </ul>
	</div>
</div>
{{end}}