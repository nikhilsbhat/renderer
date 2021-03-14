package buildercli

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/nikhilsbhat/config/decode"
	"github.com/urfave/cli/v2"

	"regexp"
)

type config struct {
	envs         map[string]string
	templatePath string
	targetPath   string
	templateData string
}

func Render(c *cli.Context) error {
	parsedEnvs, err := getEnvs(c.StringSlice(flagEnvVars))
	if err != nil {
		return err
	}

	newConfig := &config{
		envs:         parsedEnvs,
		templatePath: c.String(flagTemplatePath),
		targetPath:   c.String(flagTargetPath),
	}

	file, err := newConfig.validateAndGetConfig()
	if err != nil {
		return err
	}

	if err := newConfig.renderGoTemplate(file); err != nil {
		return err
	}

	return nil
}

func getEnvs(unparsedEnvs []string) (map[string]string, error) {
	parsedEnvs := make(map[string]string, len(unparsedEnvs))
	for _, env := range unparsedEnvs {
		regex := regexp.MustCompile(`=`)
		value := regex.Split(env, -1)
		if len(value) > 2 || len(value) < 2 {
			return nil, fmt.Errorf(fmt.Sprintf("environment variable '%s' was not passed correctly", env))
		}
		parsedEnvs[value[0]] = value[len(value)-1]
	}
	return parsedEnvs, nil
}

func (c *config) validateAndGetConfig() (io.Writer, error) {
	if _, direrr := os.Stat(c.templatePath); os.IsNotExist(direrr) {
		return nil, direrr
	}

	configContent, err := decode.ReadFile(c.templatePath)
	if err != nil {
		return nil, err
	}
	c.templateData = string(configContent)

	target, err := filepath.Abs(c.targetPath)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	file, err := os.Create(filepath.Join(target))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *config) renderGoTemplate(file io.Writer) error {
	configTemplate := template.Must(template.New("gotemplate").Parse(c.templateData))
	if err := configTemplate.Execute(file, c.envs); err != nil {
		return err
	}
	return nil
}
