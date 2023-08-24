package projector_test

import (
	"reflect"
	"testing"

	"github.com/barlevalon/polyglot-course/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func TestConfigPrint(t *testing.T) {
  testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
  testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
  testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
  testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
  opts := getOpts(args)
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("Expected args %+v, got %+v", expectedArgs, config.Args)
	}
	if config.Operation != operation {
		t.Errorf("operation expected was %+v but got %+v", operation, config.Operation)
	}
}
