/*
Copyright The ORAS Authors.
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
package remote

import (
	"fmt"
	"strings"

	"oras.land/oras-go/v2/registry"
)

// buildScheme returns HTTP scheme used to access the remote registry.
func buildScheme(plainHTTP bool) string {
	if plainHTTP {
		return "http"
	}
	return "https"
}

// buildRegistryBaseURL builds the URL for accessing the base API.
// Format: <scheme>://<registry>/v2/
// Reference: https://docs.docker.com/registry/spec/api/#base
func buildRegistryBaseURL(plainHTTP bool, ref registry.Reference) string {
	return fmt.Sprintf("%s://%s/v2/", buildScheme(plainHTTP), ref.Host())
}

// buildRegistryCatalogURL builds the URL for accessing the catalog API.
// Format: <scheme>://<registry>/v2/_catalog
// Reference: https://docs.docker.com/registry/spec/api/#catalog
func buildRegistryCatalogURL(plainHTTP bool, ref registry.Reference) string {
	return fmt.Sprintf("%s://%s/v2/_catalog", buildScheme(plainHTTP), ref.Host())
}

// buildRepositoryBaseURL builds the base endpoint of the remote repository.
// Format: <scheme>://<registry>/v2/<repository>
func buildRepositoryBaseURL(plainHTTP bool, ref registry.Reference) string {
	return fmt.Sprintf("%s://%s/v2/%s", buildScheme(plainHTTP), ref.Host(), ref.Repository)
}

// buildRepositoryTagListURL builds the URL for accessing the tag list API.
// Format: <scheme>://<registry>/v2/<repository>/tags/list
// Reference: https://docs.docker.com/registry/spec/api/#tags
func buildRepositoryTagListURL(plainHTTP bool, ref registry.Reference) string {
	return buildRepositoryBaseURL(plainHTTP, ref) + "/tags/list"
}

// buildRepositoryManifestURL builds the URL for accessing the manifest API.
// Format: <scheme>://<registry>/v2/<repository>/manifests/<digest_or_tag>
// Reference: https://docs.docker.com/registry/spec/api/#manifest
func buildRepositoryManifestURL(plainHTTP bool, ref registry.Reference) string {
	return strings.Join([]string{
		buildRepositoryBaseURL(plainHTTP, ref),
		"manifests",
		ref.Reference,
	}, "/")
}

// buildRepositoryBlobURL builds the URL for accessing the blob API.
// Format: <scheme>://<registry>/v2/<repository>/blobs/<digest>
// Reference: https://docs.docker.com/registry/spec/api/#blob
func buildRepositoryBlobURL(plainHTTP bool, ref registry.Reference) string {
	return strings.Join([]string{
		buildRepositoryBaseURL(plainHTTP, ref),
		"blobs",
		ref.Reference,
	}, "/")
}

// buildRepositoryBlobUploadURL builds the URL for blob uploading.
// Format: <scheme>://<registry>/v2/<repository>/blobs/uploads/
// Reference: https://docs.docker.com/registry/spec/api/#initiate-blob-upload
func buildRepositoryBlobUploadURL(plainHTTP bool, ref registry.Reference) string {
	return buildRepositoryBaseURL(plainHTTP, ref) + "/blobs/uploads/"
}

// buildArtifactReferrerURL builds the URL for accessing the manifest referrers
// API.
// Format: <scheme>://<registry>/oras/artifacts/v1/<repository>/manifests/<digest>/referrers
// Reference: https://github.com/oras-project/artifacts-spec/blob/main/manifest-referrers-api.md
func buildArtifactReferrerURL(plainHTTP bool, ref registry.Reference) string {
	return fmt.Sprintf(
		"%s://%s/oras/artifacts/v1/%s/manifests/%s/referrers",
		buildScheme(plainHTTP),
		ref.Host(),
		ref.Repository,
		ref.Reference,
	)
}
