package main

import (
	"github.com/buildkite/yaml"
	"testing"
)

type (
	// Build configures a Docker build.
	Build struct {
		Args       map[string]string `json:"args,omitempty"`
		CacheFrom  []string          `json:"cache_from,omitempty" yaml:"cache_from"`
		Context    string            `json:"context,omitempty"`
		Dockerfile string            `json:"dockerfile,omitempty"`
		Image      string            `json:"image,omitempty"`
		Labels     map[string]string `json:"labels,omitempty"`
	}

	// build is a tempoary type used to unmarshal
	// the Build struct when long format is used.
	build struct {
		Args       map[string]string
		CacheFrom  []string `yaml:"cache_from"`
		Context    string
		Dockerfile string
		Image      string
		Labels     map[string]string
	}
)

func TestBuild(t *testing.T) {
	tests := []struct {
		yaml       string
		image      string
		dockerfile string
	}{
		{
			yaml:       "{ image: bar, dockerfile: 123 , dockerfile: 321 }",
			image:      "bar",
			dockerfile: "1231",
		},
	}
	for _, test := range tests {
		in := []byte(test.yaml)
		out := new(Build)
		err := yaml.Unmarshal(in, out)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := out.Image, test.image; got != want {
			t.Errorf("Want image %q, got %q", want, got)
		}
		if got, want := out.Dockerfile, test.dockerfile; got != want {
			t.Errorf("Want dockerfile %q, got %q", want, got)
		}
	}
}

func TestBuildError(t *testing.T) {
	in := []byte("[a]")
	out := new(Build)
	err := yaml.Unmarshal(in, out)
	if err == nil {
		t.Errorf("Expect unmarshal error")
	} else {
		t.Log("suc")
	}
}
