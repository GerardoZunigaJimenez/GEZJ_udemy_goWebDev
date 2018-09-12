Enctype
“When you make a POST request, you have to encode the data that forms the body of the request in some way. HTML forms provide three methods of encoding:

application/x-www-form-urlencoded (the default) 
multipart/form-data 
text/plain 

Work was being done on adding application/json, but that has been abandoned. The specifics of the formats don't matter to most developers. The important points are: When you are writing client-side code, all you need to know is use multipart/form-data when your form includes any <input type="file"> elements. 

text/plain Never use text/plain. If you are writing (or debugging) a library for parsing or generating the raw data, then you need to start worrying about the format. You might also want to know about it for interest's sake. text/plain is introduced by HTML 5 and is useful only for debugging — from the spec: They are not reliably interpretable by computer — and I'd argue that the others combined with tools (like the Net tab in the developer tools of most browsers) are better for that).


application/x-www-form-urlencoded is more or less the same as a query string on the end of the URL. 

multipart/form-data is significantly more complicated but it allows entire files to be included in the data.”
http://stackoverflow.com/questions/4526273/what-does-enctype-multipart-form-data-mean 
https://www.w3.org/TR/html5/forms.html#text/plain-encoding-algorithm
https://www.ietf.org/rfc/rfc1867.txt
https://www.ietf.org/rfc/rfc2388.txt


