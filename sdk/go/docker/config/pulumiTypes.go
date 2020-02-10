// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RegistryAuth struct {
	Address string `pulumi:"address"`
	ConfigFile *string `pulumi:"configFile"`
	ConfigFileContent *string `pulumi:"configFileContent"`
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

type RegistryAuthInput interface {
	pulumi.Input

	ToRegistryAuthOutput() RegistryAuthOutput
	ToRegistryAuthOutputWithContext(context.Context) RegistryAuthOutput
}

type RegistryAuthArgs struct {
	Address pulumi.StringInput `pulumi:"address"`
	ConfigFile pulumi.StringPtrInput `pulumi:"configFile"`
	ConfigFileContent pulumi.StringPtrInput `pulumi:"configFileContent"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (RegistryAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryAuth)(nil)).Elem()
}

func (i RegistryAuthArgs) ToRegistryAuthOutput() RegistryAuthOutput {
	return i.ToRegistryAuthOutputWithContext(context.Background())
}

func (i RegistryAuthArgs) ToRegistryAuthOutputWithContext(ctx context.Context) RegistryAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryAuthOutput)
}

type RegistryAuthArrayInput interface {
	pulumi.Input

	ToRegistryAuthArrayOutput() RegistryAuthArrayOutput
	ToRegistryAuthArrayOutputWithContext(context.Context) RegistryAuthArrayOutput
}

type RegistryAuthArray []RegistryAuthInput

func (RegistryAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryAuth)(nil)).Elem()
}

func (i RegistryAuthArray) ToRegistryAuthArrayOutput() RegistryAuthArrayOutput {
	return i.ToRegistryAuthArrayOutputWithContext(context.Background())
}

func (i RegistryAuthArray) ToRegistryAuthArrayOutputWithContext(ctx context.Context) RegistryAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryAuthArrayOutput)
}

type RegistryAuthOutput struct { *pulumi.OutputState }

func (RegistryAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryAuth)(nil)).Elem()
}

func (o RegistryAuthOutput) ToRegistryAuthOutput() RegistryAuthOutput {
	return o
}

func (o RegistryAuthOutput) ToRegistryAuthOutputWithContext(ctx context.Context) RegistryAuthOutput {
	return o
}

func (o RegistryAuthOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func (v RegistryAuth) string { return v.Address }).(pulumi.StringOutput)
}

func (o RegistryAuthOutput) ConfigFile() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegistryAuth) *string { return v.ConfigFile }).(pulumi.StringPtrOutput)
}

func (o RegistryAuthOutput) ConfigFileContent() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegistryAuth) *string { return v.ConfigFileContent }).(pulumi.StringPtrOutput)
}

func (o RegistryAuthOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegistryAuth) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o RegistryAuthOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RegistryAuth) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type RegistryAuthArrayOutput struct { *pulumi.OutputState}

func (RegistryAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryAuth)(nil)).Elem()
}

func (o RegistryAuthArrayOutput) ToRegistryAuthArrayOutput() RegistryAuthArrayOutput {
	return o
}

func (o RegistryAuthArrayOutput) ToRegistryAuthArrayOutputWithContext(ctx context.Context) RegistryAuthArrayOutput {
	return o
}

func (o RegistryAuthArrayOutput) Index(i pulumi.IntInput) RegistryAuthOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) RegistryAuth {
		return vs[0].([]RegistryAuth)[vs[1].(int)]
	}).(RegistryAuthOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryAuthOutput{})
	pulumi.RegisterOutputType(RegistryAuthArrayOutput{})
}