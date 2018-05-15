# Super Smart Card(sscard)

Super Smart Card API on top of pcsc.

## TODO

``` bash
# Linux: install pcsc library
sudo apt-get install pcscd

# run pcscd 
sudo LIBCCID_ifdLogLevel=0x000F pcscd --foreground --apdu --color | tee log.txt

# goget
go get -u github.com/Napat/sscard/sscard

# go build
go build -o sscard github.com/Napat/sscard/main

./sscard  # ./sscard.exe on windows

# go run hack
go run $(find ./ | grep ./main/)

```

About requirement and other platform see: `docs/INSTALLATION_xxx.md`

## Platforms

- Windows 10
- Linux / Ubuntu

## Reference

- [PCSC in golang](https://ludovicrousseau.blogspot.fr/2016/09/pcsc-sample-in-go.html)
- [APDU command for Thai ID card](https://github.com/Napat/ThaiNationalIDCard/blob/master/APDU.md)
- [APDU RESP List](https://www.eftlab.co.uk/index.php/site-map/knowledge-base/118-apdu-response-list)
