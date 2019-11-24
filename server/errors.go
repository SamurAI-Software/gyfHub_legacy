package main

import "errors"

// ErrEmailExistIssue is for bad request
var ErrEmailExistIssue = errors.New("EMAIL_EXIST_ISSUE")

// ErrUsernameExistIssue is for bad request
var ErrUsernameExistIssue = errors.New("USERNAME_EXIST_ISSUE")

// ErrMobileExistIssue is for bad request
var ErrMobileExistIssue = errors.New("MOBILE_EXIST_ISSUE")

// ErrVerifiedIssue is for unverified error
var ErrVerifiedIssue = errors.New("VERIFIED_ISSUE")

// ErrorHubNotExistIssue is for exist Hub error
var ErrorHubNotExistIssue = errors.New("HUB_NOT_EXIST_ISSUE")
