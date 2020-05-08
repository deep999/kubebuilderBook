
> make test
go: creating new go.mod: module tmp
go: found sigs.k8s.io/controller-tools/cmd/controller-gen in sigs.k8s.io/controller-tools v0.2.5
/home/user/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
/home/user/go/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
go test ./... -coverprofile cover.out
?       tutorial        [no test files]
?       tutorial/api/v1 [no test files]
ok      tutorial/controllers    4.675s  coverage: 0.0% of statements


> /home/user/go/bin/ginkgo controllers
Running Suite: Controller Suite
===============================
Random Seed: 1588969443
Will run 1 of 1 specs

•

Ran 1 of 1 Specs in 5.074 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 8.627214577s
Test Suite Passed