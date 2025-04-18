# gin-generator

#### 介绍
💡 gin-generator🐔 可以快速把你建立起一个gin+GORM的项目
💡 包括自动生成数据库模型、service+DTO、controller、router的代码

#### 软件架构
1. gin 框架
2. GORM 数据库
3. cobra 命令行工具
4. viper 配置管理工具


#### 安装教程

go install github.com/eddylee1010/gin-generator@latest

#### 使用说明

1. gin-generator -h 获取帮助
2. gin-generator version 查看当前版本号
3. gin-generator gen 所有生成命令的父级命令
   * gin-generator gen project 生成项目结构
   * gin-generator gen model 创建数据库模型文件
   * gin-generator gen service 创建service文件
   * gin-generator gen controller 创建controller文件
   * gin-generator gen router 创建router文件
