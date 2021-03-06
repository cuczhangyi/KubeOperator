// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xcb\x6e\xeb\x36\x10\xdd\xf3\x2b\x26\xf2\xb6\xe8\x07\x78\xc7\x2b\x31\x89\x7a\x65\x49\xd0\xe3\xb6\xe9\x46\xa0\xa9\x71\xcc\x5e\x99\x14\x48\x0a\x69\xba\xeb\x7f\xf5\x9f\xfa\x0b\x05\x65\xda\x56\x6e\x1c\x24\x41\x57\x86\x01\xcd\xcc\x39\x73\xce\x1c\xae\x84\x3e\x1c\xb4\x22\x39\xdd\xb0\x8e\xfd\x96\xd6\x4d\xbd\x86\x28\xe7\x07\x04\x3e\x18\xe4\xfd\x33\xe0\x9f\xd2\x3a\x1b\x11\xb2\x12\xc3\x64\x1d\x1a\x12\x67\x6d\xdd\xb0\xaa\x4b\x58\xc6\x1a\xd6\xdd\xd2\x34\x63\xc9\x1a\x22\xc1\x15\x28\xed\xa0\xc7\x01\x1d\x42\xf8\x1c\xa4\x02\x31\x19\x83\xca\x81\x75\xdc\x61\x44\x56\xc2\x60\x8f\xca\x49\x3e\x90\x17\x4d\xba\x8a\xd5\x45\x5b\xc5\x6c\x0d\xd1\x2d\x97\x03\xf6\xe0\x74\xe8\x77\x03\xcd\x1e\x0d\x82\xb4\xc0\x15\x70\x6b\xb5\x90\xdc\x61\x0f\x7b\x6d\x1d\x4c\xaa\x47\x03\x6e\x2f\x2d\x7c\xc7\xe7\xe8\x8d\xb6\xdd\xef\x45\xfe\xa9\xde\x7f\x69\x85\x57\x7a\xdf\xd2\x36\x6b\xba\xb8\x62\x09\xcb\x9b\x94\x66\x5d\x4c\xf3\x2e\x2f\x9a\xb0\x92\x35\x44\x09\xee\xf8\x34\x38\xb8\x30\x05\xc1\x95\xdf\xce\x16\xc3\xd0\xde\xef\xd4\x83\x27\xf7\x45\xdd\x74\x34\xab\x18\x4d\x1e\x2e\x2a\xdc\x7b\x5e\x3f\xaa\x10\x78\xcd\x15\xe7\xc5\x5f\xa5\x73\xdc\x8b\x67\x14\x5a\x2c\x68\x3d\x49\xb7\x07\xb7\x3f\x6b\x74\xad\x6f\xf7\xe5\xa1\x2b\xab\xe2\x17\x16\x37\xff\x6b\xc4\x68\xf4\x1f\x28\x5c\x44\xea\x87\xba\x61\x9b\x2e\x2d\xe7\x4d\xdd\x16\x6d\xee\xb1\x97\x03\x72\x8b\xb0\x93\xc3\xe0\x9d\xe2\x2b\x06\x2d\xf8\x00\x69\x09\x3b\x69\xac\xfb\xf7\x9f\xbf\xfd\xa2\x26\x8b\x86\x94\xb4\xae\x7f\x2d\xaa\x64\xee\xb0\xa1\x4d\x7c\xbf\x86\xc8\x03\x19\xb9\xb5\x4f\xda\xf4\x1e\x8c\x54\x42\x1b\x33\x8f\x2c\xaa\xf4\x2e\xcd\x69\xf6\xea\x7b\x6d\xe4\xa3\x54\x7c\x78\xab\xb0\xad\x59\xd5\xa5\xf5\x5c\x47\xe3\x26\xfd\xc6\x42\xa1\x87\xe1\xbf\xf5\x4a\xa2\xe2\xdb\x01\xfb\x1b\x08\x1c\x84\x56\x8e\x0b\x37\x73\xe0\xfd\x41\x2a\x69\x9d\xe1\x4e\x9b\x9b\xd0\x70\xc9\x3b\xd7\x60\x27\xb1\x9f\x1b\xce\x14\x69\xb2\x49\xf3\xd7\x46\xf2\x43\xfb\x60\xa6\xb9\xe9\x11\xc2\x2b\x33\xdd\xbc\x04\x5d\xb1\x8c\x36\x2c\x59\x28\xd8\xfa\xb2\x3d\xf7\xd0\x97\x3a\x05\x79\xae\x40\x68\xcb\x84\x7e\x0c\x02\xf6\xf2\x88\x80\x90\x95\xc1\x47\xa9\xd5\xc9\x4f\x15\xbb\x4b\x8b\xfc\xa3\xd7\x0d\xc7\xe2\xf7\x1c\xe5\x8f\xd2\x5b\xc2\xff\x9e\x06\xf9\xc3\xfe\xf0\x98\xf9\xaa\xdf\xb3\xed\xc0\x95\x1f\xa2\x47\x54\xd6\x71\xf1\x9d\xdc\xb1\xe6\xc4\x87\x55\x55\x51\x1d\x45\x0c\x90\x77\x7a\x52\xf3\x41\x87\x7d\x6e\xf0\xb0\x45\x73\x96\x84\x26\xc9\x52\x82\x2d\xa2\x02\xde\xf7\xb8\x2c\x39\xe7\x4a\xd0\xec\xed\x50\x09\x05\xd7\x12\xe5\x54\x7b\x4f\xeb\x2e\xe4\xb4\x3f\xb2\x50\xb0\x20\x1a\x2e\xff\x06\xe2\x2b\x4e\x22\x64\xa5\x74\x8f\x24\x2f\x12\x76\x0e\xa6\xaa\xcd\xf3\x34\xbf\xeb\x1a\x5a\x7f\x5d\x43\x44\xfb\x1e\xfc\x47\xa0\xcd\x29\xf1\xe7\xbf\xa7\xa5\x9a\x49\x29\xa9\x1e\x7f\x1a\x8f\xc7\xf1\xc4\xa5\x03\xe9\xa0\xd7\x0a\x7f\xf6\xac\xb7\x5c\x7c\x9f\x46\x2a\x84\x9e\x94\x23\x25\xad\xe8\xa6\x63\x9b\xb2\x79\x58\x43\x94\x2a\x3b\xed\x76\x52\x48\xff\x68\x8c\xdc\xf0\x03\x3a\x34\x36\x22\x7e\x1d\x75\x5b\x96\x45\x35\x5b\x5a\xd9\x69\x1c\xb5\xf1\x7c\xdc\xf3\x88\x11\x89\xef\x59\xfc\xf5\x12\x8c\xdf\xd0\xc8\x9d\x14\xdc\xcd\x12\xcd\x5e\x38\x47\xdd\x17\x1a\x7f\x6d\xcb\x8e\xc6\x71\xd1\xe6\x9f\x09\xbd\x17\xc0\x3f\x9c\x7e\x64\xe5\x1d\xf5\xc3\xbb\xf4\x81\x69\xbe\xea\x13\x43\x82\xac\x5f\x66\x8c\x24\x70\xbc\x4d\x33\x76\x8c\x85\xbc\xcd\xb2\x4b\xe8\x06\xd9\xdc\x99\x94\x8f\x61\x84\x2d\xee\xb4\x41\xb0\x4f\xd2\x89\xbd\x54\x8f\xcb\x0f\xf8\x91\xf6\xc2\xb6\x15\x5a\x3d\x19\x81\xaf\xa9\x2d\xa6\xbf\x3b\xf3\x12\xf5\xd7\x5f\xa2\x8b\x97\xcf\xef\xce\x69\x23\x5b\x1c\xb4\x7a\xb4\x7e\x71\x8b\x27\xcd\x6f\x4f\x5a\xd0\x23\x9a\xa3\xfc\x97\x63\x19\xd1\xec\xb4\x39\x04\xa7\x13\xf2\x5f\x00\x00\x00\xff\xff\x64\x78\x74\xc6\x05\x09\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 2309, mode: os.FileMode(420), modTime: time.Unix(1597299768, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x53\xdb\x56\x14\xdd\xeb\x57\x78\xec\x6d\xa7\x3f\x20\xbb\x87\xf4\x00\x15\x59\xd2\x3c\x49\x6d\xe9\x46\xe3\x18\x4d\x4b\x03\x16\xe3\x8f\x4d\x57\x35\x14\x7f\x10\x3b\xa6\x89\x49\xa1\x78\x82\x9d\x9a\xe2\x49\x62\x9b\x8f\x0c\x06\x84\xc9\x9f\xd1\x7b\x92\x57\xfc\x85\xce\x93\x6c\x23\xe3\xd2\x7a\xab\xb9\xf7\xdc\xf3\xce\xb9\xf7\x28\x12\x37\xd7\xd7\xcd\x04\x23\x82\x28\xd4\xe1\xf7\xbc\xa2\x2a\xcf\x42\x61\xbc\x5b\x76\x4e\x4e\x71\xef\x1c\xb7\xf7\x71\xad\x15\x66\x98\x48\x7c\x2d\x93\x4a\x1b\x49\x86\x15\x34\x45\x85\x48\xe7\xa0\x00\x55\xa8\xcf\x03\x5e\x80\xdc\xb3\x50\x98\xfc\x51\x27\x17\x7b\xb8\x50\x1f\x1c\x34\x71\xff\x0d\x2e\x96\x9d\x9d\x4b\xf2\x6b\xd6\xf9\xf3\xb7\xc1\x61\xce\xb9\x6b\x7a\x20\x49\x63\xc5\x48\xa4\x57\x63\x6b\xcc\x44\xbf\x8e\xa0\x22\x69\x88\x85\x74\xb6\x0f\xd1\x3c\x73\x3f\x1f\xdf\xdf\x66\xdd\xee\xb1\x73\xb2\x3f\x78\x7d\x6c\x5f\xbd\x24\xb5\x22\xde\xbe\x70\xb3\x55\xfb\xca\x22\xb5\x9b\xf0\x13\x20\xfa\x0f\x92\x38\x2b\x12\xae\x74\x9d\x6a\x0b\x97\x3c\xb0\x79\xa0\x09\xaa\xce\x22\xc8\x41\x51\xe5\x81\xa0\xb3\x40\xd4\x45\x49\x1d\x3e\xf6\x59\x28\x3c\xb0\xf6\xdd\x4e\x13\xe7\xdb\xa4\xdc\xb1\xaf\xca\xee\x56\xdf\x1f\x72\x7f\x9b\xa5\xef\xfb\xc9\x4c\xa5\x99\x45\x49\x51\x75\x20\x20\x08\xb8\xe5\x07\x49\x7d\xca\x01\x49\x87\xdc\xbd\xea\xb1\x8a\xd3\x94\xc7\x7d\x8e\x55\xf1\x29\x8f\xe4\x9c\x06\xd0\xe7\x96\x75\x19\x49\xdf\x40\x56\x9d\x15\xab\x71\xed\x1c\x76\x3c\xf6\xca\xb2\xa2\xc2\xa8\xce\xcb\xde\x8b\xe7\x25\x4d\xa4\x84\xdc\x6e\x0f\x6f\x17\x70\xe3\x23\xce\x1d\x90\xda\x27\x52\xbb\xe1\xe5\xe1\x63\x33\x29\x23\xc9\xc8\x40\x51\xbe\x93\x10\xe7\x35\x45\x81\xca\x2e\xd2\xc9\xdd\x9c\x53\xcf\x0e\xaa\x07\x6e\xb7\x1b\x66\x24\xc4\x2f\xf0\x22\x10\x26\x4b\x5e\x1d\x4d\x56\x69\x0a\x44\x3a\xaf\x78\x45\x80\x55\xf9\x6f\xa9\xdc\x4e\xb5\x45\x0a\x3d\x52\xfb\x80\x77\xa9\x4d\xde\x2b\x7a\x6e\xb6\xea\x5c\x58\x4e\xa7\xe1\xec\xe6\xf0\xef\xfb\x1e\x1b\xaf\x3b\xc8\x9b\xae\x63\xbb\xe9\xf7\x7b\x15\x80\x8b\xf2\xe2\x53\x86\x86\x62\x2b\xeb\xab\x89\x90\x5f\xee\xfb\xea\xbe\xff\x18\xb0\x36\xc8\x0e\x41\x01\xa8\x90\x0b\x28\x3d\xa4\x79\xde\x18\xaf\x95\xaf\xeb\xe3\xa9\x9a\xcc\x01\x6f\x2a\x98\x1a\x67\x7f\xe9\x90\xea\xb5\xaf\x2c\x13\x49\x1a\x3f\xae\x9a\x89\x91\xc3\x08\x2e\xf0\x92\x38\xd3\xa9\xe0\xd2\x0d\x3e\x3a\x0a\x3a\x1c\x58\x70\x26\xf2\x8b\x99\x30\x46\xa8\xf4\x48\x66\xc3\x1c\x21\x4c\x2c\xce\x56\xcb\xe9\x9f\xbb\x9d\x06\x2e\xbc\xa6\xc8\xe6\x86\x91\x48\xa5\x63\xf1\x17\xcc\x02\x54\x47\x8c\x21\x42\x12\xa2\x66\x14\xef\xec\xab\x32\x2e\x9c\xfa\xf4\x68\xfd\x46\xd2\xfc\xd9\x88\xa7\xa3\xc6\xfa\x73\x23\x39\x96\x17\x70\xdc\x58\x4e\x7f\x1a\xe9\x59\x78\xa7\x1e\xe8\x18\x9f\xe9\x50\xfe\xa7\x2c\xf5\x1d\x98\xba\xd1\x51\xd7\x22\x50\xf4\x61\x94\xd1\x16\xaf\x38\x78\x5e\xf7\xb7\xd9\x7f\xb9\xef\x84\xb9\x62\x30\xa2\xc4\xc1\xf1\x7d\x23\x4d\x14\x79\x71\x41\x57\x81\xb2\x44\xd5\xdb\xbe\xb4\xad\xb7\xee\xce\xa6\xb3\x79\x4d\xf6\x4e\x07\xf9\x0a\x79\x53\xb6\xfb\x35\xd2\xfe\x0b\xd7\x5a\xa4\x78\xe2\x36\x4a\x5f\x85\xdc\x6e\xcf\x69\x17\xf1\xdd\x36\xee\x6c\xd9\xd6\x27\xff\x33\xee\x94\x48\x77\xef\x6b\x3a\xe6\x79\x2c\xfe\x22\xb3\x01\xe2\x71\x33\x93\x48\x33\x32\x40\x20\xaa\xc3\xa8\xac\x2e\xd3\x09\x95\x4d\xb2\x77\x4a\xb9\x5d\x5e\x84\x19\xfa\x70\x45\x93\x65\x09\xa9\x5e\xc6\x94\x49\xb5\x4b\x4a\x34\x74\x9d\x33\x0b\xbf\x7b\x19\x66\xd8\x45\xc8\x2e\x05\x62\xba\xde\x18\x7c\x28\x8d\xed\x1d\xa7\xc8\x1c\x60\x97\x34\x59\x07\x2c\x2b\x69\xe2\xac\x79\x82\x9b\x79\xdb\xea\xbb\x9f\xff\xc6\x95\xde\x13\xa9\xc2\x44\x36\xd6\x62\x89\x47\x41\xfd\x3f\xb0\xc1\xdd\x9a\x86\x0d\xfc\x8c\xe6\x3c\x9d\x98\x21\xf7\x79\x5e\x80\xfe\x85\x8a\x9a\x20\x3c\x64\xd7\x10\x9f\x52\x25\x6f\xf3\xb6\x75\x89\x73\x65\x5c\xc8\x93\xf2\xfb\x20\xff\xc0\x8a\x21\x23\x65\x66\x92\x71\x63\x9a\x75\x60\xd2\x7f\xe0\x07\x75\x7d\x94\xce\x0f\x2b\x37\x91\xc5\x67\xef\xec\x9b\x57\x8f\x16\xcf\xfd\x72\xe8\x36\x4a\xa4\xdd\xf4\x37\x68\x18\x0d\xcc\x3f\x01\x00\x00\xff\xff\xa1\xbc\xaa\xba\xb1\x07\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 1969, mode: os.FileMode(420), modTime: time.Unix(1597299745, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
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
