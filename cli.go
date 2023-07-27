package mural

type CLI struct {
	name        string
	description string
	version     Version
}

func NewCLI(name, version, description string) (*CLI, error) {
	semver, err := NewVersion(version)
	if err != nil {
		return nil, err
	}
	return &CLI{
		name:        name,
		description: description,
		version:     semver,
	}, nil
}

func (cli *CLI) Version() Version {
	return cli.version
}
