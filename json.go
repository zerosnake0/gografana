package gografana

import (
	"net/http"

	"github.com/zerosnake0/jzon"
)

var (
	jzonDecoder = jzon.NewDecoder(&jzon.DecoderOption{
		CaseSensitive:         true,
		OnlyTaggedField:       true,
		DisallowUnknownFields: false,
	})
	jzonEncoder = jzon.NewEncoder(&jzon.EncoderOption{
		EscapeHTML:      true,
		OnlyTaggedField: true,
	})
)

func jsonUnmarshal(data []byte, o interface{}) error {
	return jzonDecoder.Unmarshal(data, o)
}

func jsonMarshal(o interface{}) ([]byte, error) {
	return jzonEncoder.Marshal(o)
}

func setRequestContentTypeJson(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}
