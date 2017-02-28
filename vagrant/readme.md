## Vagrant Commands

### Boxes
` cd ~/.vagrant.d/boxes `

```
vagrant box list

vagrant box add centos65-x86_64-20140116 https://github.com/2creatives/vagrant-centos/releases/download/v6.5.3/centos65-x86_64-20140116.box

vagrant init centos65-x86_64-20140116

vagrant status

vagrant up

vagrant status

vagrant ssh
```

### Bridged Networking
When enabled, virtualbox connects to one of your installed network cards and exchanges network packets directly, circumventing your host operating system's network stack.

## Vagrant networking
* Forward ports
* Private network
* Public network

  ### Forward ports
  ```
  config.vm.network "forwarded_port", guest: 80, host: 8080
  ```

  ### Public networks
  Allows general public access to the machine.
  eg: Visible to the host network
      SSH-ing from host machine


  Setting DHCP
  ```
  config.vm.network "public_network"
  ```
  To set default network interface. If the default network interface is
  available it will ask to pick from a list of available network
  interfaces.
  ```
  config.vm.network "public_network", bridge: "en0: Wi-Fi (Airport)"
  ```

  Setting Static IP
  ```
  config.vm.network "public_network", ip: "192.168.0.17"
  ```

  ### Private networks

  Setting DHCP
  ```
  config.vm.network "private_network", type: "dhcp"
  ```

  Static IP
  ```
  config.vm.network "private_network", ip: "192.168.50.4"
  ```

  To disable-auto config
  ```
  config.vm.network "private_network", ip: "192.168.50.4", auto_config:
  false
  ```
