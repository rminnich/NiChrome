#!/bin/bash
if [ -z "${GOPATH}" ]; then
	export GOPATH=/home/travis/gopath
fi
set -e

# make sure we have the right one.
wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64
mv dep-linux-amd64 dep
chmod +x dep

echo "Check vendored dependencies"
(./dep status)

(cd usb && go build .)
./usb/usb --apt=true --fetch=true --dev=/dev/null

# in case of emergency break glass
# cpio -ivt < initramfs.linux_amd64.cpio
# cpio -ivt < linux-stabld/initramfs.linux_amd64.cpio
