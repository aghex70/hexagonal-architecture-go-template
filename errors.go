package hexagonal_architecture_go_template

import "errors"

var ProvideValueError = errors.New("please provide a default value")
var ProjectFileError = errors.New("path file already existent")
