# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `PENDO_TRACKING_STATE`: Enable or Disable User Analytics Tracking. (choice)
    - `off`: Off
    - `anonymous`: Anonymous
    - `detailed`: Detailed
* `CUSTOM_LOGIN_INFO`: If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported. (string)
* `CUSTOM_LOGO`: To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported. (string)
* `MAX_UI_JOB_EVENTS`: Maximum number of job events for the UI to retrieve within a single request. (integer)
* `UI_LIVE_UPDATES_ENABLED`: If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details. (boolean)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:



* `CUSTOM_LOGIN_INFO`: If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported. (string, default=`""`)
* `CUSTOM_LOGO`: To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported. (string, default=`""`)
* `MAX_UI_JOB_EVENTS`: Maximum number of job events for the UI to retrieve within a single request. (integer, required)
* `UI_LIVE_UPDATES_ENABLED`: If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details. (boolean, required)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.