## demo & how to use

- endpoint: https://yomiganafy.appspot.com/yomiganafy
- method: POST
- params:
    - name(required): string
    

```bash
$ curl -X POST -d name="山田健二" https://yomiganafy.appspot.com/yomiganafy | jq     

#  {
#    "result": {
#      "Name": "山田健二",
#      "Yomi": "ヤマダ ケンジ"
#    }
#  }
```

## setup
- 最新のgcloudを入れる

```bash
# 移動
$ cd path/to/yomiganafy

# ローカルサーバー起動
$ dev_appserver.py app/app.yaml
```


## deploy
```bash
$ gcloud config configurations activate yomiganafy

$ gcloud app deploy app/app.yaml
```

## TODO
- [ ] mecabの辞書を人名に絞る
- [ ] 辞書を追加する手順を整理する
- [ ] バリデーション
- [ ] エラーハンドリング
- [ ] responseもstructで管理する
- [ ] やたら遅いのでパフォーマンスチューニングの仕方(計測の仕方)を調べる
- [ ] ロジックの置き場所整理/ディレクトリ構成の整理
- [ ] app.yamlの設定値解釈

## よくある問題
- dev_appserver.pyでsetup moduleが見つからない
    - gcloudはpython2系前提なのでpyenvで変更する
- Failed parsing input: app file root.go conflicts with same file imported from GOPATH
    - app/にmain.goをおいておけば多分大丈夫なはず
    - http://otiai10.hatenablog.com/entry/2016/05/08/022515
- handlersが見つからないって
    - projectディレクトリは`$GOPATH/src/github.com/vsanna/yomiganafy`におくこと
- 依存関係に有るpkgが無いと言われる
    - それに依存している外部ライブラリを$GOPATH/srcから`rm -rf`した後に再度`go get some_pkg`すると解決する(ことがある)

