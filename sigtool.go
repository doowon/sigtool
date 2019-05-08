package sigtool

import (
	"debug/pe"
	"os"
)

// ExtractDigitalSignature extracts a digital signature specified in a signed PE file.
// It returns a digital signature (pkcs#7) in bytes.
func ExtractDigitalSignature(filePath string) (buf []byte, err error) {
	pefile, err := pe.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer pefile.Close()

	var vAddr uint32
	var size uint32
	switch t := pefile.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		vAddr = t.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_SECURITY].VirtualAddress
		size = t.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_SECURITY].Size
	case *pe.OptionalHeader64:
		vAddr = t.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_SECURITY].VirtualAddress
		size = t.DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_SECURITY].Size
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf = make([]byte, int64(size))
	f.ReadAt(buf, int64(vAddr+8))

	return buf, nil
}
