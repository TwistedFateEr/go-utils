#!/bin/bash


cat << EOF >> /etc/hosts
192.168.147.253 master
EOF

cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
modprobe br_netfilter
sysctl -p /etc/sysctl.d/k8s.conf



cat <<EOF > /etc/docker/daemon.json
{
    "registry-mirrors": ["https://v16stybc.mirror.aliyuncs.com"],
    "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF

systemctl daemon-reload
systemctl restart docker


cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

yum clean all
yum -y makecache


yum list kubelet --showduplicates | sort -r
yum install -y kubelet-1.14.2
yum install -y kubeadm-1.14.2
yum install -y kubectl-1.14.2

systemctl enable kubelet && systemctl start kubelet

#kubelet命令补全
echo "source <(kubectl completion bash)" >> ~/.bash_profile
source .bash_profile



url=registry.cn-hangzhou.aliyuncs.com/google_containers
version=v1.21.0
images=(`kubeadm config images list --kubernetes-version=$version|awk -F '/' '{print $2}'`)
for imagename in ${images[@]} ; do
  docker pull $url/$imagename
  docker tag $url/$imagename k8s.gcr.io/$imagename
  docker rmi -f $url/$imagename
done


kubeadm init --apiserver-advertise-address 192.168.147.253 --pod-network-cidr=10.244.0.0/16

wget https://www.backendcloud.cn/files/centos7-7-k8s-3nodes/kube-flannel.tar.gz
tar -axvf kube-flannel.tar.gz

kubectl apply -f kube-flannel.yml


#kubeadm join 192.168.147.253:6443 --token rys0lb.xht0u9sxvu7lskdq \
#    --discovery-token-ca-cert-hash sha256:de866c666c435d9c9b74ae0b78a6df2357639551a0a116dc5b78002c59dc1167