module github.com/el10savio/lwwset-crdt

go 1.15

require (
	github.com/el10savio/twoPSet-crdt v0.0.0-20201120171725-e585748e5362
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/el10savio/lwwset-crdt/handlers => ./handlers
