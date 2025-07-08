Lookup-dns-txt
==============
This repository is to create a simple tool to query dns TXT record.

It should be used with hesiod, since it looks for `/etc/hesiod.conf` to compose the dns query.

Usage
=============
This tool can be used in openssh server configuration, add below line in the `sshd_config` file:
    
    AuthorizedKeysCommand /etc/ssh/lookup-dns-txt %u

The `sshd` will call the tool for user's authorized key which is stored in the DNS as txt record.
