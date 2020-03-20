#!/bin/sh

sudo su

yum install -y yum -y install lshw
modprobe uio_pci_generic
lsmod | grep uio
lshw -class network -businfo


wget https://github.com/contiv/vpp/raw/master/k8s/setup-node.sh

wget https://github.com/contiv/vpp/raw/master/k8s/contiv-vpp.yaml

chmod +x setup-node.sh
chmod +x bootstrap_centos.sh

./bootstrap_centos.sh



