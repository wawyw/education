## 环境配置

1. 将`GOPATH`设置为`/root/go`

```
export GOPATH=/root/go 
```

2. 拉取项目

```
cd $GOPATH/src && git clone https://github.com/wawyw/education.git
```

3. 在`/etc/hosts`中添加：

```
127.0.0.1  orderer.example.com
127.0.0.1  peer0.org1.example.com
127.0.0.1  peer1.org1.example.com
```

4. 添加依赖：

```
cd education && go mod tidy
```

## 启动项目

运行启动脚本

```
./start.sh
```

启动区块链浏览器(在另一终端执行)

```
cd fixtures/explorer
./start.sh
```

关闭区块链浏览器

```
./stop.sh
```

项目主页：访问 localhost:9000

区块链浏览器：访问 localhost:8080

## 停止项目

education/下执行

```
./stop.sh
```
