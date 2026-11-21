import { defineConfig } from 'vitepress'

const notes = [
  { text: 'Overview', link: '/notes/' },
  { text: 'Array', link: '/notes/array' },
  { text: 'Sorting', link: '/notes/sorts' },
  { text: 'Stack', link: '/notes/stack' },
  { text: 'Linked List', link: '/notes/linkedlist' },
  { text: 'Queue', link: '/notes/queue' },
  { text: 'Binary Search Tree', link: '/notes/tree' },
  { text: 'Binary Heap', link: '/notes/heap' },
  { text: 'Hash Map', link: '/notes/hash' },
  { text: 'Graph', link: '/notes/graph' },
]

const zhNotes = [
  { text: '总览', link: '/zh/notes/' },
  { text: '数组', link: '/zh/notes/array' },
  { text: '排序', link: '/zh/notes/sorts' },
  { text: '栈', link: '/zh/notes/stack' },
  { text: '链表', link: '/zh/notes/linkedlist' },
  { text: '队列', link: '/zh/notes/queue' },
  { text: '二叉搜索树', link: '/zh/notes/tree' },
  { text: '二叉堆', link: '/zh/notes/heap' },
  { text: '哈希表', link: '/zh/notes/hash' },
  { text: '图', link: '/zh/notes/graph' },
]

const enSidebar = {
  '/notes/': [{ text: 'Notes', items: notes }],
}

const zhSidebar = {
  '/zh/notes/': [{ text: '笔记', items: zhNotes }],
}

export default defineConfig({
  lang: 'en-US',
  title: 'Algorithm Notes',
  description: 'Data structures and algorithms in Go',
  base: '/algorithms/',
  lastUpdated: true,
  ignoreDeadLinks: true,
  themeConfig: {
    logo: '/logo.svg',
    siteTitle: 'Algorithm Notes',
  },
  locales: {
    root: {
      label: 'English',
      lang: 'en-US',
      title: 'Algorithm Notes',
      description: 'Data structures and algorithms in Go',
      themeConfig: {
        nav: [
          { text: 'Home', link: '/' },
          { text: 'Notes', link: '/notes/', activeMatch: '/notes/' },
        ],
        sidebar: enSidebar,
        outline: { level: [2, 3] },
        docFooter: { prev: 'Previous', next: 'Next' },
      },
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN',
      link: '/zh/',
      title: '算法笔记',
      description: 'Go 数据结构与算法',
      themeConfig: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: '笔记', link: '/zh/notes/', activeMatch: '/zh/notes/' },
        ],
        sidebar: zhSidebar,
        outline: { level: [2, 3], label: '本页目录' },
        docFooter: { prev: '上一页', next: '下一页' },
      },
    },
  },
})
