# Render Integration Setup Guide for opencomply

This guide provides instructions to integrate your Render account with opencomply by creating a personal access token for API authorization.

## Steps

### 1. Create a Personal Access Token in Render

#### Log In

- Sign in to the Render Cloud Manager.

#### Access API Tokens

- Click your username at the top of the screen and select **API Tokens**.

#### Create a Personal Access Token

- Click **Create a Personal Access Token**.

#### Configure the Token

- **Label**: Enter a label for the token to identify its intended use, like **opencomply Integration**.
- **Expiry**: Select an appropriate expiration time for the token.
- **Permissions**:
  - For each product or service, select **ReadOnly** access.
  - For VPCs, select **Read/Write** (as Render does not offer Read-Only access for VPCs).

#### Generate and Save Token

- After configuration, ensure the token is saved securely as it will only be displayed once.

### 2. Configure opencomply with the Token

#### Access opencomply

- Log in to the opencomply portal using your admin credentials.

#### Add Integration

- Navigate to **Integrations**, select **Render**, and click **Add New Integration**.

#### Enter API Token

- Paste the personal access token you generated from Render.

#### Complete Integration

- Click **Next**, review the integration details, and then **Confirm** to establish the connection.

By following these steps, you have successfully integrated your Render account with opencomply, allowing read access and necessary write permissions for VPCs to enhance governance and compliance within the platform.
```