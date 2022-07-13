module rikuto110511/elegaku-push-message

go 1.18

replace local.packages/src/elegaku => ./src/elegaku

require github.com/PuerkitoBio/goquery v1.8.0

require (
	github.com/aws/aws-lambda-go v1.32.0
	github.com/line/line-bot-sdk-go v7.8.0+incompatible
)

require (
	github.com/aws/aws-sdk-go v1.44.42
	github.com/guregu/dynamo v1.15.1
	local.packages/src/elegaku v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
)
