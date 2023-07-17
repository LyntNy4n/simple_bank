package gapi

import (
	"context"
	db "simple_bank/db/sqlc"
	"simple_bank/pb"
	"simple_bank/validate"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	txResult, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error verifying email: %v", err)
	}

	response := &pb.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return response, nil
}

func validateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, filedViolation("eamil_id", err))
	}
	if err := validate.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, filedViolation("secret_code", err))
	}

	return violations
}
