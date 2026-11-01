# algorithms

<p align="center">
  <a href="https://github.com/cuihairu/algorithms/releases"><img alt="status" src="https://img.shields.io/badge/status-in%20development-orange"></a>
  <a href="https://cuihairu.github.io/algorithms/"><img alt="docs" src="https://img.shields.io/badge/docs-english-blue"></a>
</p>

> **English** | [中文](README.zh.md)

Data structures and algorithms implemented in Go, split into independent Go modules, each with unit tests, plus a bilingual VitePress documentation site.

## Modules

| Module | Package path | Contents | Tests |
| --- | --- | --- | --- |
| array | `array/go` | generic array wrapper with bounds checks | 4 |
| sorts | `sorts/go` | bubble / insertion / selection / merge / quick / counting / bucket sort | 7 |
| stack | `stack/go` | `Stack` interface + array stack / linked list stack | 6 |
| linkedlist | `linkedlist/go` | singly linked list with head sentinel + RWLock + fast/slow pointers | 11 |
| queue | `queue/go` | `Queue` interface + array queue (head cursor) / linked list queue | 6 |
| tree | `tree/go` | generic BST (insert / search / remove / min / max / in-order / height / balance) | 8 |
| heap | `heap/go` | generic min-heap (sift-up / sift-down / O(n) heapify) | 6 |
| hash | `hash/go` (module `hashmap`) | chaining hash map with FNV-1a + load-factor doubling resize | 5 |
| graph | `graph/go` | adjacency-list directed/undirected graph + BFS / DFS / shortest path | 9 |

## Layout

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
├── docs/          # VitePress site (a note for every module)
├── go.work        # all nine modules in one workspace
└── package.json   # docs scripts
```

## Running the tests

All nine modules are included in `go.work`; run inside each module directory:

```bash
cd <module>/go && go test ./...
```

To run the linked list tests with the race detector:

```bash
cd linkedlist/go && go test -race ./...
```

## Documentation site

`docs/` is a VitePress site with a note for every module; English lives at the site root and Simplified Chinese under `/zh/`.

```bash
pnpm install
pnpm run docs:dev     # local dev server
pnpm run docs:build   # build into docs/.vitepress/dist
```

## Go version

`go.work` targets Go 1.19.
