//go:build linux || windows
// +build linux windows

/*
Copyright 2021 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package core

import (
	"context"
	"time"

	"github.com/Mirantis/cri-dockerd/utils"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// ImageFsStatsCache caches imagefs stats.
var ImageFsStatsCache utils.Cache

const imageFsStatsMinTTL = 30 * time.Second

// ImageFsInfo returns information of the filesystem that is used to store images.
func (ds *dockerService) ImageFsInfo(
	_ context.Context,
	_ *runtimeapi.ImageFsInfoRequest,
) (*runtimeapi.ImageFsInfoResponse, error) {

	res, err := ImageFsStatsCache.Memoize("imagefs", imageFsStatsMinTTL, func() (interface{}, error) {
		return ds.imageFsInfo()
	})
	if err != nil {
		return nil, err
	}
	stats := res.(*runtimeapi.ImageFsInfoResponse)
	return stats, nil

}
