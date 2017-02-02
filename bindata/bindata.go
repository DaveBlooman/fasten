// Code generated by go-bindata.
// sources:
// libraries/amzlinux/nodejs/install.yaml
// libraries/amzlinux/nodejs/nodejs.sh
// libraries/amzlinux/ruby/install.yaml
// libraries/amzlinux/ruby/ruby.sh
// libraries/definition.yaml
// libraries/ubuntu1604/nodejs.sh
// libraries/ubuntu1604/ruby.sh
// DO NOT EDIT!

package files

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _librariesAmzlinuxNodejsInstallYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x28\x4a\xd5\x4d\xce\xcf\xcd\x4d\xcc\x4b\xb1\x52\xc8\x29\xe6\x82\xb3\xf3\x0a\x72\x15\x32\xf3\x8a\x4b\x12\x73\x72\xb8\x00\x01\x00\x00\xff\xff\x6c\x32\x79\xd0\x25\x00\x00\x00")

func librariesAmzlinuxNodejsInstallYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxNodejsInstallYaml,
		"libraries/amzlinux/nodejs/install.yaml",
	)
}

func librariesAmzlinuxNodejsInstallYaml() (*asset, error) {
	bytes, err := librariesAmzlinuxNodejsInstallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/nodejs/install.yaml", size: 37, mode: os.FileMode(420), modTime: time.Unix(1483959172, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesAmzlinuxNodejsNodejsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x4b\xaa\xc3\x30\x0c\x85\xe1\xb9\x57\xa1\xcb\x1d\x06\xc5\xc3\x42\x37\x53\x1c\x45\x34\x6e\xfd\x22\x92\x20\x81\x2e\xbe\x76\x86\x1d\xff\x87\xef\xfc\xff\xf9\x25\x16\xbf\x04\xd9\x9c\x13\x56\x40\xb6\xc3\xb9\xd3\x32\x58\x5b\x83\x32\xe0\xe9\xc8\xf6\x04\x88\x12\x13\x97\xbe\xc0\x54\x29\x68\xac\x05\x36\xd5\x26\x77\xef\xf7\x96\xe7\x52\x57\x96\x6a\x3b\xf1\x4c\x35\xfb\x6e\x59\x7b\xdc\xe6\x03\x3e\x30\x74\xc0\x0b\x8d\x45\x34\xa4\xae\x9d\xf0\x24\x42\x9a\x26\xc8\xe1\x7d\xbd\xfc\xe4\xe1\xbd\x64\x84\x6f\x00\x00\x00\xff\xff\x53\x9c\xdb\x1e\xa4\x00\x00\x00")

func librariesAmzlinuxNodejsNodejsShBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxNodejsNodejsSh,
		"libraries/amzlinux/nodejs/nodejs.sh",
	)
}

func librariesAmzlinuxNodejsNodejsSh() (*asset, error) {
	bytes, err := librariesAmzlinuxNodejsNodejsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/nodejs/nodejs.sh", size: 164, mode: os.FileMode(420), modTime: time.Unix(1484753114, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesAmzlinuxRubyInstallYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x28\x4a\x8d\x4f\xce\xcf\xcd\x4d\xcc\x4b\xb1\x52\x48\x4f\xcd\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\x48\x2a\xcd\x4b\xc9\x49\x2d\xe2\x82\x4b\x42\x04\x60\xf2\x5c\x80\x00\x00\x00\xff\xff\x98\xd3\x76\xaf\x39\x00\x00\x00")

func librariesAmzlinuxRubyInstallYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxRubyInstallYaml,
		"libraries/amzlinux/ruby/install.yaml",
	)
}

func librariesAmzlinuxRubyInstallYaml() (*asset, error) {
	bytes, err := librariesAmzlinuxRubyInstallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/ruby/install.yaml", size: 57, mode: os.FileMode(420), modTime: time.Unix(1483442956, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesAmzlinuxRubyRubySh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\xcf\xca\xc2\x30\x10\xc4\xef\xfb\x14\xfb\xf5\xf3\x1a\x23\xf5\x26\x58\xf0\xe6\xd1\x83\x2f\x90\xb4\x83\x2d\xf4\x1f\xd9\x44\x1a\x9f\x5e\x13\x2d\x82\x97\x59\x66\x86\x1d\x7e\xff\x7f\xda\x76\xa3\xb6\x46\x5a\x22\x81\x67\x85\xb0\x10\xc5\x30\x70\x98\x1b\xe3\xc1\x2a\x66\x07\x67\x04\xec\x82\x8d\xe5\x2e\x9f\xb5\xe8\x46\xf1\xa6\xef\xdf\xd5\x5e\x35\xb8\x63\x35\xfc\xe8\x3b\x9b\xe5\x13\xbf\x5e\x6e\xf8\x2e\x2b\x89\xe2\x31\x10\xea\x76\xe2\x02\xcb\x3c\x39\xcf\x97\xd3\xf5\x7c\xdc\x24\x3d\xe8\x76\x1a\xa0\x51\x97\x2a\x08\x5c\x02\x2d\xb8\xaa\xf8\x27\xde\x26\x78\x57\xd3\x33\x00\x00\xff\xff\x02\x2a\x0e\x8b\xcb\x00\x00\x00")

func librariesAmzlinuxRubyRubyShBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxRubyRubySh,
		"libraries/amzlinux/ruby/ruby.sh",
	)
}

func librariesAmzlinuxRubyRubySh() (*asset, error) {
	bytes, err := librariesAmzlinuxRubyRubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/ruby/ruby.sh", size: 203, mode: os.FileMode(420), modTime: time.Unix(1484753114, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesDefinitionYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xcc\xad\xca\xc9\xcc\x2b\xad\xb0\xe2\x52\x50\xc8\xc8\xcf\x4d\xb5\x52\xd0\x07\x51\xfa\xa9\xc9\x46\xba\xa5\xc5\xa9\x45\x5c\xa5\x49\xa5\x79\x25\xa5\x86\x66\x06\x26\xe8\x4a\x20\x32\x5c\x80\x00\x00\x00\xff\xff\x24\x55\x8b\xd6\x42\x00\x00\x00")

func librariesDefinitionYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesDefinitionYaml,
		"libraries/definition.yaml",
	)
}

func librariesDefinitionYaml() (*asset, error) {
	bytes, err := librariesDefinitionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/definition.yaml", size: 66, mode: os.FileMode(420), modTime: time.Unix(1486055873, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesUbuntu1604NodejsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func librariesUbuntu1604NodejsShBytes() ([]byte, error) {
	return bindataRead(
		_librariesUbuntu1604NodejsSh,
		"libraries/ubuntu1604/nodejs.sh",
	)
}

func librariesUbuntu1604NodejsSh() (*asset, error) {
	bytes, err := librariesUbuntu1604NodejsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/ubuntu1604/nodejs.sh", size: 0, mode: os.FileMode(420), modTime: time.Unix(1483958389, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _librariesUbuntu1604RubySh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\xe2\x2a\x2e\x4d\xc9\x57\xa8\x2c\xcd\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\x28\x2a\x4d\xaa\xd4\x4d\x49\x2d\x4b\x85\x30\x15\xd2\xf3\x15\x2a\x92\x93\x74\x4b\x4b\x32\x73\x14\x74\x2b\xb9\x00\x01\x00\x00\xff\xff\xa6\x58\x98\xfc\x3d\x00\x00\x00")

func librariesUbuntu1604RubyShBytes() ([]byte, error) {
	return bindataRead(
		_librariesUbuntu1604RubySh,
		"libraries/ubuntu1604/ruby.sh",
	)
}

func librariesUbuntu1604RubySh() (*asset, error) {
	bytes, err := librariesUbuntu1604RubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/ubuntu1604/ruby.sh", size: 61, mode: os.FileMode(420), modTime: time.Unix(1483958389, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"libraries/amzlinux/nodejs/install.yaml": librariesAmzlinuxNodejsInstallYaml,
	"libraries/amzlinux/nodejs/nodejs.sh": librariesAmzlinuxNodejsNodejsSh,
	"libraries/amzlinux/ruby/install.yaml": librariesAmzlinuxRubyInstallYaml,
	"libraries/amzlinux/ruby/ruby.sh": librariesAmzlinuxRubyRubySh,
	"libraries/definition.yaml": librariesDefinitionYaml,
	"libraries/ubuntu1604/nodejs.sh": librariesUbuntu1604NodejsSh,
	"libraries/ubuntu1604/ruby.sh": librariesUbuntu1604RubySh,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"libraries": &bintree{nil, map[string]*bintree{
		"amzlinux": &bintree{nil, map[string]*bintree{
			"nodejs": &bintree{nil, map[string]*bintree{
				"install.yaml": &bintree{librariesAmzlinuxNodejsInstallYaml, map[string]*bintree{
				}},
				"nodejs.sh": &bintree{librariesAmzlinuxNodejsNodejsSh, map[string]*bintree{
				}},
			}},
			"ruby": &bintree{nil, map[string]*bintree{
				"install.yaml": &bintree{librariesAmzlinuxRubyInstallYaml, map[string]*bintree{
				}},
				"ruby.sh": &bintree{librariesAmzlinuxRubyRubySh, map[string]*bintree{
				}},
			}},
		}},
		"definition.yaml": &bintree{librariesDefinitionYaml, map[string]*bintree{
		}},
		"ubuntu1604": &bintree{nil, map[string]*bintree{
			"nodejs.sh": &bintree{librariesUbuntu1604NodejsSh, map[string]*bintree{
			}},
			"ruby.sh": &bintree{librariesUbuntu1604RubySh, map[string]*bintree{
			}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

