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
// api/pkg/db/migrations/sqlite/000006_alter_posts_followers_groups.down.sql
// api/pkg/db/migrations/sqlite/000006_alter_posts_followers_groups.up.sql
// api/pkg/db/migrations/sqlite/000007_add_image_to_users_and_groups.down.sql
// api/pkg/db/migrations/sqlite/000007_add_image_to_users_and_groups.up.sql
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

var __000001_users_followersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xcd\x8e\x82\x30\x14\x85\xd7\xed\x53\xdc\xa5\x26\xbc\x81\x2b\xd4\x8b\x69\x06\x61\xa6\x96\x44\x57\xa4\x40\xd5\xc6\x4a\x09\x2d\x63\xe6\xed\x27\x18\x9d\xc1\xf8\xb7\xec\xb9\xe7\x9c\xf6\xbb\x9d\x71\x0c\x05\x82\x08\xa7\x31\x02\x8b\x20\x49\x05\xe0\x9a\xad\xc4\x0a\x3a\xa7\x5a\x07\x23\x4a\x74\x05\x2c\x11\xb8\x40\x0e\x9f\x9c\x2d\x43\xbe\x81\x0f\xdc\x04\x94\x6c\x6d\x5b\xcb\xa3\x02\x81\x6b\x71\x4e\x26\x59\x1c\x07\x94\xb8\xee\xb1\xae\x8e\x52\x9b\x5b\x15\xb2\x84\x7d\x65\x18\x50\xd2\x48\xe7\x4e\xb6\xad\xee\x52\x85\x6e\xfd\xbe\x92\x3f\x30\x0f\x05\x0a\xb6\xec\xcd\xb5\x2e\x0f\x7f\x57\x04\x94\xc8\xc2\x76\xfe\x7a\xd0\x47\xb9\x53\x79\x23\xfd\xfe\xaa\x94\xad\x92\x5e\x55\xb9\xf4\xc3\x12\xed\xf2\xa6\x2b\x8c\x2e\x61\x9a\xa6\xf1\xff\x93\xe6\x18\x85\x59\x2c\x60\x2b\x8d\x53\x74\x3c\xa1\xf4\xcd\x96\x72\xa7\x9c\xd3\xb6\x76\xaf\x96\x75\x36\x0e\xa6\x03\x42\x6f\x0f\xaa\x1e\x72\x1b\xf3\xf8\xd1\xc3\x10\x89\x52\x8e\x6c\x91\xf4\xf5\x30\xba\xb4\x8f\x29\x21\x84\x63\x84\x1c\x93\x19\x5e\xfe\x70\xd4\xeb\x67\x8e\x57\x20\x5b\x6b\x8c\x3d\xf5\xf6\x97\x3f\xde\x9b\x74\xbd\x7b\x42\x72\x2d\x79\x32\x96\x65\xa9\x1a\xaf\x2a\x28\xac\x35\xb7\x03\xaf\xbf\xd5\x9d\x7c\xcb\x38\xbc\xfc\x29\x28\x79\x9c\x7a\xb3\x9d\xf1\x84\xfe\x06\x00\x00\xff\xff\x70\xc3\x45\xc3\x0a\x03\x00\x00")

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

	info := bindataFileInfo{name: "000001_users_followers.up.sql", size: 778, mode: os.FileMode(420), modTime: time.Unix(1681673427, 0)}
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

var __000002_posts_commentsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x8f\x9b\x30\x10\x3d\xdb\xbf\x62\x94\x13\x48\xf4\xd0\xf6\xd8\x13\x4d\x27\x91\x55\x42\x5a\x63\xaa\xe4\x84\x08\x8c\x52\x4b\x7c\x09\x9c\x76\xf3\xef\x57\x66\xd1\x8a\x90\x25\xc9\x1e\xf6\xe8\x99\xf7\x98\x79\x6f\x1e\x4b\x89\xbe\x42\x50\xfe\xf7\x00\x41\xac\x20\xdc\x2a\xc0\x9d\x88\x54\x04\x4d\xdd\x99\x0e\x1c\xce\x74\x0e\x22\x54\xb8\x46\x09\xbf\xa4\xd8\xf8\x72\x0f\x3f\x71\xef\x71\x76\xea\xa8\x4d\x46\x5d\x4b\x0e\xe3\x20\xf0\x38\x6b\x5a\xfd\x2f\xcd\xce\x89\x39\x37\x34\x03\x31\xda\x14\x04\x0a\x77\x6a\x5c\xcd\xea\xca\x50\x65\xae\xeb\x2d\xa5\x86\xf2\x24\x35\xf0\xc3\x57\xa8\xc4\x06\xc7\x6d\x5d\xa6\x47\x4a\x9a\xd4\xfc\xed\x99\x1e\x67\x6c\xb5\x95\x28\xd6\xa1\xdd\x15\x9c\x61\x55\x17\x38\x63\x4c\xe2\x0a\x25\x86\x4b\x8c\xc0\xd6\x3b\x47\xe7\xee\x94\x30\x11\x70\x45\x1c\xf7\x7b\xbe\xfb\x8d\x73\x7e\xcb\xce\x11\x61\x70\x55\x57\x86\x8e\xd4\x8e\x5d\x1d\x6b\xaa\xd2\x92\xc0\xd0\x93\x81\x38\x14\xbf\xe3\x0b\xbd\x39\x75\x59\xab\x1b\xa3\xeb\xaa\x87\xf4\xe3\x45\x18\xa1\x54\xd6\xeb\xed\x64\x9c\xce\x3d\xb0\x9f\x73\xf9\x1f\x3f\x88\x31\x02\xee\x7c\xf6\x60\xd1\x9c\x0e\x85\xce\x16\xae\xc7\x9d\x2f\xf6\x69\x39\x86\xfa\xf7\x57\x0f\x16\xdd\xe9\xf0\xe9\xb5\x76\x47\x5e\x5a\x14\xf5\x7f\xca\x93\x01\x9f\xdc\x4f\x8f\x45\xcc\x44\xe3\x46\xb0\x66\xaf\xfa\xf6\x51\x2f\x6f\xfa\x32\x71\x0a\xef\x57\x7d\xe8\x86\x59\x5d\x96\x54\x7d\x88\xae\xb9\xdc\x5f\x07\xfb\xce\x9f\xf0\xb0\x41\xf0\x4e\x87\x60\xb0\xe8\x39\x00\x00\xff\xff\x0e\xf6\x62\xe4\x33\x04\x00\x00")

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

	info := bindataFileInfo{name: "000002_posts_comments.up.sql", size: 1075, mode: os.FileMode(420), modTime: time.Unix(1688555718, 0)}
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

	info := bindataFileInfo{name: "000003_groups_events.up.sql", size: 1084, mode: os.FileMode(420), modTime: time.Unix(1685890554, 0)}
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

var __000004_messagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x4d\x6e\xc2\x30\x10\x85\xd7\xf6\x29\x66\x99\x48\xb9\x41\x57\x2e\x9d\x20\xab\xc1\x54\xce\x54\x82\x55\x64\xe1\x51\xea\x05\x26\xb2\xcd\xfd\x2b\xe8\x8f\x4c\x56\x6c\xe7\xd3\x37\xef\xbd\x8d\x45\x45\x08\xa4\x5e\x07\x04\xdd\x83\xd9\x13\xe0\x41\x8f\x34\xc2\x99\x73\x76\x33\xe7\x46\x8a\xe0\x41\x1b\xc2\x2d\x5a\xf8\xb0\x7a\xa7\xec\x11\xde\xf1\xd8\x49\x91\x39\x7a\x4e\x53\xc5\x6f\x0f\xcc\xe7\x30\x74\x52\x24\x3e\x85\x25\x70\x2c\x15\xef\xa4\x98\xd3\xe5\xba\x3c\x9e\x4e\x97\x58\x38\x16\x20\x3c\x50\xfd\x21\x9c\xdd\xcc\xd3\xe2\xca\xd7\x1d\xfd\x04\x96\xc9\x15\x78\x53\x84\xa4\x77\xf8\x98\xe7\x7c\xcd\x3a\x29\xfa\xbd\x45\xbd\x35\xb7\xb6\xd0\xfc\x97\x6d\x41\x0a\x61\xb1\x47\x8b\x66\x83\x23\x5c\x33\xa7\x0c\x4d\xf0\xed\xca\xa8\x17\x3c\x2d\xfd\xed\x5b\x0b\xf7\xfb\xaf\xd1\xbe\x7c\x07\x00\x00\xff\xff\xea\x36\x03\x21\x79\x01\x00\x00")

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

	info := bindataFileInfo{name: "000004_messages.up.sql", size: 377, mode: os.FileMode(420), modTime: time.Unix(1688448572, 0)}
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

var __000005_notificationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4f\x6f\x82\x40\x10\xc5\xcf\xec\xa7\x98\x70\x82\x84\x43\xff\x1c\x7b\x42\x3b\x1a\x52\x04\xb3\xac\x8d\x9e\x08\x81\xb1\xd9\xc4\x2e\x96\x5d\x6d\xfc\xf6\x0d\xf8\xa7\x88\x58\x4d\x8f\xcc\x7b\xf3\xf6\xf1\xdb\x1d\x72\xf4\x05\x82\xf0\x07\x21\x42\x30\x82\x28\x16\x80\xf3\x20\x11\x09\xa8\xd2\xc8\xa5\xcc\x33\x23\x4b\x95\x9a\xdd\x9a\xb4\xc3\x2c\x59\x40\x10\x09\x1c\x23\x87\x29\x0f\x26\x3e\x5f\xc0\x1b\x2e\x3c\x06\x00\xa0\xb2\x4f\x02\x81\x73\xd1\xa4\x44\xb3\x30\xdc\xcf\x49\x19\x69\x76\xe7\x0a\x73\x5f\x18\x0b\xa2\x04\xb9\xa8\x03\xe3\x9e\xd3\xc0\x91\x85\xd7\x84\x7a\x87\x08\x97\xbd\xfb\xe1\x0c\x13\x60\xce\x83\x07\xf6\xb2\x5c\xad\xca\xef\xb4\xa2\xaf\x0d\x69\x63\x9f\x26\x54\x69\xdb\xf5\x98\xf3\xe8\x81\xfd\x51\x95\x9b\x75\xdb\xd2\x0c\xf6\xfa\xd3\x49\x97\x6a\x2b\x0d\x75\xe4\x67\x0f\x6c\xda\x92\x32\x5d\x39\x6d\xa6\xda\xae\x7f\x81\xb1\x7b\x09\x16\x64\x32\xb9\xba\xc5\x50\x93\x2a\xa8\x4a\x5b\x96\x73\x96\x17\x98\xae\x5b\xf7\xcc\xfa\x75\x2b\xaf\x28\x33\x54\xa4\x99\x81\xe1\x8c\x73\x8c\x44\x2a\x82\x09\x26\xc2\x9f\x4c\xe1\xd5\x17\x58\x7f\xb5\x17\x46\x31\xc7\x60\x1c\xd5\x4d\xc1\x39\xb5\x74\x81\x59\x16\xc7\x11\x72\x8c\x86\x98\xc0\x46\x53\xd5\x5c\x9c\xdb\xd9\xe8\xeb\xdd\x5d\xee\x7f\x02\x6e\xf3\x54\xee\xc5\x7c\x8b\x6f\x45\x39\xc9\xed\xdd\x84\x0f\x97\x76\xdd\xad\x89\x54\x0d\xf1\x88\xec\x78\x4a\x96\xd7\xfb\x30\x88\xe3\x0b\x78\xad\x0a\xff\xc3\xf7\x5b\xea\x4f\x82\x07\xdb\x91\xe1\x4f\x00\x00\x00\xff\xff\x06\x1f\x7e\x65\xea\x03\x00\x00")

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

	info := bindataFileInfo{name: "000005_notifications.up.sql", size: 1002, mode: os.FileMode(420), modTime: time.Unix(1688143474, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_alter_posts_followers_groupsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc9\x2c\xc9\x49\xb5\xe6\xe2\xc2\x54\xe4\x12\xe4\x1f\x00\x53\x95\x5e\x94\x5f\x5a\x10\x9f\x99\x82\xa6\x10\x2c\x8c\xaa\x32\x33\x37\x31\x3d\x35\xbe\x20\xb1\x24\x03\x4d\x6d\x5a\x7e\x4e\x4e\x7e\x79\x6a\x11\x8a\xed\x89\xc9\x25\x99\x65\xa9\x0a\x4e\xfe\xfe\x3e\xd6\x80\x00\x00\x00\xff\xff\xd1\x71\x90\xd4\xa7\x00\x00\x00")

func _000006_alter_posts_followers_groupsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_alter_posts_followers_groupsDownSql,
		"000006_alter_posts_followers_groups.down.sql",
	)
}

func _000006_alter_posts_followers_groupsDownSql() (*asset, error) {
	bytes, err := _000006_alter_posts_followers_groupsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_alter_posts_followers_groups.down.sql", size: 167, mode: os.FileMode(420), modTime: time.Unix(1688555718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_alter_posts_followers_groupsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x2c\xf5\x0c\x59\xc5\x66\x14\x21\xa6\x12\x23\xb8\x2b\xc1\xc6\x3a\x10\x49\x68\x46\xbd\xbe\x8b\x22\x58\xbb\x7e\xef\x3f\xbe\x32\x1e\x1d\x78\xb5\x31\x08\x25\x57\xae\xa0\x5d\x7b\x84\xa6\x35\xe7\x83\x05\x26\x4e\x51\x0a\xb1\xb4\x84\xd2\xfa\x6b\x0d\x63\x7e\x96\x8e\x7a\xd8\x5b\x8f\x3b\x74\x20\x1c\x6e\xd1\xa1\x6d\xf0\x34\xc1\xba\xa2\x7e\x2d\xc5\x3c\x34\x91\xdf\x10\x3d\xc2\x10\xbb\x12\xf8\x0e\x1e\x2f\xfe\x7f\x70\xcb\x29\xe5\x77\x1c\xe7\x1f\xc3\x95\xe9\x15\xe5\x27\x00\x00\xff\xff\xd2\x1d\x6f\x6b\xc9\x00\x00\x00")

func _000006_alter_posts_followers_groupsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_alter_posts_followers_groupsUpSql,
		"000006_alter_posts_followers_groups.up.sql",
	)
}

func _000006_alter_posts_followers_groupsUpSql() (*asset, error) {
	bytes, err := _000006_alter_posts_followers_groupsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_alter_posts_followers_groups.up.sql", size: 201, mode: os.FileMode(420), modTime: time.Unix(1688555718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_add_image_to_users_and_groupsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcc\x4d\x4c\x4f\x8d\x4f\x4a\x2c\x4e\x35\x33\xb1\xe6\xe2\x42\x56\x9c\x5e\x94\x5f\x5a\x80\x4f\x35\x20\x00\x00\xff\xff\x8c\xc7\x37\xbb\x5a\x00\x00\x00")

func _000007_add_image_to_users_and_groupsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_add_image_to_users_and_groupsDownSql,
		"000007_add_image_to_users_and_groups.down.sql",
	)
}

func _000007_add_image_to_users_and_groupsDownSql() (*asset, error) {
	bytes, err := _000007_add_image_to_users_and_groupsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_add_image_to_users_and_groups.down.sql", size: 90, mode: os.FileMode(420), modTime: time.Unix(1688570748, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_add_image_to_users_and_groupsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2f\xca\x2f\x2d\x28\x56\xe0\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcc\x4d\x4c\x4f\x8d\x4f\x4a\x2c\x4e\x35\x33\x51\x08\x71\x8d\x08\xb1\xe6\xe2\x42\xd6\x52\x5a\x9c\x5a\x54\x8c\x5f\x03\x20\x00\x00\xff\xff\x91\x8b\x7c\x4f\x62\x00\x00\x00")

func _000007_add_image_to_users_and_groupsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_add_image_to_users_and_groupsUpSql,
		"000007_add_image_to_users_and_groups.up.sql",
	)
}

func _000007_add_image_to_users_and_groupsUpSql() (*asset, error) {
	bytes, err := _000007_add_image_to_users_and_groupsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_add_image_to_users_and_groups.up.sql", size: 98, mode: os.FileMode(420), modTime: time.Unix(1688570699, 0)}
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
	"000006_alter_posts_followers_groups.down.sql": _000006_alter_posts_followers_groupsDownSql,
	"000006_alter_posts_followers_groups.up.sql": _000006_alter_posts_followers_groupsUpSql,
	"000007_add_image_to_users_and_groups.down.sql": _000007_add_image_to_users_and_groupsDownSql,
	"000007_add_image_to_users_and_groups.up.sql": _000007_add_image_to_users_and_groupsUpSql,
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
	"000006_alter_posts_followers_groups.down.sql": &bintree{_000006_alter_posts_followers_groupsDownSql, map[string]*bintree{}},
	"000006_alter_posts_followers_groups.up.sql": &bintree{_000006_alter_posts_followers_groupsUpSql, map[string]*bintree{}},
	"000007_add_image_to_users_and_groups.down.sql": &bintree{_000007_add_image_to_users_and_groupsDownSql, map[string]*bintree{}},
	"000007_add_image_to_users_and_groups.up.sql": &bintree{_000007_add_image_to_users_and_groupsUpSql, map[string]*bintree{}},
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

