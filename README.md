<div align="center">
<strong>
<samp>

[English](https://github.com/Jeremy-Run/myrrh/blob/main/README.md) | [中文版](https://github.com/Jeremy-Run/myrrh/blob/main/README_CN.md)

</samp>
</strong>
</div>

## What is myrrh?
myrrh is a lightweight rules engine

## Doc(chinese)
[Document](https://zhuanlan.zhihu.com/p/668588745)

## How to use it?

 First, you must ensure that you have a golang running environment, and then download the code locally

```bash
git clone https://github.com/Jeremy-Run/myrrh.git
```

Then execute the following code in the home directory of this project：

```bash
go run main.go
```

Then you can see the execution results：

```bash
>>> Start......
>>> Simple Expression: (user.LoginTimes() >= 5) 
>>> [user.LoginTimes() >= 5] execute result is: true 
>>> Congrats! You have completed the activity and received ten percent off coupon reward 
>>> The requirement execute result is: true 
>>> End......
```

If you want to debug it, you can go to the [config](config/config.go) file and modify the `CaseActivity`.

😝Please click on star~

## Other Case

```bash
go run main.go --num=2
```

```bash
>>> Start......
>>> Simple Expression: ((user.FollowTimes() >= 9) || (user.InvitationTimes() >= 2)) 
>>> [user.FollowTimes() >= 9] execute result is: false 
>>> [user.InvitationTimes() >= 2] execute result is: true 
>>> Congrats! You have completed the activity and received ten percent off coupon reward 
>>> The requirement execute result is: true 
>>> End......
```

```bash
go run main.go --num=3
```

```bash
>>> Start......
>>> Simple Expression: (((user.LoginTimes() >= 5) && (user.OrderTimes() >= 1) && (user.CommentTimes() >= 1)) || ((user.BrowseTimes() >= 15) && (user.PostTimes() >= 3) && ((user.FollowTimes() >= 8) || (user.InvitationTimes() >= 2)))) 
>>> [user.FollowTimes() >= 8] execute result is: true 
>>> [user.InvitationTimes() >= 2] execute result is: true 
>>> [user.BrowseTimes() >= 15] execute result is: true 
>>> [user.PostTimes() >= 3] execute result is: true 
>>> [user.LoginTimes() >= 5] execute result is: true 
>>> [user.OrderTimes() >= 1] execute result is: true 
>>> [user.CommentTimes() >= 1] execute result is: true 
>>> Congrats! You have completed the activity and received ten percent off coupon reward 
>>> The requirement execute result is: true 
>>> End......
```
