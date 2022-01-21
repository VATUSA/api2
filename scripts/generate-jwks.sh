#!/bin/bash

fname=$1

if [[ -z "$fname" ]]; then
    fname="jwks.json"
fi	
echo "- Generating JWKS file: $fname"

echo "- Downloading jwkgen"
curl -L https://github.com/rakutentech/jwkgen/releases/download/v1.4.7/jwkgen-linux-x86_64.tar.gz -o /tmp/jwkgen-linux-x86_64.tar.gz &>/dev/null
if [[ $? -ne 0 ]]; then
    echo "Failed to download jwkgen"
    exit 1
fi

echo "- Extracting jwkgen"
tar -xvzf /tmp/jwkgen-linux-x86_64.tar.gz -C /tmp &>/dev/null
if [[ $? -ne 0 ]]; then
    echo "Failed to extract jwkgen"
    exit 1
fi

echo "- Generating jwks"
key1=$(/tmp/jwkgen --jwk | jq '.kid=1')
key2=$(/tmp/jwkgen -b 2048 rsa --jwk | jq '.kid=2')

echo "- Writing jwks to $fname"
cat > $fname << EOF
{
  "keys": [
    $key1,
    $key2
  ]
}
EOF
jq . $fname > $fname.tmp
mv $fname.tmp $fname

echo "- Cleaning up"
rm -rf /tmp/jwkgen-linux-x86_64.tar.gz /tmp/jwkgen-linux-x86_64

echo "- Done"