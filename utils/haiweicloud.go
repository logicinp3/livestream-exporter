package utils

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	live "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2"
	liveModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
	liveRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/region"
	"time"
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

func (h *Haiweicloud) GetStreamFrameRate(domain, appName, streamName string) (*liveModel.ListSingleStreamFramerateResponse, error) {
	auth, err := basic.NewCredentialsBuilder().
		WithAk(h.ak).
		WithSk(h.sk).
		WithProjectId(h.ProjectId).
		SafeBuild()
	if err != nil {
		fmt.Printf("Failed to init credential: %v\n", err)
	}

	region, err := liveRegion.SafeValueOf(SigaporeRegion)
	if err != nil {
		fmt.Printf("Failed to get region info: %v\n", err)
	}

    // build client
	hcClient, err := live.LiveClientBuilder().
		WithRegion(region).
		WithCredential(auth).
		SafeBuild()
	if err != nil {
		fmt.Printf("Failed to init client: %v\n", err)
	}
	client := live.NewLiveClient(hcClient)

    // build request
	currentTime := time.Now().UTC()
	startTime := currentTime.Add(-1 * time.Minute)
	startTimeRequest := startTime.Format("2006-01-02T15:04:05Z")
	request := &liveModel.ListSingleStreamFramerateRequest{
        Domain: domain,
        App: appName,
        Stream: streamName,
        StartTime: &startTimeRequest,
    }
	response, err := client.ListSingleStreamFramerate(request)
	return response, err
}

func (h *Haiweicloud) GetStreamBitRate(domain, appName, streamName string) (*liveModel.ListSingleStreamBitrateResponse, error) {
	auth, err := basic.NewCredentialsBuilder().
		WithAk(h.ak).
		WithSk(h.sk).
		WithProjectId(h.ProjectId).
		SafeBuild()
	if err != nil {
		fmt.Printf("Failed to init credential: %v\n", err)
	}

	region, err := liveRegion.SafeValueOf(SigaporeRegion)
	if err != nil {
		fmt.Printf("Failed to get region info: %v\n", err)
	}

    // build client
	hcClient, err := live.LiveClientBuilder().
		WithRegion(region).
		WithCredential(auth).
		SafeBuild()
	if err != nil {
		fmt.Printf("Failed to init client: %v\n", err)
	}
	client := live.NewLiveClient(hcClient)

    // build request
	currentTime := time.Now().UTC()
	startTime := currentTime.Add(-1 * time.Minute)
	startTimeRequest := startTime.Format("2006-01-02T15:04:05Z")
	request := &liveModel.ListSingleStreamBitrateRequest{
        Domain: domain,
        App: appName,
        Stream: streamName,
        StartTime: &startTimeRequest,
    }
	response, err := client.ListSingleStreamBitrate(request)
	return response, err
}
