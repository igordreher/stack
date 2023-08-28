package clients

import (
	"flag"
	"fmt"
	"github.com/formancehq/fctl/pkg/config"

	"github.com/formancehq/fctl/pkg/modelutils"

	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	useDelete   = "delete <client-id>"
	shortDelete = "Delete client with id"
)

type DeleteStore struct {
	Success bool `json:"success"`
}

func NewDeleteStore() *DeleteStore {
	return &DeleteStore{}
}

func NewDeleteConfig() *config.ControllerConfig {
	flags := flag.NewFlagSet(useDelete, flag.ExitOnError)

	fctl.WithConfirmFlag(flags)
	return config.NewControllerConfig(
		useDelete,
		shortDelete,
		shortDelete,
		[]string{
			"del", "d",
		},
		flags,
		config.Organization, config.Stack,
	)
}

var _ config.Controller[*DeleteStore] = (*DeleteController)(nil)

type DeleteController struct {
	store  *DeleteStore
	config *config.ControllerConfig
}

func NewDeleteController(config *config.ControllerConfig) *DeleteController {
	return &DeleteController{
		store:  NewDeleteStore(),
		config: config,
	}
}

func (c *DeleteController) GetStore() *DeleteStore {
	return c.store
}

func (c *DeleteController) GetConfig() *config.ControllerConfig {
	return c.config
}

func (c *DeleteController) Run() (modelutils.Renderable, error) {

	flags := c.config.GetAllFLags()
	ctx := c.config.GetContext()
	args := c.config.GetArgs()
	out := c.config.GetOut()
	cfg, err := fctl.GetConfig(flags)
	if err != nil {
		return nil, err
	}

	organizationID, err := fctl.ResolveOrganizationID(flags, ctx, cfg, out)
	if err != nil {
		return nil, err
	}

	stack, err := fctl.ResolveStack(flags, ctx, cfg, organizationID, out)
	if err != nil {
		return nil, err
	}

	if !fctl.CheckStackApprobation(flags, stack, "You are about to delete an OAuth2 client") {
		return nil, fctl.ErrMissingApproval
	}

	authClient, err := fctl.NewStackClient(flags, ctx, cfg, stack, out)
	if err != nil {
		return nil, err
	}

	request := operations.DeleteClientRequest{
		ClientID: args[0],
	}
	response, err := authClient.Auth.DeleteClient(ctx, request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	c.store.Success = true

	return c, nil
}

func (c *DeleteController) Render() error {
	pterm.Success.WithWriter(c.config.GetOut()).Printfln("Client deleted!")
	return nil
}

func NewDeleteCommand() *cobra.Command {

	config := NewDeleteConfig()
	return fctl.NewCommand(config.GetUse(),
		fctl.WithArgs(cobra.ExactArgs(1)),
		fctl.WithController[*DeleteStore](NewDeleteController(config)),
	)
}
