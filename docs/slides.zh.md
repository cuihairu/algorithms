---
title: 幻灯片页
icon: slides
layout: Slide
---

> [English](slides.md) | **中文**

<!-- markdownlint-disable MD024 MD033 MD051 -->

@slidestart

<!-- .slide: data-transition="slide" -->

## 幻灯片演示

<!-- .element: class="r-fit-text" -->

一个简单的幻灯片演示与实用提示。

<!-- .element: class="r-fit-text" -->

> By Mr.Hope. 请向下滚动鼠标滚轮查看下一页幻灯片

---

<!-- .slide: data-transition="slide" data-auto-animate -->

## 标记幻灯片

<!-- .element: class="r-fit-text" -->

[👇](#/1/1)

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 标记幻灯片

<!-- .element: class="r-fit-text" -->

使用 `---` 标记水平幻灯片

<!-- .element: class="fragment fade-in" -->

使用 `--` 在水平幻灯片内分隔垂直幻灯片。

<!-- .element: class="fragment fade-in" -->

使用 `<!-- .slide: ... -->` 为幻灯片添加属性

<!-- .element: class="fragment fade-in" -->

使用 `<!-- .element: ... -->` 为前一个 html 元素添加属性

<!-- .element: class="fragment fade-in" -->

---

<!-- .slide: data-transition="slide" data-auto-animate -->

## Markdown

<!-- .element: class="r-fit-text" -->

你可以在幻灯片中使用各种标记语法。

<!-- .element: class="r-fit-text" -->

--

<!-- .slide: data-auto-animate -->

## Markdown

你可以在幻灯片中使用各种标记语法。

### 这是一个 H3

标题默认会转换为大写。

这里是包含 **加粗**、_斜体_、~~删除线~~ 文本以及 [链接](https://mrhope.site) 的段落，它可以自动换行，因此无需担心长度。

--

<!-- .slide: data-auto-animate -->

## Markdown

你可以在幻灯片中使用各种标记语法。

列表默认为 `inline-block`。

- 条目
- 条目
- 条目

1. 条目 1
1. 条目 2
1. 条目 3

--

<!-- .slide: data-auto-animate -->

## Markdown

你可以在幻灯片中使用各种标记语法。

启用 `highlight` 插件后，代码块会自动高亮。

```js
const a = 1;
```

--

<!-- .slide: data-auto-animate -->

## Markdown

你可以在幻灯片中使用各种标记语法。

启用 `math` 插件后，还可以使用 tex 语法编写数学公式。

$$
J(\theta_0,\theta_1) = \sum_{i=0}
$$

--

<!-- .slide: data-auto-animate -->

## Markdown

你可以在幻灯片中使用各种标记语法。

⚠**注意**：不支持表格、分隔线以及其他非标准 Markdown 语法。

---

<!-- .slide: data-transition="slide" data-auto-animate -->

## 布局

<!-- .element: class="r-fit-text" -->

--

<!-- .slide: data-auto-animate  -->

## 布局

<!-- .element: class="r-fit-text" -->

👆 `r-fit-text` 类让文字在不溢出幻灯片的前提下尽可能大。

--

<!-- .slide: data-auto-animate  -->

## 布局

![Logo](/logo.svg)

<!-- .element: class="r-stretch" -->

👆 `r-stretch` 辅助类可以调整元素（如图片或视频）的尺寸，使其填满幻灯片的剩余垂直空间。

--

<!-- .slide: data-auto-animate data-background-color="rgb(70, 70, 255)" -->

## 布局

### 背景

通过为幻灯片添加 `data-background` 属性自定义背景。

---

<!-- .slide: data-auto-animate -->

## Fragment

<!-- .element: class="r-fit-text" -->

--

<!-- .slide: data-auto-animate -->

## Fragment

<!-- .element: class="r-fit-text" -->

Fragment 用于高亮或逐个显示幻灯片上的单个元素。

为元素添加 `fragment` 和动画类。

--

<!-- .slide: data-auto-animate -->

## Fragment

### 动画类

- `fade-in`
<!-- .element: class="fragment fade-in" -->

- `fade-out`
<!-- .element: class="fragment fade-out" -->

- `fade-up`
<!-- .element: class="fragment fade-up" -->

<!-- list break -->

- `fade-down`
<!-- .element: class="fragment fade-down" -->

- `fade-left`
<!-- .element: class="fragment fade-left" -->

- `fade-right`
<!-- .element: class="fragment fade-right" -->

<!-- list break -->

- `fade-in-then-out`
<!-- .element: class="fragment fade-in-then-out" -->

- `fade-in-then-semi-out`
<!-- .element: class="fragment fade-in-then-semi-out" -->

--

<!-- .slide: data-auto-animate -->

## Fragment

### 动画类

- `grow`
<!-- .element: class="fragment grow" -->

- `shrink`
<!-- .element: class="fragment shrink" -->

- `strike`
<!-- .element: class="fragment strike" -->

<!-- list break -->

- `highlight-red`
<!-- .element: class="fragment highlight-red" -->

- `highlight-green`
<!-- .element: class="fragment highlight-green" -->

- `highlight-blue`
<!-- .element: class="fragment highlight-blue" -->

<!-- list break -->

- `highlight-current-red`
<!-- .element: class="fragment highlight-current-red" -->

- `highlight-current-green`
<!-- .element: class="fragment highlight-current-green" -->

- `highlight-current-blue`
<!-- .element: class="fragment highlight-current-blue" -->

--

<!-- .slide: data-auto-animate -->

## Fragment

### 多个 Fragment

通过嵌套包裹，可以对同一元素按顺序应用多个 fragment

<span class="fragment fade-in">
  <span class="fragment highlight-red">
    <span class="fragment fade-out">
      淡入 > 变红 > 淡出
    </span>
  </span>
</span>

--

<!-- .slide: data-auto-animate -->

## Fragment

### 顺序

可以使用 `data-fragment-index` 属性改变顺序。

多个元素可以在相同索引处同时出现。

- 最后出现
<!-- .element: class="fragment" data-fragment-index="3"-->

- 第二个出现
<!-- .element: class="fragment" data-fragment-index="2"-->

<!-- list break -->

- 最先出现
<!-- .element: class="fragment" data-fragment-index="1"-->

- 第二个出现
<!-- .element: class="fragment" data-fragment-index="2"-->

---

<!-- .slide: data-transition="slide" data-auto-animate -->

## 过渡

<!-- .element: class="r-fit-text" -->

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 过渡

<!-- .element: class="r-fit-text" -->

过渡效果可以通过全局设置 `transition` 配置项，或在特定幻灯片上设置 `data-transition` 属性来改变。

可选值：

- none
- fade
- slide

<!-- list break -->

- convex
- concave
- zoom

--

<!-- .slide: data-auto-animate -->

## 过渡

<!-- .element: class="r-fit-text" -->

### 自动动画

在相邻幻灯片上添加 `data-auto-animate`，可以对未变化的元素产生动画效果。

---

<!-- .slide: data-transition="slide" data-auto-animate -->

## 功能

<!-- .element: class="r-fit-text" -->

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 功能

<!-- .element: class="r-fit-text" -->

### 代码

启用 `highlight` 插件后，可以高亮代码块。

可以使用 `[a-b|c-d]` 语法按步骤高亮行。

```js [1-2|3|4]
let a = 1;
let b = 2;
let c = (x) => 1 + 2 + x;
c(3);
```

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 功能

<!-- .element: class="r-fit-text" -->

### 总览

演示进行中，按 `Esc` 或 `O` 进入总览模式

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 功能

<!-- .element: class="r-fit-text" -->

### 全屏

演示进行中，按 `F` 或 `F11` 进入全屏

--

<!-- .slide: data-transition="slide" data-auto-animate -->

## 功能

<!-- .element: class="r-fit-text" -->

### 缩放

按住 `alt` 键（Linux 下为 `ctrl`）并点击任意元素即可向该处缩放。

再次点击即可缩放回去。

---

<!-- .element: class="r-fit-text" -->

## 完

@slideend
