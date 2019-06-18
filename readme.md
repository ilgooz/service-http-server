# http-server (ID: http-server)

serve content over http

## Contents

- [Installation](#Installation)
  - [MESG Engine](#MESG-Core)
  - [Deploy the Service](#Service)
- [Definitions](#Definitions)
  - [Events](#Events)
    - [request](#request)
  - [Tasks](#Tasks)
    - [completeSession](#completeSession)
    - [cache](#cache)
    - [breakCache](#breakCache)

## Installation

### MESG Engine

This service requires [MESG Engine](https://github.com/mesg-foundation/core) to be installed first.

You can install MESG Engine by running the following command or [follow the installation guide](https://docs.mesg.com/guide/start-here/installation.html).

```bash
bash <(curl -fsSL https://mesg.com/install)
```

### Deploy the Service

To deploy this service, go to [this service page](https://marketplace.mesg.com/services/http-server) on the [MESG Marketplace](https://marketplace.mesg.com) and click the button "get/buy this service".

## Definitions

### Events

#### request

Event key: `request`

This event is emited on every page request.

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **sessionID** | `sessionID` | `String` | Session ID of corresponding page request |
| **path** | `path` | `String` | Path of HTTP request. |
| **method** | `method` | `String` | Method of HTTP request. |
| **qs** | `qs` | `Any` | Query string of HTTP request. |
| **body** | `body` | `Any` | Body of HTTP request. |
| **host** | `host` | `String` | Host of HTTP request. |
| **ip** | `ip` | `String` | IP address of HTTP request. |

### Tasks

#### completeSession

Task key: `completeSession`

This task should be called to response a pending page requests.

##### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **sessionID** | `sessionID` | `String` | Session ID of corresponding page request |
| **error** | `error` | `Object` | **`optional`** Error should be used to response with pre-existing errors |
| **code** | `code` | `Number` | **`optional`** Status code of HTTP response. |
| **mimeType** | `mimeType` | `String` | **`optional`** MIME type of HTTP response. |
| **content** | `content` | `String` | **`optional`** Content of HTTP response. |
| **cache** | `cache` | `Boolean` | **`optional`** Optionally cache this response for same requests. |
  
##### Outputs

###### sessionID

Output key: `sessionID`

Session ID of corresponding page request

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |

###### elapsedTime

Output key: `elapsedTime`

Elapsed time in nanoseconds for page request to complete

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |

#### cache

Task key: `cache`

Cache an http request with pre created response

##### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **method** | `method` | `String` | Method of HTTP request |
| **path** | `path` | `String` | Path of HTTP request |
| **code** | `code` | `Number` | **`optional`** Status code of HTTP response |
| **mimeType** | `mimeType` | `String` | **`optional`** MIME type of HTTP response |
| **content** | `content` | `String` | **`optional`** Content of HTTP response |
  
##### Outputs

###### message

Output key: `message`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |

#### breakCache

Task key: `breakCache`

Break cache for http request

##### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **path** | `path` | `String` | Path of HTTP request to break cache. |
| **method** | `method` | `String` | Method of HTTP request to break cache. |
  
##### Outputs

###### message

Output key: `message`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |


