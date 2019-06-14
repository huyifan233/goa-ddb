package calcapi

import (
	"context"
	calc "goa-ddb/gen/calc"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return
}

// Addresume implements addresume.
func (s *calcsrvc) Addresume(ctx context.Context, p []*calc.Resume) (res []int, err error) {
	s.logger.Print("calc.addresume")
	return
}

// List implements list.
func (s *calcsrvc) List(ctx context.Context) (res calc.StoredResumeCollection, err error) {
	s.logger.Print("calc.list")
	return
}
