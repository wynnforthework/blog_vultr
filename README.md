# 博客
> 在vultr上通过hexo建立博客

在谷歌浏览器上访问时，总是在域名前自动加上https，导致无法无法发问很蛋疼。

# 计划
- [-] 让服务器支持https
	- [iris listen tls] (https://iris-go.com/v8/recipe#Listen%20Tls76)
	- [一个提供免费HTTPS证书申请的网站](https://freessl.org/)
	- [centOS系统生成SSL数字证书](https://www.cnblogs.com/CKiller/p/5355039.html)
- [-] 安装hexo
- [-] 配置next主题
- [ ] [next主题个性化](https://segmentfault.com/a/1190000009544924)
- [ ] 增加简历页面
- [ ] 增加前端页面写markdown
- [ ] 用go写登录页面
- [ ] 用go实现登录之后，在前端页面写markdown，并发布文章。

# 遇到的问题

1. SSL证书
reessl是一个很好的网站，生成证书很方便，但是语焉不详，生成完之后我一度一招莫展。  

生成完证书后，页面出现了三个文本框，分别写着“CA证书”、“证书”和“私钥”，然后点击证书下载，下载到两个文件“full_chain.pem”和“private.key“。

真的很蛋疼，官方博客里有篇“SSL 证书安装
”对于用Apache、Nginx和Tomcat的人很有帮助，对于go就莫名其妙了。

2. 在go iris中使用证书
官方案例中的代码  
```
target, _ := url.Parse("https://127.0.1:443")  
go host.NewProxy("127.0.0.1:80",target).ListenAndServe()  
app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
```
需要改成，其中mycert.cert是证书生成完页面的第二个文本框内容，自己复制出来保存成xxx.cert即可。
```
target, _ := url.Parse("https://localhost:443")  
go host.NewProxy("localhost:80", target).ListenAndServe()  
app.Run(iris.TLS(":443", "mycert.cert", "mykey.key"))  
```
