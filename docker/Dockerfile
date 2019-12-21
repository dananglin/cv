FROM golang:1.13.5-buster

# Upgrade packages and install dependencies
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        fonts-crosextra-carlito \
        rsync \
    && \
    apt-get clean autoclean && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/* /tmp/*

# Install ConTeXt standalone
RUN \
    mkdir /opt/context && \
    cd /opt/context && \
    curl -LO http://minimals.contextgarden.net/setup/first-setup.sh && \
    sh ./first-setup.sh \
        --context=current \
        --engine=luatex \
    && \
    rm -rf /opt/context/tex/texmf-context/doc

ENV PATH=${PATH}:/opt/context/tex/texmf-linux-64/bin \
    OSFONTDIR=/usr/share/fonts

RUN mtxrun --script fonts --reload
