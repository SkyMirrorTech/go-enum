// Code generated by go-enum DO NOT EDIT.
// Version: example
// Revision: example
// Build Date: example
// Built By: example

package example

import (
	"fmt"
)

const (
	// OceanColorCerulean is a OceanColor of type Cerulean.
	OceanColorCerulean OceanColor = iota
	// OceanColorBlue is a OceanColor of type Blue.
	OceanColorBlue
	// OceanColorGreen is a OceanColor of type Green.
	OceanColorGreen
)

const _OceanColorName = "CeruleanBlueGreen"

var _OceanColorMap = map[OceanColor]string{
	OceanColorCerulean: _OceanColorName[0:8],
	OceanColorBlue:     _OceanColorName[8:12],
	OceanColorGreen:    _OceanColorName[12:17],
}

// String implements the Stringer interface.
func (x OceanColor) String() string {
	if str, ok := _OceanColorMap[x]; ok {
		return str
	}
	return fmt.Sprintf("OceanColor(%d)", x)
}

var _OceanColorValue = map[string]OceanColor{
	_OceanColorName[0:8]:   OceanColorCerulean,
	_OceanColorName[8:12]:  OceanColorBlue,
	_OceanColorName[12:17]: OceanColorGreen,
}

// ParseOceanColor attempts to convert a string to a OceanColor.
func ParseOceanColor(name string) (OceanColor, error) {
	if x, ok := _OceanColorValue[name]; ok {
		return x, nil
	}
	return OceanColor(0), fmt.Errorf("%s is not a valid OceanColor", name)
}

func ParseOceanColorGlobbedExample() bool {
	return true
}
func ParseOceanColorGlobbedExample2() bool {
	return true
}

// Additional template
func ParseOceanColorExample() bool {
	return true
}
