#!/bin/sh

#yum-config-manager --add-repo \
#  https://download.docker.com/linux/centos/docker-ce.repo

cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

swapoff -a

yum -y update
yum install -y yum-utils device-mapper-persistent-data lvm2
yum install -y lshw
yum install -y firewalld
yum install -y docker kubeadm kubectl kubernetes-cli ntp

systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet
systemctl enable ntpd && systemctl start ntpd
systemctl unmask firewalld
systemctl enable firewalld
systemctl start firewalld

yum -y update
systemctl daemon-reload
systemctl restart kubelet

modprobe uio_pci_generic
lsmod | grep uio
lshw -class network -businfo

sysctl -w vm.nr_hugepages=2048

wget https://github.com/contiv/vpp/raw/master/k8s/setup-node.sh
wget https://github.com/contiv/vpp/raw/master/k8s/contiv-vpp.yaml
wget https://raw.githubusercontent.com/DPDK/dpdk/master/usertools/dpdk-setup.sh

chmod +x setup-node.sh


chown $USER:$USER /var/run/docker.sock
docker images

firewall-cmd --permanent --add-port=12379-12380/tcp

mkdir /etc/vpp/
cp ./contiv-vswitch.conf /etc/vpp/

# setup alias and completion
source /usr/share/bash-completion/bash_completion
echo 'source <(kubectl completion bash)' >>~/.bashrc
echo 'alias k=kubectl' >>~/.bashrc
echo 'complete -F __start_kubectl k' >>~/.bashrc

echo 'now create your /etc/vpp/contiv-vswitch.conf file according to the guide here https://github.com/contiv/vpp/blob/master/docs/setup/VPP_CONFIG.md#creating-vpp-startup-configuration '

