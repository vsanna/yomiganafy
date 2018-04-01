https://yomiganafy.appspot.com/yomiganafy?name=山田健二

## setup
- さいしんのgcloudをいれる
- dev_appserver.py app/app.yaml てlocal server が起動する


## よくあるもんたい



## memo

gcloudをいれるとdev_appserver.pyがinstallされ、そこへのパスが通される
`$ dev_appserver.py app.yaml` だけでいい。もしsetupないと言われたらpython2系にする

手元のプロジェクトをgaeでbuildするところをシュミレートし、それをlocalでアクセスできるようにしてくれる


yomiganafy/main.goにしてしまうとmain.go内のimportと、gae側が行うimportが衝突するとかほざく
> Failed parsing input: app file root.go conflicts with same file imported from GOPATH

なのでapp/に切り出す

http://otiai10.hatenablog.com/entry/2016/05/08/022515



```bash

$ dev_appserver.py app/app.yaml

# gcloud app deploy app/app.yaml
```
