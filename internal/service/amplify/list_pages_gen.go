// Code generated by "internal/generate/listpages/main.go -ListOps=ListApps"; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
)

func listAppsPages(ctx context.Context, conn amplifyiface.AmplifyAPI, input *amplify.ListAppsInput, fn func(*amplify.ListAppsOutput, bool) bool) error {
	for {
		output, err := conn.ListAppsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}