# http-server

serve content over http

# Contents

- [Installation](#Installation)
- [Definitions](#Definitions)
  - [Events](#Events)
    - [request](#request)
  - [Tasks](#Tasks)
    - [breakCache](#breakcache)
    - [cache](#cache)
    - [completeSession](#completesession)

# Installation

## MESG Core

This service requires [MESG Core](https://github.com/mesg-foundation/core) to be installed first.

You can install MESG Core by running the following command or [follow the installation guide](https://docs.mesg.com/guide/start-here/installation.html).

```bash
bash <(curl -fsSL https://mesg.com/install)
```

## Service

To deploy this service, run the following command:
```bash
mesg-core service deploy https://github.com/ilgooz/service-http-server
```

# Definitions

# Events

## request

Event key: `request`

This event is emited on every page request.

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **body** | `body` | `Any` | Body of HTTP request. |
| **host** | `host` | `String` | Host of HTTP request. |
| **ip** | `ip` | `String` | IP address of HTTP request. |
| **method** | `method` | `String` | Method of HTTP request. |
| **path** | `path` | `String` | Path of HTTP request. |
| **qs** | `qs` | `Any` | Query string of HTTP request. |
| **sessionID** | `sessionID` | `String` | Session ID of corresponding page request |

# Tasks

## breakCache

Task key: `breakCache`

Break cache for http request

### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **method** | `method` | `String` | Method of HTTP request to break cache. |
| **path** | `path` | `String` | Path of HTTP request to break cache. |

### Outputs

#### error

Output key: `error`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` |  |

#### success

Output key: `success`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` |  |


## cache

Task key: `cache`

Cache an http request with pre created response

### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **code** | `code` | `Number` | **`optional`** Status code of HTTP response |
| **content** | `content` | `String` | **`optional`** Content of HTTP response |
| **method** | `method` | `String` | Method of HTTP request |
| **mimeType** | `mimeType` | `String` | **`optional`** MIME type of HTTP response |
| **path** | `path` | `String` | Path of HTTP request |

### Outputs

#### error

Output key: `error`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` |  |

#### success

Output key: `success`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` |  |


## completeSession

Task key: `completeSession`

This task should be called to response a pending page requests.

### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **code** | `code` | `Number` | **`optional`** Status code of HTTP response. |
| **content** | `content` | `String` | **`optional`** Content of HTTP response. |
| **error** | `error` | `Object` | **`optional`** Error should be used to response with pre-existing errors |
| **mimeType** | `mimeType` | `String` | **`optional`** MIME type of HTTP response. |
| **sessionID** | `sessionID` | `String` | Session ID of corresponding page request |

### Outputs

#### error

Output key: `error`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` |  |

#### success

Output key: `success`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **elapsedTime** | `elapsedTime` | `Number` | Elapsed time in nanoseconds for page request to complete |
| **sessionID** | `sessionID` | `String` | Session ID of corresponding page request |


