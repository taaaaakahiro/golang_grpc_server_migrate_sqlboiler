
# golang_grpc_server

## init
```sh
# ローカルにプライベートリポジトリをimportする
## https方式
$ git config --global url."https://${{ GITHUB_PERSONAL_ACESS_TOKEN }}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
git config --global url."https://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:x-oauth-basic@github.com/".insteadOf "https://github.com/"
## ssh
$ git config --global url."git@github.com:".insteadOf "https://github.com/"

# 環境設定 go get2回目以降
$ go env -w GOPRIVATE="github.com/<アカウント名>/*"

# funcをget
$ go get github.com/taaaaakahiro/golang_grpc_proto

```

## grpc command
```sh
$ grpcurl -plaintext localhost:<PORT> list # list registered service
$ grpcurl -plaintext localhost:8080 list UserService # list registered method
```

## .gitconfig
```sh
## ssh
[url "git@github.com:"]
    insteadOf = https://github.com/
## https
[url "https://ghp_XXXX:x-oauth-basic@github.com"]
    insteadOf = https://github.com
```

## Article
 - grpc
   - https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/server
 - slog
   - https://gihyo.jp/article/2023/02/tukinami-go-04