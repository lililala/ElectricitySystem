# ElectricitySystem

一个独立出来的用于获取电费的小程序的单用户版本

# 下载

~~在 [GitHub Release](https://github.com/wsndshx/ElectricitySystem/releases/latest) 页面根据你系统的架构获取最新发布的二进制文件。~~

```shell
#解压获取到的主程序
tar -zxvf ElectricitySystem_VERSION_OS_ARCH.tar.gz

# 赋予执行权限
chmod +x ./ElectricitySystem

# 启动 ElectricitySystem
./ElectricitySystem -room [房间号] -time [检查时间间隔(s)]
```

**由于项目有用到CGO，无法顺利设置交叉编译并发布，咕。**

# 编译

## Debian

```shell
sudo apt update
sudo apt install gcc -y
git clone git@github.com:wsndshx/ElectricitySystem.git
cd ElectricitySystem
go build
```

# 未来的计划

这里是我未来打算增加的功能的（不完整）列表，其不具备任何保证。可能会单纯因为“一个飞蛾频繁在我写的时候骚扰我”等原因而取消相关计划。

![image.png](https://i.loli.net/2020/09/29/iWhrdvIaLqC7s9J.png)

- [ ] 计划中
- [x] 已完成
- [ ] ~~已废弃~~

## 主线

- [ ] 任务队列
- [ ] RPC调用

## 支线

- [ ] Web控制面板
- [ ] 基于上一次抄表时间决定下一次抄表时间
- [ ] 多维度用电统计与用电预测