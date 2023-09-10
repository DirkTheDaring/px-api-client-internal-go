/*
ProxMox VE API

Testing NodesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pxapiobject

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/DirkTheDaring/px-api-client-internal-go"
)

func Test_pxapiobject_NodesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NodesApiService CreateContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.CreateContainer(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateContainerSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.CreateContainerSnapshot(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateNodesSingleStorageSingleContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var storage string

		resp, httpRes, err := apiClient.NodesApi.CreateNodesSingleStorageSingleContent(context.Background(), node, storage).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.CreateVM(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService CreateVMSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.CreateVMSnapshot(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.DeleteContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteContainerSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.DeleteContainerSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.DeleteVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService DeleteVMSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.DeleteVMSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetContainerConfig(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerConfigPending", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetContainerConfigPending(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.GetContainerSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerSnapshotConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.GetContainerSnapshotConfig(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetContainerSnapshots(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetContainerStatus(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetContainers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.GetContainers(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetCurrentContainerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetCurrentContainerStatus(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetCurrentVMStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetCurrentVMStatus(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNodeTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var upid string

		resp, httpRes, err := apiClient.NodesApi.GetNodeTask(context.Background(), node, upid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNodeTaskLog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var upid string

		resp, httpRes, err := apiClient.NodesApi.GetNodeTaskLog(context.Background(), node, upid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNodeTaskStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var upid string

		resp, httpRes, err := apiClient.NodesApi.GetNodeTaskStatus(context.Background(), node, upid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetNodeTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.GetNodeTasks(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetStorageContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var storage string

		resp, httpRes, err := apiClient.NodesApi.GetStorageContent(context.Background(), node, storage).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetStorages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.GetStorages(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetVMConfig(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMConfigPending", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetVMConfigPending(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.GetVMSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMSnapshotConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.GetVMSnapshotConfig(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.GetVMSnapshots(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService GetVMs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string

		resp, httpRes, err := apiClient.NodesApi.GetVMs(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService RebootContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.RebootContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService RebootVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.RebootVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ResizeContainerDisk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ResizeContainerDisk(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ResizeVMDisk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ResizeVMDisk(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ResumeContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ResumeContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ResumeVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ResumeVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService RollbackContainerSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.RollbackContainerSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService RollbackVMSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.RollbackVMSnapshot(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ShutdownContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ShutdownContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService ShutdownVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.ShutdownVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService StartContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.StartContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService StartVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.StartVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService StopContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.StopContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService StopNodeTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var upid string

		resp, httpRes, err := apiClient.NodesApi.StopNodeTask(context.Background(), node, upid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService StopVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.StopVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService SuspendContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.SuspendContainer(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService SuspendVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.SuspendVM(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateContainerConfigSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.UpdateContainerConfigSync(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateContainerSnapshotConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.UpdateContainerSnapshotConfig(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateVMConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.UpdateVMConfig(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateVMConfigSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64

		resp, httpRes, err := apiClient.NodesApi.UpdateVMConfigSync(context.Background(), node, vmid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UpdateVMSnapshotConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var vmid int64
		var snapname string

		resp, httpRes, err := apiClient.NodesApi.UpdateVMSnapshotConfig(context.Background(), node, vmid, snapname).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodesApiService UploadFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var node string
		var storage string

		resp, httpRes, err := apiClient.NodesApi.UploadFile(context.Background(), node, storage).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}