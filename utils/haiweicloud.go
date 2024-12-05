package utils

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	live "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/region"
)

const (
	HongkongRegion string = "ap-southeast-1"
	SigaporeRegion string = "ap-southeast-3"
)

type Haiweicloud struct {
	ak        string
	sk        string
	ProjectId string
}

func NewHaiweicloud(ak, sk, projectId string) *Haiweicloud {
	return &Haiweicloud{
		ak:        ak,
		sk:        sk,
		ProjectId: projectId,
	}
}

// func (h *Haiweicloud) GetAuth() {
// 	fmt.Println(h.ProjectId)
// 	fmt.Println("here is get stream frame rate data...")
// }

func (h *Haiweicloud) GetStreamFrameRate() (*model.ListSingleStreamFramerateResponse, error) {
	auth, err := basic.NewCredentialsBuilder().
		WithAk(h.ak).
		WithSk(h.sk).
		WithProjectId(h.ProjectId).
		SafeBuild()
	if err != nil {
		fmt.Printf("Failed to init haiweicloud : %v", err)
	}

	client := live.NewLiveClient(
		live.LiveClientBuilder().
			WithRegion(region.ValueOf(SigaporeRegion)).
			WithCredential(auth).
			Build())

	request := &model.ListSingleStreamFramerateRequest{}
	response, err := client.ListSingleStreamFramerate(request)
	return response, err
}
