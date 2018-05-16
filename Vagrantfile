# -*- mode: ruby -*-
# vi: set ft=ruby :
# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

  if Vagrant.has_plugin?("vagrant-proxyconf")
    config.proxy.http     = "http://www-proxy-sin.sg.oracle.com:80"
    config.proxy.https    = "http://www-proxy-sin.sg.oracle.com:80"
    config.proxy.no_proxy = "oracle.com,oraclecorp.com"
  end

  # For use with local box
  config.vm.box = "OracleCode-SG"
  config.vm.box_url = "file://E:\\VagrantBoxes\\oraclecode.box"  
  config.vm.network "forwarded_port", guest: 5000, host: 5000, "id": "golang-port"
  config.vm.network "forwarded_port", guest: 22, host: 2222, "id": "ssh"  
  
  config.vm.provision "git_config", type: "shell", run: "once" do |s|    
    s.inline =  "git config --global user.name darrelchia; git config --global user.email darrel.chia@oracle.com"
  end
  # For use with remote box. Box is retrieved from vagrantcloud
  # config.vm.box = "darrelchia/terraform-dev-env"
  # config.vm.box_version = "1.0.0"
  
  config.vm.synced_folder "C:\\Users\\dchia\\Desktop\\OracleCode\\demo", "/home/vagrant/workspace/", group: "vagrant", owner: "vagrant", mount_options: ["dmode=750", "fmode=750"]

  config.vm.provider "virtualbox" do |vb|
    # vb.gui = false
    vb.name = "OracleCodeSG"
    vb.memory = 4092
    vb.cpus = 1
  end
end