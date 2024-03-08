package state

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Shim struct {
	Sha256 []byte
	Path   string
}

func (s *Shim) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(&struct {
		Sha256 string `json:"sha256"`
		Path   string `json:"path"`
	}{
		Sha256: hex.EncodeToString(s.Sha256),
		Path:   s.Path,
	})
	return j, fmt.Errorf("error marshalling shim: %w", err)
}

func (s *Shim) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Sha256 string `json:"sha256"`
		Path   string `json:"path"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("error unmarshalling shim: %w", err)
	}
	s.Path = aux.Path
	sha256, err := hex.DecodeString(aux.Sha256)
	if err != nil {
		return fmt.Errorf("error decoding sha256: %w", err)
	}
	s.Sha256 = sha256
	return nil
}
