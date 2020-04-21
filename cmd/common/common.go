package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func MustParseConfig(conf interface{}, prefix, path string) {
	if len(path) == 0 {
		if err := envconfig.Process(prefix, conf); err != nil {
			fmt.Printf("failed to get config from env: %v", err)
			os.Exit(1)
		}
	}
	f, err := os.Open(MaybeSymlink(path))
	if err != nil {
		fmt.Printf("can't open config at %s. err %v\n", path, err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("can't read config at %s. err %v\n", path, err)
		os.Exit(1)
	}
	_ = f.Close()
	if err := json.Unmarshal(data, conf); err != nil {
		fmt.Printf("can't unmarshal config %s data. err %v\n", path, err)
		os.Exit(1)
	}
}

func MaybeSymlink(path string) string {
	path = strings.TrimSpace(path)
	sympath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path
	}
	return sympath
}

func MustDecryptKey(log *logrus.Entry, path, pwpath string) *keystore.Key {
	f, err := os.Open(MaybeSymlink(path))
	if err != nil {
		log.Fatalf("faield to open keyfile %s: %v", path, err)
	}
	data, err := ioutil.ReadAll(f)
	_ = f.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	password := ""
	if len(pwpath) != 0 {
		f, err = os.Open(pwpath)
		if err != nil {
			log.Fatalf("can't open %s: %v", pwpath, err)
		}
		data, err := ioutil.ReadAll(f)
		_ = f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		password = strings.TrimRightFunc(string(data), func(r rune) bool {
			return r == '\r' || r == '\n'
		})
	}
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		log.Fatalf("failed to decrypt a key: %v", err)
	}
	return key
}
