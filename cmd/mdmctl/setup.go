package main

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/micromdm/micromdm/platform/api/server/apply"
	"github.com/micromdm/micromdm/platform/api/server/list"
	"github.com/micromdm/micromdm/platform/blueprint"
	"github.com/micromdm/micromdm/platform/profile"
	"github.com/micromdm/micromdm/platform/remove"
	"github.com/micromdm/micromdm/platform/user"
)

type remoteServices struct {
	profilesvc   profile.Service
	blueprintsvc blueprint.Service
	blocksvc     remove.Service
	usersvc      user.Service
	applysvc     apply.Service
	list         list.Service
}

func setupClient(logger log.Logger) (*remoteServices, error) {
	cfg, err := LoadServerConfig()
	if err != nil {
		return nil, err
	}

	profilesvc, err := profile.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	blueprintsvc, err := blueprint.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	blocksvc, err := remove.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	usersvc, err := user.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	applysvc, err := apply.NewClient(
		cfg.ServerURL, logger, cfg.APIToken,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	listsvc, err := list.NewClient(
		cfg.ServerURL, logger, cfg.APIToken,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	return &remoteServices{
		profilesvc:   profilesvc,
		blueprintsvc: blueprintsvc,
		blocksvc:     blocksvc,
		usersvc:      usersvc,
		applysvc:     applysvc,
		list:         listsvc,
	}, nil
}