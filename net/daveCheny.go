package main

import (
	"net"
	"time"
	"crypto/tls"
)

//https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type Server struct {
	listerner net.Listener
	timeout   time.Duration
}

func (s *Server) Addr() {}

func (s *Server) ShutDown() {

}

func NewServer(addr string) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := Server{
		listerner: l,
	}
	go srv.run()
	return &srv, nil
}

//fix once
func NewServer1(addr string, clintTimeout time.Duration, maxconns, maxconcurrent int, cert *tls.Certificate) {

}

//fix twice
func NewServer2(addr string) (*Server, error) {
	return nil, nil
}

func NewTLSServer(addr string, cert *tls.Certificate) (*Server, error) {
	return nil, nil
}

func NewServerWithTimeout(addr string, timeout time.Duration) (*Server, error) {
	return nil, nil
}

func NewTLSServerWithTimeout(addr string, certificate *tls.Certificate, timeout time.Duration) (*Server, error) {
	return nil, nil
}

//fix third
type Config struct {
	Timeout time.Duration
	Cert    *tls.Certificate
}

func NewServer3(addr string, config Config) (*Server, error) {
	return nil, nil
}

func New(addr string) {
	srv, _ := NewServer3(addr, Config{}) //使用零值
}

//fix fifth
func NewServer5(addr string, config *Config) (*Server, error) {
	return nil, nil
}

func New1() {
	srv, _ := NewServer5("localhost", nil) //accept default

	config := Config{
		Timeout: time.Second,
	}
	srv2, _ := NewServer5("localhost", &config)
	config.Timeout = time.Hour
}

//fix sixth
func NewServer6(addr string, config ...Config) (*Server, error) {
	return nil, nil
}

func New2() {
	srv, _ := NewServer6("localhost")

	srv2, _ := NewServer6("localhost", Config{
		Timeout: 300 * time.Second,
	})
	//doSomething
}

//fix seventh
func NewServer7(addr string, options ...func(*Server)) (*Server, error) {
	return nil, nil
}

func New3() {
	srv, _ := NewServer7("localhost") //defaults

	timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
	}

	srv2, _ := NewServer7("localhost", timeout)
}

//fix eighth
func NewServer8(addr string, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{listerner: l}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}
