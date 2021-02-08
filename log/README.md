# package log
gkit 的基础日志

## TODO
支持 logrus 的 WithFields 方法


## 功能
可以传入配置文件，增强原生的日志组件。

也可以利用hook写入文件，但这里遇到的问题是：写入到日志的内容不是格式化过的，可能以后需要自定义

## 用法
```
_, _ = SetConfigFile("mate.conf", "default") // 设置配置文件, default 表示默认的日志实例

Info("default log use hijack Info") // 使用默认日志实例调用
Logger().Info("default log from Glogger") // 使用默认日志实例调用
Logger("mike").Info("mike log from Glogger") // 手动选择实例调用
SetReportCaller(true) // 添加方法、文件以及行数打印;logrus_meta 里没有找到如何配置，手动提供方法

```

## 配置文件
其中mike、default 表示不同的日志实例

level : trace\debug\info\warn\error\fatal\panic 如果设置成info，那么debug级别的是看不到的（使用场景：生产、测试环境区分）

formatter.name : json/text

timestamp-format : 在json格式下设置失败，不知道为何

hooks 钩子,详情看 https://github.com/gogap/logrus_mate 以及 https://github.com/gogap/logrus_mate/tree/master/example 的配置文件
```
mike {
    formatter.name = "text"
}

default {

        level = "trace"

        formatter.name = "text"
        formatter.options  {
                            timestamp-format  = "2006-01-02  15:04:05"
        }

        hooks {
                file {
                    filename = "1.log"
                    daily = true
                    rotate = true
                    level = 4
                }


        }
}
```