package state

import (
	"encoding/hex"
	"encoding/json"
)

type Shim struct {
	Sha256 []byte
	Path   string
}

func (s *Shim) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Sha256 string `json:"sha256"`
		Path   string `json:"path"`
	}{
		Sha256: hex.EncodeToString(s.Sha256),
		Path:   s.Path,
	})
}

func (s *Shim) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Sha256 string `json:"sha256"`
		Path   string `json:"path"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.Path = aux.Path
	sha256, err := hex.DecodeString(aux.Sha256)
	if err != nil {
		return err
	}
	s.Sha256 = sha256
	return nil
}
