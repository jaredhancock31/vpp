#!/bin/sh


yum install -y yum -y install lshw
modprobe uio_pci_generic
lsmod | grep uio
lshw -class network -businfo


wget https://github.com/contiv/vpp/raw/master/k8s/setup-node.sh

wget https://github.com/contiv/vpp/raw/master/k8s/contiv-vpp.yaml

wget https://raw.githubusercontent.com/DPDK/dpdk/master/usertools/dpdk-setup.sh

chmod +x setup-node.sh
chmod +x bootstrap_centos.sh

./bootstrap_centos.sh

yum -y update

docker images

yum install firewalld
systemctl unmask firewalld
systemctl enable firewalld
systemctl start firewalld

firewall-cmd --permanent --add-port=12379-12380/tcp

