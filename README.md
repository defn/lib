Create an AWS account
    - Visit https://aws.amazon.com/free
        - root email: aws-ORG@DOMAIN

Configure root account:
    - Enable IAM access to billing:
        - https://us-east-1.console.aws.amazon.com/billing/home?region=us-east-1#/account
    - Enable MFA:
        - https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/home

Create Administrator IAM user:
    - Visit https://us-east-1.console.aws.amazon.com/iam/home#/users$new?step=details
    - Add roles:
        - AdministratorAccess
        - AWSSSOMasterAccountAdministrator
    - Generate access keys
        - AWS_ACCESS_KEY_ID
        - AWS_SECRET_ACCESS_KEY
        - Add to pass
            - _${stack} suffix
        - Add to aws-vault
            - aws-vault add ORG
        - Enable MFA

Configure organization:
    - Visit https://us-east-1.console.aws.amazon.com/singlesignon/identity/home
        - Select the correct region
        - Enable SSO, which creates the organization
    - Create the Administrators group
    - Add a user to the group
    - Record the SSO url

Create Terraform cloud workspace, named after the org
    - Configure workspace for local execution mode

Add AWS_ACCESS_KEY_ID_${org},AWS_SECRET_ACCESS_KEY_${org} to lib/Earthfile

Add stack to src/defn/cli.py

Then cdktf initial accounts
    - make synth stack=${org}
    - make import stack=${org}
    - make plan stack=${org}
    - make apply stack=${org}

Generate .aws/config
    - export region=us-west-1 sso_region=us-west-2 url=https://.../start name=curl
    - bin/awsconfig >> ~/.aws/config


kubectl --context pod patch -n vc1 service kourier-internal-x-kourier-system-x-vc1 -p '{"metadata":{"annotations":{"traefik.ingress.kubernetes.io/service.serversscheme":"h2c"}}}'

v write pki/roles/gyre.defn.dev allowed_domains=gyre.defn.dev,demo.svc.cluster.local,sslip.io,remocal.net allow_subdomains=true max_ttl=120h
v write pki/issue/gyre.defn.dev common_name="remocal.net" ip_sans="169.254.32.1" alt_names="hello.demo.svc.cluster.local,169-254-32-1.sslip.io" tl=1h -format=json | jq .data > meh.json
k --context pod patch -n traefik secret default-certificate --type='json' -p='[{"op" : "replace" ,"path" : "/data/tls.key" ,"value" : "'$(cat meh.json | jq -r '.private_key | @base64')'"}]'
k --context pod patch -n traefik secret default-certificate --type='json' -p='[{"op" : "replace" ,"path" : "/data/tls.crt" ,"value" : "'$(cat meh.json | jq -r '.certificate | @base64')'"}]'
