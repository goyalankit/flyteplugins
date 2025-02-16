// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"
	io "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/io"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flytestdlib/storage"
)

// PluginContext is an autogenerated mock type for the PluginContext type
type PluginContext struct {
	mock.Mock
}

type PluginContext_DataStore struct {
	*mock.Call
}

func (_m PluginContext_DataStore) Return(_a0 *storage.DataStore) *PluginContext_DataStore {
	return &PluginContext_DataStore{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnDataStore() *PluginContext_DataStore {
	c_call := _m.On("DataStore")
	return &PluginContext_DataStore{Call: c_call}
}

func (_m *PluginContext) OnDataStoreMatch(matchers ...interface{}) *PluginContext_DataStore {
	c_call := _m.On("DataStore", matchers...)
	return &PluginContext_DataStore{Call: c_call}
}

// DataStore provides a mock function with given fields:
func (_m *PluginContext) DataStore() *storage.DataStore {
	ret := _m.Called()

	var r0 *storage.DataStore
	if rf, ok := ret.Get(0).(func() *storage.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.DataStore)
		}
	}

	return r0
}

type PluginContext_InputReader struct {
	*mock.Call
}

func (_m PluginContext_InputReader) Return(_a0 io.InputReader) *PluginContext_InputReader {
	return &PluginContext_InputReader{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnInputReader() *PluginContext_InputReader {
	c_call := _m.On("InputReader")
	return &PluginContext_InputReader{Call: c_call}
}

func (_m *PluginContext) OnInputReaderMatch(matchers ...interface{}) *PluginContext_InputReader {
	c_call := _m.On("InputReader", matchers...)
	return &PluginContext_InputReader{Call: c_call}
}

// InputReader provides a mock function with given fields:
func (_m *PluginContext) InputReader() io.InputReader {
	ret := _m.Called()

	var r0 io.InputReader
	if rf, ok := ret.Get(0).(func() io.InputReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.InputReader)
		}
	}

	return r0
}

type PluginContext_MaxDatasetSizeBytes struct {
	*mock.Call
}

func (_m PluginContext_MaxDatasetSizeBytes) Return(_a0 int64) *PluginContext_MaxDatasetSizeBytes {
	return &PluginContext_MaxDatasetSizeBytes{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnMaxDatasetSizeBytes() *PluginContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes")
	return &PluginContext_MaxDatasetSizeBytes{Call: c_call}
}

func (_m *PluginContext) OnMaxDatasetSizeBytesMatch(matchers ...interface{}) *PluginContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes", matchers...)
	return &PluginContext_MaxDatasetSizeBytes{Call: c_call}
}

// MaxDatasetSizeBytes provides a mock function with given fields:
func (_m *PluginContext) MaxDatasetSizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type PluginContext_OutputWriter struct {
	*mock.Call
}

func (_m PluginContext_OutputWriter) Return(_a0 io.OutputWriter) *PluginContext_OutputWriter {
	return &PluginContext_OutputWriter{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnOutputWriter() *PluginContext_OutputWriter {
	c_call := _m.On("OutputWriter")
	return &PluginContext_OutputWriter{Call: c_call}
}

func (_m *PluginContext) OnOutputWriterMatch(matchers ...interface{}) *PluginContext_OutputWriter {
	c_call := _m.On("OutputWriter", matchers...)
	return &PluginContext_OutputWriter{Call: c_call}
}

// OutputWriter provides a mock function with given fields:
func (_m *PluginContext) OutputWriter() io.OutputWriter {
	ret := _m.Called()

	var r0 io.OutputWriter
	if rf, ok := ret.Get(0).(func() io.OutputWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.OutputWriter)
		}
	}

	return r0
}

type PluginContext_TaskExecutionMetadata struct {
	*mock.Call
}

func (_m PluginContext_TaskExecutionMetadata) Return(_a0 core.TaskExecutionMetadata) *PluginContext_TaskExecutionMetadata {
	return &PluginContext_TaskExecutionMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnTaskExecutionMetadata() *PluginContext_TaskExecutionMetadata {
	c_call := _m.On("TaskExecutionMetadata")
	return &PluginContext_TaskExecutionMetadata{Call: c_call}
}

func (_m *PluginContext) OnTaskExecutionMetadataMatch(matchers ...interface{}) *PluginContext_TaskExecutionMetadata {
	c_call := _m.On("TaskExecutionMetadata", matchers...)
	return &PluginContext_TaskExecutionMetadata{Call: c_call}
}

// TaskExecutionMetadata provides a mock function with given fields:
func (_m *PluginContext) TaskExecutionMetadata() core.TaskExecutionMetadata {
	ret := _m.Called()

	var r0 core.TaskExecutionMetadata
	if rf, ok := ret.Get(0).(func() core.TaskExecutionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskExecutionMetadata)
		}
	}

	return r0
}

type PluginContext_TaskReader struct {
	*mock.Call
}

func (_m PluginContext_TaskReader) Return(_a0 core.TaskReader) *PluginContext_TaskReader {
	return &PluginContext_TaskReader{Call: _m.Call.Return(_a0)}
}

func (_m *PluginContext) OnTaskReader() *PluginContext_TaskReader {
	c_call := _m.On("TaskReader")
	return &PluginContext_TaskReader{Call: c_call}
}

func (_m *PluginContext) OnTaskReaderMatch(matchers ...interface{}) *PluginContext_TaskReader {
	c_call := _m.On("TaskReader", matchers...)
	return &PluginContext_TaskReader{Call: c_call}
}

// TaskReader provides a mock function with given fields:
func (_m *PluginContext) TaskReader() core.TaskReader {
	ret := _m.Called()

	var r0 core.TaskReader
	if rf, ok := ret.Get(0).(func() core.TaskReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskReader)
		}
	}

	return r0
}
