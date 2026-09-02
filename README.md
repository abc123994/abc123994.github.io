# 日記

這是一個由 Git 維護的靜態日記網站。

## 新增日記

建立日期資料夾與 Markdown 檔案：

```text
diary/YYYY-MM-DD/title.md
```

檔案內容：

```markdown
---
title: XXX
date: YYYY-MM-DD HH:MM:SS
summary: XXX
---

XXX
```

圖片可放在同一個資料夾，再用 Markdown 引用：

```markdown
![照片](photo.jpg)
```

提交並推送後，網站會自動建置發布。首頁會依日期倒序列出日記。

## 本機預覽

```bash
bundle exec jekyll serve
```
