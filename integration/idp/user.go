package idp

import (
	//"fmt"
	//"github.com/sugarcrm/multiverse/apis/iam/user/v1alpha"
	"context" // todo use one ????
	//"os"
	"github.com/sugarcrm/multiverse/apis/iam/user/v1alpha"     // "//apis/iam/user/v1alpha:go_default_library",
	disco 	"github.com/sugarcrm/multiverse/projects/discovery/client" //     "//projects/discovery/client:go_default_library",
	"github.com/sugarcrm/multiverse/projects/idm/pkg/sdk"            // "//projects/idm/pkg/sdk:go_default_library",
	"github.com/sugarcrm/multiverse/projects/idm/pkg/srn" // "//projects/idm/pkg/srn:go_default_library"
	"github.com/sugarcrm/multiverse/projects/golib/grpc"             // "//projects/golib/grpc:go_default_library",
	"github.com/sugarcrm/multiverse/projects/golib/oauth2"           // "//projects/golib/oauth2:go_default_library",

)

func GetUser(userSRN string) (*user.User, error)  {
	rn, err := srn.Create(userSRN)
	//userSrn, err := srn.Build(partition, region).User(tenantId, userSRN)
	if err != nil {
		return nil, err
	}

	sdkClient, err := getSDKClient()
	if err != nil {
		return nil, err
	}

	defer sdkClient.Close() // todo it may be should be moved up


	idpc, err := sdkClient.UserAPI(rn.Region) // TODO where i should get Region
	if err != nil {
		return nil, err
	}

	return idpc.GetUser(context.Background(), &user.GetUserRequest{Name: userSRN})
}

func getSDKClient() (*sdk.Client, error) {

	appClient, err := oauth2.LoadClientFromFile(file)
	if err != nil {
		return nil, err
	}

	tokenSource, err := oauth2.NewClientCredentialsTokenSource(context.Background(), appClient, scope)
	if err != nil {
		return nil, err
	}

	discovery := disco.New(discoveryURL)
	discovery.Refresh(context.Background()) // TODO ????

	grpcFactory := grpc.NewClientFactory(grpc.WithInsecure(), grpc.WithDisco(discovery))
	return sdk.NewClient(tokenSource, grpcFactory), nil
}
