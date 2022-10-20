# Bookings and Reservations

This is the repository for bookings project

This project use:

 - Build in Go version 1.19;
 - Use the [Bootstrap](https://getbootstrap.com/), version 5.2.2 CSS;
 - Use the [alex edwards SCS](github.com/alexedwards/scs/v2) version 2.5.0 session manager;
 - Use the [chi router](github.com/go-chi/chi) version 1.5.4;
 - Use [nosurf](github.com/justinas/nosurf) version 1.1.1.
 - Add [Vanilla JS Datepicker](github.com/mymth/vanillajs-datepicker)
 - Add [notie](https://github.com/jaredreich/notie)
 - Add [sweetalert2](https://sweetalert2.github.io/)


Commit 2022/10/20 - Changes in this commit:
- cmd/web/main.go:
    - Set session.Cookie.secure to app.InProduction. 

- cmd/web/route.go:
    - added generals-quarters, majors-suite, search-availability, contact, make reservation handlers to mux;
    - added /static fileServer to mux handler to find files in /static directory.

- pkg/handlers/handlers.go:
    - changed variable name req by request;
    - added methods to Repository class:
        - Generals;
        - Reservations;
        - Majors;
        - Contact;
        - Availability.

- pkg/render/render.go:
    - changed files extensions in ./templetes from .html to .tmpl;
    - added functions var in line 72 to create a new template.

- added files in /static
    - added styles.css and images in /static/css and /static/images

- /templetes directory:
    - added news pages templates:
        - contact.page.tmpl;
        - generals.layout.tmpl;
        - majors.layout.tmpl;
        - make-reservation.layout.tmpl;
        - search-availability.layout.tmpl.
    - changes in other templates and replace extensions from .html to .tmpl.