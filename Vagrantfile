Vagrant.configure(2) do |config|
	config.vm.box = "elwinar/golang"
	config.vm.hostname = "elwinar"
	
	config.vm.network "forwarded_port", 
		guest: 8080, 
		host: 8080,
		auto_correct: true
	
	config.vm.synced_folder ".", 
		"/home/vagrant/src/github.com/elwinar/elwinar"
	
	config.vm.provision "shell", inline: <<EOS
		chown vagrant:vagrant -R /home/vagrant/src
		pacman -Syy --noconfirm nodejs npm tmux most
		npm install -g bower gulp
EOS
end
