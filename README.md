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

## 閱讀導覽

左側索引維持最新到最舊，方便追蹤更新。`About` 是序章，正文結尾提供「上一則／下一則」導覽，閱讀順序為 `About → 最舊日記 → 最新日記`；新增日記會自動接到閱讀序列末端。

手機版會優先顯示日記正文；點擊頂部的 `INDEX` 可從左側滑出文章索引。選擇文章、點擊遮罩或按下 `Esc` 都會關閉索引，切換後會回到文章頂部。

## 本機預覽

```bash
bundle exec jekyll serve
```
