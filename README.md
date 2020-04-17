![Go](https://github.com/RossMerr/jsonschema/workflows/Go/badge.svg)

# jsonschema

Code generation from json schemas with basic support for 2019-09:


## Installation

* `go install github.com/RossMerr/jsonschema/cmd/jsonschema`


### Notes for usage

##### $id
  * Top level $id must follow the absolute-URI convention of being a canonical URL. The old 'id' field is not supported.
  
##### $ref 
  * Can resolve to any point in the document or another local file using a JSON Pointer.
  
  `"#/definitions/address"`
    
  
  `"http://www.sample.com/definitions.json#/address"`

##### $def
   * The old 'definitions' field is still supported but merged into the new $def field.
   
##### allOf
  * Will generate a struct with all of the subschemas embedded.
   
##### anyOf
  * Will generate an array of a interface with all subschemas implementing its method.
   
##### oneOf
  * Will generate a interface with all subschemas implementing its method.
  
### Validation

All generated go types will be a value or reference type depending on the 'required' field.
All fields and struct's also get generated validation tags mostly conforming to the Well-known struct tags for [validate](https://github.com/go-playground/validator).
Additional values for the validate tag include:
* allof
* anyof
* oneof
* regex

Which you can then use for any custom validators.

## Not or partial support

##### String
   
  * [Format](https://json-schema.org/understanding-json-schema/reference/string.html#format)
  
##### Regular Expressions

You can add support with a custom [validate](https://github.com/go-playground/validator) field tag. But
[a regex validator won't be added because commas and = signs can be part of a regex which conflict with the validation definitions](https://godoc.org/gopkg.in/go-playground/validator.v9#hdr-Alias_Validators_and_Tags).
      
##### Numeric

  * [Multiples](https://json-schema.org/understanding-json-schema/reference/numeric.html#multiples)
      
##### Object
 
  * [Property names](https://json-schema.org/understanding-json-schema/reference/object.html#property-names)
  * [Size](https://json-schema.org/understanding-json-schema/reference/object.html#size)
  * [Dependencies](https://json-schema.org/understanding-json-schema/reference/object.html#dependencies)
  * [Pattern Properties](https://json-schema.org/understanding-json-schema/reference/object.html#pattern-properties)
  
##### Array

  * [List validation - contains](https://json-schema.org/understanding-json-schema/reference/array.html#list-validation)
  * [Tuple validation](https://json-schema.org/understanding-json-schema/reference/array.html#tuple-validation)
  * [Length](https://json-schema.org/understanding-json-schema/reference/array.html#length)
  * [Uniqueness](https://json-schema.org/understanding-json-schema/reference/array.html#uniqueness)
  
##### Generic

  * [Enumerated values](https://json-schema.org/understanding-json-schema/reference/generic.html#enumerated-values)
  Require a type of string  `"type": "string"`
  * [Constant values](https://json-schema.org/understanding-json-schema/reference/generic.html#constant-values)
  
##### [Media: string-encoding non-JSON data](https://json-schema.org/understanding-json-schema/reference/non_json_data.html)
  
##### Combining schemas
  * [Not](https://json-schema.org/understanding-json-schema/reference/combining.html#not)  
 
##### [Applying subschemas conditionally](https://json-schema.org/understanding-json-schema/reference/conditionals.html) 

##### External JSON Pointer
All schemas must be local

 