douyin

# 启动：
启动server/cmd/api下的main  
server/cmd/interact下执行 sh output/bootstrap.sh  
server/cmd/social下执行 sh output/bootstrap.sh  
server/cmd/user下执行 sh output/bootstrap.sh  
server/cmd/video下执行 sh output/bootstrap.sh  
# 调用关系
![图片alt](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/341beef01d0146a48962bb826f49dbc1~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?"调用关系")
# 服务关系
![图片alt](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4c7037984fa74c71a14e2d8311cb68e8~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)
# 界面展示
![图片alt](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/838953fbd2f9455aba1ea8e8ca4bae16~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)
![图片alt](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/9cc0c0e07cd543d692b503e3aa7bef7a~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)
![图片alt](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/aaa0927a722c46cb8ecdcf51928d1b99~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)
![图片alt](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f0a923fd9d3d4d83aa8876c356753718~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)
    ![图片alt](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/506552d65a3e4901bcc4af4c28ffb38e~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)  
# 服务介绍
|种类|介绍|  
|---|:---|  
|api|基于Hertz的网关服务|  
|user|用户服务，包括注册，登录以及用户信息|
|video|视频服务，包括视频流推送，视频发布列表（视频发布放在api中）|
|interact|交互服务，包括点赞，点赞列表，评论和评论列表|
|social|社交服务，包括关注，关注与被关注列表，朋友列表以及消息发送和消息列表|  
# 数据库表设计
![图片alt](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c479d82b643047c29d2108ea05c561aa~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp?)