package opensearchhandler

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/elastic/go-ucfg"
	ucfgjson "github.com/elastic/go-ucfg/json"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// StandardDiff permit to compare objects
func StandardDiff(actual, expected any, log *logrus.Entry, ignore map[string]any) (diff string, err error) {
	acualByte, err := json.Marshal(actual)
	if err != nil {
		return diff, errors.Wrap(err, "Error when marshall actual")
	}
	expectedByte, err := json.Marshal(expected)
	if err != nil {
		return diff, errors.Wrap(err, "Error whe marshall expected")
	}

	actualConf, err := ucfgjson.NewConfig(acualByte, ucfg.PathSep("."))
	if err != nil {
		return diff, errors.Wrap(err, "Error when init config object from actual")
	}
	if err = ignoreDiff(actualConf, ignore); err != nil {
		return diff, errors.Wrap(err, "Error when ignore diff on actual")
	}
	actualUnpack := reflect.New(reflect.TypeOf(actual)).Interface()
	if err = actualConf.Unpack(actualUnpack, ucfg.StructTag("json")); err != nil {
		return diff, errors.Wrap(err, "Error when unpack on actual")
	}
	expectedConf, err := ucfgjson.NewConfig(expectedByte, ucfg.PathSep("."))
	if err != nil {
		return diff, errors.Wrap(err, "Error when init config object from expected")
	}
	if err = ignoreDiff(expectedConf, ignore); err != nil {
		return diff, errors.Wrap(err, "Error when ignore diff on expected")
	}
	expectedUnpack := reflect.New(reflect.TypeOf(expected)).Interface()
	if err = expectedConf.Unpack(expectedUnpack, ucfg.StructTag("json")); err != nil {
		return diff, errors.Wrap(err, "Error when unpack on expected")
	}

	/*
		test := map[string]any{}
		if err = expectedConf.Unpack(&test); err != nil {
			return diff, errors.Wrap(err, "Error when test")
		}
	*/

	return cmp.Diff(actualUnpack, expectedUnpack), nil
}

func ignoreDiff(c *ucfg.Config, ignore map[string]any) (err error) {
	for key, value := range ignore {
		hasField, err := c.Has(key, -1, ucfg.PathSep("."))
		if err != nil {
			return errors.Wrapf(err, "Error when check if field %s exist", key)
		}
		if hasField {
			needRemoveKey := false
			if value == nil {
				needRemoveKey = true
			} else {
				var v any
				switch t := value.(type) {
				case bool:
					v, err = c.Bool(key, -1, ucfg.PathSep("."))
					if err != nil {
						return errors.Wrapf(err, "Error when get bool value for key: %s", key)
					}
				case string:
					v, err = c.String(key, -1, ucfg.PathSep("."))
					if err != nil {
						return errors.Wrapf(err, "Error when get string value for key: %s", key)
					}
				case int64:
					v, err = c.Int(key, -1, ucfg.PathSep("."))
					if err != nil {
						return errors.Wrapf(err, "Error when get int64 value for key: %s", key)
					}
				case float64:
					v, err = c.Float(key, -1, ucfg.PathSep("."))
					if err != nil {
						return errors.Wrapf(err, "Error when get float64 value for key: %s", key)
					}
				default:
					return errors.Errorf("Type %T not supported", t)
				}

				if v == value {
					needRemoveKey = true
				}
			}
			if needRemoveKey {

				if _, err = c.Remove(key, -1, ucfg.PathSep(".")); err != nil {
					return errors.Wrapf(err, "Error when remove key %s", key)
				}

				// Check if needed to remove parent key because of is empty
				parentPaths := strings.Split(key, ".")
				if len(parentPaths) > 1 {
					parentPath := strings.Join(parentPaths[:len(parentPaths)-1], ".")
					parent, err := c.Child(parentPath, -1, ucfg.PathSep("."))
					if err != nil {
						return errors.Wrapf(err, "Error when get parent path %s for key %s", parentPath, key)
					}
					nb := len(parent.GetFields())
					// Remove parent if no children
					if nb == 0 {
						if _, err = c.Remove(parentPath, -1, ucfg.PathSep(".")); err != nil {
							return errors.Wrapf(err, "Error when when remove parent key %s", parentPath)
						}
					}
				}

			}
		}
	}

	return nil
}
