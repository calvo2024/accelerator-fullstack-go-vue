Serves static content from "static" on "/". This static content is built into the resulting binary.
Serves go function API on "/api/".
The API endpoint sets some sample Cookie. The intent is to use them to authenticate requests.
https://web.dev/articles/introduction-to-fetch
I haven't tested this yet. I still need to learn about security best practices (e.g. anti-CSRF).
