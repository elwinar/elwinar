Vagrant.configure(2) do |config|
	config.vm.box = "elwinar/laravel"
	config.vm.hostname = "elwinar"
	
	config.vm.network "forwarded_port", 
			guest: 80, 
			host: 8080,
			auto_correct: true
	
	config.vm.network "private_network",
			type: "dhcp"
	
	config.vm.synced_folder '.', '/vagrant', 
			nfs: true,
			mount_options: ['actimeo=1']
end
