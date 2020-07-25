Use Docker to deploy a web site supporting HTTPS with Tomcat as the server. The principle is similar to that of deploying on a common server. You only need to put the software package developed by yourself in the Tomcat project directory. The difference is that the Tomcat server must be started in the container, and the corresponding project source code must also be imported into the container.

## 1.Choose a basic image
    docker pull tomcat
    docker pull -a tomcat   #pull all versions of tomcat

## 2.Make a HTTPS server image
To deploy an HTTPS WEB site, the first step is to make the HTTP Tomcat base image support HTTPS.
### 2.1 Generate the certificate required for HTTPS
    keytool -genkeypair \
        -alias tomcat \
        -keyalg RSA \
        -keysize 4096 \
        -keypass 123456 \
        -dname "cn=localhost,ou=localhost,o=localhost,l=tj,st=tj,c=CN" \
        -validity 365 \
        -keystore tomcat.keystore \
        -storepass 123456

### 2.2 Import the certificate
    # docker run -it -v /root/ssl/:/tmp/ tomcat bash
    root@8f8c200f632c:/usr/local/tomcat# ls
    LICENSE  NOTICE  RELEASE-NOTES  RUNNING.txt  bin  conf  include  lib  logs  native-jni-lib  temp  webapps  work
    root@8f8c200f632c:/usr/local/tomcat# ls /tmp/
    tomcat.keystore
    root@8f8c200f632c:/usr/local/tomcat# mkdir keys
    root@8f8c200f632c:/usr/local/tomcat# cp /tmp/tomcat.keystore keys/  

### 2.3 Modify tomcat configuration and commit
    <Connector port="8080" protocol="HTTP/1.1"
     connectionTimeout="20000"
     redirectPort="8443" 
     SSLEnabled="true" scheme="https" secure="true" clientAuth="false" sslProtocol="TLS" keystoreFile="/usr/local/tomcat/keys/tomcat.keystore" keystorePass="123456" /> 
     
     docker commit 8f8c200f632c tomcat:https

## 3. Import the web source code into the Tomcat image
### 3.1 Import static
    COPY  ./websrc   /usr/local/tomcat/webapps/myproj/
### 3.2 Dynamic mount
    RUN mkdir -p /usr/local/tomcat/webapps/myproj
    VOLUME /usr/local/tomcat/webapps/myproj

    docker run -ti -v $(pwd)/websrc:/usr/local/tomcat/webapps/myproj

## 4. Deployment and Verification
    docker run -d -p 80:8080 -v websrc:/usr/local/tomcat/webapps/myproj myweb:v1
    
Now we can use browsers like Chrome to visit this site.
