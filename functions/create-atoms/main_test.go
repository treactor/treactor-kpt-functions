package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_create.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, len(resourceList.Items), 9)
	//yaml, err := resourceList.ToYAML()
	//assert.NoError(t, err)
	//println(string(yaml))
}

func TestUpdate(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_update.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, len(resourceList.Items), 9)
}
