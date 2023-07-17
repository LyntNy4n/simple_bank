package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func filedViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "invalid argument")
	br := &errdetails.BadRequest{
		FieldViolations: violations,
	}

	st, err := st.WithDetails(br)
	if err != nil {
		return st.Err()
	}
	return st.Err()
}

func unauthenticatedError(err error) error {
	return status.Error(codes.Unauthenticated, err.Error())
}
