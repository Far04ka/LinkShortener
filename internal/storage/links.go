package storage

import "errors"

type Repo interface {
	GetLink(id string) (string, error)
	GetId(url string) string
	SetLink(id string, url string) bool
}

type SimpleStorage struct {
	Lnks map[string]string
}

func (stor *SimpleStorage) GetLink(id string) (string, error) {
	for key, val := range stor.Lnks {
		if val == id {
			return key, nil
		}
	}
	return "", errors.New("not found")
}

func (stor *SimpleStorage) GetId(url string) string {
	for key, val := range stor.Lnks {
		if key == url {
			return val
		}
	}
	return ""
}

func (stor *SimpleStorage) SetLink(id string, url string) bool {
	for _, val := range stor.Lnks {
		if val == id {
			return false
		}
	}
	stor.Lnks[url] = id
	return true
}

var Storage SimpleStorage
