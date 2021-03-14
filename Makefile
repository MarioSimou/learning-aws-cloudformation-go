SHELL := bash
ARTIFACTS_DIR ?= ./artifacts
ARTIFACTS_BUCKET ?= serverless-architecture-code-snippets
DEFAULT_LAMBDA ?= createMedia

build:
	@for filename in {createMedia,changeMediaPermissions}; do \
		go build -o main ./lambdas/$${filename}/main.go; \
		zip -r -j ${ARTIFACTS_DIR}/$${filename}.zip main; \
		rm -rf main &> /dev/null; \
	done

test:
	@go build -o main ./lambdas/${DEFAULT_LAMBDA}/main.go
	@sam local invoke -e ./lambdas/${DEFAULT_LAMBDA}/event.json ${DEFAULT_LAMBDA}Lambda
	@rm -rf main &> /dev/null

deploy:
	@sam validate -t template.yaml
	@make build
	@aws s3 cp ./artifacts s3://${ARTIFACTS_BUCKET}/ --recursive --exclude ".gitkeep" --include "*.zip"
	@sam deploy -t template.yaml --stack-name serverless-architecture --region eu-west-1

drop:
	@aws cloudformation delete-stack --stack-name serverless-architecture --region eu-west-1

format:
	@go fmt ./...