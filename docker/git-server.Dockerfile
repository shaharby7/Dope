FROM ubuntu
USER 0
RUN  apt-get update
RUN apt-get install git -y

RUN useradd -ms /bin/bash git

RUN mkdir -p /usr/local/git
RUN chown git:git /usr/local/git
RUN apt-get update
RUN echo -e "    AllowUsers git" |  tee -a /etc/ssh/ssh_config
RUN mkdir -p /home/git/.ssh/
RUN chown -R git:git /home/git/.ssh/
RUN chmod 700 /home/git/.ssh/

USER git
EXPOSE 22