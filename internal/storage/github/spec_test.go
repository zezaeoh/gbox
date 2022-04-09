package github

import (
	"reflect"
	"testing"
)

func Test_wrapSpec(t *testing.T) {
	tests := []struct {
		name    string
		spec    map[string]interface{}
		want    spec
		wantErr bool
	}{
		{
			name: "should parse owner, repo from url 01",
			spec: map[string]interface{}{
				"url":    "https://github.com/zezaeoh/gbox-test",
				"branch": "main",
			},
			want: spec{
				URL:    "https://github.com/zezaeoh/gbox-test",
				Branch: "main",
				owner:  "zezaeoh",
				repo:   "gbox-test",
			},
		},
		{
			name: "should parse owner, repo from url 02",
			spec: map[string]interface{}{
				"url":    "https://github.com/zezaeoh/gbox-test.git",
				"branch": "main",
			},
			want: spec{
				URL:    "https://github.com/zezaeoh/gbox-test.git",
				Branch: "main",
				owner:  "zezaeoh",
				repo:   "gbox-test",
			},
		},
		{
			name: "need token if auth type is https",
			spec: map[string]interface{}{
				"url":      "https://github.com/zezaeoh/gbox-test",
				"branch":   "main",
				"authType": "https",
			},
			wantErr: true,
		},
		{
			name: "need url",
			spec: map[string]interface{}{
				"branch":   "main",
				"authType": "https",
			},
			wantErr: true,
		},
		{
			name: "need branch",
			spec: map[string]interface{}{
				"url":      "https://github.com/zezaeoh/gbox-test",
				"authType": "https",
			},
			wantErr: true,
		},
		{
			name: "unknown auth type",
			spec: map[string]interface{}{
				"url":      "https://github.com/zezaeoh/gbox-test",
				"branch":   "main",
				"authType": "test",
			},
			wantErr: true,
		},
		{
			name: "invalid url 01",
			spec: map[string]interface{}{
				"url":      "https://zezaeoh.io/zezaeoh/gbox-test",
				"branch":   "main",
				"authType": "https",
			},
			wantErr: true,
		},
		{
			name: "invalid url 02",
			spec: map[string]interface{}{
				"url":      "https://github.com/zezaeoh",
				"branch":   "main",
				"authType": "https",
			},
			wantErr: true,
		},
		{
			name: "invalid url 03",
			spec: map[string]interface{}{
				"url":      "https://github.com/zezaeoh/gbox-test/gg",
				"branch":   "main",
				"authType": "https",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wrapSpec(tt.spec)
			if tt.wantErr {
				if err == nil {
					t.Error("wrapSpec(m) want error but passed")
				}
			} else {
				if err != nil {
					t.Errorf("wrapSpec(m) return err %v, want %v", err, tt.want)
				} else if !reflect.DeepEqual(*got, tt.want) {
					t.Errorf("wrapSpec(m) => %v, want %v", *got, tt.want)
				}
			}
		})
	}
}
