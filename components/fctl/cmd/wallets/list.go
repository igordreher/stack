package wallets

import (
	"flag"
	"fmt"
	"github.com/formancehq/fctl/pkg/config"

	"github.com/formancehq/fctl/pkg/modelutils"

	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	useList   = "list"
	shortList = "List all wallets"
)

type ListStore struct {
	Wallets []shared.Wallet `json:"wallets"`
}
type ListController struct {
	store  *ListStore
	config *config.ControllerConfig
}

func NewListStore() *ListStore {
	return &ListStore{}
}
func NewListConfig() *config.ControllerConfig {
	flags := flag.NewFlagSet(useList, flag.ExitOnError)
	fctl.WithMetadataFlag(flags)
	return config.NewControllerConfig(
		useList,
		shortList,
		shortList,
		[]string{
			"list",
			"ls",
		},
		flags,
		config.Organization, config.Stack,
	)
}

var _ config.Controller[*ListStore] = (*ListController)(nil)

func NewListController(config *config.ControllerConfig) *ListController {
	return &ListController{
		store:  NewListStore(),
		config: config,
	}
}

func (c *ListController) GetStore() *ListStore {
	return c.store
}

func (c *ListController) GetConfig() *config.ControllerConfig {
	return c.config
}

func (c *ListController) Run() (modelutils.Renderable, error) {
	flags := c.config.GetAllFLags()
	ctx := c.config.GetContext()
	out := c.config.GetOut()
	cfg, err := fctl.GetConfig(flags)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving config")
	}

	organizationID, err := fctl.ResolveOrganizationID(flags, ctx, cfg, out)
	if err != nil {
		return nil, err
	}

	stack, err := fctl.ResolveStack(flags, ctx, cfg, organizationID, out)
	if err != nil {
		return nil, err
	}

	client, err := fctl.NewStackClient(flags, ctx, cfg, stack, out)
	if err != nil {
		return nil, errors.Wrap(err, "creating stack client")
	}

	metadata, err := fctl.ParseMetadata(fctl.GetStringSlice(flags, config.MetadataFlag))
	if err != nil {
		return nil, err
	}

	response, err := client.Wallets.ListWallets(ctx, operations.ListWalletsRequest{
		Metadata: metadata,
	})
	if err != nil {
		return nil, errors.Wrap(err, "listing wallet")
	}

	if response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	c.store.Wallets = response.ListWalletsResponse.Cursor.Data

	return c, nil
}

func (c *ListController) Render() error {
	if len(c.store.Wallets) == 0 {
		fmt.Fprintln(c.config.GetOut(), "No wallets found.")
		return nil
	}

	if err := pterm.DefaultTable.
		WithHasHeader(true).
		WithWriter(c.config.GetOut()).
		WithData(
			fctl.Prepend(
				fctl.Map(c.store.Wallets,
					func(src shared.Wallet) []string {
						return []string{
							src.ID,
							src.Name,
						}
					}),
				[]string{"ID", "Name"},
			),
		).Render(); err != nil {
		return errors.Wrap(err, "rendering table")
	}
	return nil
}
func NewListCommand() *cobra.Command {
	c := NewListConfig()
	return fctl.NewCommand(c.GetUse(),
		fctl.WithArgs(cobra.ExactArgs(0)),
		fctl.WithController[*ListStore](NewListController(c)),
	)
}
