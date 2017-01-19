package rubydebug

import (
	"encoding/json"
	"testing"
)

func TestToJSON(t *testing.T) {
	cases := []struct {
		rubydebug string
		json      string
	}{
		{
			rubydebug: `{
       "message" => "test",
      "@version" => "1",
    "@timestamp" => "2017-01-19T19:19:55.517Z",
           "bla" => {
        "foo" => "bar",
        "baz" => [
            [0] "bar",
            [1] "123"
        ]
    },
          "host" => "localhost",
          "tags" => [
        [0] "tag"
    ]
}`,
			json: `{"@timestamp":"2017-01-19T19:19:55.517Z","@version":"1","bla":{"baz":["bar","123"],"foo":"bar"},"host":"localhost","message":"test","tags":["tag"]}`,
		},
		{
			rubydebug: `{
       "message" => "test",
      "@version" => "1",
    "@timestamp" => "2017-01-19T19:28:11.720Z",
           "bla" => {
        "foo" => "bar",
        "baz" => [
            [0] "bar",
            [1] "123"
        ]
    },
          "host" => "localhost",
        "number" => "456",
          "tags" => [
        [0] "tag"
    ],
     "@metadata" => {
        "test" => "metadatafield"
    }
}`,
			json: `{"@metadata":{"test":"metadatafield"},"@timestamp":"2017-01-19T19:28:11.720Z","@version":"1","bla":{"baz":["bar","123"],"foo":"bar"},"host":"localhost","message":"test","number":"456","tags":["tag"]}`,
		},
	}
	for _, test := range cases {
		toJSON := ToJSON(test.rubydebug)

		var jsonUnmarshal interface{}
		err := json.Unmarshal([]byte(toJSON), &jsonUnmarshal)
		if err != nil {
			t.Errorf("Failed to unmarshal json with error: %s, %s", err, toJSON)
		}

		jsonMarshal, err := json.Marshal(jsonUnmarshal)
		got := string(jsonMarshal)
		if err != nil {
			t.Error(err)
		}
		if got != test.json {
			t.Errorf("Expected %s to equal %s with input %s", got, test.json, test.rubydebug)
		}
	}
}
