#!/bin/bash
export GOPATH=/home/travis/gopath
set -e

# go get is not helpful.
# go get -d github.com/u-root/u-root
U=$GOPATH/src/github.com/u-root
(cd $U && git clone https://github.com/u-root/u-root && cd u-root/bb && go build . && ./bb)

rm -rf linux-stable
git clone --depth 1 -b v4.12.7 git://git.kernel.org/pub/scm/linux/kernel/git/stable/linux-stable.git
cp /tmp/initramfs.linux_amd64.cpio linux-stable
(cd linux-stable && cp ../CONFIG .config && make oldconfig &&make -j 8)

git clone -b 3f3a496a23088731e4ab5654b02fbc13a6881c65 https://chromium.googlesource.com/chromiumos/platform/vboot_reference
(cd vboot_reference && make)

