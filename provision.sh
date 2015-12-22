#! /usr/bin/env sh

curl -fsS https://www.archlinux.org/mirrorlist/?country=all > /tmp/mirrolist
grep '^#Server' /tmp/mirrolist | sort -R | head -n 50 | sed 's/^#//' > /tmp/mirrolist.50
rankmirrors -v /tmp/mirrolist.50 | tee /etc/pacman.d/mirrorlist

sudo pacman -S --noconfirm go gcc git subversion mercurial
echo "export GOPATH=/home/vagrant" >> /home/vagrant/.bashrc && export GOPATH=/home/vagrant
echo "export PATH=$GOPATH/bin:$PATH" >> /home/vagrant/.bashrc && export PATH=$GOPATH/bin:$PATH
chown vagrant:vagrant -R /home/vagrant/src
go get github.com/jteeuwen/go-bindata/...

sudo pacman -S --noconfirm nodejs npm
npm install -g bower gulp

