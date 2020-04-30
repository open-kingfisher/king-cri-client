package main

import (
	"fmt"
	"github.com/open-kingfisher/king-cri-client/remote"
	runtimeapi "k8s.io/kubernetes/pkg/kubelet/apis/cri/runtime/v1alpha2"
	"time"
)

func main() {
	endpoint := "unix:///var/run/dockershim.sock"
	if client, err := remote.NewRemoteImageService(endpoint, 60*time.Second); err != nil {
		fmt.Println("remote.NewRemoteImageService error:", err)
		return
	} else {
		images, err := client.ListImages(&runtimeapi.ImageFilter{})
		if err != nil {
			fmt.Println("client.ListImages error:", err)
			return
		}
		for _, i := range images {
			fmt.Printf("ID:%s RepoTags:%s Size:%d\n", i.Id, i.RepoTags, i.Size_)
		}
	}
}
