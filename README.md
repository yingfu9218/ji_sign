### 几鸡自动签到

#### 使用编译好的二进制文件方式

```
1. 在 https://github.com/yingfu9218/ji_sign/releases/tag/1.0  页面下载对应平台的应用，解压，进入目录
2. 修改 config.yaml中的账号密码
3. 执行 ./sign 即可，或以守护进程方式执行 nohup ./sign &


```

#### 使用源码编译方式

```
1. git clone https://github.com/yingfu9218/ji_sign.git
2. cd ji_sign
3. 编译 go build -o sign main.go 得到sign 执行文件
4. 修改 config.yaml中的账号密码
5. 执行 ./sign 即可，或以守护进程方式执行 nohup ./sign &


```