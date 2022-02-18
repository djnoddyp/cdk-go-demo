package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	gateway "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	lambda "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CoffeeshopStackProps struct {
	awscdk.StackProps
}

func NewCoffeeshopStack(scope constructs.Construct, id string, props *CoffeeshopStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("CoffeeshopQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	// The Go executables can be made smaller by stripping out the symbol table and debug
	// information (-s) and omitting the DWARF symbol table (-w). This doesn’t make stack traces
	// unreadable, so it’s quite appropriate for production use.
	bundlingOptions := &lambda.BundlingOptions{
		GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
	}

	orderLambda := lambda.NewGoFunction(stack, jsii.String("OrderHandler"), &lambda.GoFunctionProps{
		Entry:    jsii.String("lambdas/order.go"),
		Bundling: bundlingOptions,
	})

	gateway.NewLambdaRestApi(stack, jsii.String("CoffeeShopApi"), &gateway.LambdaRestApiProps{
		Handler: orderLambda,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCoffeeshopStack(app, "CoffeeshopStack", &CoffeeshopStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
