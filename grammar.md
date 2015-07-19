Grammar
=======

# Parsing Expressions

  see src/sysla.peg

# Examples

## Forms

### Form I

  QUALIFIER SUBJECT ATTRIBUTE PREDICATE VALUE

#### Examples

  Qualitifer = one ('the') | no | all

  the X has Y matching V1
  the X has Y not matching V2
  the X has Y equals V3

  no I has J matching K

  all R has S matching T

#### Example phrases

    The interface "eth0" has "netmask" matching /\d+/.
    The interface "eth1" has "mac_address" equals "12:34:56:78:9A:BC".
    The device "/dev/sd1" have "free_space_bytes" more than 30000000.
    All interfaces have "address" matching  /.*/.

### Form II

  QUALIFIER SUBJECT ATTRIBUTE PREDICATE ('exist' / 'exists')

#### Examples

  The process with "name" matching "/tmp/proc" exists.
  The process with "name" not matching "/tmp/foo" exists.
  No process with "name" matching "/tmp/foo" exists.
