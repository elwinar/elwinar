Vagrant.configure(2) do |config|
	config.vm.box = "elwinar/golang"
	config.vm.hostname = "elwinar"
	
	config.vm.network "forwarded_port", 
			guest: 8080, 
			host: 8080,
			auto_correct: true
	
	config.vm.network "private_network",
			type: "dhcp"
	
	config.vm.synced_folder ".", 
			"/home/vagrant/src/github.com/elwinar/elwinar", 
			nfs: true,
			mount_options: ["actimeo=1"]
	
	config.vm.provision "shell", inline: <<EOS
		pacman -Suyy --noconfirm
		pacman -S --noconfirm nodejs npm tmux most
		npm install -g bower gulp
EOS
end
