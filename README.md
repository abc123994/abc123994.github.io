# 日記

這個 GitHub Pages 網站是唯讀日記；Git 是編輯入口。

## 新增一篇日記

建立日期資料夾與 Markdown 檔案：

```text
diary/2026-09-02/charging-cable.md
```

檔案內容使用這個最小格式：

```markdown
---
title: 買了一條充電線
date: 2026-09-02 23:41:00
summary: 下午才起床，出門買了條充電線。
---

下午才起床。

出門買了條充電線。
```

圖片可以放在同一個資料夾，並從 Markdown 引用：

```markdown
![照片](charging-cable.jpg)
```

完成後 `git add`、`git commit`、`git push`；GitHub Actions 會建置並發布。首頁會自動依 `date` 倒序列出文章。

## 封存

轉置前的舊站封存於 Git tag：`archive/pre-diary-2026-09-02`。

## 本機預覽

```bash
bundle exec jekyll serve
```
