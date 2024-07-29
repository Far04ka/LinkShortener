package conf

import (
	"errors"
	"flag"
	"strconv"
	"strings"

	"github.com/Far04ka/LinkShortener/internal/storage"
)

type AddrField struct {
	Val string
}

func (field *AddrField) String() string {
	return field.Val
}

func (field *AddrField) Set(val string) error {
	addr := val
	if strings.Contains(val, "//") {
		addr = strings.Split(val, "//")[1]
	}

	url, port := strings.Split(addr, ":")[0], strings.Split(addr, ":")[1]
	if url != storage.BaseURL {
		return errors.New("wrong url")
	} else if _, err := strconv.Atoi(port); err != nil {
		return errors.New("wrong port")
	}
	field.Val = addr
	return nil
}

type FinalAddrField struct {
	Val       string
	ShortAddr string
}

func (field *FinalAddrField) String() string {
	return field.Val
}

func (field *FinalAddrField) Set(val string) error {

	mainAddr := ""
	addr := ""
	if strings.Contains(val, "//") {
		mainAddr, addr = strings.Split(strings.Split(val, "//")[1], "/")[0], strings.Join(strings.Split(strings.Split(val, "//")[1], "/")[1:], "")
	}
	if mainAddr != Conf.Addr.Val {
		return errors.New("wrong base addres")
	}

	field.Val = val + "/"
	field.ShortAddr = "/" + addr + "/"
	return nil
}

type Config struct {
	Addr      *AddrField
	Finaladdr *FinalAddrField
}

var Conf *Config

func CreateConfig() {
	flag.Var(Conf.Addr, "a", "base URL")
	flag.Var(Conf.Finaladdr, "b", "URL before short link")
	flag.Parse()
}
