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
1. [exeファイル](https://github.com/fujimaru240/mvfile/blob/master/mvfile.exe)をダウンロードする。
1. コマンドプロンプトを開く
    - ［Windows］＋［R］キーを押し、［ファイル名を指定して実行］ダイアログを開く。 次に［名前］入力ボックスに「cmd」と入力して、［Enter］キーを押す
1. コマンドプロンプトで、ダウンロードしたファイル（ツール）の場所へ移動する
    ```bash
    cd [ダウンロードしたファイルの場所]

    例：C:\Users\xxx\Downloads    
    ```
1. コマンドプロンプトで、ツールを実行する
    ```bash
    mvfile.exe [元の親フォルダ] [コピー先のフォルダ]

    例：mvfile.exe C:\Users\input C:\Users\output
    ```