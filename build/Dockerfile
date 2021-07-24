FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/awecloud-sysoperator \
    USER_UID=1001 \
    USER_NAME=awecloud-sysoperator

# install operator binary
COPY build/_output/bin/awecloud-sysoperator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
