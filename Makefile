# 定义一些变量
MYSQL_IMAGE = mysql:latest
MYSQL_CONTAINER_NAME = mysql-container
MYSQL_ROOT_PASSWORD = root

# 拉取 MySQL 镜像
pull_sql:
	docker pull $(MYSQL_IMAGE)

# 创建并启动 MySQL 容器
run_sql:
	docker run --name $(MYSQL_CONTAINER_NAME) \
		-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
		-p 3306:3306 \
		-d $(MYSQL_IMAGE)

# 停止 MySQL 容器
stop_sql:
	docker stop $(MYSQL_CONTAINER_NAME)

# 删除 MySQL 容器
rm_sql:
	docker rm $(MYSQL_CONTAINER_NAME)

# 清理容器和镜像
clean_sql: stop rm
	docker rmi $(MYSQL_IMAGE)

# 打印当前容器和镜像状态
status:
	docker ps -a