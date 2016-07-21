{{define "page_buttom"}}
            {{template "footer"}}
        </div>
    </div>

</body>
<script type="text/javascript">
$(document).ready(function(){
        var msg_text=$('#alert_msg').text();
        if(msg_text!=''){
                $("#alert_msg").show();
        }
});

</script>
	{{template "foot_script"}}
</html>
{{end}}