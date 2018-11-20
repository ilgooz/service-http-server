## service-website [![CircleCI](https://img.shields.io/circleci/project/github/ilgooz/service-website.svg)](https://github.com/ilgooz/service-website) [![codecov](https://codecov.io/gh/ilgooz/service-website/branch/master/graph/badge.svg)](https://codecov.io/gh/ilgooz/service-website)
A MESG service to serve website content.

serve your website over http

```bash
mesg-core service deploy https://github.com/ilgooz/service-website
```

# Events

## request

Event key: `request`

This event is emited on every page request.

| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **host** | `String` | Host of HTTP request. |
| **ip** | `String` | IP address of HTTP request. |
| **method** | `String` | Method of HTTP request. |
| **path** | `String` | Path of HTTP request. |
| **sessionID** | `String` | Session ID of corresponding page request |


# Tasks

## completeSession

Task key: `completeSession`

This task should be called to response a pending page requests.

### Inputs

| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **code** | `Number` | Status code of HTTP response. |
| **content** | `String` | Content of HTTP response. |
| **error** | `Object` | Error should be used to response with pre-existing errors |
| **mimeType** | `String` | MIME type of HTTP response. |
| **sessionID** | `String` | Session ID of corresponding page request |


### Outputs

##### error

Output key: `error`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **message** | `String` |  |

##### success

Output key: `success`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **elapsedTime** | `String` | Elapsed time in nanoseconds for page request to complete |
| **sessionID** | `String` | Session ID of corresponding page request |




