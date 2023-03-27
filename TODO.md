# TODO list

## build/

- Build custom cloud image for Yandex Cloud
    - Current image does not read cloud-init configuration from Yandex
    - sshd fails to start. Network problem?
- Create CI pipeline with GitHub actions to push updated images to S3 (on schedule)
- Figure out Docker-in-Docker CI problems


## deploy/

- Automate S3 bucket management
    - Bucket is persistent while the rest of the infra is not
    - Use separate terraform module? Or better define it in the same module,
      just don't use scaling with `count=0`
- Package into a Docker container
- Deploy fleet manager to home server
- Look into generating a pre-signed URL for VM image on fleet manager.
  That would allow to make the S3 bucket private.
  Be careful: changing URL (GET params) would trigger tf to rebuild the image,
  which in turn could(?) trigger VM rebuilds.


## scale/

- Rewrite in Go: calculate scaling actions and populate tfvars
- Optional: use external data source in terraform to call scaler automatically