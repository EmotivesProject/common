# Common
## Introduction
This is a repo that holds multiple different go modules that are shared between service. This is useful to reduce code duplication across multiple services.

A lot of common modules use a configuration of some kind, this way the modules can be shared and each service that requires a small change can still share it.

## Changes
If you want to make a change in a specific go module, these are the steps.
1. Make changes in the go module.
2. Push the changes to the main branch.
3. Get the latest config hash, for example
   ```
   git log
   commit d9a3cbff66fe217a2964e77eb154b05c89f68a55 (HEAD, origin/main)
	Author: Tom Bowyer <tom4310@gmail.com>
	Date:   Sun Oct 24 15:08:25 2021 +1000
    Fix null
   ```
	That would be d9a3cbff66fe217a2964e77eb154b05c89f68a55
4. Go back into the repo you wanted the change for. e.g. uacl/chatter
5. Install the latest change
   ```
   go get github.com/EmotivesProject/common/logger@d9a3cbff66fe217a2964e77eb154b05c89f68a55
   ```
## Creating a new common module
If you wanted to create a brand new common module, these are the steps
1. Create a folder
2. Initialize the go module with the github name
   ```
   go mod init github.com/EmotivesProject/common/new
   ```
3. Make changes and push to the main branch
4. Optional. Install the new module in the relevant service