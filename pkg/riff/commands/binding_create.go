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
	servicev1 "github.com/trisberg/service/pkg/apis/service/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BindingCreateOptions struct {
	cli.ResourceOptions

	SecretRef   string
	URI         string
	URIKey      string
	Host        string
	HostKey     string
	Port        string
	PortKey     string
	Username    string
	UsernameKey string
	PasswordKey string

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
	binding := &servicev1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: opts.Namespace,
			Name:      opts.Name,
		},
	}

	if opts.SecretRef != "" {
		binding.Spec.SecretRef = opts.SecretRef
	}
	if opts.URI != "" {
		binding.Spec.URI = opts.URI
	}
	if opts.URIKey != "" {
		binding.Spec.URIKey = opts.URIKey
	}
	if opts.Host != "" {
		binding.Spec.Host = opts.Host
	}
	if opts.HostKey != "" {
		binding.Spec.HostKey = opts.HostKey
	}
	if opts.Port != "" {
		binding.Spec.Port = opts.Port
	}
	if opts.PortKey != "" {
		binding.Spec.PortKey = opts.PortKey
	}
	if opts.Username != "" {
		binding.Spec.Username = opts.Username
	}
	if opts.UsernameKey != "" {
		binding.Spec.UsernameKey = opts.UsernameKey
	}
	if opts.PasswordKey != "" {
		binding.Spec.PasswordKey = opts.PasswordKey
	}

	if opts.DryRun {
		//cli.DryRunResource(ctx, handler, handler.GetGroupVersionKind())
	} else {
		var err error
		binding, err = c.Service().Bindings(opts.Namespace).Create(binding)
		if err != nil {
			return err
		}
	}
	c.Successf("Created binding.service %q\n", binding.Name)
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
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.Args(cmd,
		cli.NameArg(&opts.Name),
	)

	cli.NamespaceFlag(cmd, c, &opts.Namespace)
	cmd.Flags().StringVar(&opts.SecretRef, cli.StripDash(cli.SecretRefFlagName), "", "name of secret for the service")
	cmd.Flags().StringVar(&opts.URI, cli.StripDash(cli.URIFlagName), "", "URI of service to bind")
	cmd.Flags().StringVar(&opts.URIKey, cli.StripDash(cli.URIKeyFlagName), "", "the key for URI in the secret")
	cmd.Flags().StringVar(&opts.Host, cli.StripDash(cli.HostFlagName), "", "hostname of service to bind")
	cmd.Flags().StringVar(&opts.HostKey, cli.StripDash(cli.HostKeyFlagName), "", "the key for hostname in the secret")
	cmd.Flags().StringVar(&opts.Port, cli.StripDash(cli.PortFlagName), "", "port of service to bind")
	cmd.Flags().StringVar(&opts.PortKey, cli.StripDash(cli.PortKeyFlagName), "", "the key for port in the secret")
	cmd.Flags().StringVar(&opts.Username, cli.StripDash(cli.UsernameFlagName), "", "username to use when connecting to service")
	cmd.Flags().StringVar(&opts.UsernameKey, cli.StripDash(cli.UsernameKeyFlagName), "", "the key for username in the secret")
	cmd.Flags().StringVar(&opts.PasswordKey, cli.StripDash(cli.PasswordKeyFlagName), "", "the key for the password in the secret")

	return cmd
}
