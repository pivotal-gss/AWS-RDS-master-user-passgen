# AWS Relational Database Service master user passgen

This tool regenerates the original masterusername password for an AWS Relational
Database Service PostgreSQL Instance originally created by the PCF Service Broker
for AWS provided a service instance guid and master salt key.

Please see Pivotal knowledge base on this tool here:
https://discuss.pivotal.io/hc/en-us/articles/360001356494

```
Usage: passGen -i [identity] -s [salt]
```

Standard security recommendations apply to distribution of the generated
password.

This tool is provided as a general service and is not under any official
supported capacity. There is no implied or guaranteed warranty or statement of
support.

Released under MIT license,	copyright 2018 Tyler Ramer
