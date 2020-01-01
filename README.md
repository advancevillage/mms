#### 项目名称
 * mms

#### 项目描述
    商品中心
    @link: http://www.woshipm.com/pd/508411.html

#### 代码规范
    1: 命名规范
       - 变量命名
         * 小驼峰  eg: Id service GoodsId
       - 常量命名
         * 前缀知意 eg: ErrorQueueXXX = errors.New("xxx")
       - 函数命名
         * 动词+名词 eg: QueryMerchandiseById() error
         * 至少返回一个error类型
    2: SQL规范
         * fmt.Sprintf()
         eg: sql := fmt.Sprintf("" +
                    "select * " +
                    "from %s m " +
                    "where m.a = '%s' " +
                    ";",
                    TableName, param)
    3: DSL规范

#### 目录结构
    mms
       |
       |--- docs      API文档 Swagger
       |--- config    配置文件
       |--- deploy    部署脚本
       |--- template  通知消息模版
       |--- src       源代码
             ｜
             ｜--- init  初始化
             ｜--- main  启动入口
             ｜--- route 路由
             ｜--- ...
       |--- README.md 说明文件

     每个目录组成
       - model.go    定义数据类型
       - subject.go  文件主题
     eg:  init
            |--- model.go
            |--- single.go      //初始化信号
            |--- log.go         //初始化日志
            |--- config.go      //初始化配置
            |--- args.go        //初始化命令行参数

#### 构建编译
    go build -o bin/mms -gcflags "-N -l" -ldflags "-X main.commit=4d399017 -X main.version=v5.0.0 -X main.mode=http"  src/main/mms.go

#### 参考文件
    1: swag Download: https://github.com/swaggo/swag/releases  [1.6.2]
       $ cp swag  /usr/sbin/
              
