# Boonterm board


# FLASH NEW Firmware

### REF

https://rd.forth.co.th/projects/boonterm-team/wiki/MBT-XX-XX


``` bash
use sdcard p''off

ตั้งค่า DIP <<<Switch DIP>>> SD Card
# in boot loader
setenv bootargs_mmc 'setenv bootargs ${bootargs} root=/dev/mmcblk1p1 rootwait'
setenv bootcmd_mmc 'run bootargs_mmc; mmc dev 1; mmc read ${loadaddr} 0x800 0x2000; bootm'
saveenv 
boot

# in linux

mkdir /media/usb
mount /dev/sda1 /media/usb
cd /media/usb
bash flash.sh


<<<Switch DIP>>> EMMC

reboot

mkdir /media/usb
mount /dev/sda1 /media/usb
cd /media/usb
bash initial.sh
```

# Set up board

```bash
=========================================
useradd likhit
usermod -aG sudo likhit

ifconfig eth0 192.168.10.123 netmask 255.255.254.0

route add default gw 192.168.10.1
mkdir /media/usb
mount /dev/sda1 /media/usb
date -s "15 MAY 2018 13:20:49"
=====================

wget http://192.168.10.16/libusb-1.0.9.tar.bz2
wget http://192.168.10.16/pcsc-lite-1.8.22.tar.bz2
wget http://192.168.10.16/ccid-1.4.23.tar.bz2

```
## INSTALL LIBRARY
1. libusb-1.0.9 [ wget http://192.168.10.16/libusb-1.0.9.tar.bz2 ]
2. pcsc-lite-1.8.22 [ wget http://192.168.10.16/pcsc-lite-1.8.22.tar.bz2 ]
3. ccid-1.4.23 [ wget http://192.168.10.16/ccid-1.4.23.tar.bz2 ]

### LIBUSB

```
https://github.com/libusb/libusb/wiki
https://sourceforge.net/projects/libusb/files/libusb-1.0/

    Version 1.0.9
    ./configure
    make
    make install

    cd /
    apt-cache policy libusb-1.0*
    libusb-1.0-0:

```

### PCSC

```
    Version 1.8.22
    ./configure
    make
    make install
    # Run service
    pcscd --foreground --debug --apdu --color | tee log1.txt
```

### CCID 

```
    Version 1.4.23
    ./configure
    make
    make install
```

# Developer Reference

https://github.com/golang/go/wiki/GoArm
https://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi

https://muscle.apdu.fr/
1.  instatll pcsc-lite
	https://pcsclite.apdu.fr/
	http://tech.springcard.com/guides/pcsc-unix-with-pcsclite/
2. install ccid driver 
	https://ccid.apdu.fr/

http://tech.springcard.com/guides/pcsc-unix-with-pcsclite/

http://www.somkiat.cc/learn-smart-card-apdu/
http://www.somkiat.cc/read-data-form-sim-card/
http://www.blackhat.com/presentations/bh-usa-08/Buetler/BH_US_08_Buetler_SmartCard_APDU_Analysis_V1_0_2.pdf


http://rebelsimcard.com/sim-commands.html

==================================================

## pySim Tool for test read sim card

```
git clone git://git.osmocom.org/pysim.git
apt-get install swig
apt-get install python2.7-dev
apt-get install python-pip
wget http://192.168.10.16/pyscard-1.9.6.tar.gz

pip install --global-option=build_ext --global-option="-I/usr/local/include/PCSC:/usr/include/python2.7 -lpython2.7" pyscard-1.9.6.tar.gz

./pySim-read.py --pcsc-device=0

```
