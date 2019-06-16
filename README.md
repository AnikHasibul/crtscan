# CRTScan

Scan subdomains from certificate transparency logs (https://crt.sh)

## Installation

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