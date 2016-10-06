package idl

import (
	"fmt"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _grpcbasic0_swagger_json = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "User Service",
    "description": "User Service API provides create, read, and read (many) access to a list of \nusers.",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/user": {
      "post": {
        "operationId": "RecordUser",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idlUser"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/idlUserRecordReq"
            }
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    },
    "/v1/user/{id}": {
      "get": {
        "operationId": "GetUser",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idlUser"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    },
    "/v1/users": {
      "get": {
        "operationId": "GetUsers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/idlUsers"
            }
          }
        },
        "parameters": [
          {
            "name": "start",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "count",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "desc",
            "in": "query",
            "required": false,
            "type": "boolean"
          }
        ],
        "tags": [
          "UserService"
        ]
      }
    }
  },
  "definitions": {
    "idlUser": {
      "type": "object",
      "properties": {
        "age": {
          "type": "string",
          "format": "int64"
        },
        "fortune": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "idlUserGetReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "idlUserRecordReq": {
      "type": "object",
      "properties": {
        "age": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "idlUsers": {
      "type": "object",
      "properties": {
        "users": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/idlUser"
          }
        }
      }
    },
    "idlUsersGetReq": {
      "type": "object",
      "properties": {
        "count": {
          "type": "string",
          "format": "int64"
        },
        "desc": {
          "type": "boolean",
          "format": "boolean"
        },
        "start": {
          "type": "string",
          "format": "int64"
        }
      }
    }
  }
}
`)

func grpcbasic0_swagger_json_bytes() ([]byte, error) {
	return _grpcbasic0_swagger_json, nil
}

func grpcbasic0_swagger_json() (*asset, error) {
	bytes, err := grpcbasic0_swagger_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "grpcbasic0.swagger.json", size: 3626, mode: os.FileMode(436), modTime: time.Unix(1475773316, 0)}
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
	"grpcbasic0.swagger.json": grpcbasic0_swagger_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"grpcbasic0.swagger.json": &_bintree_t{grpcbasic0_swagger_json, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

