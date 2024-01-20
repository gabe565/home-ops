# Kubernetes Apps

<!-- Begin apps table -->
<table>
  <tr>
    <th>Namespace</th>
    <th>Name</th>
    <th>Supporting Services</th>
  </tr>
  <tr>
    <td rowspan="2">adguard-home</td>
    <td><a href="../../kubernetes/apps/adguard-home/app/helmrelease.yaml">adguard-home</a></td>
    <td rowspan="2"><a href="../../kubernetes/apps/adguard-home/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/adguard-home/external-dns/helmrelease.yaml">adguard-home-external-dns</a></td>
  </tr>
  <tr>
    <td rowspan="1">ascii-movie</td>
    <td><a href="../../kubernetes/apps/ascii-movie/app/helmrelease.yaml">ascii-movie</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">authentik</td>
    <td><a href="../../kubernetes/apps/authentik/app/helmrelease.yaml">authentik</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/authentik/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/authentik/postgresql.yaml">postgresql</a>, <a href="../../kubernetes/apps/authentik/redis/helmrelease.yaml">redis</a></td>
  </tr>
  <tr>
    <td rowspan="1">bookstack</td>
    <td><a href="../../kubernetes/apps/bookstack/app/helmrelease.yaml">bookstack</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/bookstack/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">castsponsorskip</td>
    <td><a href="../../kubernetes/apps/castsponsorskip/app/helmrelease.yaml">castsponsorskip</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">cert-manager</td>
    <td><a href="../../kubernetes/apps/cert-manager/app/helmrelease.yaml">cert-manager</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">change-detection</td>
    <td><a href="../../kubernetes/apps/change-detection/app/helmrelease.yaml">change-detection</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">charm</td>
    <td><a href="../../kubernetes/apps/charm/app/helmrelease.yaml">charm</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">chatgpt</td>
    <td><a href="../../kubernetes/apps/chatgpt/app/helmrelease.yaml">chatgpt</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">cnpg-system</td>
    <td><a href="../../kubernetes/apps/cnpg/app/helmrelease.yaml">cnpg</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">domain-watch</td>
    <td><a href="../../kubernetes/apps/domain-watch/app/helmrelease.yaml">domain-watch</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">esphome</td>
    <td><a href="../../kubernetes/apps/esphome/app/helmrelease.yaml">esphome</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/esphome/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="2">external-dns</td>
    <td><a href="../../kubernetes/apps/external-dns/dynamic-ip/helmrelease.yaml">dynamic-ip</a></td>
    <td rowspan="2"></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/external-dns/app/helmrelease.yaml">external-dns</a></td>
  </tr>
  <tr>
    <td rowspan="1">generic-device-plugin</td>
    <td><a href="../../kubernetes/apps/generic-device-plugin/app/helmrelease.yaml">generic-device-plugin</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">gitea</td>
    <td><a href="../../kubernetes/apps/gitea/app/helmrelease.yaml">gitea</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/gitea/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/gitea/postgresql.yaml">postgresql</a>, <a href="../../kubernetes/apps/gitea/redis/helmrelease.yaml">redis</a></td>
  </tr>
  <tr>
    <td rowspan="1">hammond</td>
    <td><a href="../../kubernetes/apps/hammond/app/helmrelease.yaml">hammond</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">headscale</td>
    <td><a href="../../kubernetes/apps/headscale/app/helmrelease.yaml">headscale</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/headscale/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/headscale/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">healthchecks</td>
    <td><a href="../../kubernetes/apps/healthchecks/app/helmrelease.yaml">healthchecks</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/healthchecks/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/healthchecks/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="2">home-assistant</td>
    <td><a href="../../kubernetes/apps/home-assistant/app/helmrelease.yaml">home-assistant</a></td>
    <td rowspan="2"><a href="../../kubernetes/apps/home-assistant/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/home-assistant/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/home-assistant/piper/helmrelease.yaml">piper</a></td>
  </tr>
  <tr>
    <td rowspan="1">homepage</td>
    <td><a href="../../kubernetes/apps/homepage/app/helmrelease.yaml">homepage</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">immich</td>
    <td><a href="../../kubernetes/apps/immich/app/helmrelease.yaml">immich</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/immich/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/immich/postgresql.yaml">postgresql</a>, <a href="../../kubernetes/apps/immich/redis/helmrelease.yaml">redis</a></td>
  </tr>
  <tr>
    <td rowspan="1">ingress-nginx</td>
    <td><a href="../../kubernetes/apps/ingress-nginx/app/helmrelease.yaml">ingress-nginx</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">kromgo</td>
    <td><a href="../../kubernetes/apps/prometheus/kromgo/helmrelease.yaml">kromgo</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="3">kube-system</td>
    <td><a href="../../kubernetes/apps/cilium/app/helmrelease.yaml">cilium</a></td>
    <td rowspan="3"></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/kube-vip/app/helmrelease.yaml">kube-vip</a></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/metrics-server/app/helmrelease.yaml">metrics-server</a></td>
  </tr>
  <tr>
    <td rowspan="1">lidarr</td>
    <td><a href="../../kubernetes/apps/lidarr/app/helmrelease.yaml">lidarr</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">linkding</td>
    <td><a href="../../kubernetes/apps/linkding/app/helmrelease.yaml">linkding</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/linkding/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/linkding/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">longhorn-system</td>
    <td><a href="../../kubernetes/apps/longhorn/app/helmrelease.yaml">longhorn</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">matrimony</td>
    <td><a href="../../kubernetes/apps/matrimony/app/helmrelease.yaml">matrimony</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/matrimony/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">memos</td>
    <td><a href="../../kubernetes/apps/memos/app/helmrelease.yaml">memos</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/memos/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/memos/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">metallb-system</td>
    <td><a href="../../kubernetes/apps/metallb/app/helmrelease.yaml">metallb</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">miniflux</td>
    <td><a href="../../kubernetes/apps/miniflux/app/helmrelease.yaml">miniflux</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/miniflux/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/miniflux/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">minio</td>
    <td><a href="../../kubernetes/apps/minio/app/helmrelease.yaml">minio</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/minio/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">monica</td>
    <td><a href="../../kubernetes/apps/monica/app/helmrelease.yaml">monica</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/monica/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">mosquitto</td>
    <td><a href="../../kubernetes/apps/mosquitto/app/helmrelease.yaml">mosquitto</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">nextcloud</td>
    <td><a href="../../kubernetes/apps/nextcloud/app/helmrelease.yaml">nextcloud</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/nextcloud/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/nextcloud/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">nfs-provisioner</td>
    <td><a href="../../kubernetes/apps/nfs-subdir-external-provisioner/app/helmrelease.yaml">nfs-subdir-external-provisioner</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">nightscout</td>
    <td><a href="../../kubernetes/apps/nightscout/app/helmrelease.yaml">nightscout</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/nightscout/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">node-feature-discovery</td>
    <td><a href="../../kubernetes/apps/node-feature-discovery/app/helmrelease.yaml">node-feature-discovery</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">nvidia-device-plugin</td>
    <td><a href="../../kubernetes/apps/nvidia-device-plugin/app/helmrelease.yaml">nvidia-device-plugin</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">obico</td>
    <td><a href="../../kubernetes/apps/obico/app/helmrelease.yaml">obico</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/obico/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">overseerr</td>
    <td><a href="../../kubernetes/apps/overseerr/app/helmrelease.yaml">overseerr</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/overseerr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">paperless-ngx</td>
    <td><a href="../../kubernetes/apps/paperless-ngx/app/helmrelease.yaml">paperless-ngx</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/paperless-ngx/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/paperless-ngx/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">pictshare</td>
    <td><a href="../../kubernetes/apps/pictshare/app/helmrelease.yaml">pictshare</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">plex</td>
    <td><a href="../../kubernetes/apps/plex/app/helmrelease.yaml">plex</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/plex/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">portfolio</td>
    <td><a href="../../kubernetes/apps/portfolio/app/helmrelease.yaml">portfolio</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/portfolio/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="2">prometheus</td>
    <td><a href="../../kubernetes/apps/prometheus/app/helmrelease.yaml">kube-prometheus-stack</a></td>
    <td rowspan="2"><a href="../../kubernetes/apps/prometheus/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td><a href="../../kubernetes/apps/prometheus/app/nut-exporter.yaml">nut-exporter</a></td>
  </tr>
  <tr>
    <td rowspan="1">prowlarr</td>
    <td><a href="../../kubernetes/apps/prowlarr/app/helmrelease.yaml">prowlarr</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/prowlarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">qbittorrent</td>
    <td><a href="../../kubernetes/apps/qbittorrent/app/helmrelease.yaml">qbittorrent</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/qbittorrent/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">radarr</td>
    <td><a href="../../kubernetes/apps/radarr/app/helmrelease.yaml">radarr</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/radarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">reflector</td>
    <td><a href="../../kubernetes/apps/reflector/app/helmrelease.yaml">reflector</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">relax-sounds</td>
    <td><a href="../../kubernetes/apps/relax-sounds/app/helmrelease.yaml">relax-sounds</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/relax-sounds/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">renovate</td>
    <td><a href="../../kubernetes/apps/renovate/app/helmrelease.yaml">renovate</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">scanservjs</td>
    <td><a href="../../kubernetes/apps/scanservjs/app/helmrelease.yaml">scanservjs</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">searxng</td>
    <td><a href="../../kubernetes/apps/searxng/app/helmrelease.yaml">searxng</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/searxng/redis/helmrelease.yaml">redis</a></td>
  </tr>
  <tr>
    <td rowspan="1">shlink</td>
    <td><a href="../../kubernetes/apps/shlink/app/helmrelease.yaml">shlink</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/shlink/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/shlink/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">sonarr</td>
    <td><a href="../../kubernetes/apps/sonarr/app/helmrelease.yaml">sonarr</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/sonarr/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">stable-diffusion</td>
    <td><a href="../../kubernetes/apps/stable-diffusion/app/helmrelease.yaml">stable-diffusion-webui</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">tandoor</td>
    <td><a href="../../kubernetes/apps/tandoor/app/helmrelease.yaml">tandoor</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/tandoor/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/tandoor/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">tautulli</td>
    <td><a href="../../kubernetes/apps/tautulli/app/helmrelease.yaml">tautulli</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/tautulli/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">transsmute</td>
    <td><a href="../../kubernetes/apps/transsmute/app/helmrelease.yaml">transsmute</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">uptime-kuma</td>
    <td><a href="../../kubernetes/apps/uptime-kuma/app/helmrelease.yaml">uptime-kuma</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/uptime-kuma/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">vaultwarden</td>
    <td><a href="../../kubernetes/apps/vaultwarden/app/helmrelease.yaml">vaultwarden</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/vaultwarden/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/vaultwarden/postgresql.yaml">postgresql</a></td>
  </tr>
  <tr>
    <td rowspan="1">vikunja</td>
    <td><a href="../../kubernetes/apps/vikunja/app/helmrelease.yaml">vikunja</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/vikunja/borgmatic/helmrelease.yaml">borgmatic</a>, <a href="../../kubernetes/apps/vikunja/postgresql.yaml">postgresql</a>, <a href="../../kubernetes/apps/vikunja/redis/helmrelease.yaml">redis</a></td>
  </tr>
  <tr>
    <td rowspan="1">weave-gitops</td>
    <td><a href="../../kubernetes/apps/weave-gitops/app/helmrelease.yaml">weave-gitops</a></td>
    <td rowspan="1"></td>
  </tr>
  <tr>
    <td rowspan="1">zigbee2mqtt</td>
    <td><a href="../../kubernetes/apps/zigbee2mqtt/app/helmrelease.yaml">zigbee2mqtt</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/zigbee2mqtt/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
  <tr>
    <td rowspan="1">zwave-js-ui</td>
    <td><a href="../../kubernetes/apps/zwave-js-ui/app/helmrelease.yaml">zwave-js-ui</a></td>
    <td rowspan="1"><a href="../../kubernetes/apps/zwave-js-ui/borgmatic/helmrelease.yaml">borgmatic</a></td>
  </tr>
</table>
<!-- End apps table -->
