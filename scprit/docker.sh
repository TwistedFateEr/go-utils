#!/bin/bash


function echoEmpty() {
    echo ""
    echo ""
    echo ""
}

wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
yum clean all
yum makecache
yum update

# ssh-copy-id -i .ssh/id_rsa.pub root@

hostnamectl set-hostname master

yum install -y vim wget yum-utils device-mapper-persistent-data lvm2 bash-completion
source /etc/profile.d/bash_completion.sh

systemctl stop firewalld
systemctl disable firewalld
setenforce 0
sed -i 's/SELINUX=.*/SELINUX=disabled/g' /etc/selinux/config

echoEmpty

cat /etc/selinux/config

echoEmpty

swapoff -a
sed -i.bak '/swap/s/^/#/' /etc/fstab

echoEmpty
cat  /etc/fstab


yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum list docker-ce --showduplicates | sort -r
yum install -y docker-ce-18.09.6 docker-ce-cli-18.09.6 containerd.io
systemctl start docker
systemctl enable docker
mkdir -p /etc/docker
cat <<EOF > /etc/docker/daemon.json
{
  "registry-mirrors": ["https://w6pljua0.mirror.aliyuncs.com"]
}
EOF

docker --version
docker info

