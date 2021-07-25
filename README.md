# awecloud-sysoperator

```bash
# operator-sdk
operator-sdk-v0.17 generate k8s

# tag
git tag v0.5.2 -f
git push origin v0.5.2 -f

# protos
go-to-protobuf \
  --proto-import="./protos" \
  --packages="github.com/mengkzhaoyun/awecloud-sysoperator/pkg/apis/bcc/v1" \
  --go-header-file "./hack/boilerplate/boilerplate.generatego.txt"
```
