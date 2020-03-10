module github.com/aianman4823/sample-layered-architecture

go 1.14

replace local.packages/domain => ./domain
replace local.packages/domain/repository => ./domain/repository
replace local.packages/config => ./config
replace local.packages/interfaces/response => ./interfaces/response
replace local.packages/interfaces/handler => ./interfaces/handler
replace local.packages/usecase => ./usecase
replace local.packages/infrastructure/persistence => ./infrastructure/persistence
