package secret

import "errors"

var ErrUnsupportedSecretType = errors.New("unsupported secret type")

var ErrAccessProhibited = errors.New("secret does not belong to user")

var ErrRecordNotFound = errors.New("response provides no content")
