# Center Creator

Hands On Assignment

Create a Web Application for Center registration and Center profile management, and also write required server-side code:

·         New Center registration page:
o    This page will display the Center registration form to submit centers information, the required information to collect:
Center Name
Owner Name
Owner Mobile Name
Owner Email Id
Learning Facilitator Name (LF)
LF Email ID
LF Mobile No
City
Pin Code
Address
Password
And Confirm Password(Password length minimum 8 characters)

And while Saving the details Generate 8 Digit Unique Center Code as Login ID and Display it to user as his Center Code 
 
o    Your application shall store the data in a database (MySQL or MongoDB) after proper validation, also you have to store the date-time on which the registration was done.  
o    You should not allow duplicate values for Mobile Number, Email ID, and Center Code
o    Once one Center is registered, display appropriate information to the user and take the user to the login page.
 
·         Login Page:
o    User will enter CenterCode and password and you should validate the login information in the database and on a successful match you should take the user to the user’s dashboard page.
·         Centers Dashboard:
o    The user should be able to view his profile information.
o    The user will have options to update the following personal information:

Owner Name
Owner Mobile Name
Owner Email Id
Learning Facilitator Name
LF Email ID
LF Mobile No
City
Pin Code
Address

o    There should be an option to change the password
o    There should be an option to logout from the user’s dashboard and terminate the session.

If time permits you can give provision to add multiple LF

# mini

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
