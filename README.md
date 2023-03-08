# Test parser

Simple utility for parsing E2E test results (Wdio) producing prometheus metrics.
Output from this utility can be pushed to prometheus push gateway.

## example

```bash
go run cmd/metrics/main.go -file ./test.json | curl --data-binary @- http://localhost:9091/metrics/job/some_job
```

## test.json

```json
{
  "start":"2023-02-28T14:40:26.065Z",
  "end":"2023-02-28T14:40:34.144Z",
  "capabilities":
    {
      "acceptInsecureCerts":false,
      "browserName":"chrome",
      "browserVersion":"110.0.5481.178",
      "chrome":
        {
          "chromedriverVersion":"110.0.5481.77 (65ed616c6e8ee3fe0ad64fe83796c020644d42af-refs/branch-heads/5481@{#839})"
        },
      "goog:chromeOptions":
        {
          "debuggerAddress":"localhost:15447"
        },
      "networkConnectionEnabled":false,
      "pageLoadStrategy":"normal",
      "platformName":"windows","proxy":{},
      "setWindowRect":true,
      "strictFileInteractability":false,
      "timeouts":
        {
          "implicit":0,
          "pageLoad":300000,
          "script":30000
        },
      "unhandledPromptBehavior":"dismiss and notify",
      "webauthn:extension:credBlob":true,
      "webauthn:extension:largeBlob":true,
      "webauthn:virtualAuthenticators":true,
      "sessionId":"fca7f7f3a6a08c33e67e8fd3deebc85a"
    },
  "host":"localhost",
  "port":9515,
  "baseUrl":"https://test.something.cz",
  "framework":"mocha",
  "suites":
    [
      {
        "name":"Login Page",
        "duration":8025,
        "start":"2023-02-28T14:40:26.069Z",
        "end":"2023-02-28T14:40:34.094Z",
        "sessionId":"fca7f7f3a6a08c33e67e8fd3deebc85a",
        "tests":
          [
            {
              "name":"Should log in",
              "start":"2023-02-28T14:40:26.069Z",
              "end":"2023-02-28T14:40:29.907Z",
              "duration":3838,
              "state":"passed"
            },
            {
              "name":"Should log in via API",
              "start":"2023-02-28T14:40:29.908Z",
              "duration":0,
              "state":"skipped"
            },
            {
              "name":"Shouldn't log in - wrong credentials",
              "start":"2023-02-28T14:40:29.908Z",
              "end":"2023-02-28T14:40:31.811Z",
              "duration":1903,
              "state":"passed"
            },
            {
              "name":"Shouldn't log in - No Email",
              "start":"2023-02-28T14:40:31.811Z",
              "end":"2023-02-28T14:40:32.906Z",
              "duration":1095,
              "state":"passed"
            },
            {
              "name":"Shouldn't log in - No Password",
              "start":"2023-02-28T14:40:32.906Z",
              "end":"2023-02-28T14:40:34.094Z",
              "duration":1188,
              "state":"passed"
            }
          ],
        "hooks":
          [
            {
              "start":"2023-02-28T14:40:26.069Z",
              "end":"2023-02-28T14:40:28.553Z",
              "duration":2484,
              "title":"\"before each\" hook for Login Page",
              "associatedSuite":"Login Page",
              "associatedTest":"Should log in",
              "state":"passed"
            },
            {
              "start":"2023-02-28T14:40:29.908Z",
              "end":"2023-02-28T14:40:30.600Z",
              "duration":692,
              "title":"\"before each\" hook for Login Page",
              "associatedSuite":"Login Page",
              "associatedTest":"Shouldn't log in - wrong credentials",
              "state":"passed"
            },
            {
              "start":"2023-02-28T14:40:31.811Z",
              "end":"2023-02-28T14:40:32.351Z",
              "duration":540,
              "title":"\"before each\" hook for Login Page",
              "associatedSuite":"Login Page",
              "associatedTest":"Shouldn't log in - No Email",
              "state":"passed"
            },
            {
              "start":"2023-02-28T14:40:32.906Z",
              "end":"2023-02-28T14:40:33.527Z",
              "duration":621,
              "title":"\"before each\" hook for Login Page",
              "associatedSuite":"Login Page",
              "associatedTest":"Shouldn't log in - No Password",
              "state":"passed"
            }
          ]
      }
    ],
  "specs":["file:///C:/git/something/tests-gui/test/specs/newBank/login.test.js"],
  "state":
    {
      "passed":4,
      "failed":0,
      "skipped":1
    }
}
```