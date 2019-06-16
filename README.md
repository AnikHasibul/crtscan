# CRTScan

Scan subdomains from certificate transparency logs (https://crt.sh)

## Installation

### Download binary

Download the binary based on your OS and architecture from the [release](https://github.com/AnikHasibul/crtscan/releases) page.

### From command line

```bash
bash -c 'wget https://github.com/AnikHasibul/crtscan/releases/download/v1.0.0/crtscan_1.0.0_Linux_x86_64.tar.gz; tar -xvzf crtscan_1.0.0_Linux_x86_64.tar.gz; sudo mv crtscan /usr/local/bin/crtscan'
```

### From source

To install/build from source with [Golang](https://golang.org):

```bash
go get github.com/AnikHasibul/crtscan
```

## Usage

```txt
$ crtscan google.com
```

Or with wildcard search:

```txt
$ crtscan '%.google.com'
```

If you want to write the output into file

```txt
$ crtscan '%.google.com' > output.txt
```

TO perform a visual recon/inspection with [aquatone](https://github.com/michenriksen/aquatone):

```txt
$ crtscan '%.google.com' | aquatone
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)