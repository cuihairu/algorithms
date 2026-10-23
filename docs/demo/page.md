---
# This is the article title
title: Page config
# This is the page icon
icon: page
# This is the sidebar order
order: 1
# Set the author
author: Ms.Hope
# Set the writing date
date: 2020-01-01
# A page can have multiple categories
category:
  - Guide
# A page can have multiple tags
tag:
  - Page config
  - Guide
# This page will be pinned in the article list
sticky: true
# This page will appear in the article favorites
star: true
# You can customize the footer
footer: This is a test footer
# You can customize the copyright info
copyright: No copyright
---

> **English** | [中文](/zh/demo/page)

Content before the `more` comment is treated as the article excerpt.

<!-- more -->

## Page info

You can set page information in the Markdown Frontmatter.

- The author is set to Ms.Hope.
- The writing date is January 1st, 2020
- The category is "Guide"
- The tags are "Page config" and "Guide"

## Page content

You can freely write your Markdown here.

::: tip

- You can place images and Markdown files together, but you need to reference them with **relative links** `./`.

- For images in the `.vuepress/public` folder, please reference them with absolute links `/`.

:::

The theme contains a custom badge that can be used:

> There should be a badge in dark blue at the end of the text. <Badge text="badge" color="#242378" />

## Page structure

This page should contain:

- [Pathnav](https://vuepress-theme-hope.github.io/v2/zh/guide/layout/breadcrumb.html)
- [Title and page info](https://vuepress-theme-hope.github.io/v2/zh/guide/feature/page-info.html)
- [TOC (list of article titles)](https://vuepress-theme-hope.github.io/v2/zh/guide/layout/page.html#标题列表)
- [Contributors, last updated time and other page meta info](https://vuepress-theme-hope.github.io/v2/guide/feature/meta.html)
- [Comment](https://vuepress-theme-hope.github.io/v2/zh/guide/feature/comment.html)
- [Navbar](https://vuepress-theme-hope.github.io/v2/zh/guide/layout/navbar.html)
- [Sidebar](https://vuepress-theme-hope.github.io/v2/zh/guide/layout/sidebar.html)
- [Footer](https://vuepress-theme-hope.github.io/v2/zh/guide/layout/footer.html)
- Back to top button

You can customize them through theme options and page Frontmatter.
