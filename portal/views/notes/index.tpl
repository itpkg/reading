<br/>
<div class="row">
    {{range $k, $v := .notes}}
    <div class="col-md-3">
        <a href="/notes/{{$k}}" target="_blank" class="thumbnail">{{$v}}</a>
    </div>
    {{end}}

</div>