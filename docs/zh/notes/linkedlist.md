---
title: 单链表
icon: link
---

# 单链表

> [English](/notes/linkedlist) | **中文**

`linkedlist/go` 模块实现带头哨兵节点的单链表，读写锁保护。

## 实现位置

- 源码：`linkedlist/go/linkedlist.go`
- 测试：`linkedlist/go/linkedlist_test.go`（11 个测试函数）

## 接口

构造函数：`NewLinkedList()`（带头哨兵节点的链表）。

| 方法 | 说明 |
| --- | --- |
| `AddFirst / AddLast` | 头插 / 尾插 |
| `GetFirst / GetLast` | 读首 / 读尾 |
| `RemoveFirst / RemoveLast` | 删首 / 删尾，空表返回 nil |
| `Insert(index, element)` | 1 基下标插入；`index<=0` 不动作；越界等价尾插 |
| `DeleteByIndex(index)` | 按下标删除，越界不动作 |
| `DeleteByBottomIndex(index)` | 删倒数第 N 个（快慢指针） |
| `Contains(element)` | 是否包含 |
| `Reverse()` | 原地反转 |
| `HasCycle()` | 快慢指针判环 |
| `FindMiddleNode()` | 快慢指针取中点；偶数长度返回前一个中点 |
| `Size / IsEmpty / String` | 常规 |

## 历史坑位（已修复，回归用例在测试里）

- **String**：遍历忘了 `cur = cur.next`，死循环。
- **Contains**：循环条件 `next.next != nil` 漏判尾节点。
- **HasCycle**：条件缺 `fast != nil`，单节点链表 `nil.next` panic。
- **FindMiddleNode**：`fast.next.next` 未判空，偶数长度 panic。
- **Insert**：持 `RLock` 写链表（数据竞争），已改 `Lock`，`-race` 验证通过。

## 测试状态

```bash
cd linkedlist/go && go test -race ./...
```

当前 11 个测试全部通过（含手工构造的带环链表与偶数长度取中点用例）。
