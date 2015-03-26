Vagrant.configure(2) do |config|
	config.vm.box = "elwinar/laravel"
	config.vm.hostname = "elwinar"
	
	config.vm.network "forwarded_port", 
			guest: 80, 
			host: 42000,
			auto_correct: true
	
	config.vm.network "private_network",
			type: "dhcp"
	
	config.vm.synced_folder '.', '/vagrant', 
			nfs: true,
			mount_options: ['actimeo=1']
	
	config.vm.provision "shell", inline: <<EOS
		pacman -S --noconfirm php-sqlite
		sed -ri 's/;extension=pdo_sqlite.so/extension=pdo_sqlite.so/' /etc/php/php.ini
		systemctl restart php-fpm

		rm -f /etc/nginx/nginx.conf
		ln -s /vagrant/nginx.conf /etc/nginx/nginx.conf
		systemctl restart nginx
EOS
end
