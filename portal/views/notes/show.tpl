<div class="col-md-9">
    <br/>
    <div class="markdown">
    {{.body}}
    </div>
</div>
<div class="col-md-3">
    <h4>aaa</h4>
    <hr/>
    <ul>
    {{range $k, $v := .notes}}
    <li>
    <a href="/notes/{{$k}}" target="_blank">{{$v}}</a>
    </li>
    {{end}}
    </ul>
</div>