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
	bndv1 "github.com/trisberg/binding"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BindingCreateOptions struct {
	cli.ResourceOptions

	Host      string
	Port      string
	SecretRef string

	DryRun bool
}

var (
	_ cli.Validatable = (*BindingCreateOptions)(nil)
	_ cli.Executable  = (*BindingCreateOptions)(nil)
	_ cli.DryRunable  = (*BindingCreateOptions)(nil)
)

func (opts *BindingCreateOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := cli.EmptyFieldError

	errs = errs.Also(opts.ResourceOptions.Validate((ctx)))

	return errs
}

func (opts *BindingCreateOptions) Exec(ctx context.Context, c *cli.Config) error {
	binding := &bndv1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: opts.Namespace,
			Name:      opts.Name,
		},
	}

	if opts.SecretRef != "" {
		//handler.Spec.Build = &requestv1alpha1.Build{
		//	ApplicationRef: opts.SecretRef,
		//}
	}
	if opts.Host != "" {
		//handler.Spec.Template.Containers[0].Image = opts.Host
	}

	if opts.DryRun {
		//cli.DryRunResource(ctx, handler, handler.GetGroupVersionKind())
	} else {
		var err error
		binding, err = c.Core().Secrets(opts.Namespace).Create(binding)
		if err != nil {
			return err
		}
	}
	c.Successf("Created binding %q\n", binding.Name)
	return nil
}

func (opts *BindingCreateOptions) IsDryRun() bool {
	return opts.DryRun
}

func NewBindingCreateCommand(ctx context.Context, c *cli.Config) *cobra.Command {
	opts := &BindingCreateOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a binding to map connection properties for a backing service to an application or function",
		Long: strings.TrimSpace(`
<todo>
`),
		Example: strings.Join([]string{
			fmt.Sprintf("%s binding create my-binding %s my-service.default.svc.cluster.local %s 1234 %s my-service-secret", c.Name, cli.HostFlagName, cli.PortFlagName, cli.SecretRefFlagName),
		}, "\n"),
		Args: cli.Args(
			cli.NameArg(&opts.Name),
		),
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.NamespaceFlag(cmd, c, &opts.Namespace)
	cmd.Flags().StringVar(&opts.Host, cli.StripDash(cli.HostFlagName), "", "hostname of service to bind")
	cmd.Flags().StringVar(&opts.Port, cli.StripDash(cli.PortFlagName), "", "port of service to bind")
	cmd.Flags().StringVar(&opts.SecretRef, cli.StripDash(cli.SecretRefFlagName), "", "`name` of secret for the service")

	return cmd
}
