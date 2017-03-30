package main

import "fmt"

func main() {
	a := []int{57, 68, 59, 52}

	insertSort(a)

	fmt.Println(a)
}

func insertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		var j int
		//		for j = i - 1; j >= 0 && tmp < a[j]; j-- {
		//			a[j+1] = a[j]
		//		}

		for j = i - 1; j >= 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
			} else {
				break
			}

		}
		a[j+1] = tmp
	}
	return a
}



beta.kubernetes.io/arch=amd64,   beta.kubernetes.io/instance-type=rancher,
beta.kubernetes.io/os=linux,

failure-domain.beta.kubernetes.io/region=Region1,
failure-domain.beta.kubernetes.io/zone=FailureDomain1,

io.rancher.host.docker_version=1.12,
io.rancher.host.linux_kernel_version=3.13,

kubernetes.io/hostname=node,orchestration=true



beta.kubernetes.io/arch=amd64,
beta.kubernetes.io/os=linux,
kubernetes.io/hostname=10.10.101.83


beta.kubernetes.io/instance-type=n1-standard-2,

failure-domain.beta.kubernetes.io/region=hc-central1,
failure-domain.beta.kubernetes.io/zone=hc-central1-a,

kubernetes.io/hostname=kubernetes-minion-9vlv

kubectl label nodes 10.10.101.83 failure-domain.beta.kubernetes.io/region=hc failure-domain.beta.kubernetes.io/zone=hc-aasd --overwrite
kubectl label nodes 10.10.101.84 failure-domain.beta.kubernetes.io/region=hc failure-domain.beta.kubernetes.io/zone=hc-aasd --overwrite

kubectl label nodes 10.10.101.85 failure-domain.beta.kubernetes.io/region=hc failure-domain.beta.kubernetes.io/zone=hc-bvvvv --overwrite
kubectl label nodes 10.10.101.86 failure-domain.beta.kubernetes.io/region=hc failure-domain.beta.kubernetes.io/zone=hc-bvvvv --overwrite


kubectl label nodes 10.10.101.83 failure-domain.beta.kubernetes.io/region -
kubectl label nodes 10.10.101.84 failure-domain.beta.kubernetes.io/region -
                                                                         
kubectl label nodes 10.10.101.85 failure-domain.beta.kubernetes.io/region -
kubectl label nodes 10.10.101.86 failure-domain.beta.kubernetes.io/region -


kubectl label nodes 10.10.101.83 failure-domain.beta.kubernetes.io/zone -
kubectl label nodes 10.10.101.84 failure-domain.beta.kubernetes.io/zone -
                                                                         
kubectl label nodes 10.10.101.85 failure-domain.beta.kubernetes.io/zone -
kubectl label nodes 10.10.101.86 failure-domain.beta.kubernetes.io/zone -