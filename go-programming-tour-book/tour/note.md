1 
- 你遇到的问题，别人可能也有。网上开源的工具集就是这样的。
- 做一个简单的通用工具集，这是最直接的方式，我们用它解决在平时工作中经常会遇到的一些小麻烦

Flag
- 命令行参数的解析
  - CLI 应用程序是什么？
- flag 的 StringVar 方法实现了对命令行参数 name 的解析和绑定
  - 在其他包怎么接收到flag.Args()的参数？
    - 解决：因为我一开始没有写flag.Parse()
  - flag.NewFlagSet
  - flag.StringVar

WordConvert (Cobra)
- Bug
  - 在测试四种单词的转换模式的功能的时候，是用go run ... 没有把所有涉及的文件都写上去。一直出现unknown shorthand flag:
    - 把word包也加上去 就可以了


Time

做一个时间相关的工具，尽可能的优化我们日常获取时间的相关手工行为效率。
- 每一个子命令，都是需要到 rootCmd 中进行注册的，否则将无法使用。

