// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package bolttest

import (
	"database/sql/driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"sync"
	"time"
)

var (
	lockConnMockBegin           sync.RWMutex
	lockConnMockClose           sync.RWMutex
	lockConnMockExecNeo         sync.RWMutex
	lockConnMockExecPipeline    sync.RWMutex
	lockConnMockPrepareNeo      sync.RWMutex
	lockConnMockPreparePipeline sync.RWMutex
	lockConnMockQueryNeo        sync.RWMutex
	lockConnMockQueryNeoAll     sync.RWMutex
	lockConnMockQueryPipeline   sync.RWMutex
	lockConnMockSetChunkSize    sync.RWMutex
	lockConnMockSetTimeout      sync.RWMutex
)

// ConnMock is a mock implementation of Conn.
//
//     func TestSomethingThatUsesConn(t *testing.T) {
//
//         // make and configure a mocked Conn
//         mockedConn := &ConnMock{
//             BeginFunc: func() (driver.Tx, error) {
// 	               panic("TODO: mock out the Begin method")
//             },
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             ExecNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
// 	               panic("TODO: mock out the ExecNeo method")
//             },
//             ExecPipelineFunc: func(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error) {
// 	               panic("TODO: mock out the ExecPipeline method")
//             },
//             PrepareNeoFunc: func(query string) (golangNeo4jBoltDriver.Stmt, error) {
// 	               panic("TODO: mock out the PrepareNeo method")
//             },
//             PreparePipelineFunc: func(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error) {
// 	               panic("TODO: mock out the PreparePipeline method")
//             },
//             QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
// 	               panic("TODO: mock out the QueryNeo method")
//             },
//             QueryNeoAllFunc: func(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error) {
// 	               panic("TODO: mock out the QueryNeoAll method")
//             },
//             QueryPipelineFunc: func(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error) {
// 	               panic("TODO: mock out the QueryPipeline method")
//             },
//             SetChunkSizeFunc: func(in1 uint16)  {
// 	               panic("TODO: mock out the SetChunkSize method")
//             },
//             SetTimeoutFunc: func(in1 time.Duration)  {
// 	               panic("TODO: mock out the SetTimeout method")
//             },
//         }
//
//         // TODO: use mockedConn in code that requires Conn
//         //       and then make assertions.
//
//     }
type ConnMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (driver.Tx, error)

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ExecNeoFunc mocks the ExecNeo method.
	ExecNeoFunc func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error)

	// ExecPipelineFunc mocks the ExecPipeline method.
	ExecPipelineFunc func(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error)

	// PrepareNeoFunc mocks the PrepareNeo method.
	PrepareNeoFunc func(query string) (golangNeo4jBoltDriver.Stmt, error)

	// PreparePipelineFunc mocks the PreparePipeline method.
	PreparePipelineFunc func(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error)

	// QueryNeoFunc mocks the QueryNeo method.
	QueryNeoFunc func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error)

	// QueryNeoAllFunc mocks the QueryNeoAll method.
	QueryNeoAllFunc func(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error)

	// QueryPipelineFunc mocks the QueryPipeline method.
	QueryPipelineFunc func(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error)

	// SetChunkSizeFunc mocks the SetChunkSize method.
	SetChunkSizeFunc func(in1 uint16)

	// SetTimeoutFunc mocks the SetTimeout method.
	SetTimeoutFunc func(in1 time.Duration)

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// ExecNeo holds details about calls to the ExecNeo method.
		ExecNeo []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// ExecPipeline holds details about calls to the ExecPipeline method.
		ExecPipeline []struct {
			// Query is the query argument value.
			Query []string
			// Params is the params argument value.
			Params []map[string]interface{}
		}
		// PrepareNeo holds details about calls to the PrepareNeo method.
		PrepareNeo []struct {
			// Query is the query argument value.
			Query string
		}
		// PreparePipeline holds details about calls to the PreparePipeline method.
		PreparePipeline []struct {
			// Query is the query argument value.
			Query []string
		}
		// QueryNeo holds details about calls to the QueryNeo method.
		QueryNeo []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// QueryNeoAll holds details about calls to the QueryNeoAll method.
		QueryNeoAll []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// QueryPipeline holds details about calls to the QueryPipeline method.
		QueryPipeline []struct {
			// Query is the query argument value.
			Query []string
			// Params is the params argument value.
			Params []map[string]interface{}
		}
		// SetChunkSize holds details about calls to the SetChunkSize method.
		SetChunkSize []struct {
			// In1 is the in1 argument value.
			In1 uint16
		}
		// SetTimeout holds details about calls to the SetTimeout method.
		SetTimeout []struct {
			// In1 is the in1 argument value.
			In1 time.Duration
		}
	}
}

// Begin calls BeginFunc.
func (mock *ConnMock) Begin() (driver.Tx, error) {
	if mock.BeginFunc == nil {
		panic("moq: ConnMock.BeginFunc is nil but Conn.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockConnMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockConnMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedConn.BeginCalls())
func (mock *ConnMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockConnMockBegin.RLock()
	calls = mock.calls.Begin
	lockConnMockBegin.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *ConnMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: ConnMock.CloseFunc is nil but Conn.Close was just called")
	}
	callInfo := struct {
	}{}
	lockConnMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockConnMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedConn.CloseCalls())
func (mock *ConnMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockConnMockClose.RLock()
	calls = mock.calls.Close
	lockConnMockClose.RUnlock()
	return calls
}

// ExecNeo calls ExecNeoFunc.
func (mock *ConnMock) ExecNeo(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
	if mock.ExecNeoFunc == nil {
		panic("moq: ConnMock.ExecNeoFunc is nil but Conn.ExecNeo was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockConnMockExecNeo.Lock()
	mock.calls.ExecNeo = append(mock.calls.ExecNeo, callInfo)
	lockConnMockExecNeo.Unlock()
	return mock.ExecNeoFunc(query, params)
}

// ExecNeoCalls gets all the calls that were made to ExecNeo.
// Check the length with:
//     len(mockedConn.ExecNeoCalls())
func (mock *ConnMock) ExecNeoCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockConnMockExecNeo.RLock()
	calls = mock.calls.ExecNeo
	lockConnMockExecNeo.RUnlock()
	return calls
}

// ExecPipeline calls ExecPipelineFunc.
func (mock *ConnMock) ExecPipeline(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error) {
	if mock.ExecPipelineFunc == nil {
		panic("moq: ConnMock.ExecPipelineFunc is nil but Conn.ExecPipeline was just called")
	}
	callInfo := struct {
		Query  []string
		Params []map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockConnMockExecPipeline.Lock()
	mock.calls.ExecPipeline = append(mock.calls.ExecPipeline, callInfo)
	lockConnMockExecPipeline.Unlock()
	return mock.ExecPipelineFunc(query, params...)
}

// ExecPipelineCalls gets all the calls that were made to ExecPipeline.
// Check the length with:
//     len(mockedConn.ExecPipelineCalls())
func (mock *ConnMock) ExecPipelineCalls() []struct {
	Query  []string
	Params []map[string]interface{}
} {
	var calls []struct {
		Query  []string
		Params []map[string]interface{}
	}
	lockConnMockExecPipeline.RLock()
	calls = mock.calls.ExecPipeline
	lockConnMockExecPipeline.RUnlock()
	return calls
}

// PrepareNeo calls PrepareNeoFunc.
func (mock *ConnMock) PrepareNeo(query string) (golangNeo4jBoltDriver.Stmt, error) {
	if mock.PrepareNeoFunc == nil {
		panic("moq: ConnMock.PrepareNeoFunc is nil but Conn.PrepareNeo was just called")
	}
	callInfo := struct {
		Query string
	}{
		Query: query,
	}
	lockConnMockPrepareNeo.Lock()
	mock.calls.PrepareNeo = append(mock.calls.PrepareNeo, callInfo)
	lockConnMockPrepareNeo.Unlock()
	return mock.PrepareNeoFunc(query)
}

// PrepareNeoCalls gets all the calls that were made to PrepareNeo.
// Check the length with:
//     len(mockedConn.PrepareNeoCalls())
func (mock *ConnMock) PrepareNeoCalls() []struct {
	Query string
} {
	var calls []struct {
		Query string
	}
	lockConnMockPrepareNeo.RLock()
	calls = mock.calls.PrepareNeo
	lockConnMockPrepareNeo.RUnlock()
	return calls
}

// PreparePipeline calls PreparePipelineFunc.
func (mock *ConnMock) PreparePipeline(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error) {
	if mock.PreparePipelineFunc == nil {
		panic("moq: ConnMock.PreparePipelineFunc is nil but Conn.PreparePipeline was just called")
	}
	callInfo := struct {
		Query []string
	}{
		Query: query,
	}
	lockConnMockPreparePipeline.Lock()
	mock.calls.PreparePipeline = append(mock.calls.PreparePipeline, callInfo)
	lockConnMockPreparePipeline.Unlock()
	return mock.PreparePipelineFunc(query...)
}

// PreparePipelineCalls gets all the calls that were made to PreparePipeline.
// Check the length with:
//     len(mockedConn.PreparePipelineCalls())
func (mock *ConnMock) PreparePipelineCalls() []struct {
	Query []string
} {
	var calls []struct {
		Query []string
	}
	lockConnMockPreparePipeline.RLock()
	calls = mock.calls.PreparePipeline
	lockConnMockPreparePipeline.RUnlock()
	return calls
}

// QueryNeo calls QueryNeoFunc.
func (mock *ConnMock) QueryNeo(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
	if mock.QueryNeoFunc == nil {
		panic("moq: ConnMock.QueryNeoFunc is nil but Conn.QueryNeo was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockConnMockQueryNeo.Lock()
	mock.calls.QueryNeo = append(mock.calls.QueryNeo, callInfo)
	lockConnMockQueryNeo.Unlock()
	return mock.QueryNeoFunc(query, params)
}

// QueryNeoCalls gets all the calls that were made to QueryNeo.
// Check the length with:
//     len(mockedConn.QueryNeoCalls())
func (mock *ConnMock) QueryNeoCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockConnMockQueryNeo.RLock()
	calls = mock.calls.QueryNeo
	lockConnMockQueryNeo.RUnlock()
	return calls
}

// QueryNeoAll calls QueryNeoAllFunc.
func (mock *ConnMock) QueryNeoAll(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error) {
	if mock.QueryNeoAllFunc == nil {
		panic("moq: ConnMock.QueryNeoAllFunc is nil but Conn.QueryNeoAll was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockConnMockQueryNeoAll.Lock()
	mock.calls.QueryNeoAll = append(mock.calls.QueryNeoAll, callInfo)
	lockConnMockQueryNeoAll.Unlock()
	return mock.QueryNeoAllFunc(query, params)
}

// QueryNeoAllCalls gets all the calls that were made to QueryNeoAll.
// Check the length with:
//     len(mockedConn.QueryNeoAllCalls())
func (mock *ConnMock) QueryNeoAllCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockConnMockQueryNeoAll.RLock()
	calls = mock.calls.QueryNeoAll
	lockConnMockQueryNeoAll.RUnlock()
	return calls
}

// QueryPipeline calls QueryPipelineFunc.
func (mock *ConnMock) QueryPipeline(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error) {
	if mock.QueryPipelineFunc == nil {
		panic("moq: ConnMock.QueryPipelineFunc is nil but Conn.QueryPipeline was just called")
	}
	callInfo := struct {
		Query  []string
		Params []map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockConnMockQueryPipeline.Lock()
	mock.calls.QueryPipeline = append(mock.calls.QueryPipeline, callInfo)
	lockConnMockQueryPipeline.Unlock()
	return mock.QueryPipelineFunc(query, params...)
}

// QueryPipelineCalls gets all the calls that were made to QueryPipeline.
// Check the length with:
//     len(mockedConn.QueryPipelineCalls())
func (mock *ConnMock) QueryPipelineCalls() []struct {
	Query  []string
	Params []map[string]interface{}
} {
	var calls []struct {
		Query  []string
		Params []map[string]interface{}
	}
	lockConnMockQueryPipeline.RLock()
	calls = mock.calls.QueryPipeline
	lockConnMockQueryPipeline.RUnlock()
	return calls
}

// SetChunkSize calls SetChunkSizeFunc.
func (mock *ConnMock) SetChunkSize(in1 uint16) {
	if mock.SetChunkSizeFunc == nil {
		panic("moq: ConnMock.SetChunkSizeFunc is nil but Conn.SetChunkSize was just called")
	}
	callInfo := struct {
		In1 uint16
	}{
		In1: in1,
	}
	lockConnMockSetChunkSize.Lock()
	mock.calls.SetChunkSize = append(mock.calls.SetChunkSize, callInfo)
	lockConnMockSetChunkSize.Unlock()
	mock.SetChunkSizeFunc(in1)
}

// SetChunkSizeCalls gets all the calls that were made to SetChunkSize.
// Check the length with:
//     len(mockedConn.SetChunkSizeCalls())
func (mock *ConnMock) SetChunkSizeCalls() []struct {
	In1 uint16
} {
	var calls []struct {
		In1 uint16
	}
	lockConnMockSetChunkSize.RLock()
	calls = mock.calls.SetChunkSize
	lockConnMockSetChunkSize.RUnlock()
	return calls
}

// SetTimeout calls SetTimeoutFunc.
func (mock *ConnMock) SetTimeout(in1 time.Duration) {
	if mock.SetTimeoutFunc == nil {
		panic("moq: ConnMock.SetTimeoutFunc is nil but Conn.SetTimeout was just called")
	}
	callInfo := struct {
		In1 time.Duration
	}{
		In1: in1,
	}
	lockConnMockSetTimeout.Lock()
	mock.calls.SetTimeout = append(mock.calls.SetTimeout, callInfo)
	lockConnMockSetTimeout.Unlock()
	mock.SetTimeoutFunc(in1)
}

// SetTimeoutCalls gets all the calls that were made to SetTimeout.
// Check the length with:
//     len(mockedConn.SetTimeoutCalls())
func (mock *ConnMock) SetTimeoutCalls() []struct {
	In1 time.Duration
} {
	var calls []struct {
		In1 time.Duration
	}
	lockConnMockSetTimeout.RLock()
	calls = mock.calls.SetTimeout
	lockConnMockSetTimeout.RUnlock()
	return calls
}
