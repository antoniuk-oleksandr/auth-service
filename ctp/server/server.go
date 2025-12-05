package server

import (
	"context"
	customprotoc "github.com/antoniuk-oleksandr/auth-service/ctp/internal"
	"github.com/antoniuk-oleksandr/auth-service/ctp/types"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

type server struct {
	Config   types.Config
	Handlers map[string][]types.HandlerFunc
	Reader   types.RequestReader
	Writer   types.ResponseWriter
	Logger   types.Logger
}

type ServerOption func(*server)

func NewServer(opts ...ServerOption) types.Server {
	s := &server{
		Config:   DefaultConfig,
		Handlers: make(map[string][]types.HandlerFunc),
		Reader:   customprotoc.NewRequestReader(DefaultConfig.MaxBodySize),
		Writer:   customprotoc.NewResponseWriter(),
		Logger:   &customprotoc.DefaultLogger{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *server) RegisterHandler(command string, handlers ...types.HandlerFunc) {
	s.Handlers[command] = handlers
}

func (s *server) Start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer listener.Close()

	s.Logger.Info(
		"server started",
		types.Field{Key: "port", Value: port, Type: types.StringField},
	)

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.Logger.Error(
				"accept connection",
				types.Field{Key: "port", Value: err.Error(), Type: types.ErrorField},
			)
			continue
		}
		go s.HandleConnection(conn)
	}
}

func (s *server) HandleConnection(conn io.ReadWriteCloser) {
	defer conn.Close()

	req, err := s.Reader.ReadRequest(conn)
	if err != nil {
		s.Logger.Error(
			"accept connection",
			types.Field{Key: "err", Value: err, Type: types.ErrorField},
		)
		return
	}

	resp := s.ProcessRequest(req)

	if err := s.Writer.WriteResponse(conn, resp); err != nil {
		s.Logger.Error(
			"write responsen",
			types.Field{Key: "err", Value: err.Error(), Type: types.ErrorField},
		)
	}
}

func (s *server) ProcessRequest(req *types.Request) *types.Response {
	command := strings.TrimSpace(req.Command)

	handlers, exists := s.Handlers[command]
	if !exists {
		return &types.Response{
			Status: types.StatusNotFound,
			Body:   mustMsgpack(map[string]string{"error": "command not found"}),
		}
	}

	ctx := customprotoc.NewCtx(req, context.Background(), handlers)

	if err := handlers[0](ctx); err != nil {
		if resp, ok := err.(*types.Response); ok {
			return resp
		}
		return &types.Response{
			Status: types.StatusInternalError,
			Body:   mustMsgpack(map[string]string{"error": err.Error()}),
		}
	}

	return &types.Response{
		Status: types.StatusOK,
		Body:   []byte{},
	}
}

var DefaultConfig = types.Config{
	MaxBodySize: customprotoc.MaxBodySize,
}

func WithConfig(cfg types.Config) ServerOption {
	return func(s *server) {
		s.Config = cfg
	}
}

func WithLogger(logger types.Logger) ServerOption {
	return func(s *server) {
		s.Logger = logger
	}
}

func WithRequestReader(reader types.RequestReader) ServerOption {
	return func(s *server) {
		s.Reader = reader
	}
}

func WithResponseWriter(writer types.ResponseWriter) ServerOption {
	return func(s *server) {
		s.Writer = writer
	}
}

func mustMsgpack(data any) []byte {
	bytes, err := msgpack.Marshal(data)
	if err != nil {
		panic("msgpack marshal failed" + err.Error())
	}

	return bytes
}
