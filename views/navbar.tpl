{{define "navbar"}}
    <div class="container">
      <h1>Seraphim Space</h1>
      <div class="navbar navbar-inverse">
          <ul class="nav navbar-nav">
             <li {{if .IsHome}}class="active"{{end}}><a href="/">Home</a></li>
             <li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a></li>
             <li {{if .IsTopics}}class="active"{{end}}><a href="/topic">Topics</a></li>
          </ul>

     <div class="pull-right">
      <ul class="nav navbar-nav">
       {{if .IsLogin}}
       {{else}}
       <li><a href="/register">register</a></li>
       {{end}}
       {{if .IsLogin}}
       <li><a href="/login?exit=true">{{.uname}}  quit</a></li>
       {{else}}
       <li><a href="/login">login</a></li>
       {{end}}
      </ul>
     </div>
    </div>
    </div>
{{end}}
