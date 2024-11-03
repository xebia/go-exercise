package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	dbMock "github.com/xebia/go-exercise/internal/db/testdata"
	emailMock "github.com/xebia/go-exercise/internal/email/testdata"
)

type ServerSuite struct {
	//nolint:unused
	dbMock *dbMock.Mock
	//nolint:unused
	emailMock *emailMock.Mock
	suite.Suite
	server *Server
}

func (ss *ServerSuite) SetupSuite() {
	addr := "localhost:10000"
	dm := &dbMock.Mock{}
	em := &emailMock.Mock{}
	ss.server = NewServer(dm, em)

	go func() {
		ss.server.ListenAndServe(addr)
	}()
	time.Sleep(10 * time.Millisecond)
}

func (ss *ServerSuite) TearDownSuite() {
	ss.server.Shutdown()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
