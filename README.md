# king-cri-client
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/open-kingfisher/king-cri-client)](https://goreportcard.com/report/github.com/open-kingfisher/king-cri-client)
Kubernetes CRI Client

## 说明

- 从Kubelet中拷贝出来的CRI Client，此Client实现了 CRI Image和Runtime的grpc client
- 可以在装有kubelet的节点上运行，用于管理容器和镜像
- 此包依赖于k8s.io/kubernetes v1.13.1，更高版本的k8s.io/kubernetes已经不在支持go module安装，所有方法需要重新实现
