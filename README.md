# First-step-of-deployment
サーバーのデプロイをするためにまず基本的なクソサーバーのデプロイを行いたいと思う。

# サーバーの動かし方
## ローカル　　
1. ターミナルに入力  　
> export PORT="8080"  

2. go run main.go

## Heroku
1. heroku open

# Herokuにデプロイするためにつまずいたこと
1. Procfileの設定
2. デプロイのことを考えるとportの部分のコードを
> 
    port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
    
にするべき