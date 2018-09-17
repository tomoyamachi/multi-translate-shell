multi translation tool by golang
---

# Demo

![demo](https://user-images.githubusercontent.com/890517/45605983-4ae47f80-ba7c-11e8-8dd7-0966e90b7884.gif)

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
$ go get github.com/atotto/clipboard
```

# Extending using languages

You can add languages to main.go and run `go build`.
[Here](https://cloud.google.com/translate/docs/languages) is usable ISO-639-1 code.

```
languages = []translator.TranslateLanguage{
    // {from language code, to language code}
	{"de", "en"},
	{"de", "ja"},
	...
}
```

# Shortcut
- Alt+Enter : Execute translate
- Ctrl+v : Paste from clipboard
- Ctrl+c or Esc : Kill this application
- Ctrl+a : Jump to beginning of the line.
- Ctrl+e : Go to end of the line.
- Ctrl+k : Delete the line.

# Attention

This library doesn't use any clever algorithm for word wrapping. The wrapping is actually very naive: whenever there is whitespace or an explicit linebreak. The goal of this library is for word wrapping CLI output, so the input is typically pretty well controlled human language. Because of this, the naive approach typically works just fine.
