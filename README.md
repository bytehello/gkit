# gkit
基于golang的基础库，主要留存一些golang项目的基础功能
## 宗旨
kit 库必须具有的特点：
- 统一
- 标准库方式布局
- 高度抽象
- 支持插件
- 尽量减少依赖
- 持续维护

理论参考 https://lailin.xyz/post/go-training-week3-goroutine.html

项目组织参考:https://github.com/go-kit/kit

## 组件介绍
不传入任何参数，默认使用被logrus替代了的log，正常打印，带时间戳

可以传入配置文件，使用配置文件中的打印，如果没有指定，那么就选择第一个文件来打印（看看logrus_meta是如何选择的）

可以准入配置文件，指定具体的配置打印

可以动态新增 【携带配置文件的】 logger对象

对应 another_log_test.go 中的 TestLogMate TestHijactLoggerByMate 这两个方法好好想一想如何组织代码