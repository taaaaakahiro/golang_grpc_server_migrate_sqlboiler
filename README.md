
# golang_grpc_servera

## init
```sh
# ローカルにプライベートリポジトリをimportする
## https方式
$ git config --global url."https://${{ GITHUB_PERSONAL_ACESS_TOKEN }}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
## ssh
$ git config --global url."git@github.com:".insteadOf "https://github.com/"

# 環境設定 go get2回目以降
$ go env -w GOPRIVATE="github.com/<アカウント名>/*"

```

# .gitconfig
```sh
## ssh
[url "git@github.com:"]
    insteadOf = https://github.com/
## https
[url "https://ghp_XXXX:x-oauth-basic@github.com"]
    insteadOf = https://github.com
```
