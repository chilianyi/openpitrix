// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package gerr

import "fmt"

type ErrorMessage struct {
	Name string
	En   string
}

func (em ErrorMessage) Message(locale string, a ...interface{}) string {
	format := ""
	switch locale {
	case EN:
		format = em.En
	}
	return fmt.Sprintf(format, a...)
}

var (
	ErrorCreateResourcesFailed = ErrorMessage{
		Name: "create_resources_failed",
		En:   "create resources failed",
	}
	ErrorDeleteResourcesFailed = ErrorMessage{
		Name: "delete_resources_failed",
		En:   "delete resources failed",
	}
	ErrorDeleteResourceFailed = ErrorMessage{
		Name: "delete_resource_failed",
		En:   "delete resource [%s] failed",
	}
	ErrorUpgradeResourceFailed = ErrorMessage{
		Name: "upgrade_resource_failed",
		En:   "upgrade resource [%s] failed",
	}
	ErrorRollbackResourceFailed = ErrorMessage{
		Name: "rollback_resource_failed",
		En:   "rollback resource [%s] failed",
	}
	ErrorResizeResourceFailed = ErrorMessage{
		Name: "resize_resource_failed",
		En:   "resize resource [%s] failed",
	}
	ErrorAddResourceNodeFailed = ErrorMessage{
		Name: "add_resource_node_failed",
		En:   "add resource [%s] node failed",
	}
	ErrorDeleteResourceNodeFailed = ErrorMessage{
		Name: "delete_resource_node_failed",
		En:   "delete resource [%s] node failed",
	}
	ErrorUpdateResourceEnvFailed = ErrorMessage{
		Name: "update_resource_env_failed",
		En:   "update resource [%s] env failed",
	}
	ErrorStopResourceFailed = ErrorMessage{
		Name: "stop_resource_failed",
		En:   "stop resource [%s] failed",
	}
	ErrorStartResourceFailed = ErrorMessage{
		Name: "start_resource_failed",
		En:   "start resource [%s] failed",
	}
	ErrorRecoverResourceFailed = ErrorMessage{
		Name: "recover_resource_failed",
		En:   "recover resource [%s] failed",
	}
	ErrorCeaseResourceFailed = ErrorMessage{
		Name: "cease_resource_failed",
		En:   "cease resource [%s] failed",
	}
	ErrorRetryTaskFailed = ErrorMessage{
		Name: "retry_task_failed",
		En:   "retry task [%s] failed",
	}
	ErrorDescribeResourcesFailed = ErrorMessage{
		Name: "describe_resources_failed",
		En:   "describe resources failed",
	}
	ErrorModifyResourcesFailed = ErrorMessage{
		Name: "modify_resources_failed",
		En:   "modify resources failed",
	}
	ErrorModifyResourceFailed = ErrorMessage{
		Name: "modify_resource_failed",
		En:   "modify resource [%s] failed",
	}
	ErrorResourceNotFound = ErrorMessage{
		Name: "resource_not_found",
		En:   "resource [%s] not found",
	}
	ErrorInternalError = ErrorMessage{
		Name: "internal_error",
		En:   "internal error",
	}
	ErrorMissingParameter = ErrorMessage{
		Name: "missing_parameter",
		En:   "missing parameter [%s]",
	}
	ErrorValidateFailed = ErrorMessage{
		Name: "validate_failed",
		En:   "validate failed",
	}
	ErrorParameterParseFailed = ErrorMessage{
		Name: "parameter_parse_failed",
		En:   "parameter [%s] parse failed",
	}
	ErrorResourceAlreadyDeleted = ErrorMessage{
		Name: "resource_already_deleted",
		En:   "resource [%s] has already been deleted",
	}
	ErrorIllegalParameterLength = ErrorMessage{
		Name: "illegal_parameter_length",
		En:   "illegal parameter [%s] length",
	}
	ErrorUnsupportedParameterValue = ErrorMessage{
		Name: "unsupported_parameter_value",
		En:   "unsupported parameter [%s] value [%s]",
	}
	ErrorIllegalUrlFormat = ErrorMessage{
		Name: "illegal_url_format",
		En:   "illegal URL format [%s]",
	}
	ErrorIllegalLabelFormat = ErrorMessage{
		Name: "illegal_label_format",
		En:   "illegal label format",
	}
)