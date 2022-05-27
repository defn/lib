Create an AWS account at aws.amazon.com/free:
    - root email: aws-ORG@DOMAIN
    - account name: ORG

Configure root account:
    - enable IAM access to billing: 
        - https://us-east-1.console.aws.amazon.com/billing/home?region=us-east-1#/account
    - enable MFA:
        - https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/home

Create Administrator IAM user:
    - enable MFA
    - add roles:
        - AdministratorAccess
        - AWSSSOMasterAccountAdministrator
    - Generate access keys
        - AWS_ACCESS_KEY_ID
        - AWS_SECRET_ACCESS_KEY

Configure organization:
    - Enable SSO in the right region, which creates the organization
    - Create the Administrators group
        - https://us-west-2.console.aws.amazon.com/singlesignon/identity/home?region=us-west-2#!/groups
