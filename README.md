# douyin
2022字节跳动青训营 抖音项目  

## 接口实现
- [x] 基础接口  
- [x] 扩展接口-I  
- [x] 扩展接口-II  

## 架构设计
TODO

## TODO
- [ ] 用户密码加解密与验证
- [ ] Token生成、存储、验证
- [ ] 传入字符串验证（合法性｜安全性）
- [ ] 新的资源存储与分发方式（视频｜封面）
- [ ] ...

## 使用方法
0. 安装ffmpeg（用于上传视频后的处理）  
MacOS:`brew install ffmepg`  
Windows: 待补充
2. 复制`config/config.example.yaml`为`config/config.yaml`
3. 修改配置文件`config/config.yaml`  
domain配置项用于上传视频后生成的`play_url`与`cover_url` 注意将域名解析到后端所监听的IP  
mysql相关配置只需要建立数据库并分配用户权限 数据表会在首次启动时自动生成
# Credits
https://github.com/gin-gonic/gin  
https://github.com/go-gorm/gorm  
https://github.com/spf13/viper  
