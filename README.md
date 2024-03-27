<img src="./docs/assets/logo.svg" alt="HomeOps Logo" width="144px" height="144px" align="left"/>

<div align="center">

### My Home Operations Repository :octocat:

_... managed with Flux and Renovate, and GitHub Actions_ :robot:

[![Node-Count](https://img.shields.io/endpoint?url=https%3A%2F%2Fkromgo.gabe565.com%2Fquery%3Fmetric%3Dcluster_node_count&style=flat&label=Nodes)](https://github.com/kashalls/kromgo/)&nbsp;&nbsp;
[![Pod-Count](https://img.shields.io/endpoint?url=https%3A%2F%2Fkromgo.gabe565.com%2Fquery%3Fmetric%3Dcluster_pod_count&style=flat&label=Pods)](https://github.com/kashalls/kromgo/)&nbsp;&nbsp;
[![CPU-Usage](https://img.shields.io/endpoint?url=https%3A%2F%2Fkromgo.gabe565.com%2Fquery%3Fmetric%3Dcluster_cpu_usage&style=flat&label=CPU)](https://github.com/kashalls/kromgo/)&nbsp;&nbsp;
[![Memory-Usage](https://img.shields.io/endpoint?url=https%3A%2F%2Fkromgo.gabe565.com%2Fquery%3Fmetric%3Dcluster_memory_usage&style=flat&label=Memory)](https://github.com/kashalls/kromgo/)&nbsp;&nbsp;
[![Power-Usage](https://img.shields.io/endpoint?url=https%3A%2F%2Fkromgo.gabe565.com%2Fquery%3Fmetric%3Dcluster_power_usage&style=flat&label=Power)](https://github.com/kashalls/kromgo/)

</div>

## 📖 Overview

This is a mono repository for my home infrastructure and Kubernetes cluster. I try to adhere to Infrastructure as Code (IaC) and GitOps practices using tools like [Terraform](https://www.terraform.io/), [Kubernetes](https://kubernetes.io/), [FluxCD](https://github.com/fluxcd/flux2), [Renovate](https://github.com/renovatebot/renovate), and [GitHub Actions](https://github.com/features/actions).

## ⛵ Kubernetes

There is a template over at [onedr0p/flux-cluster-template](https://github.com/onedr0p/flux-cluster-template) if you want to try and follow along with some of the practices I use here.

### Installation

My cluster is [k3s](https://k3s.io/) provisioned overtop bare-metal Ubuntu. This is a semi-hyper-converged cluster, workloads and block storage are sharing the same available resources on my nodes while I have a separate server with BTRFS for NFS/SMB shares, bulk file storage and backups.

### Core Components

- [cilium](https://github.com/cilium/cilium): internal Kubernetes networking plugin
- [cert-manager](https://cert-manager.io/docs/): creates SSL certificates for services in my cluster
- [external-dns](https://github.com/kubernetes-sigs/external-dns): automatically syncs DNS records from my cluster ingresses to a DNS provider
- [ingress-nginx](https://github.com/kubernetes/ingress-nginx/): ingress controller for Kubernetes using NGINX as a reverse proxy and load balancer
- [sops](https://toolkit.fluxcd.io/guides/mozilla-sops/): managed secrets for Kubernetes, Ansible, and Terraform which are committed to Git

### GitOps

[FluxCD](https://github.com/fluxcd/flux2) watches the clusters in my [kubernetes](./kubernetes/) folder (see Directories below) and makes the changes to my clusters based on the state of my Git repository.

The way Flux works for me here is it will recursively search the `kubernetes/${cluster}/apps` folder until it finds the most top level `kustomization.yaml` per directory and then apply all the resources listed in it. That aforementioned `kustomization.yaml` will generally only have a namespace resource and one or many Flux kustomizations (`ks.yaml`). Under the control of those Flux kustomizations there will be a `HelmRelease` or other resources related to the application which will be applied.

[Renovate](https://github.com/renovatebot/renovate) watches my **entire** repository looking for dependency updates, when they are found a PR is automatically created. When some PRs are merged Flux applies the changes to my cluster.

### Directories

This Git repository contains the following directories under [Kubernetes](./kubernetes/).

```sh
📁 kubernetes
├── 📁 apps           # applications
├── 📁 bootstrap      # bootstrap procedures
├── 📁 flux           # core flux configuration
└── 📁 templates      # re-useable components
```

### App Listing

<!-- Begin apps section -->
<table>
  <tr>
    <th>Namespace</th>
    <th>Name</th>
    <th>Supporting Services</th>
  </tr>
  <tr>
    <td rowspan="2">adguard-home</td>
    <td><a href="kubernetes/apps/adguard-home/app/helmrelease.yaml">adguard-home</a></td>
    <td rowspan="2"><a href="kubernetes/apps/adguard-home/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/adguard-home/external-dns/helmrelease.yaml">adguard-home-external-dns</a></td>
  </tr>
  <tr>
    <td>ascii-movie</td>
    <td><a href="kubernetes/apps/ascii-movie/app/helmrelease.yaml">ascii-movie</a></td>
    <td></td>
  </tr>
  <tr>
    <td>authentik</td>
    <td><a href="kubernetes/apps/authentik/app/helmrelease.yaml">authentik</a></td>
    <td><a href="kubernetes/apps/authentik/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/authentik/app/keydb.yaml">keydb</a>, <a href="kubernetes/apps/authentik/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>bookstack</td>
    <td><a href="kubernetes/apps/bookstack/app/helmrelease.yaml">bookstack</a></td>
    <td><a href="kubernetes/apps/bookstack/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>castsponsorskip</td>
    <td><a href="kubernetes/apps/castsponsorskip/app/helmrelease.yaml">castsponsorskip</a></td>
    <td></td>
  </tr>
  <tr>
    <td>cert-manager</td>
    <td><a href="kubernetes/apps/cert-manager/app/helmrelease.yaml">cert-manager</a></td>
    <td></td>
  </tr>
  <tr>
    <td>change-detection</td>
    <td><a href="kubernetes/apps/change-detection/app/helmrelease.yaml">change-detection</a></td>
    <td></td>
  </tr>
  <tr>
    <td>charm</td>
    <td><a href="kubernetes/apps/charm/app/helmrelease.yaml">charm</a></td>
    <td></td>
  </tr>
  <tr>
    <td>chatgpt</td>
    <td><a href="kubernetes/apps/chatgpt/app/helmrelease.yaml">chatgpt</a></td>
    <td></td>
  </tr>
  <tr>
    <td>cnpg-system</td>
    <td><a href="kubernetes/apps/cnpg/app/helmrelease.yaml">cnpg</a></td>
    <td></td>
  </tr>
  <tr>
    <td>domain-watch</td>
    <td><a href="kubernetes/apps/domain-watch/app/helmrelease.yaml">domain-watch</a></td>
    <td></td>
  </tr>
  <tr>
    <td>esphome</td>
    <td><a href="kubernetes/apps/esphome/app/helmrelease.yaml">esphome</a></td>
    <td><a href="kubernetes/apps/esphome/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="2">external-dns</td>
    <td><a href="kubernetes/apps/external-dns/dynamic-ip/helmrelease.yaml">dynamic-ip</a></td>
    <td rowspan="2"></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/external-dns/app/helmrelease.yaml">external-dns</a></td>
  </tr>
  <tr>
    <td>generic-device-plugin</td>
    <td><a href="kubernetes/apps/generic-device-plugin/app/helmrelease.yaml">generic-device-plugin</a></td>
    <td></td>
  </tr>
  <tr>
    <td>gitea</td>
    <td><a href="kubernetes/apps/gitea/app/helmrelease.yaml">gitea</a></td>
    <td><a href="kubernetes/apps/gitea/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/gitea/app/keydb.yaml">keydb</a>, <a href="kubernetes/apps/gitea/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>hammond</td>
    <td><a href="kubernetes/apps/hammond/app/helmrelease.yaml">hammond</a></td>
    <td></td>
  </tr>
  <tr>
    <td>headscale</td>
    <td><a href="kubernetes/apps/headscale/app/helmrelease.yaml">headscale</a></td>
    <td><a href="kubernetes/apps/headscale/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/headscale/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>healthchecks</td>
    <td><a href="kubernetes/apps/healthchecks/app/helmrelease.yaml">healthchecks</a></td>
    <td><a href="kubernetes/apps/healthchecks/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/healthchecks/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="2">home-assistant</td>
    <td><a href="kubernetes/apps/home-assistant/app/helmrelease.yaml">home-assistant</a></td>
    <td rowspan="2"><a href="kubernetes/apps/home-assistant/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/home-assistant/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/home-assistant/piper/helmrelease.yaml">piper</a></td>
  </tr>
  <tr>
    <td>homepage</td>
    <td><a href="kubernetes/apps/homepage/app/helmrelease.yaml">homepage</a></td>
    <td></td>
  </tr>
  <tr>
    <td>immich</td>
    <td><a href="kubernetes/apps/immich/app/helmrelease.yaml">immich</a></td>
    <td><a href="kubernetes/apps/immich/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/immich/app/keydb.yaml">keydb</a>, <a href="kubernetes/apps/immich/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>ingress-nginx</td>
    <td><a href="kubernetes/apps/ingress-nginx/app/helmrelease.yaml">ingress-nginx</a></td>
    <td></td>
  </tr>
  <tr>
    <td>kromgo</td>
    <td><a href="kubernetes/apps/prometheus/kromgo/helmrelease.yaml">kromgo</a></td>
    <td></td>
  </tr>
  <tr>
    <td rowspan="3">kube-system</td>
    <td><a href="kubernetes/apps/cilium/app/helmrelease.yaml">cilium</a></td>
    <td rowspan="3"></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/kube-vip/app/helmrelease.yaml">kube-vip</a></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/metrics-server/app/helmrelease.yaml">metrics-server</a></td>
  </tr>
  <tr>
    <td>lidarr</td>
    <td><a href="kubernetes/apps/lidarr/app/helmrelease.yaml">lidarr</a></td>
    <td></td>
  </tr>
  <tr>
    <td>linkding</td>
    <td><a href="kubernetes/apps/linkding/app/helmrelease.yaml">linkding</a></td>
    <td><a href="kubernetes/apps/linkding/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/linkding/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>longhorn-system</td>
    <td><a href="kubernetes/apps/longhorn/app/helmrelease.yaml">longhorn</a></td>
    <td></td>
  </tr>
  <tr>
    <td>matrimony</td>
    <td><a href="kubernetes/apps/matrimony/app/helmrelease.yaml">matrimony</a></td>
    <td><a href="kubernetes/apps/matrimony/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>memos</td>
    <td><a href="kubernetes/apps/memos/app/helmrelease.yaml">memos</a></td>
    <td><a href="kubernetes/apps/memos/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/memos/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>metallb-system</td>
    <td><a href="kubernetes/apps/metallb/app/helmrelease.yaml">metallb</a></td>
    <td></td>
  </tr>
  <tr>
    <td>miniflux</td>
    <td><a href="kubernetes/apps/miniflux/app/helmrelease.yaml">miniflux</a></td>
    <td><a href="kubernetes/apps/miniflux/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/miniflux/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>minio</td>
    <td><a href="kubernetes/apps/minio/app/helmrelease.yaml">minio</a></td>
    <td><a href="kubernetes/apps/minio/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>monica</td>
    <td><a href="kubernetes/apps/monica/app/helmrelease.yaml">monica</a></td>
    <td><a href="kubernetes/apps/monica/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>mosquitto</td>
    <td><a href="kubernetes/apps/mosquitto/app/helmrelease.yaml">mosquitto</a></td>
    <td></td>
  </tr>
  <tr>
    <td>nextcloud</td>
    <td><a href="kubernetes/apps/nextcloud/app/helmrelease.yaml">nextcloud</a></td>
    <td><a href="kubernetes/apps/nextcloud/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/nextcloud/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>nfs-provisioner</td>
    <td><a href="kubernetes/apps/nfs-subdir-external-provisioner/app/helmrelease.yaml">nfs-subdir-external-provisioner</a></td>
    <td></td>
  </tr>
  <tr>
    <td>nightscout</td>
    <td><a href="kubernetes/apps/nightscout/app/helmrelease.yaml">nightscout</a></td>
    <td><a href="kubernetes/apps/nightscout/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>node-feature-discovery</td>
    <td><a href="kubernetes/apps/node-feature-discovery/app/helmrelease.yaml">node-feature-discovery</a></td>
    <td></td>
  </tr>
  <tr>
    <td>nvidia-device-plugin</td>
    <td><a href="kubernetes/apps/nvidia-device-plugin/app/helmrelease.yaml">nvidia-device-plugin</a></td>
    <td></td>
  </tr>
  <tr>
    <td>obico</td>
    <td><a href="kubernetes/apps/obico/app/helmrelease.yaml">obico</a></td>
    <td><a href="kubernetes/apps/obico/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>overseerr</td>
    <td><a href="kubernetes/apps/overseerr/app/helmrelease.yaml">overseerr</a></td>
    <td><a href="kubernetes/apps/overseerr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>paperless-ngx</td>
    <td><a href="kubernetes/apps/paperless-ngx/app/helmrelease.yaml">paperless-ngx</a></td>
    <td><a href="kubernetes/apps/paperless-ngx/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/paperless-ngx/app/keydb.yaml">keydb</a>, <a href="kubernetes/apps/paperless-ngx/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>pictshare</td>
    <td><a href="kubernetes/apps/pictshare/app/helmrelease.yaml">pictshare</a></td>
    <td></td>
  </tr>
  <tr>
    <td>plex</td>
    <td><a href="kubernetes/apps/plex/app/helmrelease.yaml">plex</a></td>
    <td><a href="kubernetes/apps/plex/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>portfolio</td>
    <td><a href="kubernetes/apps/portfolio/app/helmrelease.yaml">portfolio</a></td>
    <td><a href="kubernetes/apps/portfolio/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="2">prometheus</td>
    <td><a href="kubernetes/apps/prometheus/app/helmrelease.yaml">kube-prometheus-stack</a></td>
    <td rowspan="2"><a href="kubernetes/apps/prometheus/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td><a href="kubernetes/apps/prometheus/app/nut-exporter.yaml">nut-exporter</a></td>
  </tr>
  <tr>
    <td>prowlarr</td>
    <td><a href="kubernetes/apps/prowlarr/app/helmrelease.yaml">prowlarr</a></td>
    <td><a href="kubernetes/apps/prowlarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>qbittorrent</td>
    <td><a href="kubernetes/apps/qbittorrent/app/helmrelease.yaml">qbittorrent</a></td>
    <td><a href="kubernetes/apps/qbittorrent/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>radarr</td>
    <td><a href="kubernetes/apps/radarr/app/helmrelease.yaml">radarr</a></td>
    <td><a href="kubernetes/apps/radarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>reflector</td>
    <td><a href="kubernetes/apps/reflector/app/helmrelease.yaml">reflector</a></td>
    <td></td>
  </tr>
  <tr>
    <td>relax-sounds</td>
    <td><a href="kubernetes/apps/relax-sounds/app/helmrelease.yaml">relax-sounds</a></td>
    <td><a href="kubernetes/apps/relax-sounds/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>reloader</td>
    <td><a href="kubernetes/apps/reloader/app/helmrelease.yaml">reloader</a></td>
    <td></td>
  </tr>
  <tr>
    <td>renovate</td>
    <td><a href="kubernetes/apps/renovate/app/helmrelease.yaml">renovate</a></td>
    <td></td>
  </tr>
  <tr>
    <td>scanservjs</td>
    <td><a href="kubernetes/apps/scanservjs/app/helmrelease.yaml">scanservjs</a></td>
    <td></td>
  </tr>
  <tr>
    <td>shlink</td>
    <td><a href="kubernetes/apps/shlink/app/helmrelease.yaml">shlink</a></td>
    <td><a href="kubernetes/apps/shlink/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/shlink/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>sonarr</td>
    <td><a href="kubernetes/apps/sonarr/app/helmrelease.yaml">sonarr</a></td>
    <td><a href="kubernetes/apps/sonarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>stable-diffusion</td>
    <td><a href="kubernetes/apps/stable-diffusion/app/helmrelease.yaml">stable-diffusion-webui</a></td>
    <td></td>
  </tr>
  <tr>
    <td>tandoor</td>
    <td><a href="kubernetes/apps/tandoor/app/helmrelease.yaml">tandoor</a></td>
    <td><a href="kubernetes/apps/tandoor/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/tandoor/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>tautulli</td>
    <td><a href="kubernetes/apps/tautulli/app/helmrelease.yaml">tautulli</a></td>
    <td><a href="kubernetes/apps/tautulli/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>transsmute</td>
    <td><a href="kubernetes/apps/transsmute/app/helmrelease.yaml">transsmute</a></td>
    <td></td>
  </tr>
  <tr>
    <td>uptime-kuma</td>
    <td><a href="kubernetes/apps/uptime-kuma/app/helmrelease.yaml">uptime-kuma</a></td>
    <td><a href="kubernetes/apps/uptime-kuma/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>vaultwarden</td>
    <td><a href="kubernetes/apps/vaultwarden/app/helmrelease.yaml">vaultwarden</a></td>
    <td><a href="kubernetes/apps/vaultwarden/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/vaultwarden/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>vikunja</td>
    <td><a href="kubernetes/apps/vikunja/app/helmrelease.yaml">vikunja</a></td>
    <td><a href="kubernetes/apps/vikunja/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="kubernetes/apps/vikunja/app/keydb.yaml">keydb</a>, <a href="kubernetes/apps/vikunja/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td>weave-gitops</td>
    <td><a href="kubernetes/apps/weave-gitops/app/helmrelease.yaml">weave-gitops</a></td>
    <td></td>
  </tr>
  <tr>
    <td>zigbee2mqtt</td>
    <td><a href="kubernetes/apps/zigbee2mqtt/app/helmrelease.yaml">zigbee2mqtt</a></td>
    <td><a href="kubernetes/apps/zigbee2mqtt/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td>zwave-js-ui</td>
    <td><a href="kubernetes/apps/zwave-js-ui/app/helmrelease.yaml">zwave-js-ui</a></td>
    <td><a href="kubernetes/apps/zwave-js-ui/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
</table>
<!-- End apps section -->
