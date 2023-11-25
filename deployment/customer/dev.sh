apk --no-cache add tzdata

go get github.com/cespare/reflex

go install github.com/cespare/reflex

reflex -r '\.go' -s -- sh -c 'go run /usr/local/go/src/ShopBillBuddy/customer/cmd/main.go -e development'