package service

import "fmt"

type Version struct {
	Version string `json:"version"`
	Branch  string `json:"branch"`
	Hash    string `json:"hash"`
}

func NewVersion(version, branch, hash string) *Version {
	return &Version{
		Version: version,
		Branch:  branch,
		Hash:    hash,
	}
}

func (v *Version) String() string {
	return fmt.Sprintf("%s-%s (%s)", v.Version, v.Branch, v.Hash)
}
