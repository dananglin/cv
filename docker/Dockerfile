FROM golang:1.14.6-buster

# Upgrade packages and install dependencies
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        fonts-crosextra-carlito=20130920-1 \
        rsync=3.1.3-6 \
    && \
    apt-get clean autoclean && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/* /tmp/* && \
    mkdir /opt/context

WORKDIR /opt/context

# Install ConTeXt standalone
RUN \
    curl -L http://minimals.contextgarden.net/setup/first-setup.sh -o /opt/context/first-setup.sh && \
    sh /opt/context/first-setup.sh \
        --context=current \
        --engine=luatex \
    && \
    rm -rf /opt/context/tex/texmf-context/doc

ENV PATH=${PATH}:/opt/context/tex/texmf-linux-64/bin \
    OSFONTDIR=/usr/share/fonts

RUN mtxrun --script fonts --reload
