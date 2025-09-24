package cli

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/coder/coder/v2/cli/cliui"
	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/serpent"
)

func (r *RootCmd) aibridge() *serpent.Command {
	cmd := &serpent.Command{
		Use:   "aibridge",
		Short: "Manage aibridge.",
		Handler: func(inv *serpent.Invocation) error {
			return inv.Command.HelpHandler(inv)
		},
		Children: []*serpent.Command{
			r.listInvocations(),
		},
	}
	return cmd
}

func (r *RootCmd) listInvocations() *serpent.Command {
	var (
		formatter = cliui.NewOutputFormatter(
			cliui.TableFormat(
				[]WorkspaceListRow{},
				[]string{},
			),
			cliui.JSONFormat(),
		)
		start      string
		end        string
		initiatior string
		limit      int64
		cursor     string
	)

	return &serpent.Command{
		Use:   "list-invocations",
		Short: "List aibridge interceptions",
		Options: serpent.OptionSet{
			{
				Flag:        "period-start",
				Description: "Time .",
				Default:     time.Time{}.Format(time.RFC3339),
				Value:       serpent.StringOf(&start),
			},
			{
				Flag:        "period-end",
				Description: "Time .",
				Default:     time.Time{}.Format(time.RFC3339),
				Value:       serpent.StringOf(&end),
			},
			{
				Flag:        "limit",
				Description: "Limit.",
				Default:     "100",
				Value:       serpent.Int64Of(&limit),
			},
			{
				Flag:        "initiator",
				Description: "Initiator.",
				Default:     uuid.Nil.String(),
				Value:       serpent.StringOf(&initiatior),
			},
			{
				Flag:        "cursor",
				Description: "Cursor.",
				Default:     "{}",
				Value:       serpent.StringOf(&cursor),
			},
		},
		Handler: func(inv *serpent.Invocation) error {
			client, err := r.InitClient(inv)
			if err != nil {
				return err
			}

			initID, err := uuid.Parse(initiatior)
			if err != nil {
				return err
			}
			startTime, err := time.Parse(time.RFC3339, start)
			if err != nil {
				return err
			}
			endTime, err := time.Parse(time.RFC3339, end)
			if err != nil {
				return err
			}
			cursorObj := &codersdk.AIBridgeListInterceptionsCursor{}
			err = json.Unmarshal([]byte(cursor), cursorObj)
			if err != nil {
				return err
			}
			if limit > math.MaxInt32 {
				return xerrors.Errorf("limit value to high")
			}
			if limit < 0 {
				return xerrors.Errorf("limit value negative")
			}

			exCli := codersdk.NewExperimentalClient(client)
			resp, err := exCli.AIBridgeListInterceptions(inv.Context(), codersdk.AIBridgeListInterceptionsRequest{
				PeriodStart: startTime,
				PeriodEnd:   endTime,
				InitiatorID: initID,
				Limit:       int32(limit),
				Cursor:      *cursorObj,
			})
			if err != nil {
				return err
			}

			out, err := formatter.Format(inv.Context(), resp)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(inv.Stdout, out)
			return err
		},
	}
}
