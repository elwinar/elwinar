#! /usr/bin/env sh

# Ensure the mirrors are up-to-date
curl -fsS https://www.archlinux.org/mirrorlist/?country=all > /tmp/mirrolist
grep '^#Server' /tmp/mirrolist | sort -R | head -n 50 | sed 's/^#//' > /tmp/mirrolist.50
rankmirrors -v /tmp/mirrolist.50 | tee /etc/pacman.d/mirrorlist

# Update the system globally
pacman -Suyy --noconfirm

# General tools
pacman -S --noconfirm make

# Golang tools
pacman -S --noconfirm go gcc git subversion mercurial upx ucl
echo "export GOPATH=/home/vagrant" >> /home/vagrant/.bashrc && export GOPATH=/home/vagrant
echo "export PATH=$GOPATH/bin:$PATH" >> /home/vagrant/.bashrc && export PATH=$GOPATH/bin:$PATH
chown vagrant:vagrant -R /home/vagrant/src
go get github.com/jteeuwen/go-bindata/...
go get github.com/pwaller/goupx
go get github.com/elwinar/rambler

# Javascript tools
pacman -S --noconfirm nodejs npm
npm install -g bower gulp

# Docker tools
pacman -S --noconfirm docker
systemctl enable docker
systemctl start docker
gpasswd -a vagrant docker
