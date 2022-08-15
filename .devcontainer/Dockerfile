FROM amazonlinux:2.0.20220606.1 as base

# install node
RUN curl --silent --location https://rpm.nodesource.com/setup_16.x | bash -

RUN yum update -y \
    && amazon-linux-extras install -y \
    docker \
    postgresql14 \
    epel \
    && amazon-linux-extras enable -y \
    python3.8 \
    && yum install -y \
    figlet \
    file \
    gcc \
    git \
    golang \
    nodejs \
    python38 \
    ruby \
    tar \
    unzip \
    wget \
    which \
    && yum -y clean all \
    && rm -rf /var/cache

RUN python3.8 -m pip install --no-cache-dir pre-commit

RUN wget https://github.com/busyloop/lolcat/archive/master.zip \
    && unzip master.zip \
    && pushd lolcat-master/bin \
    && gem install lolcat \
    && popd \
    && rm -rf master.zip lolcat-master

FROM base as terraform
ENV TERRAFORM_VERSION=1.2.7
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    curl -sL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_arm64.zip -o terraform.zip; \
    else \
    curl -sL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip -o terraform.zip; \
    fi \
    && unzip terraform.zip -d /usr/local/bin \
    && chmod +x \
    /usr/local/bin/terraform \
    && rm -rf \
    terraform.zip

FROM base as taskfile
ENV TASKFILE_VERSION=3.14.1
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    curl -sL https://github.com/go-task/task/releases/download/v${TASKFILE_VERSION}/task_linux_arm64.tar.gz -o taskfile.tar.gz; \
    else \
    curl -sL https://github.com/go-task/task/releases/download/v${TASKFILE_VERSION}/task_linux_amd64.tar.gz -o taskfile.tar.gz; \
    fi \
    && tar -zxvf taskfile.tar.gz \
    && mv task /usr/local/bin/task \
    && chmod +x \
    /usr/local/bin/task \
    && rm -rf \
    taskfile.tar.gz

FROM base as tflint
ENV TFLINK_VERSION=0.39.3
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    curl -sL https://github.com/terraform-linters/tflint/releases/download/v${TFLINK_VERSION}/tflint_linux_arm64.zip -o tflint.zip; \
    else \
    curl -sL https://github.com/terraform-linters/tflint/releases/download/v${TFLINK_VERSION}/tflint_linux_amd64.zip -o tflint.zip; \
    fi \
    && unzip tflint.zip -d /usr/local/bin \
    && chmod +x \
    /usr/local/bin/tflint \
    && rm -rf \
    tflint.zip

FROM base as shellcheck
ENV SHELLCHECK_VERSION=0.8.0
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    curl -sL https://github.com/koalaman/shellcheck/releases/download/v${SHELLCHECK_VERSION}/shellcheck-v${SHELLCHECK_VERSION}.linux.aarch64.tar.xz -o shellcheck.tar.xz; \
    else \
    curl -sL https://github.com/koalaman/shellcheck/releases/download/v${SHELLCHECK_VERSION}/shellcheck-v${SHELLCHECK_VERSION}.linux.x86_64.tar.xz -o shellcheck.tar.xz; \
    fi \
    && tar -xvf shellcheck.tar.xz \
    && mv shellcheck-v${SHELLCHECK_VERSION}/shellcheck /usr/local/bin/shellcheck \
    && chmod +x \
    /usr/local/bin/shellcheck \
    && rm -rf \
    shellcheck.tar.xz

FROM base as docker-compose
RUN export LATEST_COMPOSE_VERSION=$(curl -sSL "https://api.github.com/repos/docker/compose/releases/latest" | grep -o -P '(?<="tag_name": ").+(?=")') \
    && curl -sSL "https://github.com/docker/compose/releases/download/${LATEST_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && chmod +x /usr/local/bin/docker-compose

FROM base as dockerize
ENV DOCKERIZE_VERSION=0.6.1
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    echo "CPU architecture not supported"; \
    else \
    curl -sL https://github.com/jwilder/dockerize/releases/download/v${DOCKERIZE_VERSION}/dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz -o dockerize.tar.gz; \
    fi \
    && tar -C /usr/local/bin -xzvf dockerize.tar.gz \
    && chmod +x /usr/local/bin/dockerize \
    && rm -rf dockerize.tar.gz

FROM base as main
COPY --from=terraform /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=taskfile /usr/local/bin/task /usr/local/bin/task
COPY --from=tflint /usr/local/bin/tflint /usr/local/bin/tflint
COPY --from=shellcheck /usr/local/bin/shellcheck /usr/local/bin/shellcheck
COPY --from=docker-compose /usr/local/bin/docker-compose /usr/local/bin/docker-compose
COPY --from=dockerize /usr/local/bin/dockerize /usr/local/bin/dockerize
COPY .bashrc /root/.bashrc