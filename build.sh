go build -o terraform-provider-teststack
rm $GOPATH/bin/terraform-provider-teststack
cp $GOPATH/src/github.com/xuyaming0800/terraform-provider-teststack/terraform-provider-teststack $GOPATH/bin/
# 如果terraform版本高于或者等于0.13
# 需要执行如下三条指令 来映射CLI到本地路径
# 如果小于此版本可以不做这三个操作
# mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/teststack/0.0.1/darwin_amd64/
# rm ~/.terraform.d/plugins/registry.terraform.io/hashicorp/teststack/0.0.1/darwin_amd64/terraform-provider-teststack_v0.0.1
# cp $GOPATH/src/github.com/xuyaming0800/terraform-provider-teststack/terraform-provider-teststack ~/.terraform.d/plugins/registry.terraform.io/hashicorp/teststack/0.0.1/darwin_amd64/terraform-provider-teststack_v0.0.1

rm terraform-provider-teststack