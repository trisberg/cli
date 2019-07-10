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
	"time"

	"github.com/projectriff/cli/pkg/cli"
	"github.com/projectriff/cli/pkg/cli/printers"
	"github.com/spf13/cobra"
	servicev1alpha1 "github.com/trisberg/service/pkg/apis/service/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

type BindingListOptions struct {
	cli.ListOptions
}

var (
	_ cli.Validatable = (*BindingListOptions)(nil)
	_ cli.Executable  = (*BindingListOptions)(nil)
)

func (opts *BindingListOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := cli.EmptyFieldError

	errs = errs.Also(opts.ListOptions.Validate(ctx))

	return errs
}

func (opts *BindingListOptions) Exec(ctx context.Context, c *cli.Config) error {
	bindings, err := c.Service().Bindings(opts.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	if len(bindings.Items) == 0 {
		c.Infof("No bindings found.\n")
		return nil
	}

	tablePrinter := printers.NewTablePrinter(printers.PrintOptions{
		WithNamespace: opts.AllNamespaces,
	}).With(func(h printers.PrintHandler) {
		columns := printBindingColumns()
		h.TableHandler(columns, printBindingList)
		h.TableHandler(columns, printBinding)
	})

	bindings = bindings.DeepCopy()
	cli.SortByNamespaceAndName(bindings.Items)

	return tablePrinter.PrintObj(bindings, c.Stdout)
}

func NewBindingListCommand(ctx context.Context, c *cli.Config) *cobra.Command {
	opts := &BindingListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "table listing of bindings",
		Long: strings.TrimSpace(`
<todo>
`),
		Example: strings.Join([]string{
			fmt.Sprintf("%s binding list", c.Name),
			fmt.Sprintf("%s binding list %s", c.Name, cli.AllNamespacesFlagName),
		}, "\n"),
		Args:    cli.Args(),
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.AllNamespacesFlag(cmd, c, &opts.Namespace, &opts.AllNamespaces)

	return cmd
}

func printBindingList(bindings *servicev1alpha1.BindingList, opts printers.PrintOptions) ([]metav1beta1.TableRow, error) {
	rows := make([]metav1beta1.TableRow, 0, len(bindings.Items))
	for i := range bindings.Items {
		r, err := printBinding(&bindings.Items[i], opts)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printBinding(binding *servicev1alpha1.Binding, opts printers.PrintOptions) ([]metav1beta1.TableRow, error) {
	now := time.Now()
	row := metav1beta1.TableRow{
		Object: runtime.RawExtension{Object: binding},
	}
	refType, refValue := bindingRef(binding)
	row.Cells = append(row.Cells,
		binding.Name,
		refType + ":" +refValue,
		cli.FormatTimestampSince(binding.CreationTimestamp, now),
	)
	return []metav1beta1.TableRow{row}, nil
}

func printBindingColumns() []metav1beta1.TableColumnDefinition {
	return []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string"},
		{Name: "Ref", Type: "string"},
		{Name: "Age", Type: "string"},
	}
}

func bindingRef(binding *servicev1alpha1.Binding) (string, string) {
	if binding.Spec.SecretRef != "" {
		return "secret", binding.Spec.SecretRef
	}
	return cli.Swarnf("secret"), cli.Swarnf("<unknown>")
}
