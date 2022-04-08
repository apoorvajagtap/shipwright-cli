## shp build upload

Run a Build with local data

### Synopsis


Creates a new BuildRun instance and instructs the Build Controller to wait for the data streamed,
instead of executing "git clone". Therefore, you can employ Shipwright Builds from a local repository
clone.

The upload skips the ".git" directory completely, and it follows the ".gitignore" directives, when
the file is found at the root of the directory uploaded.

	$ shp buildrun upload <build-name>
	$ shp buildrun upload <build-name> /path/to/repository


```
shp build upload <build-name> [path/to/source|.] [flags]
```

### Options

```
      --buildref-apiversion string            API version of build resource to reference
      --buildref-name string                  name of build resource to reference
  -e, --env stringArray                       specify a key-value pair for an environment variable to set for the build container (default [])
  -F, --follow                                Start a build and watch its log until it completes or fails.
  -h, --help                                  help for upload
      --output-credentials-secret string      name of the secret with builder-image pull credentials
      --output-image string                   image employed during the building process
      --output-image-annotation stringArray   specify a set of key-value pairs that correspond to annotations to set on the output image (default [])
      --output-image-label stringArray        specify a set of key-value pairs that correspond to labels to set on the output image (default [])
      --sa-generate                           generate a Kubernetes service-account for the build
      --sa-name string                        Kubernetes service-account name
      --timeout duration                      build process timeout
```

### Options inherited from parent commands

```
      --kubeconfig string        Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string         If present, the namespace scope for this CLI request
      --request-timeout string   The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
```

### SEE ALSO

* [shp build](shp_build.md)	 - Manage Builds
