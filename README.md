# Vercel Account Manager (VAM)

Vercel Account Manager (VAM) is a CLI tool built to easily manage multiple Vercel accounts via the official Vercel CLI. With VAM, you can add, list, remove, and switch between Vercel accounts, making it perfect for developers juggling multiple accounts.

## Features

- **Add new accounts**
- **List** all the accounts you've added
- **Remove** accounts when no longer needed
- **Switch** between accounts instantly

## Installation

### Option 1: Download Pre-built Binaries

Pre-built binaries for Linux, macOS, and Windows are available on the [Releases page](https://github.com/owbird/vercel-account-manager/releases).

### Option 2: Build from Source

If you prefer to build the tool from source, follow these steps:

```bash
git clone https://github.com/owbird/vercel-account-manager.git
cd vercel-account-manager
go build -o vam
```

## Usage

Here’s how you can use the Vercel Account Manager (VAM):

### Add a new account

To add a new Vercel account, use the `create` command.

```bash
vam create [args]
```

The args will be forwarded to the vercel CLI for login.

### List all accounts

To see all the Vercel accounts you’ve added, use the `ls` command.

```bash
vam ls
```

### Remove an account

To remove an account, use the `rm` command followed by the account name.

```bash
vam rm <account_name>
```

### Switch accounts

To switch between accounts, use the `checkout` command followed by the account name.

```bash
vam checkout <account_name>
```

## Commands Overview

| Command       | Description                           |
|---------------|---------------------------------------|
| `vam create`  | Add a new Vercel account              |
| `vam ls`      | List all added accounts               |
| `vam rm`      | Remove a specific account             |
| `vam checkout`| Switch to a different Vercel account  |

## Contributing

Feel free to open issues or submit pull requests!
