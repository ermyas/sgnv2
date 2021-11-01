# Create AWS NLB

## for test
```
aws elbv2 create-load-balancer --type network --ip-address-type ipv4 --subnets subnet-0b843e588229579f5  --name sgnv2-NLB-test
aws elbv2 create-target-group --name sgnv2-TG-test --protocol TCP --port 80 --vpc-id vpc-04554b46166d7c9b2
aws elbv2 register-targets --target-group-arn arn:aws:elasticloadbalancing:us-west-2:356032637240:targetgroup/sgnv2-TG-test/b5c25e89ee7001f9 --targets Id=i-059a577bef705a4f0
```
in register-targests, target-group-arn value is from create-target-group response json. targets Id is the EC2 instance ID that has nginx and gateway

last step is to let NLB terminate TLS using certificate. CertificateArn is the certificate in AWS Certificate Manager that matches desired domain name, celer.network in this case

```
aws elbv2 create-listener \
    --load-balancer-arn arn:aws:elasticloadbalancing:us-west-2:356032637240:loadbalancer/net/sgnv2-NLB-test/2f211b88ca579568 \
    --protocol TLS --port 443 --certificates CertificateArn=arn:aws:acm:us-west-2:356032637240:certificate/325581e7-c867-41c4-882d-008e1090b7cb \
    --ssl-policy ELBSecurityPolicy-2016-08 \
    --default-actions Type=forward,TargetGroupArn=arn:aws:elasticloadbalancing:us-west-2:356032637240:targetgroup/sgnv2-TG-test/b5c25e89ee7001f9
```

update celer.network DSN, add CNAME entry cbridge-v2-test points to sgnv2-nlb-test-2f211b88ca579568.elb.us-west-2.amazonaws.com (DNSName in create-load-balancer response)
