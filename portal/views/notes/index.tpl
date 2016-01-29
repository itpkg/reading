<br/>
<div class="row">
    {{range $k, $v := .items}}
    <div class="col-md-3">
        <a href="/notes/{{$k}}" target="_blank" class="thumbnail">{{$v}}</a>
    </div>
    {{end}}

</div>