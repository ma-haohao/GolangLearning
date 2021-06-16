# golang 依赖管理——go module

## Project

在正式进入module介绍前。通俗的讲，一个project可以有多个module组成，module可以作为独立 的project被别的project作为依赖引用，如下golang工程 `demo` 中就包含了 `foo` 和 `bar` 两个module

```
demo
├── bar
│   └── go.mod
└── foo
    └── go.mod
```

## module介绍

go module(以下称：module、模块、工程模块) 是golang中已发布版本的package的集合，是Go管理依赖的一种方式。

在go.mod中，其包含了main module的module路径、module依赖及其关联信息（版本等），如果一个工程模块需要以 `go module mode`（module模式）开发，在工程模块的根目录下必须包含 `go.mod` 文件。

