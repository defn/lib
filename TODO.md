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
        - store with _${stack} suffix
    - Enable MFA

Configure organization:
    - Visit https://us-east-1.console.aws.amazon.com/singlesignon/identity/home
        - Select the correct region
        - Enable SSO, which creates the organization
    - Create the Administrators group

Create Terraform cloud workspace, named after the org
    - Configure workspace for local execution mode
