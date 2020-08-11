install-minikube:
	chmod +x bin/minikube && sudo mkdir -p /usr/local/bin/ && sudo install bin/minikube /usr/local/bin/
install-kubectl:
	chmod +x bin/kubectl && sudo mkdir -p /usr/local/bin/ && sudo install bin/kubectl /usr/local/bin/
install-docker:
	curl -fsSL https://get.docker.com -o get-docker.sh && sudo sh get-docker.sh && sudo usermod -aG docker `whoami`
install-others:
	sudo apt-get install conntrack -y
env-setup:install-minikube install-kubectl install-docker install-others
minikube-start:
	minikube start --driver=none --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'
minikube-dashboard:
	minikube dashboard

