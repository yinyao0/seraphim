# seraphim
这是一个基于beego实现的简单网页
部署nginx

server {
   listen 80;
   server_name localhost;
   charset utf-8;
   location / {
      proxy_pass http://127.0.0.1:8080;
  }
}
