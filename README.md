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

- Deploy from github (recommended)
- Deploy from cli

![vercel dashboard](https://raw.githubusercontent.com/ojii3/vercel-go-test/main/media/vercel-go-test-build-settings.png)

### Deploy from GitHub

1. Write code
2. Push to GitHub
3. At vercel dashboard, add project from GitHub

Vercel will automatically deploy when you push to GitHub.

They also setup preview environment(with different url) for each pull request.

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

---

# vercel go test

これは、私が vercel と go を使う方法を学ぶためのテストプロジェクトです。

- <https://ojii3-vercel-go-test.vercel.app/> (public/index.html)
- <https://ojii3-vercel-go-test.vercel.app/api/hello> (api/hello.go)
- <https://ojii3-vercel-go-test.vercel.app/api/index> (api/index.go)

## 次のステップ

- [ ] データベース(vercel kv)
- [ ] フロントエンド(tinygo wasm)
- [ ] GraphQL？

これらは、異なるブランチで試してみます。

## デプロイ方法

デプロイする方法は 2 つあります。

- GitHub からデプロイする（推奨）
- CLI からデプロイする

### GitHub からデプロイする

1. コードを書く
2. GitHub にプッシュする
3. Vercel のダッシュボードで、GitHub からプロジェクトを追加する

GitHub にプッシュすると、Vercel が自動的にデプロイします。

また、各プルリクエストごとにプレビュー環境（異なる URL）もセットアップされます。

### CLI からデプロイする

Node.js と npm が必要です。

1. Vercel CLI をインストールする

```sh
npm install -g vercel
```

プロジェクトに package.json を作成しないように注意してください。

2. ログインする

```sh
vercel login
```

3. デプロイする

```sh
vercel deploy
```

## ローカルでの開発方法

1. Vercel CLI をインストールする

```sh
npm install -g vercel
```

2. ログインする

```sh
Copy code
vercel login
```

3. 実行する

```sh
vercel dev
```
