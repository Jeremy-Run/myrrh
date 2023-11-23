# myrrh
myrrh is a lightweight rules engine

## How to use it?

 First, you must ensure that you have a golang operating environment, and then download the code locally

```bash
git clone https://github.com/Jeremy-Run/myrrh.git
```

Then execute the following code in the home directory of this project：

```bash
go run main.go
```

Then you can see the execution results：

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

If you want to debug it, you can go to the [config](config/config.go) file and modify the `SimulationActivity`.

