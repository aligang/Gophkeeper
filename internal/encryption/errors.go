package encryption

import "errors"

var ErrCipherCreation = errors.New("error during cipher creating")
var ErrDataDecoding = errors.New("error decoding base64 encoded data")
var ErrFailedDecryption = errors.New("could not decrypt data")
