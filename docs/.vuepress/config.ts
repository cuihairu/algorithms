import { defineUserConfig } from "vuepress";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/",

  locales: {
    "/": {
      lang: "zh-CN",
      title: "文档演示",
      description: "算法文档",
    },
  },

  theme,

  shouldPrefetch: false,
});
