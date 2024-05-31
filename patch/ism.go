package patch

import (
	"github.com/disaster37/opensearch/v2"
	json "github.com/json-iterator/go"
)

func CleanIsmTemplate(actualByte []byte, expectedByte []byte) ([]byte, []byte, error) {
	actual := &opensearch.IsmPutPolicy{}
	expected := &opensearch.IsmPutPolicy{}
	var err error

	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal(actualByte, actual); err != nil {
		return nil, nil, err
	}

	if err = json.ConfigCompatibleWithStandardLibrary.Unmarshal(expectedByte, expected); err != nil {
		return nil, nil, err
	}

	// Remove lastUpdatedTime from each ismTemplate
	for i, ismTemplate := range actual.Policy.IsmTemplate {
		ismTemplate.LastUpdatedTime = nil
		actual.Policy.IsmTemplate[i] = ismTemplate
	}

	// Inject default value for retry action if not defined
	for i, state := range expected.Policy.States {
		for j, actions := range state.Actions {
			var (
				retryAction      map[string]any
				rolloverAction   map[string]any
				shrinkAction     map[string]any
				allocationAction map[string]any
			)
			for actionName, action := range actions {
				switch actionName {
				case "retry":
					retryAction = action.(map[string]any)
				case "rollover":
					rolloverAction = action.(map[string]any)
				case "shrink":
					shrinkAction = action.(map[string]any)
				case "allocation":
					allocationAction = action.(map[string]any)
				}
			}
			if retryAction == nil {
				actions["retry"] = map[string]any{
					"count":   3,
					"backoff": "exponential",
					"delay":   "1m",
				}
			} else {
				if _, ok := retryAction["count"]; !ok {
					retryAction["count"] = 3
				}
				if _, ok := retryAction["backoff"]; !ok {
					retryAction["backoff"] = "exponential"
				}
				if _, ok := retryAction["delay"]; !ok {
					retryAction["delay"] = "1m"
				}
				actions["retry"] = retryAction

			}

			if rolloverAction != nil {
				if _, ok := rolloverAction["copy_alias"]; !ok {
					rolloverAction["copy_alias"] = false
				}
				actions["rollover"] = rolloverAction

			}

			if shrinkAction != nil {
				if _, ok := shrinkAction["target_index_name_template"]; ok {
					targetIndexNameTemplate := shrinkAction["target_index_name_template"].(map[string]any)
					if _, ok := targetIndexNameTemplate["lang"]; !ok {
						targetIndexNameTemplate["lang"] = "mustache"
						shrinkAction["target_index_name_template"] = targetIndexNameTemplate
						actions["shrink"] = shrinkAction
					}
				}
			}

			if allocationAction != nil {
				if _, ok := allocationAction["require"]; !ok {
					allocationAction["require"] = map[string]any{}
				}
				if _, ok := allocationAction["include"]; !ok {
					allocationAction["include"] = map[string]any{}
				}
				if _, ok := allocationAction["exclude"]; !ok {
					allocationAction["exclude"] = map[string]any{}
				}
				if _, ok := allocationAction["wait_for"]; !ok {
					allocationAction["wait_for"] = false
				}
				actions["allocation"] = allocationAction
			}

			state.Actions[j] = actions
			expected.Policy.States[i] = state
		}
	}

	actualByte, err = json.ConfigCompatibleWithStandardLibrary.Marshal(actual)
	if err != nil {
		return nil, nil, err
	}

	expectedByte, err = json.ConfigCompatibleWithStandardLibrary.Marshal(expected)
	if err != nil {
		return nil, nil, err
	}

	return actualByte, expectedByte, nil
}
