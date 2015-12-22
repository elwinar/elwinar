Vagrant.configure(2) do |config|
	config.vm.box = "kaorimatz/archlinux-x86_64"
	config.vm.hostname = "elwinar"
	config.vm.network "forwarded_port", guest: 8080, host: 8080, auto_correct: true
	config.vm.synced_folder ".", "/home/vagrant/src/github.com/elwinar/elwinar"
	config.vm.provision "shell", path: "provision.sh"
end
