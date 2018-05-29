FROM centos
WORKDIR /auth
ADD ./simple-auth /auth/
ADD ./conf.ini /auth/
ADD ./registry.html /auth/
EXPOSE 9999
EXPOSE 19999
CMD /bin/bash -c /auth/simple-auth
