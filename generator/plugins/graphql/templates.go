// Code generated by go-bindata.
// sources:
// templates/schemas_body.gohtml
// templates/schemas_head.gohtml
// templates/types_body.gohtml
// templates/types_head.gohtml
// DO NOT EDIT!

package graphql

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesSchemas_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcb\x6e\xdb\x3a\x10\x5d\xcb\x5f\x31\x30\x8c\x0b\xfb\x42\xd1\x07\x18\xc8\xca\x68\xd3\x2e\x9a\x47\x93\x5d\x51\x14\x8c\x3c\x56\xd4\xd2\xa4\x4b\x51\x2e\x82\x01\xff\xbd\x20\x87\xa2\x28\x3b\x31\xaa\x15\x35\xcf\x73\xce\x70\x68\x5f\x0f\x08\x44\x8b\xea\xb1\x7e\xc1\xbd\xb8\x15\x7b\x74\x8e\xcf\x1b\xd9\xa2\xb2\x1d\x74\xd6\xf4\xb5\x05\x9a\x15\x44\x60\x84\x6a\x10\x16\x1d\x9a\x63\x5b\x23\xac\xaf\x61\x51\x3d\xf2\x4f\x07\x57\xce\xcd\x8a\x82\x68\x70\x57\x5c\x8e\x0b\x01\x51\xa3\x9f\x7c\xbb\xe4\x66\xc7\x4d\xb0\xfa\x54\x22\x40\xb5\x0d\x65\xdc\x6c\xd7\xab\x1a\x6e\xd0\xbe\x0d\x6e\x59\xcb\xee\x32\xee\x12\xda\x17\xf8\x9f\xa8\x55\x16\x4d\x8d\x07\xab\x4d\x77\xff\xab\x71\xae\xfa\x3c\x5a\x3e\x09\xb5\x95\x68\x88\xae\xa0\xdd\xc1\xa2\x7a\x32\xa2\x46\xf3\x41\x89\x67\x89\x01\x48\x09\xd6\x00\x91\x0d\x76\x4e\xe7\x18\x9f\x12\xc1\xae\x60\x49\xd4\xfc\x96\xec\x66\x10\x25\xa0\x31\xda\xac\xfe\x5d\xb6\xa3\xf0\x8d\x4e\xa4\xfb\xd8\xa2\xdc\x76\x70\x0d\x44\xd1\xb1\xd1\x8a\x07\xa2\xcd\x58\x6f\xe1\x9c\x17\xa4\x7a\x47\x7a\x2f\xc5\x65\x8a\x39\x9b\x88\xe5\x47\xe8\xfa\x26\x9e\xc9\xa4\x32\x7a\xfa\xf9\x27\xd6\x96\xd9\xdd\x85\xf3\x09\x39\x0e\x88\xb5\x42\xfd\x24\xdb\x2d\xfe\xe1\x94\x5c\x4b\xb6\x6c\xb4\xda\xb5\x0d\xcd\x8a\xa2\xf0\x99\x6b\x98\x9f\x96\x9a\x97\xde\x49\xe4\x19\xaa\x04\xa4\x7a\xe8\xb5\xc5\xed\x46\xef\xf7\xfe\xfe\xcd\xe7\x11\x4c\x51\x44\xd3\x3a\xc3\x34\x89\x75\x6e\x28\x98\x58\x16\x05\x73\x5f\xe7\xa0\xd9\x14\x90\xc5\xee\x12\x55\x6a\x1f\x87\x37\x34\xcd\x84\xda\xc9\x6d\x50\xe9\x9d\xc0\x58\xcb\x87\x0d\xb7\x24\x38\xe1\xc2\xe7\x35\xf1\xf1\x51\x90\x40\x2d\xcb\x9f\xcc\xef\xdb\x49\xf0\xf7\x72\xec\x8b\xb2\xc3\x0c\x49\x71\x56\xf7\xbf\x53\xfe\x34\x84\x66\xd3\xc9\x32\xca\xe4\xf6\x6b\x9e\x70\xdd\xe5\xf3\x1b\x63\xbe\x62\xa7\xe5\x11\xd7\xe0\xd7\x7f\x79\xc8\xc5\x8e\xae\x7b\x61\xc4\xbe\x5b\xc1\x32\x2c\xf6\x4e\xd4\x48\x2e\xdf\xb6\xe1\x33\x68\x7b\xa3\xe2\xeb\x45\xce\x07\xa9\x56\x26\xff\xd8\xd3\xe5\xec\xc7\x71\x9f\xfd\x9f\x89\x33\x57\x9a\xf5\xbc\xa4\x4a\xd4\x24\x85\x0e\xcd\x06\x31\xc6\x77\xc3\x9a\x56\x35\xd1\x1d\x31\x4d\x11\x04\xa3\xdf\xcf\xdc\x1c\x69\x4e\x37\x29\xbe\x91\x67\xaf\xd2\xb8\x49\x0f\x3d\x9a\xd7\x30\x8c\x2a\x1c\x79\x1c\x3c\x89\x78\xfb\xaa\x2f\xbd\x15\xb6\xd5\x8a\x7d\x03\x8a\xc1\xca\xc9\xd3\x98\x94\x9f\xe0\xb9\xd5\xcc\xfd\x0d\x00\x00\xff\xff\x87\xe6\x38\xcf\x64\x06\x00\x00")

func templatesSchemas_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_bodyGohtml,
		"templates/schemas_body.gohtml",
	)
}

func templatesSchemas_bodyGohtml() (*asset, error) {
	bytes, err := templatesSchemas_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_body.gohtml", size: 1636, mode: os.FileMode(420), modTime: time.Unix(1532351561, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemas_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xc1\x4a\xc4\x30\x14\x85\xe1\xb5\x79\x8a\xc3\xd0\x85\x2e\x26\x01\x97\x82\x0b\x61\x06\xe9\xc6\xba\x88\x0f\x70\xdb\x5e\xd3\xd0\x36\xad\x49\x44\xe4\x72\xdf\x5d\xd0\x82\xbb\x9f\xef\x1c\xe7\xe0\xa7\x58\xf0\x1e\x17\xc6\x17\x15\x04\x4e\x9c\xa9\xf2\x88\xfe\x1b\x21\xd6\xe9\xb3\xb7\xc3\xb6\xba\xeb\xb3\x3f\xbf\xcd\x99\x62\x62\x17\xb6\xfb\xf0\xb1\x58\x5c\x3a\xbc\x74\x1e\xd7\x4b\xeb\xd1\x7a\xb3\xd3\x30\x53\x60\x88\x34\xf6\x68\x55\x63\xe2\xba\x6f\xb9\xe2\xd6\x88\x64\x4a\x81\xd1\x1c\xf2\xf0\x88\xc6\xfe\x75\xc1\x59\xd5\xdc\x88\x1c\x9b\x7d\x5a\x22\x15\x55\x9c\xfe\xe9\x95\xea\xa4\x7a\x32\x22\x9c\xc6\xdf\xff\x9d\xf9\x09\x00\x00\xff\xff\xc2\xe3\x00\xe2\xbf\x00\x00\x00")

func templatesSchemas_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemas_headGohtml,
		"templates/schemas_head.gohtml",
	)
}

func templatesSchemas_headGohtml() (*asset, error) {
	bytes, err := templatesSchemas_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schemas_head.gohtml", size: 191, mode: os.FileMode(420), modTime: time.Unix(1532510936, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_bodyGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x73\xe3\xb6\x11\x7f\xa6\xfe\x8a\x0d\x9b\xdc\x90\x19\x46\xee\xb3\x32\x7a\xf0\xb8\xe7\xeb\x4d\x7a\x1f\x3d\x7b\x92\x07\xc7\xe3\xc0\x14\x24\xa1\xa2\x40\x1a\xa4\x64\x7b\x38\xfc\xdf\x3b\xf8\x22\x01\x12\xa4\x28\x47\x6e\xaf\x9d\xe3\xc3\xc5\x02\x17\xd8\xc5\x62\xf7\x87\xfd\x60\xce\xce\xe0\x2d\xdd\x6d\xf3\x49\x59\x32\x44\x57\x18\xbe\xc7\x74\xb7\x85\xd9\x1c\xa6\x97\x24\xc1\x53\xf1\x12\x7e\xaa\xaa\x89\xb7\x47\x0c\xca\x52\xbc\x9f\xfe\x8a\x18\x41\xf7\x09\xfe\x88\xb6\xb8\xaa\x60\x0e\x65\xb9\x7a\x48\x3e\x6f\x56\x55\x35\xfd\x88\x1f\xf9\xac\xc0\x18\xe2\xbf\x2f\x52\xba\x24\xab\x72\xe2\x79\x7c\xd2\x0c\xd4\xe3\xeb\x25\xdf\x31\x94\xad\xff\xf9\x0f\xb9\xa2\x1f\x4d\x3c\xaf\x2c\x81\x2c\xa5\x40\xd3\x8b\x74\xbb\xc5\xb4\x90\x92\x78\xde\xdf\x70\x1e\x33\x92\x15\x24\xa5\xb3\x5a\x28\x45\x53\x55\x6a\x32\xa6\x0b\x45\xff\x2b\x4a\x76\x38\x9f\x41\x4b\x24\x31\x2c\xe5\xfa\x80\xb2\x72\x02\xc6\x53\xeb\x63\xcf\x89\xb8\x42\xf4\xce\xf9\x52\x5a\x0e\x8f\x8b\x2f\x28\xa6\x4a\xf0\x19\xbc\xe9\xe7\x52\x8a\x39\x52\x1c\x21\xb7\x9c\x2a\x7e\x4b\xb1\x3d\xbd\x6d\xaa\x59\xd7\x5b\xff\xc3\xf7\xff\xa8\xf9\x76\x55\x60\x11\x9b\x8b\x35\x6a\xf0\x3c\x39\x5c\x96\xcd\x18\x1f\xa9\xc2\x49\x33\x34\x99\x9c\x9d\xc1\x7b\x9a\xed\x0a\x48\xef\xff\x85\xe3\x62\x52\x96\xa0\x74\x21\x07\x1a\xeb\x10\x64\x9f\xc4\xa0\x6d\x24\x92\xf0\x90\x99\x18\xd3\x4d\x6b\x31\x86\xb5\xda\xf4\xa9\x48\xd3\xf1\x1b\x0e\x6d\xab\xd1\x84\x97\x04\x27\x0b\xfb\xc4\x3b\xcb\x0a\x1a\x7e\xf0\x95\x9c\x57\x85\x13\x6f\xb9\xa3\x31\x10\x4a\x8a\x20\x2c\x85\x15\xa9\x9d\x2f\x39\xad\xb0\x02\xc5\x58\x32\xd0\x8a\xed\xdb\xb2\x22\x0b\xc2\x1b\x2e\xb3\x58\x44\x1b\xca\x2d\xcc\x2d\x53\x31\xc4\x13\x93\xca\xcf\x8c\xec\x51\x81\x9b\x2d\x5b\xd3\x23\xb8\x7e\xce\x84\x11\xc5\x28\x49\x94\x80\x53\x3e\x06\xdf\x57\x42\x28\xe3\x94\xab\x89\x61\x07\xed\xf3\xcd\x81\xe1\x3c\x4d\xf6\x98\xe5\xc6\x51\xeb\x31\xe7\x61\x7f\xd1\x13\xc4\x82\x42\x67\x65\x59\x4f\x99\x5e\xee\x68\xcc\xed\x52\x8a\x1a\x28\x3f\x9e\x5e\x33\x14\x63\xf6\x96\x72\xf5\x2c\xa0\xaa\xa0\xe0\xb6\x52\x88\x51\xa9\x03\x49\x11\x81\x10\xbd\xaa\xe2\xe2\x89\xef\xaf\x78\x92\x6f\x2f\x52\x5a\xe0\xa7\x22\x02\x02\x84\x16\x98\x2d\x51\x8c\xcb\x2a\x84\xe0\x8e\x1f\x73\x2a\xf7\x5e\x0b\xf1\x69\x57\x64\xbb\xe2\x9d\x18\xae\xaa\x08\x18\x66\x0c\x30\x63\x29\xe3\x27\xeb\x96\x49\x9e\x66\x9e\x21\xca\xb7\x5d\xb0\xe9\x05\xc3\xa8\xc0\x17\x6b\x92\x2c\xae\x32\x44\x2f\x59\xba\x55\x52\x04\x71\xf1\x14\x89\x63\xe9\xd9\xb6\x1f\x4e\x3c\x6f\x81\x97\x98\x01\x5f\x70\x7a\x49\x28\xc9\xd7\x41\x33\xca\xd5\x26\xad\xcc\x23\x4b\xc8\xb8\x78\xb3\x39\x30\x1c\xa7\x7b\xcc\x82\xf0\x67\x39\xf4\xdd\x1c\x28\x49\x40\x22\x87\x58\xe8\x0a\x17\xd7\x68\x15\xf8\x62\x2f\x7e\x04\x7e\xc1\x76\xd8\x0f\xed\xf1\xbb\x2d\xce\x73\xb4\xc2\x7e\x24\x96\x69\xbf\xcd\x0b\x14\x6f\xfc\x08\xf2\x82\x11\xba\x0a\xca\x72\x81\xef\x77\x2b\xa9\xe6\x2b\xfe\x2e\x08\x43\x2e\xa9\x57\x29\xf1\xd8\x69\x64\xe1\xcb\x4c\xdf\xf2\xc1\xa0\x59\xbf\xe2\x4a\x31\xac\x95\x2c\x81\xc0\xbc\xe1\xc5\x70\xb1\x63\x94\xff\x8c\xf8\x3f\xdc\x9c\x3d\xc4\x56\x39\x57\x17\x99\x06\x5b\x94\xdd\xc8\x7d\xdc\x9a\x66\x31\xf1\xee\x60\x0e\x9c\x4e\xa2\x12\xc3\xf9\x2e\x29\x60\x0e\x14\x3f\x06\xda\x5e\x2e\x53\xf6\x11\x3f\xf6\x5a\x8d\x90\x0b\x3a\x10\xd0\x9c\xb8\x09\x02\xca\xa4\xa4\x23\xbe\xcf\xb9\xad\x9c\x73\x29\x15\x44\x90\xa5\x10\xc6\x00\x02\x85\x5d\xc2\xaf\xc4\x4a\x35\x32\x18\x8a\xb6\x01\x5c\xf2\x20\xf9\x39\x63\xe8\x59\xf3\x92\xd2\xd6\x18\x4f\x84\xf1\x8e\xe5\x35\x0d\x6e\x5a\x6a\xf3\xb8\xc6\xb9\xae\xa6\xf5\x6c\xa9\x15\x63\x1e\xcc\x61\x8b\x36\x38\x68\xfc\xce\x14\x85\xbb\x5b\x82\x69\x40\xa8\x3c\x64\x6f\x99\x32\x20\x11\xec\x51\x22\x4c\x5c\xe8\x93\x50\x65\x48\xb6\xe2\x34\xb4\xfc\x46\x8a\xb5\x30\x14\xa8\x6f\xbc\x7d\x04\xca\x49\x6c\xd0\x13\xd7\xa7\x9e\x07\xfe\x1e\x25\x3e\xc7\x40\x35\x8b\x2c\xa1\x63\xbb\x9e\x6d\x54\x65\x29\xec\x34\x97\xf6\xff\x1b\x43\x59\x80\x19\x8b\xc0\x5f\x22\xc2\x61\xa1\x48\x35\x44\x02\x31\x80\x13\x04\x7b\x3f\x54\x4b\x6a\x86\x07\x95\x77\x43\x38\xf6\xef\x9b\x0b\x3a\xc9\xb1\x71\xb1\x8f\x9c\x7f\x84\x0e\xda\x41\xc0\xc4\xc1\xf7\xc0\x21\xd4\xe2\x8d\x3b\x84\x20\x4e\x69\x8c\x0a\xf0\x85\x19\xfe\xee\xfb\x30\x64\x87\xe0\xff\xee\xdf\xfa\x61\x23\xb0\xfb\xcc\x4e\x7e\x64\x8a\xdb\x18\x6b\xdf\x4f\x9c\x87\x35\x66\xea\xeb\xe9\xa9\x83\x0d\xcd\xaf\x03\x60\x54\xb5\x90\xc5\xfe\x5b\x47\x7b\x14\xa7\x4b\x1b\xf0\x3e\x51\xfc\x69\xd9\x42\x3d\x45\x4d\xe8\x02\x3f\x45\x56\xa4\xc4\xe7\x77\x02\x25\x7e\xb6\x0f\x8a\x1c\xfe\x6a\xa0\xe3\x21\xac\xba\x8b\x20\xdd\x1c\x03\x6d\x3f\x73\xfa\x37\x6f\x0e\x2f\xdc\x18\x1a\xb4\x9e\x11\x5e\xd1\x9e\xc2\x9f\xe3\x9c\x64\xf8\xc8\xef\xd4\x81\xbb\xf8\xb4\xfd\xc4\x45\xc3\x9f\x17\x7a\x4e\xca\x4f\xdb\xf2\x9c\x83\xba\xf4\x43\xa7\x10\x5d\xf1\x4d\x57\x72\xaa\xf0\x75\xb5\xd7\x84\x1b\x83\x47\x7e\x9e\xe7\x64\x45\x09\x5d\x71\x3d\x65\xb8\xff\xc4\x1b\x20\x90\x56\x7f\x18\x08\x3a\x4b\xfb\x7b\xbf\x47\xd4\x61\x4d\x8d\x61\xbd\x77\xae\xda\x60\x47\x55\x96\x9a\x87\x64\xf6\xcd\x1f\xbf\xf9\xa3\xad\xc2\x6f\xfe\x38\x52\x53\xa7\xf2\x47\x9d\xae\xab\x4c\x5e\x5e\xcd\xf2\x8f\x89\x8e\x85\x24\x2f\x99\x16\x59\x49\xfe\xd9\x19\x48\xbe\x3a\xc9\x77\x56\x71\xbe\x97\x99\xbd\xa4\x7c\x79\x1d\xa7\x5b\xc2\xb1\xab\x37\xba\xe4\x37\x50\xb7\xf1\x3c\x47\xc5\x46\x0e\x95\xaa\x44\xd5\xad\xcd\x74\x33\xb3\xe3\x8a\x33\xe7\x8b\x85\xa0\x94\x72\x06\x8e\x1a\xcb\x9b\xb6\x34\x32\x16\xed\xab\xc9\x88\x97\x43\x75\x19\x49\xa1\x7c\x66\x26\x4b\x00\x99\xb9\x65\xf5\xea\x33\x62\x68\x9b\x87\x10\x18\xa9\x59\xa4\xaa\x17\x06\xae\x78\xe2\x9f\xfc\x91\x14\xf1\x1a\x72\x16\x73\x1d\x64\xd3\xab\x74\xc7\x62\x3c\x0d\x8a\xe7\x0c\x87\x3a\x78\x8e\x51\x8e\xe1\xc7\x26\x5f\xd3\xe7\xa0\x12\xb6\x99\xce\x8b\xc8\x52\x2c\x34\x6f\xa5\x4b\xdd\x24\xdc\x4e\x7a\x3c\x91\x8d\xff\x98\xb3\x58\x0f\x88\x48\x4f\xeb\x07\xe3\xc5\x05\xca\x0b\x23\xcb\xa9\x57\xac\x95\xc8\x09\xae\x53\x51\x2f\xea\xa0\x0c\xf8\xb9\x70\xc6\xd0\x62\xde\x5c\x5a\xdd\x45\x7b\x97\x68\xad\x60\x65\x46\x4a\x4d\x87\xb5\x74\xea\xcd\xb1\xf8\x04\xdb\x53\x8b\xb4\xd7\x30\xf7\x57\xa7\x2d\x7d\xf7\xcf\x47\xfc\x18\xf8\xb9\x30\x20\x48\x97\xb0\xa3\x1b\x9a\x3e\x52\xe0\xa6\xa4\xb2\x26\x69\xc3\xa2\x50\xd0\xce\x38\xfa\xdc\xf1\x03\xca\xbe\x79\xe4\x7f\xce\x23\x5d\x97\x92\x7e\x3a\x6e\xaa\x4a\x63\x70\x73\xeb\xae\xa3\x69\xba\x65\xca\x60\x83\x9f\x45\x11\x47\xf6\x43\xe4\x69\x0f\x78\x9a\x25\x6b\x0e\x73\xe0\xf7\x2a\x5d\x04\x0c\xe7\x11\xb8\xb9\x35\x33\x3c\x7f\x83\x9f\xfd\x19\x80\xe0\x6a\x0c\x0b\xfe\xfe\x4c\xca\xd1\xbc\xa8\xc2\x0e\x24\x35\x77\xa4\xa9\xa0\x91\x1e\xfe\x3a\x7a\x51\x2e\xfa\x55\x6a\xe6\xf5\xb0\xa1\x1d\x99\x7c\x40\x59\x6e\x95\x65\xdc\xd1\x89\x0c\x4e\x3e\xa0\xec\xff\xa3\xcd\x74\xbd\xde\xd1\x4d\x20\xab\xfe\xa3\x26\xb4\x12\x88\x1a\xfb\x47\x74\xb4\x26\x86\x9d\x0c\x35\x99\xd4\xc6\x6d\x80\x54\x7b\xfe\x05\x3f\x4b\xca\xba\xa9\x24\xf7\xed\x69\x81\x6a\x83\x7b\x29\x07\xe1\x15\x4e\x1e\xfa\x69\xc0\xac\x0a\xeb\x26\xdd\x21\x63\x1a\xd9\xcf\xb2\x0d\xcb\x6e\x69\x79\x5f\x41\x4f\x8b\xe3\x80\xa3\xaf\xf5\x0b\x7e\xd6\x60\x75\xeb\x78\x2d\x74\xda\xd7\xf5\x1a\x6c\x7b\x9d\xb2\xef\xd5\xd3\xf8\xea\x76\xbe\x46\xb7\xbe\xfe\x0b\xbd\x2f\x85\x88\xce\xee\xd7\x89\xda\x5f\x92\x85\x68\x80\x59\xa1\x5a\xbb\x05\xe6\xb8\xec\x05\x19\x55\x3d\xb0\x4e\x0f\x47\x75\xbb\x66\xaa\x4f\x73\x02\x5b\xe2\xab\xaa\x3e\x0e\x71\x36\x72\xee\x60\x0e\x64\x22\xae\x4e\xf1\x96\x53\x0d\x74\xe7\x3c\x6f\x13\xc9\xaa\xc2\x1e\x25\x37\x02\xab\x6e\x23\xf9\xb7\xc4\x95\x5b\xb1\x68\x04\x7c\xdd\x4d\xb4\xaf\xab\xc7\x96\xf8\xfd\x9d\x8a\xcd\xa6\x53\xf4\x71\x4d\x04\x7f\xd3\x74\x4a\x9c\x7d\x87\xe1\x62\xcd\xb2\xb7\x5a\xf3\x97\x1f\x16\xdc\x87\x01\x27\x58\x7c\xbf\xc1\x77\x18\x01\x09\xcd\x1e\x8c\x15\xe0\x6f\x36\xc7\x08\x6b\x45\xf6\x2d\xc5\x58\x95\x99\xae\x6a\xf6\xdd\x7a\x98\x7b\x6a\x5d\x03\x79\x75\xe5\xc8\x23\x1f\x54\xcf\x7e\x7f\x9c\xc0\x96\x82\xa4\x3f\xdc\x6c\x36\xb7\xf3\xfd\x5e\x79\x8f\xa3\x86\xd2\xfd\x52\x42\x5c\x2f\xe9\xe1\x52\x4a\x7d\xa7\x7c\x05\xc5\x14\x9d\x8b\xf4\x97\x54\x9a\x0f\x7f\xba\x55\x95\x91\xa9\x99\x34\x67\x77\x3a\xa6\x24\x14\x24\x93\x4e\x02\xd6\x17\x5f\x08\xd2\x3f\x9d\x89\x09\x70\x6e\x27\x5e\x03\x28\xe4\xce\xaf\x5c\xc9\x55\x35\x31\xde\xe4\x2c\xae\x41\x4b\x11\x88\x0d\xa8\xe0\x77\xa4\x12\xb5\xd9\x0f\xaa\x51\x11\xf5\x2a\xd2\x11\x46\xfd\xef\xa9\x52\x61\x7e\x57\x99\x95\xf9\x6d\xda\xd9\x19\x5c\x61\xb6\x27\x31\x36\xbd\x30\x97\x43\x8d\x1b\x6a\x1a\x23\x96\x7b\x87\x8b\xb2\xd4\x94\xaa\x3c\xa0\xc8\x3e\xe0\x62\x9d\x2e\xf2\x20\x36\xd2\x42\x4d\x78\x81\x92\xe4\xbd\xde\x26\x0f\xa8\xc8\x9a\xe7\xec\x62\xe7\x31\xce\x8a\x1a\xee\xde\x37\x23\x7f\x47\x74\x91\x60\x06\xbd\x31\x62\xd4\x13\x24\xea\x18\x31\xec\x7a\x2c\x18\xd1\x9b\x96\x4d\xc9\xdd\x40\x5c\x3b\x3f\x68\x9c\xdd\xe8\xd3\x6e\xc5\x2c\x51\x9e\xe9\x59\x48\x7e\xd5\x28\xe9\xdc\x9f\x35\x1a\x85\x17\x13\x8e\xec\x29\x3a\x03\xb5\x6d\x56\xd1\x28\xc8\x92\x60\x69\xd7\x62\xea\x7d\x2a\xd2\x73\xb6\xda\xf1\x3b\x22\x37\x2b\x61\xe7\x6c\xe5\x00\x36\xe9\x56\x7a\x02\x17\xb0\x5d\xff\x30\xeb\x54\x88\xad\x84\x1a\x86\xf8\x48\x5d\x20\xb6\x72\x2b\x42\xcf\x71\x66\x39\x7c\x96\x91\xd7\xe8\x05\x3b\x35\x47\xcf\xdc\xb8\xfd\xea\x08\x17\xbe\x03\xcb\x89\x8d\xc0\x7f\xa0\x11\xc5\x1f\x9e\x91\x08\x0f\x57\x11\xfe\x20\x31\x0f\xc6\xe2\xe2\xc9\xa8\x9f\xf6\xe7\x13\x75\x9c\x7c\x38\xa9\xc8\x9a\x44\xc8\xef\xb8\xe9\xb4\x6d\x5a\xa0\xaf\xfb\xfa\xdb\x9b\xde\x84\xc3\x78\x67\x66\x1d\x5e\x7f\x60\x7f\xca\xe8\xde\xae\xba\xa8\x38\x5f\x29\xae\x65\x03\x3c\xe2\x5f\x77\x8b\x7c\xb6\x2b\x7c\xc1\x0f\x3b\x9c\xd7\xb9\x6a\x6f\x0b\xca\xb2\x78\x86\x1f\x3a\xe1\x5e\xcf\x7a\x7e\xc6\x2d\x5a\x56\xef\xca\xf2\x27\xf1\x1d\x72\x5a\xf4\x91\xab\xc0\x92\x4b\xa1\x2e\x8d\x80\x92\x24\xd4\x4d\xaa\x43\x82\xc1\x11\x0d\xd3\xd6\x86\x9a\x0b\x05\x33\x36\x6a\xd2\x38\x71\x5a\xe5\x75\xb5\xed\x8b\x84\x60\x5a\x48\x90\xe4\xf7\x01\x57\x54\xec\x83\xcf\xf0\x83\xf9\xa1\x99\xeb\x5b\x2e\x01\xcb\x0f\x5c\xf3\xae\xaf\x0e\x6d\xb5\xea\x34\x6b\xec\xfe\xff\x8c\x94\x2d\xfc\xa9\x8d\x50\x61\xc1\x1b\xd7\x0d\xa7\x3c\xb4\xb1\x4d\x75\x7d\xce\x1c\x2e\xeb\x37\x70\x27\x25\x1a\xba\x20\x3c\x4f\x82\xd8\x0c\xb2\xc8\x51\xaa\x6c\xac\x97\xac\x35\xea\x71\x2b\x0d\x88\x28\x45\x08\xb7\x16\x82\x3b\xef\xe5\x1a\x57\x28\x7e\x2a\xc0\x45\x61\x2c\xf9\x9e\xee\xd3\x0d\x66\x21\x04\x2a\x7f\x6e\xc7\x45\x76\x6c\x74\x32\x07\x75\x9d\xe7\xab\xfa\xe6\x01\x8b\x95\xfe\x35\xd6\x62\xad\xf6\x83\xcb\xbc\x42\x03\xe4\x9c\x80\xfb\xc2\xef\x21\x10\xd7\x47\x6f\x75\x9b\xac\x45\xf4\xa6\xcc\x44\xd8\xd1\x58\x5b\xe1\x3e\x6b\x1d\x7d\xaf\xf1\x70\x0e\xd2\xc2\x5f\x60\x3b\x4c\x7f\x41\xc3\xf0\xc3\xd4\xf8\x96\xd6\x8d\x0b\x5e\xa3\xc4\xef\xd2\x8d\x79\x5d\x0d\x97\xec\xcb\x72\xb9\x2d\x54\x79\x2b\x63\x84\x16\xcb\xc0\xff\x62\x68\x10\x8c\x5d\x29\x04\xcc\xe1\x1e\x2d\xb8\x54\x9c\xbf\x28\xf0\x07\x3f\x5c\x87\x53\xb8\x5a\xa7\xbb\x64\x01\xf7\x22\xc8\x19\x92\x56\xdc\x86\x0f\xc6\x2d\x68\x9a\x9c\xe1\x2f\x9f\xd1\x73\x92\xa2\x85\xb8\x34\x2f\xd6\x38\xde\x8c\xf4\x19\x09\xd6\x79\xdf\xbd\xd6\x8b\x85\xee\x2f\x3a\x5c\xcf\x4b\x6e\x26\x68\x7d\xda\x73\xba\xcb\x49\x7e\x1f\x69\xed\xd1\xa5\x3a\x9f\xe1\xdc\x17\xed\xa5\xb1\xe2\xda\xf0\x65\xae\x79\x1e\xc7\x38\xcf\x07\x3e\xaa\xea\x7b\xb8\x7f\x59\x2b\x41\xe7\x7c\x9c\x7c\x94\xf0\xc7\x88\x3e\xf8\x1d\xce\x31\xf2\x31\x9c\x1f\xc5\xd7\x06\x38\xd3\xbc\xad\xde\xda\x18\x03\x30\xaa\x5c\x63\xb8\x1f\x79\xf9\xb7\xaf\xfe\x3e\x68\xd6\x29\x88\xfa\x6f\xeb\x6b\xa4\x49\xbb\x18\xd7\x58\xf9\xc0\xff\x68\xf4\xef\x00\x00\x00\xff\xff\xef\x2e\xf1\x55\x5c\x38\x00\x00")

func templatesTypes_bodyGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_bodyGohtml,
		"templates/types_body.gohtml",
	)
}

func templatesTypes_bodyGohtml() (*asset, error) {
	bytes, err := templatesTypes_bodyGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_body.gohtml", size: 14428, mode: os.FileMode(420), modTime: time.Unix(1532342222, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTypes_headGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xc1\x4a\xc4\x30\x14\x85\xe1\xb5\x79\x8a\xc3\xd0\x85\x2e\x26\x01\x97\x82\x0b\x61\x06\xe9\xc6\xba\x88\x0f\x70\xdb\x5e\xd3\xd0\x36\xad\x49\x44\xe4\x72\xdf\x5d\xd0\x82\xbb\x9f\xef\x1c\xe7\xe0\xa7\x58\xf0\x1e\x17\xc6\x17\x15\x04\x4e\x9c\xa9\xf2\x88\xfe\x1b\x21\xd6\xe9\xb3\xb7\xc3\xb6\xba\xeb\xb3\x3f\xbf\xcd\x99\x62\x62\x17\xb6\xfb\xf0\xb1\x58\x5c\x3a\xbc\x74\x1e\xd7\x4b\xeb\xd1\x7a\xb3\xd3\x30\x53\x60\x88\x34\xf6\x68\x55\x63\xe2\xba\x6f\xb9\xe2\xd6\x88\x64\x4a\x81\xd1\x1c\xf2\xf0\x88\xc6\xfe\x75\xc1\x59\xd5\xdc\x88\x1c\x9b\x7d\x5a\x22\x15\x55\x9c\xfe\xe9\x95\xea\xa4\x7a\x32\x22\x9c\xc6\xdf\xff\x9d\xf9\x09\x00\x00\xff\xff\xc2\xe3\x00\xe2\xbf\x00\x00\x00")

func templatesTypes_headGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTypes_headGohtml,
		"templates/types_head.gohtml",
	)
}

func templatesTypes_headGohtml() (*asset, error) {
	bytes, err := templatesTypes_headGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/types_head.gohtml", size: 191, mode: os.FileMode(420), modTime: time.Unix(1532510936, 0)}
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
	"templates/schemas_body.gohtml": templatesSchemas_bodyGohtml,
	"templates/schemas_head.gohtml": templatesSchemas_headGohtml,
	"templates/types_body.gohtml": templatesTypes_bodyGohtml,
	"templates/types_head.gohtml": templatesTypes_headGohtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"schemas_body.gohtml": &bintree{templatesSchemas_bodyGohtml, map[string]*bintree{}},
		"schemas_head.gohtml": &bintree{templatesSchemas_headGohtml, map[string]*bintree{}},
		"types_body.gohtml": &bintree{templatesTypes_bodyGohtml, map[string]*bintree{}},
		"types_head.gohtml": &bintree{templatesTypes_headGohtml, map[string]*bintree{}},
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

