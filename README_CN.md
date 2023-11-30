<div align="center">
<strong>
<samp>

[English](https://github.com/Jeremy-Run/myrrh/blob/main/README.md) | [中文版](https://github.com/Jeremy-Run/myrrh/blob/main/README_CN.md)

</samp>
</strong>
</div>

## myrrh是什么？
myrrh是一个轻量级规则引擎


## 详细的解说
[文档](https://zhuanlan.zhihu.com/p/668588745)

## 如何使用它?

首先，你必须确保您有一个golang的运行环境，然后把代码下载到本地

```bash
git clone https://github.com/Jeremy-Run/myrrh.git
```

然后在本项目的主目录下执行以下代码:

```bash
go run main.go
```

这样你就可以看到执行结果:

```bash
>>> Start executing requirement......
>>> [user.FollowTimes() >= 8] execute result is: true 
>>> [user.InvitationTimes() >= 2] execute result is: true 
>>> [user.BrowseTimes() >= 15] execute result is: true 
>>> [user.PostTimes() >= 3] execute result is: true 
>>> [user.LoginTimes() >= 5] execute result is: true 
>>> [user.OrderTimes() >= 1] execute result is: true 
>>> [user.CommentTimes() >= 1] execute result is: true 
>>> Congrats! You have completed the activity and received 10% Off Coupon reward 
>>> The requirement execute result is: true 
>>> End of requirement execution......
```

如果你想调试它，你可以跳到[config](config/config.go)文件并修改`SimulationActivity`。

😝点个star支持一下~

代码还会持续更新哦......

