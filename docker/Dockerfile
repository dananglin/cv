FROM golang:1.15.2-buster AS builder

RUN git clone https://github.com/magefile/mage "${GOPATH}/src/mage"

WORKDIR ${GOPATH}/src/mage

RUN go run bootstrap.go

COPY go.mod ${GOPATH}/cv-builder/
COPY go.sum ${GOPATH}/cv-builder/
COPY magefile.go ${GOPATH}/cv-builder/
COPY helpers ${GOPATH}/cv-builder/helpers

WORKDIR ${GOPATH}/cv-builder

RUN mage -compile /usr/local/bin/cv-make

FROM debian:buster

COPY --from=builder /usr/local/bin/cv-make /usr/local/bin

# Install dependencies
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        fonts-crosextra-carlito=20130920-1 \
        rsync=3.1.3-6 \
        curl=7.64.0-4+deb10u1 \
        aspell=0.60.7~20110707-6 \
        aspell-en=2018.04.16-0-1 \
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
