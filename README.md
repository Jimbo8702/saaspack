folder structure for saas in golang fiber templates:

root:
/data - all data related files (types, stores)
/handlers - all api handlers
/lib - internal libraries made for the software (could be called pkg)
/logger - global logging
/settings - saas settings
/static - static files to serve
/sup - supabase config
/util - utility functions for the app
/views - html templates to render
main.go - application entry point

what do i want to do?

- I want to make my ability to build Saas fast
- I can use svelte or django templates
- tempaltes i know already, and can handle simple saas
- svelte is a better option tho, as its a separate frontend pacakge that is still fast

- i need basic auth, login, logout, signin, signout --> i want to use thirdy party auth. Auth0, supabase, firebaseAuth
- i need a basic logger, logger that can take in the kind of log, and produce it to files or just command line
- because im using third party auth, i need a user account that is separate from the auth data, so a Account struct with a UserID that poitns to the thirdparty user
- using fiber => basic framework for easy integration
- crud for each data object needed in the software
- make file for easy building
- error handling for direct responses

- lets integrate auth0 in go as an authHandler
- or we use supabase auth ...

- just build a json api for the saas, have auth routes (sign up, sign in, sign out, delete account, upgrade membership)
- basic payment and stripe integration, standard three payment tiers

- lets build a renting software first

# The Idea

- a business can upload items its offering to rent
- manage items (update price, availibity)
- manage sales
- manage customers
- a business logs in -> create a offer item -> make that item availible in your store -> user rents the item -> keep track of user dues, paid and unpaid
  --> all information about payment, invoice, and return dates.
- cloud document store to keep track of customer documents, members, item documents, aggreements

--lets niche this up

- car subscription service
- upload cars to offer
- have cars load up on front end store
- see which cars are rented, not rented, in repair shop.
- track of payments by customer

-- start with crud api for cars

- create, update, delete, and read cars
  -- create, update, delete, admin only
- subcription obj
- user & car. Stripe invoice, next payment, last payment, started on, updated at.

-- user crud, create read update delete users admin only, tho sign up wil be through a thirdparty, on success, we have a account created if it does not exists in our db

-- auth handler, sign in, sign out, sign up with email password, signup with google => create an account for user

-- need an admin dashboard
-- need a userfacing dashboard
# saaspack
