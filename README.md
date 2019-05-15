# sigtool
This a sigtool in Go for signed PE files.
Currently, only extracting digital signatures embeded in a PKCS#7 `SignedData` in a signed PE is supported.
Adding and deleting digital signatures will be supported soon.


## Command Line Usage
Example sigtool CLI usage:

	go get github.com/doowon/sigitool
	sigtool test.exe
