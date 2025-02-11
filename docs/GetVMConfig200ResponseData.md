# GetVMConfig200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acpi** | Pointer to **int32** | Enable/disable ACPI. | [optional] 
**Affinity** | Pointer to **string** | List of host cores used to execute guest processes, for example: 0,5,8-11 | [optional] 
**Agent** | Pointer to **string** | Enable/disable communication with the QEMU Guest Agent and its properties. | [optional] 
**AmdSev** | Pointer to **string** | Secure Encrypted Virtualization (SEV) features by AMD CPUs | [optional] 
**Arch** | Pointer to **string** | Virtual processor architecture. Defaults to the host. | [optional] 
**Args** | Pointer to **string** | Arbitrary arguments passed to kvm. | [optional] 
**Audio0** | Pointer to **string** | Configure a audio device, useful in combination with QXL/Spice. | [optional] 
**Autostart** | Pointer to **int32** | Automatic restart after crash (currently ignored). | [optional] 
**Balloon** | Pointer to **int64** | Amount of target RAM for the VM in MiB. Using zero disables the ballon driver. | [optional] 
**Bios** | Pointer to **string** | Select BIOS implementation. | [optional] 
**Boot** | Pointer to **string** | Specify guest boot order. Use the &#39;order&#x3D;&#39; sub-property as usage with no key or &#39;legacy&#x3D;&#39; is deprecated. | [optional] 
**Bootdisk** | Pointer to **string** | Enable booting from specified disk. Deprecated: Use &#39;boot: order&#x3D;foo;bar&#39; instead. | [optional] 
**Cdrom** | Pointer to **string** | This is an alias for option -ide2 | [optional] 
**Cicustom** | Pointer to **string** | cloud-init: Specify custom files to replace the automatically generated ones at start. | [optional] 
**Cipassword** | Pointer to **string** | cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords. | [optional] 
**Citype** | Pointer to **string** | Specifies the cloud-init configuration format. The default depends on the configured operating system type (&#x60;ostype&#x60;. We use the &#x60;nocloud&#x60; format for Linux, and &#x60;configdrive2&#x60; for windows. | [optional] 
**Ciupgrade** | Pointer to **int32** | cloud-init: do an automatic package upgrade after the first boot. | [optional] 
**Ciuser** | Pointer to **string** | cloud-init: User name to change ssh keys and password for instead of the image&#39;s configured default user. | [optional] 
**Cores** | Pointer to **int64** | The number of cores per socket. | [optional] 
**Cpu** | Pointer to **string** | Emulated CPU type. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Description** | Pointer to **string** | Description for the VM. Shown in the web-interface VM&#39;s summary. This is saved as comment inside the configuration file. | [optional] 
**Digest** | Pointer to **string** | SHA1 digest of configuration file. This can be used to prevent concurrent modifications. | [optional] 
**Efidisk0** | Pointer to **string** | Configure a disk for storing EFI vars. | [optional] 
**Freeze** | Pointer to **int32** | Freeze CPU at startup (use &#39;c&#39; monitor command to start execution). | [optional] 
**Hookscript** | Pointer to **string** | Script that will be executed during various steps in the vms lifetime. | [optional] 
**Hostpci0** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci1** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci2** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci3** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci4** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci5** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci6** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci7** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci8** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci9** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci10** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci11** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci12** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci13** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci14** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci15** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci16** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci17** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci18** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci19** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci20** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci21** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci22** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci23** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci24** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci25** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci26** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci27** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci28** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci29** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hotplug** | Pointer to **string** | Selectively enable hotplug features. This is a comma separated list of hotplug features: &#39;network&#39;, &#39;disk&#39;, &#39;cpu&#39;, &#39;memory&#39;, &#39;usb&#39; and &#39;cloudinit&#39;. Use &#39;0&#39; to disable hotplug completely. Using &#39;1&#39; as value is an alias for the default &#x60;network,disk,usb&#x60;. USB hotplugging is possible for guests with machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7. | [optional] 
**Hugepages** | Pointer to **string** | Enable/disable hugepages memory. | [optional] 
**Ide0** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). | [optional] 
**Ide1** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). | [optional] 
**Ide2** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). | [optional] 
**Ide3** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). | [optional] 
**Ipconfig0** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig1** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig2** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig3** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig4** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig5** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig6** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig7** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig8** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig9** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig10** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig11** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig12** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig13** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig14** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig15** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig16** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig17** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig18** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig19** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig20** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig21** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig22** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig23** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig24** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig25** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig26** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig27** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig28** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig29** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ivshmem** | Pointer to **string** | Inter-VM shared memory. Useful for direct communication between VMs, or to the host. | [optional] 
**Keephugepages** | Pointer to **int32** | Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts. | [optional] 
**Keyboard** | Pointer to **string** | Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS. | [optional] 
**Kvm** | Pointer to **int32** | Enable/disable KVM hardware virtualization. | [optional] 
**Localtime** | Pointer to **int32** | Set the real time clock (RTC) to local time. This is enabled by default if the &#x60;ostype&#x60; indicates a Microsoft Windows OS. | [optional] 
**Lock** | Pointer to **string** | Lock/unlock the VM. | [optional] 
**Machine** | Pointer to **string** | Specify the QEMU machine. | [optional] 
**Memory** | Pointer to **int64** | Memory properties. | [optional] 
**MigrateDowntime** | Pointer to **float32** | Set maximum tolerated downtime (in seconds) for migrations. Should the migration not be able to converge in the very end, because too much newly dirtied RAM needs to be transferred, the limit will be increased automatically step-by-step until migration can converge. | [optional] 
**MigrateSpeed** | Pointer to **int64** | Set maximum speed (in MB/s) for migrations. Value 0 is no limit. | [optional] 
**Name** | Pointer to **string** | Set a name for the VM. Only used on the configuration web interface. | [optional] 
**Nameserver** | Pointer to **string** | cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Net0** | Pointer to **string** | Specify network devices. | [optional] 
**Net1** | Pointer to **string** | Specify network devices. | [optional] 
**Net2** | Pointer to **string** | Specify network devices. | [optional] 
**Net3** | Pointer to **string** | Specify network devices. | [optional] 
**Net4** | Pointer to **string** | Specify network devices. | [optional] 
**Net5** | Pointer to **string** | Specify network devices. | [optional] 
**Net6** | Pointer to **string** | Specify network devices. | [optional] 
**Net7** | Pointer to **string** | Specify network devices. | [optional] 
**Net8** | Pointer to **string** | Specify network devices. | [optional] 
**Net9** | Pointer to **string** | Specify network devices. | [optional] 
**Net10** | Pointer to **string** | Specify network devices. | [optional] 
**Net11** | Pointer to **string** | Specify network devices. | [optional] 
**Net12** | Pointer to **string** | Specify network devices. | [optional] 
**Net13** | Pointer to **string** | Specify network devices. | [optional] 
**Net14** | Pointer to **string** | Specify network devices. | [optional] 
**Net15** | Pointer to **string** | Specify network devices. | [optional] 
**Net16** | Pointer to **string** | Specify network devices. | [optional] 
**Net17** | Pointer to **string** | Specify network devices. | [optional] 
**Net18** | Pointer to **string** | Specify network devices. | [optional] 
**Net19** | Pointer to **string** | Specify network devices. | [optional] 
**Net20** | Pointer to **string** | Specify network devices. | [optional] 
**Net21** | Pointer to **string** | Specify network devices. | [optional] 
**Net22** | Pointer to **string** | Specify network devices. | [optional] 
**Net23** | Pointer to **string** | Specify network devices. | [optional] 
**Net24** | Pointer to **string** | Specify network devices. | [optional] 
**Net25** | Pointer to **string** | Specify network devices. | [optional] 
**Net26** | Pointer to **string** | Specify network devices. | [optional] 
**Net27** | Pointer to **string** | Specify network devices. | [optional] 
**Net28** | Pointer to **string** | Specify network devices. | [optional] 
**Net29** | Pointer to **string** | Specify network devices. | [optional] 
**Net30** | Pointer to **string** | Specify network devices. | [optional] 
**Net31** | Pointer to **string** | Specify network devices. | [optional] 
**Numa** | Pointer to **int32** | Enable/disable NUMA. | [optional] 
**Numa0** | Pointer to **string** | NUMA topology. | [optional] 
**Numa1** | Pointer to **string** | NUMA topology. | [optional] 
**Numa2** | Pointer to **string** | NUMA topology. | [optional] 
**Numa3** | Pointer to **string** | NUMA topology. | [optional] 
**Numa4** | Pointer to **string** | NUMA topology. | [optional] 
**Numa5** | Pointer to **string** | NUMA topology. | [optional] 
**Numa6** | Pointer to **string** | NUMA topology. | [optional] 
**Numa7** | Pointer to **string** | NUMA topology. | [optional] 
**Numa8** | Pointer to **string** | NUMA topology. | [optional] 
**Numa9** | Pointer to **string** | NUMA topology. | [optional] 
**Numa10** | Pointer to **string** | NUMA topology. | [optional] 
**Numa11** | Pointer to **string** | NUMA topology. | [optional] 
**Numa12** | Pointer to **string** | NUMA topology. | [optional] 
**Numa13** | Pointer to **string** | NUMA topology. | [optional] 
**Numa14** | Pointer to **string** | NUMA topology. | [optional] 
**Numa15** | Pointer to **string** | NUMA topology. | [optional] 
**Numa16** | Pointer to **string** | NUMA topology. | [optional] 
**Numa17** | Pointer to **string** | NUMA topology. | [optional] 
**Numa18** | Pointer to **string** | NUMA topology. | [optional] 
**Numa19** | Pointer to **string** | NUMA topology. | [optional] 
**Numa20** | Pointer to **string** | NUMA topology. | [optional] 
**Numa21** | Pointer to **string** | NUMA topology. | [optional] 
**Numa22** | Pointer to **string** | NUMA topology. | [optional] 
**Numa23** | Pointer to **string** | NUMA topology. | [optional] 
**Numa24** | Pointer to **string** | NUMA topology. | [optional] 
**Numa25** | Pointer to **string** | NUMA topology. | [optional] 
**Numa26** | Pointer to **string** | NUMA topology. | [optional] 
**Numa27** | Pointer to **string** | NUMA topology. | [optional] 
**Numa28** | Pointer to **string** | NUMA topology. | [optional] 
**Numa29** | Pointer to **string** | NUMA topology. | [optional] 
**Onboot** | Pointer to **int32** | Specifies whether a VM will be started during system bootup. | [optional] 
**Ostype** | Pointer to **string** | Specify guest operating system. | [optional] 
**Parallel0** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel1** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel2** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel3** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Protection** | Pointer to **int32** | Sets the protection flag of the VM. This will disable the remove VM and remove disk operations. | [optional] 
**Reboot** | Pointer to **int32** | Allow reboot. If set to &#39;0&#39; the VM exit on reboot. | [optional] 
**Rng0** | Pointer to **string** | Configure a VirtIO-based Random Number Generator. | [optional] 
**Sata0** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Sata1** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Sata2** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Sata3** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Sata4** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Sata5** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). | [optional] 
**Scsi0** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi1** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi2** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi3** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi4** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi5** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi6** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi7** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi8** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi9** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi10** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi11** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi12** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi13** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi14** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi15** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi16** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi17** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi18** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi19** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi20** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi21** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi22** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi23** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi24** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi25** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi26** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi27** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi28** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsi29** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). | [optional] 
**Scsihw** | Pointer to **string** | SCSI controller model | [optional] 
**Searchdomain** | Pointer to **string** | cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Serial0** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial1** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial2** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial3** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Shares** | Pointer to **int64** | Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd. | [optional] 
**Smbios1** | Pointer to **string** | Specify SMBIOS type 1 fields. | [optional] 
**Smp** | Pointer to **int64** | The number of CPUs. Please use option -sockets instead. | [optional] 
**Sockets** | Pointer to **int64** | The number of CPU sockets. | [optional] 
**SpiceEnhancements** | Pointer to **string** | Configure additional enhancements for SPICE. | [optional] 
**Sshkeys** | Pointer to **string** | cloud-init: Setup public SSH keys (one key per line, OpenSSH format). | [optional] 
**Startdate** | Pointer to **string** | Set the initial date of the real time clock. Valid format for date are:&#39;now&#39; or &#39;2006-06-17T16:01:21&#39; or &#39;2006-06-17&#39;. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Tablet** | Pointer to **int32** | Enable/disable the USB tablet device. | [optional] 
**Tags** | Pointer to **string** | Tags of the VM. This is only meta information. | [optional] 
**Tdf** | Pointer to **int32** | Enable/disable time drift fix. | [optional] 
**Template** | Pointer to **int32** | Enable/disable Template. | [optional] 
**Tpmstate0** | Pointer to **string** | Configure a Disk for storing TPM state. The format is fixed to &#39;raw&#39;. | [optional] 
**Unused0** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused1** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused2** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused3** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused4** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused5** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused6** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused7** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused8** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused9** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused10** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused11** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused12** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused13** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused14** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused15** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused16** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused17** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused18** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused19** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused20** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused21** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused22** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused23** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused24** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused25** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused26** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused27** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused28** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused29** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Usb0** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb1** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb2** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb3** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Vcpus** | Pointer to **int64** | Number of hotplugged vcpus. | [optional] 
**Vga** | Pointer to **string** | Configure the VGA hardware. | [optional] 
**Virtio0** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio1** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio2** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio3** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio4** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio5** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio6** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio7** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio8** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio9** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio10** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio11** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio12** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio13** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio14** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Virtio15** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). | [optional] 
**Vmgenid** | Pointer to **string** | Set VM Generation ID. Use &#39;1&#39; to autogenerate on create or update, pass &#39;0&#39; to disable explicitly. | [optional] 
**Vmstatestorage** | Pointer to **string** | Default storage for VM state volumes/files. | [optional] 
**Watchdog** | Pointer to **string** | Create a virtual hardware watchdog device. | [optional] 

## Methods

### NewGetVMConfig200ResponseData

`func NewGetVMConfig200ResponseData() *GetVMConfig200ResponseData`

NewGetVMConfig200ResponseData instantiates a new GetVMConfig200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataWithDefaults

`func NewGetVMConfig200ResponseDataWithDefaults() *GetVMConfig200ResponseData`

NewGetVMConfig200ResponseDataWithDefaults instantiates a new GetVMConfig200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *GetVMConfig200ResponseData) GetAcpi() int32`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *GetVMConfig200ResponseData) GetAcpiOk() (*int32, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *GetVMConfig200ResponseData) SetAcpi(v int32)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *GetVMConfig200ResponseData) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetAffinity

`func (o *GetVMConfig200ResponseData) GetAffinity() string`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *GetVMConfig200ResponseData) GetAffinityOk() (*string, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *GetVMConfig200ResponseData) SetAffinity(v string)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *GetVMConfig200ResponseData) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAgent

`func (o *GetVMConfig200ResponseData) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *GetVMConfig200ResponseData) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *GetVMConfig200ResponseData) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *GetVMConfig200ResponseData) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAmdSev

`func (o *GetVMConfig200ResponseData) GetAmdSev() string`

GetAmdSev returns the AmdSev field if non-nil, zero value otherwise.

### GetAmdSevOk

`func (o *GetVMConfig200ResponseData) GetAmdSevOk() (*string, bool)`

GetAmdSevOk returns a tuple with the AmdSev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdSev

`func (o *GetVMConfig200ResponseData) SetAmdSev(v string)`

SetAmdSev sets AmdSev field to given value.

### HasAmdSev

`func (o *GetVMConfig200ResponseData) HasAmdSev() bool`

HasAmdSev returns a boolean if a field has been set.

### GetArch

`func (o *GetVMConfig200ResponseData) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *GetVMConfig200ResponseData) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *GetVMConfig200ResponseData) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *GetVMConfig200ResponseData) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *GetVMConfig200ResponseData) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GetVMConfig200ResponseData) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GetVMConfig200ResponseData) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GetVMConfig200ResponseData) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAudio0

`func (o *GetVMConfig200ResponseData) GetAudio0() string`

GetAudio0 returns the Audio0 field if non-nil, zero value otherwise.

### GetAudio0Ok

`func (o *GetVMConfig200ResponseData) GetAudio0Ok() (*string, bool)`

GetAudio0Ok returns a tuple with the Audio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio0

`func (o *GetVMConfig200ResponseData) SetAudio0(v string)`

SetAudio0 sets Audio0 field to given value.

### HasAudio0

`func (o *GetVMConfig200ResponseData) HasAudio0() bool`

HasAudio0 returns a boolean if a field has been set.

### GetAutostart

`func (o *GetVMConfig200ResponseData) GetAutostart() int32`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *GetVMConfig200ResponseData) GetAutostartOk() (*int32, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *GetVMConfig200ResponseData) SetAutostart(v int32)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *GetVMConfig200ResponseData) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBalloon

`func (o *GetVMConfig200ResponseData) GetBalloon() int64`

GetBalloon returns the Balloon field if non-nil, zero value otherwise.

### GetBalloonOk

`func (o *GetVMConfig200ResponseData) GetBalloonOk() (*int64, bool)`

GetBalloonOk returns a tuple with the Balloon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloon

`func (o *GetVMConfig200ResponseData) SetBalloon(v int64)`

SetBalloon sets Balloon field to given value.

### HasBalloon

`func (o *GetVMConfig200ResponseData) HasBalloon() bool`

HasBalloon returns a boolean if a field has been set.

### GetBios

`func (o *GetVMConfig200ResponseData) GetBios() string`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *GetVMConfig200ResponseData) GetBiosOk() (*string, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *GetVMConfig200ResponseData) SetBios(v string)`

SetBios sets Bios field to given value.

### HasBios

`func (o *GetVMConfig200ResponseData) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetBoot

`func (o *GetVMConfig200ResponseData) GetBoot() string`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *GetVMConfig200ResponseData) GetBootOk() (*string, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *GetVMConfig200ResponseData) SetBoot(v string)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *GetVMConfig200ResponseData) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBootdisk

`func (o *GetVMConfig200ResponseData) GetBootdisk() string`

GetBootdisk returns the Bootdisk field if non-nil, zero value otherwise.

### GetBootdiskOk

`func (o *GetVMConfig200ResponseData) GetBootdiskOk() (*string, bool)`

GetBootdiskOk returns a tuple with the Bootdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootdisk

`func (o *GetVMConfig200ResponseData) SetBootdisk(v string)`

SetBootdisk sets Bootdisk field to given value.

### HasBootdisk

`func (o *GetVMConfig200ResponseData) HasBootdisk() bool`

HasBootdisk returns a boolean if a field has been set.

### GetCdrom

`func (o *GetVMConfig200ResponseData) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *GetVMConfig200ResponseData) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *GetVMConfig200ResponseData) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *GetVMConfig200ResponseData) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetCicustom

`func (o *GetVMConfig200ResponseData) GetCicustom() string`

GetCicustom returns the Cicustom field if non-nil, zero value otherwise.

### GetCicustomOk

`func (o *GetVMConfig200ResponseData) GetCicustomOk() (*string, bool)`

GetCicustomOk returns a tuple with the Cicustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicustom

`func (o *GetVMConfig200ResponseData) SetCicustom(v string)`

SetCicustom sets Cicustom field to given value.

### HasCicustom

`func (o *GetVMConfig200ResponseData) HasCicustom() bool`

HasCicustom returns a boolean if a field has been set.

### GetCipassword

`func (o *GetVMConfig200ResponseData) GetCipassword() string`

GetCipassword returns the Cipassword field if non-nil, zero value otherwise.

### GetCipasswordOk

`func (o *GetVMConfig200ResponseData) GetCipasswordOk() (*string, bool)`

GetCipasswordOk returns a tuple with the Cipassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipassword

`func (o *GetVMConfig200ResponseData) SetCipassword(v string)`

SetCipassword sets Cipassword field to given value.

### HasCipassword

`func (o *GetVMConfig200ResponseData) HasCipassword() bool`

HasCipassword returns a boolean if a field has been set.

### GetCitype

`func (o *GetVMConfig200ResponseData) GetCitype() string`

GetCitype returns the Citype field if non-nil, zero value otherwise.

### GetCitypeOk

`func (o *GetVMConfig200ResponseData) GetCitypeOk() (*string, bool)`

GetCitypeOk returns a tuple with the Citype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitype

`func (o *GetVMConfig200ResponseData) SetCitype(v string)`

SetCitype sets Citype field to given value.

### HasCitype

`func (o *GetVMConfig200ResponseData) HasCitype() bool`

HasCitype returns a boolean if a field has been set.

### GetCiupgrade

`func (o *GetVMConfig200ResponseData) GetCiupgrade() int32`

GetCiupgrade returns the Ciupgrade field if non-nil, zero value otherwise.

### GetCiupgradeOk

`func (o *GetVMConfig200ResponseData) GetCiupgradeOk() (*int32, bool)`

GetCiupgradeOk returns a tuple with the Ciupgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiupgrade

`func (o *GetVMConfig200ResponseData) SetCiupgrade(v int32)`

SetCiupgrade sets Ciupgrade field to given value.

### HasCiupgrade

`func (o *GetVMConfig200ResponseData) HasCiupgrade() bool`

HasCiupgrade returns a boolean if a field has been set.

### GetCiuser

`func (o *GetVMConfig200ResponseData) GetCiuser() string`

GetCiuser returns the Ciuser field if non-nil, zero value otherwise.

### GetCiuserOk

`func (o *GetVMConfig200ResponseData) GetCiuserOk() (*string, bool)`

GetCiuserOk returns a tuple with the Ciuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiuser

`func (o *GetVMConfig200ResponseData) SetCiuser(v string)`

SetCiuser sets Ciuser field to given value.

### HasCiuser

`func (o *GetVMConfig200ResponseData) HasCiuser() bool`

HasCiuser returns a boolean if a field has been set.

### GetCores

`func (o *GetVMConfig200ResponseData) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *GetVMConfig200ResponseData) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *GetVMConfig200ResponseData) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *GetVMConfig200ResponseData) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpu

`func (o *GetVMConfig200ResponseData) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetVMConfig200ResponseData) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetVMConfig200ResponseData) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GetVMConfig200ResponseData) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpulimit

`func (o *GetVMConfig200ResponseData) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *GetVMConfig200ResponseData) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *GetVMConfig200ResponseData) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *GetVMConfig200ResponseData) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *GetVMConfig200ResponseData) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *GetVMConfig200ResponseData) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *GetVMConfig200ResponseData) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *GetVMConfig200ResponseData) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDescription

`func (o *GetVMConfig200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVMConfig200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVMConfig200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVMConfig200ResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *GetVMConfig200ResponseData) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GetVMConfig200ResponseData) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GetVMConfig200ResponseData) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *GetVMConfig200ResponseData) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEfidisk0

`func (o *GetVMConfig200ResponseData) GetEfidisk0() string`

GetEfidisk0 returns the Efidisk0 field if non-nil, zero value otherwise.

### GetEfidisk0Ok

`func (o *GetVMConfig200ResponseData) GetEfidisk0Ok() (*string, bool)`

GetEfidisk0Ok returns a tuple with the Efidisk0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfidisk0

`func (o *GetVMConfig200ResponseData) SetEfidisk0(v string)`

SetEfidisk0 sets Efidisk0 field to given value.

### HasEfidisk0

`func (o *GetVMConfig200ResponseData) HasEfidisk0() bool`

HasEfidisk0 returns a boolean if a field has been set.

### GetFreeze

`func (o *GetVMConfig200ResponseData) GetFreeze() int32`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *GetVMConfig200ResponseData) GetFreezeOk() (*int32, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *GetVMConfig200ResponseData) SetFreeze(v int32)`

SetFreeze sets Freeze field to given value.

### HasFreeze

`func (o *GetVMConfig200ResponseData) HasFreeze() bool`

HasFreeze returns a boolean if a field has been set.

### GetHookscript

`func (o *GetVMConfig200ResponseData) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *GetVMConfig200ResponseData) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *GetVMConfig200ResponseData) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *GetVMConfig200ResponseData) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostpci0

`func (o *GetVMConfig200ResponseData) GetHostpci0() string`

GetHostpci0 returns the Hostpci0 field if non-nil, zero value otherwise.

### GetHostpci0Ok

`func (o *GetVMConfig200ResponseData) GetHostpci0Ok() (*string, bool)`

GetHostpci0Ok returns a tuple with the Hostpci0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci0

`func (o *GetVMConfig200ResponseData) SetHostpci0(v string)`

SetHostpci0 sets Hostpci0 field to given value.

### HasHostpci0

`func (o *GetVMConfig200ResponseData) HasHostpci0() bool`

HasHostpci0 returns a boolean if a field has been set.

### GetHostpci1

`func (o *GetVMConfig200ResponseData) GetHostpci1() string`

GetHostpci1 returns the Hostpci1 field if non-nil, zero value otherwise.

### GetHostpci1Ok

`func (o *GetVMConfig200ResponseData) GetHostpci1Ok() (*string, bool)`

GetHostpci1Ok returns a tuple with the Hostpci1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci1

`func (o *GetVMConfig200ResponseData) SetHostpci1(v string)`

SetHostpci1 sets Hostpci1 field to given value.

### HasHostpci1

`func (o *GetVMConfig200ResponseData) HasHostpci1() bool`

HasHostpci1 returns a boolean if a field has been set.

### GetHostpci2

`func (o *GetVMConfig200ResponseData) GetHostpci2() string`

GetHostpci2 returns the Hostpci2 field if non-nil, zero value otherwise.

### GetHostpci2Ok

`func (o *GetVMConfig200ResponseData) GetHostpci2Ok() (*string, bool)`

GetHostpci2Ok returns a tuple with the Hostpci2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci2

`func (o *GetVMConfig200ResponseData) SetHostpci2(v string)`

SetHostpci2 sets Hostpci2 field to given value.

### HasHostpci2

`func (o *GetVMConfig200ResponseData) HasHostpci2() bool`

HasHostpci2 returns a boolean if a field has been set.

### GetHostpci3

`func (o *GetVMConfig200ResponseData) GetHostpci3() string`

GetHostpci3 returns the Hostpci3 field if non-nil, zero value otherwise.

### GetHostpci3Ok

`func (o *GetVMConfig200ResponseData) GetHostpci3Ok() (*string, bool)`

GetHostpci3Ok returns a tuple with the Hostpci3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci3

`func (o *GetVMConfig200ResponseData) SetHostpci3(v string)`

SetHostpci3 sets Hostpci3 field to given value.

### HasHostpci3

`func (o *GetVMConfig200ResponseData) HasHostpci3() bool`

HasHostpci3 returns a boolean if a field has been set.

### GetHostpci4

`func (o *GetVMConfig200ResponseData) GetHostpci4() string`

GetHostpci4 returns the Hostpci4 field if non-nil, zero value otherwise.

### GetHostpci4Ok

`func (o *GetVMConfig200ResponseData) GetHostpci4Ok() (*string, bool)`

GetHostpci4Ok returns a tuple with the Hostpci4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci4

`func (o *GetVMConfig200ResponseData) SetHostpci4(v string)`

SetHostpci4 sets Hostpci4 field to given value.

### HasHostpci4

`func (o *GetVMConfig200ResponseData) HasHostpci4() bool`

HasHostpci4 returns a boolean if a field has been set.

### GetHostpci5

`func (o *GetVMConfig200ResponseData) GetHostpci5() string`

GetHostpci5 returns the Hostpci5 field if non-nil, zero value otherwise.

### GetHostpci5Ok

`func (o *GetVMConfig200ResponseData) GetHostpci5Ok() (*string, bool)`

GetHostpci5Ok returns a tuple with the Hostpci5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci5

`func (o *GetVMConfig200ResponseData) SetHostpci5(v string)`

SetHostpci5 sets Hostpci5 field to given value.

### HasHostpci5

`func (o *GetVMConfig200ResponseData) HasHostpci5() bool`

HasHostpci5 returns a boolean if a field has been set.

### GetHostpci6

`func (o *GetVMConfig200ResponseData) GetHostpci6() string`

GetHostpci6 returns the Hostpci6 field if non-nil, zero value otherwise.

### GetHostpci6Ok

`func (o *GetVMConfig200ResponseData) GetHostpci6Ok() (*string, bool)`

GetHostpci6Ok returns a tuple with the Hostpci6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci6

`func (o *GetVMConfig200ResponseData) SetHostpci6(v string)`

SetHostpci6 sets Hostpci6 field to given value.

### HasHostpci6

`func (o *GetVMConfig200ResponseData) HasHostpci6() bool`

HasHostpci6 returns a boolean if a field has been set.

### GetHostpci7

`func (o *GetVMConfig200ResponseData) GetHostpci7() string`

GetHostpci7 returns the Hostpci7 field if non-nil, zero value otherwise.

### GetHostpci7Ok

`func (o *GetVMConfig200ResponseData) GetHostpci7Ok() (*string, bool)`

GetHostpci7Ok returns a tuple with the Hostpci7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci7

`func (o *GetVMConfig200ResponseData) SetHostpci7(v string)`

SetHostpci7 sets Hostpci7 field to given value.

### HasHostpci7

`func (o *GetVMConfig200ResponseData) HasHostpci7() bool`

HasHostpci7 returns a boolean if a field has been set.

### GetHostpci8

`func (o *GetVMConfig200ResponseData) GetHostpci8() string`

GetHostpci8 returns the Hostpci8 field if non-nil, zero value otherwise.

### GetHostpci8Ok

`func (o *GetVMConfig200ResponseData) GetHostpci8Ok() (*string, bool)`

GetHostpci8Ok returns a tuple with the Hostpci8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci8

`func (o *GetVMConfig200ResponseData) SetHostpci8(v string)`

SetHostpci8 sets Hostpci8 field to given value.

### HasHostpci8

`func (o *GetVMConfig200ResponseData) HasHostpci8() bool`

HasHostpci8 returns a boolean if a field has been set.

### GetHostpci9

`func (o *GetVMConfig200ResponseData) GetHostpci9() string`

GetHostpci9 returns the Hostpci9 field if non-nil, zero value otherwise.

### GetHostpci9Ok

`func (o *GetVMConfig200ResponseData) GetHostpci9Ok() (*string, bool)`

GetHostpci9Ok returns a tuple with the Hostpci9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci9

`func (o *GetVMConfig200ResponseData) SetHostpci9(v string)`

SetHostpci9 sets Hostpci9 field to given value.

### HasHostpci9

`func (o *GetVMConfig200ResponseData) HasHostpci9() bool`

HasHostpci9 returns a boolean if a field has been set.

### GetHostpci10

`func (o *GetVMConfig200ResponseData) GetHostpci10() string`

GetHostpci10 returns the Hostpci10 field if non-nil, zero value otherwise.

### GetHostpci10Ok

`func (o *GetVMConfig200ResponseData) GetHostpci10Ok() (*string, bool)`

GetHostpci10Ok returns a tuple with the Hostpci10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci10

`func (o *GetVMConfig200ResponseData) SetHostpci10(v string)`

SetHostpci10 sets Hostpci10 field to given value.

### HasHostpci10

`func (o *GetVMConfig200ResponseData) HasHostpci10() bool`

HasHostpci10 returns a boolean if a field has been set.

### GetHostpci11

`func (o *GetVMConfig200ResponseData) GetHostpci11() string`

GetHostpci11 returns the Hostpci11 field if non-nil, zero value otherwise.

### GetHostpci11Ok

`func (o *GetVMConfig200ResponseData) GetHostpci11Ok() (*string, bool)`

GetHostpci11Ok returns a tuple with the Hostpci11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci11

`func (o *GetVMConfig200ResponseData) SetHostpci11(v string)`

SetHostpci11 sets Hostpci11 field to given value.

### HasHostpci11

`func (o *GetVMConfig200ResponseData) HasHostpci11() bool`

HasHostpci11 returns a boolean if a field has been set.

### GetHostpci12

`func (o *GetVMConfig200ResponseData) GetHostpci12() string`

GetHostpci12 returns the Hostpci12 field if non-nil, zero value otherwise.

### GetHostpci12Ok

`func (o *GetVMConfig200ResponseData) GetHostpci12Ok() (*string, bool)`

GetHostpci12Ok returns a tuple with the Hostpci12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci12

`func (o *GetVMConfig200ResponseData) SetHostpci12(v string)`

SetHostpci12 sets Hostpci12 field to given value.

### HasHostpci12

`func (o *GetVMConfig200ResponseData) HasHostpci12() bool`

HasHostpci12 returns a boolean if a field has been set.

### GetHostpci13

`func (o *GetVMConfig200ResponseData) GetHostpci13() string`

GetHostpci13 returns the Hostpci13 field if non-nil, zero value otherwise.

### GetHostpci13Ok

`func (o *GetVMConfig200ResponseData) GetHostpci13Ok() (*string, bool)`

GetHostpci13Ok returns a tuple with the Hostpci13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci13

`func (o *GetVMConfig200ResponseData) SetHostpci13(v string)`

SetHostpci13 sets Hostpci13 field to given value.

### HasHostpci13

`func (o *GetVMConfig200ResponseData) HasHostpci13() bool`

HasHostpci13 returns a boolean if a field has been set.

### GetHostpci14

`func (o *GetVMConfig200ResponseData) GetHostpci14() string`

GetHostpci14 returns the Hostpci14 field if non-nil, zero value otherwise.

### GetHostpci14Ok

`func (o *GetVMConfig200ResponseData) GetHostpci14Ok() (*string, bool)`

GetHostpci14Ok returns a tuple with the Hostpci14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci14

`func (o *GetVMConfig200ResponseData) SetHostpci14(v string)`

SetHostpci14 sets Hostpci14 field to given value.

### HasHostpci14

`func (o *GetVMConfig200ResponseData) HasHostpci14() bool`

HasHostpci14 returns a boolean if a field has been set.

### GetHostpci15

`func (o *GetVMConfig200ResponseData) GetHostpci15() string`

GetHostpci15 returns the Hostpci15 field if non-nil, zero value otherwise.

### GetHostpci15Ok

`func (o *GetVMConfig200ResponseData) GetHostpci15Ok() (*string, bool)`

GetHostpci15Ok returns a tuple with the Hostpci15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci15

`func (o *GetVMConfig200ResponseData) SetHostpci15(v string)`

SetHostpci15 sets Hostpci15 field to given value.

### HasHostpci15

`func (o *GetVMConfig200ResponseData) HasHostpci15() bool`

HasHostpci15 returns a boolean if a field has been set.

### GetHostpci16

`func (o *GetVMConfig200ResponseData) GetHostpci16() string`

GetHostpci16 returns the Hostpci16 field if non-nil, zero value otherwise.

### GetHostpci16Ok

`func (o *GetVMConfig200ResponseData) GetHostpci16Ok() (*string, bool)`

GetHostpci16Ok returns a tuple with the Hostpci16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci16

`func (o *GetVMConfig200ResponseData) SetHostpci16(v string)`

SetHostpci16 sets Hostpci16 field to given value.

### HasHostpci16

`func (o *GetVMConfig200ResponseData) HasHostpci16() bool`

HasHostpci16 returns a boolean if a field has been set.

### GetHostpci17

`func (o *GetVMConfig200ResponseData) GetHostpci17() string`

GetHostpci17 returns the Hostpci17 field if non-nil, zero value otherwise.

### GetHostpci17Ok

`func (o *GetVMConfig200ResponseData) GetHostpci17Ok() (*string, bool)`

GetHostpci17Ok returns a tuple with the Hostpci17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci17

`func (o *GetVMConfig200ResponseData) SetHostpci17(v string)`

SetHostpci17 sets Hostpci17 field to given value.

### HasHostpci17

`func (o *GetVMConfig200ResponseData) HasHostpci17() bool`

HasHostpci17 returns a boolean if a field has been set.

### GetHostpci18

`func (o *GetVMConfig200ResponseData) GetHostpci18() string`

GetHostpci18 returns the Hostpci18 field if non-nil, zero value otherwise.

### GetHostpci18Ok

`func (o *GetVMConfig200ResponseData) GetHostpci18Ok() (*string, bool)`

GetHostpci18Ok returns a tuple with the Hostpci18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci18

`func (o *GetVMConfig200ResponseData) SetHostpci18(v string)`

SetHostpci18 sets Hostpci18 field to given value.

### HasHostpci18

`func (o *GetVMConfig200ResponseData) HasHostpci18() bool`

HasHostpci18 returns a boolean if a field has been set.

### GetHostpci19

`func (o *GetVMConfig200ResponseData) GetHostpci19() string`

GetHostpci19 returns the Hostpci19 field if non-nil, zero value otherwise.

### GetHostpci19Ok

`func (o *GetVMConfig200ResponseData) GetHostpci19Ok() (*string, bool)`

GetHostpci19Ok returns a tuple with the Hostpci19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci19

`func (o *GetVMConfig200ResponseData) SetHostpci19(v string)`

SetHostpci19 sets Hostpci19 field to given value.

### HasHostpci19

`func (o *GetVMConfig200ResponseData) HasHostpci19() bool`

HasHostpci19 returns a boolean if a field has been set.

### GetHostpci20

`func (o *GetVMConfig200ResponseData) GetHostpci20() string`

GetHostpci20 returns the Hostpci20 field if non-nil, zero value otherwise.

### GetHostpci20Ok

`func (o *GetVMConfig200ResponseData) GetHostpci20Ok() (*string, bool)`

GetHostpci20Ok returns a tuple with the Hostpci20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci20

`func (o *GetVMConfig200ResponseData) SetHostpci20(v string)`

SetHostpci20 sets Hostpci20 field to given value.

### HasHostpci20

`func (o *GetVMConfig200ResponseData) HasHostpci20() bool`

HasHostpci20 returns a boolean if a field has been set.

### GetHostpci21

`func (o *GetVMConfig200ResponseData) GetHostpci21() string`

GetHostpci21 returns the Hostpci21 field if non-nil, zero value otherwise.

### GetHostpci21Ok

`func (o *GetVMConfig200ResponseData) GetHostpci21Ok() (*string, bool)`

GetHostpci21Ok returns a tuple with the Hostpci21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci21

`func (o *GetVMConfig200ResponseData) SetHostpci21(v string)`

SetHostpci21 sets Hostpci21 field to given value.

### HasHostpci21

`func (o *GetVMConfig200ResponseData) HasHostpci21() bool`

HasHostpci21 returns a boolean if a field has been set.

### GetHostpci22

`func (o *GetVMConfig200ResponseData) GetHostpci22() string`

GetHostpci22 returns the Hostpci22 field if non-nil, zero value otherwise.

### GetHostpci22Ok

`func (o *GetVMConfig200ResponseData) GetHostpci22Ok() (*string, bool)`

GetHostpci22Ok returns a tuple with the Hostpci22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci22

`func (o *GetVMConfig200ResponseData) SetHostpci22(v string)`

SetHostpci22 sets Hostpci22 field to given value.

### HasHostpci22

`func (o *GetVMConfig200ResponseData) HasHostpci22() bool`

HasHostpci22 returns a boolean if a field has been set.

### GetHostpci23

`func (o *GetVMConfig200ResponseData) GetHostpci23() string`

GetHostpci23 returns the Hostpci23 field if non-nil, zero value otherwise.

### GetHostpci23Ok

`func (o *GetVMConfig200ResponseData) GetHostpci23Ok() (*string, bool)`

GetHostpci23Ok returns a tuple with the Hostpci23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci23

`func (o *GetVMConfig200ResponseData) SetHostpci23(v string)`

SetHostpci23 sets Hostpci23 field to given value.

### HasHostpci23

`func (o *GetVMConfig200ResponseData) HasHostpci23() bool`

HasHostpci23 returns a boolean if a field has been set.

### GetHostpci24

`func (o *GetVMConfig200ResponseData) GetHostpci24() string`

GetHostpci24 returns the Hostpci24 field if non-nil, zero value otherwise.

### GetHostpci24Ok

`func (o *GetVMConfig200ResponseData) GetHostpci24Ok() (*string, bool)`

GetHostpci24Ok returns a tuple with the Hostpci24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci24

`func (o *GetVMConfig200ResponseData) SetHostpci24(v string)`

SetHostpci24 sets Hostpci24 field to given value.

### HasHostpci24

`func (o *GetVMConfig200ResponseData) HasHostpci24() bool`

HasHostpci24 returns a boolean if a field has been set.

### GetHostpci25

`func (o *GetVMConfig200ResponseData) GetHostpci25() string`

GetHostpci25 returns the Hostpci25 field if non-nil, zero value otherwise.

### GetHostpci25Ok

`func (o *GetVMConfig200ResponseData) GetHostpci25Ok() (*string, bool)`

GetHostpci25Ok returns a tuple with the Hostpci25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci25

`func (o *GetVMConfig200ResponseData) SetHostpci25(v string)`

SetHostpci25 sets Hostpci25 field to given value.

### HasHostpci25

`func (o *GetVMConfig200ResponseData) HasHostpci25() bool`

HasHostpci25 returns a boolean if a field has been set.

### GetHostpci26

`func (o *GetVMConfig200ResponseData) GetHostpci26() string`

GetHostpci26 returns the Hostpci26 field if non-nil, zero value otherwise.

### GetHostpci26Ok

`func (o *GetVMConfig200ResponseData) GetHostpci26Ok() (*string, bool)`

GetHostpci26Ok returns a tuple with the Hostpci26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci26

`func (o *GetVMConfig200ResponseData) SetHostpci26(v string)`

SetHostpci26 sets Hostpci26 field to given value.

### HasHostpci26

`func (o *GetVMConfig200ResponseData) HasHostpci26() bool`

HasHostpci26 returns a boolean if a field has been set.

### GetHostpci27

`func (o *GetVMConfig200ResponseData) GetHostpci27() string`

GetHostpci27 returns the Hostpci27 field if non-nil, zero value otherwise.

### GetHostpci27Ok

`func (o *GetVMConfig200ResponseData) GetHostpci27Ok() (*string, bool)`

GetHostpci27Ok returns a tuple with the Hostpci27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci27

`func (o *GetVMConfig200ResponseData) SetHostpci27(v string)`

SetHostpci27 sets Hostpci27 field to given value.

### HasHostpci27

`func (o *GetVMConfig200ResponseData) HasHostpci27() bool`

HasHostpci27 returns a boolean if a field has been set.

### GetHostpci28

`func (o *GetVMConfig200ResponseData) GetHostpci28() string`

GetHostpci28 returns the Hostpci28 field if non-nil, zero value otherwise.

### GetHostpci28Ok

`func (o *GetVMConfig200ResponseData) GetHostpci28Ok() (*string, bool)`

GetHostpci28Ok returns a tuple with the Hostpci28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci28

`func (o *GetVMConfig200ResponseData) SetHostpci28(v string)`

SetHostpci28 sets Hostpci28 field to given value.

### HasHostpci28

`func (o *GetVMConfig200ResponseData) HasHostpci28() bool`

HasHostpci28 returns a boolean if a field has been set.

### GetHostpci29

`func (o *GetVMConfig200ResponseData) GetHostpci29() string`

GetHostpci29 returns the Hostpci29 field if non-nil, zero value otherwise.

### GetHostpci29Ok

`func (o *GetVMConfig200ResponseData) GetHostpci29Ok() (*string, bool)`

GetHostpci29Ok returns a tuple with the Hostpci29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci29

`func (o *GetVMConfig200ResponseData) SetHostpci29(v string)`

SetHostpci29 sets Hostpci29 field to given value.

### HasHostpci29

`func (o *GetVMConfig200ResponseData) HasHostpci29() bool`

HasHostpci29 returns a boolean if a field has been set.

### GetHotplug

`func (o *GetVMConfig200ResponseData) GetHotplug() string`

GetHotplug returns the Hotplug field if non-nil, zero value otherwise.

### GetHotplugOk

`func (o *GetVMConfig200ResponseData) GetHotplugOk() (*string, bool)`

GetHotplugOk returns a tuple with the Hotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplug

`func (o *GetVMConfig200ResponseData) SetHotplug(v string)`

SetHotplug sets Hotplug field to given value.

### HasHotplug

`func (o *GetVMConfig200ResponseData) HasHotplug() bool`

HasHotplug returns a boolean if a field has been set.

### GetHugepages

`func (o *GetVMConfig200ResponseData) GetHugepages() string`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *GetVMConfig200ResponseData) GetHugepagesOk() (*string, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *GetVMConfig200ResponseData) SetHugepages(v string)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *GetVMConfig200ResponseData) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetIde0

`func (o *GetVMConfig200ResponseData) GetIde0() string`

GetIde0 returns the Ide0 field if non-nil, zero value otherwise.

### GetIde0Ok

`func (o *GetVMConfig200ResponseData) GetIde0Ok() (*string, bool)`

GetIde0Ok returns a tuple with the Ide0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde0

`func (o *GetVMConfig200ResponseData) SetIde0(v string)`

SetIde0 sets Ide0 field to given value.

### HasIde0

`func (o *GetVMConfig200ResponseData) HasIde0() bool`

HasIde0 returns a boolean if a field has been set.

### GetIde1

`func (o *GetVMConfig200ResponseData) GetIde1() string`

GetIde1 returns the Ide1 field if non-nil, zero value otherwise.

### GetIde1Ok

`func (o *GetVMConfig200ResponseData) GetIde1Ok() (*string, bool)`

GetIde1Ok returns a tuple with the Ide1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde1

`func (o *GetVMConfig200ResponseData) SetIde1(v string)`

SetIde1 sets Ide1 field to given value.

### HasIde1

`func (o *GetVMConfig200ResponseData) HasIde1() bool`

HasIde1 returns a boolean if a field has been set.

### GetIde2

`func (o *GetVMConfig200ResponseData) GetIde2() string`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *GetVMConfig200ResponseData) GetIde2Ok() (*string, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *GetVMConfig200ResponseData) SetIde2(v string)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *GetVMConfig200ResponseData) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetIde3

`func (o *GetVMConfig200ResponseData) GetIde3() string`

GetIde3 returns the Ide3 field if non-nil, zero value otherwise.

### GetIde3Ok

`func (o *GetVMConfig200ResponseData) GetIde3Ok() (*string, bool)`

GetIde3Ok returns a tuple with the Ide3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde3

`func (o *GetVMConfig200ResponseData) SetIde3(v string)`

SetIde3 sets Ide3 field to given value.

### HasIde3

`func (o *GetVMConfig200ResponseData) HasIde3() bool`

HasIde3 returns a boolean if a field has been set.

### GetIpconfig0

`func (o *GetVMConfig200ResponseData) GetIpconfig0() string`

GetIpconfig0 returns the Ipconfig0 field if non-nil, zero value otherwise.

### GetIpconfig0Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig0Ok() (*string, bool)`

GetIpconfig0Ok returns a tuple with the Ipconfig0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig0

`func (o *GetVMConfig200ResponseData) SetIpconfig0(v string)`

SetIpconfig0 sets Ipconfig0 field to given value.

### HasIpconfig0

`func (o *GetVMConfig200ResponseData) HasIpconfig0() bool`

HasIpconfig0 returns a boolean if a field has been set.

### GetIpconfig1

`func (o *GetVMConfig200ResponseData) GetIpconfig1() string`

GetIpconfig1 returns the Ipconfig1 field if non-nil, zero value otherwise.

### GetIpconfig1Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig1Ok() (*string, bool)`

GetIpconfig1Ok returns a tuple with the Ipconfig1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig1

`func (o *GetVMConfig200ResponseData) SetIpconfig1(v string)`

SetIpconfig1 sets Ipconfig1 field to given value.

### HasIpconfig1

`func (o *GetVMConfig200ResponseData) HasIpconfig1() bool`

HasIpconfig1 returns a boolean if a field has been set.

### GetIpconfig2

`func (o *GetVMConfig200ResponseData) GetIpconfig2() string`

GetIpconfig2 returns the Ipconfig2 field if non-nil, zero value otherwise.

### GetIpconfig2Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig2Ok() (*string, bool)`

GetIpconfig2Ok returns a tuple with the Ipconfig2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig2

`func (o *GetVMConfig200ResponseData) SetIpconfig2(v string)`

SetIpconfig2 sets Ipconfig2 field to given value.

### HasIpconfig2

`func (o *GetVMConfig200ResponseData) HasIpconfig2() bool`

HasIpconfig2 returns a boolean if a field has been set.

### GetIpconfig3

`func (o *GetVMConfig200ResponseData) GetIpconfig3() string`

GetIpconfig3 returns the Ipconfig3 field if non-nil, zero value otherwise.

### GetIpconfig3Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig3Ok() (*string, bool)`

GetIpconfig3Ok returns a tuple with the Ipconfig3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig3

`func (o *GetVMConfig200ResponseData) SetIpconfig3(v string)`

SetIpconfig3 sets Ipconfig3 field to given value.

### HasIpconfig3

`func (o *GetVMConfig200ResponseData) HasIpconfig3() bool`

HasIpconfig3 returns a boolean if a field has been set.

### GetIpconfig4

`func (o *GetVMConfig200ResponseData) GetIpconfig4() string`

GetIpconfig4 returns the Ipconfig4 field if non-nil, zero value otherwise.

### GetIpconfig4Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig4Ok() (*string, bool)`

GetIpconfig4Ok returns a tuple with the Ipconfig4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig4

`func (o *GetVMConfig200ResponseData) SetIpconfig4(v string)`

SetIpconfig4 sets Ipconfig4 field to given value.

### HasIpconfig4

`func (o *GetVMConfig200ResponseData) HasIpconfig4() bool`

HasIpconfig4 returns a boolean if a field has been set.

### GetIpconfig5

`func (o *GetVMConfig200ResponseData) GetIpconfig5() string`

GetIpconfig5 returns the Ipconfig5 field if non-nil, zero value otherwise.

### GetIpconfig5Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig5Ok() (*string, bool)`

GetIpconfig5Ok returns a tuple with the Ipconfig5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig5

`func (o *GetVMConfig200ResponseData) SetIpconfig5(v string)`

SetIpconfig5 sets Ipconfig5 field to given value.

### HasIpconfig5

`func (o *GetVMConfig200ResponseData) HasIpconfig5() bool`

HasIpconfig5 returns a boolean if a field has been set.

### GetIpconfig6

`func (o *GetVMConfig200ResponseData) GetIpconfig6() string`

GetIpconfig6 returns the Ipconfig6 field if non-nil, zero value otherwise.

### GetIpconfig6Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig6Ok() (*string, bool)`

GetIpconfig6Ok returns a tuple with the Ipconfig6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig6

`func (o *GetVMConfig200ResponseData) SetIpconfig6(v string)`

SetIpconfig6 sets Ipconfig6 field to given value.

### HasIpconfig6

`func (o *GetVMConfig200ResponseData) HasIpconfig6() bool`

HasIpconfig6 returns a boolean if a field has been set.

### GetIpconfig7

`func (o *GetVMConfig200ResponseData) GetIpconfig7() string`

GetIpconfig7 returns the Ipconfig7 field if non-nil, zero value otherwise.

### GetIpconfig7Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig7Ok() (*string, bool)`

GetIpconfig7Ok returns a tuple with the Ipconfig7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig7

`func (o *GetVMConfig200ResponseData) SetIpconfig7(v string)`

SetIpconfig7 sets Ipconfig7 field to given value.

### HasIpconfig7

`func (o *GetVMConfig200ResponseData) HasIpconfig7() bool`

HasIpconfig7 returns a boolean if a field has been set.

### GetIpconfig8

`func (o *GetVMConfig200ResponseData) GetIpconfig8() string`

GetIpconfig8 returns the Ipconfig8 field if non-nil, zero value otherwise.

### GetIpconfig8Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig8Ok() (*string, bool)`

GetIpconfig8Ok returns a tuple with the Ipconfig8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig8

`func (o *GetVMConfig200ResponseData) SetIpconfig8(v string)`

SetIpconfig8 sets Ipconfig8 field to given value.

### HasIpconfig8

`func (o *GetVMConfig200ResponseData) HasIpconfig8() bool`

HasIpconfig8 returns a boolean if a field has been set.

### GetIpconfig9

`func (o *GetVMConfig200ResponseData) GetIpconfig9() string`

GetIpconfig9 returns the Ipconfig9 field if non-nil, zero value otherwise.

### GetIpconfig9Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig9Ok() (*string, bool)`

GetIpconfig9Ok returns a tuple with the Ipconfig9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig9

`func (o *GetVMConfig200ResponseData) SetIpconfig9(v string)`

SetIpconfig9 sets Ipconfig9 field to given value.

### HasIpconfig9

`func (o *GetVMConfig200ResponseData) HasIpconfig9() bool`

HasIpconfig9 returns a boolean if a field has been set.

### GetIpconfig10

`func (o *GetVMConfig200ResponseData) GetIpconfig10() string`

GetIpconfig10 returns the Ipconfig10 field if non-nil, zero value otherwise.

### GetIpconfig10Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig10Ok() (*string, bool)`

GetIpconfig10Ok returns a tuple with the Ipconfig10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig10

`func (o *GetVMConfig200ResponseData) SetIpconfig10(v string)`

SetIpconfig10 sets Ipconfig10 field to given value.

### HasIpconfig10

`func (o *GetVMConfig200ResponseData) HasIpconfig10() bool`

HasIpconfig10 returns a boolean if a field has been set.

### GetIpconfig11

`func (o *GetVMConfig200ResponseData) GetIpconfig11() string`

GetIpconfig11 returns the Ipconfig11 field if non-nil, zero value otherwise.

### GetIpconfig11Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig11Ok() (*string, bool)`

GetIpconfig11Ok returns a tuple with the Ipconfig11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig11

`func (o *GetVMConfig200ResponseData) SetIpconfig11(v string)`

SetIpconfig11 sets Ipconfig11 field to given value.

### HasIpconfig11

`func (o *GetVMConfig200ResponseData) HasIpconfig11() bool`

HasIpconfig11 returns a boolean if a field has been set.

### GetIpconfig12

`func (o *GetVMConfig200ResponseData) GetIpconfig12() string`

GetIpconfig12 returns the Ipconfig12 field if non-nil, zero value otherwise.

### GetIpconfig12Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig12Ok() (*string, bool)`

GetIpconfig12Ok returns a tuple with the Ipconfig12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig12

`func (o *GetVMConfig200ResponseData) SetIpconfig12(v string)`

SetIpconfig12 sets Ipconfig12 field to given value.

### HasIpconfig12

`func (o *GetVMConfig200ResponseData) HasIpconfig12() bool`

HasIpconfig12 returns a boolean if a field has been set.

### GetIpconfig13

`func (o *GetVMConfig200ResponseData) GetIpconfig13() string`

GetIpconfig13 returns the Ipconfig13 field if non-nil, zero value otherwise.

### GetIpconfig13Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig13Ok() (*string, bool)`

GetIpconfig13Ok returns a tuple with the Ipconfig13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig13

`func (o *GetVMConfig200ResponseData) SetIpconfig13(v string)`

SetIpconfig13 sets Ipconfig13 field to given value.

### HasIpconfig13

`func (o *GetVMConfig200ResponseData) HasIpconfig13() bool`

HasIpconfig13 returns a boolean if a field has been set.

### GetIpconfig14

`func (o *GetVMConfig200ResponseData) GetIpconfig14() string`

GetIpconfig14 returns the Ipconfig14 field if non-nil, zero value otherwise.

### GetIpconfig14Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig14Ok() (*string, bool)`

GetIpconfig14Ok returns a tuple with the Ipconfig14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig14

`func (o *GetVMConfig200ResponseData) SetIpconfig14(v string)`

SetIpconfig14 sets Ipconfig14 field to given value.

### HasIpconfig14

`func (o *GetVMConfig200ResponseData) HasIpconfig14() bool`

HasIpconfig14 returns a boolean if a field has been set.

### GetIpconfig15

`func (o *GetVMConfig200ResponseData) GetIpconfig15() string`

GetIpconfig15 returns the Ipconfig15 field if non-nil, zero value otherwise.

### GetIpconfig15Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig15Ok() (*string, bool)`

GetIpconfig15Ok returns a tuple with the Ipconfig15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig15

`func (o *GetVMConfig200ResponseData) SetIpconfig15(v string)`

SetIpconfig15 sets Ipconfig15 field to given value.

### HasIpconfig15

`func (o *GetVMConfig200ResponseData) HasIpconfig15() bool`

HasIpconfig15 returns a boolean if a field has been set.

### GetIpconfig16

`func (o *GetVMConfig200ResponseData) GetIpconfig16() string`

GetIpconfig16 returns the Ipconfig16 field if non-nil, zero value otherwise.

### GetIpconfig16Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig16Ok() (*string, bool)`

GetIpconfig16Ok returns a tuple with the Ipconfig16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig16

`func (o *GetVMConfig200ResponseData) SetIpconfig16(v string)`

SetIpconfig16 sets Ipconfig16 field to given value.

### HasIpconfig16

`func (o *GetVMConfig200ResponseData) HasIpconfig16() bool`

HasIpconfig16 returns a boolean if a field has been set.

### GetIpconfig17

`func (o *GetVMConfig200ResponseData) GetIpconfig17() string`

GetIpconfig17 returns the Ipconfig17 field if non-nil, zero value otherwise.

### GetIpconfig17Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig17Ok() (*string, bool)`

GetIpconfig17Ok returns a tuple with the Ipconfig17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig17

`func (o *GetVMConfig200ResponseData) SetIpconfig17(v string)`

SetIpconfig17 sets Ipconfig17 field to given value.

### HasIpconfig17

`func (o *GetVMConfig200ResponseData) HasIpconfig17() bool`

HasIpconfig17 returns a boolean if a field has been set.

### GetIpconfig18

`func (o *GetVMConfig200ResponseData) GetIpconfig18() string`

GetIpconfig18 returns the Ipconfig18 field if non-nil, zero value otherwise.

### GetIpconfig18Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig18Ok() (*string, bool)`

GetIpconfig18Ok returns a tuple with the Ipconfig18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig18

`func (o *GetVMConfig200ResponseData) SetIpconfig18(v string)`

SetIpconfig18 sets Ipconfig18 field to given value.

### HasIpconfig18

`func (o *GetVMConfig200ResponseData) HasIpconfig18() bool`

HasIpconfig18 returns a boolean if a field has been set.

### GetIpconfig19

`func (o *GetVMConfig200ResponseData) GetIpconfig19() string`

GetIpconfig19 returns the Ipconfig19 field if non-nil, zero value otherwise.

### GetIpconfig19Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig19Ok() (*string, bool)`

GetIpconfig19Ok returns a tuple with the Ipconfig19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig19

`func (o *GetVMConfig200ResponseData) SetIpconfig19(v string)`

SetIpconfig19 sets Ipconfig19 field to given value.

### HasIpconfig19

`func (o *GetVMConfig200ResponseData) HasIpconfig19() bool`

HasIpconfig19 returns a boolean if a field has been set.

### GetIpconfig20

`func (o *GetVMConfig200ResponseData) GetIpconfig20() string`

GetIpconfig20 returns the Ipconfig20 field if non-nil, zero value otherwise.

### GetIpconfig20Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig20Ok() (*string, bool)`

GetIpconfig20Ok returns a tuple with the Ipconfig20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig20

`func (o *GetVMConfig200ResponseData) SetIpconfig20(v string)`

SetIpconfig20 sets Ipconfig20 field to given value.

### HasIpconfig20

`func (o *GetVMConfig200ResponseData) HasIpconfig20() bool`

HasIpconfig20 returns a boolean if a field has been set.

### GetIpconfig21

`func (o *GetVMConfig200ResponseData) GetIpconfig21() string`

GetIpconfig21 returns the Ipconfig21 field if non-nil, zero value otherwise.

### GetIpconfig21Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig21Ok() (*string, bool)`

GetIpconfig21Ok returns a tuple with the Ipconfig21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig21

`func (o *GetVMConfig200ResponseData) SetIpconfig21(v string)`

SetIpconfig21 sets Ipconfig21 field to given value.

### HasIpconfig21

`func (o *GetVMConfig200ResponseData) HasIpconfig21() bool`

HasIpconfig21 returns a boolean if a field has been set.

### GetIpconfig22

`func (o *GetVMConfig200ResponseData) GetIpconfig22() string`

GetIpconfig22 returns the Ipconfig22 field if non-nil, zero value otherwise.

### GetIpconfig22Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig22Ok() (*string, bool)`

GetIpconfig22Ok returns a tuple with the Ipconfig22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig22

`func (o *GetVMConfig200ResponseData) SetIpconfig22(v string)`

SetIpconfig22 sets Ipconfig22 field to given value.

### HasIpconfig22

`func (o *GetVMConfig200ResponseData) HasIpconfig22() bool`

HasIpconfig22 returns a boolean if a field has been set.

### GetIpconfig23

`func (o *GetVMConfig200ResponseData) GetIpconfig23() string`

GetIpconfig23 returns the Ipconfig23 field if non-nil, zero value otherwise.

### GetIpconfig23Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig23Ok() (*string, bool)`

GetIpconfig23Ok returns a tuple with the Ipconfig23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig23

`func (o *GetVMConfig200ResponseData) SetIpconfig23(v string)`

SetIpconfig23 sets Ipconfig23 field to given value.

### HasIpconfig23

`func (o *GetVMConfig200ResponseData) HasIpconfig23() bool`

HasIpconfig23 returns a boolean if a field has been set.

### GetIpconfig24

`func (o *GetVMConfig200ResponseData) GetIpconfig24() string`

GetIpconfig24 returns the Ipconfig24 field if non-nil, zero value otherwise.

### GetIpconfig24Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig24Ok() (*string, bool)`

GetIpconfig24Ok returns a tuple with the Ipconfig24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig24

`func (o *GetVMConfig200ResponseData) SetIpconfig24(v string)`

SetIpconfig24 sets Ipconfig24 field to given value.

### HasIpconfig24

`func (o *GetVMConfig200ResponseData) HasIpconfig24() bool`

HasIpconfig24 returns a boolean if a field has been set.

### GetIpconfig25

`func (o *GetVMConfig200ResponseData) GetIpconfig25() string`

GetIpconfig25 returns the Ipconfig25 field if non-nil, zero value otherwise.

### GetIpconfig25Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig25Ok() (*string, bool)`

GetIpconfig25Ok returns a tuple with the Ipconfig25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig25

`func (o *GetVMConfig200ResponseData) SetIpconfig25(v string)`

SetIpconfig25 sets Ipconfig25 field to given value.

### HasIpconfig25

`func (o *GetVMConfig200ResponseData) HasIpconfig25() bool`

HasIpconfig25 returns a boolean if a field has been set.

### GetIpconfig26

`func (o *GetVMConfig200ResponseData) GetIpconfig26() string`

GetIpconfig26 returns the Ipconfig26 field if non-nil, zero value otherwise.

### GetIpconfig26Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig26Ok() (*string, bool)`

GetIpconfig26Ok returns a tuple with the Ipconfig26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig26

`func (o *GetVMConfig200ResponseData) SetIpconfig26(v string)`

SetIpconfig26 sets Ipconfig26 field to given value.

### HasIpconfig26

`func (o *GetVMConfig200ResponseData) HasIpconfig26() bool`

HasIpconfig26 returns a boolean if a field has been set.

### GetIpconfig27

`func (o *GetVMConfig200ResponseData) GetIpconfig27() string`

GetIpconfig27 returns the Ipconfig27 field if non-nil, zero value otherwise.

### GetIpconfig27Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig27Ok() (*string, bool)`

GetIpconfig27Ok returns a tuple with the Ipconfig27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig27

`func (o *GetVMConfig200ResponseData) SetIpconfig27(v string)`

SetIpconfig27 sets Ipconfig27 field to given value.

### HasIpconfig27

`func (o *GetVMConfig200ResponseData) HasIpconfig27() bool`

HasIpconfig27 returns a boolean if a field has been set.

### GetIpconfig28

`func (o *GetVMConfig200ResponseData) GetIpconfig28() string`

GetIpconfig28 returns the Ipconfig28 field if non-nil, zero value otherwise.

### GetIpconfig28Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig28Ok() (*string, bool)`

GetIpconfig28Ok returns a tuple with the Ipconfig28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig28

`func (o *GetVMConfig200ResponseData) SetIpconfig28(v string)`

SetIpconfig28 sets Ipconfig28 field to given value.

### HasIpconfig28

`func (o *GetVMConfig200ResponseData) HasIpconfig28() bool`

HasIpconfig28 returns a boolean if a field has been set.

### GetIpconfig29

`func (o *GetVMConfig200ResponseData) GetIpconfig29() string`

GetIpconfig29 returns the Ipconfig29 field if non-nil, zero value otherwise.

### GetIpconfig29Ok

`func (o *GetVMConfig200ResponseData) GetIpconfig29Ok() (*string, bool)`

GetIpconfig29Ok returns a tuple with the Ipconfig29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig29

`func (o *GetVMConfig200ResponseData) SetIpconfig29(v string)`

SetIpconfig29 sets Ipconfig29 field to given value.

### HasIpconfig29

`func (o *GetVMConfig200ResponseData) HasIpconfig29() bool`

HasIpconfig29 returns a boolean if a field has been set.

### GetIvshmem

`func (o *GetVMConfig200ResponseData) GetIvshmem() string`

GetIvshmem returns the Ivshmem field if non-nil, zero value otherwise.

### GetIvshmemOk

`func (o *GetVMConfig200ResponseData) GetIvshmemOk() (*string, bool)`

GetIvshmemOk returns a tuple with the Ivshmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvshmem

`func (o *GetVMConfig200ResponseData) SetIvshmem(v string)`

SetIvshmem sets Ivshmem field to given value.

### HasIvshmem

`func (o *GetVMConfig200ResponseData) HasIvshmem() bool`

HasIvshmem returns a boolean if a field has been set.

### GetKeephugepages

`func (o *GetVMConfig200ResponseData) GetKeephugepages() int32`

GetKeephugepages returns the Keephugepages field if non-nil, zero value otherwise.

### GetKeephugepagesOk

`func (o *GetVMConfig200ResponseData) GetKeephugepagesOk() (*int32, bool)`

GetKeephugepagesOk returns a tuple with the Keephugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeephugepages

`func (o *GetVMConfig200ResponseData) SetKeephugepages(v int32)`

SetKeephugepages sets Keephugepages field to given value.

### HasKeephugepages

`func (o *GetVMConfig200ResponseData) HasKeephugepages() bool`

HasKeephugepages returns a boolean if a field has been set.

### GetKeyboard

`func (o *GetVMConfig200ResponseData) GetKeyboard() string`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *GetVMConfig200ResponseData) GetKeyboardOk() (*string, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *GetVMConfig200ResponseData) SetKeyboard(v string)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *GetVMConfig200ResponseData) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetKvm

`func (o *GetVMConfig200ResponseData) GetKvm() int32`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *GetVMConfig200ResponseData) GetKvmOk() (*int32, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *GetVMConfig200ResponseData) SetKvm(v int32)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *GetVMConfig200ResponseData) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLocaltime

`func (o *GetVMConfig200ResponseData) GetLocaltime() int32`

GetLocaltime returns the Localtime field if non-nil, zero value otherwise.

### GetLocaltimeOk

`func (o *GetVMConfig200ResponseData) GetLocaltimeOk() (*int32, bool)`

GetLocaltimeOk returns a tuple with the Localtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaltime

`func (o *GetVMConfig200ResponseData) SetLocaltime(v int32)`

SetLocaltime sets Localtime field to given value.

### HasLocaltime

`func (o *GetVMConfig200ResponseData) HasLocaltime() bool`

HasLocaltime returns a boolean if a field has been set.

### GetLock

`func (o *GetVMConfig200ResponseData) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetVMConfig200ResponseData) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetVMConfig200ResponseData) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetVMConfig200ResponseData) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMachine

`func (o *GetVMConfig200ResponseData) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *GetVMConfig200ResponseData) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *GetVMConfig200ResponseData) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *GetVMConfig200ResponseData) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *GetVMConfig200ResponseData) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetVMConfig200ResponseData) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetVMConfig200ResponseData) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetVMConfig200ResponseData) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrateDowntime

`func (o *GetVMConfig200ResponseData) GetMigrateDowntime() float32`

GetMigrateDowntime returns the MigrateDowntime field if non-nil, zero value otherwise.

### GetMigrateDowntimeOk

`func (o *GetVMConfig200ResponseData) GetMigrateDowntimeOk() (*float32, bool)`

GetMigrateDowntimeOk returns a tuple with the MigrateDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateDowntime

`func (o *GetVMConfig200ResponseData) SetMigrateDowntime(v float32)`

SetMigrateDowntime sets MigrateDowntime field to given value.

### HasMigrateDowntime

`func (o *GetVMConfig200ResponseData) HasMigrateDowntime() bool`

HasMigrateDowntime returns a boolean if a field has been set.

### GetMigrateSpeed

`func (o *GetVMConfig200ResponseData) GetMigrateSpeed() int64`

GetMigrateSpeed returns the MigrateSpeed field if non-nil, zero value otherwise.

### GetMigrateSpeedOk

`func (o *GetVMConfig200ResponseData) GetMigrateSpeedOk() (*int64, bool)`

GetMigrateSpeedOk returns a tuple with the MigrateSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateSpeed

`func (o *GetVMConfig200ResponseData) SetMigrateSpeed(v int64)`

SetMigrateSpeed sets MigrateSpeed field to given value.

### HasMigrateSpeed

`func (o *GetVMConfig200ResponseData) HasMigrateSpeed() bool`

HasMigrateSpeed returns a boolean if a field has been set.

### GetName

`func (o *GetVMConfig200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVMConfig200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVMConfig200ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVMConfig200ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *GetVMConfig200ResponseData) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *GetVMConfig200ResponseData) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *GetVMConfig200ResponseData) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *GetVMConfig200ResponseData) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *GetVMConfig200ResponseData) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *GetVMConfig200ResponseData) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *GetVMConfig200ResponseData) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *GetVMConfig200ResponseData) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *GetVMConfig200ResponseData) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *GetVMConfig200ResponseData) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *GetVMConfig200ResponseData) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *GetVMConfig200ResponseData) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *GetVMConfig200ResponseData) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *GetVMConfig200ResponseData) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *GetVMConfig200ResponseData) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *GetVMConfig200ResponseData) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *GetVMConfig200ResponseData) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *GetVMConfig200ResponseData) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *GetVMConfig200ResponseData) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *GetVMConfig200ResponseData) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *GetVMConfig200ResponseData) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *GetVMConfig200ResponseData) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *GetVMConfig200ResponseData) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *GetVMConfig200ResponseData) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *GetVMConfig200ResponseData) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *GetVMConfig200ResponseData) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *GetVMConfig200ResponseData) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *GetVMConfig200ResponseData) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *GetVMConfig200ResponseData) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *GetVMConfig200ResponseData) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *GetVMConfig200ResponseData) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *GetVMConfig200ResponseData) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *GetVMConfig200ResponseData) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *GetVMConfig200ResponseData) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *GetVMConfig200ResponseData) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *GetVMConfig200ResponseData) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *GetVMConfig200ResponseData) GetNet8() string`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *GetVMConfig200ResponseData) GetNet8Ok() (*string, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *GetVMConfig200ResponseData) SetNet8(v string)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *GetVMConfig200ResponseData) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *GetVMConfig200ResponseData) GetNet9() string`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *GetVMConfig200ResponseData) GetNet9Ok() (*string, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *GetVMConfig200ResponseData) SetNet9(v string)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *GetVMConfig200ResponseData) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *GetVMConfig200ResponseData) GetNet10() string`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *GetVMConfig200ResponseData) GetNet10Ok() (*string, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *GetVMConfig200ResponseData) SetNet10(v string)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *GetVMConfig200ResponseData) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *GetVMConfig200ResponseData) GetNet11() string`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *GetVMConfig200ResponseData) GetNet11Ok() (*string, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *GetVMConfig200ResponseData) SetNet11(v string)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *GetVMConfig200ResponseData) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *GetVMConfig200ResponseData) GetNet12() string`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *GetVMConfig200ResponseData) GetNet12Ok() (*string, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *GetVMConfig200ResponseData) SetNet12(v string)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *GetVMConfig200ResponseData) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *GetVMConfig200ResponseData) GetNet13() string`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *GetVMConfig200ResponseData) GetNet13Ok() (*string, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *GetVMConfig200ResponseData) SetNet13(v string)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *GetVMConfig200ResponseData) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *GetVMConfig200ResponseData) GetNet14() string`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *GetVMConfig200ResponseData) GetNet14Ok() (*string, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *GetVMConfig200ResponseData) SetNet14(v string)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *GetVMConfig200ResponseData) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *GetVMConfig200ResponseData) GetNet15() string`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *GetVMConfig200ResponseData) GetNet15Ok() (*string, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *GetVMConfig200ResponseData) SetNet15(v string)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *GetVMConfig200ResponseData) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *GetVMConfig200ResponseData) GetNet16() string`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *GetVMConfig200ResponseData) GetNet16Ok() (*string, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *GetVMConfig200ResponseData) SetNet16(v string)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *GetVMConfig200ResponseData) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *GetVMConfig200ResponseData) GetNet17() string`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *GetVMConfig200ResponseData) GetNet17Ok() (*string, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *GetVMConfig200ResponseData) SetNet17(v string)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *GetVMConfig200ResponseData) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *GetVMConfig200ResponseData) GetNet18() string`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *GetVMConfig200ResponseData) GetNet18Ok() (*string, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *GetVMConfig200ResponseData) SetNet18(v string)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *GetVMConfig200ResponseData) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *GetVMConfig200ResponseData) GetNet19() string`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *GetVMConfig200ResponseData) GetNet19Ok() (*string, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *GetVMConfig200ResponseData) SetNet19(v string)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *GetVMConfig200ResponseData) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *GetVMConfig200ResponseData) GetNet20() string`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *GetVMConfig200ResponseData) GetNet20Ok() (*string, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *GetVMConfig200ResponseData) SetNet20(v string)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *GetVMConfig200ResponseData) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *GetVMConfig200ResponseData) GetNet21() string`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *GetVMConfig200ResponseData) GetNet21Ok() (*string, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *GetVMConfig200ResponseData) SetNet21(v string)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *GetVMConfig200ResponseData) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *GetVMConfig200ResponseData) GetNet22() string`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *GetVMConfig200ResponseData) GetNet22Ok() (*string, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *GetVMConfig200ResponseData) SetNet22(v string)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *GetVMConfig200ResponseData) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *GetVMConfig200ResponseData) GetNet23() string`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *GetVMConfig200ResponseData) GetNet23Ok() (*string, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *GetVMConfig200ResponseData) SetNet23(v string)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *GetVMConfig200ResponseData) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *GetVMConfig200ResponseData) GetNet24() string`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *GetVMConfig200ResponseData) GetNet24Ok() (*string, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *GetVMConfig200ResponseData) SetNet24(v string)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *GetVMConfig200ResponseData) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *GetVMConfig200ResponseData) GetNet25() string`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *GetVMConfig200ResponseData) GetNet25Ok() (*string, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *GetVMConfig200ResponseData) SetNet25(v string)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *GetVMConfig200ResponseData) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *GetVMConfig200ResponseData) GetNet26() string`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *GetVMConfig200ResponseData) GetNet26Ok() (*string, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *GetVMConfig200ResponseData) SetNet26(v string)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *GetVMConfig200ResponseData) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *GetVMConfig200ResponseData) GetNet27() string`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *GetVMConfig200ResponseData) GetNet27Ok() (*string, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *GetVMConfig200ResponseData) SetNet27(v string)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *GetVMConfig200ResponseData) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *GetVMConfig200ResponseData) GetNet28() string`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *GetVMConfig200ResponseData) GetNet28Ok() (*string, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *GetVMConfig200ResponseData) SetNet28(v string)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *GetVMConfig200ResponseData) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *GetVMConfig200ResponseData) GetNet29() string`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *GetVMConfig200ResponseData) GetNet29Ok() (*string, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *GetVMConfig200ResponseData) SetNet29(v string)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *GetVMConfig200ResponseData) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *GetVMConfig200ResponseData) GetNet30() string`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *GetVMConfig200ResponseData) GetNet30Ok() (*string, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *GetVMConfig200ResponseData) SetNet30(v string)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *GetVMConfig200ResponseData) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *GetVMConfig200ResponseData) GetNet31() string`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *GetVMConfig200ResponseData) GetNet31Ok() (*string, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *GetVMConfig200ResponseData) SetNet31(v string)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *GetVMConfig200ResponseData) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetNuma

`func (o *GetVMConfig200ResponseData) GetNuma() int32`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *GetVMConfig200ResponseData) GetNumaOk() (*int32, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *GetVMConfig200ResponseData) SetNuma(v int32)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *GetVMConfig200ResponseData) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetNuma0

`func (o *GetVMConfig200ResponseData) GetNuma0() string`

GetNuma0 returns the Numa0 field if non-nil, zero value otherwise.

### GetNuma0Ok

`func (o *GetVMConfig200ResponseData) GetNuma0Ok() (*string, bool)`

GetNuma0Ok returns a tuple with the Numa0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma0

`func (o *GetVMConfig200ResponseData) SetNuma0(v string)`

SetNuma0 sets Numa0 field to given value.

### HasNuma0

`func (o *GetVMConfig200ResponseData) HasNuma0() bool`

HasNuma0 returns a boolean if a field has been set.

### GetNuma1

`func (o *GetVMConfig200ResponseData) GetNuma1() string`

GetNuma1 returns the Numa1 field if non-nil, zero value otherwise.

### GetNuma1Ok

`func (o *GetVMConfig200ResponseData) GetNuma1Ok() (*string, bool)`

GetNuma1Ok returns a tuple with the Numa1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma1

`func (o *GetVMConfig200ResponseData) SetNuma1(v string)`

SetNuma1 sets Numa1 field to given value.

### HasNuma1

`func (o *GetVMConfig200ResponseData) HasNuma1() bool`

HasNuma1 returns a boolean if a field has been set.

### GetNuma2

`func (o *GetVMConfig200ResponseData) GetNuma2() string`

GetNuma2 returns the Numa2 field if non-nil, zero value otherwise.

### GetNuma2Ok

`func (o *GetVMConfig200ResponseData) GetNuma2Ok() (*string, bool)`

GetNuma2Ok returns a tuple with the Numa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma2

`func (o *GetVMConfig200ResponseData) SetNuma2(v string)`

SetNuma2 sets Numa2 field to given value.

### HasNuma2

`func (o *GetVMConfig200ResponseData) HasNuma2() bool`

HasNuma2 returns a boolean if a field has been set.

### GetNuma3

`func (o *GetVMConfig200ResponseData) GetNuma3() string`

GetNuma3 returns the Numa3 field if non-nil, zero value otherwise.

### GetNuma3Ok

`func (o *GetVMConfig200ResponseData) GetNuma3Ok() (*string, bool)`

GetNuma3Ok returns a tuple with the Numa3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma3

`func (o *GetVMConfig200ResponseData) SetNuma3(v string)`

SetNuma3 sets Numa3 field to given value.

### HasNuma3

`func (o *GetVMConfig200ResponseData) HasNuma3() bool`

HasNuma3 returns a boolean if a field has been set.

### GetNuma4

`func (o *GetVMConfig200ResponseData) GetNuma4() string`

GetNuma4 returns the Numa4 field if non-nil, zero value otherwise.

### GetNuma4Ok

`func (o *GetVMConfig200ResponseData) GetNuma4Ok() (*string, bool)`

GetNuma4Ok returns a tuple with the Numa4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma4

`func (o *GetVMConfig200ResponseData) SetNuma4(v string)`

SetNuma4 sets Numa4 field to given value.

### HasNuma4

`func (o *GetVMConfig200ResponseData) HasNuma4() bool`

HasNuma4 returns a boolean if a field has been set.

### GetNuma5

`func (o *GetVMConfig200ResponseData) GetNuma5() string`

GetNuma5 returns the Numa5 field if non-nil, zero value otherwise.

### GetNuma5Ok

`func (o *GetVMConfig200ResponseData) GetNuma5Ok() (*string, bool)`

GetNuma5Ok returns a tuple with the Numa5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma5

`func (o *GetVMConfig200ResponseData) SetNuma5(v string)`

SetNuma5 sets Numa5 field to given value.

### HasNuma5

`func (o *GetVMConfig200ResponseData) HasNuma5() bool`

HasNuma5 returns a boolean if a field has been set.

### GetNuma6

`func (o *GetVMConfig200ResponseData) GetNuma6() string`

GetNuma6 returns the Numa6 field if non-nil, zero value otherwise.

### GetNuma6Ok

`func (o *GetVMConfig200ResponseData) GetNuma6Ok() (*string, bool)`

GetNuma6Ok returns a tuple with the Numa6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma6

`func (o *GetVMConfig200ResponseData) SetNuma6(v string)`

SetNuma6 sets Numa6 field to given value.

### HasNuma6

`func (o *GetVMConfig200ResponseData) HasNuma6() bool`

HasNuma6 returns a boolean if a field has been set.

### GetNuma7

`func (o *GetVMConfig200ResponseData) GetNuma7() string`

GetNuma7 returns the Numa7 field if non-nil, zero value otherwise.

### GetNuma7Ok

`func (o *GetVMConfig200ResponseData) GetNuma7Ok() (*string, bool)`

GetNuma7Ok returns a tuple with the Numa7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma7

`func (o *GetVMConfig200ResponseData) SetNuma7(v string)`

SetNuma7 sets Numa7 field to given value.

### HasNuma7

`func (o *GetVMConfig200ResponseData) HasNuma7() bool`

HasNuma7 returns a boolean if a field has been set.

### GetNuma8

`func (o *GetVMConfig200ResponseData) GetNuma8() string`

GetNuma8 returns the Numa8 field if non-nil, zero value otherwise.

### GetNuma8Ok

`func (o *GetVMConfig200ResponseData) GetNuma8Ok() (*string, bool)`

GetNuma8Ok returns a tuple with the Numa8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma8

`func (o *GetVMConfig200ResponseData) SetNuma8(v string)`

SetNuma8 sets Numa8 field to given value.

### HasNuma8

`func (o *GetVMConfig200ResponseData) HasNuma8() bool`

HasNuma8 returns a boolean if a field has been set.

### GetNuma9

`func (o *GetVMConfig200ResponseData) GetNuma9() string`

GetNuma9 returns the Numa9 field if non-nil, zero value otherwise.

### GetNuma9Ok

`func (o *GetVMConfig200ResponseData) GetNuma9Ok() (*string, bool)`

GetNuma9Ok returns a tuple with the Numa9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma9

`func (o *GetVMConfig200ResponseData) SetNuma9(v string)`

SetNuma9 sets Numa9 field to given value.

### HasNuma9

`func (o *GetVMConfig200ResponseData) HasNuma9() bool`

HasNuma9 returns a boolean if a field has been set.

### GetNuma10

`func (o *GetVMConfig200ResponseData) GetNuma10() string`

GetNuma10 returns the Numa10 field if non-nil, zero value otherwise.

### GetNuma10Ok

`func (o *GetVMConfig200ResponseData) GetNuma10Ok() (*string, bool)`

GetNuma10Ok returns a tuple with the Numa10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma10

`func (o *GetVMConfig200ResponseData) SetNuma10(v string)`

SetNuma10 sets Numa10 field to given value.

### HasNuma10

`func (o *GetVMConfig200ResponseData) HasNuma10() bool`

HasNuma10 returns a boolean if a field has been set.

### GetNuma11

`func (o *GetVMConfig200ResponseData) GetNuma11() string`

GetNuma11 returns the Numa11 field if non-nil, zero value otherwise.

### GetNuma11Ok

`func (o *GetVMConfig200ResponseData) GetNuma11Ok() (*string, bool)`

GetNuma11Ok returns a tuple with the Numa11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma11

`func (o *GetVMConfig200ResponseData) SetNuma11(v string)`

SetNuma11 sets Numa11 field to given value.

### HasNuma11

`func (o *GetVMConfig200ResponseData) HasNuma11() bool`

HasNuma11 returns a boolean if a field has been set.

### GetNuma12

`func (o *GetVMConfig200ResponseData) GetNuma12() string`

GetNuma12 returns the Numa12 field if non-nil, zero value otherwise.

### GetNuma12Ok

`func (o *GetVMConfig200ResponseData) GetNuma12Ok() (*string, bool)`

GetNuma12Ok returns a tuple with the Numa12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma12

`func (o *GetVMConfig200ResponseData) SetNuma12(v string)`

SetNuma12 sets Numa12 field to given value.

### HasNuma12

`func (o *GetVMConfig200ResponseData) HasNuma12() bool`

HasNuma12 returns a boolean if a field has been set.

### GetNuma13

`func (o *GetVMConfig200ResponseData) GetNuma13() string`

GetNuma13 returns the Numa13 field if non-nil, zero value otherwise.

### GetNuma13Ok

`func (o *GetVMConfig200ResponseData) GetNuma13Ok() (*string, bool)`

GetNuma13Ok returns a tuple with the Numa13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma13

`func (o *GetVMConfig200ResponseData) SetNuma13(v string)`

SetNuma13 sets Numa13 field to given value.

### HasNuma13

`func (o *GetVMConfig200ResponseData) HasNuma13() bool`

HasNuma13 returns a boolean if a field has been set.

### GetNuma14

`func (o *GetVMConfig200ResponseData) GetNuma14() string`

GetNuma14 returns the Numa14 field if non-nil, zero value otherwise.

### GetNuma14Ok

`func (o *GetVMConfig200ResponseData) GetNuma14Ok() (*string, bool)`

GetNuma14Ok returns a tuple with the Numa14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma14

`func (o *GetVMConfig200ResponseData) SetNuma14(v string)`

SetNuma14 sets Numa14 field to given value.

### HasNuma14

`func (o *GetVMConfig200ResponseData) HasNuma14() bool`

HasNuma14 returns a boolean if a field has been set.

### GetNuma15

`func (o *GetVMConfig200ResponseData) GetNuma15() string`

GetNuma15 returns the Numa15 field if non-nil, zero value otherwise.

### GetNuma15Ok

`func (o *GetVMConfig200ResponseData) GetNuma15Ok() (*string, bool)`

GetNuma15Ok returns a tuple with the Numa15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma15

`func (o *GetVMConfig200ResponseData) SetNuma15(v string)`

SetNuma15 sets Numa15 field to given value.

### HasNuma15

`func (o *GetVMConfig200ResponseData) HasNuma15() bool`

HasNuma15 returns a boolean if a field has been set.

### GetNuma16

`func (o *GetVMConfig200ResponseData) GetNuma16() string`

GetNuma16 returns the Numa16 field if non-nil, zero value otherwise.

### GetNuma16Ok

`func (o *GetVMConfig200ResponseData) GetNuma16Ok() (*string, bool)`

GetNuma16Ok returns a tuple with the Numa16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma16

`func (o *GetVMConfig200ResponseData) SetNuma16(v string)`

SetNuma16 sets Numa16 field to given value.

### HasNuma16

`func (o *GetVMConfig200ResponseData) HasNuma16() bool`

HasNuma16 returns a boolean if a field has been set.

### GetNuma17

`func (o *GetVMConfig200ResponseData) GetNuma17() string`

GetNuma17 returns the Numa17 field if non-nil, zero value otherwise.

### GetNuma17Ok

`func (o *GetVMConfig200ResponseData) GetNuma17Ok() (*string, bool)`

GetNuma17Ok returns a tuple with the Numa17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma17

`func (o *GetVMConfig200ResponseData) SetNuma17(v string)`

SetNuma17 sets Numa17 field to given value.

### HasNuma17

`func (o *GetVMConfig200ResponseData) HasNuma17() bool`

HasNuma17 returns a boolean if a field has been set.

### GetNuma18

`func (o *GetVMConfig200ResponseData) GetNuma18() string`

GetNuma18 returns the Numa18 field if non-nil, zero value otherwise.

### GetNuma18Ok

`func (o *GetVMConfig200ResponseData) GetNuma18Ok() (*string, bool)`

GetNuma18Ok returns a tuple with the Numa18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma18

`func (o *GetVMConfig200ResponseData) SetNuma18(v string)`

SetNuma18 sets Numa18 field to given value.

### HasNuma18

`func (o *GetVMConfig200ResponseData) HasNuma18() bool`

HasNuma18 returns a boolean if a field has been set.

### GetNuma19

`func (o *GetVMConfig200ResponseData) GetNuma19() string`

GetNuma19 returns the Numa19 field if non-nil, zero value otherwise.

### GetNuma19Ok

`func (o *GetVMConfig200ResponseData) GetNuma19Ok() (*string, bool)`

GetNuma19Ok returns a tuple with the Numa19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma19

`func (o *GetVMConfig200ResponseData) SetNuma19(v string)`

SetNuma19 sets Numa19 field to given value.

### HasNuma19

`func (o *GetVMConfig200ResponseData) HasNuma19() bool`

HasNuma19 returns a boolean if a field has been set.

### GetNuma20

`func (o *GetVMConfig200ResponseData) GetNuma20() string`

GetNuma20 returns the Numa20 field if non-nil, zero value otherwise.

### GetNuma20Ok

`func (o *GetVMConfig200ResponseData) GetNuma20Ok() (*string, bool)`

GetNuma20Ok returns a tuple with the Numa20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma20

`func (o *GetVMConfig200ResponseData) SetNuma20(v string)`

SetNuma20 sets Numa20 field to given value.

### HasNuma20

`func (o *GetVMConfig200ResponseData) HasNuma20() bool`

HasNuma20 returns a boolean if a field has been set.

### GetNuma21

`func (o *GetVMConfig200ResponseData) GetNuma21() string`

GetNuma21 returns the Numa21 field if non-nil, zero value otherwise.

### GetNuma21Ok

`func (o *GetVMConfig200ResponseData) GetNuma21Ok() (*string, bool)`

GetNuma21Ok returns a tuple with the Numa21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma21

`func (o *GetVMConfig200ResponseData) SetNuma21(v string)`

SetNuma21 sets Numa21 field to given value.

### HasNuma21

`func (o *GetVMConfig200ResponseData) HasNuma21() bool`

HasNuma21 returns a boolean if a field has been set.

### GetNuma22

`func (o *GetVMConfig200ResponseData) GetNuma22() string`

GetNuma22 returns the Numa22 field if non-nil, zero value otherwise.

### GetNuma22Ok

`func (o *GetVMConfig200ResponseData) GetNuma22Ok() (*string, bool)`

GetNuma22Ok returns a tuple with the Numa22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma22

`func (o *GetVMConfig200ResponseData) SetNuma22(v string)`

SetNuma22 sets Numa22 field to given value.

### HasNuma22

`func (o *GetVMConfig200ResponseData) HasNuma22() bool`

HasNuma22 returns a boolean if a field has been set.

### GetNuma23

`func (o *GetVMConfig200ResponseData) GetNuma23() string`

GetNuma23 returns the Numa23 field if non-nil, zero value otherwise.

### GetNuma23Ok

`func (o *GetVMConfig200ResponseData) GetNuma23Ok() (*string, bool)`

GetNuma23Ok returns a tuple with the Numa23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma23

`func (o *GetVMConfig200ResponseData) SetNuma23(v string)`

SetNuma23 sets Numa23 field to given value.

### HasNuma23

`func (o *GetVMConfig200ResponseData) HasNuma23() bool`

HasNuma23 returns a boolean if a field has been set.

### GetNuma24

`func (o *GetVMConfig200ResponseData) GetNuma24() string`

GetNuma24 returns the Numa24 field if non-nil, zero value otherwise.

### GetNuma24Ok

`func (o *GetVMConfig200ResponseData) GetNuma24Ok() (*string, bool)`

GetNuma24Ok returns a tuple with the Numa24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma24

`func (o *GetVMConfig200ResponseData) SetNuma24(v string)`

SetNuma24 sets Numa24 field to given value.

### HasNuma24

`func (o *GetVMConfig200ResponseData) HasNuma24() bool`

HasNuma24 returns a boolean if a field has been set.

### GetNuma25

`func (o *GetVMConfig200ResponseData) GetNuma25() string`

GetNuma25 returns the Numa25 field if non-nil, zero value otherwise.

### GetNuma25Ok

`func (o *GetVMConfig200ResponseData) GetNuma25Ok() (*string, bool)`

GetNuma25Ok returns a tuple with the Numa25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma25

`func (o *GetVMConfig200ResponseData) SetNuma25(v string)`

SetNuma25 sets Numa25 field to given value.

### HasNuma25

`func (o *GetVMConfig200ResponseData) HasNuma25() bool`

HasNuma25 returns a boolean if a field has been set.

### GetNuma26

`func (o *GetVMConfig200ResponseData) GetNuma26() string`

GetNuma26 returns the Numa26 field if non-nil, zero value otherwise.

### GetNuma26Ok

`func (o *GetVMConfig200ResponseData) GetNuma26Ok() (*string, bool)`

GetNuma26Ok returns a tuple with the Numa26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma26

`func (o *GetVMConfig200ResponseData) SetNuma26(v string)`

SetNuma26 sets Numa26 field to given value.

### HasNuma26

`func (o *GetVMConfig200ResponseData) HasNuma26() bool`

HasNuma26 returns a boolean if a field has been set.

### GetNuma27

`func (o *GetVMConfig200ResponseData) GetNuma27() string`

GetNuma27 returns the Numa27 field if non-nil, zero value otherwise.

### GetNuma27Ok

`func (o *GetVMConfig200ResponseData) GetNuma27Ok() (*string, bool)`

GetNuma27Ok returns a tuple with the Numa27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma27

`func (o *GetVMConfig200ResponseData) SetNuma27(v string)`

SetNuma27 sets Numa27 field to given value.

### HasNuma27

`func (o *GetVMConfig200ResponseData) HasNuma27() bool`

HasNuma27 returns a boolean if a field has been set.

### GetNuma28

`func (o *GetVMConfig200ResponseData) GetNuma28() string`

GetNuma28 returns the Numa28 field if non-nil, zero value otherwise.

### GetNuma28Ok

`func (o *GetVMConfig200ResponseData) GetNuma28Ok() (*string, bool)`

GetNuma28Ok returns a tuple with the Numa28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma28

`func (o *GetVMConfig200ResponseData) SetNuma28(v string)`

SetNuma28 sets Numa28 field to given value.

### HasNuma28

`func (o *GetVMConfig200ResponseData) HasNuma28() bool`

HasNuma28 returns a boolean if a field has been set.

### GetNuma29

`func (o *GetVMConfig200ResponseData) GetNuma29() string`

GetNuma29 returns the Numa29 field if non-nil, zero value otherwise.

### GetNuma29Ok

`func (o *GetVMConfig200ResponseData) GetNuma29Ok() (*string, bool)`

GetNuma29Ok returns a tuple with the Numa29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma29

`func (o *GetVMConfig200ResponseData) SetNuma29(v string)`

SetNuma29 sets Numa29 field to given value.

### HasNuma29

`func (o *GetVMConfig200ResponseData) HasNuma29() bool`

HasNuma29 returns a boolean if a field has been set.

### GetOnboot

`func (o *GetVMConfig200ResponseData) GetOnboot() int32`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *GetVMConfig200ResponseData) GetOnbootOk() (*int32, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *GetVMConfig200ResponseData) SetOnboot(v int32)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *GetVMConfig200ResponseData) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *GetVMConfig200ResponseData) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *GetVMConfig200ResponseData) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *GetVMConfig200ResponseData) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *GetVMConfig200ResponseData) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetParallel0

`func (o *GetVMConfig200ResponseData) GetParallel0() string`

GetParallel0 returns the Parallel0 field if non-nil, zero value otherwise.

### GetParallel0Ok

`func (o *GetVMConfig200ResponseData) GetParallel0Ok() (*string, bool)`

GetParallel0Ok returns a tuple with the Parallel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel0

`func (o *GetVMConfig200ResponseData) SetParallel0(v string)`

SetParallel0 sets Parallel0 field to given value.

### HasParallel0

`func (o *GetVMConfig200ResponseData) HasParallel0() bool`

HasParallel0 returns a boolean if a field has been set.

### GetParallel1

`func (o *GetVMConfig200ResponseData) GetParallel1() string`

GetParallel1 returns the Parallel1 field if non-nil, zero value otherwise.

### GetParallel1Ok

`func (o *GetVMConfig200ResponseData) GetParallel1Ok() (*string, bool)`

GetParallel1Ok returns a tuple with the Parallel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel1

`func (o *GetVMConfig200ResponseData) SetParallel1(v string)`

SetParallel1 sets Parallel1 field to given value.

### HasParallel1

`func (o *GetVMConfig200ResponseData) HasParallel1() bool`

HasParallel1 returns a boolean if a field has been set.

### GetParallel2

`func (o *GetVMConfig200ResponseData) GetParallel2() string`

GetParallel2 returns the Parallel2 field if non-nil, zero value otherwise.

### GetParallel2Ok

`func (o *GetVMConfig200ResponseData) GetParallel2Ok() (*string, bool)`

GetParallel2Ok returns a tuple with the Parallel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel2

`func (o *GetVMConfig200ResponseData) SetParallel2(v string)`

SetParallel2 sets Parallel2 field to given value.

### HasParallel2

`func (o *GetVMConfig200ResponseData) HasParallel2() bool`

HasParallel2 returns a boolean if a field has been set.

### GetParallel3

`func (o *GetVMConfig200ResponseData) GetParallel3() string`

GetParallel3 returns the Parallel3 field if non-nil, zero value otherwise.

### GetParallel3Ok

`func (o *GetVMConfig200ResponseData) GetParallel3Ok() (*string, bool)`

GetParallel3Ok returns a tuple with the Parallel3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel3

`func (o *GetVMConfig200ResponseData) SetParallel3(v string)`

SetParallel3 sets Parallel3 field to given value.

### HasParallel3

`func (o *GetVMConfig200ResponseData) HasParallel3() bool`

HasParallel3 returns a boolean if a field has been set.

### GetProtection

`func (o *GetVMConfig200ResponseData) GetProtection() int32`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *GetVMConfig200ResponseData) GetProtectionOk() (*int32, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *GetVMConfig200ResponseData) SetProtection(v int32)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *GetVMConfig200ResponseData) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetReboot

`func (o *GetVMConfig200ResponseData) GetReboot() int32`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *GetVMConfig200ResponseData) GetRebootOk() (*int32, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *GetVMConfig200ResponseData) SetReboot(v int32)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *GetVMConfig200ResponseData) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRng0

`func (o *GetVMConfig200ResponseData) GetRng0() string`

GetRng0 returns the Rng0 field if non-nil, zero value otherwise.

### GetRng0Ok

`func (o *GetVMConfig200ResponseData) GetRng0Ok() (*string, bool)`

GetRng0Ok returns a tuple with the Rng0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng0

`func (o *GetVMConfig200ResponseData) SetRng0(v string)`

SetRng0 sets Rng0 field to given value.

### HasRng0

`func (o *GetVMConfig200ResponseData) HasRng0() bool`

HasRng0 returns a boolean if a field has been set.

### GetSata0

`func (o *GetVMConfig200ResponseData) GetSata0() string`

GetSata0 returns the Sata0 field if non-nil, zero value otherwise.

### GetSata0Ok

`func (o *GetVMConfig200ResponseData) GetSata0Ok() (*string, bool)`

GetSata0Ok returns a tuple with the Sata0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata0

`func (o *GetVMConfig200ResponseData) SetSata0(v string)`

SetSata0 sets Sata0 field to given value.

### HasSata0

`func (o *GetVMConfig200ResponseData) HasSata0() bool`

HasSata0 returns a boolean if a field has been set.

### GetSata1

`func (o *GetVMConfig200ResponseData) GetSata1() string`

GetSata1 returns the Sata1 field if non-nil, zero value otherwise.

### GetSata1Ok

`func (o *GetVMConfig200ResponseData) GetSata1Ok() (*string, bool)`

GetSata1Ok returns a tuple with the Sata1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata1

`func (o *GetVMConfig200ResponseData) SetSata1(v string)`

SetSata1 sets Sata1 field to given value.

### HasSata1

`func (o *GetVMConfig200ResponseData) HasSata1() bool`

HasSata1 returns a boolean if a field has been set.

### GetSata2

`func (o *GetVMConfig200ResponseData) GetSata2() string`

GetSata2 returns the Sata2 field if non-nil, zero value otherwise.

### GetSata2Ok

`func (o *GetVMConfig200ResponseData) GetSata2Ok() (*string, bool)`

GetSata2Ok returns a tuple with the Sata2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata2

`func (o *GetVMConfig200ResponseData) SetSata2(v string)`

SetSata2 sets Sata2 field to given value.

### HasSata2

`func (o *GetVMConfig200ResponseData) HasSata2() bool`

HasSata2 returns a boolean if a field has been set.

### GetSata3

`func (o *GetVMConfig200ResponseData) GetSata3() string`

GetSata3 returns the Sata3 field if non-nil, zero value otherwise.

### GetSata3Ok

`func (o *GetVMConfig200ResponseData) GetSata3Ok() (*string, bool)`

GetSata3Ok returns a tuple with the Sata3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata3

`func (o *GetVMConfig200ResponseData) SetSata3(v string)`

SetSata3 sets Sata3 field to given value.

### HasSata3

`func (o *GetVMConfig200ResponseData) HasSata3() bool`

HasSata3 returns a boolean if a field has been set.

### GetSata4

`func (o *GetVMConfig200ResponseData) GetSata4() string`

GetSata4 returns the Sata4 field if non-nil, zero value otherwise.

### GetSata4Ok

`func (o *GetVMConfig200ResponseData) GetSata4Ok() (*string, bool)`

GetSata4Ok returns a tuple with the Sata4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata4

`func (o *GetVMConfig200ResponseData) SetSata4(v string)`

SetSata4 sets Sata4 field to given value.

### HasSata4

`func (o *GetVMConfig200ResponseData) HasSata4() bool`

HasSata4 returns a boolean if a field has been set.

### GetSata5

`func (o *GetVMConfig200ResponseData) GetSata5() string`

GetSata5 returns the Sata5 field if non-nil, zero value otherwise.

### GetSata5Ok

`func (o *GetVMConfig200ResponseData) GetSata5Ok() (*string, bool)`

GetSata5Ok returns a tuple with the Sata5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata5

`func (o *GetVMConfig200ResponseData) SetSata5(v string)`

SetSata5 sets Sata5 field to given value.

### HasSata5

`func (o *GetVMConfig200ResponseData) HasSata5() bool`

HasSata5 returns a boolean if a field has been set.

### GetScsi0

`func (o *GetVMConfig200ResponseData) GetScsi0() string`

GetScsi0 returns the Scsi0 field if non-nil, zero value otherwise.

### GetScsi0Ok

`func (o *GetVMConfig200ResponseData) GetScsi0Ok() (*string, bool)`

GetScsi0Ok returns a tuple with the Scsi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi0

`func (o *GetVMConfig200ResponseData) SetScsi0(v string)`

SetScsi0 sets Scsi0 field to given value.

### HasScsi0

`func (o *GetVMConfig200ResponseData) HasScsi0() bool`

HasScsi0 returns a boolean if a field has been set.

### GetScsi1

`func (o *GetVMConfig200ResponseData) GetScsi1() string`

GetScsi1 returns the Scsi1 field if non-nil, zero value otherwise.

### GetScsi1Ok

`func (o *GetVMConfig200ResponseData) GetScsi1Ok() (*string, bool)`

GetScsi1Ok returns a tuple with the Scsi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi1

`func (o *GetVMConfig200ResponseData) SetScsi1(v string)`

SetScsi1 sets Scsi1 field to given value.

### HasScsi1

`func (o *GetVMConfig200ResponseData) HasScsi1() bool`

HasScsi1 returns a boolean if a field has been set.

### GetScsi2

`func (o *GetVMConfig200ResponseData) GetScsi2() string`

GetScsi2 returns the Scsi2 field if non-nil, zero value otherwise.

### GetScsi2Ok

`func (o *GetVMConfig200ResponseData) GetScsi2Ok() (*string, bool)`

GetScsi2Ok returns a tuple with the Scsi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi2

`func (o *GetVMConfig200ResponseData) SetScsi2(v string)`

SetScsi2 sets Scsi2 field to given value.

### HasScsi2

`func (o *GetVMConfig200ResponseData) HasScsi2() bool`

HasScsi2 returns a boolean if a field has been set.

### GetScsi3

`func (o *GetVMConfig200ResponseData) GetScsi3() string`

GetScsi3 returns the Scsi3 field if non-nil, zero value otherwise.

### GetScsi3Ok

`func (o *GetVMConfig200ResponseData) GetScsi3Ok() (*string, bool)`

GetScsi3Ok returns a tuple with the Scsi3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi3

`func (o *GetVMConfig200ResponseData) SetScsi3(v string)`

SetScsi3 sets Scsi3 field to given value.

### HasScsi3

`func (o *GetVMConfig200ResponseData) HasScsi3() bool`

HasScsi3 returns a boolean if a field has been set.

### GetScsi4

`func (o *GetVMConfig200ResponseData) GetScsi4() string`

GetScsi4 returns the Scsi4 field if non-nil, zero value otherwise.

### GetScsi4Ok

`func (o *GetVMConfig200ResponseData) GetScsi4Ok() (*string, bool)`

GetScsi4Ok returns a tuple with the Scsi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi4

`func (o *GetVMConfig200ResponseData) SetScsi4(v string)`

SetScsi4 sets Scsi4 field to given value.

### HasScsi4

`func (o *GetVMConfig200ResponseData) HasScsi4() bool`

HasScsi4 returns a boolean if a field has been set.

### GetScsi5

`func (o *GetVMConfig200ResponseData) GetScsi5() string`

GetScsi5 returns the Scsi5 field if non-nil, zero value otherwise.

### GetScsi5Ok

`func (o *GetVMConfig200ResponseData) GetScsi5Ok() (*string, bool)`

GetScsi5Ok returns a tuple with the Scsi5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi5

`func (o *GetVMConfig200ResponseData) SetScsi5(v string)`

SetScsi5 sets Scsi5 field to given value.

### HasScsi5

`func (o *GetVMConfig200ResponseData) HasScsi5() bool`

HasScsi5 returns a boolean if a field has been set.

### GetScsi6

`func (o *GetVMConfig200ResponseData) GetScsi6() string`

GetScsi6 returns the Scsi6 field if non-nil, zero value otherwise.

### GetScsi6Ok

`func (o *GetVMConfig200ResponseData) GetScsi6Ok() (*string, bool)`

GetScsi6Ok returns a tuple with the Scsi6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi6

`func (o *GetVMConfig200ResponseData) SetScsi6(v string)`

SetScsi6 sets Scsi6 field to given value.

### HasScsi6

`func (o *GetVMConfig200ResponseData) HasScsi6() bool`

HasScsi6 returns a boolean if a field has been set.

### GetScsi7

`func (o *GetVMConfig200ResponseData) GetScsi7() string`

GetScsi7 returns the Scsi7 field if non-nil, zero value otherwise.

### GetScsi7Ok

`func (o *GetVMConfig200ResponseData) GetScsi7Ok() (*string, bool)`

GetScsi7Ok returns a tuple with the Scsi7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi7

`func (o *GetVMConfig200ResponseData) SetScsi7(v string)`

SetScsi7 sets Scsi7 field to given value.

### HasScsi7

`func (o *GetVMConfig200ResponseData) HasScsi7() bool`

HasScsi7 returns a boolean if a field has been set.

### GetScsi8

`func (o *GetVMConfig200ResponseData) GetScsi8() string`

GetScsi8 returns the Scsi8 field if non-nil, zero value otherwise.

### GetScsi8Ok

`func (o *GetVMConfig200ResponseData) GetScsi8Ok() (*string, bool)`

GetScsi8Ok returns a tuple with the Scsi8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi8

`func (o *GetVMConfig200ResponseData) SetScsi8(v string)`

SetScsi8 sets Scsi8 field to given value.

### HasScsi8

`func (o *GetVMConfig200ResponseData) HasScsi8() bool`

HasScsi8 returns a boolean if a field has been set.

### GetScsi9

`func (o *GetVMConfig200ResponseData) GetScsi9() string`

GetScsi9 returns the Scsi9 field if non-nil, zero value otherwise.

### GetScsi9Ok

`func (o *GetVMConfig200ResponseData) GetScsi9Ok() (*string, bool)`

GetScsi9Ok returns a tuple with the Scsi9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi9

`func (o *GetVMConfig200ResponseData) SetScsi9(v string)`

SetScsi9 sets Scsi9 field to given value.

### HasScsi9

`func (o *GetVMConfig200ResponseData) HasScsi9() bool`

HasScsi9 returns a boolean if a field has been set.

### GetScsi10

`func (o *GetVMConfig200ResponseData) GetScsi10() string`

GetScsi10 returns the Scsi10 field if non-nil, zero value otherwise.

### GetScsi10Ok

`func (o *GetVMConfig200ResponseData) GetScsi10Ok() (*string, bool)`

GetScsi10Ok returns a tuple with the Scsi10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi10

`func (o *GetVMConfig200ResponseData) SetScsi10(v string)`

SetScsi10 sets Scsi10 field to given value.

### HasScsi10

`func (o *GetVMConfig200ResponseData) HasScsi10() bool`

HasScsi10 returns a boolean if a field has been set.

### GetScsi11

`func (o *GetVMConfig200ResponseData) GetScsi11() string`

GetScsi11 returns the Scsi11 field if non-nil, zero value otherwise.

### GetScsi11Ok

`func (o *GetVMConfig200ResponseData) GetScsi11Ok() (*string, bool)`

GetScsi11Ok returns a tuple with the Scsi11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi11

`func (o *GetVMConfig200ResponseData) SetScsi11(v string)`

SetScsi11 sets Scsi11 field to given value.

### HasScsi11

`func (o *GetVMConfig200ResponseData) HasScsi11() bool`

HasScsi11 returns a boolean if a field has been set.

### GetScsi12

`func (o *GetVMConfig200ResponseData) GetScsi12() string`

GetScsi12 returns the Scsi12 field if non-nil, zero value otherwise.

### GetScsi12Ok

`func (o *GetVMConfig200ResponseData) GetScsi12Ok() (*string, bool)`

GetScsi12Ok returns a tuple with the Scsi12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi12

`func (o *GetVMConfig200ResponseData) SetScsi12(v string)`

SetScsi12 sets Scsi12 field to given value.

### HasScsi12

`func (o *GetVMConfig200ResponseData) HasScsi12() bool`

HasScsi12 returns a boolean if a field has been set.

### GetScsi13

`func (o *GetVMConfig200ResponseData) GetScsi13() string`

GetScsi13 returns the Scsi13 field if non-nil, zero value otherwise.

### GetScsi13Ok

`func (o *GetVMConfig200ResponseData) GetScsi13Ok() (*string, bool)`

GetScsi13Ok returns a tuple with the Scsi13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi13

`func (o *GetVMConfig200ResponseData) SetScsi13(v string)`

SetScsi13 sets Scsi13 field to given value.

### HasScsi13

`func (o *GetVMConfig200ResponseData) HasScsi13() bool`

HasScsi13 returns a boolean if a field has been set.

### GetScsi14

`func (o *GetVMConfig200ResponseData) GetScsi14() string`

GetScsi14 returns the Scsi14 field if non-nil, zero value otherwise.

### GetScsi14Ok

`func (o *GetVMConfig200ResponseData) GetScsi14Ok() (*string, bool)`

GetScsi14Ok returns a tuple with the Scsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi14

`func (o *GetVMConfig200ResponseData) SetScsi14(v string)`

SetScsi14 sets Scsi14 field to given value.

### HasScsi14

`func (o *GetVMConfig200ResponseData) HasScsi14() bool`

HasScsi14 returns a boolean if a field has been set.

### GetScsi15

`func (o *GetVMConfig200ResponseData) GetScsi15() string`

GetScsi15 returns the Scsi15 field if non-nil, zero value otherwise.

### GetScsi15Ok

`func (o *GetVMConfig200ResponseData) GetScsi15Ok() (*string, bool)`

GetScsi15Ok returns a tuple with the Scsi15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi15

`func (o *GetVMConfig200ResponseData) SetScsi15(v string)`

SetScsi15 sets Scsi15 field to given value.

### HasScsi15

`func (o *GetVMConfig200ResponseData) HasScsi15() bool`

HasScsi15 returns a boolean if a field has been set.

### GetScsi16

`func (o *GetVMConfig200ResponseData) GetScsi16() string`

GetScsi16 returns the Scsi16 field if non-nil, zero value otherwise.

### GetScsi16Ok

`func (o *GetVMConfig200ResponseData) GetScsi16Ok() (*string, bool)`

GetScsi16Ok returns a tuple with the Scsi16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi16

`func (o *GetVMConfig200ResponseData) SetScsi16(v string)`

SetScsi16 sets Scsi16 field to given value.

### HasScsi16

`func (o *GetVMConfig200ResponseData) HasScsi16() bool`

HasScsi16 returns a boolean if a field has been set.

### GetScsi17

`func (o *GetVMConfig200ResponseData) GetScsi17() string`

GetScsi17 returns the Scsi17 field if non-nil, zero value otherwise.

### GetScsi17Ok

`func (o *GetVMConfig200ResponseData) GetScsi17Ok() (*string, bool)`

GetScsi17Ok returns a tuple with the Scsi17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi17

`func (o *GetVMConfig200ResponseData) SetScsi17(v string)`

SetScsi17 sets Scsi17 field to given value.

### HasScsi17

`func (o *GetVMConfig200ResponseData) HasScsi17() bool`

HasScsi17 returns a boolean if a field has been set.

### GetScsi18

`func (o *GetVMConfig200ResponseData) GetScsi18() string`

GetScsi18 returns the Scsi18 field if non-nil, zero value otherwise.

### GetScsi18Ok

`func (o *GetVMConfig200ResponseData) GetScsi18Ok() (*string, bool)`

GetScsi18Ok returns a tuple with the Scsi18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi18

`func (o *GetVMConfig200ResponseData) SetScsi18(v string)`

SetScsi18 sets Scsi18 field to given value.

### HasScsi18

`func (o *GetVMConfig200ResponseData) HasScsi18() bool`

HasScsi18 returns a boolean if a field has been set.

### GetScsi19

`func (o *GetVMConfig200ResponseData) GetScsi19() string`

GetScsi19 returns the Scsi19 field if non-nil, zero value otherwise.

### GetScsi19Ok

`func (o *GetVMConfig200ResponseData) GetScsi19Ok() (*string, bool)`

GetScsi19Ok returns a tuple with the Scsi19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi19

`func (o *GetVMConfig200ResponseData) SetScsi19(v string)`

SetScsi19 sets Scsi19 field to given value.

### HasScsi19

`func (o *GetVMConfig200ResponseData) HasScsi19() bool`

HasScsi19 returns a boolean if a field has been set.

### GetScsi20

`func (o *GetVMConfig200ResponseData) GetScsi20() string`

GetScsi20 returns the Scsi20 field if non-nil, zero value otherwise.

### GetScsi20Ok

`func (o *GetVMConfig200ResponseData) GetScsi20Ok() (*string, bool)`

GetScsi20Ok returns a tuple with the Scsi20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi20

`func (o *GetVMConfig200ResponseData) SetScsi20(v string)`

SetScsi20 sets Scsi20 field to given value.

### HasScsi20

`func (o *GetVMConfig200ResponseData) HasScsi20() bool`

HasScsi20 returns a boolean if a field has been set.

### GetScsi21

`func (o *GetVMConfig200ResponseData) GetScsi21() string`

GetScsi21 returns the Scsi21 field if non-nil, zero value otherwise.

### GetScsi21Ok

`func (o *GetVMConfig200ResponseData) GetScsi21Ok() (*string, bool)`

GetScsi21Ok returns a tuple with the Scsi21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi21

`func (o *GetVMConfig200ResponseData) SetScsi21(v string)`

SetScsi21 sets Scsi21 field to given value.

### HasScsi21

`func (o *GetVMConfig200ResponseData) HasScsi21() bool`

HasScsi21 returns a boolean if a field has been set.

### GetScsi22

`func (o *GetVMConfig200ResponseData) GetScsi22() string`

GetScsi22 returns the Scsi22 field if non-nil, zero value otherwise.

### GetScsi22Ok

`func (o *GetVMConfig200ResponseData) GetScsi22Ok() (*string, bool)`

GetScsi22Ok returns a tuple with the Scsi22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi22

`func (o *GetVMConfig200ResponseData) SetScsi22(v string)`

SetScsi22 sets Scsi22 field to given value.

### HasScsi22

`func (o *GetVMConfig200ResponseData) HasScsi22() bool`

HasScsi22 returns a boolean if a field has been set.

### GetScsi23

`func (o *GetVMConfig200ResponseData) GetScsi23() string`

GetScsi23 returns the Scsi23 field if non-nil, zero value otherwise.

### GetScsi23Ok

`func (o *GetVMConfig200ResponseData) GetScsi23Ok() (*string, bool)`

GetScsi23Ok returns a tuple with the Scsi23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi23

`func (o *GetVMConfig200ResponseData) SetScsi23(v string)`

SetScsi23 sets Scsi23 field to given value.

### HasScsi23

`func (o *GetVMConfig200ResponseData) HasScsi23() bool`

HasScsi23 returns a boolean if a field has been set.

### GetScsi24

`func (o *GetVMConfig200ResponseData) GetScsi24() string`

GetScsi24 returns the Scsi24 field if non-nil, zero value otherwise.

### GetScsi24Ok

`func (o *GetVMConfig200ResponseData) GetScsi24Ok() (*string, bool)`

GetScsi24Ok returns a tuple with the Scsi24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi24

`func (o *GetVMConfig200ResponseData) SetScsi24(v string)`

SetScsi24 sets Scsi24 field to given value.

### HasScsi24

`func (o *GetVMConfig200ResponseData) HasScsi24() bool`

HasScsi24 returns a boolean if a field has been set.

### GetScsi25

`func (o *GetVMConfig200ResponseData) GetScsi25() string`

GetScsi25 returns the Scsi25 field if non-nil, zero value otherwise.

### GetScsi25Ok

`func (o *GetVMConfig200ResponseData) GetScsi25Ok() (*string, bool)`

GetScsi25Ok returns a tuple with the Scsi25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi25

`func (o *GetVMConfig200ResponseData) SetScsi25(v string)`

SetScsi25 sets Scsi25 field to given value.

### HasScsi25

`func (o *GetVMConfig200ResponseData) HasScsi25() bool`

HasScsi25 returns a boolean if a field has been set.

### GetScsi26

`func (o *GetVMConfig200ResponseData) GetScsi26() string`

GetScsi26 returns the Scsi26 field if non-nil, zero value otherwise.

### GetScsi26Ok

`func (o *GetVMConfig200ResponseData) GetScsi26Ok() (*string, bool)`

GetScsi26Ok returns a tuple with the Scsi26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi26

`func (o *GetVMConfig200ResponseData) SetScsi26(v string)`

SetScsi26 sets Scsi26 field to given value.

### HasScsi26

`func (o *GetVMConfig200ResponseData) HasScsi26() bool`

HasScsi26 returns a boolean if a field has been set.

### GetScsi27

`func (o *GetVMConfig200ResponseData) GetScsi27() string`

GetScsi27 returns the Scsi27 field if non-nil, zero value otherwise.

### GetScsi27Ok

`func (o *GetVMConfig200ResponseData) GetScsi27Ok() (*string, bool)`

GetScsi27Ok returns a tuple with the Scsi27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi27

`func (o *GetVMConfig200ResponseData) SetScsi27(v string)`

SetScsi27 sets Scsi27 field to given value.

### HasScsi27

`func (o *GetVMConfig200ResponseData) HasScsi27() bool`

HasScsi27 returns a boolean if a field has been set.

### GetScsi28

`func (o *GetVMConfig200ResponseData) GetScsi28() string`

GetScsi28 returns the Scsi28 field if non-nil, zero value otherwise.

### GetScsi28Ok

`func (o *GetVMConfig200ResponseData) GetScsi28Ok() (*string, bool)`

GetScsi28Ok returns a tuple with the Scsi28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi28

`func (o *GetVMConfig200ResponseData) SetScsi28(v string)`

SetScsi28 sets Scsi28 field to given value.

### HasScsi28

`func (o *GetVMConfig200ResponseData) HasScsi28() bool`

HasScsi28 returns a boolean if a field has been set.

### GetScsi29

`func (o *GetVMConfig200ResponseData) GetScsi29() string`

GetScsi29 returns the Scsi29 field if non-nil, zero value otherwise.

### GetScsi29Ok

`func (o *GetVMConfig200ResponseData) GetScsi29Ok() (*string, bool)`

GetScsi29Ok returns a tuple with the Scsi29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi29

`func (o *GetVMConfig200ResponseData) SetScsi29(v string)`

SetScsi29 sets Scsi29 field to given value.

### HasScsi29

`func (o *GetVMConfig200ResponseData) HasScsi29() bool`

HasScsi29 returns a boolean if a field has been set.

### GetScsihw

`func (o *GetVMConfig200ResponseData) GetScsihw() string`

GetScsihw returns the Scsihw field if non-nil, zero value otherwise.

### GetScsihwOk

`func (o *GetVMConfig200ResponseData) GetScsihwOk() (*string, bool)`

GetScsihwOk returns a tuple with the Scsihw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsihw

`func (o *GetVMConfig200ResponseData) SetScsihw(v string)`

SetScsihw sets Scsihw field to given value.

### HasScsihw

`func (o *GetVMConfig200ResponseData) HasScsihw() bool`

HasScsihw returns a boolean if a field has been set.

### GetSearchdomain

`func (o *GetVMConfig200ResponseData) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *GetVMConfig200ResponseData) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *GetVMConfig200ResponseData) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *GetVMConfig200ResponseData) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSerial0

`func (o *GetVMConfig200ResponseData) GetSerial0() string`

GetSerial0 returns the Serial0 field if non-nil, zero value otherwise.

### GetSerial0Ok

`func (o *GetVMConfig200ResponseData) GetSerial0Ok() (*string, bool)`

GetSerial0Ok returns a tuple with the Serial0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial0

`func (o *GetVMConfig200ResponseData) SetSerial0(v string)`

SetSerial0 sets Serial0 field to given value.

### HasSerial0

`func (o *GetVMConfig200ResponseData) HasSerial0() bool`

HasSerial0 returns a boolean if a field has been set.

### GetSerial1

`func (o *GetVMConfig200ResponseData) GetSerial1() string`

GetSerial1 returns the Serial1 field if non-nil, zero value otherwise.

### GetSerial1Ok

`func (o *GetVMConfig200ResponseData) GetSerial1Ok() (*string, bool)`

GetSerial1Ok returns a tuple with the Serial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial1

`func (o *GetVMConfig200ResponseData) SetSerial1(v string)`

SetSerial1 sets Serial1 field to given value.

### HasSerial1

`func (o *GetVMConfig200ResponseData) HasSerial1() bool`

HasSerial1 returns a boolean if a field has been set.

### GetSerial2

`func (o *GetVMConfig200ResponseData) GetSerial2() string`

GetSerial2 returns the Serial2 field if non-nil, zero value otherwise.

### GetSerial2Ok

`func (o *GetVMConfig200ResponseData) GetSerial2Ok() (*string, bool)`

GetSerial2Ok returns a tuple with the Serial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial2

`func (o *GetVMConfig200ResponseData) SetSerial2(v string)`

SetSerial2 sets Serial2 field to given value.

### HasSerial2

`func (o *GetVMConfig200ResponseData) HasSerial2() bool`

HasSerial2 returns a boolean if a field has been set.

### GetSerial3

`func (o *GetVMConfig200ResponseData) GetSerial3() string`

GetSerial3 returns the Serial3 field if non-nil, zero value otherwise.

### GetSerial3Ok

`func (o *GetVMConfig200ResponseData) GetSerial3Ok() (*string, bool)`

GetSerial3Ok returns a tuple with the Serial3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial3

`func (o *GetVMConfig200ResponseData) SetSerial3(v string)`

SetSerial3 sets Serial3 field to given value.

### HasSerial3

`func (o *GetVMConfig200ResponseData) HasSerial3() bool`

HasSerial3 returns a boolean if a field has been set.

### GetShares

`func (o *GetVMConfig200ResponseData) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetVMConfig200ResponseData) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetVMConfig200ResponseData) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *GetVMConfig200ResponseData) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSmbios1

`func (o *GetVMConfig200ResponseData) GetSmbios1() string`

GetSmbios1 returns the Smbios1 field if non-nil, zero value otherwise.

### GetSmbios1Ok

`func (o *GetVMConfig200ResponseData) GetSmbios1Ok() (*string, bool)`

GetSmbios1Ok returns a tuple with the Smbios1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbios1

`func (o *GetVMConfig200ResponseData) SetSmbios1(v string)`

SetSmbios1 sets Smbios1 field to given value.

### HasSmbios1

`func (o *GetVMConfig200ResponseData) HasSmbios1() bool`

HasSmbios1 returns a boolean if a field has been set.

### GetSmp

`func (o *GetVMConfig200ResponseData) GetSmp() int64`

GetSmp returns the Smp field if non-nil, zero value otherwise.

### GetSmpOk

`func (o *GetVMConfig200ResponseData) GetSmpOk() (*int64, bool)`

GetSmpOk returns a tuple with the Smp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmp

`func (o *GetVMConfig200ResponseData) SetSmp(v int64)`

SetSmp sets Smp field to given value.

### HasSmp

`func (o *GetVMConfig200ResponseData) HasSmp() bool`

HasSmp returns a boolean if a field has been set.

### GetSockets

`func (o *GetVMConfig200ResponseData) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *GetVMConfig200ResponseData) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *GetVMConfig200ResponseData) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *GetVMConfig200ResponseData) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpiceEnhancements

`func (o *GetVMConfig200ResponseData) GetSpiceEnhancements() string`

GetSpiceEnhancements returns the SpiceEnhancements field if non-nil, zero value otherwise.

### GetSpiceEnhancementsOk

`func (o *GetVMConfig200ResponseData) GetSpiceEnhancementsOk() (*string, bool)`

GetSpiceEnhancementsOk returns a tuple with the SpiceEnhancements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnhancements

`func (o *GetVMConfig200ResponseData) SetSpiceEnhancements(v string)`

SetSpiceEnhancements sets SpiceEnhancements field to given value.

### HasSpiceEnhancements

`func (o *GetVMConfig200ResponseData) HasSpiceEnhancements() bool`

HasSpiceEnhancements returns a boolean if a field has been set.

### GetSshkeys

`func (o *GetVMConfig200ResponseData) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *GetVMConfig200ResponseData) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *GetVMConfig200ResponseData) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *GetVMConfig200ResponseData) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.

### GetStartdate

`func (o *GetVMConfig200ResponseData) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *GetVMConfig200ResponseData) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *GetVMConfig200ResponseData) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *GetVMConfig200ResponseData) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetStartup

`func (o *GetVMConfig200ResponseData) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *GetVMConfig200ResponseData) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *GetVMConfig200ResponseData) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *GetVMConfig200ResponseData) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetTablet

`func (o *GetVMConfig200ResponseData) GetTablet() int32`

GetTablet returns the Tablet field if non-nil, zero value otherwise.

### GetTabletOk

`func (o *GetVMConfig200ResponseData) GetTabletOk() (*int32, bool)`

GetTabletOk returns a tuple with the Tablet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablet

`func (o *GetVMConfig200ResponseData) SetTablet(v int32)`

SetTablet sets Tablet field to given value.

### HasTablet

`func (o *GetVMConfig200ResponseData) HasTablet() bool`

HasTablet returns a boolean if a field has been set.

### GetTags

`func (o *GetVMConfig200ResponseData) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetVMConfig200ResponseData) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetVMConfig200ResponseData) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetVMConfig200ResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTdf

`func (o *GetVMConfig200ResponseData) GetTdf() int32`

GetTdf returns the Tdf field if non-nil, zero value otherwise.

### GetTdfOk

`func (o *GetVMConfig200ResponseData) GetTdfOk() (*int32, bool)`

GetTdfOk returns a tuple with the Tdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdf

`func (o *GetVMConfig200ResponseData) SetTdf(v int32)`

SetTdf sets Tdf field to given value.

### HasTdf

`func (o *GetVMConfig200ResponseData) HasTdf() bool`

HasTdf returns a boolean if a field has been set.

### GetTemplate

`func (o *GetVMConfig200ResponseData) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetVMConfig200ResponseData) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetVMConfig200ResponseData) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetVMConfig200ResponseData) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTpmstate0

`func (o *GetVMConfig200ResponseData) GetTpmstate0() string`

GetTpmstate0 returns the Tpmstate0 field if non-nil, zero value otherwise.

### GetTpmstate0Ok

`func (o *GetVMConfig200ResponseData) GetTpmstate0Ok() (*string, bool)`

GetTpmstate0Ok returns a tuple with the Tpmstate0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmstate0

`func (o *GetVMConfig200ResponseData) SetTpmstate0(v string)`

SetTpmstate0 sets Tpmstate0 field to given value.

### HasTpmstate0

`func (o *GetVMConfig200ResponseData) HasTpmstate0() bool`

HasTpmstate0 returns a boolean if a field has been set.

### GetUnused0

`func (o *GetVMConfig200ResponseData) GetUnused0() string`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *GetVMConfig200ResponseData) GetUnused0Ok() (*string, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *GetVMConfig200ResponseData) SetUnused0(v string)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *GetVMConfig200ResponseData) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *GetVMConfig200ResponseData) GetUnused1() string`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *GetVMConfig200ResponseData) GetUnused1Ok() (*string, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *GetVMConfig200ResponseData) SetUnused1(v string)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *GetVMConfig200ResponseData) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *GetVMConfig200ResponseData) GetUnused2() string`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *GetVMConfig200ResponseData) GetUnused2Ok() (*string, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *GetVMConfig200ResponseData) SetUnused2(v string)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *GetVMConfig200ResponseData) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *GetVMConfig200ResponseData) GetUnused3() string`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *GetVMConfig200ResponseData) GetUnused3Ok() (*string, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *GetVMConfig200ResponseData) SetUnused3(v string)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *GetVMConfig200ResponseData) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *GetVMConfig200ResponseData) GetUnused4() string`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *GetVMConfig200ResponseData) GetUnused4Ok() (*string, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *GetVMConfig200ResponseData) SetUnused4(v string)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *GetVMConfig200ResponseData) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *GetVMConfig200ResponseData) GetUnused5() string`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *GetVMConfig200ResponseData) GetUnused5Ok() (*string, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *GetVMConfig200ResponseData) SetUnused5(v string)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *GetVMConfig200ResponseData) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *GetVMConfig200ResponseData) GetUnused6() string`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *GetVMConfig200ResponseData) GetUnused6Ok() (*string, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *GetVMConfig200ResponseData) SetUnused6(v string)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *GetVMConfig200ResponseData) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *GetVMConfig200ResponseData) GetUnused7() string`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *GetVMConfig200ResponseData) GetUnused7Ok() (*string, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *GetVMConfig200ResponseData) SetUnused7(v string)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *GetVMConfig200ResponseData) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *GetVMConfig200ResponseData) GetUnused8() string`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *GetVMConfig200ResponseData) GetUnused8Ok() (*string, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *GetVMConfig200ResponseData) SetUnused8(v string)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *GetVMConfig200ResponseData) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *GetVMConfig200ResponseData) GetUnused9() string`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *GetVMConfig200ResponseData) GetUnused9Ok() (*string, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *GetVMConfig200ResponseData) SetUnused9(v string)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *GetVMConfig200ResponseData) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *GetVMConfig200ResponseData) GetUnused10() string`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *GetVMConfig200ResponseData) GetUnused10Ok() (*string, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *GetVMConfig200ResponseData) SetUnused10(v string)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *GetVMConfig200ResponseData) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *GetVMConfig200ResponseData) GetUnused11() string`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *GetVMConfig200ResponseData) GetUnused11Ok() (*string, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *GetVMConfig200ResponseData) SetUnused11(v string)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *GetVMConfig200ResponseData) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *GetVMConfig200ResponseData) GetUnused12() string`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *GetVMConfig200ResponseData) GetUnused12Ok() (*string, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *GetVMConfig200ResponseData) SetUnused12(v string)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *GetVMConfig200ResponseData) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *GetVMConfig200ResponseData) GetUnused13() string`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *GetVMConfig200ResponseData) GetUnused13Ok() (*string, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *GetVMConfig200ResponseData) SetUnused13(v string)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *GetVMConfig200ResponseData) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *GetVMConfig200ResponseData) GetUnused14() string`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *GetVMConfig200ResponseData) GetUnused14Ok() (*string, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *GetVMConfig200ResponseData) SetUnused14(v string)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *GetVMConfig200ResponseData) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *GetVMConfig200ResponseData) GetUnused15() string`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *GetVMConfig200ResponseData) GetUnused15Ok() (*string, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *GetVMConfig200ResponseData) SetUnused15(v string)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *GetVMConfig200ResponseData) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *GetVMConfig200ResponseData) GetUnused16() string`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *GetVMConfig200ResponseData) GetUnused16Ok() (*string, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *GetVMConfig200ResponseData) SetUnused16(v string)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *GetVMConfig200ResponseData) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *GetVMConfig200ResponseData) GetUnused17() string`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *GetVMConfig200ResponseData) GetUnused17Ok() (*string, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *GetVMConfig200ResponseData) SetUnused17(v string)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *GetVMConfig200ResponseData) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *GetVMConfig200ResponseData) GetUnused18() string`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *GetVMConfig200ResponseData) GetUnused18Ok() (*string, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *GetVMConfig200ResponseData) SetUnused18(v string)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *GetVMConfig200ResponseData) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *GetVMConfig200ResponseData) GetUnused19() string`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *GetVMConfig200ResponseData) GetUnused19Ok() (*string, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *GetVMConfig200ResponseData) SetUnused19(v string)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *GetVMConfig200ResponseData) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *GetVMConfig200ResponseData) GetUnused20() string`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *GetVMConfig200ResponseData) GetUnused20Ok() (*string, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *GetVMConfig200ResponseData) SetUnused20(v string)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *GetVMConfig200ResponseData) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *GetVMConfig200ResponseData) GetUnused21() string`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *GetVMConfig200ResponseData) GetUnused21Ok() (*string, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *GetVMConfig200ResponseData) SetUnused21(v string)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *GetVMConfig200ResponseData) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *GetVMConfig200ResponseData) GetUnused22() string`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *GetVMConfig200ResponseData) GetUnused22Ok() (*string, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *GetVMConfig200ResponseData) SetUnused22(v string)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *GetVMConfig200ResponseData) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *GetVMConfig200ResponseData) GetUnused23() string`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *GetVMConfig200ResponseData) GetUnused23Ok() (*string, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *GetVMConfig200ResponseData) SetUnused23(v string)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *GetVMConfig200ResponseData) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *GetVMConfig200ResponseData) GetUnused24() string`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *GetVMConfig200ResponseData) GetUnused24Ok() (*string, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *GetVMConfig200ResponseData) SetUnused24(v string)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *GetVMConfig200ResponseData) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *GetVMConfig200ResponseData) GetUnused25() string`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *GetVMConfig200ResponseData) GetUnused25Ok() (*string, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *GetVMConfig200ResponseData) SetUnused25(v string)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *GetVMConfig200ResponseData) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *GetVMConfig200ResponseData) GetUnused26() string`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *GetVMConfig200ResponseData) GetUnused26Ok() (*string, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *GetVMConfig200ResponseData) SetUnused26(v string)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *GetVMConfig200ResponseData) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *GetVMConfig200ResponseData) GetUnused27() string`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *GetVMConfig200ResponseData) GetUnused27Ok() (*string, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *GetVMConfig200ResponseData) SetUnused27(v string)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *GetVMConfig200ResponseData) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *GetVMConfig200ResponseData) GetUnused28() string`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *GetVMConfig200ResponseData) GetUnused28Ok() (*string, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *GetVMConfig200ResponseData) SetUnused28(v string)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *GetVMConfig200ResponseData) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *GetVMConfig200ResponseData) GetUnused29() string`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *GetVMConfig200ResponseData) GetUnused29Ok() (*string, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *GetVMConfig200ResponseData) SetUnused29(v string)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *GetVMConfig200ResponseData) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.

### GetUsb0

`func (o *GetVMConfig200ResponseData) GetUsb0() string`

GetUsb0 returns the Usb0 field if non-nil, zero value otherwise.

### GetUsb0Ok

`func (o *GetVMConfig200ResponseData) GetUsb0Ok() (*string, bool)`

GetUsb0Ok returns a tuple with the Usb0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb0

`func (o *GetVMConfig200ResponseData) SetUsb0(v string)`

SetUsb0 sets Usb0 field to given value.

### HasUsb0

`func (o *GetVMConfig200ResponseData) HasUsb0() bool`

HasUsb0 returns a boolean if a field has been set.

### GetUsb1

`func (o *GetVMConfig200ResponseData) GetUsb1() string`

GetUsb1 returns the Usb1 field if non-nil, zero value otherwise.

### GetUsb1Ok

`func (o *GetVMConfig200ResponseData) GetUsb1Ok() (*string, bool)`

GetUsb1Ok returns a tuple with the Usb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb1

`func (o *GetVMConfig200ResponseData) SetUsb1(v string)`

SetUsb1 sets Usb1 field to given value.

### HasUsb1

`func (o *GetVMConfig200ResponseData) HasUsb1() bool`

HasUsb1 returns a boolean if a field has been set.

### GetUsb2

`func (o *GetVMConfig200ResponseData) GetUsb2() string`

GetUsb2 returns the Usb2 field if non-nil, zero value otherwise.

### GetUsb2Ok

`func (o *GetVMConfig200ResponseData) GetUsb2Ok() (*string, bool)`

GetUsb2Ok returns a tuple with the Usb2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb2

`func (o *GetVMConfig200ResponseData) SetUsb2(v string)`

SetUsb2 sets Usb2 field to given value.

### HasUsb2

`func (o *GetVMConfig200ResponseData) HasUsb2() bool`

HasUsb2 returns a boolean if a field has been set.

### GetUsb3

`func (o *GetVMConfig200ResponseData) GetUsb3() string`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *GetVMConfig200ResponseData) GetUsb3Ok() (*string, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *GetVMConfig200ResponseData) SetUsb3(v string)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *GetVMConfig200ResponseData) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.

### GetVcpus

`func (o *GetVMConfig200ResponseData) GetVcpus() int64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *GetVMConfig200ResponseData) GetVcpusOk() (*int64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *GetVMConfig200ResponseData) SetVcpus(v int64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *GetVMConfig200ResponseData) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVga

`func (o *GetVMConfig200ResponseData) GetVga() string`

GetVga returns the Vga field if non-nil, zero value otherwise.

### GetVgaOk

`func (o *GetVMConfig200ResponseData) GetVgaOk() (*string, bool)`

GetVgaOk returns a tuple with the Vga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVga

`func (o *GetVMConfig200ResponseData) SetVga(v string)`

SetVga sets Vga field to given value.

### HasVga

`func (o *GetVMConfig200ResponseData) HasVga() bool`

HasVga returns a boolean if a field has been set.

### GetVirtio0

`func (o *GetVMConfig200ResponseData) GetVirtio0() string`

GetVirtio0 returns the Virtio0 field if non-nil, zero value otherwise.

### GetVirtio0Ok

`func (o *GetVMConfig200ResponseData) GetVirtio0Ok() (*string, bool)`

GetVirtio0Ok returns a tuple with the Virtio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio0

`func (o *GetVMConfig200ResponseData) SetVirtio0(v string)`

SetVirtio0 sets Virtio0 field to given value.

### HasVirtio0

`func (o *GetVMConfig200ResponseData) HasVirtio0() bool`

HasVirtio0 returns a boolean if a field has been set.

### GetVirtio1

`func (o *GetVMConfig200ResponseData) GetVirtio1() string`

GetVirtio1 returns the Virtio1 field if non-nil, zero value otherwise.

### GetVirtio1Ok

`func (o *GetVMConfig200ResponseData) GetVirtio1Ok() (*string, bool)`

GetVirtio1Ok returns a tuple with the Virtio1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio1

`func (o *GetVMConfig200ResponseData) SetVirtio1(v string)`

SetVirtio1 sets Virtio1 field to given value.

### HasVirtio1

`func (o *GetVMConfig200ResponseData) HasVirtio1() bool`

HasVirtio1 returns a boolean if a field has been set.

### GetVirtio2

`func (o *GetVMConfig200ResponseData) GetVirtio2() string`

GetVirtio2 returns the Virtio2 field if non-nil, zero value otherwise.

### GetVirtio2Ok

`func (o *GetVMConfig200ResponseData) GetVirtio2Ok() (*string, bool)`

GetVirtio2Ok returns a tuple with the Virtio2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio2

`func (o *GetVMConfig200ResponseData) SetVirtio2(v string)`

SetVirtio2 sets Virtio2 field to given value.

### HasVirtio2

`func (o *GetVMConfig200ResponseData) HasVirtio2() bool`

HasVirtio2 returns a boolean if a field has been set.

### GetVirtio3

`func (o *GetVMConfig200ResponseData) GetVirtio3() string`

GetVirtio3 returns the Virtio3 field if non-nil, zero value otherwise.

### GetVirtio3Ok

`func (o *GetVMConfig200ResponseData) GetVirtio3Ok() (*string, bool)`

GetVirtio3Ok returns a tuple with the Virtio3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio3

`func (o *GetVMConfig200ResponseData) SetVirtio3(v string)`

SetVirtio3 sets Virtio3 field to given value.

### HasVirtio3

`func (o *GetVMConfig200ResponseData) HasVirtio3() bool`

HasVirtio3 returns a boolean if a field has been set.

### GetVirtio4

`func (o *GetVMConfig200ResponseData) GetVirtio4() string`

GetVirtio4 returns the Virtio4 field if non-nil, zero value otherwise.

### GetVirtio4Ok

`func (o *GetVMConfig200ResponseData) GetVirtio4Ok() (*string, bool)`

GetVirtio4Ok returns a tuple with the Virtio4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio4

`func (o *GetVMConfig200ResponseData) SetVirtio4(v string)`

SetVirtio4 sets Virtio4 field to given value.

### HasVirtio4

`func (o *GetVMConfig200ResponseData) HasVirtio4() bool`

HasVirtio4 returns a boolean if a field has been set.

### GetVirtio5

`func (o *GetVMConfig200ResponseData) GetVirtio5() string`

GetVirtio5 returns the Virtio5 field if non-nil, zero value otherwise.

### GetVirtio5Ok

`func (o *GetVMConfig200ResponseData) GetVirtio5Ok() (*string, bool)`

GetVirtio5Ok returns a tuple with the Virtio5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio5

`func (o *GetVMConfig200ResponseData) SetVirtio5(v string)`

SetVirtio5 sets Virtio5 field to given value.

### HasVirtio5

`func (o *GetVMConfig200ResponseData) HasVirtio5() bool`

HasVirtio5 returns a boolean if a field has been set.

### GetVirtio6

`func (o *GetVMConfig200ResponseData) GetVirtio6() string`

GetVirtio6 returns the Virtio6 field if non-nil, zero value otherwise.

### GetVirtio6Ok

`func (o *GetVMConfig200ResponseData) GetVirtio6Ok() (*string, bool)`

GetVirtio6Ok returns a tuple with the Virtio6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio6

`func (o *GetVMConfig200ResponseData) SetVirtio6(v string)`

SetVirtio6 sets Virtio6 field to given value.

### HasVirtio6

`func (o *GetVMConfig200ResponseData) HasVirtio6() bool`

HasVirtio6 returns a boolean if a field has been set.

### GetVirtio7

`func (o *GetVMConfig200ResponseData) GetVirtio7() string`

GetVirtio7 returns the Virtio7 field if non-nil, zero value otherwise.

### GetVirtio7Ok

`func (o *GetVMConfig200ResponseData) GetVirtio7Ok() (*string, bool)`

GetVirtio7Ok returns a tuple with the Virtio7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio7

`func (o *GetVMConfig200ResponseData) SetVirtio7(v string)`

SetVirtio7 sets Virtio7 field to given value.

### HasVirtio7

`func (o *GetVMConfig200ResponseData) HasVirtio7() bool`

HasVirtio7 returns a boolean if a field has been set.

### GetVirtio8

`func (o *GetVMConfig200ResponseData) GetVirtio8() string`

GetVirtio8 returns the Virtio8 field if non-nil, zero value otherwise.

### GetVirtio8Ok

`func (o *GetVMConfig200ResponseData) GetVirtio8Ok() (*string, bool)`

GetVirtio8Ok returns a tuple with the Virtio8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio8

`func (o *GetVMConfig200ResponseData) SetVirtio8(v string)`

SetVirtio8 sets Virtio8 field to given value.

### HasVirtio8

`func (o *GetVMConfig200ResponseData) HasVirtio8() bool`

HasVirtio8 returns a boolean if a field has been set.

### GetVirtio9

`func (o *GetVMConfig200ResponseData) GetVirtio9() string`

GetVirtio9 returns the Virtio9 field if non-nil, zero value otherwise.

### GetVirtio9Ok

`func (o *GetVMConfig200ResponseData) GetVirtio9Ok() (*string, bool)`

GetVirtio9Ok returns a tuple with the Virtio9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio9

`func (o *GetVMConfig200ResponseData) SetVirtio9(v string)`

SetVirtio9 sets Virtio9 field to given value.

### HasVirtio9

`func (o *GetVMConfig200ResponseData) HasVirtio9() bool`

HasVirtio9 returns a boolean if a field has been set.

### GetVirtio10

`func (o *GetVMConfig200ResponseData) GetVirtio10() string`

GetVirtio10 returns the Virtio10 field if non-nil, zero value otherwise.

### GetVirtio10Ok

`func (o *GetVMConfig200ResponseData) GetVirtio10Ok() (*string, bool)`

GetVirtio10Ok returns a tuple with the Virtio10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio10

`func (o *GetVMConfig200ResponseData) SetVirtio10(v string)`

SetVirtio10 sets Virtio10 field to given value.

### HasVirtio10

`func (o *GetVMConfig200ResponseData) HasVirtio10() bool`

HasVirtio10 returns a boolean if a field has been set.

### GetVirtio11

`func (o *GetVMConfig200ResponseData) GetVirtio11() string`

GetVirtio11 returns the Virtio11 field if non-nil, zero value otherwise.

### GetVirtio11Ok

`func (o *GetVMConfig200ResponseData) GetVirtio11Ok() (*string, bool)`

GetVirtio11Ok returns a tuple with the Virtio11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio11

`func (o *GetVMConfig200ResponseData) SetVirtio11(v string)`

SetVirtio11 sets Virtio11 field to given value.

### HasVirtio11

`func (o *GetVMConfig200ResponseData) HasVirtio11() bool`

HasVirtio11 returns a boolean if a field has been set.

### GetVirtio12

`func (o *GetVMConfig200ResponseData) GetVirtio12() string`

GetVirtio12 returns the Virtio12 field if non-nil, zero value otherwise.

### GetVirtio12Ok

`func (o *GetVMConfig200ResponseData) GetVirtio12Ok() (*string, bool)`

GetVirtio12Ok returns a tuple with the Virtio12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio12

`func (o *GetVMConfig200ResponseData) SetVirtio12(v string)`

SetVirtio12 sets Virtio12 field to given value.

### HasVirtio12

`func (o *GetVMConfig200ResponseData) HasVirtio12() bool`

HasVirtio12 returns a boolean if a field has been set.

### GetVirtio13

`func (o *GetVMConfig200ResponseData) GetVirtio13() string`

GetVirtio13 returns the Virtio13 field if non-nil, zero value otherwise.

### GetVirtio13Ok

`func (o *GetVMConfig200ResponseData) GetVirtio13Ok() (*string, bool)`

GetVirtio13Ok returns a tuple with the Virtio13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio13

`func (o *GetVMConfig200ResponseData) SetVirtio13(v string)`

SetVirtio13 sets Virtio13 field to given value.

### HasVirtio13

`func (o *GetVMConfig200ResponseData) HasVirtio13() bool`

HasVirtio13 returns a boolean if a field has been set.

### GetVirtio14

`func (o *GetVMConfig200ResponseData) GetVirtio14() string`

GetVirtio14 returns the Virtio14 field if non-nil, zero value otherwise.

### GetVirtio14Ok

`func (o *GetVMConfig200ResponseData) GetVirtio14Ok() (*string, bool)`

GetVirtio14Ok returns a tuple with the Virtio14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio14

`func (o *GetVMConfig200ResponseData) SetVirtio14(v string)`

SetVirtio14 sets Virtio14 field to given value.

### HasVirtio14

`func (o *GetVMConfig200ResponseData) HasVirtio14() bool`

HasVirtio14 returns a boolean if a field has been set.

### GetVirtio15

`func (o *GetVMConfig200ResponseData) GetVirtio15() string`

GetVirtio15 returns the Virtio15 field if non-nil, zero value otherwise.

### GetVirtio15Ok

`func (o *GetVMConfig200ResponseData) GetVirtio15Ok() (*string, bool)`

GetVirtio15Ok returns a tuple with the Virtio15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio15

`func (o *GetVMConfig200ResponseData) SetVirtio15(v string)`

SetVirtio15 sets Virtio15 field to given value.

### HasVirtio15

`func (o *GetVMConfig200ResponseData) HasVirtio15() bool`

HasVirtio15 returns a boolean if a field has been set.

### GetVmgenid

`func (o *GetVMConfig200ResponseData) GetVmgenid() string`

GetVmgenid returns the Vmgenid field if non-nil, zero value otherwise.

### GetVmgenidOk

`func (o *GetVMConfig200ResponseData) GetVmgenidOk() (*string, bool)`

GetVmgenidOk returns a tuple with the Vmgenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmgenid

`func (o *GetVMConfig200ResponseData) SetVmgenid(v string)`

SetVmgenid sets Vmgenid field to given value.

### HasVmgenid

`func (o *GetVMConfig200ResponseData) HasVmgenid() bool`

HasVmgenid returns a boolean if a field has been set.

### GetVmstatestorage

`func (o *GetVMConfig200ResponseData) GetVmstatestorage() string`

GetVmstatestorage returns the Vmstatestorage field if non-nil, zero value otherwise.

### GetVmstatestorageOk

`func (o *GetVMConfig200ResponseData) GetVmstatestorageOk() (*string, bool)`

GetVmstatestorageOk returns a tuple with the Vmstatestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstatestorage

`func (o *GetVMConfig200ResponseData) SetVmstatestorage(v string)`

SetVmstatestorage sets Vmstatestorage field to given value.

### HasVmstatestorage

`func (o *GetVMConfig200ResponseData) HasVmstatestorage() bool`

HasVmstatestorage returns a boolean if a field has been set.

### GetWatchdog

`func (o *GetVMConfig200ResponseData) GetWatchdog() string`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *GetVMConfig200ResponseData) GetWatchdogOk() (*string, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *GetVMConfig200ResponseData) SetWatchdog(v string)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *GetVMConfig200ResponseData) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


