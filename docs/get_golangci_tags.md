# 최신 가져오기

```bash
curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest | jq .tag_name

# 
curl -s https://api.github.com/repos/golang/go/releases/latest 

git ls-remote --tags https://github.com/golang/go.git | grep "refs/tags/go"

git ls-remote --tags https://github.com/golang/go.git | awk '{print $2}' | sed 's|refs/tags/||'  | grep go | sed 's|go||' | grep -v "beta" | grep -v "rc" | sort -V | tail -n 1

curl -s https://go.dev/dl/?mode=json | jq -r '.[0].version'
```
