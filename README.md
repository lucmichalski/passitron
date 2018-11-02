[![Build Status](https://img.shields.io/travis/function61/pi-security-module.svg?style=for-the-badge)](https://travis-ci.org/function61/pi-security-module)
[![Download](https://img.shields.io/bintray/v/function61/pi-security-module/main.svg?style=for-the-badge&label=Download)](https://bintray.com/function61/pi-security-module/main/_latestVersion#files)

What is this?
-------------

Software for a separate trusted hardware device ("hardware security module") which
essentially acts just like [Keepass](http://keepass.info/) and only serves the
function of storing secrets.

If you use Keepass on your PC and your PC gets compromised by a virus or a hacker,
it's game over.

If you use a separate device for storing secrets, your PC compromise does not
expose your secrets. This software only exposes your secret when you physically
press a button on the device - and only exposes one secret per push acknowledge.


Links
-----

- [Developer docs](https://docs.function61.com/pi-security-module)


Features
--------

- Supported secrets:
	* Passwords
	* OTP tokens (Google Authenticator)
	* SSH keys (via SSH agent protocol)
	* Keylists (["printed OTP list"](https://en.wikipedia.org/wiki/One-time_password#Hardcopy))
- Create, view and list secrets in a folder hierarchy.
- Export database to Keepass format (for viewing in mobile devices when traveling etc.)
- Import data from Keepass format


Recommended hardware
--------------------

![](docs/pi-zero-in-wood-case.png)

I'm using [Raspberry Zero W](https://www.raspberrypi.org/products/pi-zero-w/)
with [wooden case](https://thepihut.com/products/zebra-zero-for-raspberry-pi-zero-wood)
and a [capacitive pushbutton](http://www.ebay.com/sch/?_nkw=ttp223).

It doesn't matter much which hardware you use, as long as you don't run anything else on
that system - to minimize the attack surface. For such a light use Raspberry Pi is
economical, although this project runs across processor architectures and operating systems
because Golang is so awesome. :)


Download & running
------------------

Click the "Download" badge at top of this readme and locate the binary for your OS/arch combo:

- For Raspberry Pi, download `pism_linux-arm`
- For Linux PC, download `pism_linux-amd64`

Note: don't worry about `public.tar.gz` - it's downloaded automatically if it doesn't exist.

Rename the downloaded binary to `pism`.

Pro-tip: you can download this directly to your Pi from command line:

```
$ mkdir pi-security-module/
$ cd pi-security-module
$ curl --fail --location -o pism <url to pism_linux-arm from Bintray>

# mark the binary as executable
$ chmod +x pism
```

Installation & running:

```
# this only writes a Systemd unit file
$ sudo ./pism install

# configure Systemd to start this on system reboot
$ sudo systemctl enable pi-security-module

# start it manually (reboot would have the same effect)
$ sudo systemctl start pi-security-module

# check logs for succesful start
$ sudo systemctl status pi-security-module
Mar 19 12:44:52 raspberrypi systemd[1]: Started Pi security module.
Mar 19 12:44:52 raspberrypi pism[15626]: 2018/03/19 12:44:52 extractPublicFiles: extracting public files from public.tar.gz
Mar 19 12:44:53 raspberrypi pism[15626]: 2018/03/19 12:44:53 Starting in port 80
```

Looks good. You should now be able to access the web interface at `http://<ip of your pi>`.


How to build & develop
----------------------

[How to build & develop](https://github.com/function61/turbobob/blob/master/docs/external-how-to-build-and-dev.md)
(with Turbo Bob, our build tool). It's easy and simple!


TODO
----

- Tags to .JS command definitions
- Enter to confirm command dialog
- Data types for command fields (password)
