// Code generated by go-bindata.
// sources:
// query.sql
// DO NOT EDIT!

package ods

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

var _querySql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x41\x6f\xa3\x3a\x10\xbe\xf3\x2b\x7c\x6b\x2a\xf5\x21\xbd\x77\xed\x8b\xf4\x28\x21\xaf\xe9\xa6\x50\x41\xda\x6a\x4f\x91\x43\x06\xea\xac\x6b\x90\x71\x22\xe5\xdf\xaf\x6c\x0c\xb5\x21\xd0\xe6\xb0\xab\x56\x5a\x6e\x9e\xf9\x66\xe6\xf3\x78\xfc\xe1\x59\xe0\x2f\xbd\x38\x40\xff\xa5\x98\x73\x02\x7c\x9d\x16\x5b\x40\x4f\x5e\xec\xdf\x7a\xf1\xe4\x9f\x4b\x34\x45\x17\x4f\x37\x17\xd7\x8e\x93\x04\xcb\xc0\x5f\x39\x08\x11\xd7\xaf\xb1\xbe\x84\xaa\xcf\x4b\x90\x8e\xbf\x52\x80\x19\x94\x98\x8b\x3d\x87\x44\x60\x41\x0a\x26\x01\x05\x27\x39\x61\xb5\xdf\xe3\x9c\x1c\x30\x6d\xbc\x2a\xc1\x16\x2a\x41\x98\x32\xd4\xa0\x39\x25\xf9\x8b\x08\xf7\xaf\x1b\xe0\x4d\x95\x4c\xd9\xd6\x4c\x19\x6b\x58\xb2\x9a\x21\xeb\x93\xb9\xb0\x80\xf5\xb6\x21\xa1\x71\x80\x45\x95\x14\x74\x6b\xe0\x2a\x69\x5b\x57\x05\xdd\x1a\x18\xef\x80\x09\xc5\x1b\x0a\x26\x06\x37\x46\x09\xe4\x6e\x0c\x07\x60\x7b\xb0\x8a\xf2\xda\x26\x01\xb8\x8f\xf0\x12\x84\x59\x4a\x28\xc5\xfc\xb8\xd6\x50\x67\x1e\x47\xf7\x68\xe2\x68\x48\xdb\x61\xfd\x11\x6a\x76\xfa\xca\x72\x74\x3b\x6c\x7b\xed\xfe\xda\x3e\xb3\xad\xb6\x27\x59\xcd\x0c\x43\xf2\x78\x3f\x21\xa9\xeb\x53\x5c\xa9\xb6\x5d\xa2\x5e\x9b\xdb\x96\x5e\x75\x48\x97\x38\x25\xe2\x88\xfe\x3a\x91\xa4\x09\x6b\xbb\xdc\x84\xaa\x5e\x2c\xd8\x01\x98\x28\xf8\x71\x09\xb9\x84\x12\xfa\x96\x79\x11\x86\x41\x8c\xee\xa2\x45\x68\xc1\x54\x6e\x85\x4d\xd1\xf3\x62\x75\x8b\x26\x28\x8c\x96\x91\xff\x0d\x5d\x3a\x06\xdd\x28\x94\xd4\xcc\xc0\xc5\x0c\x4d\x11\x49\x3b\xb6\x26\xe6\xf9\x36\x88\x83\xa1\xb3\x40\x53\xfb\xc2\xbc\xe1\xbc\x70\xa6\x5a\x29\xb0\xd8\x57\xe8\xdf\x29\xfa\xbb\xf1\xfd\x1f\x47\x8f\x0f\xe8\xe6\xfb\xaf\x3e\x5f\xfb\x14\x47\x0f\xbc\x39\xa8\xda\xa8\xce\x86\x38\x66\xa3\x1d\xa4\xa7\xd3\x1a\xcd\xd3\xbc\x47\x48\x0f\x32\x1e\x60\x67\xef\x42\xce\x50\xb9\x4b\x5d\xff\x05\xf3\x1c\xbc\xd7\x62\xcf\x84\xe2\xaa\xaf\x98\xd3\x4e\xcf\x4d\x51\xfc\x20\x4c\x0d\xce\xc6\xe9\x0d\x8d\xf6\x3e\xe0\xaa\x02\x96\x03\x57\xb0\x72\x70\x66\xa2\x10\x6d\x5c\x1d\xa3\x46\x65\x53\xbe\x2d\xfb\xd9\xdb\xb4\x77\xc5\x9e\x33\x38\x26\x90\xbf\x02\x13\xb2\x48\xb9\xab\x46\xab\x94\x6e\x1b\xac\x0a\x95\xbb\xca\xb4\xbc\x5f\xab\xee\x4c\x5d\x6a\xf8\x12\x44\x61\x37\xb3\xaa\x95\x9e\xa8\xd5\x8e\xb2\x0c\xd0\x3b\x69\xe1\xed\xfa\x7d\x62\xfa\x16\x97\x3b\x7a\x2e\x2b\x7a\x1e\x2b\x3a\xc6\xaa\x2f\x2a\xe3\x6c\x4e\xc9\x44\xd7\xd6\xa7\x44\xdd\x25\xe4\xfa\x6f\xa5\x02\xda\xa5\xc2\x9a\x7a\xf2\x21\x31\x19\x50\x92\x8e\x8c\xfc\xa6\xbb\xe8\x68\x79\xa8\xf7\x22\xb5\xb4\xb3\x03\x6e\xae\x1d\xa3\x31\x27\x1e\x03\x12\xdd\x35\xda\x21\x9d\xf7\x81\x0c\xb0\x4d\x36\xdc\x7a\x29\x48\xb0\x69\xb0\xa1\xf2\xb5\x20\x11\xf5\x9e\x3e\xa3\xd2\x65\x5f\x46\xe9\xe6\x50\x6b\x4e\x76\x9e\xba\x65\x1f\x14\xb7\x39\x80\x21\x6c\xd9\xb8\xb0\x65\xbd\x2a\xe3\xba\x96\xb9\x73\x80\x76\x62\x24\xba\x5d\x7f\x2e\x6d\xff\x92\x12\x9a\xfd\x51\xd0\x9e\x82\x36\x12\x8a\x07\x35\x14\x9f\x27\xa2\xf8\x6c\x15\xc5\x67\xc9\x28\x7e\x5f\x47\xb1\x12\xd2\x6b\xe7\x67\x00\x00\x00\xff\xff\x6b\x74\x9c\xcb\x46\x0e\x00\x00")

func querySqlBytes() ([]byte, error) {
	return bindataRead(
		_querySql,
		"query.sql",
	)
}

func querySql() (*asset, error) {
	bytes, err := querySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.sql", size: 3654, mode: os.FileMode(420), modTime: time.Unix(1521024067, 0)}
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
	"query.sql": querySql,
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
	"query.sql": &bintree{querySql, map[string]*bintree{}},
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

