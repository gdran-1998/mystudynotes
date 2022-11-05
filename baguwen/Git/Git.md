# Git

## 创建版本库

版本库又名仓库，英文名repository。可以理解成一个目录，这个目录里面的所有文件都可以被Git管理起来，每个文件的修改、删除，Git都能跟踪，以便任何时刻都可以追踪历史，或者在将来某个时刻可以“还原”。

git init  把目录变成Git可以管理的仓库

把一个文件放到Git仓库只需两步：

git add	告诉Git，把文件添加到仓库；

git commit	告诉Git，把文件提交到仓库。



初始化一个Git仓库，使用`git init`命令。

添加文件到Git仓库，分两步：

1. 使用命令`git add <file>`，注意，可反复多次使用，添加多个文件；
2. 使用命令`git commit -m <message>`，完成。



## 时光穿梭机

git status	时刻掌握仓库当前的状态

git diff	查看difference，显示的格式是Unix通用的diff格式



小结

​	1.要随时掌握工作区的状态，使用`git status`命令。

​	2.如果`git status`告诉你有文件被修改过，用`git diff`可以查看修改内容。



### 版本回退

git log	显示从最近到最远的提交日志

git reset 回退到哪个版本





小结

现在总结一下：

- `HEAD`指向的版本就是当前版本，因此，Git允许我们在版本的历史之间穿梭，使用命令`git reset --hard commit_id`。
- 穿梭前，用`git log`可以查看提交历史，以便确定要回退到哪个版本。
- 要重返未来，用`git reflog`查看命令历史，以便确定要回到未来的哪个版本。



### 工作区和暂存区

工作区（Working Directory）	就是电脑里能看见的目录

版本库（Repository）	工作区有个隐藏目录`.git`,这不算工作区，而实Git的版本库。Git的版本库里存了很多东西，其中最重要的就是称为stage（或者叫index）的暂存区，还有Git为我们自动创建的第一个分支`master`，以及指向`master`的一个指针叫`HEAD`。



### 管理修改

为什么Git比其他版本控制系统设计得优秀，因为Git跟踪并管理的是修改，而非文件。



### 撤销修改

git checkout	把文件在工作区的修改全部撤销

git reset	可以把暂存区的修改撤销掉（unstage），重新放回工作区。`git reset`命令既可以回退版本，也可以把暂存区的修改回退到工作区。当我们用`HEAD`时，表示最新的版本。



小结

又到了小结时间。

场景1：当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时，用命令`git checkout -- file`。

场景2：当你不但改乱了工作区某个文件的内容，还添加到了暂存区时，想丢弃修改，分两步，第一步用命令`git reset HEAD <file>`，就回到了场景1，第二步按场景1操作。

场景3：已经提交了不合适的修改到版本库时，想要撤销本次提交，参考版本回退一节，不过前提是没有推送到远程库。



### 删除文件

一般情况下，你通常直接在文件管理器中把没用的文件删了，或者用`rm`命令删了：`rm test.txt`。

现在你有两个选择：

一是确实要从版本库中删除该文件，那就用命令`git rm`删掉，并且`git commit`：`git rm test.txt`，`git commit -m " remove test.txt"`，

二是删错了，因为版本库里还有呢，所以可以很轻松地把误删的文件恢复到最新版本：`git checkout -- test.txt`。



小结

命令`git rm`用于删除一个文件。如果一个文件已经被提交到版本库，那么你永远不用担心误删，但是要小心，你只能恢复文件到最新版本，你会丢失**最近一次提交后你修改的内容**。





## 远程仓库

Github	提供Git仓库托管服务。即只要注册一个GitHub账号，就可以免费获得Git远程仓库。

你的本地Git仓库和GitHub仓库之间的传输是通过SSH加密的，所以，需要一点设置：

第1步：创建SSH Key。

第2步：登录Github，添加“SSH Key”



### 添加远程库

在Github上创建一个Git仓库

把一个已有的本地仓库与之关联，然后，把本地仓库的内容推送到GitHub仓库:`git remote add origin git@github.com:eryuesenlin/learngit.git`

下一步，就可以把本地库的所有内容推送到远程库上：`git push -u origin master`；失败的话 `git pull --rebase origin master`

从现在起，只要本地作了提交，就可以通过命令：`git push origin master ` 把本地`master`分支的最新修改推送至GitHub

#### 删除远程库

先使用 `git remote -v` 查看远程库信息，然后根据名字 `origin` 删除 `git remote rm origin`

此处的“删除”其实是解除了本地和远程的绑定关系，并不是物理上删除了远程库。远程库本身并没有任何改动。要真正删除远程库，需要登录到GitHub，在后台页面找到删除按钮再删除。

#### 小结

要关联一个远程库，使用命令`git remote add origin git@server-name:path/repo-name.git`；

关联一个远程库时必须给远程库指定一个名字，`origin`是默认习惯命名；

关联后，使用命令`git push -u origin master`第一次推送master分支的所有内容；

此后，每次本地提交后，只要有必要，就可以使用命令`git push origin master`推送最新修改；

分布式版本系统的最大好处之一是在本地工作完全不需要考虑远程库的存在，也就是有没有联网都可以正常工作，而SVN在没有联网的时候是拒绝干活的！当有网络的时候，再把本地提交推送一下就完成了同步，真是太方便了！



### 从远程克隆

git clone

#### 小结

要克隆一个仓库，首先必须知道仓库的地址，然后使用`git clone`命令克隆。

Git支持多种协议，包括`https`，但`ssh`协议速度最快。



## 分支管理

假设你准备开发一个新功能，但是需要两周才能完成，第一周你写了50%的代码，如果立刻提交，由于代码还没写完，不完整的代码库会导致别人不能干活了。如果等代码全部写完再一次提交，又存在丢失每天进度的巨大风险。

现在有了分支，就不用怕了。你创建了一个属于你自己的分支，别人看不到，还继续在原来的分支上正常工作，而你在自己的分支上干活，想提交就提交，直到开发完毕后，再一次性合并到原来的分支上，这样，既安全，又不影响别人工作。



### 创建与合并分支

#### 创建

一开始， `master`	分支是一条线，Git用 `master` 指向最新的提交，再用 `HEAD` 指向 `master` ，就能确定当前分支，以及当前分支的提交点。每次提交，`master ` 分支都会向前移动一步，这样，随着你不断提交，`master` 分支的线也越来越长。

当我们创建新的分支，例如`dev`时，Git新建了一个指针叫`dev`，指向`master`相同的提交，再把`HEAD`指向`dev`，就表示当前分支在`dev`上。

从现在开始，对工作区的修改和提交就是针对`dev`分支了，比如新提交一次后，`dev`指针往前移动一步，而`master`指针不变。

#### 合并

最简单的方法，就是直接把`master`指向`dev`的当前提交，就完成了合并。

合并完分支后，甚至可以删除`dev`分支。删除`dev`分支就是把`dev`指针给删掉，删掉后，我们就剩下了一条`master`分支。



`git checkout -b dev`  创建dev分支，然后切换到dev分支，参数`-b`表示创建并切换

相当于两条命令：`git branch dev`加上`git checkout dev`

`git branch` 查看当前分支

`git checkout master ` 或`git switch master`切回`master`分支

`git merge dev` 把`dev`分支的工作成果合并到`master`分支上

`git branch -d dev`合并完成后,删除`dev`分支



#### 小结

Git鼓励大量使用分支：

查看分支：`git branch`

创建分支：`git branch <name>`

切换分支：`git checkout <name>`或者`git switch <name>`

创建+切换分支：`git checkout -b <name>`或者`git switch -c <name>`

合并某分支到当前分支：`git merge <name>`

删除分支：`git branch -d <name>`



### 总结

1、（先进入项目文件夹）通过命令 git init 把这个目录变成git可以管理的仓库

git init
2、把文件添加到版本库中，使用命令 git add .添加到暂存区里面去，不要忘记后面的小数点“.”，意为添加文件夹下的所有文件

git add .
3、用命令 git commit告诉Git，把文件提交到仓库。引号内为提交说明

git commit -m 'first commit'
4. 添加一个分支

git branch -m master
5、关联到远程库

git remote add origin 你的远程库地址
如：

git remote add origin https://github.com/githubusername/demo.git
6、获取远程库与本地同步合并（如果远程库不为空必须做这一步，否则后面的提交会失败）

git pull --rebase origin master
7、把本地库的内容推送到远程，使用 git push命令，实际上是把当前分支master推送到远程。执行此命令后会要求输入用户名、密码，验证通过后即开始上传。

git push -u origin master




