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
	"context"

	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

// IfNotExist
func (i *InstallOptions) CreateNamespace() error {
	namespaceClient := i.KubeClientSet.CoreV1().Namespaces()
	namespaceList, err := namespaceClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, nsList := range namespaceList.Items {
		if i.Namespace == nsList.Name {
			klog.Infof("Namespace %s already exists.", i.Namespace)
			return nil
		}
	}

	n := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: i.Namespace,
		},
	}

	_, err = namespaceClient.Create(context.TODO(), n, metav1.CreateOptions{})
	if err != nil {
		return errors.Errorf("Create namespace %s failed: %v\n", i.Namespace, err)
	}
	klog.Infof("Create Namespace '%s' successfully.", i.Namespace)
	return nil
}