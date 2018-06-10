# Go project skeleton :skull:

Go project skeleton template with [Butler](https://github.com/netzkern/butler).

## Installation

See [here](https://github.com/netzkern/butler#usage)

## Usage

#### 1. Configure:

Create a `butler.yml` file in your working dir. See more options [here](https://github.com/netzkern/butler/blob/master/docs/config.md#config-places)

Example:
```yaml
templates:
  - name: skeleton
    url: https://github.com/THE108/go-skeleton.git

variables:
  repoPath: <your repo path>
  useKafka: true
  useCassandra: true
  usePostgres: true
```

`repoPath` - path for the repo you want to create (i.e `$GOPATH/src/your-org`)

`use*` - specify `true` if the needed component must be generated

Now supported:
 - `kafka`
 - `cassandra`
 - `postgresql`

#### 2. Run `butler`:
```bash
$ butler
```

#### 3. Go through the wizard's steps:
 - Create a new project
 - Specify a project name
 - Provide a description
 - Set project destination to your project path (must be `repoPath/your-project-name`, i.e. `$GOPATH/src/your-org/your-project`)
 - Press `y`
 - Enjoy!

## LICENCE

MIT
