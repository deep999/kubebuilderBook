package controllers

import (
	"fmt"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
	crdtypes "tutorial/api/v1"
)

var _ = Describe("Job", func() {
	const timeout = time.Second * 180
	const interval = time.Second * 2

	Describe("SAMPLE tests", func() {
		It("Schould be OK", func() {
			spec := crdtypes.CronJobSpec{Foo: "testFOO"}
			key := types.NamespacedName{
				Name:      "funnyname",
				Namespace: "default",
			}
			job := &crdtypes.CronJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      key.Name,
					Namespace: key.Namespace,
				},
				Spec: spec,
			}

			Expect(k8sClient.Create(context.Background(), job)).Should(Succeed())
			Expect(k8sClient.Get(context.Background(), key, job)).Should(BeNil())
		})
	})

})
