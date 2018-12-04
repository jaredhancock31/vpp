#!/bin/sh
echo "Starting Contiv VPP ..."
echo ""

cd ../vagrant
sh ./vagrant-dummy-start

echo "Setting Kubernetes proxy ..."
echo ""
vagrant ssh k8s-master -c "kubectl proxy --port=8080 &"

echo "Setting VMs port forwarding ..."

cd ../ui
sh ./setVMs.sh

echo "Port forwarding has been set."

echo "Starting vagrant ..."

cd ./vagrant
vagrant up

echo ""
echo "Application has been deployed. Open URL http://localhost:4200 in your browser with disabled security."