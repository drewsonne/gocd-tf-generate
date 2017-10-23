package cli

import (
	"testing"
	//ct "github.com/drewsonne/go-gocd/testcli"
	//"flag"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/stretchr/testify/assert"
	//"github.com/urfave/cli"
	"flag"
	"github.com/urfave/cli"
)

//func TestConfigure(t *testing.T) {
//	ct.Test(t, ct.TestCase{
//		Steps: []ct.Step{
//			{
//				Command: "configure",
//			},
//		},
//	})
//}

func TestConfigure(t *testing.T) {
	pfs := flag.NewFlagSet("test", flag.ContinueOnError)
	pfs.String("profile", "", "")
	pfs.Set("profile", "test-profile")
	pctx := cli.NewContext(nil, pfs, nil)

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	ctx := cli.NewContext(nil, fs, pctx)

	r := configActionRunner{
		context: ctx,
		generateConfig: func() (cfg *gocd.Configuration, err error) {
			return &gocd.Configuration{
				Server: "other-test",
			}, nil
		},
		loadConfigs: func() (cfgs map[string]*gocd.Configuration, err error) {
			return map[string]*gocd.Configuration{
				"test-profile": {
					Server: "my-test",
				},
			}, nil
		},
		writeConfigs: func(b []byte) (err error) {
			assert.Equal(t,
				`test-profile:
  server: other-test
`, string(b))
			return nil
		},
	}.run()

	assert.Nil(t, r)
}
