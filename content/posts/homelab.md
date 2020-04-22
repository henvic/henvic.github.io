---
title: "Homelab: Intel NUC with the ESXi hypervisor"
type: post
description: "In this blog post I tell my experience mounting a minimalist homelab with a Intel NUC and the ESXi hypervisor."
date: "2020-04-20"
image: "/img/posts/homelab/intel-nuc-10.jpg"
hashtags: "vmware,esxi,homelab"
---
In this blog post, I'm going to talk a little about my experience running multiple operating systems with an Intel NUC I recently bought and the ESXi 7 hypervisor. My main idea was to use this homelab for:

* Network Area Storage (NAS)

* Kubernetes cluster

* FreeBSD playground

* Managing backup

* Stage and development machine for Go projects

## Hardware

I thought I'd get a full-size ATX tower case with plenty of space for expansion and components, but I was always open to any [form factor](https://en.wikipedia.org/wiki/Computer_form_factor). When browsing [r/homelab](https://www.reddit.com/r/homelab/) and asking for opinions from people elsewhere, I noticed Intel's [Next Unit of Computing](https://en.wikipedia.org/wiki/Next_Unit_of_Computing) (also known as Intel NUC) line of [barebone computers](https://en.wikipedia.org/wiki/Barebone_computer) stood out as a popular choice, and I decided to give it a chance.

[![Intel NUC Performance Kit photo](/img/posts/homelab/intel-nuc-10_small.jpg)](/img/posts/homelab/intel-nuc-10.jpg)

### If you want to buy the same thing

* [INTEL® NUC 10 Performance Kit — NUC10I7FNH](https://www.intel.com/content/www/us/en/products/boards-kits/nuc/kits/nuc10i7fnh.html) (Buy [from Amazon](https://amzn.to/3ewroRS))

* [Samsung 970 EVO Plus NVMe M.2 500GB](https://www.samsung.com/us/computing/memory-storage/solid-state-drives/ssd-970-evo-plus-nvme-m-2-500gb-mz-v7s500b-am/) (Buy [from Amazon](https://amzn.to/34LvpNK))

* [Crucial 32GB Kit (2 x 16GB) DDR4 2666MHz SODIMM](https://www.crucial.com/memory/ddr4/ct2k16g4s266m) (Buy [from Amazon](https://amzn.to/2xyv4Sx))

* [LaCie d2 Professional 8TB](https://www.lacie.com/products/d2/) (Buy [from Amazon](https://amzn.to/2yn5S1b))

_Disclaimer: As an Amazon Associate I earn from qualifying purchases._ <!-- Amazon requires me to post this -->

The best thing about this kit is the small factor. Unfortunately, it comes with a low-end GPU — otherwise, I'd use it to run the [X-Plane 11](https://www.x-plane.com) flight simulator too. I wished they made a model that supported [ECC memory](https://en.wikipedia.org/wiki/ECC_memory) for reliability too…

[![Next Unit of Computing (or NUC) is an Intel line of small-factor barebone computers.](/img/posts/homelab/intel-nuc-10-inside_small.jpg)](/img/posts/homelab/intel-nuc-10-inside.jpg)

*Next Unit of Computing (or NUC) is an Intel line of small-factor barebone computers.*

I'm going to use the LaCie d2 Professional external drive with [FreeNAS](https://www.freenas.org/), mostly to store my Lightroom library files.

In an ideal scenario, I'd be using at least an SSD, maybe with RAID 0. However, this would be too expensive, and this is a great alternative.

The Intel NUC has a 2.5-inch hard-drive bay unused by now, but I should eventually add an extra SSD there — perhaps, a small and fast SSD dedicated to Lightroom previews and video editing.

At this time, I'm using Backblaze for backups, but for this setup, I'll be moving off the cloud as this will be way cheaper (break-even is months), and I'll be under the control of my data.

<div id="amzn-assoc-ad-28708ac8-a880-4aff-b5fd-2649b98d4954"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=28708ac8-a880-4aff-b5fd-2649b98d4954"></script>

## Hypervisor and operating systems

I decided to use [VMware ESXi](https://www.vmware.com/products/esxi-and-esx.html), a type-1 or native [hypervisor](https://en.wikipedia.org/wiki/Hypervisor), to be able to run multiple operating systems on my new Intel NUC. ESXi is free for personal use, and you can download it from VMware's website after registering for an account and getting a trial license.

[![If you plug a display on your machine, all you'll see is the ESXi Direct Console User Interface.](/img/posts/homelab/dcui_small.png)](/img/posts/homelab/dcui.png)

*If you plug a display on your machine, all you'll see is the ESXi Direct Console User Interface.*

ESXi is a host for virtual machines, instead of an operating system for you to use directly. The only time you need to use a display on your NUC is when configuring it for the first time or if you can't access the web interface for some reason.

Note: while you can passthrough your video card control to your guest operating system, you can't passthrough your mouse and keyboard. If you want your virtual machine to control them, you'll have to passthrough an entire USB controller. Unfortunately, you won't find a slot for an extra [PCI USB controller](https://amzn.to/2VgZgud) in such a small computer.

[![You can control a few essential settings and do some troubleshooting from the DCUI, if necessary.](/img/posts/homelab/dcui-settings_small.png)](/img/posts/homelab/dcui-settings.png)

*You can control a few essential settings and do some troubleshooting from the DCUI, if necessary.*

You also have access to a shell that resembles a regular Unix system, but it [isn't one](https://blogs.vmware.com/vsphere/2013/06/its-a-unix-system-i-know-this.html). It's not a Unix-like system, and is only partially POSIX compliant.

[![Shell access is seldom required, and using it to manage your system is frowned upon. You want to keep SSH disabled.](/img/posts/homelab/ssh_small.png)](/img/posts/homelab/ssh.png)

*Shell access is seldom required, and using it to manage your system is frowned upon. You want to keep SSH disabled.*

## Web client

[![VMware ESXi 7 web client showing a virtual machine console running [Plan 9](https://plan9.io/plan9/) from Bell Labs](/img/posts/homelab/esxi_small.png)](/img/posts/homelab/esxi.png)

*VMware ESXi 7 web client showing a virtual machine console running [Plan 9](https://plan9.io/plan9/) from Bell Labs*

You should download the hypervisor and run on your server on a dedicated storage unit that can be something as simple as a USB pen-drive. You cannot install it on a device along with any other operating system. It's going to partition the disk destructively. The installation process is pretty-straightforward if everything works fine.

### Using HTTPS certificates

[![Replace the built-in certificate with one of your own.](/img/posts/homelab/gen-certificate_small.png)](/img/posts/homelab/gen-certificate.png)

*Replace the built-in certificate with one of your own.*

I used [Filippo Valsorda](https://filippo.io)'s [mkcert](https://mkcert.dev/) program to generate a HTTPS certificate to use on the hypervisor and on my [FreeNAS](https://www.freenas.org/) instance.

For me, it doesn't matter if we are talking about servers on the local network. If you are using it seriously, you should consider doing the same regardless if you're going to be accessing it from inside your own home or not. Don't trust your network.

Please notice I use a .dev TLD here, and modern browsers force .dev domains to use HTTPS — what is interesting because it will mitigate some security risks by avoiding sending data over plain HTTP.

## Back to Intel NUC 10

VMware doesn't list the Intel NUC line on its [compatibility guide](https://www.vmware.com/resources/compatibility/search.php). However, many people have a positive experience running ESXi with it. I was confident everything would work out of the box, and only during installation, I discovered the [Ethernet NIC](https://www.vmware.com/resources/compatibility/search.php) wasn't working. This hardware was released in Q4/2019 and used a new network card.

I lost some time trying to solve this issue, and here is how I fixed it:

1. Use a second (maybe virtual?) machine with Internet connection working fine to install ESXi on your boot device.

1. Open a shell or an SSH connection and execute the commands below.

1. Move the boot device to your Intel NUC 10th generation.

Thanks to William Lam and Andrew Roderos for the articles and help figuring out what was going on.

```sh
$ esxcli network firewall ruleset set -e true -r httpClient
$ wget https://download3.vmware.com/software/vmw-tools/ESXi670-NE1000-32543355-offline_bundle-15486963.zip
$ esxcli software vib install -d /ESXi670-NE1000–32543355-offline_bundle-15486963.zip
```

You can run esxcli with the **--dry-run** flag to see what changes it'll make on your system. Be aware that the drivers versioning is a little confusing.

[![Perhaps this one would work out of the box? Unfortunately, I lost mine.](/img/posts/homelab/thunderbolt-nic_small.jpg)](/imsr/posts/homelab/thunderbolt-nic.jpg)

*Perhaps this one would work out of the box? Unfortunately, I lost mine.*

While I was at it, I bought an Ethernet dongle. It didn't work out of the box either, and I had to install the NIC Fling.

To manage the network with it, you need some initialization scripts changes. See [USB Network Native Driver for ESXi](https://flings.vmware.com/usb-network-native-driver-for-esxi) for details.

### Virtual switch
[![Virtual switch screenshot](/img/posts/homelab/virtual-switch_small.png)](/img/posts/homelab/virtual-switch.png)

#### Fail-over
An exciting surprise is that by using two NICs, I've now, out of the box, fail-over. You can check your virtual switch settings for more options and play with active/standby possibilities.

```sh
[root@homelab:~] esxcfg-nics --list
Name    PCI             Driver  Link    Speed       Duplex MAC Address       MTU    Description
vmnic0  0000:00:1f.6    ne1000  Up      1000Mbps    Full   1c:69:7a:61:30:c0 1500   Intel Corporation Ethernet Connection (10) 1219-V
vusb0   Pseudo          uether  Up      1000Mbps    Full   00:90:9e:9d:a6:7a 1500   ASIX Elec. Corp. AX88179
```

See also:
`$ esxcli network list`

### More Ethernet issues

I'm using my cablemodem Ethernet ports to connect everything. I experienced the following problems so far:

* Sometimes I'm getting only 100Mbps when using my cablemodem's CAT5e cable (on any interface).

* The cheap Ethernet dongle works on my server and my Macbook. However, if I use the dongle on my Macbook, the NUC becomes unreachable.

* Once, the server was unreachable when using both NICs. I removed the USB NIC from the server, immediately everything worked (through the internal NIC).

I also had some issues reserving IP addresses on my cablemodem admin page, so I suspect the routing problem might be on the modem. I don't have a way to test for it, though.

I don't know when I'll be getting a new Ethernet dongle later for my laptop, though, because my 5GHz Wi-Fi connection is good enough so far. If I get it at all will be to avoid losing packets and for lower jitter and latency (for better NAS experience).

## Running virtual machines

[![Web client showing multiple virtual machines](/img/posts/homelab/vms_small.png)](/img/posts/homelab/vms.png)

You can create or register a VM on the web client or with the vCenter server, but I haven't tried it.

If you are running only a single ESXi host with a few guests (operating system on your virtual machines), I find that using the web client is pretty straight-forward. However, integration with VMware Fusion Pro 11 is also possible — and albeit a little limited, it provides a great user experience.

[![VMware Fusion screenshot](/img/posts/homelab/fusion_small.png)](/img/posts/homelab/fusion.png)

I was able to move the virtual machines I had on my laptop to ESXi drag & dropping. I still need to migrate a Ubuntu image not listed there that I used to use a decade ago or so :)

[![Uploading VM screenshot](/img/posts/homelab/uploading-vm_small.png)](/img/posts/homelab/uploading-vm.png)

## Operating systems

### FreeBSD

[FreeBSD](https://www.freebsd.org/) is an open-source operating system liked for its robustness and security model. I want to learn more about it, so I'll be trying to use it first instead of any other operating systems for deploying my projects — and given [Netflix's positive experience with it](https://papers.freebsd.org/2019/FOSDEM/looney-Netflix_and_FreeBSD.files/FOSDEM_2019_Netflix_and_FreeBSD.pdf), I know it's a great choice.

![FreeBSD boot screenshot](/img/posts/homelab/freebsd.png)

FreeBSD supports installation of [packages and ports](https://www.freshports.org/faq.php#port). Packages are compiled ports. Installing port for the first time in FreeBSD:

`$ portsnap fetch extract`

You might want to read [Using ports](https://www.freebsd.org/doc/en_US.ISO8859-1/books/handbook/ports-using.html) for more details.

To use [Tailscale](https://tailscale.com) with FreeBSD, you can do this:

`$ cd /usr/ports/security/tailscale/ && make install clean && kg install tailscale`

Once installed, you want to run it as a service:

`$ service tailscaled enable`

and start it with:

`$ service tailscaled start`

### FreeNAS
[![FreeNAS screenshot](/img/posts/homelab/freenas_small.png)](/img/posts/homelab/freenas.png)

I'll be using FreeNAS as if my external 8TB hard-drive was my network-attached storage appliance. It supports sharing your data using multiple protocols, including, but not limited to, Microsoft's [SMB](https://en.wikipedia.org/wiki/Server_Message_Block), Apple's [deprecated](https://eclecticlight.co/2019/12/09/can-you-still-use-afp-sharing/) [AFP](https://en.wikipedia.org/wiki/Apple_Filing_Protocol) ([announcement](https://support.apple.com/en-us/HT207828)), Sun's [NFS](https://en.wikipedia.org/wiki/Network_File_System), [rsync](http://en.wikipedia.org/wiki/rsync), and [SSH](https://en.wikipedia.org/wiki/Secure_Shell).

There's also the possibility of mounting network devices using [iSCSI](http://en.wikipedia.org/wiki/iSCSI). You might want to read more about the [difference between NAS and SAN](https://www.backblaze.com/blog/whats-the-diff-nas-vs-san/). Still, when you set up SAN with iSCSI, it means you'll be using a remote block device. When you set up a NAS, instead of controlling a storage device, you'll be accessing the files remote filesystem.

**Importing data:** I'm having a hard time reading my encrypted [LaCie Rugged USB-C 2 TB](https://amzn.to/3ey2fGm) disk formatted with [Apple File System](https://en.wikipedia.org/wiki/Apple_File_System) (APFS).

I'm using:

`rsync -auv /Volumes/media nas.henvic.dev:/mnt/henvic/media`

to transfer my files to the NAS — using SSH as the transport layer, so no need to run the rsync service on FreeNAS.

I tried to tweak this my settings a bit without much success. I gained about 1MB/s by passing the flag:

`-e "ssh -T -c aes128-gcm@openssh.com -o Compression=no"`

to rsync to tweak the SSH connection and the encryption cipher.

List the SSH ciphers available on your system by running ssh `-Q` cipher. For the sake of security, **I recommend against messing with this, though.**

* FreeNAS doesn't recognize APFS, and I decided that using my Macbook to migrate the data was safer.

* I'm leaving my MacBook turned on during nights to upload the data to the NAS. I avoid coffee, but my MacBook uses [caffeine](http://lightheadsw.com/caffeine/) to avoid sleeping.

* The transfer rate between my Macbook and the NAS is a little slower than it should, I think. Perhaps decrypting the data from disk or the SSH encryption  is getting in the way?

[![Storage transfer rate](/img/posts/homelab/storage_small.png)](/img/posts/homelab/storage.png)

[![[iStat Menus](https://bjango.com/mac/istatmenus/) showing how my transfer rate is ten times or so slower than it theoretically should be in an ideal world.](/img/posts/homelab/network_small.png)]((/img/posts/homelab/network.png))

*[iStat Menus](https://bjango.com/mac/istatmenus/) showing how my transfer rate is ten times or so slower than it theoretically should be in an ideal world.*

**Day-to-day:** I'm going to use SMB since APFS was deprecated (I only discovered this by chance).

FreeNAS also comes with a few useful services that you can enable and use, like [MinIO](https://min.io) — a cloud storage server that is API compatible with [Amazon S3](https://aws.amazon.com/s3/).

Other plugins include:

* A few to automate backing up on services or other NAS.

* Self-hosted git service using [Gogs](https://gogs.io).

* [Emby](https://emby.media) and [Plex](https://www.plex.tv) media players.

* [Zone Minder](https://zoneminder.com) is an opensource video surveilance (CCTV) software system.

* [Transmission](https://transmissionbt.com), BitTorrent client.

* [Quassel IRC](https://quassel-irc.org) bouncer

FreeBSD [Jails](https://www.ixsystems.com/documentation/freenas/11.3-U2/jails.html) are used as the basis of these plugins.

**Remote access:** I plan to run [TailScale](https://tailscale.com) on a Jail to access my files remotely when I'm not at home. I wish FreeNAS came with a plugin for it.

### Clear Linux OS

Clear Linux is one of the best Linux distributions out there. It is created and maintained by Intel and has excellent hardware compatibility, [and it's the best in the benchmarks](https://arstechnica.com/gadgets/2020/02/linux-distro-review-intels-own-clear-linux-os/). It might as well be one of the best Linux distribution I have ever used in regards to attention to detail.

It is well-documented and follows a functional style for software installation, too, with the concept of bundles — making your life supposedly easier.

[![[Clear Linux](https://clearlinux.org) is a Linux distribution created by Intel.](/img/posts/homelab/linux_small.png)](/img/posts/homelab/linux.png)

*[Clear Linux](https://clearlinux.org) is a Linux distribution created by Intel.*

I'll be using it to run a Kubernetes cluster and get rid of Docker on my main MacBook Pro. [Adieu, Docker.qcow2](https://github.com/docker/for-mac/issues/371)! 🐄

Installing [Docker](https://docs.01.org/clearlinux/latest/tutorials/docker.html) and [Kubernetes](https://docs.01.org/clearlinux/latest/tutorials/kubernetes.html):

`$ sudo swupd bundle-add cloud-native-basic bundle-add cloud-control containers-basic`

You'll probably want to read a bit about [Kata Containers](https://katacontainers.io) and the [cri-o](https://cri-o.io) runtimes.

Perhaps I'll install [TetriNET](https://en.wikipedia.org/wiki/TetriNET) and [Frozen Bubble](http://www.frozen-bubble.org/) on it too (on a second VM for the sake of isolation?).

### OpenIndiana

[OpenIndiana](https://www.openindiana.org) is a i[llumos](https://illumos.org) distribution. Illumos is an opensource operating system based in OpenSolaris, which itself is based in Solaris.

[![OpenIndiana screenshot](/img/posts/homelab/openindiana_small.png)](/img/posts/homelab/openindiana.png)

Solaris is an interesting operating system with great contributions to the opensource ecosystem such as [ZFS](https://en.wikipedia.org/wiki/ZFS) (which is also used by FreeBSD and FreeNAS).

### Haiku

[Haiku](https://www.haiku-os.org) is an opensource operating system compatible with the discontinued [BeOS](https://en.wikipedia.org/wiki/BeOS). BeOS [could have been](https://discuss.haiku-os.org/t/history-why-apple-chose-next-instead-of-beos/7133) what [NeXTSTEP](https://en.wikipedia.org/wiki/NeXTSTEP) became to Apple: a replacement for the [Classic Mac OS](https://en.wikipedia.org/wiki/Classic_Mac_OS).

[![Haiku](/img/posts/homelab/haiku_small.png)](/img/posts/homelab/haiku.png)

### TempleOS

Out of curiosity, I installed the psychedelic [TempleOS](https://en.wikipedia.org/wiki/TempleOS). An operating system created by [Terry A. Davis](https://en.wikipedia.org/wiki/Terry_A._Davis), who had schizophrenia. I had mostly no idea what I'd found. If you “take the tour” and choose to run the TempleOS Test Suite, this is what you see.

<div class="grid-x">
  <div class="cell auto large-8">
    <div class="responsive-embed">
      <iframe width="560" height="315" src="https://www.youtube.com/embed/WkA0dMudUG0" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    </div>
    <a href="https://www.youtube.com/watch?v=WkA0dMudUG0">TempleOS Test Suite</a> video on YouTube.
  </div>
</div>

**Warning:** Risk of epileptic seizures! Run this video or the operating system at your own risk!

### CPU, memory, and storage for each virtual machine

I'm still playing with fine tunning it, but for my use case, there is not much secret. I don't have a reason to have all virtual machines up at once.

[![CPU usage: it shows 12 CPUs instead of 6 because of [hyper-threading](https://en.wikipedia.org/wiki/Hyper-threading).](/img/posts/homelab/cpu-usage_small.png)](/img/posts/homelab/cpu-usage.png)

*CPU usage: it shows 12 CPUs instead of 6 because of [hyper-threading](https://en.wikipedia.org/wiki/Hyper-threading).*

### Fuchsia

I tried to run Fuchsia on ESXi, but the system never booted. It [seems to work fine](https://fuchsia.dev/fuchsia-src/development/hardware/intel_nuc) with the older Intel NUC 6 and 7, though, so I suppose it is doable with a little more effort. I had more luck just running it with [qemu](https://www.qemu.org) instead.

### Virtual machines and emulators

It's important to remind that VMware ESXi is a hypervisor that you can use to run [x86](https://en.wikipedia.org/wiki/X86) and [x86–64](https://en.wikipedia.org/wiki/X86-64) virtual machines on 64-bit PC hardware. It's not an emulator for different architectures, so it's not possible to run [AmigaOS](https://www.amigaos.net/) or [Classic Mac OS](https://en.wikipedia.org/wiki/Classic_Mac_OS) on it.

I got excited about running different things and tried to run [MorphOS](https://morphos-team.net/) without realizing it is designed for [PowerPC](https://en.wikipedia.org/wiki/PowerPC) (PPC). I only discovered this after it never booted up ¯\\\_(ツ)\_/¯.

## Accessing from the outside world

I give each of my virtual machine a hostname that I can use to access them internally. I use only a single level of subdomain to keep things tidy,  and to avoid any implication regarding HTTP cookie management.

I decided to use [Tailscale](https://tailscale.com) to expose each node individually. However, I can't install it on the host (ESXi), but this is fine because if I ever need access to it, I can always use something like:

`ssh -L 127.0.0.1:443:homelab.henvic.dev:443 gateway.henvic.dev`

and access the admin web client as usual on my web browser, pointing it to [https://localhost](https://localhost).

## Virtual Private Network (VPN)

Despite being able to access it remotely through TailScale (which uses [WireGuard](https://en.wikipedia.org/wiki/WireGuard) behind the scenes), sometimes, I may want to access my home network directly. I might want to do this to bypass IP geolocation verification if I'm traveling or to rely on a secure termination point if I'm accessing the web through a suspicious access point and want to mitigate the risk of a [man-in-the-middle attack](https://en.wikipedia.org/wiki/Man-in-the-middle_attack).

### OpenVPN

[OpenVPN](https://openvpn.net/) is a great option to solve this. I installed its appliance image on ESXi, and it worked out of the box like a charm.

[![OpenVPN web client)](/img/posts/homelab/openvpn_small.png)](/img/posts/homelab/openvpn.png)

[![You can install OpenVPN Connect on your computer or smartphone for easy installation.)](/img/posts/homelab/openvpn-connect_small.png)](/img/posts/homelab/openvpn-connect.png)

*You can install OpenVPN Connect on your computer or smartphone for easy connection.*

Now, all I need to use it from outside home is to expose it through the web. I don't have a static IP address at home, so this would be something else to solve too. In theory, I could always use Tailscale there also (for VPN inception).

## Protecting your system

I recommend you use SSH with [public-key cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography) (with the [ed25519](https://ed25519.cr.yp.to) cipher) to access everything you can, and disable password authentication or use it as a second-factor.

If you plan on having multiple systems, at first it might sound like a good idea to share your private key between them to avoid the hassle of changing files between machines, etc. However, you might want to limit the surface attack by generating a new key  with:

`$ ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -C <comment>`

and sharing them between machines (at least the ones you don't trust too much, or if you share with other people). Maybe you want to read [how to set up Multi-Factor Authentication for SSH](https://www.digitalocean.com/community/tutorials/how-to-set-up-multi-factor-authentication-for-ssh-on-ubuntu-16-04) using the [google-authenticator-libpam](https://github.com/google/google-authenticator-libpam).

## Ideas

### Importing data from the cloud

I want to use [Perkeep](https://perkeep.org) to import data from the cloud. I've photos and videos in both Google Photos and iCloud that I need to safely store on premises and off-site because we never know when a cloud provider might go away or, in practice, hijack your data — even when unwilling.

[Perkeep](https://perkeep.org) is a set of open source formats, protocols, and software for modeling, storing, searching, sharing, and syncronizing data. I'll be using it along with [gphotos-cdp](https://github.com/perkeep/gphotos-cdp) to download my data from Google Photos for safekeeping.

### Continuous Integration

I work with software development, and for the last five years or so, I use most the [Go programming language](https://golang.org/), a language known to have fast tests and build times.

However, sometimes a project might have an architecture that takes too long to test, or you might find differences between operating systems that are usually only discovered on a CI. So, perhaps running [Buildkite](https://buildkite.com/) on the Intel NUC would be a great idea. Another possibility is to write a small program that would run tests on different machines and output the results.

### Backup & restore plan

One common mistake many people make is never testing their backups. I'm guilt of this myself.

Back to around 2010, I had just started using OS X and relied on FileVault to back up my data. However, it only backed up data when you logged out your machine, but I just never logged out. I was tired from working overnight on this project of my own and removed a directory with rm -rf . It was the directory where I kept all my projects, and I wasn't using git + GitHub back then, but a local copy of [Mercurial](https://www.mercurial-scm.org) without any remotes. I freaked out, but then I remembered I had just renamed the directory minutes before.

Then a couple of years ago, I was unfortunate enough to have my backpack stolen from a beach house in California. My TimeMachine backup disk (cold) was in Brazil, and I discovered the hard way [Backblaze](https://secure.backblaze.com/r/01v1py) backup didn't restore permissions correctly. My operating system was so screwed up due to wrong permissions that I asked my parents to mail me my backup disk (and I've to admit they sent it without making a backup first), even though my data loss was minimal thanks to Backblaze.

An excellent backup & restore plan should account for things like this from happening, and consider things such as:

* Cold backup storage on multiple sites (aware of distance and jurisdiction), and syncing plan.

* Hot backup.

* Logical and physical security.

* Recovery strategy (dry-run!) and data validation.

### Processing photos and videos

Ideally, I'd love to offload processing photos and videos to my Intel NUC. I have a [Canon 6D Mark II](https://en.wikipedia.org/wiki/Canon_EOS_6D_Mark_II) and a [DJI Phantom Pro](https://www.dji.com/nl/phantom-4-pro), and I enjoy [taking pictures and sharing them](https://www.flickr.com/photos/henriquev/). Still, I keep slacking in editing and publishing them because I used to keep them on an external drive, and, as silly as it is, I hate to have a USB drive connected to my laptop because I know I'll disconnect the cable every dozen minutes or so. My computer is powerful, but my internal storage is minimal.

I use [Adobe Lightroom Classic](https://www.adobe.com/products/photoshop-lightroom.html) to manage my photo and video collections and to edit photos. Perhaps I'll move to something like [PhotoPrism](https://github.com/photoprism/photoprism) for managing my collection as I'd prefer something that had a client/server architecture, and to use Adobe Photoshop for editing.

Processing offloading would be the best option. For photos, it would mean the server could be responsible for things like AI tagging or batch tasks such as creating previews. I wonder why Lightroom doesn't have this. For videos, offloading is way more common. If I had a Mac mini instead of an Intel NUC, I could even run macOS as a guest operating system to help with this task a little, as the Adobe suite isn't available for Linux.

### ToR

[ToR](https://www.torproject.org) is a fundamental piece of the Internet ecosystem. It allows people to browse the Internet with privacy — essential to liberty and a way to avoid persecution from the state (especially important in conflict zones).

You can help the ToR project by running a [Tor relay](https://trac.torproject.org/projects/tor/wiki/TorRelayGuide) to relay traffic inside the network, or you can even set up an exit node so people can use your Internet connection as the “exit” point to the “mundane” Internet.

You can also host an onion service (formerly known as “hidden service”).

[![Using the [Tor Browser](https://www.torproject.org/download/). Facebook is one of the few major websites that has an [onion service](https://en.wikipedia.org/wiki/List_of_Tor_onion_services).](/img/posts/homelab/tor_small.png)](/img/posts/homelab/tor.png)

*Using the [Tor Browser](https://www.torproject.org/download/) to access Facebook, one of the few major websites that has an [onion service](https://en.wikipedia.org/wiki/List_of_Tor_onion_services) — meaning no exit nodes as the connection is routed end-to-end through the Tor network.*

### EternalTerminal

One problem I've with SSH is that it's easy to close my laptop lid and to lose a connection. Usually, this is not a problem, but eventually, I happen to be doing some heavy-lifting on a server that will take a while, and I get stuck with it. You can use tmux to avoid this, but why not [EternalTerminal](https://github.com/MisterTea/EternalTerminal)? I need to configure this in my primary virtual machines (FreeBSD and Linux).

## Reviews from others

* [Andrew Ronderos: VMware ESXi Home Lab — Intel NUC 10 (Frost Canyon)](https://andrewroderos.com/vmware-esxi-home-lab-intel-nuc-frost-canyon/)

* [Virten.net: Homelab — Will ESXi 7.0 run on Intel NUC?](https://www.virten.net/2020/04/homelab-will-esxi-7-0-run-on-intel-nuc/)

* [Virten.net: ESXi on 10th Gen Intel NUC (Comet Lake — Frost Canyon)](https://www.virten.net/2020/03/esxi-on-10th-gen-intel-nuc-comet-lake-frost-canyon/)

* [virtuallyGhetto: ESXi on 10th Gen Intel NUC (Frost Canyon)](https://www.virtuallyghetto.com/2020/01/esxi-on-10th-gen-intel-nuc-frost-canyon.html)

## Thanks

I hope you enjoyed this blog post. If you have any tips for me, interesting stuff to share or just want to talk about this article, please reach out.

<div id="amzn-assoc-ad-d9bdb317-57e8-4c0b-9651-1cdfe5cf0657"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=d9bdb317-57e8-4c0b-9651-1cdfe5cf0657"></script>

If you click and buy any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/).

{{< tweet 1252360422683897856 >}}
