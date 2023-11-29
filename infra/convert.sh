#!/bin/bash


for language in java python go nodejs dotnet
do
	pulumi convert --generate-only --language ${language} --out ../dest-${language} --cwd source-yaml --non-interactive
done
exit 0