package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


//Containing struct for the templates being generated
type Template struct {
	APIVersion string           `json:"apiVersion,omitempty"`
	Properties labs.Properties  `json:"properties"`
}
