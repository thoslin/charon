package main

import (
	"database/sql"

	"github.com/piotrkowalczuk/charon"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreateGroup implements charon.RPCServer interface.
func (rs *rpcServer) CreateGroup(ctx context.Context, req *charon.CreateGroupRequest) (*charon.CreateGroupResponse, error) {
	actor, err := rs.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	if !actor.permissions.Contains(charon.GroupCanCreate) {
		return nil, grpc.Errorf(codes.PermissionDenied, "charond: actor do not have permission: %s", charon.GroupCanCreate.String())
	}

	entity, err := rs.repository.group.Create(actor.user.ID, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	return &charon.CreateGroupResponse{
		Group: entity.Message(),
	}, nil
}

// ModifyGroup implements charon.RPCServer interface.
func (rs *rpcServer) ModifyGroup(ctx context.Context, req *charon.ModifyGroupRequest) (*charon.ModifyGroupResponse, error) {
	actor, err := rs.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	group, err := rs.repository.group.UpdateOneByID(int64(req.Id), actor.user.ID, req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	return &charon.ModifyGroupResponse{
		Group: group.Message(),
	}, nil
}

// DeleteGroup implements charon.RPCServer interface.
func (rs *rpcServer) DeleteGroup(ctx context.Context, req *charon.DeleteGroupRequest) (*charon.DeleteGroupResponse, error) {
	affected, err := rs.repository.group.DeleteOneByID(int64(req.Id))
	if err != nil {
		return nil, err
	}

	return &charon.DeleteGroupResponse{
		Affected: affected,
	}, nil
}

// GetGroup implements charon.RPCServer interface.
func (rs *rpcServer) GetGroup(ctx context.Context, req *charon.GetGroupRequest) (*charon.GetGroupResponse, error) {
	entity, err := rs.repository.group.FindOneByID(int64(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, grpc.Errorf(codes.NotFound, "charond: group with id %d does not exists", int64(req.Id))
		}
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &charon.GetGroupResponse{
		Group: entity.Message(),
	}, nil
}

// ListGroups implements charon.RPCServer interface.
func (rs *rpcServer) ListGroups(ctx context.Context, req *charon.ListGroupsRequest) (*charon.ListGroupsResponse, error) {

	return nil, grpc.Errorf(codes.Unimplemented, "charond: list groups endpoint is not implemented yet")
}
