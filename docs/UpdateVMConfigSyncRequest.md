# UpdateVMConfigSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acpi** | Pointer to **bool** | Enable/disable ACPI. | [optional] 
**Affinity** | Pointer to **string** | List of host cores used to execute guest processes, for example: 0,5,8-11 | [optional] 
**Agent** | Pointer to [**GetVMConfig200ResponseDataAgent**](GetVMConfig200ResponseDataAgent.md) |  | [optional] 
**Arch** | Pointer to **string** | Virtual processor architecture. Defaults to the host. | [optional] 
**Args** | Pointer to **string** | Arbitrary arguments passed to kvm. | [optional] 
**Audio0** | Pointer to [**GetVMConfig200ResponseDataAudio0**](GetVMConfig200ResponseDataAudio0.md) |  | [optional] 
**Autostart** | Pointer to **bool** | Automatic restart after crash (currently ignored). | [optional] 
**Balloon** | Pointer to **int64** | Amount of target RAM for the VM in MiB. Using zero disables the ballon driver. | [optional] 
**Bios** | Pointer to **string** | Select BIOS implementation. | [optional] 
**Boot** | Pointer to [**GetVMConfig200ResponseDataBoot**](GetVMConfig200ResponseDataBoot.md) |  | [optional] 
**Bootdisk** | Pointer to **string** | Enable booting from specified disk. Deprecated: Use &#39;boot: order&#x3D;foo;bar&#39; instead. | [optional] 
**Cdrom** | Pointer to **string** | This is an alias for option -ide2 | [optional] 
**Cicustom** | Pointer to **string** | cloud-init: Specify custom files to replace the automatically generated ones at start. | [optional] 
**Cipassword** | Pointer to **string** | cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords. | [optional] 
**Citype** | Pointer to **string** | Specifies the cloud-init configuration format. The default depends on the configured operating system type (&#x60;ostype&#x60;. We use the &#x60;nocloud&#x60; format for Linux, and &#x60;configdrive2&#x60; for windows. | [optional] 
**Ciupgrade** | Pointer to **bool** | cloud-init: do an automatic package upgrade after the first boot. | [optional] 
**Ciuser** | Pointer to **string** | cloud-init: User name to change ssh keys and password for instead of the image&#39;s configured default user. | [optional] 
**Cores** | Pointer to **int64** | The number of cores per socket. | [optional] 
**Cpu** | Pointer to **string** | Emulated CPU type. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Delete** | Pointer to **string** | A list of settings you want to delete. | [optional] 
**Description** | Pointer to **string** | Description for the VM. Shown in the web-interface VM&#39;s summary. This is saved as comment inside the configuration file. | [optional] 
**Digest** | Pointer to **string** | Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications. | [optional] 
**Efidisk0** | Pointer to [**CreateVMRequestEfidisk0**](CreateVMRequestEfidisk0.md) |  | [optional] 
**Force** | Pointer to **bool** | Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called &#39;unused[n]&#39;, which contains the volume ID. Unlink of unused[n] always cause physical removal. | [optional] 
**Freeze** | Pointer to **bool** | Freeze CPU at startup (use &#39;c&#39; monitor command to start execution). | [optional] 
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
**Ide0** | Pointer to [**CreateVMRequestIde0**](CreateVMRequestIde0.md) |  | [optional] 
**Ide1** | Pointer to [**CreateVMRequestIde0**](CreateVMRequestIde0.md) |  | [optional] 
**Ide2** | Pointer to [**CreateVMRequestIde0**](CreateVMRequestIde0.md) |  | [optional] 
**Ide3** | Pointer to [**CreateVMRequestIde0**](CreateVMRequestIde0.md) |  | [optional] 
**Ipconfig0** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig1** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig2** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig3** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig4** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig5** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig6** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig7** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig8** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig9** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig10** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig11** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig12** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig13** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig14** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig15** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig16** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig17** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig18** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig19** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig20** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig21** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig22** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig23** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig24** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig25** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig26** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig27** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig28** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ipconfig29** | Pointer to [**GetVMConfig200ResponseDataIpconfig0**](GetVMConfig200ResponseDataIpconfig0.md) |  | [optional] 
**Ivshmem** | Pointer to [**GetVMConfig200ResponseDataIvshmem**](GetVMConfig200ResponseDataIvshmem.md) |  | [optional] 
**Keephugepages** | Pointer to **bool** | Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts. | [optional] 
**Keyboard** | Pointer to **string** | Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS. | [optional] 
**Kvm** | Pointer to **bool** | Enable/disable KVM hardware virtualization. | [optional] 
**Localtime** | Pointer to **bool** | Set the real time clock (RTC) to local time. This is enabled by default if the &#x60;ostype&#x60; indicates a Microsoft Windows OS. | [optional] 
**Lock** | Pointer to **string** | Lock/unlock the VM. | [optional] 
**Machine** | Pointer to **string** | Specifies the QEMU machine type. | [optional] 
**Memory** | Pointer to [**CreateVMRequestMemory**](CreateVMRequestMemory.md) |  | [optional] 
**MigrateDowntime** | Pointer to **float32** | Set maximum tolerated downtime (in seconds) for migrations. | [optional] 
**MigrateSpeed** | Pointer to **int64** | Set maximum speed (in MB/s) for migrations. Value 0 is no limit. | [optional] 
**Name** | Pointer to **string** | Set a name for the VM. Only used on the configuration web interface. | [optional] 
**Nameserver** | Pointer to **string** | cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Net0** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net1** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net2** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net3** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net4** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net5** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net6** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net7** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net8** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net9** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net10** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net11** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net12** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net13** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net14** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net15** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net16** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net17** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net18** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net19** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net20** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net21** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net22** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net23** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net24** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net25** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net26** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net27** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net28** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net29** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net30** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Net31** | Pointer to [**GetVMConfig200ResponseDataNet0**](GetVMConfig200ResponseDataNet0.md) |  | [optional] 
**Numa** | Pointer to **bool** | Enable/disable NUMA. | [optional] 
**Numa0** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa1** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa2** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa3** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa4** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa5** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa6** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa7** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa8** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa9** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa10** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa11** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa12** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa13** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa14** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa15** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa16** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa17** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa18** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa19** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa20** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa21** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa22** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa23** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa24** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa25** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa26** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa27** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa28** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Numa29** | Pointer to [**GetVMConfig200ResponseDataNuma0**](GetVMConfig200ResponseDataNuma0.md) |  | [optional] 
**Onboot** | Pointer to **bool** | Specifies whether a VM will be started during system bootup. | [optional] 
**Ostype** | Pointer to **string** | Specify guest operating system. | [optional] 
**Parallel0** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel1** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel2** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel3** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Protection** | Pointer to **bool** | Sets the protection flag of the VM. This will disable the remove VM and remove disk operations. | [optional] 
**Reboot** | Pointer to **bool** | Allow reboot. If set to &#39;0&#39; the VM exit on reboot. | [optional] 
**Revert** | Pointer to **string** | Revert a pending change. | [optional] 
**Rng0** | Pointer to [**GetVMConfig200ResponseDataRng0**](GetVMConfig200ResponseDataRng0.md) |  | [optional] 
**Sata0** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Sata1** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Sata2** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Sata3** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Sata4** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Sata5** | Pointer to [**CreateVMRequestSata0**](CreateVMRequestSata0.md) |  | [optional] 
**Scsi0** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi1** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi2** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi3** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi4** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi5** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi6** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi7** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi8** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi9** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi10** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi11** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi12** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi13** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi14** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi15** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi16** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi17** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi18** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi19** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi20** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi21** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi22** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi23** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi24** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi25** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi26** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi27** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi28** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsi29** | Pointer to [**CreateVMRequestScsi0**](CreateVMRequestScsi0.md) |  | [optional] 
**Scsihw** | Pointer to **string** | SCSI controller model | [optional] 
**Searchdomain** | Pointer to **string** | cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Serial0** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial1** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial2** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial3** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Shares** | Pointer to **int64** | Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd. | [optional] 
**Skiplock** | Pointer to **bool** | Ignore locks - only root is allowed to use this option. | [optional] 
**Smbios1** | Pointer to **string** | Specify SMBIOS type 1 fields. | [optional] 
**Smp** | Pointer to **int64** | The number of CPUs. Please use option -sockets instead. | [optional] 
**Sockets** | Pointer to **int64** | The number of CPU sockets. | [optional] 
**SpiceEnhancements** | Pointer to [**GetVMConfig200ResponseDataSpiceEnhancements**](GetVMConfig200ResponseDataSpiceEnhancements.md) |  | [optional] 
**Sshkeys** | Pointer to **string** | cloud-init: Setup public SSH keys (one key per line, OpenSSH format). | [optional] 
**Startdate** | Pointer to **string** | Set the initial date of the real time clock. Valid format for date are:&#39;now&#39; or &#39;2006-06-17T16:01:21&#39; or &#39;2006-06-17&#39;. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Tablet** | Pointer to **bool** | Enable/disable the USB tablet device. | [optional] 
**Tags** | Pointer to **string** | Tags of the VM. This is only meta information. | [optional] 
**Tdf** | Pointer to **bool** | Enable/disable time drift fix. | [optional] 
**Template** | Pointer to **bool** | Enable/disable Template. | [optional] 
**Tpmstate0** | Pointer to [**CreateVMRequestTpmstate0**](CreateVMRequestTpmstate0.md) |  | [optional] 
**Unused0** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused1** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused2** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused3** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused4** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused5** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused6** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused7** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused8** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused9** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused10** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused11** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused12** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused13** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused14** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused15** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused16** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused17** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused18** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused19** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused20** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused21** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused22** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused23** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused24** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused25** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused26** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused27** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused28** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Unused29** | Pointer to [**GetVMConfig200ResponseDataUnused0**](GetVMConfig200ResponseDataUnused0.md) |  | [optional] 
**Usb0** | Pointer to [**GetVMConfig200ResponseDataUsb0**](GetVMConfig200ResponseDataUsb0.md) |  | [optional] 
**Usb1** | Pointer to [**GetVMConfig200ResponseDataUsb0**](GetVMConfig200ResponseDataUsb0.md) |  | [optional] 
**Usb2** | Pointer to [**GetVMConfig200ResponseDataUsb0**](GetVMConfig200ResponseDataUsb0.md) |  | [optional] 
**Usb3** | Pointer to [**GetVMConfig200ResponseDataUsb0**](GetVMConfig200ResponseDataUsb0.md) |  | [optional] 
**Vcpus** | Pointer to **int64** | Number of hotplugged vcpus. | [optional] 
**Vga** | Pointer to [**GetVMConfig200ResponseDataVga**](GetVMConfig200ResponseDataVga.md) |  | [optional] 
**Virtio0** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio1** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio2** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio3** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio4** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio5** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio6** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio7** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio8** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio9** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio10** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio11** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio12** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio13** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio14** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Virtio15** | Pointer to [**CreateVMRequestVirtio0**](CreateVMRequestVirtio0.md) |  | [optional] 
**Vmgenid** | Pointer to **string** | Set VM Generation ID. Use &#39;1&#39; to autogenerate on create or update, pass &#39;0&#39; to disable explicitly. | [optional] 
**Vmstatestorage** | Pointer to **string** | Default storage for VM state volumes/files. | [optional] 
**Watchdog** | Pointer to [**GetVMConfig200ResponseDataWatchdog**](GetVMConfig200ResponseDataWatchdog.md) |  | [optional] 

## Methods

### NewUpdateVMConfigSyncRequest

`func NewUpdateVMConfigSyncRequest() *UpdateVMConfigSyncRequest`

NewUpdateVMConfigSyncRequest instantiates a new UpdateVMConfigSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMConfigSyncRequestWithDefaults

`func NewUpdateVMConfigSyncRequestWithDefaults() *UpdateVMConfigSyncRequest`

NewUpdateVMConfigSyncRequestWithDefaults instantiates a new UpdateVMConfigSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *UpdateVMConfigSyncRequest) GetAcpi() bool`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *UpdateVMConfigSyncRequest) GetAcpiOk() (*bool, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *UpdateVMConfigSyncRequest) SetAcpi(v bool)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *UpdateVMConfigSyncRequest) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetAffinity

`func (o *UpdateVMConfigSyncRequest) GetAffinity() string`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *UpdateVMConfigSyncRequest) GetAffinityOk() (*string, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *UpdateVMConfigSyncRequest) SetAffinity(v string)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *UpdateVMConfigSyncRequest) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAgent

`func (o *UpdateVMConfigSyncRequest) GetAgent() GetVMConfig200ResponseDataAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *UpdateVMConfigSyncRequest) GetAgentOk() (*GetVMConfig200ResponseDataAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *UpdateVMConfigSyncRequest) SetAgent(v GetVMConfig200ResponseDataAgent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *UpdateVMConfigSyncRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetArch

`func (o *UpdateVMConfigSyncRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UpdateVMConfigSyncRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UpdateVMConfigSyncRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *UpdateVMConfigSyncRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *UpdateVMConfigSyncRequest) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdateVMConfigSyncRequest) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdateVMConfigSyncRequest) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *UpdateVMConfigSyncRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAudio0

`func (o *UpdateVMConfigSyncRequest) GetAudio0() GetVMConfig200ResponseDataAudio0`

GetAudio0 returns the Audio0 field if non-nil, zero value otherwise.

### GetAudio0Ok

`func (o *UpdateVMConfigSyncRequest) GetAudio0Ok() (*GetVMConfig200ResponseDataAudio0, bool)`

GetAudio0Ok returns a tuple with the Audio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio0

`func (o *UpdateVMConfigSyncRequest) SetAudio0(v GetVMConfig200ResponseDataAudio0)`

SetAudio0 sets Audio0 field to given value.

### HasAudio0

`func (o *UpdateVMConfigSyncRequest) HasAudio0() bool`

HasAudio0 returns a boolean if a field has been set.

### GetAutostart

`func (o *UpdateVMConfigSyncRequest) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *UpdateVMConfigSyncRequest) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *UpdateVMConfigSyncRequest) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *UpdateVMConfigSyncRequest) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBalloon

`func (o *UpdateVMConfigSyncRequest) GetBalloon() int64`

GetBalloon returns the Balloon field if non-nil, zero value otherwise.

### GetBalloonOk

`func (o *UpdateVMConfigSyncRequest) GetBalloonOk() (*int64, bool)`

GetBalloonOk returns a tuple with the Balloon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloon

`func (o *UpdateVMConfigSyncRequest) SetBalloon(v int64)`

SetBalloon sets Balloon field to given value.

### HasBalloon

`func (o *UpdateVMConfigSyncRequest) HasBalloon() bool`

HasBalloon returns a boolean if a field has been set.

### GetBios

`func (o *UpdateVMConfigSyncRequest) GetBios() string`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *UpdateVMConfigSyncRequest) GetBiosOk() (*string, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *UpdateVMConfigSyncRequest) SetBios(v string)`

SetBios sets Bios field to given value.

### HasBios

`func (o *UpdateVMConfigSyncRequest) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetBoot

`func (o *UpdateVMConfigSyncRequest) GetBoot() GetVMConfig200ResponseDataBoot`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *UpdateVMConfigSyncRequest) GetBootOk() (*GetVMConfig200ResponseDataBoot, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *UpdateVMConfigSyncRequest) SetBoot(v GetVMConfig200ResponseDataBoot)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *UpdateVMConfigSyncRequest) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBootdisk

`func (o *UpdateVMConfigSyncRequest) GetBootdisk() string`

GetBootdisk returns the Bootdisk field if non-nil, zero value otherwise.

### GetBootdiskOk

`func (o *UpdateVMConfigSyncRequest) GetBootdiskOk() (*string, bool)`

GetBootdiskOk returns a tuple with the Bootdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootdisk

`func (o *UpdateVMConfigSyncRequest) SetBootdisk(v string)`

SetBootdisk sets Bootdisk field to given value.

### HasBootdisk

`func (o *UpdateVMConfigSyncRequest) HasBootdisk() bool`

HasBootdisk returns a boolean if a field has been set.

### GetCdrom

`func (o *UpdateVMConfigSyncRequest) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *UpdateVMConfigSyncRequest) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *UpdateVMConfigSyncRequest) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *UpdateVMConfigSyncRequest) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetCicustom

`func (o *UpdateVMConfigSyncRequest) GetCicustom() string`

GetCicustom returns the Cicustom field if non-nil, zero value otherwise.

### GetCicustomOk

`func (o *UpdateVMConfigSyncRequest) GetCicustomOk() (*string, bool)`

GetCicustomOk returns a tuple with the Cicustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicustom

`func (o *UpdateVMConfigSyncRequest) SetCicustom(v string)`

SetCicustom sets Cicustom field to given value.

### HasCicustom

`func (o *UpdateVMConfigSyncRequest) HasCicustom() bool`

HasCicustom returns a boolean if a field has been set.

### GetCipassword

`func (o *UpdateVMConfigSyncRequest) GetCipassword() string`

GetCipassword returns the Cipassword field if non-nil, zero value otherwise.

### GetCipasswordOk

`func (o *UpdateVMConfigSyncRequest) GetCipasswordOk() (*string, bool)`

GetCipasswordOk returns a tuple with the Cipassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipassword

`func (o *UpdateVMConfigSyncRequest) SetCipassword(v string)`

SetCipassword sets Cipassword field to given value.

### HasCipassword

`func (o *UpdateVMConfigSyncRequest) HasCipassword() bool`

HasCipassword returns a boolean if a field has been set.

### GetCitype

`func (o *UpdateVMConfigSyncRequest) GetCitype() string`

GetCitype returns the Citype field if non-nil, zero value otherwise.

### GetCitypeOk

`func (o *UpdateVMConfigSyncRequest) GetCitypeOk() (*string, bool)`

GetCitypeOk returns a tuple with the Citype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitype

`func (o *UpdateVMConfigSyncRequest) SetCitype(v string)`

SetCitype sets Citype field to given value.

### HasCitype

`func (o *UpdateVMConfigSyncRequest) HasCitype() bool`

HasCitype returns a boolean if a field has been set.

### GetCiupgrade

`func (o *UpdateVMConfigSyncRequest) GetCiupgrade() bool`

GetCiupgrade returns the Ciupgrade field if non-nil, zero value otherwise.

### GetCiupgradeOk

`func (o *UpdateVMConfigSyncRequest) GetCiupgradeOk() (*bool, bool)`

GetCiupgradeOk returns a tuple with the Ciupgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiupgrade

`func (o *UpdateVMConfigSyncRequest) SetCiupgrade(v bool)`

SetCiupgrade sets Ciupgrade field to given value.

### HasCiupgrade

`func (o *UpdateVMConfigSyncRequest) HasCiupgrade() bool`

HasCiupgrade returns a boolean if a field has been set.

### GetCiuser

`func (o *UpdateVMConfigSyncRequest) GetCiuser() string`

GetCiuser returns the Ciuser field if non-nil, zero value otherwise.

### GetCiuserOk

`func (o *UpdateVMConfigSyncRequest) GetCiuserOk() (*string, bool)`

GetCiuserOk returns a tuple with the Ciuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiuser

`func (o *UpdateVMConfigSyncRequest) SetCiuser(v string)`

SetCiuser sets Ciuser field to given value.

### HasCiuser

`func (o *UpdateVMConfigSyncRequest) HasCiuser() bool`

HasCiuser returns a boolean if a field has been set.

### GetCores

`func (o *UpdateVMConfigSyncRequest) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *UpdateVMConfigSyncRequest) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *UpdateVMConfigSyncRequest) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *UpdateVMConfigSyncRequest) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpu

`func (o *UpdateVMConfigSyncRequest) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UpdateVMConfigSyncRequest) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UpdateVMConfigSyncRequest) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UpdateVMConfigSyncRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpulimit

`func (o *UpdateVMConfigSyncRequest) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *UpdateVMConfigSyncRequest) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *UpdateVMConfigSyncRequest) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *UpdateVMConfigSyncRequest) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *UpdateVMConfigSyncRequest) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *UpdateVMConfigSyncRequest) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *UpdateVMConfigSyncRequest) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *UpdateVMConfigSyncRequest) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDelete

`func (o *UpdateVMConfigSyncRequest) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateVMConfigSyncRequest) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateVMConfigSyncRequest) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateVMConfigSyncRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateVMConfigSyncRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVMConfigSyncRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVMConfigSyncRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVMConfigSyncRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *UpdateVMConfigSyncRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UpdateVMConfigSyncRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UpdateVMConfigSyncRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *UpdateVMConfigSyncRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEfidisk0

`func (o *UpdateVMConfigSyncRequest) GetEfidisk0() CreateVMRequestEfidisk0`

GetEfidisk0 returns the Efidisk0 field if non-nil, zero value otherwise.

### GetEfidisk0Ok

`func (o *UpdateVMConfigSyncRequest) GetEfidisk0Ok() (*CreateVMRequestEfidisk0, bool)`

GetEfidisk0Ok returns a tuple with the Efidisk0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfidisk0

`func (o *UpdateVMConfigSyncRequest) SetEfidisk0(v CreateVMRequestEfidisk0)`

SetEfidisk0 sets Efidisk0 field to given value.

### HasEfidisk0

`func (o *UpdateVMConfigSyncRequest) HasEfidisk0() bool`

HasEfidisk0 returns a boolean if a field has been set.

### GetForce

`func (o *UpdateVMConfigSyncRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpdateVMConfigSyncRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpdateVMConfigSyncRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpdateVMConfigSyncRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetFreeze

`func (o *UpdateVMConfigSyncRequest) GetFreeze() bool`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *UpdateVMConfigSyncRequest) GetFreezeOk() (*bool, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *UpdateVMConfigSyncRequest) SetFreeze(v bool)`

SetFreeze sets Freeze field to given value.

### HasFreeze

`func (o *UpdateVMConfigSyncRequest) HasFreeze() bool`

HasFreeze returns a boolean if a field has been set.

### GetHookscript

`func (o *UpdateVMConfigSyncRequest) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *UpdateVMConfigSyncRequest) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *UpdateVMConfigSyncRequest) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *UpdateVMConfigSyncRequest) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostpci0

`func (o *UpdateVMConfigSyncRequest) GetHostpci0() string`

GetHostpci0 returns the Hostpci0 field if non-nil, zero value otherwise.

### GetHostpci0Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci0Ok() (*string, bool)`

GetHostpci0Ok returns a tuple with the Hostpci0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci0

`func (o *UpdateVMConfigSyncRequest) SetHostpci0(v string)`

SetHostpci0 sets Hostpci0 field to given value.

### HasHostpci0

`func (o *UpdateVMConfigSyncRequest) HasHostpci0() bool`

HasHostpci0 returns a boolean if a field has been set.

### GetHostpci1

`func (o *UpdateVMConfigSyncRequest) GetHostpci1() string`

GetHostpci1 returns the Hostpci1 field if non-nil, zero value otherwise.

### GetHostpci1Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci1Ok() (*string, bool)`

GetHostpci1Ok returns a tuple with the Hostpci1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci1

`func (o *UpdateVMConfigSyncRequest) SetHostpci1(v string)`

SetHostpci1 sets Hostpci1 field to given value.

### HasHostpci1

`func (o *UpdateVMConfigSyncRequest) HasHostpci1() bool`

HasHostpci1 returns a boolean if a field has been set.

### GetHostpci2

`func (o *UpdateVMConfigSyncRequest) GetHostpci2() string`

GetHostpci2 returns the Hostpci2 field if non-nil, zero value otherwise.

### GetHostpci2Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci2Ok() (*string, bool)`

GetHostpci2Ok returns a tuple with the Hostpci2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci2

`func (o *UpdateVMConfigSyncRequest) SetHostpci2(v string)`

SetHostpci2 sets Hostpci2 field to given value.

### HasHostpci2

`func (o *UpdateVMConfigSyncRequest) HasHostpci2() bool`

HasHostpci2 returns a boolean if a field has been set.

### GetHostpci3

`func (o *UpdateVMConfigSyncRequest) GetHostpci3() string`

GetHostpci3 returns the Hostpci3 field if non-nil, zero value otherwise.

### GetHostpci3Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci3Ok() (*string, bool)`

GetHostpci3Ok returns a tuple with the Hostpci3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci3

`func (o *UpdateVMConfigSyncRequest) SetHostpci3(v string)`

SetHostpci3 sets Hostpci3 field to given value.

### HasHostpci3

`func (o *UpdateVMConfigSyncRequest) HasHostpci3() bool`

HasHostpci3 returns a boolean if a field has been set.

### GetHostpci4

`func (o *UpdateVMConfigSyncRequest) GetHostpci4() string`

GetHostpci4 returns the Hostpci4 field if non-nil, zero value otherwise.

### GetHostpci4Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci4Ok() (*string, bool)`

GetHostpci4Ok returns a tuple with the Hostpci4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci4

`func (o *UpdateVMConfigSyncRequest) SetHostpci4(v string)`

SetHostpci4 sets Hostpci4 field to given value.

### HasHostpci4

`func (o *UpdateVMConfigSyncRequest) HasHostpci4() bool`

HasHostpci4 returns a boolean if a field has been set.

### GetHostpci5

`func (o *UpdateVMConfigSyncRequest) GetHostpci5() string`

GetHostpci5 returns the Hostpci5 field if non-nil, zero value otherwise.

### GetHostpci5Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci5Ok() (*string, bool)`

GetHostpci5Ok returns a tuple with the Hostpci5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci5

`func (o *UpdateVMConfigSyncRequest) SetHostpci5(v string)`

SetHostpci5 sets Hostpci5 field to given value.

### HasHostpci5

`func (o *UpdateVMConfigSyncRequest) HasHostpci5() bool`

HasHostpci5 returns a boolean if a field has been set.

### GetHostpci6

`func (o *UpdateVMConfigSyncRequest) GetHostpci6() string`

GetHostpci6 returns the Hostpci6 field if non-nil, zero value otherwise.

### GetHostpci6Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci6Ok() (*string, bool)`

GetHostpci6Ok returns a tuple with the Hostpci6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci6

`func (o *UpdateVMConfigSyncRequest) SetHostpci6(v string)`

SetHostpci6 sets Hostpci6 field to given value.

### HasHostpci6

`func (o *UpdateVMConfigSyncRequest) HasHostpci6() bool`

HasHostpci6 returns a boolean if a field has been set.

### GetHostpci7

`func (o *UpdateVMConfigSyncRequest) GetHostpci7() string`

GetHostpci7 returns the Hostpci7 field if non-nil, zero value otherwise.

### GetHostpci7Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci7Ok() (*string, bool)`

GetHostpci7Ok returns a tuple with the Hostpci7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci7

`func (o *UpdateVMConfigSyncRequest) SetHostpci7(v string)`

SetHostpci7 sets Hostpci7 field to given value.

### HasHostpci7

`func (o *UpdateVMConfigSyncRequest) HasHostpci7() bool`

HasHostpci7 returns a boolean if a field has been set.

### GetHostpci8

`func (o *UpdateVMConfigSyncRequest) GetHostpci8() string`

GetHostpci8 returns the Hostpci8 field if non-nil, zero value otherwise.

### GetHostpci8Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci8Ok() (*string, bool)`

GetHostpci8Ok returns a tuple with the Hostpci8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci8

`func (o *UpdateVMConfigSyncRequest) SetHostpci8(v string)`

SetHostpci8 sets Hostpci8 field to given value.

### HasHostpci8

`func (o *UpdateVMConfigSyncRequest) HasHostpci8() bool`

HasHostpci8 returns a boolean if a field has been set.

### GetHostpci9

`func (o *UpdateVMConfigSyncRequest) GetHostpci9() string`

GetHostpci9 returns the Hostpci9 field if non-nil, zero value otherwise.

### GetHostpci9Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci9Ok() (*string, bool)`

GetHostpci9Ok returns a tuple with the Hostpci9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci9

`func (o *UpdateVMConfigSyncRequest) SetHostpci9(v string)`

SetHostpci9 sets Hostpci9 field to given value.

### HasHostpci9

`func (o *UpdateVMConfigSyncRequest) HasHostpci9() bool`

HasHostpci9 returns a boolean if a field has been set.

### GetHostpci10

`func (o *UpdateVMConfigSyncRequest) GetHostpci10() string`

GetHostpci10 returns the Hostpci10 field if non-nil, zero value otherwise.

### GetHostpci10Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci10Ok() (*string, bool)`

GetHostpci10Ok returns a tuple with the Hostpci10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci10

`func (o *UpdateVMConfigSyncRequest) SetHostpci10(v string)`

SetHostpci10 sets Hostpci10 field to given value.

### HasHostpci10

`func (o *UpdateVMConfigSyncRequest) HasHostpci10() bool`

HasHostpci10 returns a boolean if a field has been set.

### GetHostpci11

`func (o *UpdateVMConfigSyncRequest) GetHostpci11() string`

GetHostpci11 returns the Hostpci11 field if non-nil, zero value otherwise.

### GetHostpci11Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci11Ok() (*string, bool)`

GetHostpci11Ok returns a tuple with the Hostpci11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci11

`func (o *UpdateVMConfigSyncRequest) SetHostpci11(v string)`

SetHostpci11 sets Hostpci11 field to given value.

### HasHostpci11

`func (o *UpdateVMConfigSyncRequest) HasHostpci11() bool`

HasHostpci11 returns a boolean if a field has been set.

### GetHostpci12

`func (o *UpdateVMConfigSyncRequest) GetHostpci12() string`

GetHostpci12 returns the Hostpci12 field if non-nil, zero value otherwise.

### GetHostpci12Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci12Ok() (*string, bool)`

GetHostpci12Ok returns a tuple with the Hostpci12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci12

`func (o *UpdateVMConfigSyncRequest) SetHostpci12(v string)`

SetHostpci12 sets Hostpci12 field to given value.

### HasHostpci12

`func (o *UpdateVMConfigSyncRequest) HasHostpci12() bool`

HasHostpci12 returns a boolean if a field has been set.

### GetHostpci13

`func (o *UpdateVMConfigSyncRequest) GetHostpci13() string`

GetHostpci13 returns the Hostpci13 field if non-nil, zero value otherwise.

### GetHostpci13Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci13Ok() (*string, bool)`

GetHostpci13Ok returns a tuple with the Hostpci13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci13

`func (o *UpdateVMConfigSyncRequest) SetHostpci13(v string)`

SetHostpci13 sets Hostpci13 field to given value.

### HasHostpci13

`func (o *UpdateVMConfigSyncRequest) HasHostpci13() bool`

HasHostpci13 returns a boolean if a field has been set.

### GetHostpci14

`func (o *UpdateVMConfigSyncRequest) GetHostpci14() string`

GetHostpci14 returns the Hostpci14 field if non-nil, zero value otherwise.

### GetHostpci14Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci14Ok() (*string, bool)`

GetHostpci14Ok returns a tuple with the Hostpci14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci14

`func (o *UpdateVMConfigSyncRequest) SetHostpci14(v string)`

SetHostpci14 sets Hostpci14 field to given value.

### HasHostpci14

`func (o *UpdateVMConfigSyncRequest) HasHostpci14() bool`

HasHostpci14 returns a boolean if a field has been set.

### GetHostpci15

`func (o *UpdateVMConfigSyncRequest) GetHostpci15() string`

GetHostpci15 returns the Hostpci15 field if non-nil, zero value otherwise.

### GetHostpci15Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci15Ok() (*string, bool)`

GetHostpci15Ok returns a tuple with the Hostpci15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci15

`func (o *UpdateVMConfigSyncRequest) SetHostpci15(v string)`

SetHostpci15 sets Hostpci15 field to given value.

### HasHostpci15

`func (o *UpdateVMConfigSyncRequest) HasHostpci15() bool`

HasHostpci15 returns a boolean if a field has been set.

### GetHostpci16

`func (o *UpdateVMConfigSyncRequest) GetHostpci16() string`

GetHostpci16 returns the Hostpci16 field if non-nil, zero value otherwise.

### GetHostpci16Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci16Ok() (*string, bool)`

GetHostpci16Ok returns a tuple with the Hostpci16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci16

`func (o *UpdateVMConfigSyncRequest) SetHostpci16(v string)`

SetHostpci16 sets Hostpci16 field to given value.

### HasHostpci16

`func (o *UpdateVMConfigSyncRequest) HasHostpci16() bool`

HasHostpci16 returns a boolean if a field has been set.

### GetHostpci17

`func (o *UpdateVMConfigSyncRequest) GetHostpci17() string`

GetHostpci17 returns the Hostpci17 field if non-nil, zero value otherwise.

### GetHostpci17Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci17Ok() (*string, bool)`

GetHostpci17Ok returns a tuple with the Hostpci17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci17

`func (o *UpdateVMConfigSyncRequest) SetHostpci17(v string)`

SetHostpci17 sets Hostpci17 field to given value.

### HasHostpci17

`func (o *UpdateVMConfigSyncRequest) HasHostpci17() bool`

HasHostpci17 returns a boolean if a field has been set.

### GetHostpci18

`func (o *UpdateVMConfigSyncRequest) GetHostpci18() string`

GetHostpci18 returns the Hostpci18 field if non-nil, zero value otherwise.

### GetHostpci18Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci18Ok() (*string, bool)`

GetHostpci18Ok returns a tuple with the Hostpci18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci18

`func (o *UpdateVMConfigSyncRequest) SetHostpci18(v string)`

SetHostpci18 sets Hostpci18 field to given value.

### HasHostpci18

`func (o *UpdateVMConfigSyncRequest) HasHostpci18() bool`

HasHostpci18 returns a boolean if a field has been set.

### GetHostpci19

`func (o *UpdateVMConfigSyncRequest) GetHostpci19() string`

GetHostpci19 returns the Hostpci19 field if non-nil, zero value otherwise.

### GetHostpci19Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci19Ok() (*string, bool)`

GetHostpci19Ok returns a tuple with the Hostpci19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci19

`func (o *UpdateVMConfigSyncRequest) SetHostpci19(v string)`

SetHostpci19 sets Hostpci19 field to given value.

### HasHostpci19

`func (o *UpdateVMConfigSyncRequest) HasHostpci19() bool`

HasHostpci19 returns a boolean if a field has been set.

### GetHostpci20

`func (o *UpdateVMConfigSyncRequest) GetHostpci20() string`

GetHostpci20 returns the Hostpci20 field if non-nil, zero value otherwise.

### GetHostpci20Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci20Ok() (*string, bool)`

GetHostpci20Ok returns a tuple with the Hostpci20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci20

`func (o *UpdateVMConfigSyncRequest) SetHostpci20(v string)`

SetHostpci20 sets Hostpci20 field to given value.

### HasHostpci20

`func (o *UpdateVMConfigSyncRequest) HasHostpci20() bool`

HasHostpci20 returns a boolean if a field has been set.

### GetHostpci21

`func (o *UpdateVMConfigSyncRequest) GetHostpci21() string`

GetHostpci21 returns the Hostpci21 field if non-nil, zero value otherwise.

### GetHostpci21Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci21Ok() (*string, bool)`

GetHostpci21Ok returns a tuple with the Hostpci21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci21

`func (o *UpdateVMConfigSyncRequest) SetHostpci21(v string)`

SetHostpci21 sets Hostpci21 field to given value.

### HasHostpci21

`func (o *UpdateVMConfigSyncRequest) HasHostpci21() bool`

HasHostpci21 returns a boolean if a field has been set.

### GetHostpci22

`func (o *UpdateVMConfigSyncRequest) GetHostpci22() string`

GetHostpci22 returns the Hostpci22 field if non-nil, zero value otherwise.

### GetHostpci22Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci22Ok() (*string, bool)`

GetHostpci22Ok returns a tuple with the Hostpci22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci22

`func (o *UpdateVMConfigSyncRequest) SetHostpci22(v string)`

SetHostpci22 sets Hostpci22 field to given value.

### HasHostpci22

`func (o *UpdateVMConfigSyncRequest) HasHostpci22() bool`

HasHostpci22 returns a boolean if a field has been set.

### GetHostpci23

`func (o *UpdateVMConfigSyncRequest) GetHostpci23() string`

GetHostpci23 returns the Hostpci23 field if non-nil, zero value otherwise.

### GetHostpci23Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci23Ok() (*string, bool)`

GetHostpci23Ok returns a tuple with the Hostpci23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci23

`func (o *UpdateVMConfigSyncRequest) SetHostpci23(v string)`

SetHostpci23 sets Hostpci23 field to given value.

### HasHostpci23

`func (o *UpdateVMConfigSyncRequest) HasHostpci23() bool`

HasHostpci23 returns a boolean if a field has been set.

### GetHostpci24

`func (o *UpdateVMConfigSyncRequest) GetHostpci24() string`

GetHostpci24 returns the Hostpci24 field if non-nil, zero value otherwise.

### GetHostpci24Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci24Ok() (*string, bool)`

GetHostpci24Ok returns a tuple with the Hostpci24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci24

`func (o *UpdateVMConfigSyncRequest) SetHostpci24(v string)`

SetHostpci24 sets Hostpci24 field to given value.

### HasHostpci24

`func (o *UpdateVMConfigSyncRequest) HasHostpci24() bool`

HasHostpci24 returns a boolean if a field has been set.

### GetHostpci25

`func (o *UpdateVMConfigSyncRequest) GetHostpci25() string`

GetHostpci25 returns the Hostpci25 field if non-nil, zero value otherwise.

### GetHostpci25Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci25Ok() (*string, bool)`

GetHostpci25Ok returns a tuple with the Hostpci25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci25

`func (o *UpdateVMConfigSyncRequest) SetHostpci25(v string)`

SetHostpci25 sets Hostpci25 field to given value.

### HasHostpci25

`func (o *UpdateVMConfigSyncRequest) HasHostpci25() bool`

HasHostpci25 returns a boolean if a field has been set.

### GetHostpci26

`func (o *UpdateVMConfigSyncRequest) GetHostpci26() string`

GetHostpci26 returns the Hostpci26 field if non-nil, zero value otherwise.

### GetHostpci26Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci26Ok() (*string, bool)`

GetHostpci26Ok returns a tuple with the Hostpci26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci26

`func (o *UpdateVMConfigSyncRequest) SetHostpci26(v string)`

SetHostpci26 sets Hostpci26 field to given value.

### HasHostpci26

`func (o *UpdateVMConfigSyncRequest) HasHostpci26() bool`

HasHostpci26 returns a boolean if a field has been set.

### GetHostpci27

`func (o *UpdateVMConfigSyncRequest) GetHostpci27() string`

GetHostpci27 returns the Hostpci27 field if non-nil, zero value otherwise.

### GetHostpci27Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci27Ok() (*string, bool)`

GetHostpci27Ok returns a tuple with the Hostpci27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci27

`func (o *UpdateVMConfigSyncRequest) SetHostpci27(v string)`

SetHostpci27 sets Hostpci27 field to given value.

### HasHostpci27

`func (o *UpdateVMConfigSyncRequest) HasHostpci27() bool`

HasHostpci27 returns a boolean if a field has been set.

### GetHostpci28

`func (o *UpdateVMConfigSyncRequest) GetHostpci28() string`

GetHostpci28 returns the Hostpci28 field if non-nil, zero value otherwise.

### GetHostpci28Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci28Ok() (*string, bool)`

GetHostpci28Ok returns a tuple with the Hostpci28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci28

`func (o *UpdateVMConfigSyncRequest) SetHostpci28(v string)`

SetHostpci28 sets Hostpci28 field to given value.

### HasHostpci28

`func (o *UpdateVMConfigSyncRequest) HasHostpci28() bool`

HasHostpci28 returns a boolean if a field has been set.

### GetHostpci29

`func (o *UpdateVMConfigSyncRequest) GetHostpci29() string`

GetHostpci29 returns the Hostpci29 field if non-nil, zero value otherwise.

### GetHostpci29Ok

`func (o *UpdateVMConfigSyncRequest) GetHostpci29Ok() (*string, bool)`

GetHostpci29Ok returns a tuple with the Hostpci29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci29

`func (o *UpdateVMConfigSyncRequest) SetHostpci29(v string)`

SetHostpci29 sets Hostpci29 field to given value.

### HasHostpci29

`func (o *UpdateVMConfigSyncRequest) HasHostpci29() bool`

HasHostpci29 returns a boolean if a field has been set.

### GetHotplug

`func (o *UpdateVMConfigSyncRequest) GetHotplug() string`

GetHotplug returns the Hotplug field if non-nil, zero value otherwise.

### GetHotplugOk

`func (o *UpdateVMConfigSyncRequest) GetHotplugOk() (*string, bool)`

GetHotplugOk returns a tuple with the Hotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplug

`func (o *UpdateVMConfigSyncRequest) SetHotplug(v string)`

SetHotplug sets Hotplug field to given value.

### HasHotplug

`func (o *UpdateVMConfigSyncRequest) HasHotplug() bool`

HasHotplug returns a boolean if a field has been set.

### GetHugepages

`func (o *UpdateVMConfigSyncRequest) GetHugepages() string`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *UpdateVMConfigSyncRequest) GetHugepagesOk() (*string, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *UpdateVMConfigSyncRequest) SetHugepages(v string)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *UpdateVMConfigSyncRequest) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetIde0

`func (o *UpdateVMConfigSyncRequest) GetIde0() CreateVMRequestIde0`

GetIde0 returns the Ide0 field if non-nil, zero value otherwise.

### GetIde0Ok

`func (o *UpdateVMConfigSyncRequest) GetIde0Ok() (*CreateVMRequestIde0, bool)`

GetIde0Ok returns a tuple with the Ide0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde0

`func (o *UpdateVMConfigSyncRequest) SetIde0(v CreateVMRequestIde0)`

SetIde0 sets Ide0 field to given value.

### HasIde0

`func (o *UpdateVMConfigSyncRequest) HasIde0() bool`

HasIde0 returns a boolean if a field has been set.

### GetIde1

`func (o *UpdateVMConfigSyncRequest) GetIde1() CreateVMRequestIde0`

GetIde1 returns the Ide1 field if non-nil, zero value otherwise.

### GetIde1Ok

`func (o *UpdateVMConfigSyncRequest) GetIde1Ok() (*CreateVMRequestIde0, bool)`

GetIde1Ok returns a tuple with the Ide1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde1

`func (o *UpdateVMConfigSyncRequest) SetIde1(v CreateVMRequestIde0)`

SetIde1 sets Ide1 field to given value.

### HasIde1

`func (o *UpdateVMConfigSyncRequest) HasIde1() bool`

HasIde1 returns a boolean if a field has been set.

### GetIde2

`func (o *UpdateVMConfigSyncRequest) GetIde2() CreateVMRequestIde0`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *UpdateVMConfigSyncRequest) GetIde2Ok() (*CreateVMRequestIde0, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *UpdateVMConfigSyncRequest) SetIde2(v CreateVMRequestIde0)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *UpdateVMConfigSyncRequest) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetIde3

`func (o *UpdateVMConfigSyncRequest) GetIde3() CreateVMRequestIde0`

GetIde3 returns the Ide3 field if non-nil, zero value otherwise.

### GetIde3Ok

`func (o *UpdateVMConfigSyncRequest) GetIde3Ok() (*CreateVMRequestIde0, bool)`

GetIde3Ok returns a tuple with the Ide3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde3

`func (o *UpdateVMConfigSyncRequest) SetIde3(v CreateVMRequestIde0)`

SetIde3 sets Ide3 field to given value.

### HasIde3

`func (o *UpdateVMConfigSyncRequest) HasIde3() bool`

HasIde3 returns a boolean if a field has been set.

### GetIpconfig0

`func (o *UpdateVMConfigSyncRequest) GetIpconfig0() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig0 returns the Ipconfig0 field if non-nil, zero value otherwise.

### GetIpconfig0Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig0Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig0Ok returns a tuple with the Ipconfig0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig0

`func (o *UpdateVMConfigSyncRequest) SetIpconfig0(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig0 sets Ipconfig0 field to given value.

### HasIpconfig0

`func (o *UpdateVMConfigSyncRequest) HasIpconfig0() bool`

HasIpconfig0 returns a boolean if a field has been set.

### GetIpconfig1

`func (o *UpdateVMConfigSyncRequest) GetIpconfig1() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig1 returns the Ipconfig1 field if non-nil, zero value otherwise.

### GetIpconfig1Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig1Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig1Ok returns a tuple with the Ipconfig1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig1

`func (o *UpdateVMConfigSyncRequest) SetIpconfig1(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig1 sets Ipconfig1 field to given value.

### HasIpconfig1

`func (o *UpdateVMConfigSyncRequest) HasIpconfig1() bool`

HasIpconfig1 returns a boolean if a field has been set.

### GetIpconfig2

`func (o *UpdateVMConfigSyncRequest) GetIpconfig2() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig2 returns the Ipconfig2 field if non-nil, zero value otherwise.

### GetIpconfig2Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig2Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig2Ok returns a tuple with the Ipconfig2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig2

`func (o *UpdateVMConfigSyncRequest) SetIpconfig2(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig2 sets Ipconfig2 field to given value.

### HasIpconfig2

`func (o *UpdateVMConfigSyncRequest) HasIpconfig2() bool`

HasIpconfig2 returns a boolean if a field has been set.

### GetIpconfig3

`func (o *UpdateVMConfigSyncRequest) GetIpconfig3() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig3 returns the Ipconfig3 field if non-nil, zero value otherwise.

### GetIpconfig3Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig3Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig3Ok returns a tuple with the Ipconfig3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig3

`func (o *UpdateVMConfigSyncRequest) SetIpconfig3(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig3 sets Ipconfig3 field to given value.

### HasIpconfig3

`func (o *UpdateVMConfigSyncRequest) HasIpconfig3() bool`

HasIpconfig3 returns a boolean if a field has been set.

### GetIpconfig4

`func (o *UpdateVMConfigSyncRequest) GetIpconfig4() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig4 returns the Ipconfig4 field if non-nil, zero value otherwise.

### GetIpconfig4Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig4Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig4Ok returns a tuple with the Ipconfig4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig4

`func (o *UpdateVMConfigSyncRequest) SetIpconfig4(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig4 sets Ipconfig4 field to given value.

### HasIpconfig4

`func (o *UpdateVMConfigSyncRequest) HasIpconfig4() bool`

HasIpconfig4 returns a boolean if a field has been set.

### GetIpconfig5

`func (o *UpdateVMConfigSyncRequest) GetIpconfig5() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig5 returns the Ipconfig5 field if non-nil, zero value otherwise.

### GetIpconfig5Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig5Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig5Ok returns a tuple with the Ipconfig5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig5

`func (o *UpdateVMConfigSyncRequest) SetIpconfig5(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig5 sets Ipconfig5 field to given value.

### HasIpconfig5

`func (o *UpdateVMConfigSyncRequest) HasIpconfig5() bool`

HasIpconfig5 returns a boolean if a field has been set.

### GetIpconfig6

`func (o *UpdateVMConfigSyncRequest) GetIpconfig6() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig6 returns the Ipconfig6 field if non-nil, zero value otherwise.

### GetIpconfig6Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig6Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig6Ok returns a tuple with the Ipconfig6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig6

`func (o *UpdateVMConfigSyncRequest) SetIpconfig6(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig6 sets Ipconfig6 field to given value.

### HasIpconfig6

`func (o *UpdateVMConfigSyncRequest) HasIpconfig6() bool`

HasIpconfig6 returns a boolean if a field has been set.

### GetIpconfig7

`func (o *UpdateVMConfigSyncRequest) GetIpconfig7() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig7 returns the Ipconfig7 field if non-nil, zero value otherwise.

### GetIpconfig7Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig7Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig7Ok returns a tuple with the Ipconfig7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig7

`func (o *UpdateVMConfigSyncRequest) SetIpconfig7(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig7 sets Ipconfig7 field to given value.

### HasIpconfig7

`func (o *UpdateVMConfigSyncRequest) HasIpconfig7() bool`

HasIpconfig7 returns a boolean if a field has been set.

### GetIpconfig8

`func (o *UpdateVMConfigSyncRequest) GetIpconfig8() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig8 returns the Ipconfig8 field if non-nil, zero value otherwise.

### GetIpconfig8Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig8Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig8Ok returns a tuple with the Ipconfig8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig8

`func (o *UpdateVMConfigSyncRequest) SetIpconfig8(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig8 sets Ipconfig8 field to given value.

### HasIpconfig8

`func (o *UpdateVMConfigSyncRequest) HasIpconfig8() bool`

HasIpconfig8 returns a boolean if a field has been set.

### GetIpconfig9

`func (o *UpdateVMConfigSyncRequest) GetIpconfig9() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig9 returns the Ipconfig9 field if non-nil, zero value otherwise.

### GetIpconfig9Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig9Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig9Ok returns a tuple with the Ipconfig9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig9

`func (o *UpdateVMConfigSyncRequest) SetIpconfig9(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig9 sets Ipconfig9 field to given value.

### HasIpconfig9

`func (o *UpdateVMConfigSyncRequest) HasIpconfig9() bool`

HasIpconfig9 returns a boolean if a field has been set.

### GetIpconfig10

`func (o *UpdateVMConfigSyncRequest) GetIpconfig10() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig10 returns the Ipconfig10 field if non-nil, zero value otherwise.

### GetIpconfig10Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig10Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig10Ok returns a tuple with the Ipconfig10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig10

`func (o *UpdateVMConfigSyncRequest) SetIpconfig10(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig10 sets Ipconfig10 field to given value.

### HasIpconfig10

`func (o *UpdateVMConfigSyncRequest) HasIpconfig10() bool`

HasIpconfig10 returns a boolean if a field has been set.

### GetIpconfig11

`func (o *UpdateVMConfigSyncRequest) GetIpconfig11() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig11 returns the Ipconfig11 field if non-nil, zero value otherwise.

### GetIpconfig11Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig11Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig11Ok returns a tuple with the Ipconfig11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig11

`func (o *UpdateVMConfigSyncRequest) SetIpconfig11(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig11 sets Ipconfig11 field to given value.

### HasIpconfig11

`func (o *UpdateVMConfigSyncRequest) HasIpconfig11() bool`

HasIpconfig11 returns a boolean if a field has been set.

### GetIpconfig12

`func (o *UpdateVMConfigSyncRequest) GetIpconfig12() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig12 returns the Ipconfig12 field if non-nil, zero value otherwise.

### GetIpconfig12Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig12Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig12Ok returns a tuple with the Ipconfig12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig12

`func (o *UpdateVMConfigSyncRequest) SetIpconfig12(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig12 sets Ipconfig12 field to given value.

### HasIpconfig12

`func (o *UpdateVMConfigSyncRequest) HasIpconfig12() bool`

HasIpconfig12 returns a boolean if a field has been set.

### GetIpconfig13

`func (o *UpdateVMConfigSyncRequest) GetIpconfig13() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig13 returns the Ipconfig13 field if non-nil, zero value otherwise.

### GetIpconfig13Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig13Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig13Ok returns a tuple with the Ipconfig13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig13

`func (o *UpdateVMConfigSyncRequest) SetIpconfig13(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig13 sets Ipconfig13 field to given value.

### HasIpconfig13

`func (o *UpdateVMConfigSyncRequest) HasIpconfig13() bool`

HasIpconfig13 returns a boolean if a field has been set.

### GetIpconfig14

`func (o *UpdateVMConfigSyncRequest) GetIpconfig14() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig14 returns the Ipconfig14 field if non-nil, zero value otherwise.

### GetIpconfig14Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig14Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig14Ok returns a tuple with the Ipconfig14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig14

`func (o *UpdateVMConfigSyncRequest) SetIpconfig14(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig14 sets Ipconfig14 field to given value.

### HasIpconfig14

`func (o *UpdateVMConfigSyncRequest) HasIpconfig14() bool`

HasIpconfig14 returns a boolean if a field has been set.

### GetIpconfig15

`func (o *UpdateVMConfigSyncRequest) GetIpconfig15() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig15 returns the Ipconfig15 field if non-nil, zero value otherwise.

### GetIpconfig15Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig15Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig15Ok returns a tuple with the Ipconfig15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig15

`func (o *UpdateVMConfigSyncRequest) SetIpconfig15(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig15 sets Ipconfig15 field to given value.

### HasIpconfig15

`func (o *UpdateVMConfigSyncRequest) HasIpconfig15() bool`

HasIpconfig15 returns a boolean if a field has been set.

### GetIpconfig16

`func (o *UpdateVMConfigSyncRequest) GetIpconfig16() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig16 returns the Ipconfig16 field if non-nil, zero value otherwise.

### GetIpconfig16Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig16Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig16Ok returns a tuple with the Ipconfig16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig16

`func (o *UpdateVMConfigSyncRequest) SetIpconfig16(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig16 sets Ipconfig16 field to given value.

### HasIpconfig16

`func (o *UpdateVMConfigSyncRequest) HasIpconfig16() bool`

HasIpconfig16 returns a boolean if a field has been set.

### GetIpconfig17

`func (o *UpdateVMConfigSyncRequest) GetIpconfig17() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig17 returns the Ipconfig17 field if non-nil, zero value otherwise.

### GetIpconfig17Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig17Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig17Ok returns a tuple with the Ipconfig17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig17

`func (o *UpdateVMConfigSyncRequest) SetIpconfig17(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig17 sets Ipconfig17 field to given value.

### HasIpconfig17

`func (o *UpdateVMConfigSyncRequest) HasIpconfig17() bool`

HasIpconfig17 returns a boolean if a field has been set.

### GetIpconfig18

`func (o *UpdateVMConfigSyncRequest) GetIpconfig18() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig18 returns the Ipconfig18 field if non-nil, zero value otherwise.

### GetIpconfig18Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig18Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig18Ok returns a tuple with the Ipconfig18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig18

`func (o *UpdateVMConfigSyncRequest) SetIpconfig18(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig18 sets Ipconfig18 field to given value.

### HasIpconfig18

`func (o *UpdateVMConfigSyncRequest) HasIpconfig18() bool`

HasIpconfig18 returns a boolean if a field has been set.

### GetIpconfig19

`func (o *UpdateVMConfigSyncRequest) GetIpconfig19() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig19 returns the Ipconfig19 field if non-nil, zero value otherwise.

### GetIpconfig19Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig19Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig19Ok returns a tuple with the Ipconfig19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig19

`func (o *UpdateVMConfigSyncRequest) SetIpconfig19(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig19 sets Ipconfig19 field to given value.

### HasIpconfig19

`func (o *UpdateVMConfigSyncRequest) HasIpconfig19() bool`

HasIpconfig19 returns a boolean if a field has been set.

### GetIpconfig20

`func (o *UpdateVMConfigSyncRequest) GetIpconfig20() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig20 returns the Ipconfig20 field if non-nil, zero value otherwise.

### GetIpconfig20Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig20Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig20Ok returns a tuple with the Ipconfig20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig20

`func (o *UpdateVMConfigSyncRequest) SetIpconfig20(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig20 sets Ipconfig20 field to given value.

### HasIpconfig20

`func (o *UpdateVMConfigSyncRequest) HasIpconfig20() bool`

HasIpconfig20 returns a boolean if a field has been set.

### GetIpconfig21

`func (o *UpdateVMConfigSyncRequest) GetIpconfig21() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig21 returns the Ipconfig21 field if non-nil, zero value otherwise.

### GetIpconfig21Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig21Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig21Ok returns a tuple with the Ipconfig21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig21

`func (o *UpdateVMConfigSyncRequest) SetIpconfig21(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig21 sets Ipconfig21 field to given value.

### HasIpconfig21

`func (o *UpdateVMConfigSyncRequest) HasIpconfig21() bool`

HasIpconfig21 returns a boolean if a field has been set.

### GetIpconfig22

`func (o *UpdateVMConfigSyncRequest) GetIpconfig22() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig22 returns the Ipconfig22 field if non-nil, zero value otherwise.

### GetIpconfig22Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig22Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig22Ok returns a tuple with the Ipconfig22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig22

`func (o *UpdateVMConfigSyncRequest) SetIpconfig22(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig22 sets Ipconfig22 field to given value.

### HasIpconfig22

`func (o *UpdateVMConfigSyncRequest) HasIpconfig22() bool`

HasIpconfig22 returns a boolean if a field has been set.

### GetIpconfig23

`func (o *UpdateVMConfigSyncRequest) GetIpconfig23() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig23 returns the Ipconfig23 field if non-nil, zero value otherwise.

### GetIpconfig23Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig23Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig23Ok returns a tuple with the Ipconfig23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig23

`func (o *UpdateVMConfigSyncRequest) SetIpconfig23(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig23 sets Ipconfig23 field to given value.

### HasIpconfig23

`func (o *UpdateVMConfigSyncRequest) HasIpconfig23() bool`

HasIpconfig23 returns a boolean if a field has been set.

### GetIpconfig24

`func (o *UpdateVMConfigSyncRequest) GetIpconfig24() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig24 returns the Ipconfig24 field if non-nil, zero value otherwise.

### GetIpconfig24Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig24Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig24Ok returns a tuple with the Ipconfig24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig24

`func (o *UpdateVMConfigSyncRequest) SetIpconfig24(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig24 sets Ipconfig24 field to given value.

### HasIpconfig24

`func (o *UpdateVMConfigSyncRequest) HasIpconfig24() bool`

HasIpconfig24 returns a boolean if a field has been set.

### GetIpconfig25

`func (o *UpdateVMConfigSyncRequest) GetIpconfig25() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig25 returns the Ipconfig25 field if non-nil, zero value otherwise.

### GetIpconfig25Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig25Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig25Ok returns a tuple with the Ipconfig25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig25

`func (o *UpdateVMConfigSyncRequest) SetIpconfig25(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig25 sets Ipconfig25 field to given value.

### HasIpconfig25

`func (o *UpdateVMConfigSyncRequest) HasIpconfig25() bool`

HasIpconfig25 returns a boolean if a field has been set.

### GetIpconfig26

`func (o *UpdateVMConfigSyncRequest) GetIpconfig26() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig26 returns the Ipconfig26 field if non-nil, zero value otherwise.

### GetIpconfig26Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig26Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig26Ok returns a tuple with the Ipconfig26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig26

`func (o *UpdateVMConfigSyncRequest) SetIpconfig26(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig26 sets Ipconfig26 field to given value.

### HasIpconfig26

`func (o *UpdateVMConfigSyncRequest) HasIpconfig26() bool`

HasIpconfig26 returns a boolean if a field has been set.

### GetIpconfig27

`func (o *UpdateVMConfigSyncRequest) GetIpconfig27() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig27 returns the Ipconfig27 field if non-nil, zero value otherwise.

### GetIpconfig27Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig27Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig27Ok returns a tuple with the Ipconfig27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig27

`func (o *UpdateVMConfigSyncRequest) SetIpconfig27(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig27 sets Ipconfig27 field to given value.

### HasIpconfig27

`func (o *UpdateVMConfigSyncRequest) HasIpconfig27() bool`

HasIpconfig27 returns a boolean if a field has been set.

### GetIpconfig28

`func (o *UpdateVMConfigSyncRequest) GetIpconfig28() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig28 returns the Ipconfig28 field if non-nil, zero value otherwise.

### GetIpconfig28Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig28Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig28Ok returns a tuple with the Ipconfig28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig28

`func (o *UpdateVMConfigSyncRequest) SetIpconfig28(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig28 sets Ipconfig28 field to given value.

### HasIpconfig28

`func (o *UpdateVMConfigSyncRequest) HasIpconfig28() bool`

HasIpconfig28 returns a boolean if a field has been set.

### GetIpconfig29

`func (o *UpdateVMConfigSyncRequest) GetIpconfig29() GetVMConfig200ResponseDataIpconfig0`

GetIpconfig29 returns the Ipconfig29 field if non-nil, zero value otherwise.

### GetIpconfig29Ok

`func (o *UpdateVMConfigSyncRequest) GetIpconfig29Ok() (*GetVMConfig200ResponseDataIpconfig0, bool)`

GetIpconfig29Ok returns a tuple with the Ipconfig29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig29

`func (o *UpdateVMConfigSyncRequest) SetIpconfig29(v GetVMConfig200ResponseDataIpconfig0)`

SetIpconfig29 sets Ipconfig29 field to given value.

### HasIpconfig29

`func (o *UpdateVMConfigSyncRequest) HasIpconfig29() bool`

HasIpconfig29 returns a boolean if a field has been set.

### GetIvshmem

`func (o *UpdateVMConfigSyncRequest) GetIvshmem() GetVMConfig200ResponseDataIvshmem`

GetIvshmem returns the Ivshmem field if non-nil, zero value otherwise.

### GetIvshmemOk

`func (o *UpdateVMConfigSyncRequest) GetIvshmemOk() (*GetVMConfig200ResponseDataIvshmem, bool)`

GetIvshmemOk returns a tuple with the Ivshmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvshmem

`func (o *UpdateVMConfigSyncRequest) SetIvshmem(v GetVMConfig200ResponseDataIvshmem)`

SetIvshmem sets Ivshmem field to given value.

### HasIvshmem

`func (o *UpdateVMConfigSyncRequest) HasIvshmem() bool`

HasIvshmem returns a boolean if a field has been set.

### GetKeephugepages

`func (o *UpdateVMConfigSyncRequest) GetKeephugepages() bool`

GetKeephugepages returns the Keephugepages field if non-nil, zero value otherwise.

### GetKeephugepagesOk

`func (o *UpdateVMConfigSyncRequest) GetKeephugepagesOk() (*bool, bool)`

GetKeephugepagesOk returns a tuple with the Keephugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeephugepages

`func (o *UpdateVMConfigSyncRequest) SetKeephugepages(v bool)`

SetKeephugepages sets Keephugepages field to given value.

### HasKeephugepages

`func (o *UpdateVMConfigSyncRequest) HasKeephugepages() bool`

HasKeephugepages returns a boolean if a field has been set.

### GetKeyboard

`func (o *UpdateVMConfigSyncRequest) GetKeyboard() string`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *UpdateVMConfigSyncRequest) GetKeyboardOk() (*string, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *UpdateVMConfigSyncRequest) SetKeyboard(v string)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *UpdateVMConfigSyncRequest) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetKvm

`func (o *UpdateVMConfigSyncRequest) GetKvm() bool`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *UpdateVMConfigSyncRequest) GetKvmOk() (*bool, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *UpdateVMConfigSyncRequest) SetKvm(v bool)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *UpdateVMConfigSyncRequest) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLocaltime

`func (o *UpdateVMConfigSyncRequest) GetLocaltime() bool`

GetLocaltime returns the Localtime field if non-nil, zero value otherwise.

### GetLocaltimeOk

`func (o *UpdateVMConfigSyncRequest) GetLocaltimeOk() (*bool, bool)`

GetLocaltimeOk returns a tuple with the Localtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaltime

`func (o *UpdateVMConfigSyncRequest) SetLocaltime(v bool)`

SetLocaltime sets Localtime field to given value.

### HasLocaltime

`func (o *UpdateVMConfigSyncRequest) HasLocaltime() bool`

HasLocaltime returns a boolean if a field has been set.

### GetLock

`func (o *UpdateVMConfigSyncRequest) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *UpdateVMConfigSyncRequest) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *UpdateVMConfigSyncRequest) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *UpdateVMConfigSyncRequest) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMachine

`func (o *UpdateVMConfigSyncRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *UpdateVMConfigSyncRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *UpdateVMConfigSyncRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *UpdateVMConfigSyncRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *UpdateVMConfigSyncRequest) GetMemory() CreateVMRequestMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateVMConfigSyncRequest) GetMemoryOk() (*CreateVMRequestMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateVMConfigSyncRequest) SetMemory(v CreateVMRequestMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateVMConfigSyncRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrateDowntime

`func (o *UpdateVMConfigSyncRequest) GetMigrateDowntime() float32`

GetMigrateDowntime returns the MigrateDowntime field if non-nil, zero value otherwise.

### GetMigrateDowntimeOk

`func (o *UpdateVMConfigSyncRequest) GetMigrateDowntimeOk() (*float32, bool)`

GetMigrateDowntimeOk returns a tuple with the MigrateDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateDowntime

`func (o *UpdateVMConfigSyncRequest) SetMigrateDowntime(v float32)`

SetMigrateDowntime sets MigrateDowntime field to given value.

### HasMigrateDowntime

`func (o *UpdateVMConfigSyncRequest) HasMigrateDowntime() bool`

HasMigrateDowntime returns a boolean if a field has been set.

### GetMigrateSpeed

`func (o *UpdateVMConfigSyncRequest) GetMigrateSpeed() int64`

GetMigrateSpeed returns the MigrateSpeed field if non-nil, zero value otherwise.

### GetMigrateSpeedOk

`func (o *UpdateVMConfigSyncRequest) GetMigrateSpeedOk() (*int64, bool)`

GetMigrateSpeedOk returns a tuple with the MigrateSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateSpeed

`func (o *UpdateVMConfigSyncRequest) SetMigrateSpeed(v int64)`

SetMigrateSpeed sets MigrateSpeed field to given value.

### HasMigrateSpeed

`func (o *UpdateVMConfigSyncRequest) HasMigrateSpeed() bool`

HasMigrateSpeed returns a boolean if a field has been set.

### GetName

`func (o *UpdateVMConfigSyncRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVMConfigSyncRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVMConfigSyncRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVMConfigSyncRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *UpdateVMConfigSyncRequest) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *UpdateVMConfigSyncRequest) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *UpdateVMConfigSyncRequest) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *UpdateVMConfigSyncRequest) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *UpdateVMConfigSyncRequest) GetNet0() GetVMConfig200ResponseDataNet0`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *UpdateVMConfigSyncRequest) GetNet0Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *UpdateVMConfigSyncRequest) SetNet0(v GetVMConfig200ResponseDataNet0)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *UpdateVMConfigSyncRequest) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *UpdateVMConfigSyncRequest) GetNet1() GetVMConfig200ResponseDataNet0`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *UpdateVMConfigSyncRequest) GetNet1Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *UpdateVMConfigSyncRequest) SetNet1(v GetVMConfig200ResponseDataNet0)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *UpdateVMConfigSyncRequest) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *UpdateVMConfigSyncRequest) GetNet2() GetVMConfig200ResponseDataNet0`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *UpdateVMConfigSyncRequest) GetNet2Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *UpdateVMConfigSyncRequest) SetNet2(v GetVMConfig200ResponseDataNet0)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *UpdateVMConfigSyncRequest) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *UpdateVMConfigSyncRequest) GetNet3() GetVMConfig200ResponseDataNet0`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *UpdateVMConfigSyncRequest) GetNet3Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *UpdateVMConfigSyncRequest) SetNet3(v GetVMConfig200ResponseDataNet0)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *UpdateVMConfigSyncRequest) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *UpdateVMConfigSyncRequest) GetNet4() GetVMConfig200ResponseDataNet0`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *UpdateVMConfigSyncRequest) GetNet4Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *UpdateVMConfigSyncRequest) SetNet4(v GetVMConfig200ResponseDataNet0)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *UpdateVMConfigSyncRequest) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *UpdateVMConfigSyncRequest) GetNet5() GetVMConfig200ResponseDataNet0`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *UpdateVMConfigSyncRequest) GetNet5Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *UpdateVMConfigSyncRequest) SetNet5(v GetVMConfig200ResponseDataNet0)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *UpdateVMConfigSyncRequest) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *UpdateVMConfigSyncRequest) GetNet6() GetVMConfig200ResponseDataNet0`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *UpdateVMConfigSyncRequest) GetNet6Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *UpdateVMConfigSyncRequest) SetNet6(v GetVMConfig200ResponseDataNet0)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *UpdateVMConfigSyncRequest) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *UpdateVMConfigSyncRequest) GetNet7() GetVMConfig200ResponseDataNet0`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *UpdateVMConfigSyncRequest) GetNet7Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *UpdateVMConfigSyncRequest) SetNet7(v GetVMConfig200ResponseDataNet0)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *UpdateVMConfigSyncRequest) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *UpdateVMConfigSyncRequest) GetNet8() GetVMConfig200ResponseDataNet0`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *UpdateVMConfigSyncRequest) GetNet8Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *UpdateVMConfigSyncRequest) SetNet8(v GetVMConfig200ResponseDataNet0)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *UpdateVMConfigSyncRequest) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *UpdateVMConfigSyncRequest) GetNet9() GetVMConfig200ResponseDataNet0`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *UpdateVMConfigSyncRequest) GetNet9Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *UpdateVMConfigSyncRequest) SetNet9(v GetVMConfig200ResponseDataNet0)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *UpdateVMConfigSyncRequest) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *UpdateVMConfigSyncRequest) GetNet10() GetVMConfig200ResponseDataNet0`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *UpdateVMConfigSyncRequest) GetNet10Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *UpdateVMConfigSyncRequest) SetNet10(v GetVMConfig200ResponseDataNet0)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *UpdateVMConfigSyncRequest) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *UpdateVMConfigSyncRequest) GetNet11() GetVMConfig200ResponseDataNet0`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *UpdateVMConfigSyncRequest) GetNet11Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *UpdateVMConfigSyncRequest) SetNet11(v GetVMConfig200ResponseDataNet0)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *UpdateVMConfigSyncRequest) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *UpdateVMConfigSyncRequest) GetNet12() GetVMConfig200ResponseDataNet0`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *UpdateVMConfigSyncRequest) GetNet12Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *UpdateVMConfigSyncRequest) SetNet12(v GetVMConfig200ResponseDataNet0)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *UpdateVMConfigSyncRequest) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *UpdateVMConfigSyncRequest) GetNet13() GetVMConfig200ResponseDataNet0`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *UpdateVMConfigSyncRequest) GetNet13Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *UpdateVMConfigSyncRequest) SetNet13(v GetVMConfig200ResponseDataNet0)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *UpdateVMConfigSyncRequest) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *UpdateVMConfigSyncRequest) GetNet14() GetVMConfig200ResponseDataNet0`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *UpdateVMConfigSyncRequest) GetNet14Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *UpdateVMConfigSyncRequest) SetNet14(v GetVMConfig200ResponseDataNet0)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *UpdateVMConfigSyncRequest) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *UpdateVMConfigSyncRequest) GetNet15() GetVMConfig200ResponseDataNet0`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *UpdateVMConfigSyncRequest) GetNet15Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *UpdateVMConfigSyncRequest) SetNet15(v GetVMConfig200ResponseDataNet0)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *UpdateVMConfigSyncRequest) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *UpdateVMConfigSyncRequest) GetNet16() GetVMConfig200ResponseDataNet0`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *UpdateVMConfigSyncRequest) GetNet16Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *UpdateVMConfigSyncRequest) SetNet16(v GetVMConfig200ResponseDataNet0)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *UpdateVMConfigSyncRequest) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *UpdateVMConfigSyncRequest) GetNet17() GetVMConfig200ResponseDataNet0`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *UpdateVMConfigSyncRequest) GetNet17Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *UpdateVMConfigSyncRequest) SetNet17(v GetVMConfig200ResponseDataNet0)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *UpdateVMConfigSyncRequest) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *UpdateVMConfigSyncRequest) GetNet18() GetVMConfig200ResponseDataNet0`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *UpdateVMConfigSyncRequest) GetNet18Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *UpdateVMConfigSyncRequest) SetNet18(v GetVMConfig200ResponseDataNet0)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *UpdateVMConfigSyncRequest) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *UpdateVMConfigSyncRequest) GetNet19() GetVMConfig200ResponseDataNet0`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *UpdateVMConfigSyncRequest) GetNet19Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *UpdateVMConfigSyncRequest) SetNet19(v GetVMConfig200ResponseDataNet0)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *UpdateVMConfigSyncRequest) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *UpdateVMConfigSyncRequest) GetNet20() GetVMConfig200ResponseDataNet0`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *UpdateVMConfigSyncRequest) GetNet20Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *UpdateVMConfigSyncRequest) SetNet20(v GetVMConfig200ResponseDataNet0)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *UpdateVMConfigSyncRequest) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *UpdateVMConfigSyncRequest) GetNet21() GetVMConfig200ResponseDataNet0`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *UpdateVMConfigSyncRequest) GetNet21Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *UpdateVMConfigSyncRequest) SetNet21(v GetVMConfig200ResponseDataNet0)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *UpdateVMConfigSyncRequest) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *UpdateVMConfigSyncRequest) GetNet22() GetVMConfig200ResponseDataNet0`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *UpdateVMConfigSyncRequest) GetNet22Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *UpdateVMConfigSyncRequest) SetNet22(v GetVMConfig200ResponseDataNet0)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *UpdateVMConfigSyncRequest) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *UpdateVMConfigSyncRequest) GetNet23() GetVMConfig200ResponseDataNet0`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *UpdateVMConfigSyncRequest) GetNet23Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *UpdateVMConfigSyncRequest) SetNet23(v GetVMConfig200ResponseDataNet0)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *UpdateVMConfigSyncRequest) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *UpdateVMConfigSyncRequest) GetNet24() GetVMConfig200ResponseDataNet0`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *UpdateVMConfigSyncRequest) GetNet24Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *UpdateVMConfigSyncRequest) SetNet24(v GetVMConfig200ResponseDataNet0)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *UpdateVMConfigSyncRequest) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *UpdateVMConfigSyncRequest) GetNet25() GetVMConfig200ResponseDataNet0`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *UpdateVMConfigSyncRequest) GetNet25Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *UpdateVMConfigSyncRequest) SetNet25(v GetVMConfig200ResponseDataNet0)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *UpdateVMConfigSyncRequest) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *UpdateVMConfigSyncRequest) GetNet26() GetVMConfig200ResponseDataNet0`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *UpdateVMConfigSyncRequest) GetNet26Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *UpdateVMConfigSyncRequest) SetNet26(v GetVMConfig200ResponseDataNet0)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *UpdateVMConfigSyncRequest) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *UpdateVMConfigSyncRequest) GetNet27() GetVMConfig200ResponseDataNet0`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *UpdateVMConfigSyncRequest) GetNet27Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *UpdateVMConfigSyncRequest) SetNet27(v GetVMConfig200ResponseDataNet0)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *UpdateVMConfigSyncRequest) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *UpdateVMConfigSyncRequest) GetNet28() GetVMConfig200ResponseDataNet0`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *UpdateVMConfigSyncRequest) GetNet28Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *UpdateVMConfigSyncRequest) SetNet28(v GetVMConfig200ResponseDataNet0)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *UpdateVMConfigSyncRequest) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *UpdateVMConfigSyncRequest) GetNet29() GetVMConfig200ResponseDataNet0`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *UpdateVMConfigSyncRequest) GetNet29Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *UpdateVMConfigSyncRequest) SetNet29(v GetVMConfig200ResponseDataNet0)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *UpdateVMConfigSyncRequest) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *UpdateVMConfigSyncRequest) GetNet30() GetVMConfig200ResponseDataNet0`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *UpdateVMConfigSyncRequest) GetNet30Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *UpdateVMConfigSyncRequest) SetNet30(v GetVMConfig200ResponseDataNet0)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *UpdateVMConfigSyncRequest) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *UpdateVMConfigSyncRequest) GetNet31() GetVMConfig200ResponseDataNet0`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *UpdateVMConfigSyncRequest) GetNet31Ok() (*GetVMConfig200ResponseDataNet0, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *UpdateVMConfigSyncRequest) SetNet31(v GetVMConfig200ResponseDataNet0)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *UpdateVMConfigSyncRequest) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetNuma

`func (o *UpdateVMConfigSyncRequest) GetNuma() bool`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *UpdateVMConfigSyncRequest) GetNumaOk() (*bool, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *UpdateVMConfigSyncRequest) SetNuma(v bool)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *UpdateVMConfigSyncRequest) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetNuma0

`func (o *UpdateVMConfigSyncRequest) GetNuma0() GetVMConfig200ResponseDataNuma0`

GetNuma0 returns the Numa0 field if non-nil, zero value otherwise.

### GetNuma0Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma0Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma0Ok returns a tuple with the Numa0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma0

`func (o *UpdateVMConfigSyncRequest) SetNuma0(v GetVMConfig200ResponseDataNuma0)`

SetNuma0 sets Numa0 field to given value.

### HasNuma0

`func (o *UpdateVMConfigSyncRequest) HasNuma0() bool`

HasNuma0 returns a boolean if a field has been set.

### GetNuma1

`func (o *UpdateVMConfigSyncRequest) GetNuma1() GetVMConfig200ResponseDataNuma0`

GetNuma1 returns the Numa1 field if non-nil, zero value otherwise.

### GetNuma1Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma1Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma1Ok returns a tuple with the Numa1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma1

`func (o *UpdateVMConfigSyncRequest) SetNuma1(v GetVMConfig200ResponseDataNuma0)`

SetNuma1 sets Numa1 field to given value.

### HasNuma1

`func (o *UpdateVMConfigSyncRequest) HasNuma1() bool`

HasNuma1 returns a boolean if a field has been set.

### GetNuma2

`func (o *UpdateVMConfigSyncRequest) GetNuma2() GetVMConfig200ResponseDataNuma0`

GetNuma2 returns the Numa2 field if non-nil, zero value otherwise.

### GetNuma2Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma2Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma2Ok returns a tuple with the Numa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma2

`func (o *UpdateVMConfigSyncRequest) SetNuma2(v GetVMConfig200ResponseDataNuma0)`

SetNuma2 sets Numa2 field to given value.

### HasNuma2

`func (o *UpdateVMConfigSyncRequest) HasNuma2() bool`

HasNuma2 returns a boolean if a field has been set.

### GetNuma3

`func (o *UpdateVMConfigSyncRequest) GetNuma3() GetVMConfig200ResponseDataNuma0`

GetNuma3 returns the Numa3 field if non-nil, zero value otherwise.

### GetNuma3Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma3Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma3Ok returns a tuple with the Numa3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma3

`func (o *UpdateVMConfigSyncRequest) SetNuma3(v GetVMConfig200ResponseDataNuma0)`

SetNuma3 sets Numa3 field to given value.

### HasNuma3

`func (o *UpdateVMConfigSyncRequest) HasNuma3() bool`

HasNuma3 returns a boolean if a field has been set.

### GetNuma4

`func (o *UpdateVMConfigSyncRequest) GetNuma4() GetVMConfig200ResponseDataNuma0`

GetNuma4 returns the Numa4 field if non-nil, zero value otherwise.

### GetNuma4Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma4Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma4Ok returns a tuple with the Numa4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma4

`func (o *UpdateVMConfigSyncRequest) SetNuma4(v GetVMConfig200ResponseDataNuma0)`

SetNuma4 sets Numa4 field to given value.

### HasNuma4

`func (o *UpdateVMConfigSyncRequest) HasNuma4() bool`

HasNuma4 returns a boolean if a field has been set.

### GetNuma5

`func (o *UpdateVMConfigSyncRequest) GetNuma5() GetVMConfig200ResponseDataNuma0`

GetNuma5 returns the Numa5 field if non-nil, zero value otherwise.

### GetNuma5Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma5Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma5Ok returns a tuple with the Numa5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma5

`func (o *UpdateVMConfigSyncRequest) SetNuma5(v GetVMConfig200ResponseDataNuma0)`

SetNuma5 sets Numa5 field to given value.

### HasNuma5

`func (o *UpdateVMConfigSyncRequest) HasNuma5() bool`

HasNuma5 returns a boolean if a field has been set.

### GetNuma6

`func (o *UpdateVMConfigSyncRequest) GetNuma6() GetVMConfig200ResponseDataNuma0`

GetNuma6 returns the Numa6 field if non-nil, zero value otherwise.

### GetNuma6Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma6Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma6Ok returns a tuple with the Numa6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma6

`func (o *UpdateVMConfigSyncRequest) SetNuma6(v GetVMConfig200ResponseDataNuma0)`

SetNuma6 sets Numa6 field to given value.

### HasNuma6

`func (o *UpdateVMConfigSyncRequest) HasNuma6() bool`

HasNuma6 returns a boolean if a field has been set.

### GetNuma7

`func (o *UpdateVMConfigSyncRequest) GetNuma7() GetVMConfig200ResponseDataNuma0`

GetNuma7 returns the Numa7 field if non-nil, zero value otherwise.

### GetNuma7Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma7Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma7Ok returns a tuple with the Numa7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma7

`func (o *UpdateVMConfigSyncRequest) SetNuma7(v GetVMConfig200ResponseDataNuma0)`

SetNuma7 sets Numa7 field to given value.

### HasNuma7

`func (o *UpdateVMConfigSyncRequest) HasNuma7() bool`

HasNuma7 returns a boolean if a field has been set.

### GetNuma8

`func (o *UpdateVMConfigSyncRequest) GetNuma8() GetVMConfig200ResponseDataNuma0`

GetNuma8 returns the Numa8 field if non-nil, zero value otherwise.

### GetNuma8Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma8Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma8Ok returns a tuple with the Numa8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma8

`func (o *UpdateVMConfigSyncRequest) SetNuma8(v GetVMConfig200ResponseDataNuma0)`

SetNuma8 sets Numa8 field to given value.

### HasNuma8

`func (o *UpdateVMConfigSyncRequest) HasNuma8() bool`

HasNuma8 returns a boolean if a field has been set.

### GetNuma9

`func (o *UpdateVMConfigSyncRequest) GetNuma9() GetVMConfig200ResponseDataNuma0`

GetNuma9 returns the Numa9 field if non-nil, zero value otherwise.

### GetNuma9Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma9Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma9Ok returns a tuple with the Numa9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma9

`func (o *UpdateVMConfigSyncRequest) SetNuma9(v GetVMConfig200ResponseDataNuma0)`

SetNuma9 sets Numa9 field to given value.

### HasNuma9

`func (o *UpdateVMConfigSyncRequest) HasNuma9() bool`

HasNuma9 returns a boolean if a field has been set.

### GetNuma10

`func (o *UpdateVMConfigSyncRequest) GetNuma10() GetVMConfig200ResponseDataNuma0`

GetNuma10 returns the Numa10 field if non-nil, zero value otherwise.

### GetNuma10Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma10Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma10Ok returns a tuple with the Numa10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma10

`func (o *UpdateVMConfigSyncRequest) SetNuma10(v GetVMConfig200ResponseDataNuma0)`

SetNuma10 sets Numa10 field to given value.

### HasNuma10

`func (o *UpdateVMConfigSyncRequest) HasNuma10() bool`

HasNuma10 returns a boolean if a field has been set.

### GetNuma11

`func (o *UpdateVMConfigSyncRequest) GetNuma11() GetVMConfig200ResponseDataNuma0`

GetNuma11 returns the Numa11 field if non-nil, zero value otherwise.

### GetNuma11Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma11Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma11Ok returns a tuple with the Numa11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma11

`func (o *UpdateVMConfigSyncRequest) SetNuma11(v GetVMConfig200ResponseDataNuma0)`

SetNuma11 sets Numa11 field to given value.

### HasNuma11

`func (o *UpdateVMConfigSyncRequest) HasNuma11() bool`

HasNuma11 returns a boolean if a field has been set.

### GetNuma12

`func (o *UpdateVMConfigSyncRequest) GetNuma12() GetVMConfig200ResponseDataNuma0`

GetNuma12 returns the Numa12 field if non-nil, zero value otherwise.

### GetNuma12Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma12Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma12Ok returns a tuple with the Numa12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma12

`func (o *UpdateVMConfigSyncRequest) SetNuma12(v GetVMConfig200ResponseDataNuma0)`

SetNuma12 sets Numa12 field to given value.

### HasNuma12

`func (o *UpdateVMConfigSyncRequest) HasNuma12() bool`

HasNuma12 returns a boolean if a field has been set.

### GetNuma13

`func (o *UpdateVMConfigSyncRequest) GetNuma13() GetVMConfig200ResponseDataNuma0`

GetNuma13 returns the Numa13 field if non-nil, zero value otherwise.

### GetNuma13Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma13Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma13Ok returns a tuple with the Numa13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma13

`func (o *UpdateVMConfigSyncRequest) SetNuma13(v GetVMConfig200ResponseDataNuma0)`

SetNuma13 sets Numa13 field to given value.

### HasNuma13

`func (o *UpdateVMConfigSyncRequest) HasNuma13() bool`

HasNuma13 returns a boolean if a field has been set.

### GetNuma14

`func (o *UpdateVMConfigSyncRequest) GetNuma14() GetVMConfig200ResponseDataNuma0`

GetNuma14 returns the Numa14 field if non-nil, zero value otherwise.

### GetNuma14Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma14Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma14Ok returns a tuple with the Numa14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma14

`func (o *UpdateVMConfigSyncRequest) SetNuma14(v GetVMConfig200ResponseDataNuma0)`

SetNuma14 sets Numa14 field to given value.

### HasNuma14

`func (o *UpdateVMConfigSyncRequest) HasNuma14() bool`

HasNuma14 returns a boolean if a field has been set.

### GetNuma15

`func (o *UpdateVMConfigSyncRequest) GetNuma15() GetVMConfig200ResponseDataNuma0`

GetNuma15 returns the Numa15 field if non-nil, zero value otherwise.

### GetNuma15Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma15Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma15Ok returns a tuple with the Numa15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma15

`func (o *UpdateVMConfigSyncRequest) SetNuma15(v GetVMConfig200ResponseDataNuma0)`

SetNuma15 sets Numa15 field to given value.

### HasNuma15

`func (o *UpdateVMConfigSyncRequest) HasNuma15() bool`

HasNuma15 returns a boolean if a field has been set.

### GetNuma16

`func (o *UpdateVMConfigSyncRequest) GetNuma16() GetVMConfig200ResponseDataNuma0`

GetNuma16 returns the Numa16 field if non-nil, zero value otherwise.

### GetNuma16Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma16Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma16Ok returns a tuple with the Numa16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma16

`func (o *UpdateVMConfigSyncRequest) SetNuma16(v GetVMConfig200ResponseDataNuma0)`

SetNuma16 sets Numa16 field to given value.

### HasNuma16

`func (o *UpdateVMConfigSyncRequest) HasNuma16() bool`

HasNuma16 returns a boolean if a field has been set.

### GetNuma17

`func (o *UpdateVMConfigSyncRequest) GetNuma17() GetVMConfig200ResponseDataNuma0`

GetNuma17 returns the Numa17 field if non-nil, zero value otherwise.

### GetNuma17Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma17Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma17Ok returns a tuple with the Numa17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma17

`func (o *UpdateVMConfigSyncRequest) SetNuma17(v GetVMConfig200ResponseDataNuma0)`

SetNuma17 sets Numa17 field to given value.

### HasNuma17

`func (o *UpdateVMConfigSyncRequest) HasNuma17() bool`

HasNuma17 returns a boolean if a field has been set.

### GetNuma18

`func (o *UpdateVMConfigSyncRequest) GetNuma18() GetVMConfig200ResponseDataNuma0`

GetNuma18 returns the Numa18 field if non-nil, zero value otherwise.

### GetNuma18Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma18Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma18Ok returns a tuple with the Numa18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma18

`func (o *UpdateVMConfigSyncRequest) SetNuma18(v GetVMConfig200ResponseDataNuma0)`

SetNuma18 sets Numa18 field to given value.

### HasNuma18

`func (o *UpdateVMConfigSyncRequest) HasNuma18() bool`

HasNuma18 returns a boolean if a field has been set.

### GetNuma19

`func (o *UpdateVMConfigSyncRequest) GetNuma19() GetVMConfig200ResponseDataNuma0`

GetNuma19 returns the Numa19 field if non-nil, zero value otherwise.

### GetNuma19Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma19Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma19Ok returns a tuple with the Numa19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma19

`func (o *UpdateVMConfigSyncRequest) SetNuma19(v GetVMConfig200ResponseDataNuma0)`

SetNuma19 sets Numa19 field to given value.

### HasNuma19

`func (o *UpdateVMConfigSyncRequest) HasNuma19() bool`

HasNuma19 returns a boolean if a field has been set.

### GetNuma20

`func (o *UpdateVMConfigSyncRequest) GetNuma20() GetVMConfig200ResponseDataNuma0`

GetNuma20 returns the Numa20 field if non-nil, zero value otherwise.

### GetNuma20Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma20Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma20Ok returns a tuple with the Numa20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma20

`func (o *UpdateVMConfigSyncRequest) SetNuma20(v GetVMConfig200ResponseDataNuma0)`

SetNuma20 sets Numa20 field to given value.

### HasNuma20

`func (o *UpdateVMConfigSyncRequest) HasNuma20() bool`

HasNuma20 returns a boolean if a field has been set.

### GetNuma21

`func (o *UpdateVMConfigSyncRequest) GetNuma21() GetVMConfig200ResponseDataNuma0`

GetNuma21 returns the Numa21 field if non-nil, zero value otherwise.

### GetNuma21Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma21Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma21Ok returns a tuple with the Numa21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma21

`func (o *UpdateVMConfigSyncRequest) SetNuma21(v GetVMConfig200ResponseDataNuma0)`

SetNuma21 sets Numa21 field to given value.

### HasNuma21

`func (o *UpdateVMConfigSyncRequest) HasNuma21() bool`

HasNuma21 returns a boolean if a field has been set.

### GetNuma22

`func (o *UpdateVMConfigSyncRequest) GetNuma22() GetVMConfig200ResponseDataNuma0`

GetNuma22 returns the Numa22 field if non-nil, zero value otherwise.

### GetNuma22Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma22Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma22Ok returns a tuple with the Numa22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma22

`func (o *UpdateVMConfigSyncRequest) SetNuma22(v GetVMConfig200ResponseDataNuma0)`

SetNuma22 sets Numa22 field to given value.

### HasNuma22

`func (o *UpdateVMConfigSyncRequest) HasNuma22() bool`

HasNuma22 returns a boolean if a field has been set.

### GetNuma23

`func (o *UpdateVMConfigSyncRequest) GetNuma23() GetVMConfig200ResponseDataNuma0`

GetNuma23 returns the Numa23 field if non-nil, zero value otherwise.

### GetNuma23Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma23Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma23Ok returns a tuple with the Numa23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma23

`func (o *UpdateVMConfigSyncRequest) SetNuma23(v GetVMConfig200ResponseDataNuma0)`

SetNuma23 sets Numa23 field to given value.

### HasNuma23

`func (o *UpdateVMConfigSyncRequest) HasNuma23() bool`

HasNuma23 returns a boolean if a field has been set.

### GetNuma24

`func (o *UpdateVMConfigSyncRequest) GetNuma24() GetVMConfig200ResponseDataNuma0`

GetNuma24 returns the Numa24 field if non-nil, zero value otherwise.

### GetNuma24Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma24Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma24Ok returns a tuple with the Numa24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma24

`func (o *UpdateVMConfigSyncRequest) SetNuma24(v GetVMConfig200ResponseDataNuma0)`

SetNuma24 sets Numa24 field to given value.

### HasNuma24

`func (o *UpdateVMConfigSyncRequest) HasNuma24() bool`

HasNuma24 returns a boolean if a field has been set.

### GetNuma25

`func (o *UpdateVMConfigSyncRequest) GetNuma25() GetVMConfig200ResponseDataNuma0`

GetNuma25 returns the Numa25 field if non-nil, zero value otherwise.

### GetNuma25Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma25Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma25Ok returns a tuple with the Numa25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma25

`func (o *UpdateVMConfigSyncRequest) SetNuma25(v GetVMConfig200ResponseDataNuma0)`

SetNuma25 sets Numa25 field to given value.

### HasNuma25

`func (o *UpdateVMConfigSyncRequest) HasNuma25() bool`

HasNuma25 returns a boolean if a field has been set.

### GetNuma26

`func (o *UpdateVMConfigSyncRequest) GetNuma26() GetVMConfig200ResponseDataNuma0`

GetNuma26 returns the Numa26 field if non-nil, zero value otherwise.

### GetNuma26Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma26Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma26Ok returns a tuple with the Numa26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma26

`func (o *UpdateVMConfigSyncRequest) SetNuma26(v GetVMConfig200ResponseDataNuma0)`

SetNuma26 sets Numa26 field to given value.

### HasNuma26

`func (o *UpdateVMConfigSyncRequest) HasNuma26() bool`

HasNuma26 returns a boolean if a field has been set.

### GetNuma27

`func (o *UpdateVMConfigSyncRequest) GetNuma27() GetVMConfig200ResponseDataNuma0`

GetNuma27 returns the Numa27 field if non-nil, zero value otherwise.

### GetNuma27Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma27Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma27Ok returns a tuple with the Numa27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma27

`func (o *UpdateVMConfigSyncRequest) SetNuma27(v GetVMConfig200ResponseDataNuma0)`

SetNuma27 sets Numa27 field to given value.

### HasNuma27

`func (o *UpdateVMConfigSyncRequest) HasNuma27() bool`

HasNuma27 returns a boolean if a field has been set.

### GetNuma28

`func (o *UpdateVMConfigSyncRequest) GetNuma28() GetVMConfig200ResponseDataNuma0`

GetNuma28 returns the Numa28 field if non-nil, zero value otherwise.

### GetNuma28Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma28Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma28Ok returns a tuple with the Numa28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma28

`func (o *UpdateVMConfigSyncRequest) SetNuma28(v GetVMConfig200ResponseDataNuma0)`

SetNuma28 sets Numa28 field to given value.

### HasNuma28

`func (o *UpdateVMConfigSyncRequest) HasNuma28() bool`

HasNuma28 returns a boolean if a field has been set.

### GetNuma29

`func (o *UpdateVMConfigSyncRequest) GetNuma29() GetVMConfig200ResponseDataNuma0`

GetNuma29 returns the Numa29 field if non-nil, zero value otherwise.

### GetNuma29Ok

`func (o *UpdateVMConfigSyncRequest) GetNuma29Ok() (*GetVMConfig200ResponseDataNuma0, bool)`

GetNuma29Ok returns a tuple with the Numa29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma29

`func (o *UpdateVMConfigSyncRequest) SetNuma29(v GetVMConfig200ResponseDataNuma0)`

SetNuma29 sets Numa29 field to given value.

### HasNuma29

`func (o *UpdateVMConfigSyncRequest) HasNuma29() bool`

HasNuma29 returns a boolean if a field has been set.

### GetOnboot

`func (o *UpdateVMConfigSyncRequest) GetOnboot() bool`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *UpdateVMConfigSyncRequest) GetOnbootOk() (*bool, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *UpdateVMConfigSyncRequest) SetOnboot(v bool)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *UpdateVMConfigSyncRequest) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *UpdateVMConfigSyncRequest) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *UpdateVMConfigSyncRequest) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *UpdateVMConfigSyncRequest) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *UpdateVMConfigSyncRequest) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetParallel0

`func (o *UpdateVMConfigSyncRequest) GetParallel0() string`

GetParallel0 returns the Parallel0 field if non-nil, zero value otherwise.

### GetParallel0Ok

`func (o *UpdateVMConfigSyncRequest) GetParallel0Ok() (*string, bool)`

GetParallel0Ok returns a tuple with the Parallel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel0

`func (o *UpdateVMConfigSyncRequest) SetParallel0(v string)`

SetParallel0 sets Parallel0 field to given value.

### HasParallel0

`func (o *UpdateVMConfigSyncRequest) HasParallel0() bool`

HasParallel0 returns a boolean if a field has been set.

### GetParallel1

`func (o *UpdateVMConfigSyncRequest) GetParallel1() string`

GetParallel1 returns the Parallel1 field if non-nil, zero value otherwise.

### GetParallel1Ok

`func (o *UpdateVMConfigSyncRequest) GetParallel1Ok() (*string, bool)`

GetParallel1Ok returns a tuple with the Parallel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel1

`func (o *UpdateVMConfigSyncRequest) SetParallel1(v string)`

SetParallel1 sets Parallel1 field to given value.

### HasParallel1

`func (o *UpdateVMConfigSyncRequest) HasParallel1() bool`

HasParallel1 returns a boolean if a field has been set.

### GetParallel2

`func (o *UpdateVMConfigSyncRequest) GetParallel2() string`

GetParallel2 returns the Parallel2 field if non-nil, zero value otherwise.

### GetParallel2Ok

`func (o *UpdateVMConfigSyncRequest) GetParallel2Ok() (*string, bool)`

GetParallel2Ok returns a tuple with the Parallel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel2

`func (o *UpdateVMConfigSyncRequest) SetParallel2(v string)`

SetParallel2 sets Parallel2 field to given value.

### HasParallel2

`func (o *UpdateVMConfigSyncRequest) HasParallel2() bool`

HasParallel2 returns a boolean if a field has been set.

### GetParallel3

`func (o *UpdateVMConfigSyncRequest) GetParallel3() string`

GetParallel3 returns the Parallel3 field if non-nil, zero value otherwise.

### GetParallel3Ok

`func (o *UpdateVMConfigSyncRequest) GetParallel3Ok() (*string, bool)`

GetParallel3Ok returns a tuple with the Parallel3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel3

`func (o *UpdateVMConfigSyncRequest) SetParallel3(v string)`

SetParallel3 sets Parallel3 field to given value.

### HasParallel3

`func (o *UpdateVMConfigSyncRequest) HasParallel3() bool`

HasParallel3 returns a boolean if a field has been set.

### GetProtection

`func (o *UpdateVMConfigSyncRequest) GetProtection() bool`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *UpdateVMConfigSyncRequest) GetProtectionOk() (*bool, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *UpdateVMConfigSyncRequest) SetProtection(v bool)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *UpdateVMConfigSyncRequest) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetReboot

`func (o *UpdateVMConfigSyncRequest) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpdateVMConfigSyncRequest) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpdateVMConfigSyncRequest) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpdateVMConfigSyncRequest) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRevert

`func (o *UpdateVMConfigSyncRequest) GetRevert() string`

GetRevert returns the Revert field if non-nil, zero value otherwise.

### GetRevertOk

`func (o *UpdateVMConfigSyncRequest) GetRevertOk() (*string, bool)`

GetRevertOk returns a tuple with the Revert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevert

`func (o *UpdateVMConfigSyncRequest) SetRevert(v string)`

SetRevert sets Revert field to given value.

### HasRevert

`func (o *UpdateVMConfigSyncRequest) HasRevert() bool`

HasRevert returns a boolean if a field has been set.

### GetRng0

`func (o *UpdateVMConfigSyncRequest) GetRng0() GetVMConfig200ResponseDataRng0`

GetRng0 returns the Rng0 field if non-nil, zero value otherwise.

### GetRng0Ok

`func (o *UpdateVMConfigSyncRequest) GetRng0Ok() (*GetVMConfig200ResponseDataRng0, bool)`

GetRng0Ok returns a tuple with the Rng0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng0

`func (o *UpdateVMConfigSyncRequest) SetRng0(v GetVMConfig200ResponseDataRng0)`

SetRng0 sets Rng0 field to given value.

### HasRng0

`func (o *UpdateVMConfigSyncRequest) HasRng0() bool`

HasRng0 returns a boolean if a field has been set.

### GetSata0

`func (o *UpdateVMConfigSyncRequest) GetSata0() CreateVMRequestSata0`

GetSata0 returns the Sata0 field if non-nil, zero value otherwise.

### GetSata0Ok

`func (o *UpdateVMConfigSyncRequest) GetSata0Ok() (*CreateVMRequestSata0, bool)`

GetSata0Ok returns a tuple with the Sata0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata0

`func (o *UpdateVMConfigSyncRequest) SetSata0(v CreateVMRequestSata0)`

SetSata0 sets Sata0 field to given value.

### HasSata0

`func (o *UpdateVMConfigSyncRequest) HasSata0() bool`

HasSata0 returns a boolean if a field has been set.

### GetSata1

`func (o *UpdateVMConfigSyncRequest) GetSata1() CreateVMRequestSata0`

GetSata1 returns the Sata1 field if non-nil, zero value otherwise.

### GetSata1Ok

`func (o *UpdateVMConfigSyncRequest) GetSata1Ok() (*CreateVMRequestSata0, bool)`

GetSata1Ok returns a tuple with the Sata1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata1

`func (o *UpdateVMConfigSyncRequest) SetSata1(v CreateVMRequestSata0)`

SetSata1 sets Sata1 field to given value.

### HasSata1

`func (o *UpdateVMConfigSyncRequest) HasSata1() bool`

HasSata1 returns a boolean if a field has been set.

### GetSata2

`func (o *UpdateVMConfigSyncRequest) GetSata2() CreateVMRequestSata0`

GetSata2 returns the Sata2 field if non-nil, zero value otherwise.

### GetSata2Ok

`func (o *UpdateVMConfigSyncRequest) GetSata2Ok() (*CreateVMRequestSata0, bool)`

GetSata2Ok returns a tuple with the Sata2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata2

`func (o *UpdateVMConfigSyncRequest) SetSata2(v CreateVMRequestSata0)`

SetSata2 sets Sata2 field to given value.

### HasSata2

`func (o *UpdateVMConfigSyncRequest) HasSata2() bool`

HasSata2 returns a boolean if a field has been set.

### GetSata3

`func (o *UpdateVMConfigSyncRequest) GetSata3() CreateVMRequestSata0`

GetSata3 returns the Sata3 field if non-nil, zero value otherwise.

### GetSata3Ok

`func (o *UpdateVMConfigSyncRequest) GetSata3Ok() (*CreateVMRequestSata0, bool)`

GetSata3Ok returns a tuple with the Sata3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata3

`func (o *UpdateVMConfigSyncRequest) SetSata3(v CreateVMRequestSata0)`

SetSata3 sets Sata3 field to given value.

### HasSata3

`func (o *UpdateVMConfigSyncRequest) HasSata3() bool`

HasSata3 returns a boolean if a field has been set.

### GetSata4

`func (o *UpdateVMConfigSyncRequest) GetSata4() CreateVMRequestSata0`

GetSata4 returns the Sata4 field if non-nil, zero value otherwise.

### GetSata4Ok

`func (o *UpdateVMConfigSyncRequest) GetSata4Ok() (*CreateVMRequestSata0, bool)`

GetSata4Ok returns a tuple with the Sata4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata4

`func (o *UpdateVMConfigSyncRequest) SetSata4(v CreateVMRequestSata0)`

SetSata4 sets Sata4 field to given value.

### HasSata4

`func (o *UpdateVMConfigSyncRequest) HasSata4() bool`

HasSata4 returns a boolean if a field has been set.

### GetSata5

`func (o *UpdateVMConfigSyncRequest) GetSata5() CreateVMRequestSata0`

GetSata5 returns the Sata5 field if non-nil, zero value otherwise.

### GetSata5Ok

`func (o *UpdateVMConfigSyncRequest) GetSata5Ok() (*CreateVMRequestSata0, bool)`

GetSata5Ok returns a tuple with the Sata5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata5

`func (o *UpdateVMConfigSyncRequest) SetSata5(v CreateVMRequestSata0)`

SetSata5 sets Sata5 field to given value.

### HasSata5

`func (o *UpdateVMConfigSyncRequest) HasSata5() bool`

HasSata5 returns a boolean if a field has been set.

### GetScsi0

`func (o *UpdateVMConfigSyncRequest) GetScsi0() CreateVMRequestScsi0`

GetScsi0 returns the Scsi0 field if non-nil, zero value otherwise.

### GetScsi0Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi0Ok() (*CreateVMRequestScsi0, bool)`

GetScsi0Ok returns a tuple with the Scsi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi0

`func (o *UpdateVMConfigSyncRequest) SetScsi0(v CreateVMRequestScsi0)`

SetScsi0 sets Scsi0 field to given value.

### HasScsi0

`func (o *UpdateVMConfigSyncRequest) HasScsi0() bool`

HasScsi0 returns a boolean if a field has been set.

### GetScsi1

`func (o *UpdateVMConfigSyncRequest) GetScsi1() CreateVMRequestScsi0`

GetScsi1 returns the Scsi1 field if non-nil, zero value otherwise.

### GetScsi1Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi1Ok() (*CreateVMRequestScsi0, bool)`

GetScsi1Ok returns a tuple with the Scsi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi1

`func (o *UpdateVMConfigSyncRequest) SetScsi1(v CreateVMRequestScsi0)`

SetScsi1 sets Scsi1 field to given value.

### HasScsi1

`func (o *UpdateVMConfigSyncRequest) HasScsi1() bool`

HasScsi1 returns a boolean if a field has been set.

### GetScsi2

`func (o *UpdateVMConfigSyncRequest) GetScsi2() CreateVMRequestScsi0`

GetScsi2 returns the Scsi2 field if non-nil, zero value otherwise.

### GetScsi2Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi2Ok() (*CreateVMRequestScsi0, bool)`

GetScsi2Ok returns a tuple with the Scsi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi2

`func (o *UpdateVMConfigSyncRequest) SetScsi2(v CreateVMRequestScsi0)`

SetScsi2 sets Scsi2 field to given value.

### HasScsi2

`func (o *UpdateVMConfigSyncRequest) HasScsi2() bool`

HasScsi2 returns a boolean if a field has been set.

### GetScsi3

`func (o *UpdateVMConfigSyncRequest) GetScsi3() CreateVMRequestScsi0`

GetScsi3 returns the Scsi3 field if non-nil, zero value otherwise.

### GetScsi3Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi3Ok() (*CreateVMRequestScsi0, bool)`

GetScsi3Ok returns a tuple with the Scsi3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi3

`func (o *UpdateVMConfigSyncRequest) SetScsi3(v CreateVMRequestScsi0)`

SetScsi3 sets Scsi3 field to given value.

### HasScsi3

`func (o *UpdateVMConfigSyncRequest) HasScsi3() bool`

HasScsi3 returns a boolean if a field has been set.

### GetScsi4

`func (o *UpdateVMConfigSyncRequest) GetScsi4() CreateVMRequestScsi0`

GetScsi4 returns the Scsi4 field if non-nil, zero value otherwise.

### GetScsi4Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi4Ok() (*CreateVMRequestScsi0, bool)`

GetScsi4Ok returns a tuple with the Scsi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi4

`func (o *UpdateVMConfigSyncRequest) SetScsi4(v CreateVMRequestScsi0)`

SetScsi4 sets Scsi4 field to given value.

### HasScsi4

`func (o *UpdateVMConfigSyncRequest) HasScsi4() bool`

HasScsi4 returns a boolean if a field has been set.

### GetScsi5

`func (o *UpdateVMConfigSyncRequest) GetScsi5() CreateVMRequestScsi0`

GetScsi5 returns the Scsi5 field if non-nil, zero value otherwise.

### GetScsi5Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi5Ok() (*CreateVMRequestScsi0, bool)`

GetScsi5Ok returns a tuple with the Scsi5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi5

`func (o *UpdateVMConfigSyncRequest) SetScsi5(v CreateVMRequestScsi0)`

SetScsi5 sets Scsi5 field to given value.

### HasScsi5

`func (o *UpdateVMConfigSyncRequest) HasScsi5() bool`

HasScsi5 returns a boolean if a field has been set.

### GetScsi6

`func (o *UpdateVMConfigSyncRequest) GetScsi6() CreateVMRequestScsi0`

GetScsi6 returns the Scsi6 field if non-nil, zero value otherwise.

### GetScsi6Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi6Ok() (*CreateVMRequestScsi0, bool)`

GetScsi6Ok returns a tuple with the Scsi6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi6

`func (o *UpdateVMConfigSyncRequest) SetScsi6(v CreateVMRequestScsi0)`

SetScsi6 sets Scsi6 field to given value.

### HasScsi6

`func (o *UpdateVMConfigSyncRequest) HasScsi6() bool`

HasScsi6 returns a boolean if a field has been set.

### GetScsi7

`func (o *UpdateVMConfigSyncRequest) GetScsi7() CreateVMRequestScsi0`

GetScsi7 returns the Scsi7 field if non-nil, zero value otherwise.

### GetScsi7Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi7Ok() (*CreateVMRequestScsi0, bool)`

GetScsi7Ok returns a tuple with the Scsi7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi7

`func (o *UpdateVMConfigSyncRequest) SetScsi7(v CreateVMRequestScsi0)`

SetScsi7 sets Scsi7 field to given value.

### HasScsi7

`func (o *UpdateVMConfigSyncRequest) HasScsi7() bool`

HasScsi7 returns a boolean if a field has been set.

### GetScsi8

`func (o *UpdateVMConfigSyncRequest) GetScsi8() CreateVMRequestScsi0`

GetScsi8 returns the Scsi8 field if non-nil, zero value otherwise.

### GetScsi8Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi8Ok() (*CreateVMRequestScsi0, bool)`

GetScsi8Ok returns a tuple with the Scsi8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi8

`func (o *UpdateVMConfigSyncRequest) SetScsi8(v CreateVMRequestScsi0)`

SetScsi8 sets Scsi8 field to given value.

### HasScsi8

`func (o *UpdateVMConfigSyncRequest) HasScsi8() bool`

HasScsi8 returns a boolean if a field has been set.

### GetScsi9

`func (o *UpdateVMConfigSyncRequest) GetScsi9() CreateVMRequestScsi0`

GetScsi9 returns the Scsi9 field if non-nil, zero value otherwise.

### GetScsi9Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi9Ok() (*CreateVMRequestScsi0, bool)`

GetScsi9Ok returns a tuple with the Scsi9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi9

`func (o *UpdateVMConfigSyncRequest) SetScsi9(v CreateVMRequestScsi0)`

SetScsi9 sets Scsi9 field to given value.

### HasScsi9

`func (o *UpdateVMConfigSyncRequest) HasScsi9() bool`

HasScsi9 returns a boolean if a field has been set.

### GetScsi10

`func (o *UpdateVMConfigSyncRequest) GetScsi10() CreateVMRequestScsi0`

GetScsi10 returns the Scsi10 field if non-nil, zero value otherwise.

### GetScsi10Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi10Ok() (*CreateVMRequestScsi0, bool)`

GetScsi10Ok returns a tuple with the Scsi10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi10

`func (o *UpdateVMConfigSyncRequest) SetScsi10(v CreateVMRequestScsi0)`

SetScsi10 sets Scsi10 field to given value.

### HasScsi10

`func (o *UpdateVMConfigSyncRequest) HasScsi10() bool`

HasScsi10 returns a boolean if a field has been set.

### GetScsi11

`func (o *UpdateVMConfigSyncRequest) GetScsi11() CreateVMRequestScsi0`

GetScsi11 returns the Scsi11 field if non-nil, zero value otherwise.

### GetScsi11Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi11Ok() (*CreateVMRequestScsi0, bool)`

GetScsi11Ok returns a tuple with the Scsi11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi11

`func (o *UpdateVMConfigSyncRequest) SetScsi11(v CreateVMRequestScsi0)`

SetScsi11 sets Scsi11 field to given value.

### HasScsi11

`func (o *UpdateVMConfigSyncRequest) HasScsi11() bool`

HasScsi11 returns a boolean if a field has been set.

### GetScsi12

`func (o *UpdateVMConfigSyncRequest) GetScsi12() CreateVMRequestScsi0`

GetScsi12 returns the Scsi12 field if non-nil, zero value otherwise.

### GetScsi12Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi12Ok() (*CreateVMRequestScsi0, bool)`

GetScsi12Ok returns a tuple with the Scsi12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi12

`func (o *UpdateVMConfigSyncRequest) SetScsi12(v CreateVMRequestScsi0)`

SetScsi12 sets Scsi12 field to given value.

### HasScsi12

`func (o *UpdateVMConfigSyncRequest) HasScsi12() bool`

HasScsi12 returns a boolean if a field has been set.

### GetScsi13

`func (o *UpdateVMConfigSyncRequest) GetScsi13() CreateVMRequestScsi0`

GetScsi13 returns the Scsi13 field if non-nil, zero value otherwise.

### GetScsi13Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi13Ok() (*CreateVMRequestScsi0, bool)`

GetScsi13Ok returns a tuple with the Scsi13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi13

`func (o *UpdateVMConfigSyncRequest) SetScsi13(v CreateVMRequestScsi0)`

SetScsi13 sets Scsi13 field to given value.

### HasScsi13

`func (o *UpdateVMConfigSyncRequest) HasScsi13() bool`

HasScsi13 returns a boolean if a field has been set.

### GetScsi14

`func (o *UpdateVMConfigSyncRequest) GetScsi14() CreateVMRequestScsi0`

GetScsi14 returns the Scsi14 field if non-nil, zero value otherwise.

### GetScsi14Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi14Ok() (*CreateVMRequestScsi0, bool)`

GetScsi14Ok returns a tuple with the Scsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi14

`func (o *UpdateVMConfigSyncRequest) SetScsi14(v CreateVMRequestScsi0)`

SetScsi14 sets Scsi14 field to given value.

### HasScsi14

`func (o *UpdateVMConfigSyncRequest) HasScsi14() bool`

HasScsi14 returns a boolean if a field has been set.

### GetScsi15

`func (o *UpdateVMConfigSyncRequest) GetScsi15() CreateVMRequestScsi0`

GetScsi15 returns the Scsi15 field if non-nil, zero value otherwise.

### GetScsi15Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi15Ok() (*CreateVMRequestScsi0, bool)`

GetScsi15Ok returns a tuple with the Scsi15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi15

`func (o *UpdateVMConfigSyncRequest) SetScsi15(v CreateVMRequestScsi0)`

SetScsi15 sets Scsi15 field to given value.

### HasScsi15

`func (o *UpdateVMConfigSyncRequest) HasScsi15() bool`

HasScsi15 returns a boolean if a field has been set.

### GetScsi16

`func (o *UpdateVMConfigSyncRequest) GetScsi16() CreateVMRequestScsi0`

GetScsi16 returns the Scsi16 field if non-nil, zero value otherwise.

### GetScsi16Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi16Ok() (*CreateVMRequestScsi0, bool)`

GetScsi16Ok returns a tuple with the Scsi16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi16

`func (o *UpdateVMConfigSyncRequest) SetScsi16(v CreateVMRequestScsi0)`

SetScsi16 sets Scsi16 field to given value.

### HasScsi16

`func (o *UpdateVMConfigSyncRequest) HasScsi16() bool`

HasScsi16 returns a boolean if a field has been set.

### GetScsi17

`func (o *UpdateVMConfigSyncRequest) GetScsi17() CreateVMRequestScsi0`

GetScsi17 returns the Scsi17 field if non-nil, zero value otherwise.

### GetScsi17Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi17Ok() (*CreateVMRequestScsi0, bool)`

GetScsi17Ok returns a tuple with the Scsi17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi17

`func (o *UpdateVMConfigSyncRequest) SetScsi17(v CreateVMRequestScsi0)`

SetScsi17 sets Scsi17 field to given value.

### HasScsi17

`func (o *UpdateVMConfigSyncRequest) HasScsi17() bool`

HasScsi17 returns a boolean if a field has been set.

### GetScsi18

`func (o *UpdateVMConfigSyncRequest) GetScsi18() CreateVMRequestScsi0`

GetScsi18 returns the Scsi18 field if non-nil, zero value otherwise.

### GetScsi18Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi18Ok() (*CreateVMRequestScsi0, bool)`

GetScsi18Ok returns a tuple with the Scsi18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi18

`func (o *UpdateVMConfigSyncRequest) SetScsi18(v CreateVMRequestScsi0)`

SetScsi18 sets Scsi18 field to given value.

### HasScsi18

`func (o *UpdateVMConfigSyncRequest) HasScsi18() bool`

HasScsi18 returns a boolean if a field has been set.

### GetScsi19

`func (o *UpdateVMConfigSyncRequest) GetScsi19() CreateVMRequestScsi0`

GetScsi19 returns the Scsi19 field if non-nil, zero value otherwise.

### GetScsi19Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi19Ok() (*CreateVMRequestScsi0, bool)`

GetScsi19Ok returns a tuple with the Scsi19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi19

`func (o *UpdateVMConfigSyncRequest) SetScsi19(v CreateVMRequestScsi0)`

SetScsi19 sets Scsi19 field to given value.

### HasScsi19

`func (o *UpdateVMConfigSyncRequest) HasScsi19() bool`

HasScsi19 returns a boolean if a field has been set.

### GetScsi20

`func (o *UpdateVMConfigSyncRequest) GetScsi20() CreateVMRequestScsi0`

GetScsi20 returns the Scsi20 field if non-nil, zero value otherwise.

### GetScsi20Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi20Ok() (*CreateVMRequestScsi0, bool)`

GetScsi20Ok returns a tuple with the Scsi20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi20

`func (o *UpdateVMConfigSyncRequest) SetScsi20(v CreateVMRequestScsi0)`

SetScsi20 sets Scsi20 field to given value.

### HasScsi20

`func (o *UpdateVMConfigSyncRequest) HasScsi20() bool`

HasScsi20 returns a boolean if a field has been set.

### GetScsi21

`func (o *UpdateVMConfigSyncRequest) GetScsi21() CreateVMRequestScsi0`

GetScsi21 returns the Scsi21 field if non-nil, zero value otherwise.

### GetScsi21Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi21Ok() (*CreateVMRequestScsi0, bool)`

GetScsi21Ok returns a tuple with the Scsi21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi21

`func (o *UpdateVMConfigSyncRequest) SetScsi21(v CreateVMRequestScsi0)`

SetScsi21 sets Scsi21 field to given value.

### HasScsi21

`func (o *UpdateVMConfigSyncRequest) HasScsi21() bool`

HasScsi21 returns a boolean if a field has been set.

### GetScsi22

`func (o *UpdateVMConfigSyncRequest) GetScsi22() CreateVMRequestScsi0`

GetScsi22 returns the Scsi22 field if non-nil, zero value otherwise.

### GetScsi22Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi22Ok() (*CreateVMRequestScsi0, bool)`

GetScsi22Ok returns a tuple with the Scsi22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi22

`func (o *UpdateVMConfigSyncRequest) SetScsi22(v CreateVMRequestScsi0)`

SetScsi22 sets Scsi22 field to given value.

### HasScsi22

`func (o *UpdateVMConfigSyncRequest) HasScsi22() bool`

HasScsi22 returns a boolean if a field has been set.

### GetScsi23

`func (o *UpdateVMConfigSyncRequest) GetScsi23() CreateVMRequestScsi0`

GetScsi23 returns the Scsi23 field if non-nil, zero value otherwise.

### GetScsi23Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi23Ok() (*CreateVMRequestScsi0, bool)`

GetScsi23Ok returns a tuple with the Scsi23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi23

`func (o *UpdateVMConfigSyncRequest) SetScsi23(v CreateVMRequestScsi0)`

SetScsi23 sets Scsi23 field to given value.

### HasScsi23

`func (o *UpdateVMConfigSyncRequest) HasScsi23() bool`

HasScsi23 returns a boolean if a field has been set.

### GetScsi24

`func (o *UpdateVMConfigSyncRequest) GetScsi24() CreateVMRequestScsi0`

GetScsi24 returns the Scsi24 field if non-nil, zero value otherwise.

### GetScsi24Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi24Ok() (*CreateVMRequestScsi0, bool)`

GetScsi24Ok returns a tuple with the Scsi24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi24

`func (o *UpdateVMConfigSyncRequest) SetScsi24(v CreateVMRequestScsi0)`

SetScsi24 sets Scsi24 field to given value.

### HasScsi24

`func (o *UpdateVMConfigSyncRequest) HasScsi24() bool`

HasScsi24 returns a boolean if a field has been set.

### GetScsi25

`func (o *UpdateVMConfigSyncRequest) GetScsi25() CreateVMRequestScsi0`

GetScsi25 returns the Scsi25 field if non-nil, zero value otherwise.

### GetScsi25Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi25Ok() (*CreateVMRequestScsi0, bool)`

GetScsi25Ok returns a tuple with the Scsi25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi25

`func (o *UpdateVMConfigSyncRequest) SetScsi25(v CreateVMRequestScsi0)`

SetScsi25 sets Scsi25 field to given value.

### HasScsi25

`func (o *UpdateVMConfigSyncRequest) HasScsi25() bool`

HasScsi25 returns a boolean if a field has been set.

### GetScsi26

`func (o *UpdateVMConfigSyncRequest) GetScsi26() CreateVMRequestScsi0`

GetScsi26 returns the Scsi26 field if non-nil, zero value otherwise.

### GetScsi26Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi26Ok() (*CreateVMRequestScsi0, bool)`

GetScsi26Ok returns a tuple with the Scsi26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi26

`func (o *UpdateVMConfigSyncRequest) SetScsi26(v CreateVMRequestScsi0)`

SetScsi26 sets Scsi26 field to given value.

### HasScsi26

`func (o *UpdateVMConfigSyncRequest) HasScsi26() bool`

HasScsi26 returns a boolean if a field has been set.

### GetScsi27

`func (o *UpdateVMConfigSyncRequest) GetScsi27() CreateVMRequestScsi0`

GetScsi27 returns the Scsi27 field if non-nil, zero value otherwise.

### GetScsi27Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi27Ok() (*CreateVMRequestScsi0, bool)`

GetScsi27Ok returns a tuple with the Scsi27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi27

`func (o *UpdateVMConfigSyncRequest) SetScsi27(v CreateVMRequestScsi0)`

SetScsi27 sets Scsi27 field to given value.

### HasScsi27

`func (o *UpdateVMConfigSyncRequest) HasScsi27() bool`

HasScsi27 returns a boolean if a field has been set.

### GetScsi28

`func (o *UpdateVMConfigSyncRequest) GetScsi28() CreateVMRequestScsi0`

GetScsi28 returns the Scsi28 field if non-nil, zero value otherwise.

### GetScsi28Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi28Ok() (*CreateVMRequestScsi0, bool)`

GetScsi28Ok returns a tuple with the Scsi28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi28

`func (o *UpdateVMConfigSyncRequest) SetScsi28(v CreateVMRequestScsi0)`

SetScsi28 sets Scsi28 field to given value.

### HasScsi28

`func (o *UpdateVMConfigSyncRequest) HasScsi28() bool`

HasScsi28 returns a boolean if a field has been set.

### GetScsi29

`func (o *UpdateVMConfigSyncRequest) GetScsi29() CreateVMRequestScsi0`

GetScsi29 returns the Scsi29 field if non-nil, zero value otherwise.

### GetScsi29Ok

`func (o *UpdateVMConfigSyncRequest) GetScsi29Ok() (*CreateVMRequestScsi0, bool)`

GetScsi29Ok returns a tuple with the Scsi29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi29

`func (o *UpdateVMConfigSyncRequest) SetScsi29(v CreateVMRequestScsi0)`

SetScsi29 sets Scsi29 field to given value.

### HasScsi29

`func (o *UpdateVMConfigSyncRequest) HasScsi29() bool`

HasScsi29 returns a boolean if a field has been set.

### GetScsihw

`func (o *UpdateVMConfigSyncRequest) GetScsihw() string`

GetScsihw returns the Scsihw field if non-nil, zero value otherwise.

### GetScsihwOk

`func (o *UpdateVMConfigSyncRequest) GetScsihwOk() (*string, bool)`

GetScsihwOk returns a tuple with the Scsihw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsihw

`func (o *UpdateVMConfigSyncRequest) SetScsihw(v string)`

SetScsihw sets Scsihw field to given value.

### HasScsihw

`func (o *UpdateVMConfigSyncRequest) HasScsihw() bool`

HasScsihw returns a boolean if a field has been set.

### GetSearchdomain

`func (o *UpdateVMConfigSyncRequest) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *UpdateVMConfigSyncRequest) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *UpdateVMConfigSyncRequest) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *UpdateVMConfigSyncRequest) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSerial0

`func (o *UpdateVMConfigSyncRequest) GetSerial0() string`

GetSerial0 returns the Serial0 field if non-nil, zero value otherwise.

### GetSerial0Ok

`func (o *UpdateVMConfigSyncRequest) GetSerial0Ok() (*string, bool)`

GetSerial0Ok returns a tuple with the Serial0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial0

`func (o *UpdateVMConfigSyncRequest) SetSerial0(v string)`

SetSerial0 sets Serial0 field to given value.

### HasSerial0

`func (o *UpdateVMConfigSyncRequest) HasSerial0() bool`

HasSerial0 returns a boolean if a field has been set.

### GetSerial1

`func (o *UpdateVMConfigSyncRequest) GetSerial1() string`

GetSerial1 returns the Serial1 field if non-nil, zero value otherwise.

### GetSerial1Ok

`func (o *UpdateVMConfigSyncRequest) GetSerial1Ok() (*string, bool)`

GetSerial1Ok returns a tuple with the Serial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial1

`func (o *UpdateVMConfigSyncRequest) SetSerial1(v string)`

SetSerial1 sets Serial1 field to given value.

### HasSerial1

`func (o *UpdateVMConfigSyncRequest) HasSerial1() bool`

HasSerial1 returns a boolean if a field has been set.

### GetSerial2

`func (o *UpdateVMConfigSyncRequest) GetSerial2() string`

GetSerial2 returns the Serial2 field if non-nil, zero value otherwise.

### GetSerial2Ok

`func (o *UpdateVMConfigSyncRequest) GetSerial2Ok() (*string, bool)`

GetSerial2Ok returns a tuple with the Serial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial2

`func (o *UpdateVMConfigSyncRequest) SetSerial2(v string)`

SetSerial2 sets Serial2 field to given value.

### HasSerial2

`func (o *UpdateVMConfigSyncRequest) HasSerial2() bool`

HasSerial2 returns a boolean if a field has been set.

### GetSerial3

`func (o *UpdateVMConfigSyncRequest) GetSerial3() string`

GetSerial3 returns the Serial3 field if non-nil, zero value otherwise.

### GetSerial3Ok

`func (o *UpdateVMConfigSyncRequest) GetSerial3Ok() (*string, bool)`

GetSerial3Ok returns a tuple with the Serial3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial3

`func (o *UpdateVMConfigSyncRequest) SetSerial3(v string)`

SetSerial3 sets Serial3 field to given value.

### HasSerial3

`func (o *UpdateVMConfigSyncRequest) HasSerial3() bool`

HasSerial3 returns a boolean if a field has been set.

### GetShares

`func (o *UpdateVMConfigSyncRequest) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *UpdateVMConfigSyncRequest) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *UpdateVMConfigSyncRequest) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *UpdateVMConfigSyncRequest) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSkiplock

`func (o *UpdateVMConfigSyncRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *UpdateVMConfigSyncRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *UpdateVMConfigSyncRequest) SetSkiplock(v bool)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *UpdateVMConfigSyncRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetSmbios1

`func (o *UpdateVMConfigSyncRequest) GetSmbios1() string`

GetSmbios1 returns the Smbios1 field if non-nil, zero value otherwise.

### GetSmbios1Ok

`func (o *UpdateVMConfigSyncRequest) GetSmbios1Ok() (*string, bool)`

GetSmbios1Ok returns a tuple with the Smbios1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbios1

`func (o *UpdateVMConfigSyncRequest) SetSmbios1(v string)`

SetSmbios1 sets Smbios1 field to given value.

### HasSmbios1

`func (o *UpdateVMConfigSyncRequest) HasSmbios1() bool`

HasSmbios1 returns a boolean if a field has been set.

### GetSmp

`func (o *UpdateVMConfigSyncRequest) GetSmp() int64`

GetSmp returns the Smp field if non-nil, zero value otherwise.

### GetSmpOk

`func (o *UpdateVMConfigSyncRequest) GetSmpOk() (*int64, bool)`

GetSmpOk returns a tuple with the Smp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmp

`func (o *UpdateVMConfigSyncRequest) SetSmp(v int64)`

SetSmp sets Smp field to given value.

### HasSmp

`func (o *UpdateVMConfigSyncRequest) HasSmp() bool`

HasSmp returns a boolean if a field has been set.

### GetSockets

`func (o *UpdateVMConfigSyncRequest) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *UpdateVMConfigSyncRequest) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *UpdateVMConfigSyncRequest) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *UpdateVMConfigSyncRequest) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpiceEnhancements

`func (o *UpdateVMConfigSyncRequest) GetSpiceEnhancements() GetVMConfig200ResponseDataSpiceEnhancements`

GetSpiceEnhancements returns the SpiceEnhancements field if non-nil, zero value otherwise.

### GetSpiceEnhancementsOk

`func (o *UpdateVMConfigSyncRequest) GetSpiceEnhancementsOk() (*GetVMConfig200ResponseDataSpiceEnhancements, bool)`

GetSpiceEnhancementsOk returns a tuple with the SpiceEnhancements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnhancements

`func (o *UpdateVMConfigSyncRequest) SetSpiceEnhancements(v GetVMConfig200ResponseDataSpiceEnhancements)`

SetSpiceEnhancements sets SpiceEnhancements field to given value.

### HasSpiceEnhancements

`func (o *UpdateVMConfigSyncRequest) HasSpiceEnhancements() bool`

HasSpiceEnhancements returns a boolean if a field has been set.

### GetSshkeys

`func (o *UpdateVMConfigSyncRequest) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *UpdateVMConfigSyncRequest) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *UpdateVMConfigSyncRequest) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *UpdateVMConfigSyncRequest) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.

### GetStartdate

`func (o *UpdateVMConfigSyncRequest) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *UpdateVMConfigSyncRequest) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *UpdateVMConfigSyncRequest) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *UpdateVMConfigSyncRequest) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetStartup

`func (o *UpdateVMConfigSyncRequest) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *UpdateVMConfigSyncRequest) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *UpdateVMConfigSyncRequest) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *UpdateVMConfigSyncRequest) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetTablet

`func (o *UpdateVMConfigSyncRequest) GetTablet() bool`

GetTablet returns the Tablet field if non-nil, zero value otherwise.

### GetTabletOk

`func (o *UpdateVMConfigSyncRequest) GetTabletOk() (*bool, bool)`

GetTabletOk returns a tuple with the Tablet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablet

`func (o *UpdateVMConfigSyncRequest) SetTablet(v bool)`

SetTablet sets Tablet field to given value.

### HasTablet

`func (o *UpdateVMConfigSyncRequest) HasTablet() bool`

HasTablet returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMConfigSyncRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMConfigSyncRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMConfigSyncRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMConfigSyncRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTdf

`func (o *UpdateVMConfigSyncRequest) GetTdf() bool`

GetTdf returns the Tdf field if non-nil, zero value otherwise.

### GetTdfOk

`func (o *UpdateVMConfigSyncRequest) GetTdfOk() (*bool, bool)`

GetTdfOk returns a tuple with the Tdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdf

`func (o *UpdateVMConfigSyncRequest) SetTdf(v bool)`

SetTdf sets Tdf field to given value.

### HasTdf

`func (o *UpdateVMConfigSyncRequest) HasTdf() bool`

HasTdf returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateVMConfigSyncRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateVMConfigSyncRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateVMConfigSyncRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateVMConfigSyncRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTpmstate0

`func (o *UpdateVMConfigSyncRequest) GetTpmstate0() CreateVMRequestTpmstate0`

GetTpmstate0 returns the Tpmstate0 field if non-nil, zero value otherwise.

### GetTpmstate0Ok

`func (o *UpdateVMConfigSyncRequest) GetTpmstate0Ok() (*CreateVMRequestTpmstate0, bool)`

GetTpmstate0Ok returns a tuple with the Tpmstate0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmstate0

`func (o *UpdateVMConfigSyncRequest) SetTpmstate0(v CreateVMRequestTpmstate0)`

SetTpmstate0 sets Tpmstate0 field to given value.

### HasTpmstate0

`func (o *UpdateVMConfigSyncRequest) HasTpmstate0() bool`

HasTpmstate0 returns a boolean if a field has been set.

### GetUnused0

`func (o *UpdateVMConfigSyncRequest) GetUnused0() GetVMConfig200ResponseDataUnused0`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused0Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *UpdateVMConfigSyncRequest) SetUnused0(v GetVMConfig200ResponseDataUnused0)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *UpdateVMConfigSyncRequest) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *UpdateVMConfigSyncRequest) GetUnused1() GetVMConfig200ResponseDataUnused0`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused1Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *UpdateVMConfigSyncRequest) SetUnused1(v GetVMConfig200ResponseDataUnused0)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *UpdateVMConfigSyncRequest) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *UpdateVMConfigSyncRequest) GetUnused2() GetVMConfig200ResponseDataUnused0`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused2Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *UpdateVMConfigSyncRequest) SetUnused2(v GetVMConfig200ResponseDataUnused0)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *UpdateVMConfigSyncRequest) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *UpdateVMConfigSyncRequest) GetUnused3() GetVMConfig200ResponseDataUnused0`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused3Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *UpdateVMConfigSyncRequest) SetUnused3(v GetVMConfig200ResponseDataUnused0)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *UpdateVMConfigSyncRequest) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *UpdateVMConfigSyncRequest) GetUnused4() GetVMConfig200ResponseDataUnused0`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused4Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *UpdateVMConfigSyncRequest) SetUnused4(v GetVMConfig200ResponseDataUnused0)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *UpdateVMConfigSyncRequest) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *UpdateVMConfigSyncRequest) GetUnused5() GetVMConfig200ResponseDataUnused0`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused5Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *UpdateVMConfigSyncRequest) SetUnused5(v GetVMConfig200ResponseDataUnused0)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *UpdateVMConfigSyncRequest) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *UpdateVMConfigSyncRequest) GetUnused6() GetVMConfig200ResponseDataUnused0`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused6Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *UpdateVMConfigSyncRequest) SetUnused6(v GetVMConfig200ResponseDataUnused0)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *UpdateVMConfigSyncRequest) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *UpdateVMConfigSyncRequest) GetUnused7() GetVMConfig200ResponseDataUnused0`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused7Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *UpdateVMConfigSyncRequest) SetUnused7(v GetVMConfig200ResponseDataUnused0)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *UpdateVMConfigSyncRequest) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *UpdateVMConfigSyncRequest) GetUnused8() GetVMConfig200ResponseDataUnused0`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused8Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *UpdateVMConfigSyncRequest) SetUnused8(v GetVMConfig200ResponseDataUnused0)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *UpdateVMConfigSyncRequest) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *UpdateVMConfigSyncRequest) GetUnused9() GetVMConfig200ResponseDataUnused0`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused9Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *UpdateVMConfigSyncRequest) SetUnused9(v GetVMConfig200ResponseDataUnused0)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *UpdateVMConfigSyncRequest) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *UpdateVMConfigSyncRequest) GetUnused10() GetVMConfig200ResponseDataUnused0`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused10Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *UpdateVMConfigSyncRequest) SetUnused10(v GetVMConfig200ResponseDataUnused0)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *UpdateVMConfigSyncRequest) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *UpdateVMConfigSyncRequest) GetUnused11() GetVMConfig200ResponseDataUnused0`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused11Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *UpdateVMConfigSyncRequest) SetUnused11(v GetVMConfig200ResponseDataUnused0)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *UpdateVMConfigSyncRequest) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *UpdateVMConfigSyncRequest) GetUnused12() GetVMConfig200ResponseDataUnused0`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused12Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *UpdateVMConfigSyncRequest) SetUnused12(v GetVMConfig200ResponseDataUnused0)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *UpdateVMConfigSyncRequest) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *UpdateVMConfigSyncRequest) GetUnused13() GetVMConfig200ResponseDataUnused0`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused13Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *UpdateVMConfigSyncRequest) SetUnused13(v GetVMConfig200ResponseDataUnused0)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *UpdateVMConfigSyncRequest) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *UpdateVMConfigSyncRequest) GetUnused14() GetVMConfig200ResponseDataUnused0`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused14Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *UpdateVMConfigSyncRequest) SetUnused14(v GetVMConfig200ResponseDataUnused0)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *UpdateVMConfigSyncRequest) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *UpdateVMConfigSyncRequest) GetUnused15() GetVMConfig200ResponseDataUnused0`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused15Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *UpdateVMConfigSyncRequest) SetUnused15(v GetVMConfig200ResponseDataUnused0)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *UpdateVMConfigSyncRequest) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *UpdateVMConfigSyncRequest) GetUnused16() GetVMConfig200ResponseDataUnused0`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused16Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *UpdateVMConfigSyncRequest) SetUnused16(v GetVMConfig200ResponseDataUnused0)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *UpdateVMConfigSyncRequest) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *UpdateVMConfigSyncRequest) GetUnused17() GetVMConfig200ResponseDataUnused0`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused17Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *UpdateVMConfigSyncRequest) SetUnused17(v GetVMConfig200ResponseDataUnused0)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *UpdateVMConfigSyncRequest) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *UpdateVMConfigSyncRequest) GetUnused18() GetVMConfig200ResponseDataUnused0`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused18Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *UpdateVMConfigSyncRequest) SetUnused18(v GetVMConfig200ResponseDataUnused0)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *UpdateVMConfigSyncRequest) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *UpdateVMConfigSyncRequest) GetUnused19() GetVMConfig200ResponseDataUnused0`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused19Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *UpdateVMConfigSyncRequest) SetUnused19(v GetVMConfig200ResponseDataUnused0)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *UpdateVMConfigSyncRequest) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *UpdateVMConfigSyncRequest) GetUnused20() GetVMConfig200ResponseDataUnused0`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused20Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *UpdateVMConfigSyncRequest) SetUnused20(v GetVMConfig200ResponseDataUnused0)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *UpdateVMConfigSyncRequest) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *UpdateVMConfigSyncRequest) GetUnused21() GetVMConfig200ResponseDataUnused0`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused21Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *UpdateVMConfigSyncRequest) SetUnused21(v GetVMConfig200ResponseDataUnused0)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *UpdateVMConfigSyncRequest) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *UpdateVMConfigSyncRequest) GetUnused22() GetVMConfig200ResponseDataUnused0`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused22Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *UpdateVMConfigSyncRequest) SetUnused22(v GetVMConfig200ResponseDataUnused0)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *UpdateVMConfigSyncRequest) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *UpdateVMConfigSyncRequest) GetUnused23() GetVMConfig200ResponseDataUnused0`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused23Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *UpdateVMConfigSyncRequest) SetUnused23(v GetVMConfig200ResponseDataUnused0)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *UpdateVMConfigSyncRequest) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *UpdateVMConfigSyncRequest) GetUnused24() GetVMConfig200ResponseDataUnused0`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused24Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *UpdateVMConfigSyncRequest) SetUnused24(v GetVMConfig200ResponseDataUnused0)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *UpdateVMConfigSyncRequest) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *UpdateVMConfigSyncRequest) GetUnused25() GetVMConfig200ResponseDataUnused0`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused25Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *UpdateVMConfigSyncRequest) SetUnused25(v GetVMConfig200ResponseDataUnused0)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *UpdateVMConfigSyncRequest) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *UpdateVMConfigSyncRequest) GetUnused26() GetVMConfig200ResponseDataUnused0`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused26Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *UpdateVMConfigSyncRequest) SetUnused26(v GetVMConfig200ResponseDataUnused0)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *UpdateVMConfigSyncRequest) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *UpdateVMConfigSyncRequest) GetUnused27() GetVMConfig200ResponseDataUnused0`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused27Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *UpdateVMConfigSyncRequest) SetUnused27(v GetVMConfig200ResponseDataUnused0)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *UpdateVMConfigSyncRequest) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *UpdateVMConfigSyncRequest) GetUnused28() GetVMConfig200ResponseDataUnused0`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused28Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *UpdateVMConfigSyncRequest) SetUnused28(v GetVMConfig200ResponseDataUnused0)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *UpdateVMConfigSyncRequest) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *UpdateVMConfigSyncRequest) GetUnused29() GetVMConfig200ResponseDataUnused0`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *UpdateVMConfigSyncRequest) GetUnused29Ok() (*GetVMConfig200ResponseDataUnused0, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *UpdateVMConfigSyncRequest) SetUnused29(v GetVMConfig200ResponseDataUnused0)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *UpdateVMConfigSyncRequest) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.

### GetUsb0

`func (o *UpdateVMConfigSyncRequest) GetUsb0() GetVMConfig200ResponseDataUsb0`

GetUsb0 returns the Usb0 field if non-nil, zero value otherwise.

### GetUsb0Ok

`func (o *UpdateVMConfigSyncRequest) GetUsb0Ok() (*GetVMConfig200ResponseDataUsb0, bool)`

GetUsb0Ok returns a tuple with the Usb0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb0

`func (o *UpdateVMConfigSyncRequest) SetUsb0(v GetVMConfig200ResponseDataUsb0)`

SetUsb0 sets Usb0 field to given value.

### HasUsb0

`func (o *UpdateVMConfigSyncRequest) HasUsb0() bool`

HasUsb0 returns a boolean if a field has been set.

### GetUsb1

`func (o *UpdateVMConfigSyncRequest) GetUsb1() GetVMConfig200ResponseDataUsb0`

GetUsb1 returns the Usb1 field if non-nil, zero value otherwise.

### GetUsb1Ok

`func (o *UpdateVMConfigSyncRequest) GetUsb1Ok() (*GetVMConfig200ResponseDataUsb0, bool)`

GetUsb1Ok returns a tuple with the Usb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb1

`func (o *UpdateVMConfigSyncRequest) SetUsb1(v GetVMConfig200ResponseDataUsb0)`

SetUsb1 sets Usb1 field to given value.

### HasUsb1

`func (o *UpdateVMConfigSyncRequest) HasUsb1() bool`

HasUsb1 returns a boolean if a field has been set.

### GetUsb2

`func (o *UpdateVMConfigSyncRequest) GetUsb2() GetVMConfig200ResponseDataUsb0`

GetUsb2 returns the Usb2 field if non-nil, zero value otherwise.

### GetUsb2Ok

`func (o *UpdateVMConfigSyncRequest) GetUsb2Ok() (*GetVMConfig200ResponseDataUsb0, bool)`

GetUsb2Ok returns a tuple with the Usb2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb2

`func (o *UpdateVMConfigSyncRequest) SetUsb2(v GetVMConfig200ResponseDataUsb0)`

SetUsb2 sets Usb2 field to given value.

### HasUsb2

`func (o *UpdateVMConfigSyncRequest) HasUsb2() bool`

HasUsb2 returns a boolean if a field has been set.

### GetUsb3

`func (o *UpdateVMConfigSyncRequest) GetUsb3() GetVMConfig200ResponseDataUsb0`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *UpdateVMConfigSyncRequest) GetUsb3Ok() (*GetVMConfig200ResponseDataUsb0, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *UpdateVMConfigSyncRequest) SetUsb3(v GetVMConfig200ResponseDataUsb0)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *UpdateVMConfigSyncRequest) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.

### GetVcpus

`func (o *UpdateVMConfigSyncRequest) GetVcpus() int64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *UpdateVMConfigSyncRequest) GetVcpusOk() (*int64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *UpdateVMConfigSyncRequest) SetVcpus(v int64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *UpdateVMConfigSyncRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVga

`func (o *UpdateVMConfigSyncRequest) GetVga() GetVMConfig200ResponseDataVga`

GetVga returns the Vga field if non-nil, zero value otherwise.

### GetVgaOk

`func (o *UpdateVMConfigSyncRequest) GetVgaOk() (*GetVMConfig200ResponseDataVga, bool)`

GetVgaOk returns a tuple with the Vga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVga

`func (o *UpdateVMConfigSyncRequest) SetVga(v GetVMConfig200ResponseDataVga)`

SetVga sets Vga field to given value.

### HasVga

`func (o *UpdateVMConfigSyncRequest) HasVga() bool`

HasVga returns a boolean if a field has been set.

### GetVirtio0

`func (o *UpdateVMConfigSyncRequest) GetVirtio0() CreateVMRequestVirtio0`

GetVirtio0 returns the Virtio0 field if non-nil, zero value otherwise.

### GetVirtio0Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio0Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio0Ok returns a tuple with the Virtio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio0

`func (o *UpdateVMConfigSyncRequest) SetVirtio0(v CreateVMRequestVirtio0)`

SetVirtio0 sets Virtio0 field to given value.

### HasVirtio0

`func (o *UpdateVMConfigSyncRequest) HasVirtio0() bool`

HasVirtio0 returns a boolean if a field has been set.

### GetVirtio1

`func (o *UpdateVMConfigSyncRequest) GetVirtio1() CreateVMRequestVirtio0`

GetVirtio1 returns the Virtio1 field if non-nil, zero value otherwise.

### GetVirtio1Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio1Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio1Ok returns a tuple with the Virtio1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio1

`func (o *UpdateVMConfigSyncRequest) SetVirtio1(v CreateVMRequestVirtio0)`

SetVirtio1 sets Virtio1 field to given value.

### HasVirtio1

`func (o *UpdateVMConfigSyncRequest) HasVirtio1() bool`

HasVirtio1 returns a boolean if a field has been set.

### GetVirtio2

`func (o *UpdateVMConfigSyncRequest) GetVirtio2() CreateVMRequestVirtio0`

GetVirtio2 returns the Virtio2 field if non-nil, zero value otherwise.

### GetVirtio2Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio2Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio2Ok returns a tuple with the Virtio2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio2

`func (o *UpdateVMConfigSyncRequest) SetVirtio2(v CreateVMRequestVirtio0)`

SetVirtio2 sets Virtio2 field to given value.

### HasVirtio2

`func (o *UpdateVMConfigSyncRequest) HasVirtio2() bool`

HasVirtio2 returns a boolean if a field has been set.

### GetVirtio3

`func (o *UpdateVMConfigSyncRequest) GetVirtio3() CreateVMRequestVirtio0`

GetVirtio3 returns the Virtio3 field if non-nil, zero value otherwise.

### GetVirtio3Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio3Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio3Ok returns a tuple with the Virtio3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio3

`func (o *UpdateVMConfigSyncRequest) SetVirtio3(v CreateVMRequestVirtio0)`

SetVirtio3 sets Virtio3 field to given value.

### HasVirtio3

`func (o *UpdateVMConfigSyncRequest) HasVirtio3() bool`

HasVirtio3 returns a boolean if a field has been set.

### GetVirtio4

`func (o *UpdateVMConfigSyncRequest) GetVirtio4() CreateVMRequestVirtio0`

GetVirtio4 returns the Virtio4 field if non-nil, zero value otherwise.

### GetVirtio4Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio4Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio4Ok returns a tuple with the Virtio4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio4

`func (o *UpdateVMConfigSyncRequest) SetVirtio4(v CreateVMRequestVirtio0)`

SetVirtio4 sets Virtio4 field to given value.

### HasVirtio4

`func (o *UpdateVMConfigSyncRequest) HasVirtio4() bool`

HasVirtio4 returns a boolean if a field has been set.

### GetVirtio5

`func (o *UpdateVMConfigSyncRequest) GetVirtio5() CreateVMRequestVirtio0`

GetVirtio5 returns the Virtio5 field if non-nil, zero value otherwise.

### GetVirtio5Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio5Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio5Ok returns a tuple with the Virtio5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio5

`func (o *UpdateVMConfigSyncRequest) SetVirtio5(v CreateVMRequestVirtio0)`

SetVirtio5 sets Virtio5 field to given value.

### HasVirtio5

`func (o *UpdateVMConfigSyncRequest) HasVirtio5() bool`

HasVirtio5 returns a boolean if a field has been set.

### GetVirtio6

`func (o *UpdateVMConfigSyncRequest) GetVirtio6() CreateVMRequestVirtio0`

GetVirtio6 returns the Virtio6 field if non-nil, zero value otherwise.

### GetVirtio6Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio6Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio6Ok returns a tuple with the Virtio6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio6

`func (o *UpdateVMConfigSyncRequest) SetVirtio6(v CreateVMRequestVirtio0)`

SetVirtio6 sets Virtio6 field to given value.

### HasVirtio6

`func (o *UpdateVMConfigSyncRequest) HasVirtio6() bool`

HasVirtio6 returns a boolean if a field has been set.

### GetVirtio7

`func (o *UpdateVMConfigSyncRequest) GetVirtio7() CreateVMRequestVirtio0`

GetVirtio7 returns the Virtio7 field if non-nil, zero value otherwise.

### GetVirtio7Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio7Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio7Ok returns a tuple with the Virtio7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio7

`func (o *UpdateVMConfigSyncRequest) SetVirtio7(v CreateVMRequestVirtio0)`

SetVirtio7 sets Virtio7 field to given value.

### HasVirtio7

`func (o *UpdateVMConfigSyncRequest) HasVirtio7() bool`

HasVirtio7 returns a boolean if a field has been set.

### GetVirtio8

`func (o *UpdateVMConfigSyncRequest) GetVirtio8() CreateVMRequestVirtio0`

GetVirtio8 returns the Virtio8 field if non-nil, zero value otherwise.

### GetVirtio8Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio8Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio8Ok returns a tuple with the Virtio8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio8

`func (o *UpdateVMConfigSyncRequest) SetVirtio8(v CreateVMRequestVirtio0)`

SetVirtio8 sets Virtio8 field to given value.

### HasVirtio8

`func (o *UpdateVMConfigSyncRequest) HasVirtio8() bool`

HasVirtio8 returns a boolean if a field has been set.

### GetVirtio9

`func (o *UpdateVMConfigSyncRequest) GetVirtio9() CreateVMRequestVirtio0`

GetVirtio9 returns the Virtio9 field if non-nil, zero value otherwise.

### GetVirtio9Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio9Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio9Ok returns a tuple with the Virtio9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio9

`func (o *UpdateVMConfigSyncRequest) SetVirtio9(v CreateVMRequestVirtio0)`

SetVirtio9 sets Virtio9 field to given value.

### HasVirtio9

`func (o *UpdateVMConfigSyncRequest) HasVirtio9() bool`

HasVirtio9 returns a boolean if a field has been set.

### GetVirtio10

`func (o *UpdateVMConfigSyncRequest) GetVirtio10() CreateVMRequestVirtio0`

GetVirtio10 returns the Virtio10 field if non-nil, zero value otherwise.

### GetVirtio10Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio10Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio10Ok returns a tuple with the Virtio10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio10

`func (o *UpdateVMConfigSyncRequest) SetVirtio10(v CreateVMRequestVirtio0)`

SetVirtio10 sets Virtio10 field to given value.

### HasVirtio10

`func (o *UpdateVMConfigSyncRequest) HasVirtio10() bool`

HasVirtio10 returns a boolean if a field has been set.

### GetVirtio11

`func (o *UpdateVMConfigSyncRequest) GetVirtio11() CreateVMRequestVirtio0`

GetVirtio11 returns the Virtio11 field if non-nil, zero value otherwise.

### GetVirtio11Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio11Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio11Ok returns a tuple with the Virtio11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio11

`func (o *UpdateVMConfigSyncRequest) SetVirtio11(v CreateVMRequestVirtio0)`

SetVirtio11 sets Virtio11 field to given value.

### HasVirtio11

`func (o *UpdateVMConfigSyncRequest) HasVirtio11() bool`

HasVirtio11 returns a boolean if a field has been set.

### GetVirtio12

`func (o *UpdateVMConfigSyncRequest) GetVirtio12() CreateVMRequestVirtio0`

GetVirtio12 returns the Virtio12 field if non-nil, zero value otherwise.

### GetVirtio12Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio12Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio12Ok returns a tuple with the Virtio12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio12

`func (o *UpdateVMConfigSyncRequest) SetVirtio12(v CreateVMRequestVirtio0)`

SetVirtio12 sets Virtio12 field to given value.

### HasVirtio12

`func (o *UpdateVMConfigSyncRequest) HasVirtio12() bool`

HasVirtio12 returns a boolean if a field has been set.

### GetVirtio13

`func (o *UpdateVMConfigSyncRequest) GetVirtio13() CreateVMRequestVirtio0`

GetVirtio13 returns the Virtio13 field if non-nil, zero value otherwise.

### GetVirtio13Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio13Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio13Ok returns a tuple with the Virtio13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio13

`func (o *UpdateVMConfigSyncRequest) SetVirtio13(v CreateVMRequestVirtio0)`

SetVirtio13 sets Virtio13 field to given value.

### HasVirtio13

`func (o *UpdateVMConfigSyncRequest) HasVirtio13() bool`

HasVirtio13 returns a boolean if a field has been set.

### GetVirtio14

`func (o *UpdateVMConfigSyncRequest) GetVirtio14() CreateVMRequestVirtio0`

GetVirtio14 returns the Virtio14 field if non-nil, zero value otherwise.

### GetVirtio14Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio14Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio14Ok returns a tuple with the Virtio14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio14

`func (o *UpdateVMConfigSyncRequest) SetVirtio14(v CreateVMRequestVirtio0)`

SetVirtio14 sets Virtio14 field to given value.

### HasVirtio14

`func (o *UpdateVMConfigSyncRequest) HasVirtio14() bool`

HasVirtio14 returns a boolean if a field has been set.

### GetVirtio15

`func (o *UpdateVMConfigSyncRequest) GetVirtio15() CreateVMRequestVirtio0`

GetVirtio15 returns the Virtio15 field if non-nil, zero value otherwise.

### GetVirtio15Ok

`func (o *UpdateVMConfigSyncRequest) GetVirtio15Ok() (*CreateVMRequestVirtio0, bool)`

GetVirtio15Ok returns a tuple with the Virtio15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio15

`func (o *UpdateVMConfigSyncRequest) SetVirtio15(v CreateVMRequestVirtio0)`

SetVirtio15 sets Virtio15 field to given value.

### HasVirtio15

`func (o *UpdateVMConfigSyncRequest) HasVirtio15() bool`

HasVirtio15 returns a boolean if a field has been set.

### GetVmgenid

`func (o *UpdateVMConfigSyncRequest) GetVmgenid() string`

GetVmgenid returns the Vmgenid field if non-nil, zero value otherwise.

### GetVmgenidOk

`func (o *UpdateVMConfigSyncRequest) GetVmgenidOk() (*string, bool)`

GetVmgenidOk returns a tuple with the Vmgenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmgenid

`func (o *UpdateVMConfigSyncRequest) SetVmgenid(v string)`

SetVmgenid sets Vmgenid field to given value.

### HasVmgenid

`func (o *UpdateVMConfigSyncRequest) HasVmgenid() bool`

HasVmgenid returns a boolean if a field has been set.

### GetVmstatestorage

`func (o *UpdateVMConfigSyncRequest) GetVmstatestorage() string`

GetVmstatestorage returns the Vmstatestorage field if non-nil, zero value otherwise.

### GetVmstatestorageOk

`func (o *UpdateVMConfigSyncRequest) GetVmstatestorageOk() (*string, bool)`

GetVmstatestorageOk returns a tuple with the Vmstatestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstatestorage

`func (o *UpdateVMConfigSyncRequest) SetVmstatestorage(v string)`

SetVmstatestorage sets Vmstatestorage field to given value.

### HasVmstatestorage

`func (o *UpdateVMConfigSyncRequest) HasVmstatestorage() bool`

HasVmstatestorage returns a boolean if a field has been set.

### GetWatchdog

`func (o *UpdateVMConfigSyncRequest) GetWatchdog() GetVMConfig200ResponseDataWatchdog`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *UpdateVMConfigSyncRequest) GetWatchdogOk() (*GetVMConfig200ResponseDataWatchdog, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *UpdateVMConfigSyncRequest) SetWatchdog(v GetVMConfig200ResponseDataWatchdog)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *UpdateVMConfigSyncRequest) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


