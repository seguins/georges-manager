FROM haproxy

RUN apt-get update && apt-get install -y wget unzip && apt-get clean
WORKDIR /
RUN wget --no-check-certificate https://releases.hashicorp.com/consul-template/0.12.0/consul-template_0.12.0_linux_amd64.zip
RUN unzip consul-template_0.12.0_linux_amd64.zip
RUN rm consul-template_0.12.0_linux_amd64.zip

ADD haproxy.ctmpl /haproxy.ctmpl

ADD run.sh /run.sh

RUN chmod +x run.sh

CMD ["/run.sh"]
