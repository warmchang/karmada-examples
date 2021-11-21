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
package certs

import (
	"fmt"
	"os"
	"testing"
)

func TestCertificateGeneration(t *testing.T) {

	c := &Config{
		PkiPath:                     "./test-Certs-tmp",
		Namespace:                   "karmada-system",
		EtcdStatefulSetName:         "etcd",
		EtcdServiceName:             "etcd",
		EtcdReplicas:                3,
		KArmadaMasterIP:             "192.168.0.1",
		KArmadaApiServerServiceName: "karmada-apiserver",
		KArmadaWebhookServiceName:   "karmada-webhook",
		FlagsExternalIP:             "172.16.0.2,172.16.0.3",
	}
	filesName, err := c.CertificateGeneration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(filesName)
}