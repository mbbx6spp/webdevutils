# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.ssh.default.username = "builder"

  config.vm.define :ubuntu do |ubuntu|
    ubuntu.vm.box = "builder-ubuntu-12.04"
    # TODO: Upload image somewhere to host
    #ubuntu.vm.box_url = ""
    ubuntu.vm.network :private_network, ip: "192.168.33.10"
    ubuntu.vm.synced_folder ".", "/mnt"
    ubuntu.vm.provision :shell, :inline => 'sudo -i /mnt/configure.ubuntu.sh'
    config.vm.network :forwarded_port, guest: 22, host: 2201
    config.vm.network :forwarded_port, guest: 3000, host: 3301
  end

  config.vm.define :centos do |centos|
    centos.vm.box = "builder-centos-6.4"
    # TODO: Upload image somewhere to host
    #centos.vm.box_url = ""
    centos.vm.network :private_network, ip: "192.168.33.11"
    centos.vm.synced_folder ".", "/mnt"
    centos.vm.provision :shell, :inline => 'sudo -i /mnt/configure.centos.sh'
    config.vm.network :forwarded_port, guest: 22, host: 2202
    config.vm.network :forwarded_port, guest: 3000, host: 3302
  end

  config.vm.define :freebsd do |freebsd|
    freebsd.vm.box = "builder-freebsd-9.1"
    # TODO: Upload image somewhere to host
    #freebsd.vm.box_url = ""
    freebsd.vm.network :private_network, ip: "192.168.33.13"
    freebsd.vm.synced_folder ".", "/mnt"
    freebsd.vm.provision :shell, :inline => 'sudo -i /mnt/configure.freebsd.sh'
    config.vm.network :forwarded_port, guest: 22, host: 2203
    config.vm.network :forwarded_port, guest: 3000, host: 3303
  end
end
