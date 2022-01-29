<p align="center">
<a href="https://github.com/SoundCrusher"><img src="https://img.shields.io/badge/author-OgromWang-brightgreen"/></a>
<a href="https://theme.typora.io/theme/Drake"><img src="https://img.shields.io/badge/theme-Github-orange.svg"/></a>
<a href="https://github.com/SoundCrusher"><img src="https://img.shields.io/badge/license-OgromWang-blue"/></a>
</p>

# About
这是一个基于 fyne 开发的项目，为解决平常开发中遇到的繁琐问题

# Feature

## 1. `curl` 格式化
在开发中，QA 人员提 bug 时，会将 `curl` 命令贴到 bug 单中，但总是多次空行等问题，每次都需要手动删去很是麻烦。

然后 google 了一圈没有发现能格式化 `curl` 的工具，以此有了该项目。

## 2. 八进制字符串解码
支持解析八进制字符串，如：

`"\345\207\272\345\272\223\351\224\231\350\257\257"` = `"出库错误"`

## 3. 待开发

# Project

```text
.
├── LICENSE
├── README.md
├── app
│   ├── content
│   ├── data
│   ├── feature
│   │   ├── convert
│   │   ├── curl
│   │   ├── interface.go
│   │   └── more.go
│   ├── menu
│   │   └── menu.go
│   └── vars
├── cmd
│   ├── Makefile.mk
│   └── main.go
├── dist
├── example
│   └── main_demo.go
├── go.mod
├── go.sum
└── static
```

# Run
```shell
$ go run cmd/main.go
```

# Fyne Build
```shell
$ fyne package -os (darwin | windows | other...) -name (应用名称) -icon (应用图标)
```