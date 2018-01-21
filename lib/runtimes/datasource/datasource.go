package datasource

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jysperm/deploying/config"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/jysperm/deploying/lib/utils"
)

var swarmClient *client.Client

const RegistryAuthParam = "deploying"

func init() {
	var err error
	swarmClient, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}

func MakeRedisImage() error {
	redisAssets := utils.GetAssetFilePath("datasource-redis/")
	redisTag := fmt.Sprintf("%s/redis:latest", config.DefaultRegistry)
	buildOpts := types.ImageBuildOptions{
		Tags:           []string{redisTag},
		Dockerfile:     "Dockerfile",
		NoCache:        false,
		Remove:         true,
		SuppressOutput: false,
	}

	buildCtx, err := buildContext(redisAssets)
	if err != nil {
		return err
	}

	defer buildCtx.Close()

	res, err := swarmClient.ImageBuild(context.Background(), buildCtx, buildOpts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		return err
	}

	response, err := swarmClient.ImagePush(context.Background(), redisTag, types.ImagePushOptions{All: true, RegistryAuth: RegistryAuthParam})
	if err != nil {
		return err
	}

	defer response.Close()

	if _, err := io.Copy(os.Stdout, response); err != nil {
		return err
	}

	return nil
}

func MakeMognoImage() error {
	mongoAssets := utils.GetAssetFilePath("datasource-mongodb/")
	mongoTag := fmt.Sprintf("%s/mongodb:latest", config.DefaultRegistry)
	buildOpts := types.ImageBuildOptions{
		Tags:           []string{mongoTag},
		Dockerfile:     "Dockerfile",
		NoCache:        false,
		Remove:         true,
		SuppressOutput: false,
	}

	buildCtx, err := buildContext(mongoAssets)
	if err != nil {
		return err
	}

	defer buildCtx.Close()

	res, err := swarmClient.ImageBuild(context.Background(), buildCtx, buildOpts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		return err
	}

	response, err := swarmClient.ImagePush(context.Background(), mongoTag, types.ImagePushOptions{All: true, RegistryAuth: RegistryAuthParam})
	if err != nil {
		return err
	}

	defer response.Close()

	if _, err := io.Copy(os.Stdout, response); err != nil {
		return err
	}

	return nil
}

func buildContext(path string) (io.ReadCloser, error) {
	content, err := archive.Tar(path, archive.Gzip)
	if err != nil {
		return nil, err
	}

	return content, nil
}