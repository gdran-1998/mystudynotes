# 模块一：Go 程序是怎么跑起来的

## 工程师的学习与进步

- 多写代码，累计代码量（至少累计几十万的代码量，才能对设计有自己的观点），要总结和思考，如何对过去的工作进行改进（如自动化、系统化），积累自己的代码库、笔记库、开源项目。

- 读好书，建立知识体系（比如像Designing Data-Intensive Application 这种书，应该多读几遍）。
- 关注一些靠谱的国内外新闻源，通过问题出发，主动使用Google，主动去reddit、hackernews上参与讨论，避免被困在信息茧房里。
- 锻炼口才和演讲能力，内部分享 -> 外部分享。在公司内，该演要演，不要只是闷头干活。
- 通过输出促进输入（博客、公众号、分享），打造个人品牌，通过读者的反馈循环提升自己的认知。
- 信息源：Github Trending、reddit、medium、hacker news、acm.org、oreily，国外的领域相关大会（OSDI，SOSP，VLDB）、论文，国际一流公司的技术博客，YouTube上的国外工程师演讲。

为什么Go语言适合现代的后端编程环境？

- 服务类应用以API居多，IO密集型，且网络IO最多。
- 运行成本低，无VM。网络连接数不多的情况下内存占用低。
- 强类型语言，易上手，易维护。

为什么适合基础设施？

- kubernetes、etcd、istio、docker已经证明了Go的能力。



对Go的**启动**和**执行流程**建立简单的宏观认识

## 理解可执行文件

文本 -> 编译 -> 二进制可执行文件

编译：文本代码 -> 目标文件（.o , .a）

链接：将目标文件合并为可执行文件

可执行文件在不同的操作系统上规范不一样：Linux（ELF）、Windows（PE）、MacOS（Mach-O）

Linux的可执行文件ELF（Executable and Linkable Format）为例，ELF由几部分构成：

- ELF header
- Section header
- Sections

操作系统可执行文件的步骤（以linux为例）：

解析ELF header -> 加载文件至内存 -> 从entry point 开始执行代码

通过entry point找到Go进程的执行入口。



## Go进程的启动与初始化

CPU无法理解文本，只能执行一条一条的二进制机器码指令，每次执行完一条指令，pc寄存器就指向下一条继续执行。在64位平台上pc寄存器=rip

Go语言是一门有 runtime 的语言，可以认为 runtime 是为了实现额外的功能，而在程序启动时自动加载/运行的一些模块。

Go语言的 runtime 包括：

- Scheduler：调度器管理所有的G，M，P，在后台执行调度循环。
- Netpoll：网络轮询负责管理网络FD相关的读写、就绪事件。
- Memory Management：当代码需要内存时，负责内存分配工作。
- Garbage Collector：当内存不再需要时，负责回收内存。

这些模块中，最核心的就是Scheduler，它负责串联所有的 runtime 流程。



通过 entry point 找到Go进程的执行入口：

runtime._rt0_amd64_linux -> runtime._rt0_amd64 -> rumtime.rt0_go

 rumtime.rto_go: **开始执行用户的main函数 -> 初始化内置数据结构 -> 获取CPU核心数 -> 全局m0 g0 初始化 -> argc argv处理。**（从开始执行用户的main函数 开始进入调度循环）

m0：Go程序启动后创建的第一个线程。


