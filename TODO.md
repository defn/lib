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
    - Generate access keys
        - AWS_ACCESS_KEY_ID
        - AWS_SECRET_ACCESS_KEY