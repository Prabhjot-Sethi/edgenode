# Copyright 2019 Intel Corporation. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM centos:7.6.1810

WORKDIR /root/opae

ENV http_proxy=$http_proxy
ENV https_proxy=$https_proxy

RUN yum install -y gcc gcc-c++ cmake make autoconf automake libxml2 libxml2-devel json-c-devel boost ncurses ncurses-devel ncurses-libs boost-devel libuuid libuuid-devel python2-jsonschema doxygen hwloc-devel libpng12 rsync openssl-devel bc python-devel python-libs python-sphinx openssl unzip which wget python36 epel-release

#Instal Kernel headers
RUN wget http://linuxsoft.cern.ch/cern/centos/7/rt/CentOS-RT.repo -O /etc/yum.repos.d/CentOS-RT.repo
RUN wget http://linuxsoft.cern.ch/cern/centos/7/os/x86_64/RPM-GPG-KEY-cern -O /etc/pki/rpm-gpg/RPM-GPG-KEY-cern
RUN yum install -y "kernel-rt-devel-uname-r == $(uname -r)"
RUN mkdir -p /lib/modules/$(uname -r)
RUN ln -s /usr/src/kernels/$(uname -r) /lib/modules/$(uname -r)/build

#copy packages
COPY n3000-1-3-5-beta-rte-setup.zip .
COPY n3000-1-3-5-beta-cfg-2x2x25g-setup.zip .

#unzip packages
RUN unzip n3000-1-3-5-beta-rte-setup.zip
RUN unzip n3000-1-3-5-beta-cfg-2x2x25g-setup.zip

#make scripts executable
RUN chmod +x n3000-1.3.5-beta-rte-setup.sh
RUN chmod +x n3000-1.3.5-beta-cfg-2x2x25G-setup.sh

#install and copy expect
RUN yum install -y expect-5.45-14.el7_1.x86_64
COPY expect_script.sh .

#install OPAE packages
RUN ./expect_script.sh
RUN source /root/intelrtestack/bin/init_env.sh

#install PAC N3000 configuration
RUN ./n3000-1.3.5-beta-cfg-*-setup.sh -y

#copy module checking script
COPY check_if_modules_loaded.sh .

#install pip and intelhex
RUN easy_install pip && pip install intelhex