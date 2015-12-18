Georges Manager
===============

**This repo is work in progress.**

Demo
----

Execute the following command :
```
echo '127.0.0.1 test.preprod.local' >> /etc/hosts
docker-compose up
mkdir /tmp/test
cd /tmp/test
cat  > docker-compose.yml <<EOF
web:
  image: jlordiales/python-micro-service
  hostname: host1
EOF
tar cfz test.tar.gz docker-compose.yml
curl 127.0.0.1:8080/start-service -XPOST --data-binary '@test.tar.gz' -v
```

Find the port of haproxy :
```docker inspect --format='{{(index (index .NetworkSettings.Ports "80/tcp") 0).HostPort}}' georgesmanager_haproxy_1```

Open your browser, enter : ```http://test.preprod.local:<port>``` :
You should see ```Hello World from host1```
