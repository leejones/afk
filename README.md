# afk

A CLI to update your Slack status.

## Setup

### Install

1. Download the latest version from the Releases page on GitHub.
2. Unpack the file and add it to your PATH (e.g. to `/usr/local/bin`)

### Setup Slack API key

1. Go to your [Slack apps](https://api.slack.com/apps) page
1. Click **Create New App**
1. Click **From scratch**
1. Enter a name (e.g. afk-tool) and select the relevant workspace
1. Click **OAuth & Permissions** in the left sidebar
1. Under **User Token Scopes** add the following scopes
    1. `dnd:read`
    1. `dnd:write`
    1. `users.profile:read`
    1. `users.profile:write`
1. At the top of the page, click **Install to `your_workspace_name`** and go through the OAuth flow
1. From the top of the page, copy the **OAuth Access Token**
1. On your computer, create a new file at `~/.afk.yml`
1. Set limited permissions so it's not viewable by others: `chmod 600 ~/.afk.yml`
1. Edit the file and add the OAuth Access Token you copied earlier in the following format:

```
slackToken: TODO-ADD-YOUR-TOKEN-HERE
```

## Usage

```
afk --duration 1h --emoji :sandwich: --message "Lunch!"
```

For more details, run `afk --help`.
