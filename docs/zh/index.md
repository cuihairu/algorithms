---
layout: home

hero:
  name: 算法笔记
  text: Go 数据结构与算法
  tagline: 九个带单元测试的 Go 小模块，覆盖数组、排序、栈、链表、队列、二叉搜索树、堆、哈希表与图。
  image:
    src: /logo.svg
    alt: 算法笔记
  actions:
    - theme: brand
      text: 浏览笔记
      link: /zh/notes/
    - theme: alt
      text: English
      link: /

features:
  - title: 实现均带测试
    details: 每个数据结构与排序实现都有单元测试；用 `cd <module>/go && go test ./...` 运行单个模块。
  - title: 合适处使用泛型
    details: tree 与 heap 使用泛型；array 封装同样泛型，并显式做越界检查。
  - title: 结构一致
    details: 每篇笔记记录源码位置、公开接口、使用示例与测试状态。
  - title: 双语
    details: 英文位于站点根路径，简体中文位于 `/zh/`。
---

## 笔记

| 主题 | 模块 | 实现 |
| --- | --- | --- |
| [数组封装](/zh/notes/array) | `array/go` | 泛型数组 + 越界检查 |
| [排序算法](/zh/notes/sorts) | `sorts/go` | 冒泡/插入/选择/归并/快排/计数/桶 |
| [栈](/zh/notes/stack) | `stack/go` | 数组栈 / 链表栈 |
| [单链表](/zh/notes/linkedlist) | `linkedlist/go` | 头哨兵 + 读写锁 + 快慢指针 |
| [队列](/zh/notes/queue) | `queue/go` | 数组队列（头游标）/ 链表队列 |
| [二叉搜索树](/zh/notes/tree) | `tree/go` | 泛型 BST：插入/查/三态删除/最值/中序/高度/平衡 |
| [二叉堆](/zh/notes/heap) | `heap/go` | 泛型最小堆：上浮/下沉/O(n) 建堆 |
| [哈希表](/zh/notes/hash) | `hash/go` | 链地址 + FNV-1a + 负载因子倍增扩容 |
| [图](/zh/notes/graph) | `graph/go` | 邻接表有向/无向图 + BFS/DFS/最短路 |

站点约定见[指南](/zh/guide/)与[组件演示](/zh/demo/)。
