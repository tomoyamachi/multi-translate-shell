multi translation tool by golang
---

# Build
This service use Google Translation API.

You need signup GCP account, and enable the Cloud Translation API for your project.

```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/gcloud-auth.json
```

# Package dependencies

```
$ go get github.com/marcusolsson/tui-go
$ go get cloud.google.com/go/translate
```
