package decode

import "core"

//decoder is responsible for implementing of decode method which returns core.Form object for further usage.
type decoder interface {
	Decode() *core.Form
}