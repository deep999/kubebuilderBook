mkdir -p book/src/tutorial
cd book/src/tutorial
go mod init tutorial
kubebuilder init --domain tutorial.kubebuilder.io
kubebuilder create api --group batch --version v1 --kind CronJob

#add test
cd controllers 
/home/uros/go/bin/ginkgo generate job