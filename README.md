## Installation

![Go](https://github.com/RossMerr/jsonschema/workflows/Go/badge.svg)

* `go get -u github.com/RossMerr/jsonschema/cmd/jsonschema`



Basic support to generate go struct's, featuring:
##### $ref 
  * Can resolve to any point in the document or another local file using a JSON Pointer.
  
  `"#/definitions/address"`
  
  `"definitions.json#/address"`
  
  * External files are not supported!
  
  `"http://www.sample.com/definitions.json#/address"`
  
##### $id
  * Top level $id must follow the absolute-URI convention. The old 'id' field is not supported.
  
##### $def
   * The old 'definitions' field is still supported but merged into the new $def field.
   
##### allOf
  * Will generate a struct with all of the subschemas embedded.
   
##### anyOf
  * Will generate an array of a interface with all subschemas implementing its method.
   
##### oneOf
  * Will generate a interface with all subschemas implementing its method.
  
#### Validation

All generated go types will be a value or reference type depending on the 'required' field.
All fields and struct's also get generated validation tags mostly conforming to the Well-known struct tags for [validate](https://github.com/go-playground/validator).
Additional values for the validate tag include:
* allof
* anyof
* oneof
* regex

Which you can then use for any custom validators.

#### Not Supported

##### Not
  * No support we be provided for the 'not' keyword 

##### String-encoding 
  * 'contentMediaType' and 'contentEncoding' will not be supported
 
##### External JSON Pointer

##### Conditionally
  * If, Then and Else are not supported 
  
#### Coming soon

* Enum support for strings
* Validation feedback before parsing 