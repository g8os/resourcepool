package client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
)

type AysRepositoryRepositoryBlueprintPostRespBody struct {
	Content json.RawMessage `json:"content" validate:"nonzero"`
	Name    string          `json:"name" validate:"nonzero"`
}

func (s AysRepositoryRepositoryBlueprintPostRespBody) Validate() error {

	return validator.Validate(s)
}
