/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/projectriff/cli/pkg/cli"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BindingDeleteOptions struct {
	cli.DeleteOptions
}

var (
	_ cli.Validatable = (*BindingDeleteOptions)(nil)
	_ cli.Executable  = (*BindingDeleteOptions)(nil)
)

func (opts *BindingDeleteOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := cli.EmptyFieldError

	errs = errs.Also(opts.DeleteOptions.Validate(ctx))

	return errs
}

func (opts *BindingDeleteOptions) Exec(ctx context.Context, c *cli.Config) error {
	client := c.Service().Bindings(opts.Namespace)

	bindingNames := []string{}
	if opts.All {
		bindings, err := c.Service().Bindings(opts.Namespace).List(metav1.ListOptions{})
		if err != nil {
			return err
		}
		if len(bindings.Items) == 0 {
			c.Infof("No bindings found.\n")
			return nil
		}
		for i := range bindings.Items {
			bindingNames = append(bindingNames, bindings.Items[i].Name)
		}
	} else {
		bindingNames = append(bindingNames, opts.Names...)
	}

	for _, name := range bindingNames {
		if err := client.Delete(name, nil); err != nil {
			return err
		}
		c.Successf("Deleted binding %q\n", name)
	}

	return nil
}

func NewBindingDeleteCommand(ctx context.Context, c *cli.Config) *cobra.Command {
	opts := &BindingDeleteOptions{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete binding(s)",
		Long: strings.TrimSpace(`
<todo>
`),
		Example: strings.Join([]string{
			fmt.Sprintf("%s binding delete my-binding", c.Name),
			fmt.Sprintf("%s binding delete %s ", c.Name, cli.AllFlagName),
		}, "\n"),
		Args: cli.Args(
			cli.NamesArg(&opts.Names),
		),
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.NamespaceFlag(cmd, c, &opts.Namespace)
	cmd.Flags().BoolVar(&opts.All, cli.StripDash(cli.AllFlagName), false, "delete all bindings within the namespace")

	return cmd
}
