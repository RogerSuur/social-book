// Code generated by go-bindata.
// sources:
// api/pkg/db/migrations/sqlite/000001_users_followers.down.sql
// api/pkg/db/migrations/sqlite/000001_users_followers.up.sql
// api/pkg/db/migrations/sqlite/000002_posts_comments.down.sql
// api/pkg/db/migrations/sqlite/000002_posts_comments.up.sql
// api/pkg/db/migrations/sqlite/000003_groups_events.down.sql
// api/pkg/db/migrations/sqlite/000003_groups_events.up.sql
// api/pkg/db/migrations/sqlite/000004_messages.down.sql
// api/pkg/db/migrations/sqlite/000004_messages.up.sql
// api/pkg/db/migrations/sqlite/000005_notifications.down.sql
// api/pkg/db/migrations/sqlite/000005_notifications.up.sql
// DO NOT EDIT!

package database

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

var __000001_users_followersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xcb\xcf\xc9\xc9\x2f\x4f\x2d\x2a\xb6\xe6\xc2\xae\xa0\xb4\x18\x2c\xc9\x05\x08\x00\x00\xff\xff\xd8\x66\x13\x19\x3f\x00\x00\x00")

func _000001_users_followersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_users_followersDownSql,
		"000001_users_followers.down.sql",
	)
}

func _000001_users_followersDownSql() (*asset, error) {
	bytes, err := _000001_users_followersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_users_followers.down.sql", size: 63, mode: os.FileMode(420), modTime: time.Unix(1680720125, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_users_followersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\xcb\x6e\xc2\x30\x10\x3c\xdb\x5f\xb1\x47\x90\xf2\x07\x9c\x02\x6c\x90\xd5\x90\xb4\xc6\x91\xe0\x14\x39\x89\x01\xab\x26\x8e\x62\xa7\xa8\x7f\x5f\x85\x42\x1b\xc4\xeb\xe8\xd9\x99\xf1\xec\xec\x8c\x63\x28\x10\x44\x38\x8d\x11\x58\x04\x49\x2a\x00\xd7\x6c\x25\x56\xd0\x39\xd5\x3a\x18\x51\xa2\x2b\x60\x89\xc0\x05\x72\x78\xe7\x6c\x19\xf2\x0d\xbc\xe1\x26\xa0\x64\x6b\xdb\x5a\x1e\x14\x08\x5c\x8b\x93\x32\xc9\xe2\x38\xa0\xc4\x75\xf7\x71\x75\x90\xda\x5c\xa3\x90\x25\xec\x23\xc3\x80\x92\x46\x3a\x77\xb4\x6d\x75\xa3\x2a\x74\xeb\xf7\x95\xfc\x86\x79\x28\x50\xb0\x65\x4f\xae\x75\xf9\xf9\xf7\x45\x40\x89\x2c\x6c\xe7\x2f\x0f\x7d\x90\x3b\x95\x37\xd2\xef\x2f\x48\xd9\x2a\xe9\x55\x95\x4b\x3f\x34\xd1\x2e\x6f\xba\xc2\xe8\x12\xa6\x69\x1a\xff\x47\x9a\x63\x14\x66\xb1\x80\xad\x34\x4e\xd1\xf1\x84\xd2\x17\x2d\xe5\x4e\x39\xa7\x6d\xed\x9e\x95\x75\x22\x0e\xa6\xc3\xbe\x7e\xe5\xc3\xcd\x8d\x09\x28\x21\x51\xca\x91\x2d\x92\xde\x01\x46\x67\x83\x31\x25\x84\x70\x8c\x90\x63\x32\xc3\xf3\x99\x46\x3d\x7e\x8a\xfa\x2c\xeb\xd6\x1a\x63\x8f\x3d\xfd\xe9\x51\x7b\x92\xae\x77\x0f\xc2\x5e\x4c\x1e\x8c\x65\x59\xaa\xc6\xab\x0a\x0a\x6b\xcd\xf5\xc0\xeb\x2f\x75\x03\x5f\xef\x38\xfc\xfc\xe1\xa2\xe4\xbe\xea\x45\x3b\xe3\x09\xfd\x09\x00\x00\xff\xff\xda\x9c\xcb\x34\xed\x02\x00\x00")

func _000001_users_followersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_users_followersUpSql,
		"000001_users_followers.up.sql",
	)
}

func _000001_users_followersUpSql() (*asset, error) {
	bytes, err := _000001_users_followersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_users_followers.up.sql", size: 749, mode: os.FileMode(420), modTime: time.Unix(1680721003, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_posts_commentsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\xcf\xcd\x4d\xcd\x2b\x29\xb6\xe6\xc2\x2e\x9f\x98\x93\x93\x5f\x9e\x9a\x12\x5f\x50\x94\x59\x96\x58\x92\x1a\x5f\x90\x5f\x8c\x5b\x31\x7e\x49\x90\x09\xc9\x95\xf1\x25\x95\x05\xa9\xd6\x5c\x80\x00\x00\x00\xff\xff\xf1\xfa\xe0\xd3\x8e\x00\x00\x00")

func _000002_posts_commentsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_posts_commentsDownSql,
		"000002_posts_comments.down.sql",
	)
}

func _000002_posts_commentsDownSql() (*asset, error) {
	bytes, err := _000002_posts_commentsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_posts_comments.down.sql", size: 142, mode: os.FileMode(420), modTime: time.Unix(1680720229, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_posts_commentsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x8f\x9b\x30\x10\x3d\xdb\xbf\x62\x94\x13\x48\x54\x6a\x7b\xed\x89\xa6\x93\xc8\x2a\x21\xad\x31\x55\x72\x42\x04\x46\xa9\x25\xbe\x04\x4e\xbb\xf9\xf7\x2b\xb3\x68\x45\xc8\x92\x64\x0f\x7b\xf4\xcc\x7b\xcc\xbc\x37\x8f\xa5\x44\x5f\x21\x28\xff\x7b\x80\x20\x56\x10\x6e\x15\xe0\x4e\x44\x2a\x82\xa6\xee\x4c\x07\x0e\x67\x3a\x07\x11\x2a\x5c\xa3\x84\x5f\x52\x6c\x7c\xb9\x87\x9f\xb8\xf7\x38\x3b\x75\xd4\x26\xa3\xae\x25\x87\x71\x10\x78\x9c\x35\xad\xfe\x97\x66\xe7\xc4\x9c\x1b\x9a\x81\x18\x6d\x0a\x02\x85\x3b\x35\xae\x66\x75\x65\xa8\x32\xd7\xf5\x96\x52\x43\x79\x92\x1a\xf8\xe1\x2b\x54\x62\x83\xe3\xb6\x2e\xd3\x23\x25\x4d\x6a\xfe\xf6\x4c\x8f\x33\xb6\xda\x4a\x14\xeb\xd0\xee\x0a\xce\xb0\xaa\x0b\x9c\x31\x26\x71\x85\x12\xc3\x25\x46\x60\xeb\x9d\xa3\x73\x77\x4a\x98\x08\xb8\x22\x8e\xfb\x3d\xdf\xfd\xc6\x39\xbf\x65\xe7\x88\x30\xb8\xaa\x2b\x43\x47\x6a\xc7\xae\x8e\x35\x55\x69\x49\x60\xe8\xc9\x40\x1c\x8a\xdf\xf1\x85\xde\x9c\xba\xac\xd5\x8d\xd1\x75\xd5\x43\xfa\xf1\x22\x8c\x50\x2a\xeb\xf5\x76\x32\x4e\xe7\x1e\xd8\xcf\xb9\xfc\x8f\x1f\xc4\x18\x01\x77\x3e\x7b\xb0\x68\x4e\x87\x42\x67\x0b\xd7\xe3\xce\x17\xfb\xb4\x1c\x43\xfd\xfb\xab\x07\x8b\xee\x74\xf8\xf4\x5a\xbb\x23\x2f\x2d\x8a\xfa\x3f\xe5\xc9\x80\x4f\xee\xa7\xc7\x22\x66\xa2\x71\x23\x58\xb3\x57\x7d\xfb\xa8\x97\x37\x7d\x99\x38\x85\xf7\xab\x3e\x74\xc3\xac\x2e\x4b\xaa\x3e\x44\xd7\x5c\xee\xaf\x83\x7d\xe7\x4f\x78\xd8\x20\x78\xa7\x43\x30\x58\xf4\x1c\x00\x00\xff\xff\xc4\x24\x0d\xf2\x33\x04\x00\x00")

func _000002_posts_commentsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_posts_commentsUpSql,
		"000002_posts_comments.up.sql",
	)
}

func _000002_posts_commentsUpSql() (*asset, error) {
	bytes, err := _000002_posts_commentsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_posts_comments.up.sql", size: 1075, mode: os.FileMode(420), modTime: time.Unix(1680720216, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_groups_eventsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2f\xca\x2f\x2d\x88\x4f\x2d\x4b\xcd\x2b\x89\x4f\x2c\x29\x49\xcd\x4b\x49\xcc\x4b\x4e\xb5\xe6\x42\x52\x0d\x52\x0e\x55\x5d\x5a\x9c\x5a\x14\x0f\xd6\x52\x8c\x4b\x09\x92\x81\xf8\xd5\x14\x5b\x73\x71\x01\x02\x00\x00\xff\xff\x63\x95\xae\x25\x94\x00\x00\x00")

func _000003_groups_eventsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_groups_eventsDownSql,
		"000003_groups_events.down.sql",
	)
}

func _000003_groups_eventsDownSql() (*asset, error) {
	bytes, err := _000003_groups_eventsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_groups_events.down.sql", size: 148, mode: os.FileMode(420), modTime: time.Unix(1680720288, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_groups_eventsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xcd\x4e\xc3\x30\x10\x84\xcf\xf6\x53\xec\xb1\x91\xfa\x06\x9c\xd2\xb2\xa9\x2c\xd2\x04\x39\x46\x6a\x4f\x51\x94\xac\x2a\x23\x70\xa2\xd8\xe5\xf9\x51\xfe\x20\x85\xd4\x54\x70\xe1\xea\xd9\x99\xd1\xb7\xde\xad\xc4\x50\x21\xa8\x70\x13\x23\x88\x08\x92\x54\x01\x1e\x44\xa6\x32\x38\xb5\xf5\xb9\xb1\x2b\xce\x74\x05\x22\x51\xb8\x43\x09\x8f\x52\xec\x43\x79\x84\x07\x3c\xae\x39\x2b\x5b\x2a\x5c\xdd\xe6\xb3\x81\xce\x9f\x3c\xc5\xf1\x9a\x33\xa7\xdd\x0b\x81\xc2\x83\x9a\xbf\x56\x64\xcb\x56\x37\x4e\xd7\xa6\xd7\xa6\x18\xaa\xf2\xc2\xc1\x7d\xa8\x50\x89\x3d\xce\x1d\x2c\x4a\x25\x8a\x5d\xd2\x95\xc2\xea\xb3\x33\xe0\x8c\x31\x89\x11\x4a\x4c\xb6\x98\xc1\xd9\x52\x6b\x57\xdd\x7b\x70\xc7\xb9\x07\xac\x1b\xcc\x07\x3a\xf0\xe1\xf5\x73\xcb\x6c\xbd\xfb\x8a\xf6\x5c\x6b\x73\x9d\xe6\x02\x66\x6c\x08\x80\x7f\x27\x81\x1e\xe5\x72\x7e\xaa\xfd\x6a\x98\x60\x6e\x80\x1f\x22\xe8\x8d\x8c\xf3\xd3\xfb\x10\x3d\x9b\xf9\xe1\x33\xfb\xe2\xdc\xe9\x57\x5a\x94\x3b\xc1\x36\x85\x99\x92\x6f\x3e\xa3\x7f\xba\xe2\xbc\x70\x8e\x4c\x55\x98\x92\x7e\x7b\x6a\x43\xce\xb2\xa6\xed\x58\xa0\xcd\x09\x36\x69\xfa\x67\xfa\xa9\x6c\x91\xfe\xe3\x6a\xc6\x1d\xbc\x07\x00\x00\xff\xff\x7d\x05\xe7\x30\x3c\x04\x00\x00")

func _000003_groups_eventsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_groups_eventsUpSql,
		"000003_groups_events.up.sql",
	)
}

func _000003_groups_eventsUpSql() (*asset, error) {
	bytes, err := _000003_groups_eventsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_groups_events.up.sql", size: 1084, mode: os.FileMode(420), modTime: time.Unix(1680720335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_messagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xc3\xcb\x53\x7a\x20\x00\x00\x00")

func _000004_messagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_messagesDownSql,
		"000004_messages.down.sql",
	)
}

func _000004_messagesDownSql() (*asset, error) {
	bytes, err := _000004_messagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_messages.down.sql", size: 32, mode: os.FileMode(420), modTime: time.Unix(1680720283, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_messagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x4d\x6a\xc3\x30\x10\x85\xd7\xd2\x29\x66\x69\x83\x6f\xd0\x95\x9a\x8e\x83\xa8\xa3\x14\x79\x0a\xc9\xca\x88\x68\x70\xb5\x88\x62\x24\xe5\xfe\x25\xe9\x0f\x8a\xa1\xd0\xed\x7c\x7c\xf3\xde\xdb\x58\x54\x84\x40\xea\x79\x40\xd0\x3d\x98\x3d\x01\x1e\xf4\x48\x23\x9c\x39\x67\x37\x73\x6e\xa4\x08\x1e\xb4\x21\xdc\xa2\x85\x37\xab\x77\xca\x1e\xe1\x15\x8f\x9d\x14\x99\xa3\xe7\x34\x55\xfc\xf6\xc0\xbc\x0f\x43\x27\x45\xe2\x53\x58\x02\xc7\x52\xf1\x4e\x8a\x39\x5d\xae\xcb\xe3\xe9\x74\x89\x85\x63\x01\xc2\x03\xd5\x1f\xc2\xd9\xcd\x3c\x2d\xae\x7c\xdc\xd1\x57\x60\x99\x5c\x81\x17\x45\x48\x7a\x87\x8f\x79\xce\xff\xc5\xfa\xbd\x45\xbd\x35\xb7\xda\xd0\xfc\xb6\x6e\x41\x0a\x61\xb1\x47\x8b\x66\x83\x23\x5c\x33\xa7\x0c\x4d\xf0\xed\xca\xa8\xa7\xfc\x5b\xfa\x19\xba\x16\xee\xf7\x6f\xa3\x7d\xfa\x0c\x00\x00\xff\xff\xb1\xce\xdf\xdb\x82\x01\x00\x00")

func _000004_messagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_messagesUpSql,
		"000004_messages.up.sql",
	)
}

func _000004_messagesUpSql() (*asset, error) {
	bytes, err := _000004_messagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_messages.up.sql", size: 386, mode: os.FileMode(420), modTime: time.Unix(1680720324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_notificationsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\xa9\x2c\x48\x2d\xb6\xe6\x22\x42\x65\x4a\x6a\x49\x62\x66\x0e\x51\x6a\x8b\xad\x01\x01\x00\x00\xff\xff\xd5\x93\xef\x6f\x7a\x00\x00\x00")

func _000005_notificationsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_notificationsDownSql,
		"000005_notifications.down.sql",
	)
}

func _000005_notificationsDownSql() (*asset, error) {
	bytes, err := _000005_notificationsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_notifications.down.sql", size: 122, mode: os.FileMode(420), modTime: time.Unix(1680720295, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_notificationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xcb\x6e\xdb\x30\x10\x3c\x9b\x5f\xb1\xf0\x49\x02\x84\xa2\x8f\x63\x4e\x8a\xbb\x0e\x84\xda\x52\x40\xd1\x45\x72\x12\x08\x69\xdd\xb0\x70\x28\x57\xa4\x5d\x18\x46\xff\xbd\x20\x25\xa7\x8a\x23\xd7\x42\x4e\x7c\xec\xec\xcc\x78\xbc\xd4\x8c\x63\x2c\x10\x44\x7c\xbb\x40\x48\xe6\x90\x66\x02\xf0\x21\xc9\x45\x0e\xba\xb6\x6a\xad\x4a\x69\x55\xad\x0b\x7b\xd8\x92\x09\xd8\x44\x55\x90\xa4\x02\xef\x90\xc3\x3d\x4f\x96\x31\x7f\x84\x6f\xf8\x18\x31\x00\x00\xd2\x56\xd9\x03\x08\x7c\x10\x9e\x27\x5d\x2d\x16\x6d\xa5\x22\x53\x36\x6a\xeb\x98\x86\xca\x96\x9e\xb7\x1b\x69\xe9\x75\x8d\x85\x37\x8c\x25\x69\x8e\x5c\x38\xcd\x6c\xc0\x10\x04\xaa\x8a\x3a\xdd\xa8\xaf\x12\xbd\x70\x86\xec\x7b\xbc\x58\x61\x0e\x2c\xf8\x18\xc1\x74\x67\xa8\x31\xd3\x6e\x03\xeb\x7a\xb3\xa9\x7f\x43\x43\xbf\x76\x64\xac\xbb\x5e\xb9\xeb\xa3\x2b\x6a\xf9\x4c\x7f\xe0\x49\x1a\x30\xa4\x2d\x1c\xea\x1d\xc8\xb3\x86\x0f\xd3\x30\x62\xc1\xa7\x08\xa6\x3f\x9a\x7a\xb7\xf5\xbc\x7e\x07\x4a\xef\x95\xf5\x46\x87\x48\x7d\x95\x8c\xe7\xb4\x35\xfc\xac\x95\x06\xfb\x44\xd0\xf6\x1e\xfd\xe2\x91\x9e\xff\xf3\x00\xbf\x6f\x19\x6f\xbb\x43\x8e\x10\xfb\x72\x92\x28\x68\x4f\xda\xf6\x24\xfd\xf9\x92\x56\xd9\x90\xb4\x54\x81\xd4\x2d\x0e\x8e\x7e\x69\x01\xeb\xba\xb9\xa4\x78\xc3\x18\x63\x63\x67\xb0\x22\x2b\xd5\xe6\xda\x14\x1a\xd2\x15\x35\x45\x0f\xf2\x7a\xd8\xde\x4c\xd1\x65\x68\x3b\x58\xc3\xf5\x49\xf7\x93\x0b\x69\x61\xb6\xe2\x1c\x53\x51\x88\x64\x89\xb9\x88\x97\xf7\xf0\x35\x16\xe8\x4e\xfd\x86\x79\xc6\x31\xb9\x4b\x9d\x53\x08\x5e\x5c\x86\xc0\x26\x13\x8e\x73\xe4\x98\xce\x30\x07\x3f\xa0\x6e\xae\xc3\xb3\x8e\x21\xdf\xe7\xcd\xc3\x2f\x24\xf4\x2f\x69\x6c\xcc\xd7\xf2\x6d\xa8\x24\xb5\x1f\x9d\x70\xf7\xa7\x5d\x46\x1b\x22\xed\x42\x3c\x45\x76\x52\x91\xa5\xff\x5c\xdc\x66\xd9\x9b\xf0\x7a\x16\xde\x17\xdf\x3f\x53\xff\x4d\xb0\x83\x9d\x32\xfc\x1b\x00\x00\xff\xff\xbc\x24\xac\xb4\x2c\x05\x00\x00")

func _000005_notificationsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_notificationsUpSql,
		"000005_notifications.up.sql",
	)
}

func _000005_notificationsUpSql() (*asset, error) {
	bytes, err := _000005_notificationsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_notifications.up.sql", size: 1324, mode: os.FileMode(420), modTime: time.Unix(1680720343, 0)}
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
	"000001_users_followers.down.sql": _000001_users_followersDownSql,
	"000001_users_followers.up.sql": _000001_users_followersUpSql,
	"000002_posts_comments.down.sql": _000002_posts_commentsDownSql,
	"000002_posts_comments.up.sql": _000002_posts_commentsUpSql,
	"000003_groups_events.down.sql": _000003_groups_eventsDownSql,
	"000003_groups_events.up.sql": _000003_groups_eventsUpSql,
	"000004_messages.down.sql": _000004_messagesDownSql,
	"000004_messages.up.sql": _000004_messagesUpSql,
	"000005_notifications.down.sql": _000005_notificationsDownSql,
	"000005_notifications.up.sql": _000005_notificationsUpSql,
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
	"000001_users_followers.down.sql": &bintree{_000001_users_followersDownSql, map[string]*bintree{}},
	"000001_users_followers.up.sql": &bintree{_000001_users_followersUpSql, map[string]*bintree{}},
	"000002_posts_comments.down.sql": &bintree{_000002_posts_commentsDownSql, map[string]*bintree{}},
	"000002_posts_comments.up.sql": &bintree{_000002_posts_commentsUpSql, map[string]*bintree{}},
	"000003_groups_events.down.sql": &bintree{_000003_groups_eventsDownSql, map[string]*bintree{}},
	"000003_groups_events.up.sql": &bintree{_000003_groups_eventsUpSql, map[string]*bintree{}},
	"000004_messages.down.sql": &bintree{_000004_messagesDownSql, map[string]*bintree{}},
	"000004_messages.up.sql": &bintree{_000004_messagesUpSql, map[string]*bintree{}},
	"000005_notifications.down.sql": &bintree{_000005_notificationsDownSql, map[string]*bintree{}},
	"000005_notifications.up.sql": &bintree{_000005_notificationsUpSql, map[string]*bintree{}},
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

