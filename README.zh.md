# algorithms

[![Status](https://img.shields.io/badge/status-in%20development-orange)](https://github.com/cuihairu/algorithms/releases)

> [English](README.md) | **中文**

用 Go 实现的数据结构与算法，拆分为独立 Go module，各自带单元测试，另附一个 VuePress 文档站。

## 模块

| 模块 | 包路径 | 内容 | 测试数 |
| --- | --- | --- | --- |
| array | `array/go` | 泛型数组封装 + 越界检查 | 4 |
| sorts | `sorts/go` | 冒泡/插入/选择/归并/快排/计数/桶排序 | 7 |
| stack | `stack/go` | `Stack` 接口 + 数组栈 / 链表栈 | 6 |
| linkedlist | `linkedlist/go` | 头哨兵单链表 + 读写锁 + 快慢指针 | 11 |
| queue | `queue/go` | `Queue` 接口 + 数组队列（头游标）/ 链表队列 | 6 |
| tree | `tree/go` | 泛型 BST（插入/查/删/最值/中序/高度/平衡） | 8 |
| heap | `heap/go` | 泛型最小堆（上浮/下沉/O(n) 建堆） | 6 |
| hash | `hash/go`（模块名 `hashmap`） | 链地址哈希表 + FNV-1a + 负载因子倍增扩容 | 5 |
| graph | `graph/go` | 邻接表有向/无向图 + BFS/DFS/最短路 | 9 |

## 目录结构

```
algorithms/
├── array/go/
├── sorts/go/
├── stack/go/
├── linkedlist/go/
├── queue/go/
├── tree/go/
├── heap/go/
├── hash/go/
├── graph/go/
├── docs/          # VuePress 站点（每个模块一篇笔记）
├── go.work        # 九个模块统一工作区
└── package.json   # 文档脚本
```

## 运行测试

九个模块均已纳入 `go.work`，在各模块目录内执行：

```bash
cd <module>/go && go test ./...
```

带竞态检测跑单链表测试：

```bash
cd linkedlist/go && go test -race ./...
```

## 文档站

`docs/` 是一个 VuePress 站点（theme-hope），每个模块一篇笔记。

```bash
npm install
npm run docs:dev     # 本地开发服务器
npm run docs:build   # 构建到 docs/.vuepress/dist
```

## Go 版本

`go.work` 目标 Go 1.19。
