package cli

import (
	"encoding/json"
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// NewCliClient
func cliAgent(c *cli.Context) *gocd.Client {
	var cfg *gocd.Configuration
	var err error
	if cfg, err = gocd.LoadConfig(); err != nil {
		panic(err)
	}

	if server := c.String("server"); server != "" {
		cfg.Server = server
	}

	if username := c.String("username"); username != "" {
		cfg.Username = username
	}

	if password := c.String("password"); password != "" {
		cfg.Password = password
	}

	cfg.SslCheck = cfg.SslCheck || c.Bool("ssl_check")

	return cfg.Client()
}

func handeErrOutput(reqType string, err error) error {
	return handleOutput(nil, nil, reqType, err)
}

func handleOutput(r interface{}, hr *gocd.APIResponse, reqType string, err error) error {
	var b []byte
	var o map[string]interface{}
	if err != nil {
		o = map[string]interface{}{
			"Error": err.Error(),
		}
	} else if hr.HTTP.StatusCode >= 200 && hr.HTTP.StatusCode < 300 {
		o = map[string]interface{}{
			fmt.Sprintf("%sResponse", reqType): r,
		}
		//} else if hr.HTTP.StatusCode == 404 {
		//	o = map[string]interface{}{
		//		"Error": fmt.Sprintf("Could not find resource for '%s' action.", reqType),
		//	}
	} else {

		b1, _ := json.Marshal(hr.HTTP.Header)
		b2, _ := json.Marshal(hr.Request.HTTP.Header)
		o = map[string]interface{}{
			"Error":           "An error occurred while retrieving the resource.",
			"Status":          hr.HTTP.StatusCode,
			"ResponseHeader":  string(b1),
			"ResponseBody":    hr.Body,
			"RequestBody":     hr.Request.Body,
			"RequestEndpoint": hr.Request.HTTP.URL.String(),
			"RequestHeader":   string(b2),
		}
	}
	b, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	return nil
}
