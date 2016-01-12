package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/pqcnstr"
	"github.com/piotrkowalczuk/sklog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreateUser implements charon.RPCServer interface.
func (rs *rpcServer) CreateUser(ctx context.Context, req *charon.CreateUserRequest) (*charon.CreateUserResponse, error) {
	var err error
	defer func() {
		if err != nil {
			sklog.Error(rs.logger, err)
		} else {
			sklog.Debug(rs.logger, "user created")
		}
	}()

	actor, err := rs.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	if !actor.user.IsSuperuser {
		if req.IsSuperuser != nil && req.IsSuperuser.Valid {
			return nil, grpc.Errorf(codes.PermissionDenied, "charond: user is not allowed to create superuser")
		}

		if req.IsStaff != nil && req.IsStaff.Valid && !actor.permissions.Contains(charon.UserCanCreateStaff) {
			return nil, grpc.Errorf(codes.PermissionDenied, "charond: user is not allowed to create staff user")
		}
	}

	if req.SecurePassword == "" {
		req.SecurePassword, err = rs.passwordHasher.Hash(req.PlainPassword)
		if err != nil {
			return nil, err
		}
	} else {
		if !actor.user.IsSuperuser {
			return nil, grpc.Errorf(codes.PermissionDenied, "charond: only superuser can create an user with manualy defined secure password")
		}
	}

	entity, err := rs.repository.user.Create(
		req.Username,
		req.SecurePassword,
		req.FirstName,
		req.LastName,
		uuid.New(),
		req.IsSuperuser.BoolOr(false),
		req.IsStaff.BoolOr(false),
		req.IsActive.BoolOr(false),
		req.IsConfirmed.BoolOr(false),
	)
	if err != nil {
		return nil, mapUserError(err)
	}

	return &charon.CreateUserResponse{
		User: entity.Message(),
	}, nil
}

// ModifyUser implements charon.RPCServer interface.
func (rs *rpcServer) ModifyUser(ctx context.Context, req *charon.ModifyUserRequest) (*charon.ModifyUserResponse, error) {
	if int64(req.Id) <= 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "charond: user cannot be modified, invalid id: %d", int64(req.Id))
	}

	actor, err := rs.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	entity, err := rs.repository.user.FindOneByID(int64(req.Id))
	if err != nil {
		return nil, err
	}

	if hint, ok := modifyUserFirewall(req, entity, actor); !ok {
		return nil, grpc.Errorf(codes.PermissionDenied, "charond: "+hint)
	}

	entity, err = rs.repository.user.UpdateOneByID(
		int64(req.Id),
		req.Username,
		req.SecurePassword,
		req.FirstName,
		req.LastName,
		req.IsSuperuser,
		req.IsActive,
		req.IsStaff,
		req.IsConfirmed,
	)
	if err != nil {
		return nil, mapUserError(err)
	}

	sklog.Debug(rs.logger, "user modified", "id", int64(req.Id))

	return &charon.ModifyUserResponse{
		User: entity.Message(),
	}, nil
}

func modifyUserFirewall(req *charon.ModifyUserRequest, entity *userEntity, actor *actor) (string, bool) {
	isOwner := actor.user.ID == entity.ID

	if !actor.user.IsSuperuser {
		switch {
		case entity.IsSuperuser:
			return "only superuser can modify a superuser account", false
		case entity.IsStaff && !isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsStranger):
			return "missing permission to modify an account as a stranger", false
		case entity.IsStaff && isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsOwner):
			return "missing permission to modify an account as an owner", false
		case req.IsSuperuser != nil && req.IsSuperuser.Valid:
			return "only superuser can change existing account to superuser", false
		case req.IsStaff != nil && req.IsStaff.Valid && !actor.permissions.Contains(charon.UserCanCreateStaff):
			return "user is not allowed to create user with is_staff property that has custom value", false
		}
	}

	return "", true
}

// GetUser implements charon.RPCServer interface.
func (rs *rpcServer) GetUser(ctx context.Context, req *charon.GetUserRequest) (*charon.GetUserResponse, error) {
	user, err := rs.repository.user.FindOneByID(int64(req.Id))
	if err != nil {
		return nil, err
	}

	sklog.Debug(rs.logger, "user retrieved", "id", int64(req.Id))

	return &charon.GetUserResponse{
		User: user.Message(),
	}, nil
}

// ListUsers implements charon.RPCServer interface.
func (rs *rpcServer) ListUsers(ctx context.Context, req *charon.ListUsersRequest) (*charon.ListUsersResponse, error) {
	users, err := rs.repository.user.Find(req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	resp := &charon.ListUsersResponse{
		Users: make([]*charon.User, 0, len(users)),
	}

	for _, u := range users {
		resp.Users = append(resp.Users, u.Message())
	}

	sklog.Debug(rs.logger, "users list retrieved", "count", len(users))

	return resp, nil
}

// DeleteUser implements charon.RPCServer interface.
func (rs *rpcServer) DeleteUser(ctx context.Context, req *charon.DeleteUserRequest) (*charon.DeleteUserResponse, error) {
	if int64(req.Id) <= 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "charond: user cannot be deleted, invalid id: %d", int64(req.Id))
	}
	affected, err := rs.repository.user.DeleteOneByID(int64(req.Id))
	if err != nil {
		return nil, err
	}

	sklog.Debug(rs.logger, "users deleted", "id", int64(req.Id))

	return &charon.DeleteUserResponse{
		Affected: affected,
	}, nil
}

func mapUserError(err error) error {
	switch pqcnstr.FromError(err) {
	case tableUserConstraintPrimaryKey:
		return grpc.Errorf(codes.AlreadyExists, charon.ErrDescUserWithIDExists)
	case tableUserConstraintUniqueUsername:
		return grpc.Errorf(codes.AlreadyExists, charon.ErrDescUserWithUsernameExists)
	default:
		return err
	}
}
