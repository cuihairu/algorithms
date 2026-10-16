# algorithms 巡检台账（patrol-20261013-C1 建账）

盘点基准：4 个 Go module（array / sorts / stack / linkedlist）+ vuepress docs 站。
统一验收口径：各 module 内 `go test ./...` 全绿；docs 项以 `npm run docs:build` 成功 + 内容非模板为准。

## 一、已实现（代码 + 测试通过）

- [x] array/go 数组封装，含 4 个测试函数，`go test ./...` ok —— 证据 ede7480
- [x] sorts/go 冒泡/插入/选择/归并/快排/计数 6 种排序，各有测试且通过 —— 证据 934cb3b、71a2c03、1df32b9
- [x] linkedlist/go 双向链表 18 个方法实现完整 —— 证据 655b3bc（注意：测试为空壳，见二档）

## 二、部分实现（代码在，但有 bug 或测试缺失）

- [x] sorts/go 红测试修复（C1 完成）：bucket.go 桶序号误用 nBuckets 整除致全负输入越界（历史 14d87d4 曾修但未修对），且末尾回拷只扫 n 个桶（nBuckets 可大于 n）；counting.go 未按 min 偏移索引致负数越界（此前被 bucket panic 掩盖）。补 TestBucketSortSparse 稀疏两元素回归用例。
  验收达成：sorts 模块 `go test ./...` 7 个测试全绿。证据：本条修复 commit
- [x] go.work 收录全部 4 个 module（C2 完成）。顺带最小修复 linkedlist 空壳测试的编译错误（declared and not used），使四模块测试命令均可执行。
  验收达成：array/sorts/stack/linkedlist 四模块 `go test ./...` 均可执行。证据 3c39425
- [x] stack/go ArrayStack 补齐测试（C2 完成）。测试先红后修：Pop 原先不弹栈（等同 Top）、Flush 后复用 Size()=len(data) 失真、String 漏栈顶元素，三处全修，加编译期接口守卫 `var _ Stack`。
  验收达成：stack 模块 ArrayStack 3 测试全绿。证据 7e2f349
- [x] stack/go LinkedListStack 补 Push/Pop/Top/Size/Flush 五方法（C2 完成）。测试先行（编译红确认缺方法），头插法实现，空栈/ LIFO /Flush 复用三组用例 + 接口守卫。
  验收达成：stack 模块 6 测试全绿。证据 040b66c
- [x] linkedlist/go 测试全覆盖（C3 完成）。测试先行暴露 5 个真 bug 并全部修复：String 死循环（漏 cur=cur.next）、Contains 漏判尾节点、HasCycle 单节点 nil panic、FindMiddleNode 偶数长度 panic、Insert 持 RLock 写数据（-race 验证通过）。11 个测试 + 中点/判环边界用例。
  验收达成：linkedlist 模块 11 测试全绿（含 -race）。证据 f7a6b7a

## 三、只有设计 / 完全没有

- [x] docs/ 真实内容（C3 完成）：新增 docs/notes/ 四篇笔记（array/sorts/stack/linkedlist，含 API 表、复杂度、历史坑位回归记录）+ 索引页，接入 navbar/sidebar。vuepress beta.53 在 Node 24 下构建可用。
  验收达成：`npm run docs:build` 成功，13→18 页全渲染。证据 666a1ce
- [x] queue / tree / graph / heap / hash 结构（C4 完成）：全部按「module 骨架 + 基础实现 + 测试」落地，一个结构一个 commit，并收录进 go.work：
  - queue：Queue 接口 + 数组队列（头游标）/链表队列，6 测试 —— 证据 c231579
  - tree：泛型 BST（插入/查/删三态/最值/中序），6 测试 —— 证据 08c57bf
  - heap：泛型最小堆（上浮/下沉/O(n) 建堆），6 测试 —— 证据 a64bf41
  - hash：链地址哈希表（FNV-1a + 负载因子倍增扩容），5 测试 —— 证据 d7a45a6（模块名取 hashmap 避让标准库）
  - graph：邻接表有向/无向图 + BFS/DFS（判重边/环/连通分量用例），6 测试 —— 证据 1d1007d
  验收达成：全仓 9 模块 `go test ./...` 全绿。

## 四、后续可选（新账，非本轮验收）

- [x] docs/notes/ 为 queue/tree/heap/hash/graph 五个新模块补笔记（P1014A 完成）：五篇笔记按既有风格落地（接口/API 表、复杂度、测试状态），sidebar structure 自动收录无需改配置；notes/README.md 索引表补五行并更正「九个模块」表述。
  验收达成：`npm run docs:build` 成功，18→23 页全渲染。证据 b950a7f
- [x] BST 增加高度/平衡检查；graph 增加 BFS 最短路（边权 1）（P1014A 完成）：测试先行（编译红确认缺方法）后实现。tree 加 `Height()`（空树 0）与 `IsBalanced()`（每节点左右子树高度差 ≤1），共用 `checkBalance` 单趟同时算高与平衡，O(n) 无重算；graph 加 `ShortestPath(from,to)`，BFS 记父节点首次命中即回溯，含两端点，起点即终点返回单点，顶点缺失/不可达返回 nil。
  验收达成：tree 8 测试、graph 9 测试全绿（- count=1），全仓 9 模块 `go test ./...` 全绿；两篇笔记 API 表已同步。证据 1eb38b0

- [x] 文档双语化（来源：用户 2026-10-14，全仓统一要求：**除 hello-* 外所有仓都要双语**）
  - 现状：本仓 README 为**单语**，无 `README.zh.md`
  - 按全局规矩：**英文为 README 主文档**（`README.md` 保持英文），中文转 **`README.zh.md`**
  - **两份顶部都要加切换器且必须在首屏**（标题下第一行）：
    - `README.md` → `> **English** | [中文](README.zh.md)`
    - `README.zh.md` → `> [English](README.md) | **中文**`
  - 中文版要**完整对应，不是摘要**；技术名词、命令、路径、代码块**原样保留不译**
  - docs 下的文档若有实质内容，同样需要中文版（用户原话"其他都需要双语"）
  - 验收：zh 行数与 en 相当、逐节覆盖、切换器两份都有；只改文档不碰代码
  - 落地（用户裁定约定：**英文为基名、中文转 .zh.md**）：README.md（英，重写为完整仓库说明）+ README.zh.md（中）；docs/notes 10 篇（9 笔记 + 索引）英文为基名、中文 git mv 为 `.zh.md`，22 文件全部加首屏切换器。模板样板（docs/demo、docs/guide、docs/slides、docs/README）非实质文档，未动。
  - 验收达成：11 对文件 zh/en 行数完全一致（README 65/65、notes 各篇 30~55 逐对相等），切换器逐文件校验为 H1 下第一行；`npm run docs:build` 成功 23→33 页。证据 7d6b72c

## 五、docs 补双语（用户 2026-10-16，本轮范围仅 docs/ 根目录）

- [x] docs/ 根目录文档双语化（P1016A 完成）：根目录共 2 个文档，全部配对完成。
  - `docs/slides.md`（433 行，**本文件本身已是英文**，reveal.js 幻灯片演示）→ 新增中文版 `docs/slides.zh.md`，逐页对应翻译。
  - `docs/README.md`（站点首页，中文模板）→ git mv 为 `docs/README.zh.md`，英文基名 `docs/README.md` 重写对应。
  - 4 个文件全部加语言切换器；`slides.md`/`slides.zh.md` 的切换器在 frontmatter 结束后首行（第 7 行）；`README.md`/`README.zh.md` 的切换器为正文首行（frontmatter 含 17 个 feature 卡，占 102 行，必须占据文件开头，否则 `home: true` 布局失效——此为无法"前 3 行"的唯一例外，已在回报中注明）。
  - 排除项核对：`CHANGELOG`/`CONTRIBUTING`/`LICENSE`/`CODE_OF_CONDUCT`/`SECURITY` 本仓均不存在，无需排除。
  - 本仓为 `algorithms`，非 `hello-*` 教学仓，不在免做名单，口径无冲突。
  - 验收达成：结构骨架逐行比对一致（slides 117 行骨架仅 frontmatter `title` 显示文本随语言变化；home frontmatter 键集合一致）；行数 README 对 112/112、slides 对 436/436；英文基名正文无中文（切换器行除外）；`npm run docs:build` 成功 33→35 页，dist 生成 `slides.html`/`slides.zh.html`/`index.html`/`README.zh.html` 且切换器均在产物中。证据 3fffe15
  - 环境注记：执行时发现 `node_modules` 已被清空（上轮尚可构建），用 `npm ci` 按 lock 精确还原后构建通过，`package.json`/`package-lock.json` 未被改动。
- [x] docs 子目录双语化（R1016E 完成）：`docs/demo/` 5 篇 + `docs/guide/` 5 篇共 10 对全部落地，沿用「英文为基名、中文转 `.zh.md`」约定（git mv 保留历史），20 文件全部加首屏切换器。
  - 中文侧相对链接改指 `.zh.md` 同伴（demo 索引 4 处、guide 索引 2 处、bar/foo 各 1 处）；`markdown.md` 的 `@include(./README.md{11-17})` 因切换器插入下移 2 行，两侧同步改为 `{13-19}` 并经 dist 产物验证渲染出的就是 4 条目录项。
  - 代码块/命令/URL 原样保留（flow 图与 demo 的中文节点串、JSON `"label": "散点数据集"`、`#标题列表` 锚点 URL 等均不译）；英文内链锚点改为 `#custom-containers`，dist 中 `id="custom-containers"` 已验证存在。
  - 验收达成：10 对行数逐对相等（20/20、45/45、18/18、75/75、433/433、19/19、16/16、9/9、16/16、9/9）；切换器逐文件为正文首行（encrypt 对在 H1 标题下第一行）；英文正文中文残留仅切换器行与代码围栏/URL 内（脚本逐行校验）；`npm run docs:build` 成功 35→45 页。证据 f75a95a
  - 遗留（配置级，非文档改动可解）：sidebar `children: "structure"` 使英/中两套条目并列，且英文页的 上一页/下一页 可能落到中文页——需动 `docs/.vuepress/` 配置（locale 路由或 sidebar 白名单）才能解决，本轮未动。

