#!/bin/bash

http POST :8080/accounts id="123" customerName="Joe Doe" fiscalDocument="11111111111"
http POST :8080/accounts id="456" customerName="Jane Smith" fiscalDocument="22222222222"

http POST :8080/accounts/123/deposits amount=1000
http POST :8080/accounts/123/withdraws amount=200
http POST :8080/accounts/123/transfers amount=100 destinationId=456

http GET :8080/accounts/123
http GET :8080/accounts/456