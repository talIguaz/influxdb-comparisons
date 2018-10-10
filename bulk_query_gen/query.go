package bulk_query_gen

import (
	"fmt"
	"sync"
	"time"
)

const DefaultQueryInterval = time.Hour

type Query interface {
	Release()
	HumanLabelName() []byte
	HumanDescriptionName() []byte
	fmt.Stringer
}

var HTTPQueryPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return &HTTPQuery{
			HumanLabel:       []byte{},
			HumanDescription: []byte{},
			Method:           []byte{},
			Path:             []byte{},
			Body:             []byte{},
			StartTimestamp:   0,
			EndTimestamp:     0,
		}
	},
}

// HTTPQuery encodes an HTTP request. This will typically by serialized for use
// by the query_benchmarker program.
type HTTPQuery struct {
	HumanLabel       []byte
	HumanDescription []byte
	Method           []byte
	Path             []byte
	Body             []byte
	StartTimestamp   int64
	EndTimestamp     int64
}

func NewHTTPQuery() *HTTPQuery {
	return HTTPQueryPool.Get().(*HTTPQuery)
}

// String produces a debug-ready description of a Query.
func (q *HTTPQuery) String() string {
	return fmt.Sprintf("HumanLabel: \"%s\", HumanDescription: \"%s\", Method: \"%s\", Path: \"%s\", Body: \"%s\"", q.HumanLabel, q.HumanDescription, q.Method, q.Path, q.Body)
}

func (q *HTTPQuery) HumanLabelName() []byte {
	return q.HumanLabel
}
func (q *HTTPQuery) HumanDescriptionName() []byte {
	return q.HumanDescription
}

func (q *HTTPQuery) Release() {
	q.HumanLabel = q.HumanLabel[:0]
	q.HumanDescription = q.HumanDescription[:0]
	q.Method = q.Method[:0]
	q.Path = q.Path[:0]
	q.Body = q.Body[:0]
	q.StartTimestamp = 0
	q.EndTimestamp = 0

	HTTPQueryPool.Put(q)
}
