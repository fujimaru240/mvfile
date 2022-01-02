# ファイルコピーツール
- 対象フォルダ配下のサブフォルダに格納されたファイルを、指定の出力先フォルダへコピーするツール
- 例
    ```
    input
    ├sub1
      ├file1.png
      ├file2.png
    ├sub2
      ├file1.png
    ├sub3
      ├file1.png
      ├file3.png
    
    ↓

    output
    ├file1.png      (元: sub1/file1.png)
    ├file1_1.png    (元: sub2/file1.png)
    ├file1_2.png    (元: sub3/file1.png)
    ├file2.png
    ├file3.png
    ```

## 使用方法
1. コマンドプロンプトを開く