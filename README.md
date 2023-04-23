Frontend Sign in Screen and Enable Google Third Party Login

check figma link：https://www.figma.com/file/Kao0wcmhaatk4eQUXXd4ik/CleanLabor-Signup-and-Signin?node-id=1-440&t=uCcvSH02CyEckdZz-0
have the front end page developed
Enable Google Third Party Login
submit a PR request to main branch
备注

Register your application with the Google Developers Console and obtain a Client ID and Client Secret.
Configure the OAuth 2.0 consent screen to display the required information about your application and its usage.
In your application, redirect the user to the Google OAuth 2.0 authorization endpoint to initiate the login flow.
The user will be prompted to grant permission to your application to access their Google account information.
Google will redirect the user back to your application with an authorization code.
Exchange the authorization code for an access token and a refresh token.
Use the access token to make API requests on behalf of the user.
