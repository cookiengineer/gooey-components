
# Gooey Components

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey-components/master/assets/gooey.jpg">
</p>

[Gooey](https://github.com/cookiengineer/gooey) (GUI) is a Pure Go WebASM bindings framework.
It bridges the gaps between Go, WebASM and Browser APIs.

The [Gooey Components](https://github.com/cookiengineer/gooey-components) framework offers ready-to-use
Web Components to structure a Web Application that uses a local Web View for its UI.


# (Native) Program Architecture

- The program starts a local webserver and opens a webview pointing towards it
- The program uses `go:embed` to embed a `/public/*` folder containing all assets
- The program's user interface is built using HTML, CSS, and WebASM.


# WebView HTML Architecture

- Static elements can never be removed from the DOM
- Static elements can have DOM event listeners
- Dynamic elements can be removed from the DOM
- Dynamic elements should not have DOM event listeners
- App Layout always consists of `body > header`, `body > main`, and `body > footer` elements
- App Views always consist of nested `main > section[data-view=...]` elements
- App Views optionally contain `main > section[data-view=...] > aside` elements to represent sidebars


# WebView CSS Architecture

(In the future, it is planned to use `go.rice` and a `toolchain` that offers a bundling methods)

- For now, the default theme is located in [/design](/design). Just copy the design folder so that it
  is located at `/public/design/` and that it is served as static assets. Include the `/design/index.css`
  then and you're set.


# WebView WebASM Architecture

- [app.Main](/pkg/app/Main.go)
- [app.Client](/pkg/app/Client.go)
- [app.ClientListener](/pkg/app/ClientListener.go)
- [app.Storage](/pkg/app/Storage.go)
- [app.View](/pkg/app/View.go)

# WebView WebASM Interfaces

- [interfaces.Component](/pkg/interfaces/Component.go)
- [interfaces.View](/pkg/interfaces/View.go)

# WebView WebASM Components

- [x] [components.Button](/pkg/components/Button.go)
- [x] [components.Label](/pkg/components/Label.go)

**Layout Components**

- [ ] [components.Footer](/pkg/components/Footer.go)
- [ ] [components.Header](/pkg/components/Header.go)
- [ ] [components.Aside](/pkg/components/Aside.go)

**Content Components**

- [ ] [components.Article](/pkg/components/Article.go)
- [ ] [components.Fieldset](/pkg/components/Fieldset.go)
- [ ] [components.Table](/pkg/components/Table.go)

## Work-in-Progress

- [TODO.md](/TODO.md) documents the work-in-progress of things that will be implemented next


## Examples

The [examples](/examples) folder contains minimal test cases that show how you can
use the bindings. They also contain a separate `main.go` which is compiled into a
`main.wasm` file and a `serve.go` which reflects the local webserver.

All examples are served on `http://localhost:3000` if you execute the `build.sh`.

These examples also serve as unit tests, because `go test` cannot generate binaries
for the `syscall/js` platform.

- [app](/examples/app)


## Projects

These are the Projects using `gooey` or `gooey-components` as a library. This list is meant to
showcase how to use the library and how to integrate it with [webview/webview_go](https://github.com/webview/webview_go).

- [Git Evac](https://github.com/cookiengineer/git-evac), a Git Management Tool


# License

This project is licensed under the [MIT](./LICENSE_MIT.txt) license.

