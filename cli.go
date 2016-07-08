package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"gopkg.in/yaml.v2"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeFileOpenError
	ExitCodeRequestError
)

// CLI 構造体
type CLI struct {
	outStream, errStream io.Writer
}

// Config 構造体
type Config struct {
	URL   string
	ID    string
	Token string
}

// Run 引数処理を含めた具体的な処理
func (c *CLI) Run(args []string) int {
	// オプション引数のパース
	var version bool
	var showConfig bool
	flags := flag.NewFlagSet("kintai-cli", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")
	flags.BoolVar(&showConfig, "show-config", false, "Print config information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	// バージョン情報の表示
	if version {
		fmt.Fprintf(c.outStream, "version %s\n", Version)
		return ExitCodeOK
	}

	// read config
	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Fprint(c.errStream, err)
		return ExitCodeFileOpenError
	}

	// read config
	var config Config
	yaml.Unmarshal(buf, &config)

	// show config
	if showConfig {
		fmt.Fprintf(c.outStream, "url: %s\n", config.URL)
		fmt.Fprintf(c.outStream, "user_id: %s\n", config.ID)
		fmt.Fprintf(c.outStream, "token: %s\n", config.Token)
		return ExitCodeOK
	}

	fmt.Fprint(c.outStream, "Do kintai work\n")

	values := url.Values{}
	values.Add("user_id", config.ID)
	values.Add("access_token", config.Token)

	// get
	// resp, err := http.Get(config.URL + "?" + values.Encode())

	// post
	resp, err := http.PostForm(config.URL, values)

	if err != nil {
		fmt.Fprint(c.errStream, err)
		return ExitCodeRequestError
	}

	defer resp.Body.Close()

	execute(resp, c)

	return ExitCodeOK
}

func execute(response *http.Response, c *CLI) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprint(c.errStream, err)
	}

	fmt.Fprint(c.outStream, string(body))
}
