/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

import (
	"fmt"
	"os"
	"testing"

	"github.com/prodanlabs/kaadm/app/utils"
)

func TestWaitEtcdReplicasetInDesired(t *testing.T) {
	restConfig, err := utils.RestConfig("/home/prodan/.kube/config")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kubeClient, err := utils.NewClientSet(restConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	EtcdReplicas = 3
	if err = WaitEtcdReplicasetInDesired(kubeClient, "karmada-system", "app=etcd", 50); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestWaitPodReady(t *testing.T) {
	restConfig, err := utils.RestConfig("/home/prodan/.kube/config")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kubeClient, err := utils.NewClientSet(restConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = WaitPodReady(kubeClient, "karmada-system", "app=etcd", 50); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}