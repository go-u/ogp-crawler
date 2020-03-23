#!/bin/bash

# BaseDir=".secrets"
BaseDir="."
PassPhraseFile="${BaseDir}/masterkey/passphrase.txt"
Categories=(
"cloudsql"
"twitter"
)

pwd

for CategoryDir in "${Categories[@]}" ;do
  echo "${category}"
  PlaneDir="${BaseDir}/${CategoryDir}/plane"
  EncDir="${BaseDir}/${CategoryDir}/encrypted"
  echo "${PlaneDir}"
  for KeyPath in "${PlaneDir}"/*; do
    echo "${KeyPath}"
    KeyName=$(basename "${KeyPath}")
    EncryptedKeyPath=${EncDir}/${KeyName}.gpg
    echo "${EncryptedKeyPath}"
    gpg --output "${EncryptedKeyPath}" --batch --yes --passphrase-file "${PassPhraseFile}" --symmetric --cipher-algo AES256 "${KeyPath}"
  done
done
