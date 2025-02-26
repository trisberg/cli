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
	buildv1alpha1 "github.com/projectriff/system/pkg/apis/build/v1alpha1"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

type FunctionListOptions struct {
	cli.ListOptions
}

var (
	_ cli.Validatable = (*FunctionListOptions)(nil)
	_ cli.Executable  = (*FunctionListOptions)(nil)
)

func (opts *FunctionListOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := cli.EmptyFieldError

	errs = errs.Also(opts.ListOptions.Validate(ctx))

	return errs
}

func (opts *FunctionListOptions) Exec(ctx context.Context, c *cli.Config) error {
	functions, err := c.Build().Functions(opts.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	if len(functions.Items) == 0 {
		c.Infof("No functions found.\n")
		return nil
	}

	tablePrinter := printers.NewTablePrinter(printers.PrintOptions{
		WithNamespace: opts.AllNamespaces,
	}).With(func(h printers.PrintHandler) {
		columns := opts.printColumns()
		h.TableHandler(columns, opts.printList)
		h.TableHandler(columns, opts.print)
	})

	functions = functions.DeepCopy()
	cli.SortByNamespaceAndName(functions.Items)

	return tablePrinter.PrintObj(functions, c.Stdout)
}

func NewFunctionListCommand(ctx context.Context, c *cli.Config) *cobra.Command {
	opts := &FunctionListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "table listing of functions",
		Long: strings.TrimSpace(`
List functions in a namespace or across all namespaces.

For detail regarding the status of a single function, run:

    ` + c.Name + ` function status <function-name>
`),
		Example: strings.Join([]string{
			fmt.Sprintf("%s function list", c.Name),
			fmt.Sprintf("%s function list %s", c.Name, cli.AllNamespacesFlagName),
		}, "\n"),
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.AllNamespacesFlag(cmd, c, &opts.Namespace, &opts.AllNamespaces)

	return cmd
}

func (opts *FunctionListOptions) printList(functions *buildv1alpha1.FunctionList, printOpts printers.PrintOptions) ([]metav1beta1.TableRow, error) {
	rows := make([]metav1beta1.TableRow, 0, len(functions.Items))
	for i := range functions.Items {
		r, err := opts.print(&functions.Items[i], printOpts)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func (opts *FunctionListOptions) print(function *buildv1alpha1.Function, _ printers.PrintOptions) ([]metav1beta1.TableRow, error) {
	now := time.Now()
	row := metav1beta1.TableRow{
		Object: runtime.RawExtension{Object: function},
	}
	row.Cells = append(row.Cells,
		function.Name,
		cli.FormatEmptyString(function.Status.LatestImage),
		cli.FormatEmptyString(function.Spec.Artifact),
		cli.FormatEmptyString(function.Spec.Handler),
		cli.FormatEmptyString(function.Spec.Invoker),
		cli.FormatConditionStatus(function.Status.GetCondition(buildv1alpha1.FunctionConditionReady)),
		cli.FormatTimestampSince(function.CreationTimestamp, now),
	)
	return []metav1beta1.TableRow{row}, nil
}

func (opts *FunctionListOptions) printColumns() []metav1beta1.TableColumnDefinition {
	return []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string"},
		{Name: "Latest Image", Type: "string"},
		{Name: "Artifact", Type: "string"},
		{Name: "Handler", Type: "string"},
		{Name: "Invoker", Type: "string"},
		{Name: "Status", Type: "string"},
		{Name: "Age", Type: "string"},
	}
}
