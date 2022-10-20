Changes in this commit:
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