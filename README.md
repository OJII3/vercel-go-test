# vercel go test

This is a test project for me to learn how to use vercel with go.

- <https://ojii3-vercel-go-test.vercel.app/> (public/index.html)
- <https://ojii3-vercel-go-test.vercel.app/api/hello> (api/hello.go)
- <https://ojii3-vercel-go-test.vercel.app/api/index> (api/index.go)

## Next steps

- [ ] database(vercel kv)
- [ ] frontend(tinygo wasm)
- [ ] graphql?

I will try these at different branches.

## How to deploy

There are two ways to deploy.

- Deploy from cli
- Deploy from github

### Deploy from cli

You need nodejs and npm.

1. Install vercel cli

```sh
npm install -g vercel
```

Be careful not to create package.json to the project.

2. Login

```sh
vercel login
```

3. Deploy

```sh
vercel deploy
```

### Deploy from GitHub

1. Push to GitHub
2. At vercel dashboard, add project from GitHub

Vercel will automatically deploy when you push to GitHub.

They also setup preview environment(with different url) for each pull request.

## How to develop locally

1. Install vercel cli

```sh
npm install -g vercel
```

2. Login

```sh
vercel login
```

3. Run

```sh
vercel dev
```
