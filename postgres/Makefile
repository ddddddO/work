# WSL内で完結
fetchdata:
	./maketasks/fetch-dbdump-from-pi.sh

pubsubclean:
	./maketasks/cleanup-pubsub.sh

pubsubup:
	./maketasks/launch-pubsub.sh


# 以下からPowerShell側でも使える処理を記載
# k8s関係のタスクランナーはPowerShellからも実行できるようにshは使わないこと
pubsubbuild:
	docker build -t pub -f Dockerfiles/app/pub/Dockerfile .
	docker build -t sub -f Dockerfiles/app/sub/Dockerfile .
	docker build -t db -f Dockerfiles/db/Dockerfile .
	docker image prune -f

pubsubbuildgke: pubsubbuild
	docker tag pub asia.gcr.io/work1111/pub
	docker tag sub asia.gcr.io/work1111/sub
	docker tag db asia.gcr.io/work1111/db

pubsubpush:
	docker push asia.gcr.io/work1111/pub
	docker push asia.gcr.io/work1111/sub
	docker push asia.gcr.io/work1111/db
