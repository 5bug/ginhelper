// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/conf/conf.go
// pkg/rest/jwt.go
// pkg/rest/page.go
// pkg/rest/rest.go
// pkg/rest/server.go
// pkg/storage/storage.go
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pkgConfConfGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x51\xd1\x8a\xd5\x30\x10\x7d\xce\x7c\xc5\x38\xb0\x92\x40\x69\xc1\xc7\x2b\x7d\x54\x10\x54\x44\xd1\x17\x11\x89\x69\x92\x8d\x9b\x26\x97\xa4\x5d\x90\xa5\xff\x2e\x93\xdb\xdb\xbd\x0f\xfb\xd2\x72\x26\x67\xce\x9c\x33\x73\xd6\xe6\x41\x7b\x8b\x26\x27\x07\x10\xe6\x73\x2e\x0b\x4a\x10\x64\x93\xc9\x53\x48\x7e\xf8\x5b\x73\x22\x10\xe4\xa2\xf6\xfc\x0f\x79\x08\x79\x5d\x42\x24\x00\x41\x3e\x2c\xf7\xeb\x9f\xde\xe4\x79\xa8\xa1\xac\xe7\x6a\xd3\x10\xb3\x2f\x6b\x25\x10\xff\xf4\x1c\x91\x7c\x3e\x3f\xf8\x3e\xa4\x81\x61\xff\xf8\x86\x40\x01\x0c\x03\x7e\xcc\x7a\xc2\xc8\x1f\x1e\x1e\x3c\xb8\x35\x99\x56\x95\x06\x43\x5a\x6c\x71\xda\xd8\xa7\x4d\xa1\xb4\xa5\xa0\x2d\x25\x17\x85\x4f\x20\x26\xeb\x6c\x41\x66\xcb\x86\xc5\xa4\x17\xdd\xe1\x6f\x3c\x8d\xc8\x6e\xfb\x4f\xba\xd4\x7b\x1d\xa5\x51\x20\xc4\xc5\x4d\xff\x21\xb9\xec\x24\x5d\x46\x21\x77\x9c\xee\x2a\x75\x58\x97\x12\x92\x97\x5c\x50\x0a\xc4\x26\x15\x88\x47\x5d\xd0\x85\x68\x3f\xeb\xd9\xee\x04\x10\x9c\xbf\xff\xd6\xc0\x0f\x5d\xe4\xeb\x2b\xa1\xc3\x5d\x94\x3a\x24\x3a\xd0\x21\x40\x6a\xef\xfd\xa2\x4b\xb5\x2c\x1f\xdc\xb3\xfa\x38\x22\x51\x0b\x71\xe3\x53\xd2\x5a\x2d\x4e\xd6\xe9\x35\x2e\xfb\x76\x58\x46\x14\xbb\xac\x25\x81\xd8\xe0\xe5\x58\x2c\xfb\xf3\xae\xfe\xa2\xee\x98\xa0\x60\x5f\x0f\xef\xf0\x34\xe2\xe5\x76\xfd\x57\xab\xa7\xf7\x21\x5a\x79\xc3\x0b\xae\x91\xc6\x11\x53\x88\xcd\xd3\xb5\x82\xed\x74\xdf\xd3\xbc\x2f\xf6\xa2\x68\xd4\xdb\xf6\xfc\xea\xb9\xe1\x6a\xeb\x1d\x1f\xeb\xb0\x75\x34\x32\xfd\x44\xcd\x0b\xc7\x39\xf2\x70\xa0\x0d\xae\x70\x83\xff\x01\x00\x00\xff\xff\xa3\x56\x20\x6e\x96\x02\x00\x00")

func pkgConfConfGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgConfConfGo,
		"pkg/conf/conf.go",
	)
}

func pkgConfConfGo() (*asset, error) {
	bytes, err := pkgConfConfGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/conf/conf.go", size: 662, mode: os.FileMode(420), modTime: time.Unix(1653144430, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgRestJwtGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xc1\x6e\x24\x35\x10\x3d\xb7\xbf\xa2\x68\xa4\xc8\x1d\xf5\xba\x21\x02\x0e\x91\xe6\x10\x85\x65\x87\x95\x32\x41\x99\x5d\xf6\x80\x38\x38\xdd\x95\x1e\x27\x3d\xf6\x60\x57\x67\x12\xa2\x39\xc3\x0d\x90\x90\xf6\x82\xb4\x5f\xc0\x0d\x71\xd8\x03\xfc\x0c\x64\xe1\x2f\x50\xb9\xa7\x3b\x33\x21\xa0\x9d\x93\xc7\x7e\x7e\xfd\xea\xd5\x2b\x2f\x74\x79\xa1\x6b\x04\x8f\x81\x84\x30\xf3\x85\xf3\x04\x52\x24\x29\x7a\xef\x7c\x48\x45\x92\x5a\xa4\x62\x46\xb4\xe0\x35\x99\x39\xa6\x42\x24\x69\x6d\x68\xd6\x9e\xaa\xd2\xcd\x8b\xaa\xf6\xe6\x5c\x37\x97\xba\x38\x5f\xd2\xa3\xda\xa5\xdb\xc7\xb5\xb1\x8f\x6a\x67\x4d\xc9\xab\x54\x64\x42\x14\x05\x1c\x2f\xd0\x6b\x72\x1e\x94\x52\x82\xae\x17\x78\xb7\x13\xc8\xb7\x25\xc1\x8d\x48\xce\x97\xa4\xa6\xa4\x6d\xa5\x7d\x75\xd8\x68\x33\x0f\x22\x99\xe8\x39\x32\xc4\xd8\x5a\xac\x22\xd5\xd1\xf5\x14\x4b\x8f\x14\xa9\x2e\xb5\xbf\xdb\x18\xc1\x17\x5f\x9e\x5e\x13\xca\xf4\xdd\xf7\xdf\xdb\xfb\x60\x8c\xcd\x02\x7d\xda\x09\x78\xe6\x2e\xd0\x3e\xbe\x5a\x18\x8f\x1f\xb7\x5e\x93\x71\x16\x70\xeb\xaf\x28\x9d\x0d\xf4\x20\x70\x04\xec\x83\x1a\xbb\xd6\xc3\x2e\xec\x45\xc2\x83\x96\x66\x87\xde\x05\x96\xa1\xc4\x59\x6b\xcb\x61\x4b\x66\x50\x1b\xab\xc6\xda\x56\x0d\xfa\x4f\xf8\xe8\x46\x24\x1e\xa9\xf5\x16\x18\x29\x4b\xd8\x65\xc4\xa1\xb3\x84\x57\x94\xf1\x71\xe2\xd6\x86\xe4\x80\xde\xc3\xfe\x08\x74\x4b\x33\xe7\xcd\xd7\x51\x83\x2c\xd5\x09\x7e\xd5\x62\x20\x35\x46\x5d\xa1\x57\x4f\x90\x64\xba\x85\x49\xb3\x4c\x24\x89\x39\x8b\x04\xef\x8c\xc0\x9a\x26\x32\x27\xa5\x3a\x38\x75\x9e\x5e\x18\x9a\x4d\x49\x53\x1b\x24\xf7\x57\x75\xeb\xe7\xb6\x27\xc1\x2a\xeb\xd0\x4f\xa7\xc7\x93\xff\x82\xe4\xb0\x73\x82\xa1\x6d\xe8\x04\x17\xcd\x75\xa4\x4f\x8e\x42\xbd\xcf\x1f\x55\x8f\x39\x44\x32\xcb\x79\x77\x15\xc9\xba\xaa\x45\x92\xac\x04\x33\x4f\x59\x74\x5f\x69\x9a\x43\xbf\xcc\xe2\xe9\x04\xaf\x48\x66\x22\x59\xad\x5b\xbd\x55\x1d\x0c\x3e\x6f\x1b\x43\xdc\xb0\x29\xf9\x75\x4a\x32\x90\x0e\x76\x8f\xb7\xcc\x8c\xd9\x8e\x2e\x9b\x33\x18\xf0\xa3\x11\xa4\x69\x34\x88\x31\xa3\x0e\x15\xd4\x04\x97\xf7\x7c\x05\x13\xc0\xb6\x4d\x93\xb2\xca\xbe\xa0\x95\x48\x22\xd3\xd0\x2e\x4e\xef\x67\xda\x07\x64\x9b\xbb\xf8\x0e\xda\x72\xd8\xe9\x15\xdd\xac\xf2\x2e\x03\xf1\x0c\x76\xf9\x5a\xcc\x5c\x06\xd2\x80\xb1\x84\xfe\x4c\x97\xc8\xb0\x6d\xe9\x7d\x80\xfa\xb8\xe7\xdc\x5f\x11\x6d\xfe\x77\xcb\x37\x54\xba\x1c\xdc\x05\x0b\x8c\x1f\x54\x9d\x32\x25\x07\x8b\xba\xfb\xee\x02\x76\x76\xd6\x90\xcf\x75\x63\xaa\xfb\x34\x6f\xe1\x91\xb1\x97\x7c\x93\x6d\x5a\x5f\xec\xda\xf8\x04\x6d\xac\x10\xde\xfc\xf8\xea\xf6\xdb\xef\x9f\xbe\x78\xd6\xf5\xb1\xdf\x97\xf6\x6e\xc8\x33\x90\xdd\x22\xdf\x28\xbd\x64\xf9\x83\x81\x22\x79\xe0\xa5\x88\x41\xec\xc6\x36\x1c\xd0\x7e\x37\xb0\x13\xb7\x94\x99\x3a\xa8\x2a\xf9\xc0\x54\x67\xea\xb9\x35\x57\x32\xcb\xa1\x28\xe0\xaf\xdf\xbf\xb9\xfd\xe9\xd5\xed\xcb\x5f\xff\x7e\xf9\x0b\x33\x7d\x1a\x42\x8b\x7e\x1f\x00\x20\xdd\x78\x4b\x72\xf8\xff\x5f\x51\xc0\x9b\x9f\x7f\xfb\xf3\xbb\x1f\xfe\x78\xfd\x9a\x33\xcf\x83\xc0\xc5\xe5\x77\x79\xe9\xa3\x32\xc1\xe5\x46\x50\x62\x41\xa6\xb6\xc6\xd6\x47\x48\x33\x57\x8d\xa7\x7b\x1f\x7e\x94\x43\x39\x78\xb9\xee\x0d\x83\xb0\x9a\x46\x8b\x64\x1f\x85\x4c\xac\xc4\x3f\x01\x00\x00\xff\xff\x72\x35\x94\x07\xd9\x05\x00\x00")

func pkgRestJwtGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgRestJwtGo,
		"pkg/rest/jwt.go",
	)
}

func pkgRestJwtGo() (*asset, error) {
	bytes, err := pkgRestJwtGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/rest/jwt.go", size: 1497, mode: os.FileMode(420), modTime: time.Unix(1653142800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgRestPageGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcd\x8e\x9b\x30\x14\x85\xd7\xf8\x29\xae\x58\x54\x30\x1a\xc1\x54\x9a\xc9\x22\x2a\xab\x74\x53\xa9\x6a\x2b\xa5\x0f\x80\x45\x4c\xe2\xc6\xb1\x91\x7d\x51\x7f\x10\xef\x5e\x5d\x83\xc1\xf9\x99\xec\x7c\xee\xf1\xc9\xb9\x9f\xe9\x78\x73\xe6\x47\x01\x56\x38\x64\x4c\x5e\x3a\x63\x11\xd2\x0b\xc7\x53\xca\x58\x59\xc2\x0f\x7e\x14\xdf\x3b\x94\x46\x43\x51\x14\x0c\xff\x76\x22\xd6\x1c\xda\xbe\x41\x18\x58\x42\xe2\x17\x7d\x10\x7f\x40\x6a\x04\x00\xa8\x5b\x63\x2f\xdb\xb4\x0b\x7a\x5a\x4f\xa6\xbd\xfc\x27\xe0\x81\x89\x74\xf2\xec\xa9\x81\xff\x39\xb4\x52\x1f\x83\xc7\x19\x8b\x69\x0d\x65\xc9\x5d\x03\xc6\xc2\x41\xb8\x86\x8d\xbe\xe4\x4f\x83\x5c\xdd\xb6\x8c\xc5\x47\x35\xe7\x0a\xf5\x2f\x67\xf4\xbb\x35\xef\x4d\xa1\xa6\x8f\xdf\x99\x5e\x23\x79\x36\xaf\xc1\x83\x8b\xbe\xb8\x28\xee\x26\x09\x83\x9e\xd6\xf3\x0e\x5f\xa5\xc3\xcf\x1c\xf9\xba\xc0\xa2\xac\xed\xa3\x9d\x58\x42\x73\x4a\x15\xb6\xe5\x8d\x18\xc6\x90\xad\xa4\xc3\x25\xf6\x9b\xf8\xbd\xe4\xb4\x52\x29\x38\x70\xe4\xac\xed\x75\x13\x8f\x32\xd3\x61\xf4\xac\xcf\xe0\xfb\x4d\x8b\x3d\x83\xba\xf9\xa3\x1c\x9e\x96\xcc\x81\x25\x94\x08\xdb\x0a\x3e\x04\x71\x18\x27\xb1\x58\x61\x57\x60\x3a\x5c\xcf\xd1\xdc\x73\x5e\xc7\x74\x9c\xa7\x11\xe1\x6a\x2a\x14\x0f\x3c\xd4\x8a\x6a\x65\xf4\xb5\x16\x3b\x21\x55\xd6\x2a\xc3\x71\xf3\x9a\x79\x77\x0e\x25\x04\x21\x8e\xcf\xf3\x7c\x0e\xf2\x04\x2b\xbf\x1f\x4b\xac\xc0\xde\xea\x89\xcf\xc4\x6e\x77\x12\xcd\xd9\xbf\x87\xe7\xe5\x21\x3d\xad\x94\xf2\xc9\x90\xe5\x04\x41\xb6\xd7\x1b\xc2\xa7\x0a\x5e\x68\x90\x5c\xcb\x15\xbc\xb0\x64\xbc\xf2\x7b\x02\x77\xf6\x99\xcb\x47\xb2\x83\x50\x4e\xc0\xdd\x15\x78\x7b\x74\xe3\x8d\xf2\x47\xf6\x3f\x00\x00\xff\xff\xef\x65\x19\x81\xd9\x03\x00\x00")

func pkgRestPageGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgRestPageGo,
		"pkg/rest/page.go",
	)
}

func pkgRestPageGo() (*asset, error) {
	bytes, err := pkgRestPageGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/rest/page.go", size: 985, mode: os.FileMode(420), modTime: time.Unix(1653142893, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgRestRestGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x55\x4d\x6f\x1b\x37\x10\x3d\x2f\x7f\xc5\x84\x40\x8b\x5d\x55\xa5\xee\x32\x74\x49\xd0\xef\xc6\x36\xe4\xa4\x67\xaf\xa8\xd9\x15\xa3\x15\x29\x93\x5c\xc5\x81\x20\xa0\xc7\xa2\x28\xd0\x1e\x72\x69\x0b\x34\x3d\xb4\xe7\xa0\x40\x0f\x05\x6a\xa0\x7f\xa6\x4e\xdd\x53\xfe\x42\x30\x24\xd7\xd2\x26\x80\x7d\x31\x97\x33\x9c\x79\xef\xf1\x71\xb4\x2e\xe5\xb2\xac\x11\x2c\x3a\xcf\x98\x5a\xad\x8d\xf5\x90\xb3\x8c\x4b\xa3\x3d\x5e\x7a\xce\x32\x5e\xad\xc2\x3f\x8b\x55\x83\x32\x2e\x5b\xed\xd5\x0a\x47\x73\x9c\xb5\x35\x6d\x38\x6f\x95\xae\x1d\x67\x2c\xe3\xb5\xf2\x8b\x76\x26\xa4\x59\x8d\x6a\xa5\x3f\xac\x8d\x56\x92\x56\xbc\x1f\x73\xca\xb6\x6b\x87\x7a\xd4\x98\xda\xb6\x8e\xb3\x82\x31\xff\x6c\x8d\x50\x2b\xfd\x20\x36\x87\x09\x0c\x6a\xa5\x45\xfa\x64\x6c\x34\x82\x2e\x94\xf0\xc1\x1c\x2b\xa5\x31\x9e\xec\x62\xce\xdb\x56\x7a\xd8\xb2\x2c\x65\xdd\x56\xc8\xf6\xc5\x59\x76\xb2\x46\x5b\x7a\x63\x61\xd0\xad\xd8\x2e\xf4\x38\xc6\xa7\x5d\x29\x8d\x4f\xbb\x56\xac\x6a\xb5\x3c\x88\xe5\xb2\x87\xae\x80\x41\x77\x88\x1a\xfb\x4b\x18\x4f\xe0\xfd\xb4\xb5\xdd\x85\x2d\xb1\x27\xd6\x21\xbb\x5f\xca\x65\x6d\x4d\xab\xe7\x79\x11\x73\x7a\xfc\x25\xcb\x54\x05\x9b\x21\x98\x25\xd5\x93\xe2\x13\xf4\x39\x37\x09\x2f\x2f\x8e\x28\xb0\x65\x59\x38\x79\x4b\x68\x02\x1b\x91\xdf\xb2\x2a\x58\xb6\x63\x99\x45\xdf\x5a\x0d\xd2\x5f\x26\x96\x53\x74\x6d\xe3\xa7\xb8\x6e\x9e\x81\x10\x22\x4a\x78\xb8\xb9\x97\x31\xee\xc2\xcc\x98\x06\xd2\xdf\xf9\x13\x67\xf4\x98\xdb\x10\xe1\xe7\x2c\x7b\xe8\x6a\xda\x57\xda\xa3\xad\x4a\x89\xdb\x5d\x97\xb3\x72\x35\x3f\x4f\x4d\xcf\xd0\xa7\x62\xd4\x32\x28\x9a\x5b\x18\x1c\xb4\x2d\xf6\x39\x39\x5a\x0b\x68\xad\xb1\x05\xa1\x50\x15\x7d\xc0\xbd\x09\x68\xd5\x04\xd2\x56\xa4\x62\x13\xa8\xca\xc6\x61\xd8\x22\x20\x13\xa8\x56\x5e\x7c\x44\x47\xab\x9c\x5f\xff\xf6\xc7\xcd\x9f\xbf\xbf\xbe\xfa\xee\xff\xe7\x3f\xde\xbc\x7c\xf9\xfa\xea\xa7\xf7\x1c\x1f\x52\xb1\x98\x92\x17\x24\x11\x60\xe3\xf0\xed\xb2\xde\xb6\x87\x55\xf9\xab\x6f\x7e\xb8\xfe\xf6\x05\x27\x45\x23\xa1\x4f\x1f\x3d\x3a\x9d\xa2\x5b\x1b\xed\x10\xfe\xfb\xfb\xc5\xbf\x7f\x7d\x7d\xf3\xcf\xf3\xeb\x9f\x7f\x79\xf5\xeb\xd5\xf5\xd5\xf7\x91\xe2\x61\x12\xd9\x26\xdd\xef\x10\x16\xde\xaf\x1f\x98\x39\xc2\x40\x69\x1f\x10\xc1\x20\x10\x1e\xc2\xbc\xf4\x65\xd8\xee\xe4\x0c\x1a\xcc\xb1\x42\x0b\x54\x34\x0f\xdf\x24\x8a\x25\x67\x58\x94\x66\x83\x36\x2f\x8e\xa0\x27\x51\x16\x9f\x57\x22\x6a\x87\x10\x9f\x6a\x1e\x9e\xae\x38\xf3\xa5\x5c\xe6\x05\xf1\xcf\x92\x43\x58\x46\x6e\xd9\x91\x1b\x55\x05\x39\xc1\x29\x7a\x9a\x07\x6f\x90\xb7\x0f\x2e\x6d\x1b\xd7\xe3\x78\x0d\x43\x78\xe8\xea\x71\x3a\xdb\x49\xbc\x23\x93\x8a\xcf\xcf\x4e\x8e\xf3\x41\x47\x7b\x08\xa1\x1a\x75\xef\xc1\x94\x62\x8a\x17\x2d\x3a\x2f\x1e\x4f\xbf\x14\xa7\xa5\x5f\x7c\xc0\xa3\x11\xc6\x3c\x9d\xa1\x0b\x29\xd8\x1e\xf4\x2e\xc0\xbd\xf7\x99\x3b\x56\x4d\x3e\x20\xf1\xa2\x3e\xef\xb6\x0c\xb1\xfe\x7d\xdf\x49\x89\x2c\xb0\xeb\xb2\xc4\xde\x9c\x5a\x35\xc5\x9d\x9c\x3a\x8b\xdc\x57\x7a\x0e\x72\x81\x72\x09\x33\x5a\xb6\x56\x45\x5b\x50\xe0\xad\x29\x42\xa7\x2f\xa0\x77\xeb\x81\x77\xf2\x7f\x9a\xb3\x21\xbb\x54\xda\xe5\x52\x7c\xdc\x36\x0d\x29\x94\x17\x43\xe0\x63\xde\xb1\x3e\x5b\x98\xb6\x99\x53\x87\xc7\x56\xe5\x16\x2f\x0a\x18\x8d\x40\x9b\x46\x69\x3f\xa6\x9a\x01\x50\xa7\x1b\x19\x2f\x8c\x97\x00\x89\xb2\x8f\xde\x79\x6b\x71\x7e\xa0\xb5\x87\xe3\x44\xab\x26\xb1\x0c\xca\x27\x9a\x66\xf6\x04\x4c\xd5\x9b\x05\xca\x85\xdc\xc0\x3b\x5e\x12\x25\xf5\x88\x86\xf1\xb2\x65\xd9\x46\x45\x43\x87\x9f\x1b\xf1\x55\xd9\xb4\x78\x52\x51\x7a\xb4\xe4\x46\x89\x2f\x08\x65\x01\x93\x7d\xd6\xa9\xb7\x87\x28\x37\x4a\xc4\x26\xbd\xd9\x17\x67\xc4\x8e\xbd\x09\x00\x00\xff\xff\xf6\xe7\xbd\xe4\xf5\x06\x00\x00")

func pkgRestRestGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgRestRestGo,
		"pkg/rest/rest.go",
	)
}

func pkgRestRestGo() (*asset, error) {
	bytes, err := pkgRestRestGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/rest/rest.go", size: 1781, mode: os.FileMode(420), modTime: time.Unix(1653144221, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgRestServerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\x4f\x6f\xda\x4e\x10\x3d\x7b\x3f\xc5\xc8\x87\x9f\x6c\x44\x8c\x7e\x3d\x5a\xea\x01\x11\x5a\x2a\x25\x01\x05\xe7\x54\xf5\xe0\x2c\x63\xb3\x2a\xd9\xb5\x66\xc7\x40\x8a\xf8\xee\xd5\x78\x4d\xe2\x24\x0d\x07\x76\x3c\x33\xef\xbd\xdd\xf9\xd3\x94\xfa\x77\x59\x23\x10\x7a\x56\xca\x3c\x35\x8e\x18\x12\x15\xc5\xda\x59\xc6\x23\xc7\x2a\x8a\xd9\x3c\x61\xac\x54\x14\xd7\x86\xb7\xed\x63\xa6\xdd\xd3\xa4\x36\xf6\x4a\x52\xc8\x3c\x4e\xb4\x23\x1f\x7f\x1e\x6e\x1a\x72\xd5\x3f\xe2\xb5\xb3\x46\x8b\x15\xab\x54\x29\x7e\x6e\x50\x84\xb7\xcc\xcd\x3d\xd6\xc6\x33\xd2\xb7\xd6\x6a\xa8\x5a\xab\x93\x51\x6d\x6c\x36\xb7\xb5\xb1\x98\x4a\xf6\x64\x02\x8b\xa2\x58\xcd\x9c\xad\x40\x3b\x5b\x99\x1a\x36\x58\x19\x8b\x81\x67\xc1\xdc\x74\x31\xcf\xd4\x6a\x86\x93\x8a\x6e\x84\xd1\x82\x78\x8c\xad\x55\x34\xd5\x6c\xf6\xb8\x81\x47\xe7\x76\x2a\x5a\xad\xc8\x55\x00\xe1\xeb\xfc\xc2\xbf\x46\xda\x23\x41\x96\x65\x3d\xef\xab\xef\x95\x59\xf4\x61\x74\x91\xec\xd1\x77\x78\x18\x24\x5b\x3c\x0c\xb0\x4a\x9e\xf4\x36\x23\x79\x4b\x92\xc2\x68\x80\x3e\xa9\x88\x90\x5b\xb2\xf0\xdf\xab\xf7\xa4\xa2\x4e\x39\xef\xde\x3f\x56\xd1\xb9\x57\x5e\x73\x49\x0c\xbe\xfb\xaf\xa9\xd1\xe0\x91\xf6\x46\x63\x50\x4d\xfc\x90\x3a\x0d\xd9\x89\xe6\x23\x8c\xfa\x8e\x67\xb3\x70\x8e\x81\xfa\x36\xc0\xfb\x9e\xa4\x80\x44\x2e\x5c\xcc\xb5\x92\x91\x7f\x05\xe9\xd0\x35\x56\x65\xbb\xe3\x24\xbd\x04\xb2\x07\x8f\x49\x77\x53\xf2\xd9\x1d\x1e\x92\xce\x98\x75\x1d\x93\x17\x44\xd3\xdd\xce\x1d\x96\x64\x6a\x63\x7d\x0e\xf2\xfb\xf9\x2b\xb4\xe8\x14\x8b\x6c\x3e\x99\xc4\x63\xe8\x4c\x2f\xf6\x79\xfc\x82\xba\x45\xde\xba\xcd\x07\xd4\x72\x55\xfc\x58\xde\xad\x05\xf5\x7d\x5e\xc8\xb1\x5a\xae\xc3\xf9\x10\x8e\x69\x31\x5b\x88\x71\x3d\xbf\x99\x17\xf3\x21\xe5\x02\xcb\x0d\xd2\x07\xca\x69\xcb\x5b\x47\xe6\x4f\xc9\xc6\x59\x41\x76\x25\xb2\x7c\x75\x83\xb6\xe6\xed\xd0\x53\x3c\x37\xd8\x33\xce\x8f\x8d\xf3\x38\xa4\xfc\x8c\x71\x70\x83\x19\xe1\x06\x2d\x9b\x72\xe7\x73\x60\x6a\x71\xfc\xae\x4a\x52\xff\x3c\x2c\x85\xeb\x1c\xfd\x44\xa7\xdd\xec\x42\x57\xd4\xcb\xb8\x08\x5e\xbe\x03\xff\x6d\x79\x9c\xd6\x98\xc3\xff\x5f\x60\x04\xb2\xd3\xd9\xc2\xb5\x24\xa1\x73\x3a\x56\xd1\xdb\x9e\x49\x37\xef\x51\xbb\x3d\xd2\x73\x92\xa6\x2a\x32\x15\xf8\x4c\x46\x2d\x0b\xbb\x22\x42\xdd\x66\x67\x97\xc9\x48\x02\x3c\x95\x51\x8c\xe8\x83\xb3\xbf\x53\xaf\x71\xdf\xda\xa4\xe7\x0b\x9b\x99\xaa\xb3\xfa\x1b\x00\x00\xff\xff\xc9\x8f\xc8\x2e\x8d\x04\x00\x00")

func pkgRestServerGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgRestServerGo,
		"pkg/rest/server.go",
	)
}

func pkgRestServerGo() (*asset, error) {
	bytes, err := pkgRestServerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/rest/server.go", size: 1165, mode: os.FileMode(420), modTime: time.Unix(1653143038, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgStorageStorageGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x6f\x6b\xdb\x30\x10\xc6\x5f\x4b\x9f\x42\x33\xb4\xd8\x25\xd8\x1d\x94\x31\x0a\x61\x6b\xea\x6d\x2d\x24\xd9\x58\x3b\xfa\x5a\xb1\xce\xae\xa8\x2d\xb9\xa7\x73\x9b\x52\xf2\xdd\xc7\x29\xff\xda\xc0\x60\x7e\x11\x49\xcf\x9d\xa4\x9f\x9e\xbb\xf4\xba\x7a\xd0\x0d\xa8\x40\x1e\x75\x03\x52\xda\xae\xf7\x48\x2a\x95\x22\xa9\x3b\x4a\xa4\x48\xc8\x76\x90\x48\x29\x92\xc6\x63\x97\x5b\x5f\x18\xb4\x4f\x80\x45\xf7\x12\x1e\xdb\xe4\x8d\xce\xe3\xe1\xba\x68\x7d\xd3\x00\x26\x32\x93\xb2\x28\x54\xa9\x49\x4f\x74\x00\x95\xe7\xb9\xa4\x97\x1e\xf6\x4a\x20\x1c\x2a\x52\xaf\x52\x94\xf1\xfc\xb9\xee\xa2\x68\x5d\x23\xc5\x95\x0f\xa4\xd6\xdf\x56\xfa\xc5\x98\xeb\xcf\x3a\x92\xe2\x4f\x00\x3c\x4c\xd1\x21\xdc\x79\x34\x6f\xa4\x72\x12\xcf\x7d\x93\xb5\x8a\x60\xbf\xc1\xd8\x70\xe9\x5d\xbd\x27\xdb\x4b\x7b\xb4\x39\xd0\x9d\xc7\x87\x77\xb7\x30\xdb\x85\x31\xf8\x8f\x9b\x77\xda\x4c\x2f\xaf\x4d\x0b\x7b\xe2\x99\x5e\x5e\x54\x64\x9f\x60\x2b\x70\xf8\xd6\x76\xe0\x07\x62\xe1\xd3\xd9\x06\x6e\x0e\xcf\x3f\x3c\x76\xe5\x44\x39\x78\x56\xd1\xde\x72\x22\xeb\xc1\x55\xfb\x50\x5a\xa9\x93\xad\x97\x99\x3a\xd9\x24\x31\x73\x51\xa8\x00\x2d\x54\xa4\xe8\x1e\x94\xd1\xa4\x17\xec\x77\x7c\xa3\x0e\x2a\x96\x51\x0a\x13\x9c\x3a\x1f\xab\xba\xa3\xfc\xa6\x47\xeb\xa8\x4e\x93\xa3\x70\x7e\x14\xbe\x52\xd5\xa7\x3c\x33\x59\x71\x14\xbe\x54\xf7\x1a\x03\xd0\x78\xa0\xfa\x73\xb7\x38\x3b\xee\x79\xc9\xd0\xe3\x5b\x1c\xe0\xb8\xf5\xd5\x78\xea\x2b\xdd\x26\x23\x55\xe5\x5c\x13\x1e\xb7\x76\xf0\x9c\xed\x8a\x9a\xc7\x38\xae\x2b\x92\x49\x61\x16\x23\x05\x88\x4c\x11\xe9\x7f\xf6\xe0\xd2\x48\xb7\x9e\x9a\xe0\xb2\x91\x3a\x8e\x31\x2e\x8b\x6d\x5e\xa5\x10\xd3\xd8\x5f\xe7\x6a\xdd\x67\x79\x09\xb5\x1e\x5a\xca\xa7\xbe\x99\x79\x03\xe9\x46\xbe\x76\xb5\xcf\x46\x52\xac\x32\x29\x6c\x1d\xaf\xf9\x30\x56\xce\xb6\x6c\x90\xe8\xb5\xb3\x55\x0a\x88\x99\x14\x2b\x29\xc2\x63\x5b\x4e\x76\x2c\x66\x91\x97\x93\xf4\xbf\xf7\xe5\x37\x40\x97\xde\xb9\x99\x5e\x4e\x6d\x0d\xfc\xe7\x49\xf9\x27\xbf\xf2\x03\xe7\xed\x92\x66\x7a\xc9\xcf\xe2\xdc\x90\x7e\x3c\x3d\x3d\x88\x71\x2f\x6c\x63\x99\x14\x08\x34\xa0\x53\x66\xb1\xeb\xd7\xca\xa3\x99\x7b\xfa\xee\x07\x67\x62\xd3\xc6\x7e\x78\xaf\x33\x1b\x43\x7b\xcc\xd4\xc2\xfb\x48\xbd\x39\x89\x23\xe3\x8d\xd3\xdf\x10\xdf\xef\x93\x2b\xf9\x37\x00\x00\xff\xff\xcf\x95\x55\x82\x19\x04\x00\x00")

func pkgStorageStorageGoBytes() ([]byte, error) {
	return bindataRead(
		_pkgStorageStorageGo,
		"pkg/storage/storage.go",
	)
}

func pkgStorageStorageGo() (*asset, error) {
	bytes, err := pkgStorageStorageGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/storage/storage.go", size: 1049, mode: os.FileMode(420), modTime: time.Unix(1653144433, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"pkg/conf/conf.go":       pkgConfConfGo,
	"pkg/rest/jwt.go":        pkgRestJwtGo,
	"pkg/rest/page.go":       pkgRestPageGo,
	"pkg/rest/rest.go":       pkgRestRestGo,
	"pkg/rest/server.go":     pkgRestServerGo,
	"pkg/storage/storage.go": pkgStorageStorageGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"pkg": &bintree{nil, map[string]*bintree{
		"conf": &bintree{nil, map[string]*bintree{
			"conf.go": &bintree{pkgConfConfGo, map[string]*bintree{}},
		}},
		"rest": &bintree{nil, map[string]*bintree{
			"jwt.go":    &bintree{pkgRestJwtGo, map[string]*bintree{}},
			"page.go":   &bintree{pkgRestPageGo, map[string]*bintree{}},
			"rest.go":   &bintree{pkgRestRestGo, map[string]*bintree{}},
			"server.go": &bintree{pkgRestServerGo, map[string]*bintree{}},
		}},
		"storage": &bintree{nil, map[string]*bintree{
			"storage.go": &bintree{pkgStorageStorageGo, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
