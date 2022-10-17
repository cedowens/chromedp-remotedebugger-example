# chromedp-remotedebugger-example

A super simple example of how to use chromedp to run Chrome headless with the remote debugger port programmatically (is still a wrapper around the Chrome binary).

## Steps
1. edit the main.go file and add the username to the Chrome path
2. > go mod init example/chrome-remotedebugger-example
3. > go build
4. run the binary

***Note: I added a 1 minute chromedp sleep so that the remote debugger runs long enough for the remote debugging port to be interacted with. You can adjust that time slot upward as needed.
