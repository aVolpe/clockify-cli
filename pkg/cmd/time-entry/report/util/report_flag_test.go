package util_test

import (
	"testing"

	"github.com/lucassabreu/clockify-cli/pkg/cmd/time-entry/report/util"
	"github.com/stretchr/testify/assert"
)

func TestReportBillableFlagsChecks(t *testing.T) {
	rf := util.NewReportFlags()
	rf.Billable = true
	rf.NotBillable = true

	err := rf.Check()
	if assert.Error(t, err) {
		assert.Regexp(t,
			"can't be used together.*billable.*not-billable", err.Error())
	}

	rf.Billable = false
	rf.NotBillable = true

	assert.NoError(t, rf.Check())

	rf.Billable = true
	rf.NotBillable = false

	assert.NoError(t, rf.Check())
}

func TestReportProjectFlagsChecks(t *testing.T) {
	rf := util.NewReportFlags()
	rf.Client = "me"
	rf.Project = ""

	err := rf.Check()
	if assert.Error(t, err) {
		assert.Equal(t,
			"flag 'client' can't be used without flag 'project'", err.Error())
	}

	rf.Client = ""
	rf.Project = "mine"

	assert.NoError(t, rf.Check())

	rf.Client = "me"
	rf.Project = "mine"

	assert.NoError(t, rf.Check())
}
