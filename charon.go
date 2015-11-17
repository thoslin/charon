package charon

import (
	"errors"

	"github.com/piotrkowalczuk/mnemosyne"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	contextKeyUser  = "context_key_charon_user"
	DecisionUnknown = 0
	DecisionGranted = 1
)

// NewContext returns a new Context that carries User value.
func NewContext(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, contextKeyUser, u)
}

// FromContext returns the User value stored in context, if any.
func FromContext(ctx context.Context) (User, bool) {
	u, ok := ctx.Value(contextKeyUser).(User)
	return u, ok
}

// Error ...
type Error struct {
	StatusCode   int
	InternalCode int
	Message      string
	Validation   map[string][]string
}

// Error ...
func (e *Error) Error() string {
	return e.Message
}

// AddValidation ...
func (e *Error) AddValidation(key, value string) {
	if e.Validation[key] == nil {
		e.Validation[key] = make([]string, 0, 1)
	}

	e.Validation[key] = append(e.Validation[key], value)
}

type AuthorizationChecker func(context.Context, Permission, ...interface{}) (bool, error)

type Charon interface {
	IsGranted(context.Context, Permission, ...interface{}) (bool, error)
	IsAuthenticated(context.Context) (bool, error)
	Login(string, string) (*mnemosyne.Session, error)
	Logout(*mnemosyne.Token) error
}

type charon struct {
	client  RPCClient
	checker AuthorizationChecker
}

// CharonOpts ...
type CharonOpts struct {
	AuthorizationChecker AuthorizationChecker
}

// New allocates new Charon instance.
func New(conn *grpc.ClientConn, options CharonOpts) Charon {
	return &charon{
		client: NewRPCClient(conn),
	}
}

// IsGranted ...
func (c *charon) IsGranted(ctx context.Context, perm Permission, args ...interface{}) (bool, error) {
	var req *IsGrantedRequest

	user, ok := FromContext(ctx)
	if !ok {
		session, ok := mnemosyne.FromContext(ctx)
		if !ok {
			return false, errors.New("charon: permission cannot be checked, user nor session exists in context")
		}

		req = &IsGrantedRequest{
			Token: session.Token.String(),
		}
	} else {
		req = &IsGrantedRequest{
			UserId: user.Id,
		}
	}

	res, err := c.client.IsGranted(ctx, req)
	if err != nil {
		return false, err
	}

	if res.IsGranted {
		return true, nil
	}

	if c.checker == nil {
		return false, nil
	}

	return c.checker(ctx, perm, args...)
}

// IsAuthenticated ...
func (c *charon) IsAuthenticated(ctx context.Context) (bool, error) {
	session, ok := mnemosyne.FromContext(ctx)
	if !ok {
		return false, errors.New("charon: is not authenticated, missing session in context")
	}

	res, err := c.client.IsAuthenticated(ctx, &IsAuthenticatedRequest{
		Token: session.Token,
	})
	if err != nil {
		return false, err
	}

	return res.IsAuthenticated, nil
}

func (c *charon) Login(username, password string) (*mnemosyne.Session, error) {
	return nil, nil
}

func (c *charon) Logout(token *mnemosyne.Token) error {
	return nil
}