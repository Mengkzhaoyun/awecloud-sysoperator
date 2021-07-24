# proto

```bash
go install k8s.io/kubernetes/vendor/k8s.io/code-generator/cmd/go-to-protobuf
go install k8s.io/kubernetes/vendor/k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo

go-to-protobuf \
  --proto-import="${KUBE_ROOT}/vendor" \
  --proto-import="${KUBE_ROOT}/third_party/protobuf" \
  --packages="$(IFS=, ; echo "${APIROOTS[*]}")" \
  --go-header-file "${KUBE_ROOT}/hack/boilerplate/boilerplate.generatego.txt" \
  "$@"
```
