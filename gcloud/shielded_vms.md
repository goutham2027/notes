https://cloud.google.com/compute/shielded-vm/docs/shielded-vm

- Shielded VM offers verifiable integrity of VM instances
- Verifiable integrity is achieved through
  - Secure boot
  - Virtual trusted platform module (vTPM)
  - Integrity monitoring

### Secure boot
- ensures system only runs authentic s/w by verifying the digital
signature of all boot components
- halts the boot process if signature verification fails
- shielded VMs run firmwware signed and verfied by Google's Certificate
  Authority (CA). This ensures that the instance's firmware is
unmodified

### Virtual Trusted Platform Module (vTPM)
- specialized computer chip to protect objects, like keys and certs
  that are used to authenticate access to the system.
- vTPM enables Measured Boot by performing the measurements needed to
  create a known good boot baseline - integrity policy baseline.
- the integrity policy baseline is used for comparison with measurements
  from subsequent VM boots to determine if anything has changed

### Measured Boot
- During measured boot, a hash of each component (eg: firmware,
  bootloader or kernel) is created as the component is loaded, and that
hash is concatenated and rehased with the hashes of any components that
have already been loaded.
- During the first time boot of a VM instance, Measured boot creates the
  integrity policy baseline. Each time the VM boots, measurements are
taken again and stored in secure memory until the next reboot.
- Having 2 sets of measurements enables integrity monitoring to
  determine if there is any change in VM's boot sequence

### Integrity monitoring
- Integriry monitoring relies on the measurements created by Measured
  Boot
- Integrity monitoring compares the most recent boot measurements to the
  integrity policy baseline and returns a pair of pass/fail results for
both early boot sequence and late boot sequence
- early boot sequence - from the start of the UEFI frameware until it
  passed control to the bootloader.
- late boot - boot sequence from the bootloader until it passed control
  to the operating system kernel
- if the failure is expected, eg: after a system update, integrity
  policy baseline needs to be updated
- view integrity reports in cloud monitoring and set alerts on integrity
  failures
